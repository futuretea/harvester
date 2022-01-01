package fake

import (
	v1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/k8s.cni.cncf.io/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeK8sCniCncfIoV1 struct {
	*testing.Fake
}

func (c *FakeK8sCniCncfIoV1) NetworkAttachmentDefinitions(namespace string) v1.NetworkAttachmentDefinitionInterface {
	__traceStack()

	return &FakeNetworkAttachmentDefinitions{c, namespace}
}

func (c *FakeK8sCniCncfIoV1) RESTClient() rest.Interface {
	__traceStack()

	var ret *rest.RESTClient
	return ret
}
