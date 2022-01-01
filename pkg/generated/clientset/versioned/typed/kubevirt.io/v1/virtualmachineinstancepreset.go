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

type VirtualMachineInstancePresetsGetter interface {
	VirtualMachineInstancePresets(namespace string) VirtualMachineInstancePresetInterface
}

type VirtualMachineInstancePresetInterface interface {
	Create(ctx context.Context, virtualMachineInstancePreset *v1.VirtualMachineInstancePreset, opts metav1.CreateOptions) (*v1.VirtualMachineInstancePreset, error)
	Update(ctx context.Context, virtualMachineInstancePreset *v1.VirtualMachineInstancePreset, opts metav1.UpdateOptions) (*v1.VirtualMachineInstancePreset, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.VirtualMachineInstancePreset, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.VirtualMachineInstancePresetList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VirtualMachineInstancePreset, err error)
	VirtualMachineInstancePresetExpansion
}

type virtualMachineInstancePresets struct {
	client	rest.Interface
	ns	string
}

func newVirtualMachineInstancePresets(c *KubevirtV1Client, namespace string) *virtualMachineInstancePresets {
	__traceStack()

	return &virtualMachineInstancePresets{
		client:	c.RESTClient(),
		ns:	namespace,
	}
}

func (c *virtualMachineInstancePresets) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.VirtualMachineInstancePreset, err error) {
	__traceStack()

	result = &v1.VirtualMachineInstancePreset{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachineinstancepresets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineInstancePresets) List(ctx context.Context, opts metav1.ListOptions) (result *v1.VirtualMachineInstancePresetList, err error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.VirtualMachineInstancePresetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachineinstancepresets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineInstancePresets) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachineinstancepresets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *virtualMachineInstancePresets) Create(ctx context.Context, virtualMachineInstancePreset *v1.VirtualMachineInstancePreset, opts metav1.CreateOptions) (result *v1.VirtualMachineInstancePreset, err error) {
	__traceStack()

	result = &v1.VirtualMachineInstancePreset{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualmachineinstancepresets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineInstancePreset).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineInstancePresets) Update(ctx context.Context, virtualMachineInstancePreset *v1.VirtualMachineInstancePreset, opts metav1.UpdateOptions) (result *v1.VirtualMachineInstancePreset, err error) {
	__traceStack()

	result = &v1.VirtualMachineInstancePreset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachineinstancepresets").
		Name(virtualMachineInstancePreset.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineInstancePreset).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineInstancePresets) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	__traceStack()

	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachineinstancepresets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *virtualMachineInstancePresets) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	__traceStack()

	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachineinstancepresets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *virtualMachineInstancePresets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VirtualMachineInstancePreset, err error) {
	__traceStack()

	result = &v1.VirtualMachineInstancePreset{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualmachineinstancepresets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
