package fake

import (
	"context"

	v1beta1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

type FakeVolumes struct {
	Fake	*FakeLonghornV1beta1
	ns	string
}

var volumesResource = schema.GroupVersionResource{Group: "longhorn.io", Version: "v1beta1", Resource: "volumes"}

var volumesKind = schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta1", Kind: "Volume"}

func (c *FakeVolumes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Volume, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(volumesResource, c.ns, name), &v1beta1.Volume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Volume), err
}

func (c *FakeVolumes) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VolumeList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(volumesResource, volumesKind, c.ns, opts), &v1beta1.VolumeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.VolumeList{ListMeta: obj.(*v1beta1.VolumeList).ListMeta}
	for _, item := range obj.(*v1beta1.VolumeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeVolumes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(volumesResource, c.ns, opts))

}

func (c *FakeVolumes) Create(ctx context.Context, volume *v1beta1.Volume, opts v1.CreateOptions) (result *v1beta1.Volume, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(volumesResource, c.ns, volume), &v1beta1.Volume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Volume), err
}

func (c *FakeVolumes) Update(ctx context.Context, volume *v1beta1.Volume, opts v1.UpdateOptions) (result *v1beta1.Volume, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(volumesResource, c.ns, volume), &v1beta1.Volume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Volume), err
}

func (c *FakeVolumes) UpdateStatus(ctx context.Context, volume *v1beta1.Volume, opts v1.UpdateOptions) (*v1beta1.Volume, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(volumesResource, "status", c.ns, volume), &v1beta1.Volume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Volume), err
}

func (c *FakeVolumes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(volumesResource, c.ns, name), &v1beta1.Volume{})

	return err
}

func (c *FakeVolumes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(volumesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.VolumeList{})
	return err
}

func (c *FakeVolumes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Volume, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(volumesResource, c.ns, name, pt, data, subresources...), &v1beta1.Volume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Volume), err
}
