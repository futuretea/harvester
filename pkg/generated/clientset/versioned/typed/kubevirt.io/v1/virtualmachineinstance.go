package v1

import (
	"context"
	"time"

	scheme "github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1 "kubevirt.io/client-go/api/v1"
)

type VirtualMachineInstancesGetter interface {
	VirtualMachineInstances(namespace string) VirtualMachineInstanceInterface
}

type VirtualMachineInstanceInterface interface {
	Create(ctx context.Context, virtualMachineInstance *v1.VirtualMachineInstance, opts metav1.CreateOptions) (*v1.VirtualMachineInstance, error)
	Update(ctx context.Context, virtualMachineInstance *v1.VirtualMachineInstance, opts metav1.UpdateOptions) (*v1.VirtualMachineInstance, error)
	UpdateStatus(ctx context.Context, virtualMachineInstance *v1.VirtualMachineInstance, opts metav1.UpdateOptions) (*v1.VirtualMachineInstance, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.VirtualMachineInstance, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.VirtualMachineInstanceList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VirtualMachineInstance, err error)
	VirtualMachineInstanceExpansion
}

type virtualMachineInstances struct {
	client	rest.Interface
	ns	string
}

func newVirtualMachineInstances(c *KubevirtV1Client, namespace string) *virtualMachineInstances {
	__traceStack()

	return &virtualMachineInstances{
		client:	c.RESTClient(),
		ns:	namespace,
	}
}

func (c *virtualMachineInstances) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.VirtualMachineInstance, err error) {
	__traceStack()

	result = &v1.VirtualMachineInstance{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachineinstances").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineInstances) List(ctx context.Context, opts metav1.ListOptions) (result *v1.VirtualMachineInstanceList, err error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.VirtualMachineInstanceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachineinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineInstances) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachineinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *virtualMachineInstances) Create(ctx context.Context, virtualMachineInstance *v1.VirtualMachineInstance, opts metav1.CreateOptions) (result *v1.VirtualMachineInstance, err error) {
	__traceStack()

	result = &v1.VirtualMachineInstance{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualmachineinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineInstance).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineInstances) Update(ctx context.Context, virtualMachineInstance *v1.VirtualMachineInstance, opts metav1.UpdateOptions) (result *v1.VirtualMachineInstance, err error) {
	__traceStack()

	result = &v1.VirtualMachineInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachineinstances").
		Name(virtualMachineInstance.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineInstance).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineInstances) UpdateStatus(ctx context.Context, virtualMachineInstance *v1.VirtualMachineInstance, opts metav1.UpdateOptions) (result *v1.VirtualMachineInstance, err error) {
	__traceStack()

	result = &v1.VirtualMachineInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachineinstances").
		Name(virtualMachineInstance.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineInstance).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineInstances) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	__traceStack()

	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachineinstances").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *virtualMachineInstances) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	__traceStack()

	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachineinstances").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *virtualMachineInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VirtualMachineInstance, err error) {
	__traceStack()

	result = &v1.VirtualMachineInstance{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualmachineinstances").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
