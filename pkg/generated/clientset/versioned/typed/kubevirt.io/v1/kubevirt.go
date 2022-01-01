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

type KubeVirtsGetter interface {
	KubeVirts(namespace string) KubeVirtInterface
}

type KubeVirtInterface interface {
	Create(ctx context.Context, kubeVirt *v1.KubeVirt, opts metav1.CreateOptions) (*v1.KubeVirt, error)
	Update(ctx context.Context, kubeVirt *v1.KubeVirt, opts metav1.UpdateOptions) (*v1.KubeVirt, error)
	UpdateStatus(ctx context.Context, kubeVirt *v1.KubeVirt, opts metav1.UpdateOptions) (*v1.KubeVirt, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.KubeVirt, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.KubeVirtList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.KubeVirt, err error)
	KubeVirtExpansion
}

type kubeVirts struct {
	client	rest.Interface
	ns	string
}

func newKubeVirts(c *KubevirtV1Client, namespace string) *kubeVirts {
	__traceStack()

	return &kubeVirts{
		client:	c.RESTClient(),
		ns:	namespace,
	}
}

func (c *kubeVirts) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.KubeVirt, err error) {
	__traceStack()

	result = &v1.KubeVirt{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kubevirts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *kubeVirts) List(ctx context.Context, opts metav1.ListOptions) (result *v1.KubeVirtList, err error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.KubeVirtList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kubevirts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *kubeVirts) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kubevirts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *kubeVirts) Create(ctx context.Context, kubeVirt *v1.KubeVirt, opts metav1.CreateOptions) (result *v1.KubeVirt, err error) {
	__traceStack()

	result = &v1.KubeVirt{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kubevirts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kubeVirt).
		Do(ctx).
		Into(result)
	return
}

func (c *kubeVirts) Update(ctx context.Context, kubeVirt *v1.KubeVirt, opts metav1.UpdateOptions) (result *v1.KubeVirt, err error) {
	__traceStack()

	result = &v1.KubeVirt{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kubevirts").
		Name(kubeVirt.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kubeVirt).
		Do(ctx).
		Into(result)
	return
}

func (c *kubeVirts) UpdateStatus(ctx context.Context, kubeVirt *v1.KubeVirt, opts metav1.UpdateOptions) (result *v1.KubeVirt, err error) {
	__traceStack()

	result = &v1.KubeVirt{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kubevirts").
		Name(kubeVirt.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kubeVirt).
		Do(ctx).
		Into(result)
	return
}

func (c *kubeVirts) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	__traceStack()

	return c.client.Delete().
		Namespace(c.ns).
		Resource("kubevirts").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *kubeVirts) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	__traceStack()

	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kubevirts").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *kubeVirts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.KubeVirt, err error) {
	__traceStack()

	result = &v1.KubeVirt{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kubevirts").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
