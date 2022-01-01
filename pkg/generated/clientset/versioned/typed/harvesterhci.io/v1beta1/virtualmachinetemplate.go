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

type VirtualMachineTemplatesGetter interface {
	VirtualMachineTemplates(namespace string) VirtualMachineTemplateInterface
}

type VirtualMachineTemplateInterface interface {
	Create(ctx context.Context, virtualMachineTemplate *v1beta1.VirtualMachineTemplate, opts v1.CreateOptions) (*v1beta1.VirtualMachineTemplate, error)
	Update(ctx context.Context, virtualMachineTemplate *v1beta1.VirtualMachineTemplate, opts v1.UpdateOptions) (*v1beta1.VirtualMachineTemplate, error)
	UpdateStatus(ctx context.Context, virtualMachineTemplate *v1beta1.VirtualMachineTemplate, opts v1.UpdateOptions) (*v1beta1.VirtualMachineTemplate, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.VirtualMachineTemplate, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.VirtualMachineTemplateList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VirtualMachineTemplate, err error)
	VirtualMachineTemplateExpansion
}

type virtualMachineTemplates struct {
	client	rest.Interface
	ns	string
}

func newVirtualMachineTemplates(c *HarvesterhciV1beta1Client, namespace string) *virtualMachineTemplates {
	__traceStack()

	return &virtualMachineTemplates{
		client:	c.RESTClient(),
		ns:	namespace,
	}
}

func (c *virtualMachineTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VirtualMachineTemplate, err error) {
	__traceStack()

	result = &v1beta1.VirtualMachineTemplate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinetemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VirtualMachineTemplateList, err error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.VirtualMachineTemplateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *virtualMachineTemplates) Create(ctx context.Context, virtualMachineTemplate *v1beta1.VirtualMachineTemplate, opts v1.CreateOptions) (result *v1beta1.VirtualMachineTemplate, err error) {
	__traceStack()

	result = &v1beta1.VirtualMachineTemplate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualmachinetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineTemplate).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineTemplates) Update(ctx context.Context, virtualMachineTemplate *v1beta1.VirtualMachineTemplate, opts v1.UpdateOptions) (result *v1beta1.VirtualMachineTemplate, err error) {
	__traceStack()

	result = &v1beta1.VirtualMachineTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachinetemplates").
		Name(virtualMachineTemplate.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineTemplate).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineTemplates) UpdateStatus(ctx context.Context, virtualMachineTemplate *v1beta1.VirtualMachineTemplate, opts v1.UpdateOptions) (result *v1beta1.VirtualMachineTemplate, err error) {
	__traceStack()

	result = &v1beta1.VirtualMachineTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachinetemplates").
		Name(virtualMachineTemplate.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineTemplate).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinetemplates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *virtualMachineTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinetemplates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *virtualMachineTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VirtualMachineTemplate, err error) {
	__traceStack()

	result = &v1beta1.VirtualMachineTemplate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualmachinetemplates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
