package v1beta1

import (
	"github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v1beta1 "github.com/kubernetes-csi/external-snapshotter/v2/pkg/apis/volumesnapshot/v1beta1"
	rest "k8s.io/client-go/rest"
)

type SnapshotV1beta1Interface interface {
	RESTClient() rest.Interface
	VolumeSnapshotsGetter
	VolumeSnapshotClassesGetter
	VolumeSnapshotContentsGetter
}

type SnapshotV1beta1Client struct {
	restClient rest.Interface
}

func (c *SnapshotV1beta1Client) VolumeSnapshots(namespace string) VolumeSnapshotInterface {
	__traceStack()

	return newVolumeSnapshots(c, namespace)
}

func (c *SnapshotV1beta1Client) VolumeSnapshotClasses() VolumeSnapshotClassInterface {
	__traceStack()

	return newVolumeSnapshotClasses(c)
}

func (c *SnapshotV1beta1Client) VolumeSnapshotContents() VolumeSnapshotContentInterface {
	__traceStack()

	return newVolumeSnapshotContents(c)
}

func NewForConfig(c *rest.Config) (*SnapshotV1beta1Client, error) {
	__traceStack()

	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &SnapshotV1beta1Client{client}, nil
}

func NewForConfigOrDie(c *rest.Config) *SnapshotV1beta1Client {
	__traceStack()

	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

func New(c rest.Interface) *SnapshotV1beta1Client {
	__traceStack()

	return &SnapshotV1beta1Client{c}
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

func (c *SnapshotV1beta1Client) RESTClient() rest.Interface {
	__traceStack()

	if c == nil {
		return nil
	}
	return c.restClient
}
