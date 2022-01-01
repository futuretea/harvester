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

type VirtualMachineRestoresGetter interface {
	VirtualMachineRestores(namespace string) VirtualMachineRestoreInterface
}

type VirtualMachineRestoreInterface interface {
	Create(ctx context.Context, virtualMachineRestore *v1beta1.VirtualMachineRestore, opts v1.CreateOptions) (*v1beta1.VirtualMachineRestore, error)
	Update(ctx context.Context, virtualMachineRestore *v1beta1.VirtualMachineRestore, opts v1.UpdateOptions) (*v1beta1.VirtualMachineRestore, error)
	UpdateStatus(ctx context.Context, virtualMachineRestore *v1beta1.VirtualMachineRestore, opts v1.UpdateOptions) (*v1beta1.VirtualMachineRestore, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.VirtualMachineRestore, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.VirtualMachineRestoreList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VirtualMachineRestore, err error)
	VirtualMachineRestoreExpansion
}

type virtualMachineRestores struct {
	client	rest.Interface
	ns	string
}

func newVirtualMachineRestores(c *HarvesterhciV1beta1Client, namespace string) *virtualMachineRestores {
	__traceStack()

	return &virtualMachineRestores{
		client:	c.RESTClient(),
		ns:	namespace,
	}
}

func (c *virtualMachineRestores) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VirtualMachineRestore, err error) {
	__traceStack()

	result = &v1beta1.VirtualMachineRestore{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinerestores").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineRestores) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VirtualMachineRestoreList, err error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.VirtualMachineRestoreList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinerestores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineRestores) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinerestores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *virtualMachineRestores) Create(ctx context.Context, virtualMachineRestore *v1beta1.VirtualMachineRestore, opts v1.CreateOptions) (result *v1beta1.VirtualMachineRestore, err error) {
	__traceStack()

	result = &v1beta1.VirtualMachineRestore{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualmachinerestores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineRestore).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineRestores) Update(ctx context.Context, virtualMachineRestore *v1beta1.VirtualMachineRestore, opts v1.UpdateOptions) (result *v1beta1.VirtualMachineRestore, err error) {
	__traceStack()

	result = &v1beta1.VirtualMachineRestore{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachinerestores").
		Name(virtualMachineRestore.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineRestore).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineRestores) UpdateStatus(ctx context.Context, virtualMachineRestore *v1beta1.VirtualMachineRestore, opts v1.UpdateOptions) (result *v1beta1.VirtualMachineRestore, err error) {
	__traceStack()

	result = &v1beta1.VirtualMachineRestore{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachinerestores").
		Name(virtualMachineRestore.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineRestore).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineRestores) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinerestores").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *virtualMachineRestores) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinerestores").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *virtualMachineRestores) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VirtualMachineRestore, err error) {
	__traceStack()

	result = &v1beta1.VirtualMachineRestore{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualmachinerestores").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
