package fake

import (
	v1beta1 "github.com/harvester/harvester/pkg/generated/clientset/versioned/typed/snapshot.storage.k8s.io/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSnapshotV1beta1 struct {
	*testing.Fake
}

func (c *FakeSnapshotV1beta1) VolumeSnapshots(namespace string) v1beta1.VolumeSnapshotInterface {
	__traceStack()

	return &FakeVolumeSnapshots{c, namespace}
}

func (c *FakeSnapshotV1beta1) VolumeSnapshotClasses() v1beta1.VolumeSnapshotClassInterface {
	__traceStack()

	return &FakeVolumeSnapshotClasses{c}
}

func (c *FakeSnapshotV1beta1) VolumeSnapshotContents() v1beta1.VolumeSnapshotContentInterface {
	__traceStack()

	return &FakeVolumeSnapshotContents{c}
}

func (c *FakeSnapshotV1beta1) RESTClient() rest.Interface {
	__traceStack()

	var ret *rest.RESTClient
	return ret
}
