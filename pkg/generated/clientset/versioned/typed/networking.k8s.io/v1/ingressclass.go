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

type IngressClassesGetter interface {
	IngressClasses() IngressClassInterface
}

type IngressClassInterface interface {
	Create(ctx context.Context, ingressClass *v1.IngressClass, opts metav1.CreateOptions) (*v1.IngressClass, error)
	Update(ctx context.Context, ingressClass *v1.IngressClass, opts metav1.UpdateOptions) (*v1.IngressClass, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.IngressClass, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.IngressClassList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.IngressClass, err error)
	IngressClassExpansion
}

type ingressClasses struct {
	client rest.Interface
}

func newIngressClasses(c *NetworkingV1Client) *ingressClasses {
	__traceStack()

	return &ingressClasses{
		client: c.RESTClient(),
	}
}

func (c *ingressClasses) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.IngressClass, err error) {
	__traceStack()

	result = &v1.IngressClass{}
	err = c.client.Get().
		Resource("ingressclasses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *ingressClasses) List(ctx context.Context, opts metav1.ListOptions) (result *v1.IngressClassList, err error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.IngressClassList{}
	err = c.client.Get().
		Resource("ingressclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *ingressClasses) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("ingressclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *ingressClasses) Create(ctx context.Context, ingressClass *v1.IngressClass, opts metav1.CreateOptions) (result *v1.IngressClass, err error) {
	__traceStack()

	result = &v1.IngressClass{}
	err = c.client.Post().
		Resource("ingressclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ingressClass).
		Do(ctx).
		Into(result)
	return
}

func (c *ingressClasses) Update(ctx context.Context, ingressClass *v1.IngressClass, opts metav1.UpdateOptions) (result *v1.IngressClass, err error) {
	__traceStack()

	result = &v1.IngressClass{}
	err = c.client.Put().
		Resource("ingressclasses").
		Name(ingressClass.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(ingressClass).
		Do(ctx).
		Into(result)
	return
}

func (c *ingressClasses) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	__traceStack()

	return c.client.Delete().
		Resource("ingressclasses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *ingressClasses) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	__traceStack()

	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("ingressclasses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *ingressClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.IngressClass, err error) {
	__traceStack()

	result = &v1.IngressClass{}
	err = c.client.Patch(pt).
		Resource("ingressclasses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
