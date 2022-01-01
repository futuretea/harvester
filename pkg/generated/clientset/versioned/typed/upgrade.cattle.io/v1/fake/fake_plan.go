package fake

import (
	"context"

	upgradecattleiov1 "github.com/rancher/system-upgrade-controller/pkg/apis/upgrade.cattle.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

type FakePlans struct {
	Fake	*FakeUpgradeV1
	ns	string
}

var plansResource = schema.GroupVersionResource{Group: "upgrade.cattle.io", Version: "v1", Resource: "plans"}

var plansKind = schema.GroupVersionKind{Group: "upgrade.cattle.io", Version: "v1", Kind: "Plan"}

func (c *FakePlans) Get(ctx context.Context, name string, options v1.GetOptions) (result *upgradecattleiov1.Plan, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(plansResource, c.ns, name), &upgradecattleiov1.Plan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*upgradecattleiov1.Plan), err
}

func (c *FakePlans) List(ctx context.Context, opts v1.ListOptions) (result *upgradecattleiov1.PlanList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(plansResource, plansKind, c.ns, opts), &upgradecattleiov1.PlanList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &upgradecattleiov1.PlanList{ListMeta: obj.(*upgradecattleiov1.PlanList).ListMeta}
	for _, item := range obj.(*upgradecattleiov1.PlanList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakePlans) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(plansResource, c.ns, opts))

}

func (c *FakePlans) Create(ctx context.Context, plan *upgradecattleiov1.Plan, opts v1.CreateOptions) (result *upgradecattleiov1.Plan, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(plansResource, c.ns, plan), &upgradecattleiov1.Plan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*upgradecattleiov1.Plan), err
}

func (c *FakePlans) Update(ctx context.Context, plan *upgradecattleiov1.Plan, opts v1.UpdateOptions) (result *upgradecattleiov1.Plan, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(plansResource, c.ns, plan), &upgradecattleiov1.Plan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*upgradecattleiov1.Plan), err
}

func (c *FakePlans) UpdateStatus(ctx context.Context, plan *upgradecattleiov1.Plan, opts v1.UpdateOptions) (*upgradecattleiov1.Plan, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(plansResource, "status", c.ns, plan), &upgradecattleiov1.Plan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*upgradecattleiov1.Plan), err
}

func (c *FakePlans) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(plansResource, c.ns, name), &upgradecattleiov1.Plan{})

	return err
}

func (c *FakePlans) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(plansResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &upgradecattleiov1.PlanList{})
	return err
}

func (c *FakePlans) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *upgradecattleiov1.Plan, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(plansResource, c.ns, name, pt, data, subresources...), &upgradecattleiov1.Plan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*upgradecattleiov1.Plan), err
}
