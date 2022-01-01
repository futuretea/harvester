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

type FakeVolumeSnapshotClasses struct {
	Fake *FakeSnapshotV1beta1
}

var volumesnapshotclassesResource = schema.GroupVersionResource{Group: "snapshot.storage.k8s.io", Version: "v1beta1", Resource: "volumesnapshotclasses"}

var volumesnapshotclassesKind = schema.GroupVersionKind{Group: "snapshot.storage.k8s.io", Version: "v1beta1", Kind: "VolumeSnapshotClass"}

func (c *FakeVolumeSnapshotClasses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VolumeSnapshotClass, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(volumesnapshotclassesResource, name), &v1beta1.VolumeSnapshotClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeSnapshotClass), err
}

func (c *FakeVolumeSnapshotClasses) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VolumeSnapshotClassList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(volumesnapshotclassesResource, volumesnapshotclassesKind, opts), &v1beta1.VolumeSnapshotClassList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.VolumeSnapshotClassList{ListMeta: obj.(*v1beta1.VolumeSnapshotClassList).ListMeta}
	for _, item := range obj.(*v1beta1.VolumeSnapshotClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeVolumeSnapshotClasses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(volumesnapshotclassesResource, opts))
}

func (c *FakeVolumeSnapshotClasses) Create(ctx context.Context, volumeSnapshotClass *v1beta1.VolumeSnapshotClass, opts v1.CreateOptions) (result *v1beta1.VolumeSnapshotClass, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(volumesnapshotclassesResource, volumeSnapshotClass), &v1beta1.VolumeSnapshotClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeSnapshotClass), err
}

func (c *FakeVolumeSnapshotClasses) Update(ctx context.Context, volumeSnapshotClass *v1beta1.VolumeSnapshotClass, opts v1.UpdateOptions) (result *v1beta1.VolumeSnapshotClass, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(volumesnapshotclassesResource, volumeSnapshotClass), &v1beta1.VolumeSnapshotClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeSnapshotClass), err
}

func (c *FakeVolumeSnapshotClasses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(volumesnapshotclassesResource, name), &v1beta1.VolumeSnapshotClass{})
	return err
}

func (c *FakeVolumeSnapshotClasses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewRootDeleteCollectionAction(volumesnapshotclassesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.VolumeSnapshotClassList{})
	return err
}

func (c *FakeVolumeSnapshotClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VolumeSnapshotClass, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(volumesnapshotclassesResource, name, pt, data, subresources...), &v1beta1.VolumeSnapshotClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VolumeSnapshotClass), err
}
