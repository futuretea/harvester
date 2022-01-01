package fake

import (
	"context"

	v1beta1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

type FakeNodes struct {
	Fake	*FakeLonghornV1beta1
	ns	string
}

var nodesResource = schema.GroupVersionResource{Group: "longhorn.io", Version: "v1beta1", Resource: "nodes"}

var nodesKind = schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta1", Kind: "Node"}

func (c *FakeNodes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Node, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(nodesResource, c.ns, name), &v1beta1.Node{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Node), err
}

func (c *FakeNodes) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.NodeList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(nodesResource, nodesKind, c.ns, opts), &v1beta1.NodeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.NodeList{ListMeta: obj.(*v1beta1.NodeList).ListMeta}
	for _, item := range obj.(*v1beta1.NodeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeNodes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(nodesResource, c.ns, opts))

}

func (c *FakeNodes) Create(ctx context.Context, node *v1beta1.Node, opts v1.CreateOptions) (result *v1beta1.Node, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(nodesResource, c.ns, node), &v1beta1.Node{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Node), err
}

func (c *FakeNodes) Update(ctx context.Context, node *v1beta1.Node, opts v1.UpdateOptions) (result *v1beta1.Node, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(nodesResource, c.ns, node), &v1beta1.Node{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Node), err
}

func (c *FakeNodes) UpdateStatus(ctx context.Context, node *v1beta1.Node, opts v1.UpdateOptions) (*v1beta1.Node, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(nodesResource, "status", c.ns, node), &v1beta1.Node{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Node), err
}

func (c *FakeNodes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(nodesResource, c.ns, name), &v1beta1.Node{})

	return err
}

func (c *FakeNodes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(nodesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.NodeList{})
	return err
}

func (c *FakeNodes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Node, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(nodesResource, c.ns, name, pt, data, subresources...), &v1beta1.Node{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Node), err
}
