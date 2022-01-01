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

type FakeVirtualMachineBackups struct {
	Fake	*FakeHarvesterhciV1beta1
	ns	string
}

var virtualmachinebackupsResource = schema.GroupVersionResource{Group: "harvesterhci.io", Version: "v1beta1", Resource: "virtualmachinebackups"}

var virtualmachinebackupsKind = schema.GroupVersionKind{Group: "harvesterhci.io", Version: "v1beta1", Kind: "VirtualMachineBackup"}

func (c *FakeVirtualMachineBackups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VirtualMachineBackup, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualmachinebackupsResource, c.ns, name), &v1beta1.VirtualMachineBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineBackup), err
}

func (c *FakeVirtualMachineBackups) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VirtualMachineBackupList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualmachinebackupsResource, virtualmachinebackupsKind, c.ns, opts), &v1beta1.VirtualMachineBackupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.VirtualMachineBackupList{ListMeta: obj.(*v1beta1.VirtualMachineBackupList).ListMeta}
	for _, item := range obj.(*v1beta1.VirtualMachineBackupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeVirtualMachineBackups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualmachinebackupsResource, c.ns, opts))

}

func (c *FakeVirtualMachineBackups) Create(ctx context.Context, virtualMachineBackup *v1beta1.VirtualMachineBackup, opts v1.CreateOptions) (result *v1beta1.VirtualMachineBackup, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualmachinebackupsResource, c.ns, virtualMachineBackup), &v1beta1.VirtualMachineBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineBackup), err
}

func (c *FakeVirtualMachineBackups) Update(ctx context.Context, virtualMachineBackup *v1beta1.VirtualMachineBackup, opts v1.UpdateOptions) (result *v1beta1.VirtualMachineBackup, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualmachinebackupsResource, c.ns, virtualMachineBackup), &v1beta1.VirtualMachineBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineBackup), err
}

func (c *FakeVirtualMachineBackups) UpdateStatus(ctx context.Context, virtualMachineBackup *v1beta1.VirtualMachineBackup, opts v1.UpdateOptions) (*v1beta1.VirtualMachineBackup, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualmachinebackupsResource, "status", c.ns, virtualMachineBackup), &v1beta1.VirtualMachineBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineBackup), err
}

func (c *FakeVirtualMachineBackups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualmachinebackupsResource, c.ns, name), &v1beta1.VirtualMachineBackup{})

	return err
}

func (c *FakeVirtualMachineBackups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(virtualmachinebackupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.VirtualMachineBackupList{})
	return err
}

func (c *FakeVirtualMachineBackups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VirtualMachineBackup, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualmachinebackupsResource, c.ns, name, pt, data, subresources...), &v1beta1.VirtualMachineBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineBackup), err
}
