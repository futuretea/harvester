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

type IngressesGetter interface {
	Ingresses(namespace string) IngressInterface
}

type IngressInterface interface {
	Create(ctx context.Context, ingress *v1.Ingress, opts metav1.CreateOptions) (*v1.Ingress, error)
	Update(ctx context.Context, ingress *v1.Ingress, opts metav1.UpdateOptions) (*v1.Ingress, error)
	UpdateStatus(ctx context.Context, ingress *v1.Ingress, opts metav1.UpdateOptions) (*v1.Ingress, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Ingress, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.IngressList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Ingress, err error)
	IngressExpansion
}

type ingresses struct {
	client	rest.Interface
	ns	string
}

func newIngresses(c *NetworkingV1Client, namespace string) *ingresses {
	__traceStack()

	return &ingresses{
		client:	c.RESTClient(),
		ns:	namespace,
	}
}

func (c *ingresses) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Ingress, err error) {
	__traceStack()

	result = &v1.Ingress{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ingresses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *ingresses) List(ctx context.Context, opts metav1.ListOptions) (result *v1.IngressList, err error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.IngressList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ingresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *ingresses) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ingresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *ingresses) Create(ctx context.Context, ingress *v1.Ingress, opts metav1.CreateOptions) (result *v1.Ingress, err error) {
	__traceStack()

	result = &v1.Ingress{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ingresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ingress).
		Do(ctx).
		Into(result)
	return
}

func (c *ingresses) Update(ctx context.Context, ingress *v1.Ingress, opts metav1.UpdateOptions) (result *v1.Ingress, err error) {
	__traceStack()

	result = &v1.Ingress{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ingresses").
		Name(ingress.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ingress).
		Do(ctx).
		Into(result)
	return
}

func (c *ingresses) UpdateStatus(ctx context.Context, ingress *v1.Ingress, opts metav1.UpdateOptions) (result *v1.Ingress, err error) {
	__traceStack()

	result = &v1.Ingress{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ingresses").
		Name(ingress.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ingress).
		Do(ctx).
		Into(result)
	return
}

func (c *ingresses) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	__traceStack()

	return c.client.Delete().
		Namespace(c.ns).
		Resource("ingresses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *ingresses) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	__traceStack()

	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ingresses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *ingresses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Ingress, err error) {
	__traceStack()

	result = &v1.Ingress{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ingresses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
