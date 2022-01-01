package fake

import (
	v1alpha4 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/cluster.x-k8s.io/v1alpha4"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeClusterV1alpha4 struct {
	*testing.Fake
}

func (c *FakeClusterV1alpha4) Machines(namespace string) v1alpha4.MachineInterface {
	__traceStack()

	return &FakeMachines{c, namespace}
}

func (c *FakeClusterV1alpha4) RESTClient() rest.Interface {
	__traceStack()

	var ret *rest.RESTClient
	return ret
}
