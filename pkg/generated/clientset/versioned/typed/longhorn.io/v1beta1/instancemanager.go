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

type InstanceManagersGetter interface {
	InstanceManagers(namespace string) InstanceManagerInterface
}

type InstanceManagerInterface interface {
	Create(ctx context.Context, instanceManager *v1beta1.InstanceManager, opts v1.CreateOptions) (*v1beta1.InstanceManager, error)
	Update(ctx context.Context, instanceManager *v1beta1.InstanceManager, opts v1.UpdateOptions) (*v1beta1.InstanceManager, error)
	UpdateStatus(ctx context.Context, instanceManager *v1beta1.InstanceManager, opts v1.UpdateOptions) (*v1beta1.InstanceManager, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.InstanceManager, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.InstanceManagerList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.InstanceManager, err error)
	InstanceManagerExpansion
}

type instanceManagers struct {
	client	rest.Interface
	ns	string
}

func newInstanceManagers(c *LonghornV1beta1Client, namespace string) *instanceManagers {
	__traceStack()

	return &instanceManagers{
		client:	c.RESTClient(),
		ns:	namespace,
	}
}

func (c *instanceManagers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.InstanceManager, err error) {
	__traceStack()

	result = &v1beta1.InstanceManager{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("instancemanagers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *instanceManagers) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.InstanceManagerList, err error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.InstanceManagerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("instancemanagers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *instanceManagers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("instancemanagers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *instanceManagers) Create(ctx context.Context, instanceManager *v1beta1.InstanceManager, opts v1.CreateOptions) (result *v1beta1.InstanceManager, err error) {
	__traceStack()

	result = &v1beta1.InstanceManager{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("instancemanagers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(instanceManager).
		Do(ctx).
		Into(result)
	return
}

func (c *instanceManagers) Update(ctx context.Context, instanceManager *v1beta1.InstanceManager, opts v1.UpdateOptions) (result *v1beta1.InstanceManager, err error) {
	__traceStack()

	result = &v1beta1.InstanceManager{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("instancemanagers").
		Name(instanceManager.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(instanceManager).
		Do(ctx).
		Into(result)
	return
}

func (c *instanceManagers) UpdateStatus(ctx context.Context, instanceManager *v1beta1.InstanceManager, opts v1.UpdateOptions) (result *v1beta1.InstanceManager, err error) {
	__traceStack()

	result = &v1beta1.InstanceManager{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("instancemanagers").
		Name(instanceManager.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(instanceManager).
		Do(ctx).
		Into(result)
	return
}

func (c *instanceManagers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	return c.client.Delete().
		Namespace(c.ns).
		Resource("instancemanagers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *instanceManagers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("instancemanagers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *instanceManagers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.InstanceManager, err error) {
	__traceStack()

	result = &v1beta1.InstanceManager{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("instancemanagers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
