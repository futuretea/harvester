package versioned

import (
	"fmt"

	clusterv1alpha4 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/cluster.x-k8s.io/v1alpha4"
	harvesterhciv1beta1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/harvesterhci.io/v1beta1"
	k8scnicncfiov1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/k8s.cni.cncf.io/v1"
	kubevirtv1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/kubevirt.io/v1"
	longhornv1beta1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/longhorn.io/v1beta1"
	networkingv1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/networking.k8s.io/v1"
	snapshotv1beta1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/snapshot.storage.k8s.io/v1beta1"
	upgradev1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/upgrade.cattle.io/v1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	ClusterV1alpha4() clusterv1alpha4.ClusterV1alpha4Interface
	HarvesterhciV1beta1() harvesterhciv1beta1.HarvesterhciV1beta1Interface
	K8sCniCncfIoV1() k8scnicncfiov1.K8sCniCncfIoV1Interface
	KubevirtV1() kubevirtv1.KubevirtV1Interface
	LonghornV1beta1() longhornv1beta1.LonghornV1beta1Interface
	NetworkingV1() networkingv1.NetworkingV1Interface
	SnapshotV1beta1() snapshotv1beta1.SnapshotV1beta1Interface
	UpgradeV1() upgradev1.UpgradeV1Interface
}

type Clientset struct {
	*discovery.DiscoveryClient
	clusterV1alpha4		*clusterv1alpha4.ClusterV1alpha4Client
	harvesterhciV1beta1	*harvesterhciv1beta1.HarvesterhciV1beta1Client
	k8sCniCncfIoV1		*k8scnicncfiov1.K8sCniCncfIoV1Client
	kubevirtV1		*kubevirtv1.KubevirtV1Client
	longhornV1beta1		*longhornv1beta1.LonghornV1beta1Client
	networkingV1		*networkingv1.NetworkingV1Client
	snapshotV1beta1		*snapshotv1beta1.SnapshotV1beta1Client
	upgradeV1		*upgradev1.UpgradeV1Client
}

func (c *Clientset) ClusterV1alpha4() clusterv1alpha4.ClusterV1alpha4Interface {
	__traceStack()

	return c.clusterV1alpha4
}

func (c *Clientset) HarvesterhciV1beta1() harvesterhciv1beta1.HarvesterhciV1beta1Interface {
	__traceStack()

	return c.harvesterhciV1beta1
}

func (c *Clientset) K8sCniCncfIoV1() k8scnicncfiov1.K8sCniCncfIoV1Interface {
	__traceStack()

	return c.k8sCniCncfIoV1
}

func (c *Clientset) KubevirtV1() kubevirtv1.KubevirtV1Interface {
	__traceStack()

	return c.kubevirtV1
}

func (c *Clientset) LonghornV1beta1() longhornv1beta1.LonghornV1beta1Interface {
	__traceStack()

	return c.longhornV1beta1
}

func (c *Clientset) NetworkingV1() networkingv1.NetworkingV1Interface {
	__traceStack()

	return c.networkingV1
}

func (c *Clientset) SnapshotV1beta1() snapshotv1beta1.SnapshotV1beta1Interface {
	__traceStack()

	return c.snapshotV1beta1
}

func (c *Clientset) UpgradeV1() upgradev1.UpgradeV1Interface {
	__traceStack()

	return c.upgradeV1
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	__traceStack()

	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

func NewForConfig(c *rest.Config) (*Clientset, error) {
	__traceStack()

	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.clusterV1alpha4, err = clusterv1alpha4.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.harvesterhciV1beta1, err = harvesterhciv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.k8sCniCncfIoV1, err = k8scnicncfiov1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.kubevirtV1, err = kubevirtv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.longhornV1beta1, err = longhornv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.networkingV1, err = networkingv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.snapshotV1beta1, err = snapshotv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.upgradeV1, err = upgradev1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

func NewForConfigOrDie(c *rest.Config) *Clientset {
	__traceStack()

	var cs Clientset
	cs.clusterV1alpha4 = clusterv1alpha4.NewForConfigOrDie(c)
	cs.harvesterhciV1beta1 = harvesterhciv1beta1.NewForConfigOrDie(c)
	cs.k8sCniCncfIoV1 = k8scnicncfiov1.NewForConfigOrDie(c)
	cs.kubevirtV1 = kubevirtv1.NewForConfigOrDie(c)
	cs.longhornV1beta1 = longhornv1beta1.NewForConfigOrDie(c)
	cs.networkingV1 = networkingv1.NewForConfigOrDie(c)
	cs.snapshotV1beta1 = snapshotv1beta1.NewForConfigOrDie(c)
	cs.upgradeV1 = upgradev1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

func New(c rest.Interface) *Clientset {
	__traceStack()

	var cs Clientset
	cs.clusterV1alpha4 = clusterv1alpha4.New(c)
	cs.harvesterhciV1beta1 = harvesterhciv1beta1.New(c)
	cs.k8sCniCncfIoV1 = k8scnicncfiov1.New(c)
	cs.kubevirtV1 = kubevirtv1.New(c)
	cs.longhornV1beta1 = longhornv1beta1.New(c)
	cs.networkingV1 = networkingv1.New(c)
	cs.snapshotV1beta1 = snapshotv1beta1.New(c)
	cs.upgradeV1 = upgradev1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
