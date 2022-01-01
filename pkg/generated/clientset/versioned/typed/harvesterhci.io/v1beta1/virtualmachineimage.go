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

type VirtualMachineImagesGetter interface {
	VirtualMachineImages(namespace string) VirtualMachineImageInterface
}

type VirtualMachineImageInterface interface {
	Create(ctx context.Context, virtualMachineImage *v1beta1.VirtualMachineImage, opts v1.CreateOptions) (*v1beta1.VirtualMachineImage, error)
	Update(ctx context.Context, virtualMachineImage *v1beta1.VirtualMachineImage, opts v1.UpdateOptions) (*v1beta1.VirtualMachineImage, error)
	UpdateStatus(ctx context.Context, virtualMachineImage *v1beta1.VirtualMachineImage, opts v1.UpdateOptions) (*v1beta1.VirtualMachineImage, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.VirtualMachineImage, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.VirtualMachineImageList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VirtualMachineImage, err error)
	VirtualMachineImageExpansion
}

type virtualMachineImages struct {
	client	rest.Interface
	ns	string
}

func newVirtualMachineImages(c *HarvesterhciV1beta1Client, namespace string) *virtualMachineImages {
	__traceStack()

	return &virtualMachineImages{
		client:	c.RESTClient(),
		ns:	namespace,
	}
}

func (c *virtualMachineImages) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VirtualMachineImage, err error) {
	__traceStack()

	result = &v1beta1.VirtualMachineImage{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachineimages").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineImages) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VirtualMachineImageList, err error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.VirtualMachineImageList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachineimages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineImages) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachineimages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *virtualMachineImages) Create(ctx context.Context, virtualMachineImage *v1beta1.VirtualMachineImage, opts v1.CreateOptions) (result *v1beta1.VirtualMachineImage, err error) {
	__traceStack()

	result = &v1beta1.VirtualMachineImage{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualmachineimages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineImage).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineImages) Update(ctx context.Context, virtualMachineImage *v1beta1.VirtualMachineImage, opts v1.UpdateOptions) (result *v1beta1.VirtualMachineImage, err error) {
	__traceStack()

	result = &v1beta1.VirtualMachineImage{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachineimages").
		Name(virtualMachineImage.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineImage).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineImages) UpdateStatus(ctx context.Context, virtualMachineImage *v1beta1.VirtualMachineImage, opts v1.UpdateOptions) (result *v1beta1.VirtualMachineImage, err error) {
	__traceStack()

	result = &v1beta1.VirtualMachineImage{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachineimages").
		Name(virtualMachineImage.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineImage).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineImages) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachineimages").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *virtualMachineImages) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachineimages").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *virtualMachineImages) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VirtualMachineImage, err error) {
	__traceStack()

	result = &v1beta1.VirtualMachineImage{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualmachineimages").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
