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

type FakeSettings struct {
	Fake *FakeHarvesterhciV1beta1
}

var settingsResource = schema.GroupVersionResource{Group: "harvesterhci.io", Version: "v1beta1", Resource: "settings"}

var settingsKind = schema.GroupVersionKind{Group: "harvesterhci.io", Version: "v1beta1", Kind: "Setting"}

func (c *FakeSettings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Setting, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(settingsResource, name), &v1beta1.Setting{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Setting), err
}

func (c *FakeSettings) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.SettingList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(settingsResource, settingsKind, opts), &v1beta1.SettingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.SettingList{ListMeta: obj.(*v1beta1.SettingList).ListMeta}
	for _, item := range obj.(*v1beta1.SettingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeSettings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(settingsResource, opts))
}

func (c *FakeSettings) Create(ctx context.Context, setting *v1beta1.Setting, opts v1.CreateOptions) (result *v1beta1.Setting, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(settingsResource, setting), &v1beta1.Setting{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Setting), err
}

func (c *FakeSettings) Update(ctx context.Context, setting *v1beta1.Setting, opts v1.UpdateOptions) (result *v1beta1.Setting, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(settingsResource, setting), &v1beta1.Setting{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Setting), err
}

func (c *FakeSettings) UpdateStatus(ctx context.Context, setting *v1beta1.Setting, opts v1.UpdateOptions) (*v1beta1.Setting, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(settingsResource, "status", setting), &v1beta1.Setting{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Setting), err
}

func (c *FakeSettings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(settingsResource, name), &v1beta1.Setting{})
	return err
}

func (c *FakeSettings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewRootDeleteCollectionAction(settingsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.SettingList{})
	return err
}

func (c *FakeSettings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Setting, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(settingsResource, name, pt, data, subresources...), &v1beta1.Setting{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Setting), err
}
