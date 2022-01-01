package fake

import (
	v1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/kubevirt.io/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeKubevirtV1 struct {
	*testing.Fake
}

func (c *FakeKubevirtV1) KubeVirts(namespace string) v1.KubeVirtInterface {
	__traceStack()

	return &FakeKubeVirts{c, namespace}
}

func (c *FakeKubevirtV1) VirtualMachines(namespace string) v1.VirtualMachineInterface {
	__traceStack()

	return &FakeVirtualMachines{c, namespace}
}

func (c *FakeKubevirtV1) VirtualMachineInstances(namespace string) v1.VirtualMachineInstanceInterface {
	__traceStack()

	return &FakeVirtualMachineInstances{c, namespace}
}

func (c *FakeKubevirtV1) VirtualMachineInstanceMigrations(namespace string) v1.VirtualMachineInstanceMigrationInterface {
	__traceStack()

	return &FakeVirtualMachineInstanceMigrations{c, namespace}
}

func (c *FakeKubevirtV1) VirtualMachineInstancePresets(namespace string) v1.VirtualMachineInstancePresetInterface {
	__traceStack()

	return &FakeVirtualMachineInstancePresets{c, namespace}
}

func (c *FakeKubevirtV1) VirtualMachineInstanceReplicaSets(namespace string) v1.VirtualMachineInstanceReplicaSetInterface {
	__traceStack()

	return &FakeVirtualMachineInstanceReplicaSets{c, namespace}
}

func (c *FakeKubevirtV1) RESTClient() rest.Interface {
	__traceStack()

	var ret *rest.RESTClient
	return ret
}
