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

type FakeBackingImageManagers struct {
	Fake	*FakeLonghornV1beta1
	ns	string
}

var backingimagemanagersResource = schema.GroupVersionResource{Group: "longhorn.io", Version: "v1beta1", Resource: "backingimagemanagers"}

var backingimagemanagersKind = schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta1", Kind: "BackingImageManager"}

func (c *FakeBackingImageManagers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.BackingImageManager, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(backingimagemanagersResource, c.ns, name), &v1beta1.BackingImageManager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackingImageManager), err
}

func (c *FakeBackingImageManagers) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.BackingImageManagerList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(backingimagemanagersResource, backingimagemanagersKind, c.ns, opts), &v1beta1.BackingImageManagerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.BackingImageManagerList{ListMeta: obj.(*v1beta1.BackingImageManagerList).ListMeta}
	for _, item := range obj.(*v1beta1.BackingImageManagerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeBackingImageManagers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(backingimagemanagersResource, c.ns, opts))

}

func (c *FakeBackingImageManagers) Create(ctx context.Context, backingImageManager *v1beta1.BackingImageManager, opts v1.CreateOptions) (result *v1beta1.BackingImageManager, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(backingimagemanagersResource, c.ns, backingImageManager), &v1beta1.BackingImageManager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackingImageManager), err
}

func (c *FakeBackingImageManagers) Update(ctx context.Context, backingImageManager *v1beta1.BackingImageManager, opts v1.UpdateOptions) (result *v1beta1.BackingImageManager, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(backingimagemanagersResource, c.ns, backingImageManager), &v1beta1.BackingImageManager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackingImageManager), err
}

func (c *FakeBackingImageManagers) UpdateStatus(ctx context.Context, backingImageManager *v1beta1.BackingImageManager, opts v1.UpdateOptions) (*v1beta1.BackingImageManager, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(backingimagemanagersResource, "status", c.ns, backingImageManager), &v1beta1.BackingImageManager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackingImageManager), err
}

func (c *FakeBackingImageManagers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(backingimagemanagersResource, c.ns, name), &v1beta1.BackingImageManager{})

	return err
}

func (c *FakeBackingImageManagers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(backingimagemanagersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.BackingImageManagerList{})
	return err
}

func (c *FakeBackingImageManagers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BackingImageManager, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(backingimagemanagersResource, c.ns, name, pt, data, subresources...), &v1beta1.BackingImageManager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackingImageManager), err
}
