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

type FakeUpgrades struct {
	Fake	*FakeHarvesterhciV1beta1
	ns	string
}

var upgradesResource = schema.GroupVersionResource{Group: "harvesterhci.io", Version: "v1beta1", Resource: "upgrades"}

var upgradesKind = schema.GroupVersionKind{Group: "harvesterhci.io", Version: "v1beta1", Kind: "Upgrade"}

func (c *FakeUpgrades) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Upgrade, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(upgradesResource, c.ns, name), &v1beta1.Upgrade{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Upgrade), err
}

func (c *FakeUpgrades) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.UpgradeList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(upgradesResource, upgradesKind, c.ns, opts), &v1beta1.UpgradeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.UpgradeList{ListMeta: obj.(*v1beta1.UpgradeList).ListMeta}
	for _, item := range obj.(*v1beta1.UpgradeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeUpgrades) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(upgradesResource, c.ns, opts))

}

func (c *FakeUpgrades) Create(ctx context.Context, upgrade *v1beta1.Upgrade, opts v1.CreateOptions) (result *v1beta1.Upgrade, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(upgradesResource, c.ns, upgrade), &v1beta1.Upgrade{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Upgrade), err
}

func (c *FakeUpgrades) Update(ctx context.Context, upgrade *v1beta1.Upgrade, opts v1.UpdateOptions) (result *v1beta1.Upgrade, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(upgradesResource, c.ns, upgrade), &v1beta1.Upgrade{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Upgrade), err
}

func (c *FakeUpgrades) UpdateStatus(ctx context.Context, upgrade *v1beta1.Upgrade, opts v1.UpdateOptions) (*v1beta1.Upgrade, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(upgradesResource, "status", c.ns, upgrade), &v1beta1.Upgrade{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Upgrade), err
}

func (c *FakeUpgrades) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(upgradesResource, c.ns, name), &v1beta1.Upgrade{})

	return err
}

func (c *FakeUpgrades) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(upgradesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.UpgradeList{})
	return err
}

func (c *FakeUpgrades) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Upgrade, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(upgradesResource, c.ns, name, pt, data, subresources...), &v1beta1.Upgrade{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Upgrade), err
}
