package v1beta1

import (
	"context"
	"time"

	scheme "github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v1beta1 "github.com/kubernetes-csi/external-snapshotter/v2/pkg/apis/volumesnapshot/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

type VolumeSnapshotContentsGetter interface {
	VolumeSnapshotContents() VolumeSnapshotContentInterface
}

type VolumeSnapshotContentInterface interface {
	Create(ctx context.Context, volumeSnapshotContent *v1beta1.VolumeSnapshotContent, opts v1.CreateOptions) (*v1beta1.VolumeSnapshotContent, error)
	Update(ctx context.Context, volumeSnapshotContent *v1beta1.VolumeSnapshotContent, opts v1.UpdateOptions) (*v1beta1.VolumeSnapshotContent, error)
	UpdateStatus(ctx context.Context, volumeSnapshotContent *v1beta1.VolumeSnapshotContent, opts v1.UpdateOptions) (*v1beta1.VolumeSnapshotContent, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.VolumeSnapshotContent, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.VolumeSnapshotContentList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VolumeSnapshotContent, err error)
	VolumeSnapshotContentExpansion
}

type volumeSnapshotContents struct {
	client rest.Interface
}

func newVolumeSnapshotContents(c *SnapshotV1beta1Client) *volumeSnapshotContents {
	__traceStack()

	return &volumeSnapshotContents{
		client: c.RESTClient(),
	}
}

func (c *volumeSnapshotContents) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VolumeSnapshotContent, err error) {
	__traceStack()

	result = &v1beta1.VolumeSnapshotContent{}
	err = c.client.Get().
		Resource("volumesnapshotcontents").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *volumeSnapshotContents) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VolumeSnapshotContentList, err error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.VolumeSnapshotContentList{}
	err = c.client.Get().
		Resource("volumesnapshotcontents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *volumeSnapshotContents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("volumesnapshotcontents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *volumeSnapshotContents) Create(ctx context.Context, volumeSnapshotContent *v1beta1.VolumeSnapshotContent, opts v1.CreateOptions) (result *v1beta1.VolumeSnapshotContent, err error) {
	__traceStack()

	result = &v1beta1.VolumeSnapshotContent{}
	err = c.client.Post().
		Resource("volumesnapshotcontents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeSnapshotContent).
		Do(ctx).
		Into(result)
	return
}

func (c *volumeSnapshotContents) Update(ctx context.Context, volumeSnapshotContent *v1beta1.VolumeSnapshotContent, opts v1.UpdateOptions) (result *v1beta1.VolumeSnapshotContent, err error) {
	__traceStack()

	result = &v1beta1.VolumeSnapshotContent{}
	err = c.client.Put().
		Resource("volumesnapshotcontents").
		Name(volumeSnapshotContent.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeSnapshotContent).
		Do(ctx).
		Into(result)
	return
}

func (c *volumeSnapshotContents) UpdateStatus(ctx context.Context, volumeSnapshotContent *v1beta1.VolumeSnapshotContent, opts v1.UpdateOptions) (result *v1beta1.VolumeSnapshotContent, err error) {
	__traceStack()

	result = &v1beta1.VolumeSnapshotContent{}
	err = c.client.Put().
		Resource("volumesnapshotcontents").
		Name(volumeSnapshotContent.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeSnapshotContent).
		Do(ctx).
		Into(result)
	return
}

func (c *volumeSnapshotContents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	return c.client.Delete().
		Resource("volumesnapshotcontents").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *volumeSnapshotContents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("volumesnapshotcontents").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *volumeSnapshotContents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VolumeSnapshotContent, err error) {
	__traceStack()

	result = &v1beta1.VolumeSnapshotContent{}
	err = c.client.Patch(pt).
		Resource("volumesnapshotcontents").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
