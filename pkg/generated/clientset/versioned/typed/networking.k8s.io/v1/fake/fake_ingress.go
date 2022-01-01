package fake

import (
	"context"

	networkingv1 "k8s.io/api/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

type FakeIngresses struct {
	Fake	*FakeNetworkingV1
	ns	string
}

var ingressesResource = schema.GroupVersionResource{Group: "networking.k8s.io", Version: "v1", Resource: "ingresses"}

var ingressesKind = schema.GroupVersionKind{Group: "networking.k8s.io", Version: "v1", Kind: "Ingress"}

func (c *FakeIngresses) Get(ctx context.Context, name string, options v1.GetOptions) (result *networkingv1.Ingress, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ingressesResource, c.ns, name), &networkingv1.Ingress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.Ingress), err
}

func (c *FakeIngresses) List(ctx context.Context, opts v1.ListOptions) (result *networkingv1.IngressList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(ingressesResource, ingressesKind, c.ns, opts), &networkingv1.IngressList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &networkingv1.IngressList{ListMeta: obj.(*networkingv1.IngressList).ListMeta}
	for _, item := range obj.(*networkingv1.IngressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeIngresses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ingressesResource, c.ns, opts))

}

func (c *FakeIngresses) Create(ctx context.Context, ingress *networkingv1.Ingress, opts v1.CreateOptions) (result *networkingv1.Ingress, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ingressesResource, c.ns, ingress), &networkingv1.Ingress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.Ingress), err
}

func (c *FakeIngresses) Update(ctx context.Context, ingress *networkingv1.Ingress, opts v1.UpdateOptions) (result *networkingv1.Ingress, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ingressesResource, c.ns, ingress), &networkingv1.Ingress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.Ingress), err
}

func (c *FakeIngresses) UpdateStatus(ctx context.Context, ingress *networkingv1.Ingress, opts v1.UpdateOptions) (*networkingv1.Ingress, error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ingressesResource, "status", c.ns, ingress), &networkingv1.Ingress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.Ingress), err
}

func (c *FakeIngresses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ingressesResource, c.ns, name), &networkingv1.Ingress{})

	return err
}

func (c *FakeIngresses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(ingressesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &networkingv1.IngressList{})
	return err
}

func (c *FakeIngresses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *networkingv1.Ingress, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ingressesResource, c.ns, name, pt, data, subresources...), &networkingv1.Ingress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.Ingress), err
}
