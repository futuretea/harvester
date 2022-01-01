package v1

import (
	"github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v1 "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
	rest "k8s.io/client-go/rest"
)

type K8sCniCncfIoV1Interface interface {
	RESTClient() rest.Interface
	NetworkAttachmentDefinitionsGetter
}

type K8sCniCncfIoV1Client struct {
	restClient rest.Interface
}

func (c *K8sCniCncfIoV1Client) NetworkAttachmentDefinitions(namespace string) NetworkAttachmentDefinitionInterface {
	__traceStack()

	return newNetworkAttachmentDefinitions(c, namespace)
}

func NewForConfig(c *rest.Config) (*K8sCniCncfIoV1Client, error) {
	__traceStack()

	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &K8sCniCncfIoV1Client{client}, nil
}

func NewForConfigOrDie(c *rest.Config) *K8sCniCncfIoV1Client {
	__traceStack()

	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

func New(c rest.Interface) *K8sCniCncfIoV1Client {
	__traceStack()

	return &K8sCniCncfIoV1Client{c}
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

func (c *K8sCniCncfIoV1Client) RESTClient() rest.Interface {
	__traceStack()

	if c == nil {
		return nil
	}
	return c.restClient
}
