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

type FakeVirtualMachineImages struct {
	Fake	*FakeHarvesterhciV1beta1
	ns	string
}

var virtualmachineimagesResource = schema.GroupVersionResource{Group: "harvesterhci.io", Version: "v1beta1", Resource: "virtualmachineimages"}

var virtualmachineimagesKind = schema.GroupVersionKind{Group: "harvesterhci.io", Version: "v1beta1", Kind: "VirtualMachineImage"}

func (c *FakeVirtualMachineImages) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VirtualMachineImage, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualmachineimagesResource, c.ns, name), &v1beta1.VirtualMachineImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineImage), err
}

func (c *FakeVirtualMachineImages) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VirtualMachineImageList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualmachineimagesResource, virtualmachineimagesKind, c.ns, opts), &v1beta1.VirtualMachineImageList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.VirtualMachineImageList{ListMeta: obj.(*v1beta1.VirtualMachineImageList).ListMeta}
	for _, item := range obj.(*v1beta1.VirtualMachineImageList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeVirtualMachineImages) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualmachineimagesResource, c.ns, opts))

}

func (c *FakeVirtualMachineImages) Create(ctx context.Context, virtualMachineImage *v1beta1.VirtualMachineImage, opts v1.CreateOptions) (result *v1beta1.VirtualMachineImage, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualmachineimagesResource, c.ns, virtualMachineImage), &v1beta1.VirtualMachineImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineImage), err
}

func (c *FakeVirtualMachineImages) Update(ctx context.Context, virtualMachineImage *v1beta1.VirtualMachineImage, opts v1.UpdateOptions) (result *v1beta1.VirtualMachineImage, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualmachineimagesResource, c.ns, virtualMachineImage), &v1beta1.VirtualMachineImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineImage), err
}

func (c *FakeVirtualMachineImages) UpdateStatus(ctx context.Context, virtualMachineImage *v1beta1.VirtualMachineImage, opts v1.UpdateOptions) (*v1beta1.VirtualMachineImage, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualmachineimagesResource, "status", c.ns, virtualMachineImage), &v1beta1.VirtualMachineImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineImage), err
}

func (c *FakeVirtualMachineImages) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualmachineimagesResource, c.ns, name), &v1beta1.VirtualMachineImage{})

	return err
}

func (c *FakeVirtualMachineImages) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(virtualmachineimagesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.VirtualMachineImageList{})
	return err
}

func (c *FakeVirtualMachineImages) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VirtualMachineImage, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualmachineimagesResource, c.ns, name, pt, data, subresources...), &v1beta1.VirtualMachineImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineImage), err
}
