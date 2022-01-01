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

type FakeVersions struct {
	Fake	*FakeHarvesterhciV1beta1
	ns	string
}

var versionsResource = schema.GroupVersionResource{Group: "harvesterhci.io", Version: "v1beta1", Resource: "versions"}

var versionsKind = schema.GroupVersionKind{Group: "harvesterhci.io", Version: "v1beta1", Kind: "Version"}

func (c *FakeVersions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Version, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(versionsResource, c.ns, name), &v1beta1.Version{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Version), err
}

func (c *FakeVersions) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VersionList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(versionsResource, versionsKind, c.ns, opts), &v1beta1.VersionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.VersionList{ListMeta: obj.(*v1beta1.VersionList).ListMeta}
	for _, item := range obj.(*v1beta1.VersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeVersions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(versionsResource, c.ns, opts))

}

func (c *FakeVersions) Create(ctx context.Context, version *v1beta1.Version, opts v1.CreateOptions) (result *v1beta1.Version, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(versionsResource, c.ns, version), &v1beta1.Version{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Version), err
}

func (c *FakeVersions) Update(ctx context.Context, version *v1beta1.Version, opts v1.UpdateOptions) (result *v1beta1.Version, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(versionsResource, c.ns, version), &v1beta1.Version{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Version), err
}

func (c *FakeVersions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(versionsResource, c.ns, name), &v1beta1.Version{})

	return err
}

func (c *FakeVersions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(versionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.VersionList{})
	return err
}

func (c *FakeVersions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Version, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(versionsResource, c.ns, name, pt, data, subresources...), &v1beta1.Version{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Version), err
}
