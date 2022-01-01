package v1

import (
	"github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/api/networking/v1"
	rest "k8s.io/client-go/rest"
)

type NetworkingV1Interface interface {
	RESTClient() rest.Interface
	IngressesGetter
	IngressClassesGetter
	NetworkPoliciesGetter
}

type NetworkingV1Client struct {
	restClient rest.Interface
}

func (c *NetworkingV1Client) Ingresses(namespace string) IngressInterface {
	__traceStack()

	return newIngresses(c, namespace)
}

func (c *NetworkingV1Client) IngressClasses() IngressClassInterface {
	__traceStack()

	return newIngressClasses(c)
}

func (c *NetworkingV1Client) NetworkPolicies(namespace string) NetworkPolicyInterface {
	__traceStack()

	return newNetworkPolicies(c, namespace)
}

func NewForConfig(c *rest.Config) (*NetworkingV1Client, error) {
	__traceStack()

	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &NetworkingV1Client{client}, nil
}

func NewForConfigOrDie(c *rest.Config) *NetworkingV1Client {
	__traceStack()

	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

func New(c rest.Interface) *NetworkingV1Client {
	__traceStack()

	return &NetworkingV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	__traceStack()

	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

func (c *NetworkingV1Client) RESTClient() rest.Interface {
	__traceStack()

	if c == nil {
		return nil
	}
	return c.restClient
}
