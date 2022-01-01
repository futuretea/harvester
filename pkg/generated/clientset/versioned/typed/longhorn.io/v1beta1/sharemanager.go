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

type ShareManagersGetter interface {
	ShareManagers(namespace string) ShareManagerInterface
}

type ShareManagerInterface interface {
	Create(ctx context.Context, shareManager *v1beta1.ShareManager, opts v1.CreateOptions) (*v1beta1.ShareManager, error)
	Update(ctx context.Context, shareManager *v1beta1.ShareManager, opts v1.UpdateOptions) (*v1beta1.ShareManager, error)
	UpdateStatus(ctx context.Context, shareManager *v1beta1.ShareManager, opts v1.UpdateOptions) (*v1beta1.ShareManager, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.ShareManager, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ShareManagerList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ShareManager, err error)
	ShareManagerExpansion
}

type shareManagers struct {
	client	rest.Interface
	ns	string
}

func newShareManagers(c *LonghornV1beta1Client, namespace string) *shareManagers {
	__traceStack()

	return &shareManagers{
		client:	c.RESTClient(),
		ns:	namespace,
	}
}

func (c *shareManagers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ShareManager, err error) {
	__traceStack()

	result = &v1beta1.ShareManager{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sharemanagers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *shareManagers) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ShareManagerList, err error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ShareManagerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sharemanagers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *shareManagers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sharemanagers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *shareManagers) Create(ctx context.Context, shareManager *v1beta1.ShareManager, opts v1.CreateOptions) (result *v1beta1.ShareManager, err error) {
	__traceStack()

	result = &v1beta1.ShareManager{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sharemanagers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(shareManager).
		Do(ctx).
		Into(result)
	return
}

func (c *shareManagers) Update(ctx context.Context, shareManager *v1beta1.ShareManager, opts v1.UpdateOptions) (result *v1beta1.ShareManager, err error) {
	__traceStack()

	result = &v1beta1.ShareManager{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sharemanagers").
		Name(shareManager.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(shareManager).
		Do(ctx).
		Into(result)
	return
}

func (c *shareManagers) UpdateStatus(ctx context.Context, shareManager *v1beta1.ShareManager, opts v1.UpdateOptions) (result *v1beta1.ShareManager, err error) {
	__traceStack()

	result = &v1beta1.ShareManager{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sharemanagers").
		Name(shareManager.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(shareManager).
		Do(ctx).
		Into(result)
	return
}

func (c *shareManagers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	return c.client.Delete().
		Namespace(c.ns).
		Resource("sharemanagers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *shareManagers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sharemanagers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *shareManagers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ShareManager, err error) {
	__traceStack()

	result = &v1beta1.ShareManager{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sharemanagers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
