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

type VirtualMachinesGetter interface {
	VirtualMachines(namespace string) VirtualMachineInterface
}

type VirtualMachineInterface interface {
	Create(ctx context.Context, virtualMachine *v1.VirtualMachine, opts metav1.CreateOptions) (*v1.VirtualMachine, error)
	Update(ctx context.Context, virtualMachine *v1.VirtualMachine, opts metav1.UpdateOptions) (*v1.VirtualMachine, error)
	UpdateStatus(ctx context.Context, virtualMachine *v1.VirtualMachine, opts metav1.UpdateOptions) (*v1.VirtualMachine, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.VirtualMachine, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.VirtualMachineList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VirtualMachine, err error)
	VirtualMachineExpansion
}

type virtualMachines struct {
	client	rest.Interface
	ns	string
}

func newVirtualMachines(c *KubevirtV1Client, namespace string) *virtualMachines {
	__traceStack()

	return &virtualMachines{
		client:	c.RESTClient(),
		ns:	namespace,
	}
}

func (c *virtualMachines) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.VirtualMachine, err error) {
	__traceStack()

	result = &v1.VirtualMachine{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachines").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachines) List(ctx context.Context, opts metav1.ListOptions) (result *v1.VirtualMachineList, err error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.VirtualMachineList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachines) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *virtualMachines) Create(ctx context.Context, virtualMachine *v1.VirtualMachine, opts metav1.CreateOptions) (result *v1.VirtualMachine, err error) {
	__traceStack()

	result = &v1.VirtualMachine{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualmachines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachine).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachines) Update(ctx context.Context, virtualMachine *v1.VirtualMachine, opts metav1.UpdateOptions) (result *v1.VirtualMachine, err error) {
	__traceStack()

	result = &v1.VirtualMachine{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachines").
		Name(virtualMachine.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachine).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachines) UpdateStatus(ctx context.Context, virtualMachine *v1.VirtualMachine, opts metav1.UpdateOptions) (result *v1.VirtualMachine, err error) {
	__traceStack()

	result = &v1.VirtualMachine{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachines").
		Name(virtualMachine.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachine).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachines) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	__traceStack()

	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachines").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *virtualMachines) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	__traceStack()

	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachines").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *virtualMachines) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VirtualMachine, err error) {
	__traceStack()

	result = &v1.VirtualMachine{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualmachines").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
