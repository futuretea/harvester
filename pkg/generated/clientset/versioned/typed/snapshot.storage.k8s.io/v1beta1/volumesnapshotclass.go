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

type VolumeSnapshotClassesGetter interface {
	VolumeSnapshotClasses() VolumeSnapshotClassInterface
}

type VolumeSnapshotClassInterface interface {
	Create(ctx context.Context, volumeSnapshotClass *v1beta1.VolumeSnapshotClass, opts v1.CreateOptions) (*v1beta1.VolumeSnapshotClass, error)
	Update(ctx context.Context, volumeSnapshotClass *v1beta1.VolumeSnapshotClass, opts v1.UpdateOptions) (*v1beta1.VolumeSnapshotClass, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.VolumeSnapshotClass, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.VolumeSnapshotClassList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VolumeSnapshotClass, err error)
	VolumeSnapshotClassExpansion
}

type volumeSnapshotClasses struct {
	client rest.Interface
}

func newVolumeSnapshotClasses(c *SnapshotV1beta1Client) *volumeSnapshotClasses {
	__traceStack()

	return &volumeSnapshotClasses{
		client: c.RESTClient(),
	}
}

func (c *volumeSnapshotClasses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VolumeSnapshotClass, err error) {
	__traceStack()

	result = &v1beta1.VolumeSnapshotClass{}
	err = c.client.Get().
		Resource("volumesnapshotclasses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *volumeSnapshotClasses) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VolumeSnapshotClassList, err error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.VolumeSnapshotClassList{}
	err = c.client.Get().
		Resource("volumesnapshotclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *volumeSnapshotClasses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("volumesnapshotclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *volumeSnapshotClasses) Create(ctx context.Context, volumeSnapshotClass *v1beta1.VolumeSnapshotClass, opts v1.CreateOptions) (result *v1beta1.VolumeSnapshotClass, err error) {
	__traceStack()

	result = &v1beta1.VolumeSnapshotClass{}
	err = c.client.Post().
		Resource("volumesnapshotclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeSnapshotClass).
		Do(ctx).
		Into(result)
	return
}

func (c *volumeSnapshotClasses) Update(ctx context.Context, volumeSnapshotClass *v1beta1.VolumeSnapshotClass, opts v1.UpdateOptions) (result *v1beta1.VolumeSnapshotClass, err error) {
	__traceStack()

	result = &v1beta1.VolumeSnapshotClass{}
	err = c.client.Put().
		Resource("volumesnapshotclasses").
		Name(volumeSnapshotClass.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeSnapshotClass).
		Do(ctx).
		Into(result)
	return
}

func (c *volumeSnapshotClasses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	return c.client.Delete().
		Resource("volumesnapshotclasses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *volumeSnapshotClasses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("volumesnapshotclasses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *volumeSnapshotClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VolumeSnapshotClass, err error) {
	__traceStack()

	result = &v1beta1.VolumeSnapshotClass{}
	err = c.client.Patch(pt).
		Resource("volumesnapshotclasses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
