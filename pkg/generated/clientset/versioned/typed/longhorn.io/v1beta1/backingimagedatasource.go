package v1beta1

import (
	"context"
	"time"

	scheme "github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v1beta1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

type BackingImageDataSourcesGetter interface {
	BackingImageDataSources(namespace string) BackingImageDataSourceInterface
}

type BackingImageDataSourceInterface interface {
	Create(ctx context.Context, backingImageDataSource *v1beta1.BackingImageDataSource, opts v1.CreateOptions) (*v1beta1.BackingImageDataSource, error)
	Update(ctx context.Context, backingImageDataSource *v1beta1.BackingImageDataSource, opts v1.UpdateOptions) (*v1beta1.BackingImageDataSource, error)
	UpdateStatus(ctx context.Context, backingImageDataSource *v1beta1.BackingImageDataSource, opts v1.UpdateOptions) (*v1beta1.BackingImageDataSource, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.BackingImageDataSource, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.BackingImageDataSourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BackingImageDataSource, err error)
	BackingImageDataSourceExpansion
}

type backingImageDataSources struct {
	client	rest.Interface
	ns	string
}

func newBackingImageDataSources(c *LonghornV1beta1Client, namespace string) *backingImageDataSources {
	__traceStack()

	return &backingImageDataSources{
		client:	c.RESTClient(),
		ns:	namespace,
	}
}

func (c *backingImageDataSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.BackingImageDataSource, err error) {
	__traceStack()

	result = &v1beta1.BackingImageDataSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backingimagedatasources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *backingImageDataSources) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.BackingImageDataSourceList, err error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.BackingImageDataSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backingimagedatasources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *backingImageDataSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("backingimagedatasources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *backingImageDataSources) Create(ctx context.Context, backingImageDataSource *v1beta1.BackingImageDataSource, opts v1.CreateOptions) (result *v1beta1.BackingImageDataSource, err error) {
	__traceStack()

	result = &v1beta1.BackingImageDataSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("backingimagedatasources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backingImageDataSource).
		Do(ctx).
		Into(result)
	return
}

func (c *backingImageDataSources) Update(ctx context.Context, backingImageDataSource *v1beta1.BackingImageDataSource, opts v1.UpdateOptions) (result *v1beta1.BackingImageDataSource, err error) {
	__traceStack()

	result = &v1beta1.BackingImageDataSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backingimagedatasources").
		Name(backingImageDataSource.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backingImageDataSource).
		Do(ctx).
		Into(result)
	return
}

func (c *backingImageDataSources) UpdateStatus(ctx context.Context, backingImageDataSource *v1beta1.BackingImageDataSource, opts v1.UpdateOptions) (result *v1beta1.BackingImageDataSource, err error) {
	__traceStack()

	result = &v1beta1.BackingImageDataSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backingimagedatasources").
		Name(backingImageDataSource.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backingImageDataSource).
		Do(ctx).
		Into(result)
	return
}

func (c *backingImageDataSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	return c.client.Delete().
		Namespace(c.ns).
		Resource("backingimagedatasources").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *backingImageDataSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backingimagedatasources").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *backingImageDataSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BackingImageDataSource, err error) {
	__traceStack()

	result = &v1beta1.BackingImageDataSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("backingimagedatasources").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
