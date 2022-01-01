package cluster

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"sigs.k8s.io/kind/pkg/cluster"
	"sigs.k8s.io/kind/pkg/exec"

	"github.com/harvester/harvester/tests/framework/env"
	"github.com/harvester/harvester/tests/framework/finder"
	"github.com/harvester/harvester/tests/framework/fuzz"
	"github.com/harvester/harvester/tests/framework/logs"
)

var _ Cluster = &LocalKindCluster{}

type LocalKindCluster struct {
	ExportIngressHTTPPort	int

	ExportIngressHTTPSPort	int

	ExportImageStoragePort	int

	Image	string

	ImageMirror	string

	ClusterName	string

	ControlPlanes	int

	Workers	int

	WaitTimeout	time.Duration

	ClusterConfigPath	string
}

const (
	DefaultControlPlanes	= 1
	DefaultWorkers		= 3
)

func (c *LocalKindCluster) Startup(output io.Writer) error {
	__traceStack()

	logger := logs.NewLogger(output, 0)
	provider := cluster.NewProvider(
		cluster.ProviderWithLogger(logger),
	)

	existed, err := isClusterExisted(provider, c.ClusterName)
	if err != nil {
		return err
	}

	if existed {
		err = provider.Delete(c.ClusterName, "")
		if err != nil {
			return fmt.Errorf("failed to clean the previous cluster, %v", err)
		}
	}

	var configOption cluster.CreateOption
	if c.ClusterConfigPath == "" {
		err = c.initExportPorts()
		if err != nil {
			return err
		}
		config, err := c.generateConfiguration()
		if err != nil {
			return err
		}
		logger.V(0).Info(string(config))
		configOption = cluster.CreateWithRawConfig(config)
	} else {
		configPath, err := filepath.Abs(c.ClusterConfigPath)
		if err != nil {
			return fmt.Errorf("failed to load cluster config from path %s, %v", c.ClusterConfigPath, err)
		}
		configOption = cluster.CreateWithConfigFile(configPath)
	}

	err = provider.Create(
		c.ClusterName,
		configOption,
		cluster.CreateWithNodeImage(c.Image),
		cluster.CreateWithWaitForReady(c.WaitTimeout),
	)
	if err != nil {
		return fmt.Errorf("failed to startup, %v", err)
	}

	return nil
}

func (c LocalKindCluster) LoadImages(output io.Writer) error {
	__traceStack()

	logger := logs.NewLogger(output, 0)

	for _, image := range env.GetPreloadingImages() {
		logger.V(0).Infof("Loading image %s...", image)
		cmd := exec.Command("kind", "load", "docker-image", image, "--name", c.ClusterName)

		lines, err := exec.CombinedOutputLines(cmd)
		if err != nil {
			return err
		}
		logger.V(0).Info(strings.Join(lines, "\n"))
	}

	return nil
}

func (c LocalKindCluster) Cleanup(output io.Writer) error {
	__traceStack()

	var logger = logs.NewLogger(output, 0)
	var provider = cluster.NewProvider(
		cluster.ProviderWithLogger(logger),
	)

	var existed, err = isClusterExisted(provider, c.ClusterName)
	if err != nil {
		return err
	}

	if !existed {
		return nil
	}

	err = provider.Delete(c.ClusterName, "")
	if err != nil {
		return fmt.Errorf("failed to clean the local test cluster, %v", err)
	}
	return nil
}

func (c LocalKindCluster) GetKind() string {
	__traceStack()

	return KindClusterKind
}

func (c LocalKindCluster) String() string {
	__traceStack()

	return fmt.Sprintf("Name: %s, Kind: %s, Image: %s", c.ClusterName, c.GetKind(), c.Image)
}

func (c *LocalKindCluster) initExportPorts() error {
	__traceStack()

	if c.ExportIngressHTTPSPort == 0 || c.ExportIngressHTTPPort == 0 || c.ExportImageStoragePort == 0 {
		var ports, err = fuzz.FreePorts(3)
		if err != nil {
			return fmt.Errorf("failed to generate free ports in local, %v", err)
		}
		c.ExportIngressHTTPPort = ports[0]
		c.ExportIngressHTTPSPort = ports[1]
		c.ExportImageStoragePort = ports[2]
	}
	return nil
}

func (c LocalKindCluster) generateConfiguration() ([]byte, error) {
	__traceStack()

	var tpText = `---
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
{{- if .ImageMirror }}
containerdConfigPatches:
- |-
  [plugins."io.containerd.grpc.v1.cri".registry.mirrors."docker.io"]
    endpoint = ["{{.ImageMirror}}"]
{{- end }}
networking:
  apiServerAddress: "0.0.0.0"
nodes:
  - role: control-plane
    kubeadmConfigPatches:
    - |
      kind: InitConfiguration
      nodeRegistration:
        kubeletExtraArgs:
          node-labels: "ingress-ready=true"
    extraPortMappings:
    - containerPort: 80
      hostPort: {{ .ExportIngressHTTPPort }}
      protocol: TCP
    - containerPort: 443
      hostPort: {{ .ExportIngressHTTPSPort }}
      protocol: TCP
    - containerPort: 32000
      hostPort: {{ .ExportImageStoragePort }}
      protocol: TCP
{{- range (intRange .ControlPlanes) }}
  - role: control-plane
{{- end }}
{{- range (intRange .Workers) }}
  - role: worker
{{- end }}
---
`
	tpFuncMap := template.FuncMap{
		"intRange": func(size int) []int {
			return make([]int, size)
		},
	}
	tp, err := template.New("harvester").Funcs(tpFuncMap).Parse(tpText)
	if err != nil {
		return nil, fmt.Errorf("failed to parse configuration template, %v", err)
	}
	cp := c
	cp.ControlPlanes--
	var output bytes.Buffer
	err = tp.Execute(&output, cp)
	if err != nil {
		return nil, fmt.Errorf("failed to generate configuration, %v", err)
	}
	return output.Bytes(), nil
}

func isClusterExisted(provider *cluster.Provider, clusterName string) (bool, error) {
	__traceStack()

	var clusters, err = provider.List()
	if err != nil {
		return false, fmt.Errorf("failed to list all local clusters, %v", err)
	}

	for _, cls := range clusters {
		if cls == clusterName {
			return true, nil
		}
	}
	return false, nil
}

var (
	localKindClusterOnce	sync.Once
	localKindCluster	*LocalKindCluster
)

func NewLocalKindCluster() *LocalKindCluster {
	__traceStack()

	localKindClusterOnce.Do(func() {
		envFinder := finder.NewEnvFinder("kind")
		localKindCluster = &LocalKindCluster{
			ExportIngressHTTPPort:	envFinder.GetInt("exportIngressHttpPort", 0),
			ExportIngressHTTPSPort:	envFinder.GetInt("exportIngressHttpsPort", 0),
			ExportImageStoragePort:	envFinder.GetInt("exportImageStoragePort", 0),
			Image:			envFinder.Get("image", "kindest/node:v1.21.1"),
			ImageMirror:		envFinder.Get("imageMirror", ""),
			ClusterName:		envFinder.Get("clusterName", "harvester"),
			ControlPlanes:		envFinder.GetInt("controlPlanes", DefaultControlPlanes),
			Workers:		envFinder.GetInt("workers", DefaultWorkers),
			WaitTimeout:		envFinder.GetDuration("waitTimeout", 10*time.Minute),
			ClusterConfigPath:	envFinder.Get("clusterConfigPath", ""),
		}
	})
	return localKindCluster
}
