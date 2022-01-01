package v1

import (
	"github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
	v1 "kubevirt.io/client-go/api/v1"
)

type KubevirtV1Interface interface {
	RESTClient() rest.Interface
	KubeVirtsGetter
	VirtualMachinesGetter
	VirtualMachineInstancesGetter
	VirtualMachineInstanceMigrationsGetter
	VirtualMachineInstancePresetsGetter
	VirtualMachineInstanceReplicaSetsGetter
}

type KubevirtV1Client struct {
	restClient rest.Interface
}

func (c *KubevirtV1Client) KubeVirts(namespace string) KubeVirtInterface {
	__traceStack()

	return newKubeVirts(c, namespace)
}

func (c *KubevirtV1Client) VirtualMachines(namespace string) VirtualMachineInterface {
	__traceStack()

	return newVirtualMachines(c, namespace)
}

func (c *KubevirtV1Client) VirtualMachineInstances(namespace string) VirtualMachineInstanceInterface {
	__traceStack()

	return newVirtualMachineInstances(c, namespace)
}

func (c *KubevirtV1Client) VirtualMachineInstanceMigrations(namespace string) VirtualMachineInstanceMigrationInterface {
	__traceStack()

	return newVirtualMachineInstanceMigrations(c, namespace)
}

func (c *KubevirtV1Client) VirtualMachineInstancePresets(namespace string) VirtualMachineInstancePresetInterface {
	__traceStack()

	return newVirtualMachineInstancePresets(c, namespace)
}

func (c *KubevirtV1Client) VirtualMachineInstanceReplicaSets(namespace string) VirtualMachineInstanceReplicaSetInterface {
	__traceStack()

	return newVirtualMachineInstanceReplicaSets(c, namespace)
}

func NewForConfig(c *rest.Config) (*KubevirtV1Client, error) {
	__traceStack()

	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &KubevirtV1Client{client}, nil
}

func NewForConfigOrDie(c *rest.Config) *KubevirtV1Client {
	__traceStack()

	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

func New(c rest.Interface) *KubevirtV1Client {
	__traceStack()

	return &KubevirtV1Client{c}
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

func (c *KubevirtV1Client) RESTClient() rest.Interface {
	__traceStack()

	if c == nil {
		return nil
	}
	return c.restClient
}
