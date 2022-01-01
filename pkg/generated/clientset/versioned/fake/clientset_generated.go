package fake

import (
	clientset "github.com/harvester/harvester/pkg/generated/clientset/versioned"
	clusterv1alpha4 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/cluster.x-k8s.io/v1alpha4"
	fakeclusterv1alpha4 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/cluster.x-k8s.io/v1alpha4/fake"
	harvesterhciv1beta1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/harvesterhci.io/v1beta1"
	fakeharvesterhciv1beta1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/harvesterhci.io/v1beta1/fake"
	k8scnicncfiov1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/k8s.cni.cncf.io/v1"
	fakek8scnicncfiov1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/k8s.cni.cncf.io/v1/fake"
	kubevirtv1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/kubevirt.io/v1"
	fakekubevirtv1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/kubevirt.io/v1/fake"
	longhornv1beta1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/longhorn.io/v1beta1"
	fakelonghornv1beta1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/longhorn.io/v1beta1/fake"
	networkingv1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/networking.k8s.io/v1"
	fakenetworkingv1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/networking.k8s.io/v1/fake"
	snapshotv1beta1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/snapshot.storage.k8s.io/v1beta1"
	fakesnapshotv1beta1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/snapshot.storage.k8s.io/v1beta1/fake"
	upgradev1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/upgrade.cattle.io/v1"
	fakeupgradev1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/upgrade.cattle.io/v1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	__traceStack()

	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

type Clientset struct {
	testing.Fake
	discovery	*fakediscovery.FakeDiscovery
	tracker		testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	__traceStack()

	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	__traceStack()

	return c.tracker
}

var _ clientset.Interface = &Clientset{}

func (c *Clientset) ClusterV1alpha4() clusterv1alpha4.ClusterV1alpha4Interface {
	__traceStack()

	return &fakeclusterv1alpha4.FakeClusterV1alpha4{Fake: &c.Fake}
}

func (c *Clientset) HarvesterhciV1beta1() harvesterhciv1beta1.HarvesterhciV1beta1Interface {
	__traceStack()

	return &fakeharvesterhciv1beta1.FakeHarvesterhciV1beta1{Fake: &c.Fake}
}

func (c *Clientset) K8sCniCncfIoV1() k8scnicncfiov1.K8sCniCncfIoV1Interface {
	__traceStack()

	return &fakek8scnicncfiov1.FakeK8sCniCncfIoV1{Fake: &c.Fake}
}

func (c *Clientset) KubevirtV1() kubevirtv1.KubevirtV1Interface {
	__traceStack()

	return &fakekubevirtv1.FakeKubevirtV1{Fake: &c.Fake}
}

func (c *Clientset) LonghornV1beta1() longhornv1beta1.LonghornV1beta1Interface {
	__traceStack()

	return &fakelonghornv1beta1.FakeLonghornV1beta1{Fake: &c.Fake}
}

func (c *Clientset) NetworkingV1() networkingv1.NetworkingV1Interface {
	__traceStack()

	return &fakenetworkingv1.FakeNetworkingV1{Fake: &c.Fake}
}

func (c *Clientset) SnapshotV1beta1() snapshotv1beta1.SnapshotV1beta1Interface {
	__traceStack()

	return &fakesnapshotv1beta1.FakeSnapshotV1beta1{Fake: &c.Fake}
}

func (c *Clientset) UpgradeV1() upgradev1.UpgradeV1Interface {
	__traceStack()

	return &fakeupgradev1.FakeUpgradeV1{Fake: &c.Fake}
}
