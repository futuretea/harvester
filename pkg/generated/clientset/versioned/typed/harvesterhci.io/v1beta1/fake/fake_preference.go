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

type FakePreferences struct {
	Fake	*FakeHarvesterhciV1beta1
	ns	string
}

var preferencesResource = schema.GroupVersionResource{Group: "harvesterhci.io", Version: "v1beta1", Resource: "preferences"}

var preferencesKind = schema.GroupVersionKind{Group: "harvesterhci.io", Version: "v1beta1", Kind: "Preference"}

func (c *FakePreferences) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Preference, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(preferencesResource, c.ns, name), &v1beta1.Preference{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Preference), err
}

func (c *FakePreferences) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.PreferenceList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(preferencesResource, preferencesKind, c.ns, opts), &v1beta1.PreferenceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.PreferenceList{ListMeta: obj.(*v1beta1.PreferenceList).ListMeta}
	for _, item := range obj.(*v1beta1.PreferenceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakePreferences) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(preferencesResource, c.ns, opts))

}

func (c *FakePreferences) Create(ctx context.Context, preference *v1beta1.Preference, opts v1.CreateOptions) (result *v1beta1.Preference, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(preferencesResource, c.ns, preference), &v1beta1.Preference{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Preference), err
}

func (c *FakePreferences) Update(ctx context.Context, preference *v1beta1.Preference, opts v1.UpdateOptions) (result *v1beta1.Preference, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(preferencesResource, c.ns, preference), &v1beta1.Preference{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Preference), err
}

func (c *FakePreferences) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(preferencesResource, c.ns, name), &v1beta1.Preference{})

	return err
}

func (c *FakePreferences) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(preferencesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.PreferenceList{})
	return err
}

func (c *FakePreferences) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Preference, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(preferencesResource, c.ns, name, pt, data, subresources...), &v1beta1.Preference{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Preference), err
}
