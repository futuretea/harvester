package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	apiv1 "kubevirt.io/client-go/api/v1"
)

type FakeVirtualMachines struct {
	Fake	*FakeKubevirtV1
	ns	string
}

var virtualmachinesResource = schema.GroupVersionResource{Group: "kubevirt.io", Version: "v1", Resource: "virtualmachines"}

var virtualmachinesKind = schema.GroupVersionKind{Group: "kubevirt.io", Version: "v1", Kind: "VirtualMachine"}

func (c *FakeVirtualMachines) Get(ctx context.Context, name string, options v1.GetOptions) (result *apiv1.VirtualMachine, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualmachinesResource, c.ns, name), &apiv1.VirtualMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.VirtualMachine), err
}

func (c *FakeVirtualMachines) List(ctx context.Context, opts v1.ListOptions) (result *apiv1.VirtualMachineList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualmachinesResource, virtualmachinesKind, c.ns, opts), &apiv1.VirtualMachineList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &apiv1.VirtualMachineList{ListMeta: obj.(*apiv1.VirtualMachineList).ListMeta}
	for _, item := range obj.(*apiv1.VirtualMachineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeVirtualMachines) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualmachinesResource, c.ns, opts))

}

func (c *FakeVirtualMachines) Create(ctx context.Context, virtualMachine *apiv1.VirtualMachine, opts v1.CreateOptions) (result *apiv1.VirtualMachine, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualmachinesResource, c.ns, virtualMachine), &apiv1.VirtualMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.VirtualMachine), err
}

func (c *FakeVirtualMachines) Update(ctx context.Context, virtualMachine *apiv1.VirtualMachine, opts v1.UpdateOptions) (result *apiv1.VirtualMachine, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualmachinesResource, c.ns, virtualMachine), &apiv1.VirtualMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.VirtualMachine), err
}

func (c *FakeVirtualMachines) UpdateStatus(ctx context.Context, virtualMachine *apiv1.VirtualMachine, opts v1.UpdateOptions) (*apiv1.VirtualMachine, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualmachinesResource, "status", c.ns, virtualMachine), &apiv1.VirtualMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.VirtualMachine), err
}

func (c *FakeVirtualMachines) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualmachinesResource, c.ns, name), &apiv1.VirtualMachine{})

	return err
}

func (c *FakeVirtualMachines) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(virtualmachinesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &apiv1.VirtualMachineList{})
	return err
}

func (c *FakeVirtualMachines) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *apiv1.VirtualMachine, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualmachinesResource, c.ns, name, pt, data, subresources...), &apiv1.VirtualMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.VirtualMachine), err
}
