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

type FakeKeyPairs struct {
	Fake	*FakeHarvesterhciV1beta1
	ns	string
}

var keypairsResource = schema.GroupVersionResource{Group: "harvesterhci.io", Version: "v1beta1", Resource: "keypairs"}

var keypairsKind = schema.GroupVersionKind{Group: "harvesterhci.io", Version: "v1beta1", Kind: "KeyPair"}

func (c *FakeKeyPairs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.KeyPair, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(keypairsResource, c.ns, name), &v1beta1.KeyPair{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.KeyPair), err
}

func (c *FakeKeyPairs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.KeyPairList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(keypairsResource, keypairsKind, c.ns, opts), &v1beta1.KeyPairList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.KeyPairList{ListMeta: obj.(*v1beta1.KeyPairList).ListMeta}
	for _, item := range obj.(*v1beta1.KeyPairList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeKeyPairs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(keypairsResource, c.ns, opts))

}

func (c *FakeKeyPairs) Create(ctx context.Context, keyPair *v1beta1.KeyPair, opts v1.CreateOptions) (result *v1beta1.KeyPair, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(keypairsResource, c.ns, keyPair), &v1beta1.KeyPair{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.KeyPair), err
}

func (c *FakeKeyPairs) Update(ctx context.Context, keyPair *v1beta1.KeyPair, opts v1.UpdateOptions) (result *v1beta1.KeyPair, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(keypairsResource, c.ns, keyPair), &v1beta1.KeyPair{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.KeyPair), err
}

func (c *FakeKeyPairs) UpdateStatus(ctx context.Context, keyPair *v1beta1.KeyPair, opts v1.UpdateOptions) (*v1beta1.KeyPair, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(keypairsResource, "status", c.ns, keyPair), &v1beta1.KeyPair{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.KeyPair), err
}

func (c *FakeKeyPairs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(keypairsResource, c.ns, name), &v1beta1.KeyPair{})

	return err
}

func (c *FakeKeyPairs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(keypairsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.KeyPairList{})
	return err
}

func (c *FakeKeyPairs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.KeyPair, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(keypairsResource, c.ns, name, pt, data, subresources...), &v1beta1.KeyPair{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.KeyPair), err
}
