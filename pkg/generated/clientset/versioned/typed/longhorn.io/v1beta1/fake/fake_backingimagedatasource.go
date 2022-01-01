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

type FakeBackingImageDataSources struct {
	Fake	*FakeLonghornV1beta1
	ns	string
}

var backingimagedatasourcesResource = schema.GroupVersionResource{Group: "longhorn.io", Version: "v1beta1", Resource: "backingimagedatasources"}

var backingimagedatasourcesKind = schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta1", Kind: "BackingImageDataSource"}

func (c *FakeBackingImageDataSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.BackingImageDataSource, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(backingimagedatasourcesResource, c.ns, name), &v1beta1.BackingImageDataSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackingImageDataSource), err
}

func (c *FakeBackingImageDataSources) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.BackingImageDataSourceList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(backingimagedatasourcesResource, backingimagedatasourcesKind, c.ns, opts), &v1beta1.BackingImageDataSourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.BackingImageDataSourceList{ListMeta: obj.(*v1beta1.BackingImageDataSourceList).ListMeta}
	for _, item := range obj.(*v1beta1.BackingImageDataSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeBackingImageDataSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(backingimagedatasourcesResource, c.ns, opts))

}

func (c *FakeBackingImageDataSources) Create(ctx context.Context, backingImageDataSource *v1beta1.BackingImageDataSource, opts v1.CreateOptions) (result *v1beta1.BackingImageDataSource, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(backingimagedatasourcesResource, c.ns, backingImageDataSource), &v1beta1.BackingImageDataSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackingImageDataSource), err
}

func (c *FakeBackingImageDataSources) Update(ctx context.Context, backingImageDataSource *v1beta1.BackingImageDataSource, opts v1.UpdateOptions) (result *v1beta1.BackingImageDataSource, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(backingimagedatasourcesResource, c.ns, backingImageDataSource), &v1beta1.BackingImageDataSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackingImageDataSource), err
}

func (c *FakeBackingImageDataSources) UpdateStatus(ctx context.Context, backingImageDataSource *v1beta1.BackingImageDataSource, opts v1.UpdateOptions) (*v1beta1.BackingImageDataSource, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(backingimagedatasourcesResource, "status", c.ns, backingImageDataSource), &v1beta1.BackingImageDataSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackingImageDataSource), err
}

func (c *FakeBackingImageDataSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(backingimagedatasourcesResource, c.ns, name), &v1beta1.BackingImageDataSource{})

	return err
}

func (c *FakeBackingImageDataSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(backingimagedatasourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.BackingImageDataSourceList{})
	return err
}

func (c *FakeBackingImageDataSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BackingImageDataSource, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(backingimagedatasourcesResource, c.ns, name, pt, data, subresources...), &v1beta1.BackingImageDataSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackingImageDataSource), err
}
