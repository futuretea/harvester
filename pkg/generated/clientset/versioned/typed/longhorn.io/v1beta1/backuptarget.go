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

type BackupTargetsGetter interface {
	BackupTargets(namespace string) BackupTargetInterface
}

type BackupTargetInterface interface {
	Create(ctx context.Context, backupTarget *v1beta1.BackupTarget, opts v1.CreateOptions) (*v1beta1.BackupTarget, error)
	Update(ctx context.Context, backupTarget *v1beta1.BackupTarget, opts v1.UpdateOptions) (*v1beta1.BackupTarget, error)
	UpdateStatus(ctx context.Context, backupTarget *v1beta1.BackupTarget, opts v1.UpdateOptions) (*v1beta1.BackupTarget, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.BackupTarget, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.BackupTargetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BackupTarget, err error)
	BackupTargetExpansion
}

type backupTargets struct {
	client	rest.Interface
	ns	string
}

func newBackupTargets(c *LonghornV1beta1Client, namespace string) *backupTargets {
	__traceStack()

	return &backupTargets{
		client:	c.RESTClient(),
		ns:	namespace,
	}
}

func (c *backupTargets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.BackupTarget, err error) {
	__traceStack()

	result = &v1beta1.BackupTarget{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backuptargets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *backupTargets) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.BackupTargetList, err error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.BackupTargetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backuptargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *backupTargets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("backuptargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *backupTargets) Create(ctx context.Context, backupTarget *v1beta1.BackupTarget, opts v1.CreateOptions) (result *v1beta1.BackupTarget, err error) {
	__traceStack()

	result = &v1beta1.BackupTarget{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("backuptargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backupTarget).
		Do(ctx).
		Into(result)
	return
}

func (c *backupTargets) Update(ctx context.Context, backupTarget *v1beta1.BackupTarget, opts v1.UpdateOptions) (result *v1beta1.BackupTarget, err error) {
	__traceStack()

	result = &v1beta1.BackupTarget{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backuptargets").
		Name(backupTarget.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backupTarget).
		Do(ctx).
		Into(result)
	return
}

func (c *backupTargets) UpdateStatus(ctx context.Context, backupTarget *v1beta1.BackupTarget, opts v1.UpdateOptions) (result *v1beta1.BackupTarget, err error) {
	__traceStack()

	result = &v1beta1.BackupTarget{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backuptargets").
		Name(backupTarget.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backupTarget).
		Do(ctx).
		Into(result)
	return
}

func (c *backupTargets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	return c.client.Delete().
		Namespace(c.ns).
		Resource("backuptargets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *backupTargets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backuptargets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *backupTargets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BackupTarget, err error) {
	__traceStack()

	result = &v1beta1.BackupTarget{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("backuptargets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
