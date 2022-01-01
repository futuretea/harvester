package fake

import (
	"context"

	v1beta1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

type FakeSupportBundles struct {
	Fake	*FakeHarvesterhciV1beta1
	ns	string
}

var supportbundlesResource = schema.GroupVersionResource{Group: "harvesterhci.io", Version: "v1beta1", Resource: "supportbundles"}

var supportbundlesKind = schema.GroupVersionKind{Group: "harvesterhci.io", Version: "v1beta1", Kind: "SupportBundle"}

func (c *FakeSupportBundles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.SupportBundle, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(supportbundlesResource, c.ns, name), &v1beta1.SupportBundle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SupportBundle), err
}

func (c *FakeSupportBundles) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.SupportBundleList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(supportbundlesResource, supportbundlesKind, c.ns, opts), &v1beta1.SupportBundleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.SupportBundleList{ListMeta: obj.(*v1beta1.SupportBundleList).ListMeta}
	for _, item := range obj.(*v1beta1.SupportBundleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeSupportBundles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(supportbundlesResource, c.ns, opts))

}

func (c *FakeSupportBundles) Create(ctx context.Context, supportBundle *v1beta1.SupportBundle, opts v1.CreateOptions) (result *v1beta1.SupportBundle, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(supportbundlesResource, c.ns, supportBundle), &v1beta1.SupportBundle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SupportBundle), err
}

func (c *FakeSupportBundles) Update(ctx context.Context, supportBundle *v1beta1.SupportBundle, opts v1.UpdateOptions) (result *v1beta1.SupportBundle, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(supportbundlesResource, c.ns, supportBundle), &v1beta1.SupportBundle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SupportBundle), err
}

func (c *FakeSupportBundles) UpdateStatus(ctx context.Context, supportBundle *v1beta1.SupportBundle, opts v1.UpdateOptions) (*v1beta1.SupportBundle, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(supportbundlesResource, "status", c.ns, supportBundle), &v1beta1.SupportBundle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SupportBundle), err
}

func (c *FakeSupportBundles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(supportbundlesResource, c.ns, name), &v1beta1.SupportBundle{})

	return err
}

func (c *FakeSupportBundles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(supportbundlesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.SupportBundleList{})
	return err
}

func (c *FakeSupportBundles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.SupportBundle, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(supportbundlesResource, c.ns, name, pt, data, subresources...), &v1beta1.SupportBundle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.SupportBundle), err
}
