package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/kubernetes-csi/external-snapshotter/v2/pkg/apis/volumesnapshot/v1beta1"
	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/generic"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type VolumeSnapshotHandler func(string, *v1beta1.VolumeSnapshot) (*v1beta1.VolumeSnapshot, error)

type VolumeSnapshotController interface {
	generic.ControllerMeta
	VolumeSnapshotClient

	OnChange(ctx context.Context, name string, sync VolumeSnapshotHandler)
	OnRemove(ctx context.Context, name string, sync VolumeSnapshotHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() VolumeSnapshotCache
}

type VolumeSnapshotClient interface {
	Create(*v1beta1.VolumeSnapshot) (*v1beta1.VolumeSnapshot, error)
	Update(*v1beta1.VolumeSnapshot) (*v1beta1.VolumeSnapshot, error)

	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1beta1.VolumeSnapshot, error)
	List(namespace string, opts metav1.ListOptions) (*v1beta1.VolumeSnapshotList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.VolumeSnapshot, err error)
}

type VolumeSnapshotCache interface {
	Get(namespace, name string) (*v1beta1.VolumeSnapshot, error)
	List(namespace string, selector labels.Selector) ([]*v1beta1.VolumeSnapshot, error)

	AddIndexer(indexName string, indexer VolumeSnapshotIndexer)
	GetByIndex(indexName, key string) ([]*v1beta1.VolumeSnapshot, error)
}

type VolumeSnapshotIndexer func(obj *v1beta1.VolumeSnapshot) ([]string, error)

type volumeSnapshotController struct {
	controller	controller.SharedController
	client		*client.Client
	gvk		schema.GroupVersionKind
	groupResource	schema.GroupResource
}

func NewVolumeSnapshotController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) VolumeSnapshotController {
	__traceStack()

	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &volumeSnapshotController{
		controller:	c,
		client:		c.Client(),
		gvk:		gvk,
		groupResource: schema.GroupResource{
			Group:		gvk.Group,
			Resource:	resource,
		},
	}
}

func FromVolumeSnapshotHandlerToHandler(sync VolumeSnapshotHandler) generic.Handler {
	__traceStack()

	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1beta1.VolumeSnapshot
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1beta1.VolumeSnapshot))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *volumeSnapshotController) Updater() generic.Updater {
	__traceStack()

	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1beta1.VolumeSnapshot))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateVolumeSnapshotDeepCopyOnChange(client VolumeSnapshotClient, obj *v1beta1.VolumeSnapshot, handler func(obj *v1beta1.VolumeSnapshot) (*v1beta1.VolumeSnapshot, error)) (*v1beta1.VolumeSnapshot, error) {
	__traceStack()

	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *volumeSnapshotController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	__traceStack()

	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *volumeSnapshotController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	__traceStack()

	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *volumeSnapshotController) OnChange(ctx context.Context, name string, sync VolumeSnapshotHandler) {
	__traceStack()

	c.AddGenericHandler(ctx, name, FromVolumeSnapshotHandlerToHandler(sync))
}

func (c *volumeSnapshotController) OnRemove(ctx context.Context, name string, sync VolumeSnapshotHandler) {
	__traceStack()

	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromVolumeSnapshotHandlerToHandler(sync)))
}

func (c *volumeSnapshotController) Enqueue(namespace, name string) {
	__traceStack()

	c.controller.Enqueue(namespace, name)
}

func (c *volumeSnapshotController) EnqueueAfter(namespace, name string, duration time.Duration) {
	__traceStack()

	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *volumeSnapshotController) Informer() cache.SharedIndexInformer {
	__traceStack()

	return c.controller.Informer()
}

func (c *volumeSnapshotController) GroupVersionKind() schema.GroupVersionKind {
	__traceStack()

	return c.gvk
}

func (c *volumeSnapshotController) Cache() VolumeSnapshotCache {
	__traceStack()

	return &volumeSnapshotCache{
		indexer:	c.Informer().GetIndexer(),
		resource:	c.groupResource,
	}
}

func (c *volumeSnapshotController) Create(obj *v1beta1.VolumeSnapshot) (*v1beta1.VolumeSnapshot, error) {
	__traceStack()

	result := &v1beta1.VolumeSnapshot{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *volumeSnapshotController) Update(obj *v1beta1.VolumeSnapshot) (*v1beta1.VolumeSnapshot, error) {
	__traceStack()

	result := &v1beta1.VolumeSnapshot{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *volumeSnapshotController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	__traceStack()

	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *volumeSnapshotController) Get(namespace, name string, options metav1.GetOptions) (*v1beta1.VolumeSnapshot, error) {
	__traceStack()

	result := &v1beta1.VolumeSnapshot{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *volumeSnapshotController) List(namespace string, opts metav1.ListOptions) (*v1beta1.VolumeSnapshotList, error) {
	__traceStack()

	result := &v1beta1.VolumeSnapshotList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *volumeSnapshotController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *volumeSnapshotController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1beta1.VolumeSnapshot, error) {
	__traceStack()

	result := &v1beta1.VolumeSnapshot{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type volumeSnapshotCache struct {
	indexer		cache.Indexer
	resource	schema.GroupResource
}

func (c *volumeSnapshotCache) Get(namespace, name string) (*v1beta1.VolumeSnapshot, error) {
	__traceStack()

	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1beta1.VolumeSnapshot), nil
}

func (c *volumeSnapshotCache) List(namespace string, selector labels.Selector) (ret []*v1beta1.VolumeSnapshot, err error) {
	__traceStack()

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.VolumeSnapshot))
	})

	return ret, err
}

func (c *volumeSnapshotCache) AddIndexer(indexName string, indexer VolumeSnapshotIndexer) {
	__traceStack()

	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1beta1.VolumeSnapshot))
		},
	}))
}

func (c *volumeSnapshotCache) GetByIndex(indexName, key string) (result []*v1beta1.VolumeSnapshot, err error) {
	__traceStack()

	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1beta1.VolumeSnapshot, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1beta1.VolumeSnapshot))
	}
	return result, nil
}
