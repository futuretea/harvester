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

type FakeVirtualMachineInstanceReplicaSets struct {
	Fake	*FakeKubevirtV1
	ns	string
}

var virtualmachineinstancereplicasetsResource = schema.GroupVersionResource{Group: "kubevirt.io", Version: "v1", Resource: "virtualmachineinstancereplicasets"}

var virtualmachineinstancereplicasetsKind = schema.GroupVersionKind{Group: "kubevirt.io", Version: "v1", Kind: "VirtualMachineInstanceReplicaSet"}

func (c *FakeVirtualMachineInstanceReplicaSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *apiv1.VirtualMachineInstanceReplicaSet, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualmachineinstancereplicasetsResource, c.ns, name), &apiv1.VirtualMachineInstanceReplicaSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.VirtualMachineInstanceReplicaSet), err
}

func (c *FakeVirtualMachineInstanceReplicaSets) List(ctx context.Context, opts v1.ListOptions) (result *apiv1.VirtualMachineInstanceReplicaSetList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualmachineinstancereplicasetsResource, virtualmachineinstancereplicasetsKind, c.ns, opts), &apiv1.VirtualMachineInstanceReplicaSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &apiv1.VirtualMachineInstanceReplicaSetList{ListMeta: obj.(*apiv1.VirtualMachineInstanceReplicaSetList).ListMeta}
	for _, item := range obj.(*apiv1.VirtualMachineInstanceReplicaSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeVirtualMachineInstanceReplicaSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualmachineinstancereplicasetsResource, c.ns, opts))

}

func (c *FakeVirtualMachineInstanceReplicaSets) Create(ctx context.Context, virtualMachineInstanceReplicaSet *apiv1.VirtualMachineInstanceReplicaSet, opts v1.CreateOptions) (result *apiv1.VirtualMachineInstanceReplicaSet, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualmachineinstancereplicasetsResource, c.ns, virtualMachineInstanceReplicaSet), &apiv1.VirtualMachineInstanceReplicaSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.VirtualMachineInstanceReplicaSet), err
}

func (c *FakeVirtualMachineInstanceReplicaSets) Update(ctx context.Context, virtualMachineInstanceReplicaSet *apiv1.VirtualMachineInstanceReplicaSet, opts v1.UpdateOptions) (result *apiv1.VirtualMachineInstanceReplicaSet, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualmachineinstancereplicasetsResource, c.ns, virtualMachineInstanceReplicaSet), &apiv1.VirtualMachineInstanceReplicaSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.VirtualMachineInstanceReplicaSet), err
}

func (c *FakeVirtualMachineInstanceReplicaSets) UpdateStatus(ctx context.Context, virtualMachineInstanceReplicaSet *apiv1.VirtualMachineInstanceReplicaSet, opts v1.UpdateOptions) (*apiv1.VirtualMachineInstanceReplicaSet, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualmachineinstancereplicasetsResource, "status", c.ns, virtualMachineInstanceReplicaSet), &apiv1.VirtualMachineInstanceReplicaSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.VirtualMachineInstanceReplicaSet), err
}

func (c *FakeVirtualMachineInstanceReplicaSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualmachineinstancereplicasetsResource, c.ns, name), &apiv1.VirtualMachineInstanceReplicaSet{})

	return err
}

func (c *FakeVirtualMachineInstanceReplicaSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(virtualmachineinstancereplicasetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &apiv1.VirtualMachineInstanceReplicaSetList{})
	return err
}

func (c *FakeVirtualMachineInstanceReplicaSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *apiv1.VirtualMachineInstanceReplicaSet, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualmachineinstancereplicasetsResource, c.ns, name, pt, data, subresources...), &apiv1.VirtualMachineInstanceReplicaSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.VirtualMachineInstanceReplicaSet), err
}
