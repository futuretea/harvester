package fake

import (
	"context"

	v1beta1 "github.com/kubernetes-csi/external-snapshotter/v2/pkg/apis/volumesnapshot/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

type FakeVolumeSnapshots struct {
	Fake	*FakeSnapshotV1beta1
	ns	string
}

var volumesnapshotsResource = schema.GroupVersionResource{Group: "snapshot.storage.k8s.io", Version: "v1beta1", Resource: "volumesnapshots"}

var volumesnapshotsKind = schema.GroupVersionKind{Group: "snapshot.storage.k8s.io", Version: "v1beta1", Kind: "VolumeSnapshot"}

func (c *FakeVolumeSnapshots) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VolumeSnapshot, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(volumesnapshotsResource, c.ns, name), &v1beta1.VolumeSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeSnapshot), err
}

func (c *FakeVolumeSnapshots) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VolumeSnapshotList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(volumesnapshotsResource, volumesnapshotsKind, c.ns, opts), &v1beta1.VolumeSnapshotList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.VolumeSnapshotList{ListMeta: obj.(*v1beta1.VolumeSnapshotList).ListMeta}
	for _, item := range obj.(*v1beta1.VolumeSnapshotList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeVolumeSnapshots) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(volumesnapshotsResource, c.ns, opts))

}

func (c *FakeVolumeSnapshots) Create(ctx context.Context, volumeSnapshot *v1beta1.VolumeSnapshot, opts v1.CreateOptions) (result *v1beta1.VolumeSnapshot, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(volumesnapshotsResource, c.ns, volumeSnapshot), &v1beta1.VolumeSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeSnapshot), err
}

func (c *FakeVolumeSnapshots) Update(ctx context.Context, volumeSnapshot *v1beta1.VolumeSnapshot, opts v1.UpdateOptions) (result *v1beta1.VolumeSnapshot, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(volumesnapshotsResource, c.ns, volumeSnapshot), &v1beta1.VolumeSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeSnapshot), err
}

func (c *FakeVolumeSnapshots) UpdateStatus(ctx context.Context, volumeSnapshot *v1beta1.VolumeSnapshot, opts v1.UpdateOptions) (*v1beta1.VolumeSnapshot, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(volumesnapshotsResource, "status", c.ns, volumeSnapshot), &v1beta1.VolumeSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeSnapshot), err
}

func (c *FakeVolumeSnapshots) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(volumesnapshotsResource, c.ns, name), &v1beta1.VolumeSnapshot{})

	return err
}

func (c *FakeVolumeSnapshots) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(volumesnapshotsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.VolumeSnapshotList{})
	return err
}

func (c *FakeVolumeSnapshots) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VolumeSnapshot, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(volumesnapshotsResource, c.ns, name, pt, data, subresources...), &v1beta1.VolumeSnapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeSnapshot), err
}
