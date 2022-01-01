package fake

import (
	v1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/networking.k8s.io/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeNetworkingV1 struct {
	*testing.Fake
}

func (c *FakeNetworkingV1) Ingresses(namespace string) v1.IngressInterface {
	__traceStack()

	return &FakeIngresses{c, namespace}
}

func (c *FakeNetworkingV1) IngressClasses() v1.IngressClassInterface {
	__traceStack()

	return &FakeIngressClasses{c}
}

func (c *FakeNetworkingV1) NetworkPolicies(namespace string) v1.NetworkPolicyInterface {
	__traceStack()

	return &FakeNetworkPolicies{c, namespace}
}

func (c *FakeNetworkingV1) RESTClient() rest.Interface {
	__traceStack()

	var ret *rest.RESTClient
	return ret
}
