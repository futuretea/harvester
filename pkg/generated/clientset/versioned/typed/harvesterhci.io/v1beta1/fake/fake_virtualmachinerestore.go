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

type FakeVirtualMachineRestores struct {
	Fake	*FakeHarvesterhciV1beta1
	ns	string
}

var virtualmachinerestoresResource = schema.GroupVersionResource{Group: "harvesterhci.io", Version: "v1beta1", Resource: "virtualmachinerestores"}

var virtualmachinerestoresKind = schema.GroupVersionKind{Group: "harvesterhci.io", Version: "v1beta1", Kind: "VirtualMachineRestore"}

func (c *FakeVirtualMachineRestores) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VirtualMachineRestore, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualmachinerestoresResource, c.ns, name), &v1beta1.VirtualMachineRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineRestore), err
}

func (c *FakeVirtualMachineRestores) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VirtualMachineRestoreList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualmachinerestoresResource, virtualmachinerestoresKind, c.ns, opts), &v1beta1.VirtualMachineRestoreList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.VirtualMachineRestoreList{ListMeta: obj.(*v1beta1.VirtualMachineRestoreList).ListMeta}
	for _, item := range obj.(*v1beta1.VirtualMachineRestoreList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeVirtualMachineRestores) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualmachinerestoresResource, c.ns, opts))

}

func (c *FakeVirtualMachineRestores) Create(ctx context.Context, virtualMachineRestore *v1beta1.VirtualMachineRestore, opts v1.CreateOptions) (result *v1beta1.VirtualMachineRestore, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualmachinerestoresResource, c.ns, virtualMachineRestore), &v1beta1.VirtualMachineRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineRestore), err
}

func (c *FakeVirtualMachineRestores) Update(ctx context.Context, virtualMachineRestore *v1beta1.VirtualMachineRestore, opts v1.UpdateOptions) (result *v1beta1.VirtualMachineRestore, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualmachinerestoresResource, c.ns, virtualMachineRestore), &v1beta1.VirtualMachineRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineRestore), err
}

func (c *FakeVirtualMachineRestores) UpdateStatus(ctx context.Context, virtualMachineRestore *v1beta1.VirtualMachineRestore, opts v1.UpdateOptions) (*v1beta1.VirtualMachineRestore, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualmachinerestoresResource, "status", c.ns, virtualMachineRestore), &v1beta1.VirtualMachineRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineRestore), err
}

func (c *FakeVirtualMachineRestores) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualmachinerestoresResource, c.ns, name), &v1beta1.VirtualMachineRestore{})

	return err
}

func (c *FakeVirtualMachineRestores) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(virtualmachinerestoresResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.VirtualMachineRestoreList{})
	return err
}

func (c *FakeVirtualMachineRestores) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VirtualMachineRestore, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualmachinerestoresResource, c.ns, name, pt, data, subresources...), &v1beta1.VirtualMachineRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineRestore), err
}
