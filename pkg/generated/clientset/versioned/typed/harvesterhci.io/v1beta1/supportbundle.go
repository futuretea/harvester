package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	scheme "github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

type SupportBundlesGetter interface {
	SupportBundles(namespace string) SupportBundleInterface
}

type SupportBundleInterface interface {
	Create(ctx context.Context, supportBundle *v1beta1.SupportBundle, opts v1.CreateOptions) (*v1beta1.SupportBundle, error)
	Update(ctx context.Context, supportBundle *v1beta1.SupportBundle, opts v1.UpdateOptions) (*v1beta1.SupportBundle, error)
	UpdateStatus(ctx context.Context, supportBundle *v1beta1.SupportBundle, opts v1.UpdateOptions) (*v1beta1.SupportBundle, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.SupportBundle, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.SupportBundleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.SupportBundle, err error)
	SupportBundleExpansion
}

type supportBundles struct {
	client	rest.Interface
	ns	string
}

func newSupportBundles(c *HarvesterhciV1beta1Client, namespace string) *supportBundles {
	__traceStack()

	return &supportBundles{
		client:	c.RESTClient(),
		ns:	namespace,
	}
}

func (c *supportBundles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.SupportBundle, err error) {
	__traceStack()

	result = &v1beta1.SupportBundle{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("supportbundles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *supportBundles) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.SupportBundleList, err error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.SupportBundleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("supportbundles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *supportBundles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("supportbundles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *supportBundles) Create(ctx context.Context, supportBundle *v1beta1.SupportBundle, opts v1.CreateOptions) (result *v1beta1.SupportBundle, err error) {
	__traceStack()

	result = &v1beta1.SupportBundle{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("supportbundles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(supportBundle).
		Do(ctx).
		Into(result)
	return
}

func (c *supportBundles) Update(ctx context.Context, supportBundle *v1beta1.SupportBundle, opts v1.UpdateOptions) (result *v1beta1.SupportBundle, err error) {
	__traceStack()

	result = &v1beta1.SupportBundle{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("supportbundles").
		Name(supportBundle.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(supportBundle).
		Do(ctx).
		Into(result)
	return
}

func (c *supportBundles) UpdateStatus(ctx context.Context, supportBundle *v1beta1.SupportBundle, opts v1.UpdateOptions) (result *v1beta1.SupportBundle, err error) {
	__traceStack()

	result = &v1beta1.SupportBundle{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("supportbundles").
		Name(supportBundle.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(supportBundle).
		Do(ctx).
		Into(result)
	return
}

func (c *supportBundles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	return c.client.Delete().
		Namespace(c.ns).
		Resource("supportbundles").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *supportBundles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("supportbundles").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *supportBundles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.SupportBundle, err error) {
	__traceStack()

	result = &v1beta1.SupportBundle{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("supportbundles").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
