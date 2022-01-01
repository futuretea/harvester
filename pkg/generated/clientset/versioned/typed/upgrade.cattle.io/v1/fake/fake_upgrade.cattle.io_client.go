package fake

import (
	v1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/upgrade.cattle.io/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeUpgradeV1 struct {
	*testing.Fake
}

func (c *FakeUpgradeV1) Plans(namespace string) v1.PlanInterface {
	__traceStack()

	return &FakePlans{c, namespace}
}

func (c *FakeUpgradeV1) RESTClient() rest.Interface {
	__traceStack()

	var ret *rest.RESTClient
	return ret
}
