package v1beta1

import (
	"github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v1beta1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
	rest "k8s.io/client-go/rest"
)

type LonghornV1beta1Interface interface {
	RESTClient() rest.Interface
	BackingImagesGetter
	BackingImageDataSourcesGetter
	BackingImageManagersGetter
	BackupsGetter
	BackupTargetsGetter
	BackupVolumesGetter
	EnginesGetter
	EngineImagesGetter
	InstanceManagersGetter
	NodesGetter
	RecurringJobsGetter
	ReplicasGetter
	SettingsGetter
	ShareManagersGetter
	VolumesGetter
}

type LonghornV1beta1Client struct {
	restClient rest.Interface
}

func (c *LonghornV1beta1Client) BackingImages(namespace string) BackingImageInterface {
	__traceStack()

	return newBackingImages(c, namespace)
}

func (c *LonghornV1beta1Client) BackingImageDataSources(namespace string) BackingImageDataSourceInterface {
	__traceStack()

	return newBackingImageDataSources(c, namespace)
}

func (c *LonghornV1beta1Client) BackingImageManagers(namespace string) BackingImageManagerInterface {
	__traceStack()

	return newBackingImageManagers(c, namespace)
}

func (c *LonghornV1beta1Client) Backups(namespace string) BackupInterface {
	__traceStack()

	return newBackups(c, namespace)
}

func (c *LonghornV1beta1Client) BackupTargets(namespace string) BackupTargetInterface {
	__traceStack()

	return newBackupTargets(c, namespace)
}

func (c *LonghornV1beta1Client) BackupVolumes(namespace string) BackupVolumeInterface {
	__traceStack()

	return newBackupVolumes(c, namespace)
}

func (c *LonghornV1beta1Client) Engines(namespace string) EngineInterface {
	__traceStack()

	return newEngines(c, namespace)
}

func (c *LonghornV1beta1Client) EngineImages(namespace string) EngineImageInterface {
	__traceStack()

	return newEngineImages(c, namespace)
}

func (c *LonghornV1beta1Client) InstanceManagers(namespace string) InstanceManagerInterface {
	__traceStack()

	return newInstanceManagers(c, namespace)
}

func (c *LonghornV1beta1Client) Nodes(namespace string) NodeInterface {
	__traceStack()

	return newNodes(c, namespace)
}

func (c *LonghornV1beta1Client) RecurringJobs(namespace string) RecurringJobInterface {
	__traceStack()

	return newRecurringJobs(c, namespace)
}

func (c *LonghornV1beta1Client) Replicas(namespace string) ReplicaInterface {
	__traceStack()

	return newReplicas(c, namespace)
}

func (c *LonghornV1beta1Client) Settings(namespace string) SettingInterface {
	__traceStack()

	return newSettings(c, namespace)
}

func (c *LonghornV1beta1Client) ShareManagers(namespace string) ShareManagerInterface {
	__traceStack()

	return newShareManagers(c, namespace)
}

func (c *LonghornV1beta1Client) Volumes(namespace string) VolumeInterface {
	__traceStack()

	return newVolumes(c, namespace)
}

func NewForConfig(c *rest.Config) (*LonghornV1beta1Client, error) {
	__traceStack()

	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &LonghornV1beta1Client{client}, nil
}

func NewForConfigOrDie(c *rest.Config) *LonghornV1beta1Client {
	__traceStack()

	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

func New(c rest.Interface) *LonghornV1beta1Client {
	__traceStack()

	return &LonghornV1beta1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	__traceStack()

	gv := v1beta1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

func (c *LonghornV1beta1Client) RESTClient() rest.Interface {
	__traceStack()

	if c == nil {
		return nil
	}
	return c.restClient
}
