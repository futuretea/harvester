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

type FakeVirtualMachineTemplateVersions struct {
	Fake	*FakeHarvesterhciV1beta1
	ns	string
}

var virtualmachinetemplateversionsResource = schema.GroupVersionResource{Group: "harvesterhci.io", Version: "v1beta1", Resource: "virtualmachinetemplateversions"}

var virtualmachinetemplateversionsKind = schema.GroupVersionKind{Group: "harvesterhci.io", Version: "v1beta1", Kind: "VirtualMachineTemplateVersion"}

func (c *FakeVirtualMachineTemplateVersions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VirtualMachineTemplateVersion, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualmachinetemplateversionsResource, c.ns, name), &v1beta1.VirtualMachineTemplateVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineTemplateVersion), err
}

func (c *FakeVirtualMachineTemplateVersions) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VirtualMachineTemplateVersionList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualmachinetemplateversionsResource, virtualmachinetemplateversionsKind, c.ns, opts), &v1beta1.VirtualMachineTemplateVersionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.VirtualMachineTemplateVersionList{ListMeta: obj.(*v1beta1.VirtualMachineTemplateVersionList).ListMeta}
	for _, item := range obj.(*v1beta1.VirtualMachineTemplateVersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeVirtualMachineTemplateVersions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualmachinetemplateversionsResource, c.ns, opts))

}

func (c *FakeVirtualMachineTemplateVersions) Create(ctx context.Context, virtualMachineTemplateVersion *v1beta1.VirtualMachineTemplateVersion, opts v1.CreateOptions) (result *v1beta1.VirtualMachineTemplateVersion, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualmachinetemplateversionsResource, c.ns, virtualMachineTemplateVersion), &v1beta1.VirtualMachineTemplateVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineTemplateVersion), err
}

func (c *FakeVirtualMachineTemplateVersions) Update(ctx context.Context, virtualMachineTemplateVersion *v1beta1.VirtualMachineTemplateVersion, opts v1.UpdateOptions) (result *v1beta1.VirtualMachineTemplateVersion, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualmachinetemplateversionsResource, c.ns, virtualMachineTemplateVersion), &v1beta1.VirtualMachineTemplateVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineTemplateVersion), err
}

func (c *FakeVirtualMachineTemplateVersions) UpdateStatus(ctx context.Context, virtualMachineTemplateVersion *v1beta1.VirtualMachineTemplateVersion, opts v1.UpdateOptions) (*v1beta1.VirtualMachineTemplateVersion, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualmachinetemplateversionsResource, "status", c.ns, virtualMachineTemplateVersion), &v1beta1.VirtualMachineTemplateVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineTemplateVersion), err
}

func (c *FakeVirtualMachineTemplateVersions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualmachinetemplateversionsResource, c.ns, name), &v1beta1.VirtualMachineTemplateVersion{})

	return err
}

func (c *FakeVirtualMachineTemplateVersions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(virtualmachinetemplateversionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.VirtualMachineTemplateVersionList{})
	return err
}

func (c *FakeVirtualMachineTemplateVersions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VirtualMachineTemplateVersion, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualmachinetemplateversionsResource, c.ns, name, pt, data, subresources...), &v1beta1.VirtualMachineTemplateVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineTemplateVersion), err
}
