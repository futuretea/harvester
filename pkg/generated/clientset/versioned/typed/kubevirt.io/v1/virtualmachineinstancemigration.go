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

type VirtualMachineInstanceMigrationsGetter interface {
	VirtualMachineInstanceMigrations(namespace string) VirtualMachineInstanceMigrationInterface
}

type VirtualMachineInstanceMigrationInterface interface {
	Create(ctx context.Context, virtualMachineInstanceMigration *v1.VirtualMachineInstanceMigration, opts metav1.CreateOptions) (*v1.VirtualMachineInstanceMigration, error)
	Update(ctx context.Context, virtualMachineInstanceMigration *v1.VirtualMachineInstanceMigration, opts metav1.UpdateOptions) (*v1.VirtualMachineInstanceMigration, error)
	UpdateStatus(ctx context.Context, virtualMachineInstanceMigration *v1.VirtualMachineInstanceMigration, opts metav1.UpdateOptions) (*v1.VirtualMachineInstanceMigration, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.VirtualMachineInstanceMigration, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.VirtualMachineInstanceMigrationList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VirtualMachineInstanceMigration, err error)
	VirtualMachineInstanceMigrationExpansion
}

type virtualMachineInstanceMigrations struct {
	client	rest.Interface
	ns	string
}

func newVirtualMachineInstanceMigrations(c *KubevirtV1Client, namespace string) *virtualMachineInstanceMigrations {
	__traceStack()

	return &virtualMachineInstanceMigrations{
		client:	c.RESTClient(),
		ns:	namespace,
	}
}

func (c *virtualMachineInstanceMigrations) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.VirtualMachineInstanceMigration, err error) {
	__traceStack()

	result = &v1.VirtualMachineInstanceMigration{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachineinstancemigrations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineInstanceMigrations) List(ctx context.Context, opts metav1.ListOptions) (result *v1.VirtualMachineInstanceMigrationList, err error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.VirtualMachineInstanceMigrationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachineinstancemigrations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineInstanceMigrations) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachineinstancemigrations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *virtualMachineInstanceMigrations) Create(ctx context.Context, virtualMachineInstanceMigration *v1.VirtualMachineInstanceMigration, opts metav1.CreateOptions) (result *v1.VirtualMachineInstanceMigration, err error) {
	__traceStack()

	result = &v1.VirtualMachineInstanceMigration{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualmachineinstancemigrations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineInstanceMigration).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineInstanceMigrations) Update(ctx context.Context, virtualMachineInstanceMigration *v1.VirtualMachineInstanceMigration, opts metav1.UpdateOptions) (result *v1.VirtualMachineInstanceMigration, err error) {
	__traceStack()

	result = &v1.VirtualMachineInstanceMigration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachineinstancemigrations").
		Name(virtualMachineInstanceMigration.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineInstanceMigration).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineInstanceMigrations) UpdateStatus(ctx context.Context, virtualMachineInstanceMigration *v1.VirtualMachineInstanceMigration, opts metav1.UpdateOptions) (result *v1.VirtualMachineInstanceMigration, err error) {
	__traceStack()

	result = &v1.VirtualMachineInstanceMigration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachineinstancemigrations").
		Name(virtualMachineInstanceMigration.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineInstanceMigration).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineInstanceMigrations) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	__traceStack()

	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachineinstancemigrations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *virtualMachineInstanceMigrations) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	__traceStack()

	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachineinstancemigrations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *virtualMachineInstanceMigrations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VirtualMachineInstanceMigration, err error) {
	__traceStack()

	result = &v1.VirtualMachineInstanceMigration{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualmachineinstancemigrations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
