package v1alpha4

import (
	"context"
	"time"

	scheme "github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha4 "sigs.k8s.io/cluster-api/api/v1alpha4"
)

type MachinesGetter interface {
	Machines(namespace string) MachineInterface
}

type MachineInterface interface {
	Create(ctx context.Context, machine *v1alpha4.Machine, opts v1.CreateOptions) (*v1alpha4.Machine, error)
	Update(ctx context.Context, machine *v1alpha4.Machine, opts v1.UpdateOptions) (*v1alpha4.Machine, error)
	UpdateStatus(ctx context.Context, machine *v1alpha4.Machine, opts v1.UpdateOptions) (*v1alpha4.Machine, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha4.Machine, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha4.MachineList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha4.Machine, err error)
	MachineExpansion
}

type machines struct {
	client	rest.Interface
	ns	string
}

func newMachines(c *ClusterV1alpha4Client, namespace string) *machines {
	__traceStack()

	return &machines{
		client:	c.RESTClient(),
		ns:	namespace,
	}
}

func (c *machines) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha4.Machine, err error) {
	__traceStack()

	result = &v1alpha4.Machine{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("machines").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *machines) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha4.MachineList, err error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha4.MachineList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("machines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *machines) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("machines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *machines) Create(ctx context.Context, machine *v1alpha4.Machine, opts v1.CreateOptions) (result *v1alpha4.Machine, err error) {
	__traceStack()

	result = &v1alpha4.Machine{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("machines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(machine).
		Do(ctx).
		Into(result)
	return
}

func (c *machines) Update(ctx context.Context, machine *v1alpha4.Machine, opts v1.UpdateOptions) (result *v1alpha4.Machine, err error) {
	__traceStack()

	result = &v1alpha4.Machine{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("machines").
		Name(machine.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(machine).
		Do(ctx).
		Into(result)
	return
}

func (c *machines) UpdateStatus(ctx context.Context, machine *v1alpha4.Machine, opts v1.UpdateOptions) (result *v1alpha4.Machine, err error) {
	__traceStack()

	result = &v1alpha4.Machine{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("machines").
		Name(machine.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(machine).
		Do(ctx).
		Into(result)
	return
}

func (c *machines) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	return c.client.Delete().
		Namespace(c.ns).
		Resource("machines").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *machines) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("machines").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *machines) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha4.Machine, err error) {
	__traceStack()

	result = &v1alpha4.Machine{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("machines").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
