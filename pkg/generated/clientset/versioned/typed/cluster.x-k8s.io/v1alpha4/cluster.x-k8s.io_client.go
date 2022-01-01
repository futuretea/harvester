package v1alpha4

import (
	"github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
	v1alpha4 "sigs.k8s.io/cluster-api/api/v1alpha4"
)

type ClusterV1alpha4Interface interface {
	RESTClient() rest.Interface
	MachinesGetter
}

type ClusterV1alpha4Client struct {
	restClient rest.Interface
}

func (c *ClusterV1alpha4Client) Machines(namespace string) MachineInterface {
	__traceStack()

	return newMachines(c, namespace)
}

func NewForConfig(c *rest.Config) (*ClusterV1alpha4Client, error) {
	__traceStack()

	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ClusterV1alpha4Client{client}, nil
}

func NewForConfigOrDie(c *rest.Config) *ClusterV1alpha4Client {
	__traceStack()

	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

func New(c rest.Interface) *ClusterV1alpha4Client {
	__traceStack()

	return &ClusterV1alpha4Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	__traceStack()

	gv := v1alpha4.GroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

func (c *ClusterV1alpha4Client) RESTClient() rest.Interface {
	__traceStack()

	if c == nil {
		return nil
	}
	return c.restClient
}
