package v1

import (
	"context"
	"time"

	scheme "github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

type NetworkPoliciesGetter interface {
	NetworkPolicies(namespace string) NetworkPolicyInterface
}

type NetworkPolicyInterface interface {
	Create(ctx context.Context, networkPolicy *v1.NetworkPolicy, opts metav1.CreateOptions) (*v1.NetworkPolicy, error)
	Update(ctx context.Context, networkPolicy *v1.NetworkPolicy, opts metav1.UpdateOptions) (*v1.NetworkPolicy, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.NetworkPolicy, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.NetworkPolicyList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.NetworkPolicy, err error)
	NetworkPolicyExpansion
}

type networkPolicies struct {
	client	rest.Interface
	ns	string
}

func newNetworkPolicies(c *NetworkingV1Client, namespace string) *networkPolicies {
	__traceStack()

	return &networkPolicies{
		client:	c.RESTClient(),
		ns:	namespace,
	}
}

func (c *networkPolicies) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.NetworkPolicy, err error) {
	__traceStack()

	result = &v1.NetworkPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *networkPolicies) List(ctx context.Context, opts metav1.ListOptions) (result *v1.NetworkPolicyList, err error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.NetworkPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *networkPolicies) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("networkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *networkPolicies) Create(ctx context.Context, networkPolicy *v1.NetworkPolicy, opts metav1.CreateOptions) (result *v1.NetworkPolicy, err error) {
	__traceStack()

	result = &v1.NetworkPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("networkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkPolicy).
		Do(ctx).
		Into(result)
	return
}

func (c *networkPolicies) Update(ctx context.Context, networkPolicy *v1.NetworkPolicy, opts metav1.UpdateOptions) (result *v1.NetworkPolicy, err error) {
	__traceStack()

	result = &v1.NetworkPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkpolicies").
		Name(networkPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkPolicy).
		Do(ctx).
		Into(result)
	return
}

func (c *networkPolicies) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	__traceStack()

	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkpolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *networkPolicies) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	__traceStack()

	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkpolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *networkPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.NetworkPolicy, err error) {
	__traceStack()

	result = &v1.NetworkPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("networkpolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
