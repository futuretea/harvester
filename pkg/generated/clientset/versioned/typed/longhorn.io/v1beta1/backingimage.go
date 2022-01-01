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

type BackingImagesGetter interface {
	BackingImages(namespace string) BackingImageInterface
}

type BackingImageInterface interface {
	Create(ctx context.Context, backingImage *v1beta1.BackingImage, opts v1.CreateOptions) (*v1beta1.BackingImage, error)
	Update(ctx context.Context, backingImage *v1beta1.BackingImage, opts v1.UpdateOptions) (*v1beta1.BackingImage, error)
	UpdateStatus(ctx context.Context, backingImage *v1beta1.BackingImage, opts v1.UpdateOptions) (*v1beta1.BackingImage, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.BackingImage, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.BackingImageList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BackingImage, err error)
	BackingImageExpansion
}

type backingImages struct {
	client	rest.Interface
	ns	string
}

func newBackingImages(c *LonghornV1beta1Client, namespace string) *backingImages {
	__traceStack()

	return &backingImages{
		client:	c.RESTClient(),
		ns:	namespace,
	}
}

func (c *backingImages) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.BackingImage, err error) {
	__traceStack()

	result = &v1beta1.BackingImage{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backingimages").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *backingImages) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.BackingImageList, err error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.BackingImageList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backingimages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *backingImages) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("backingimages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *backingImages) Create(ctx context.Context, backingImage *v1beta1.BackingImage, opts v1.CreateOptions) (result *v1beta1.BackingImage, err error) {
	__traceStack()

	result = &v1beta1.BackingImage{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("backingimages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backingImage).
		Do(ctx).
		Into(result)
	return
}

func (c *backingImages) Update(ctx context.Context, backingImage *v1beta1.BackingImage, opts v1.UpdateOptions) (result *v1beta1.BackingImage, err error) {
	__traceStack()

	result = &v1beta1.BackingImage{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backingimages").
		Name(backingImage.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backingImage).
		Do(ctx).
		Into(result)
	return
}

func (c *backingImages) UpdateStatus(ctx context.Context, backingImage *v1beta1.BackingImage, opts v1.UpdateOptions) (result *v1beta1.BackingImage, err error) {
	__traceStack()

	result = &v1beta1.BackingImage{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backingimages").
		Name(backingImage.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backingImage).
		Do(ctx).
		Into(result)
	return
}

func (c *backingImages) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	return c.client.Delete().
		Namespace(c.ns).
		Resource("backingimages").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *backingImages) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backingimages").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *backingImages) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BackingImage, err error) {
	__traceStack()

	result = &v1beta1.BackingImage{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("backingimages").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
