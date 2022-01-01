package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha4 "sigs.k8s.io/cluster-api/api/v1alpha4"
)

type FakeMachines struct {
	Fake	*FakeClusterV1alpha4
	ns	string
}

var machinesResource = schema.GroupVersionResource{Group: "cluster.x-k8s.io", Version: "v1alpha4", Resource: "machines"}

var machinesKind = schema.GroupVersionKind{Group: "cluster.x-k8s.io", Version: "v1alpha4", Kind: "Machine"}

func (c *FakeMachines) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha4.Machine, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(machinesResource, c.ns, name), &v1alpha4.Machine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha4.Machine), err
}

func (c *FakeMachines) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha4.MachineList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(machinesResource, machinesKind, c.ns, opts), &v1alpha4.MachineList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha4.MachineList{ListMeta: obj.(*v1alpha4.MachineList).ListMeta}
	for _, item := range obj.(*v1alpha4.MachineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeMachines) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(machinesResource, c.ns, opts))

}

func (c *FakeMachines) Create(ctx context.Context, machine *v1alpha4.Machine, opts v1.CreateOptions) (result *v1alpha4.Machine, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(machinesResource, c.ns, machine), &v1alpha4.Machine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha4.Machine), err
}

func (c *FakeMachines) Update(ctx context.Context, machine *v1alpha4.Machine, opts v1.UpdateOptions) (result *v1alpha4.Machine, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(machinesResource, c.ns, machine), &v1alpha4.Machine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha4.Machine), err
}

func (c *FakeMachines) UpdateStatus(ctx context.Context, machine *v1alpha4.Machine, opts v1.UpdateOptions) (*v1alpha4.Machine, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(machinesResource, "status", c.ns, machine), &v1alpha4.Machine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha4.Machine), err
}

func (c *FakeMachines) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(machinesResource, c.ns, name), &v1alpha4.Machine{})

	return err
}

func (c *FakeMachines) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(machinesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha4.MachineList{})
	return err
}

func (c *FakeMachines) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha4.Machine, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(machinesResource, c.ns, name, pt, data, subresources...), &v1alpha4.Machine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha4.Machine), err
}
