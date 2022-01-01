package v1

import (
	"context"
	"time"

	scheme "github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v1 "github.com/rancher/system-upgrade-controller/pkg/apis/upgrade.cattle.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

type PlansGetter interface {
	Plans(namespace string) PlanInterface
}

type PlanInterface interface {
	Create(ctx context.Context, plan *v1.Plan, opts metav1.CreateOptions) (*v1.Plan, error)
	Update(ctx context.Context, plan *v1.Plan, opts metav1.UpdateOptions) (*v1.Plan, error)
	UpdateStatus(ctx context.Context, plan *v1.Plan, opts metav1.UpdateOptions) (*v1.Plan, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Plan, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.PlanList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Plan, err error)
	PlanExpansion
}

type plans struct {
	client	rest.Interface
	ns	string
}

func newPlans(c *UpgradeV1Client, namespace string) *plans {
	__traceStack()

	return &plans{
		client:	c.RESTClient(),
		ns:	namespace,
	}
}

func (c *plans) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Plan, err error) {
	__traceStack()

	result = &v1.Plan{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("plans").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *plans) List(ctx context.Context, opts metav1.ListOptions) (result *v1.PlanList, err error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.PlanList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("plans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *plans) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("plans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *plans) Create(ctx context.Context, plan *v1.Plan, opts metav1.CreateOptions) (result *v1.Plan, err error) {
	__traceStack()

	result = &v1.Plan{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("plans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(plan).
		Do(ctx).
		Into(result)
	return
}

func (c *plans) Update(ctx context.Context, plan *v1.Plan, opts metav1.UpdateOptions) (result *v1.Plan, err error) {
	__traceStack()

	result = &v1.Plan{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("plans").
		Name(plan.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(plan).
		Do(ctx).
		Into(result)
	return
}

func (c *plans) UpdateStatus(ctx context.Context, plan *v1.Plan, opts metav1.UpdateOptions) (result *v1.Plan, err error) {
	__traceStack()

	result = &v1.Plan{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("plans").
		Name(plan.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(plan).
		Do(ctx).
		Into(result)
	return
}

func (c *plans) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	__traceStack()

	return c.client.Delete().
		Namespace(c.ns).
		Resource("plans").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *plans) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	__traceStack()

	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("plans").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *plans) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Plan, err error) {
	__traceStack()

	result = &v1.Plan{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("plans").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
