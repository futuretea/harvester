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

type FakeVirtualMachineInstances struct {
	Fake	*FakeKubevirtV1
	ns	string
}

var virtualmachineinstancesResource = schema.GroupVersionResource{Group: "kubevirt.io", Version: "v1", Resource: "virtualmachineinstances"}

var virtualmachineinstancesKind = schema.GroupVersionKind{Group: "kubevirt.io", Version: "v1", Kind: "VirtualMachineInstance"}

func (c *FakeVirtualMachineInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *apiv1.VirtualMachineInstance, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualmachineinstancesResource, c.ns, name), &apiv1.VirtualMachineInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.VirtualMachineInstance), err
}

func (c *FakeVirtualMachineInstances) List(ctx context.Context, opts v1.ListOptions) (result *apiv1.VirtualMachineInstanceList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualmachineinstancesResource, virtualmachineinstancesKind, c.ns, opts), &apiv1.VirtualMachineInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &apiv1.VirtualMachineInstanceList{ListMeta: obj.(*apiv1.VirtualMachineInstanceList).ListMeta}
	for _, item := range obj.(*apiv1.VirtualMachineInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeVirtualMachineInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualmachineinstancesResource, c.ns, opts))

}

func (c *FakeVirtualMachineInstances) Create(ctx context.Context, virtualMachineInstance *apiv1.VirtualMachineInstance, opts v1.CreateOptions) (result *apiv1.VirtualMachineInstance, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualmachineinstancesResource, c.ns, virtualMachineInstance), &apiv1.VirtualMachineInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.VirtualMachineInstance), err
}

func (c *FakeVirtualMachineInstances) Update(ctx context.Context, virtualMachineInstance *apiv1.VirtualMachineInstance, opts v1.UpdateOptions) (result *apiv1.VirtualMachineInstance, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualmachineinstancesResource, c.ns, virtualMachineInstance), &apiv1.VirtualMachineInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.VirtualMachineInstance), err
}

func (c *FakeVirtualMachineInstances) UpdateStatus(ctx context.Context, virtualMachineInstance *apiv1.VirtualMachineInstance, opts v1.UpdateOptions) (*apiv1.VirtualMachineInstance, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualmachineinstancesResource, "status", c.ns, virtualMachineInstance), &apiv1.VirtualMachineInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.VirtualMachineInstance), err
}

func (c *FakeVirtualMachineInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualmachineinstancesResource, c.ns, name), &apiv1.VirtualMachineInstance{})

	return err
}

func (c *FakeVirtualMachineInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(virtualmachineinstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &apiv1.VirtualMachineInstanceList{})
	return err
}

func (c *FakeVirtualMachineInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *apiv1.VirtualMachineInstance, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualmachineinstancesResource, c.ns, name, pt, data, subresources...), &apiv1.VirtualMachineInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.VirtualMachineInstance), err
}
