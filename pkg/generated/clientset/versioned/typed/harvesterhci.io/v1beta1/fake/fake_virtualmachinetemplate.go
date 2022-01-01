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

type FakeVirtualMachineTemplates struct {
	Fake	*FakeHarvesterhciV1beta1
	ns	string
}

var virtualmachinetemplatesResource = schema.GroupVersionResource{Group: "harvesterhci.io", Version: "v1beta1", Resource: "virtualmachinetemplates"}

var virtualmachinetemplatesKind = schema.GroupVersionKind{Group: "harvesterhci.io", Version: "v1beta1", Kind: "VirtualMachineTemplate"}

func (c *FakeVirtualMachineTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VirtualMachineTemplate, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualmachinetemplatesResource, c.ns, name), &v1beta1.VirtualMachineTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineTemplate), err
}

func (c *FakeVirtualMachineTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VirtualMachineTemplateList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualmachinetemplatesResource, virtualmachinetemplatesKind, c.ns, opts), &v1beta1.VirtualMachineTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.VirtualMachineTemplateList{ListMeta: obj.(*v1beta1.VirtualMachineTemplateList).ListMeta}
	for _, item := range obj.(*v1beta1.VirtualMachineTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeVirtualMachineTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualmachinetemplatesResource, c.ns, opts))

}

func (c *FakeVirtualMachineTemplates) Create(ctx context.Context, virtualMachineTemplate *v1beta1.VirtualMachineTemplate, opts v1.CreateOptions) (result *v1beta1.VirtualMachineTemplate, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualmachinetemplatesResource, c.ns, virtualMachineTemplate), &v1beta1.VirtualMachineTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineTemplate), err
}

func (c *FakeVirtualMachineTemplates) Update(ctx context.Context, virtualMachineTemplate *v1beta1.VirtualMachineTemplate, opts v1.UpdateOptions) (result *v1beta1.VirtualMachineTemplate, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualmachinetemplatesResource, c.ns, virtualMachineTemplate), &v1beta1.VirtualMachineTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineTemplate), err
}

func (c *FakeVirtualMachineTemplates) UpdateStatus(ctx context.Context, virtualMachineTemplate *v1beta1.VirtualMachineTemplate, opts v1.UpdateOptions) (*v1beta1.VirtualMachineTemplate, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualmachinetemplatesResource, "status", c.ns, virtualMachineTemplate), &v1beta1.VirtualMachineTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineTemplate), err
}

func (c *FakeVirtualMachineTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualmachinetemplatesResource, c.ns, name), &v1beta1.VirtualMachineTemplate{})

	return err
}

func (c *FakeVirtualMachineTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(virtualmachinetemplatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.VirtualMachineTemplateList{})
	return err
}

func (c *FakeVirtualMachineTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VirtualMachineTemplate, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualmachinetemplatesResource, c.ns, name, pt, data, subresources...), &v1beta1.VirtualMachineTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineTemplate), err
}
