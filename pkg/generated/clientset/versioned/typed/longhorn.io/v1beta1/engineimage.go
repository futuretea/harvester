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

type EngineImagesGetter interface {
	EngineImages(namespace string) EngineImageInterface
}

type EngineImageInterface interface {
	Create(ctx context.Context, engineImage *v1beta1.EngineImage, opts v1.CreateOptions) (*v1beta1.EngineImage, error)
	Update(ctx context.Context, engineImage *v1beta1.EngineImage, opts v1.UpdateOptions) (*v1beta1.EngineImage, error)
	UpdateStatus(ctx context.Context, engineImage *v1beta1.EngineImage, opts v1.UpdateOptions) (*v1beta1.EngineImage, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.EngineImage, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.EngineImageList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.EngineImage, err error)
	EngineImageExpansion
}

type engineImages struct {
	client	rest.Interface
	ns	string
}

func newEngineImages(c *LonghornV1beta1Client, namespace string) *engineImages {
	__traceStack()

	return &engineImages{
		client:	c.RESTClient(),
		ns:	namespace,
	}
}

func (c *engineImages) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.EngineImage, err error) {
	__traceStack()

	result = &v1beta1.EngineImage{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("engineimages").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *engineImages) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.EngineImageList, err error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.EngineImageList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("engineimages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *engineImages) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("engineimages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *engineImages) Create(ctx context.Context, engineImage *v1beta1.EngineImage, opts v1.CreateOptions) (result *v1beta1.EngineImage, err error) {
	__traceStack()

	result = &v1beta1.EngineImage{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("engineimages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(engineImage).
		Do(ctx).
		Into(result)
	return
}

func (c *engineImages) Update(ctx context.Context, engineImage *v1beta1.EngineImage, opts v1.UpdateOptions) (result *v1beta1.EngineImage, err error) {
	__traceStack()

	result = &v1beta1.EngineImage{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("engineimages").
		Name(engineImage.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(engineImage).
		Do(ctx).
		Into(result)
	return
}

func (c *engineImages) UpdateStatus(ctx context.Context, engineImage *v1beta1.EngineImage, opts v1.UpdateOptions) (result *v1beta1.EngineImage, err error) {
	__traceStack()

	result = &v1beta1.EngineImage{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("engineimages").
		Name(engineImage.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(engineImage).
		Do(ctx).
		Into(result)
	return
}

func (c *engineImages) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	return c.client.Delete().
		Namespace(c.ns).
		Resource("engineimages").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *engineImages) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("engineimages").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *engineImages) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.EngineImage, err error) {
	__traceStack()

	result = &v1beta1.EngineImage{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("engineimages").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
