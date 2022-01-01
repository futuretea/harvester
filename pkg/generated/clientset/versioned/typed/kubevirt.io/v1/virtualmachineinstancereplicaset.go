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

type VirtualMachineInstanceReplicaSetsGetter interface {
	VirtualMachineInstanceReplicaSets(namespace string) VirtualMachineInstanceReplicaSetInterface
}

type VirtualMachineInstanceReplicaSetInterface interface {
	Create(ctx context.Context, virtualMachineInstanceReplicaSet *v1.VirtualMachineInstanceReplicaSet, opts metav1.CreateOptions) (*v1.VirtualMachineInstanceReplicaSet, error)
	Update(ctx context.Context, virtualMachineInstanceReplicaSet *v1.VirtualMachineInstanceReplicaSet, opts metav1.UpdateOptions) (*v1.VirtualMachineInstanceReplicaSet, error)
	UpdateStatus(ctx context.Context, virtualMachineInstanceReplicaSet *v1.VirtualMachineInstanceReplicaSet, opts metav1.UpdateOptions) (*v1.VirtualMachineInstanceReplicaSet, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.VirtualMachineInstanceReplicaSet, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.VirtualMachineInstanceReplicaSetList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VirtualMachineInstanceReplicaSet, err error)
	VirtualMachineInstanceReplicaSetExpansion
}

type virtualMachineInstanceReplicaSets struct {
	client	rest.Interface
	ns	string
}

func newVirtualMachineInstanceReplicaSets(c *KubevirtV1Client, namespace string) *virtualMachineInstanceReplicaSets {
	__traceStack()

	return &virtualMachineInstanceReplicaSets{
		client:	c.RESTClient(),
		ns:	namespace,
	}
}

func (c *virtualMachineInstanceReplicaSets) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.VirtualMachineInstanceReplicaSet, err error) {
	__traceStack()

	result = &v1.VirtualMachineInstanceReplicaSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachineinstancereplicasets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineInstanceReplicaSets) List(ctx context.Context, opts metav1.ListOptions) (result *v1.VirtualMachineInstanceReplicaSetList, err error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.VirtualMachineInstanceReplicaSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachineinstancereplicasets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineInstanceReplicaSets) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachineinstancereplicasets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *virtualMachineInstanceReplicaSets) Create(ctx context.Context, virtualMachineInstanceReplicaSet *v1.VirtualMachineInstanceReplicaSet, opts metav1.CreateOptions) (result *v1.VirtualMachineInstanceReplicaSet, err error) {
	__traceStack()

	result = &v1.VirtualMachineInstanceReplicaSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualmachineinstancereplicasets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineInstanceReplicaSet).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineInstanceReplicaSets) Update(ctx context.Context, virtualMachineInstanceReplicaSet *v1.VirtualMachineInstanceReplicaSet, opts metav1.UpdateOptions) (result *v1.VirtualMachineInstanceReplicaSet, err error) {
	__traceStack()

	result = &v1.VirtualMachineInstanceReplicaSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachineinstancereplicasets").
		Name(virtualMachineInstanceReplicaSet.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineInstanceReplicaSet).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineInstanceReplicaSets) UpdateStatus(ctx context.Context, virtualMachineInstanceReplicaSet *v1.VirtualMachineInstanceReplicaSet, opts metav1.UpdateOptions) (result *v1.VirtualMachineInstanceReplicaSet, err error) {
	__traceStack()

	result = &v1.VirtualMachineInstanceReplicaSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachineinstancereplicasets").
		Name(virtualMachineInstanceReplicaSet.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineInstanceReplicaSet).
		Do(ctx).
		Into(result)
	return
}

func (c *virtualMachineInstanceReplicaSets) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	__traceStack()

	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachineinstancereplicasets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *virtualMachineInstanceReplicaSets) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	__traceStack()

	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachineinstancereplicasets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *virtualMachineInstanceReplicaSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VirtualMachineInstanceReplicaSet, err error) {
	__traceStack()

	result = &v1.VirtualMachineInstanceReplicaSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualmachineinstancereplicasets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
