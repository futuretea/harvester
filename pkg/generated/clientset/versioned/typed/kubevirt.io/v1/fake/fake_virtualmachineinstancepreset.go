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

type FakeVirtualMachineInstancePresets struct {
	Fake	*FakeKubevirtV1
	ns	string
}

var virtualmachineinstancepresetsResource = schema.GroupVersionResource{Group: "kubevirt.io", Version: "v1", Resource: "virtualmachineinstancepresets"}

var virtualmachineinstancepresetsKind = schema.GroupVersionKind{Group: "kubevirt.io", Version: "v1", Kind: "VirtualMachineInstancePreset"}

func (c *FakeVirtualMachineInstancePresets) Get(ctx context.Context, name string, options v1.GetOptions) (result *apiv1.VirtualMachineInstancePreset, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualmachineinstancepresetsResource, c.ns, name), &apiv1.VirtualMachineInstancePreset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.VirtualMachineInstancePreset), err
}

func (c *FakeVirtualMachineInstancePresets) List(ctx context.Context, opts v1.ListOptions) (result *apiv1.VirtualMachineInstancePresetList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualmachineinstancepresetsResource, virtualmachineinstancepresetsKind, c.ns, opts), &apiv1.VirtualMachineInstancePresetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &apiv1.VirtualMachineInstancePresetList{ListMeta: obj.(*apiv1.VirtualMachineInstancePresetList).ListMeta}
	for _, item := range obj.(*apiv1.VirtualMachineInstancePresetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeVirtualMachineInstancePresets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualmachineinstancepresetsResource, c.ns, opts))

}

func (c *FakeVirtualMachineInstancePresets) Create(ctx context.Context, virtualMachineInstancePreset *apiv1.VirtualMachineInstancePreset, opts v1.CreateOptions) (result *apiv1.VirtualMachineInstancePreset, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualmachineinstancepresetsResource, c.ns, virtualMachineInstancePreset), &apiv1.VirtualMachineInstancePreset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.VirtualMachineInstancePreset), err
}

func (c *FakeVirtualMachineInstancePresets) Update(ctx context.Context, virtualMachineInstancePreset *apiv1.VirtualMachineInstancePreset, opts v1.UpdateOptions) (result *apiv1.VirtualMachineInstancePreset, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualmachineinstancepresetsResource, c.ns, virtualMachineInstancePreset), &apiv1.VirtualMachineInstancePreset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.VirtualMachineInstancePreset), err
}

func (c *FakeVirtualMachineInstancePresets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualmachineinstancepresetsResource, c.ns, name), &apiv1.VirtualMachineInstancePreset{})

	return err
}

func (c *FakeVirtualMachineInstancePresets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(virtualmachineinstancepresetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &apiv1.VirtualMachineInstancePresetList{})
	return err
}

func (c *FakeVirtualMachineInstancePresets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *apiv1.VirtualMachineInstancePreset, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualmachineinstancepresetsResource, c.ns, name, pt, data, subresources...), &apiv1.VirtualMachineInstancePreset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.VirtualMachineInstancePreset), err
}
