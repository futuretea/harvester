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

type FakeNetworkPolicies struct {
	Fake	*FakeNetworkingV1
	ns	string
}

var networkpoliciesResource = schema.GroupVersionResource{Group: "networking.k8s.io", Version: "v1", Resource: "networkpolicies"}

var networkpoliciesKind = schema.GroupVersionKind{Group: "networking.k8s.io", Version: "v1", Kind: "NetworkPolicy"}

func (c *FakeNetworkPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *networkingv1.NetworkPolicy, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkpoliciesResource, c.ns, name), &networkingv1.NetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.NetworkPolicy), err
}

func (c *FakeNetworkPolicies) List(ctx context.Context, opts v1.ListOptions) (result *networkingv1.NetworkPolicyList, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkpoliciesResource, networkpoliciesKind, c.ns, opts), &networkingv1.NetworkPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &networkingv1.NetworkPolicyList{ListMeta: obj.(*networkingv1.NetworkPolicyList).ListMeta}
	for _, item := range obj.(*networkingv1.NetworkPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeNetworkPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkpoliciesResource, c.ns, opts))

}

func (c *FakeNetworkPolicies) Create(ctx context.Context, networkPolicy *networkingv1.NetworkPolicy, opts v1.CreateOptions) (result *networkingv1.NetworkPolicy, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkpoliciesResource, c.ns, networkPolicy), &networkingv1.NetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.NetworkPolicy), err
}

func (c *FakeNetworkPolicies) Update(ctx context.Context, networkPolicy *networkingv1.NetworkPolicy, opts v1.UpdateOptions) (result *networkingv1.NetworkPolicy, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkpoliciesResource, c.ns, networkPolicy), &networkingv1.NetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.NetworkPolicy), err
}

func (c *FakeNetworkPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networkpoliciesResource, c.ns, name), &networkingv1.NetworkPolicy{})

	return err
}

func (c *FakeNetworkPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	action := testing.NewDeleteCollectionAction(networkpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &networkingv1.NetworkPolicyList{})
	return err
}

func (c *FakeNetworkPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *networkingv1.NetworkPolicy, err error) {
	__traceStack()

	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkpoliciesResource, c.ns, name, pt, data, subresources...), &networkingv1.NetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*networkingv1.NetworkPolicy), err
}
