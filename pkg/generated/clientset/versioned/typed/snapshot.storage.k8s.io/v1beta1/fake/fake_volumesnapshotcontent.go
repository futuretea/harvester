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

type FakeVolumeSnapshotContents struct {
	Fake *FakeSnapshotV1beta1
}

var volumesnapshotcontentsResource = schema.GroupVersionResource{Group: "snapshot.storage.k8s.io", Version: "v1beta1", Resource: "volumesnapshotcontents"}

var volumesnapshotcontentsKind = schema.GroupVersionKind{Group: "snapshot.storage.k8s.io", Version: "v1beta1", Kind: "VolumeSnapshotContent"}

func (c *FakeVolumeSnapshotContents) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VolumeSnapshotContent, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(volumesnapshotcontentsResource, name), &v1beta1.VolumeSnapshotContent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeSnapshotContent), err
}

func (c *FakeVolumeSnapshotContents) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VolumeSnapshotContentList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(volumesnapshotcontentsResource, volumesnapshotcontentsKind, opts), &v1beta1.VolumeSnapshotContentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.VolumeSnapshotContentList{ListMeta: obj.(*v1beta1.VolumeSnapshotContentList).ListMeta}
	for _, item := range obj.(*v1beta1.VolumeSnapshotContentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeVolumeSnapshotContents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(volumesnapshotcontentsResource, opts))
}

func (c *FakeVolumeSnapshotContents) Create(ctx context.Context, volumeSnapshotContent *v1beta1.VolumeSnapshotContent, opts v1.CreateOptions) (result *v1beta1.VolumeSnapshotContent, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(volumesnapshotcontentsResource, volumeSnapshotContent), &v1beta1.VolumeSnapshotContent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeSnapshotContent), err
}

func (c *FakeVolumeSnapshotContents) Update(ctx context.Context, volumeSnapshotContent *v1beta1.VolumeSnapshotContent, opts v1.UpdateOptions) (result *v1beta1.VolumeSnapshotContent, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(volumesnapshotcontentsResource, volumeSnapshotContent), &v1beta1.VolumeSnapshotContent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeSnapshotContent), err
}

func (c *FakeVolumeSnapshotContents) UpdateStatus(ctx context.Context, volumeSnapshotContent *v1beta1.VolumeSnapshotContent, opts v1.UpdateOptions) (*v1beta1.VolumeSnapshotContent, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(volumesnapshotcontentsResource, "status", volumeSnapshotContent), &v1beta1.VolumeSnapshotContent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeSnapshotContent), err
}

func (c *FakeVolumeSnapshotContents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(volumesnapshotcontentsResource, name), &v1beta1.VolumeSnapshotContent{})
	return err
}

func (c *FakeVolumeSnapshotContents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewRootDeleteCollectionAction(volumesnapshotcontentsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.VolumeSnapshotContentList{})
	return err
}

func (c *FakeVolumeSnapshotContents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VolumeSnapshotContent, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(volumesnapshotcontentsResource, name, pt, data, subresources...), &v1beta1.VolumeSnapshotContent{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeSnapshotContent), err
}
