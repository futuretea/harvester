package v1

import (
	"github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v1 "github.com/rancher/system-upgrade-controller/pkg/apis/upgrade.cattle.io/v1"
	rest "k8s.io/client-go/rest"
)

type UpgradeV1Interface interface {
	RESTClient() rest.Interface
	PlansGetter
}

type UpgradeV1Client struct {
	restClient rest.Interface
}

func (c *UpgradeV1Client) Plans(namespace string) PlanInterface {
	__traceStack()

	return newPlans(c, namespace)
}

func NewForConfig(c *rest.Config) (*UpgradeV1Client, error) {
	__traceStack()

	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &UpgradeV1Client{client}, nil
}

func NewForConfigOrDie(c *rest.Config) *UpgradeV1Client {
	__traceStack()

	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

func New(c rest.Interface) *UpgradeV1Client {
	__traceStack()

	return &UpgradeV1Client{c}
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

func (c *UpgradeV1Client) RESTClient() rest.Interface {
	__traceStack()

	if c == nil {
		return nil
	}
	return c.restClient
}
