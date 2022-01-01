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

type VolumeSnapshotsGetter interface {
	VolumeSnapshots(namespace string) VolumeSnapshotInterface
}

type VolumeSnapshotInterface interface {
	Create(ctx context.Context, volumeSnapshot *v1beta1.VolumeSnapshot, opts v1.CreateOptions) (*v1beta1.VolumeSnapshot, error)
	Update(ctx context.Context, volumeSnapshot *v1beta1.VolumeSnapshot, opts v1.UpdateOptions) (*v1beta1.VolumeSnapshot, error)
	UpdateStatus(ctx context.Context, volumeSnapshot *v1beta1.VolumeSnapshot, opts v1.UpdateOptions) (*v1beta1.VolumeSnapshot, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.VolumeSnapshot, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.VolumeSnapshotList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VolumeSnapshot, err error)
	VolumeSnapshotExpansion
}

type volumeSnapshots struct {
	client	rest.Interface
	ns	string
}

func newVolumeSnapshots(c *SnapshotV1beta1Client, namespace string) *volumeSnapshots {
	__traceStack()

	return &volumeSnapshots{
		client:	c.RESTClient(),
		ns:	namespace,
	}
}

func (c *volumeSnapshots) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VolumeSnapshot, err error) {
	__traceStack()

	result = &v1beta1.VolumeSnapshot{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("volumesnapshots").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *volumeSnapshots) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VolumeSnapshotList, err error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.VolumeSnapshotList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("volumesnapshots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *volumeSnapshots) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	__traceStack()

	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("volumesnapshots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *volumeSnapshots) Create(ctx context.Context, volumeSnapshot *v1beta1.VolumeSnapshot, opts v1.CreateOptions) (result *v1beta1.VolumeSnapshot, err error) {
	__traceStack()

	result = &v1beta1.VolumeSnapshot{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("volumesnapshots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeSnapshot).
		Do(ctx).
		Into(result)
	return
}

func (c *volumeSnapshots) Update(ctx context.Context, volumeSnapshot *v1beta1.VolumeSnapshot, opts v1.UpdateOptions) (result *v1beta1.VolumeSnapshot, err error) {
	__traceStack()

	result = &v1beta1.VolumeSnapshot{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("volumesnapshots").
		Name(volumeSnapshot.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeSnapshot).
		Do(ctx).
		Into(result)
	return
}

func (c *volumeSnapshots) UpdateStatus(ctx context.Context, volumeSnapshot *v1beta1.VolumeSnapshot, opts v1.UpdateOptions) (result *v1beta1.VolumeSnapshot, err error) {
	__traceStack()

	result = &v1beta1.VolumeSnapshot{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("volumesnapshots").
		Name(volumeSnapshot.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(volumeSnapshot).
		Do(ctx).
		Into(result)
	return
}

func (c *volumeSnapshots) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	__traceStack()

	return c.client.Delete().
		Namespace(c.ns).
		Resource("volumesnapshots").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *volumeSnapshots) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	__traceStack()

	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("volumesnapshots").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *volumeSnapshots) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VolumeSnapshot, err error) {
	__traceStack()

	result = &v1beta1.VolumeSnapshot{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("volumesnapshots").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
