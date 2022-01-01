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

type FakeVirtualMachineInstanceMigrations struct {
	Fake	*FakeKubevirtV1
	ns	string
}

var virtualmachineinstancemigrationsResource = schema.GroupVersionResource{Group: "kubevirt.io", Version: "v1", Resource: "virtualmachineinstancemigrations"}

var virtualmachineinstancemigrationsKind = schema.GroupVersionKind{Group: "kubevirt.io", Version: "v1", Kind: "VirtualMachineInstanceMigration"}

func (c *FakeVirtualMachineInstanceMigrations) Get(ctx context.Context, name string, options v1.GetOptions) (result *apiv1.VirtualMachineInstanceMigration, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualmachineinstancemigrationsResource, c.ns, name), &apiv1.VirtualMachineInstanceMigration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.VirtualMachineInstanceMigration), err
}

func (c *FakeVirtualMachineInstanceMigrations) List(ctx context.Context, opts v1.ListOptions) (result *apiv1.VirtualMachineInstanceMigrationList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualmachineinstancemigrationsResource, virtualmachineinstancemigrationsKind, c.ns, opts), &apiv1.VirtualMachineInstanceMigrationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &apiv1.VirtualMachineInstanceMigrationList{ListMeta: obj.(*apiv1.VirtualMachineInstanceMigrationList).ListMeta}
	for _, item := range obj.(*apiv1.VirtualMachineInstanceMigrationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeVirtualMachineInstanceMigrations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualmachineinstancemigrationsResource, c.ns, opts))

}

func (c *FakeVirtualMachineInstanceMigrations) Create(ctx context.Context, virtualMachineInstanceMigration *apiv1.VirtualMachineInstanceMigration, opts v1.CreateOptions) (result *apiv1.VirtualMachineInstanceMigration, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualmachineinstancemigrationsResource, c.ns, virtualMachineInstanceMigration), &apiv1.VirtualMachineInstanceMigration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.VirtualMachineInstanceMigration), err
}

func (c *FakeVirtualMachineInstanceMigrations) Update(ctx context.Context, virtualMachineInstanceMigration *apiv1.VirtualMachineInstanceMigration, opts v1.UpdateOptions) (result *apiv1.VirtualMachineInstanceMigration, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualmachineinstancemigrationsResource, c.ns, virtualMachineInstanceMigration), &apiv1.VirtualMachineInstanceMigration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.VirtualMachineInstanceMigration), err
}

func (c *FakeVirtualMachineInstanceMigrations) UpdateStatus(ctx context.Context, virtualMachineInstanceMigration *apiv1.VirtualMachineInstanceMigration, opts v1.UpdateOptions) (*apiv1.VirtualMachineInstanceMigration, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualmachineinstancemigrationsResource, "status", c.ns, virtualMachineInstanceMigration), &apiv1.VirtualMachineInstanceMigration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.VirtualMachineInstanceMigration), err
}

func (c *FakeVirtualMachineInstanceMigrations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualmachineinstancemigrationsResource, c.ns, name), &apiv1.VirtualMachineInstanceMigration{})

	return err
}

func (c *FakeVirtualMachineInstanceMigrations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(virtualmachineinstancemigrationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &apiv1.VirtualMachineInstanceMigrationList{})
	return err
}

func (c *FakeVirtualMachineInstanceMigrations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *apiv1.VirtualMachineInstanceMigration, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualmachineinstancemigrationsResource, c.ns, name, pt, data, subresources...), &apiv1.VirtualMachineInstanceMigration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apiv1.VirtualMachineInstanceMigration), err
}
