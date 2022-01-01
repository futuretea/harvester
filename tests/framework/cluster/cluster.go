package cluster

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"path"

	"k8s.io/client-go/tools/clientcmd"

	"github.com/harvester/harvester/tests/framework/env"
)

const (
	KindClusterKind		= "kind"
	ExistingClusterKind	= "unknown"
)

type Cluster interface {
	fmt.Stringer

	GetKind() string

	Startup(output io.Writer) error

	LoadImages(output io.Writer) error

	Cleanup(output io.Writer) error
}

func NewLocalCluster() Cluster {
	__traceStack()

	return NewLocalKindCluster()
}

func GetExistCluster() Cluster {
	__traceStack()

	return &ExistingCluster{}
}

func Start(output io.Writer) (clientcmd.ClientConfig, Cluster, error) {
	__traceStack()

	var cluster Cluster
	if env.IsUsingExistingCluster() {
		cluster = GetExistCluster()
	} else {
		cluster = NewLocalCluster()
	}
	err := cluster.Startup(output)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to startup the local cluster %s, %v", cluster, err)
	}

	err = cluster.LoadImages(output)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load images into cluster %s, %v", cluster, err)
	}

	KubeClientConfig, err := GetConfig()
	return KubeClientConfig, cluster, err
}

func Stop(output io.Writer) error {
	__traceStack()

	if env.IsUsingExistingCluster() || env.IsKeepingTestingCluster() {
		return nil
	}
	cluster := NewLocalCluster()
	err := cluster.Cleanup(output)
	if err != nil {
		return fmt.Errorf("failed to cleanup the local '%s' cluster, %v", cluster.GetKind(), err)
	}
	return nil
}

func GetConfig() (clientcmd.ClientConfig, error) {
	__traceStack()

	var loadingRules = clientcmd.NewDefaultClientConfigLoadingRules()
	if _, ok := os.LookupEnv("HOME"); !ok {
		var u, err = user.Current()
		if err != nil {
			return nil, fmt.Errorf("could not load config from current user as the current user is not found, %v", err)
		}
		loadingRules.Precedence = append(loadingRules.Precedence,
			path.Join(u.HomeDir, clientcmd.RecommendedHomeDir, clientcmd.RecommendedFileName))
	}

	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, &clientcmd.ConfigOverrides{}), nil
}
