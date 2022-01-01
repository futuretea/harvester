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

type VolumeSnapshotContentHandler func(string, *v1beta1.VolumeSnapshotContent) (*v1beta1.VolumeSnapshotContent, error)

type VolumeSnapshotContentController interface {
	generic.ControllerMeta
	VolumeSnapshotContentClient

	OnChange(ctx context.Context, name string, sync VolumeSnapshotContentHandler)
	OnRemove(ctx context.Context, name string, sync VolumeSnapshotContentHandler)
	Enqueue(name string)
	EnqueueAfter(name string, duration time.Duration)

	Cache() VolumeSnapshotContentCache
}

type VolumeSnapshotContentClient interface {
	Create(*v1beta1.VolumeSnapshotContent) (*v1beta1.VolumeSnapshotContent, error)
	Update(*v1beta1.VolumeSnapshotContent) (*v1beta1.VolumeSnapshotContent, error)

	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*v1beta1.VolumeSnapshotContent, error)
	List(opts metav1.ListOptions) (*v1beta1.VolumeSnapshotContentList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.VolumeSnapshotContent, err error)
}

type VolumeSnapshotContentCache interface {
	Get(name string) (*v1beta1.VolumeSnapshotContent, error)
	List(selector labels.Selector) ([]*v1beta1.VolumeSnapshotContent, error)

	AddIndexer(indexName string, indexer VolumeSnapshotContentIndexer)
	GetByIndex(indexName, key string) ([]*v1beta1.VolumeSnapshotContent, error)
}

type VolumeSnapshotContentIndexer func(obj *v1beta1.VolumeSnapshotContent) ([]string, error)

type volumeSnapshotContentController struct {
	controller	controller.SharedController
	client		*client.Client
	gvk		schema.GroupVersionKind
	groupResource	schema.GroupResource
}

func NewVolumeSnapshotContentController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) VolumeSnapshotContentController {
	__traceStack()

	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &volumeSnapshotContentController{
		controller:	c,
		client:		c.Client(),
		gvk:		gvk,
		groupResource: schema.GroupResource{
			Group:		gvk.Group,
			Resource:	resource,
		},
	}
}

func FromVolumeSnapshotContentHandlerToHandler(sync VolumeSnapshotContentHandler) generic.Handler {
	__traceStack()

	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1beta1.VolumeSnapshotContent
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1beta1.VolumeSnapshotContent))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *volumeSnapshotContentController) Updater() generic.Updater {
	__traceStack()

	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1beta1.VolumeSnapshotContent))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateVolumeSnapshotContentDeepCopyOnChange(client VolumeSnapshotContentClient, obj *v1beta1.VolumeSnapshotContent, handler func(obj *v1beta1.VolumeSnapshotContent) (*v1beta1.VolumeSnapshotContent, error)) (*v1beta1.VolumeSnapshotContent, error) {
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

func (c *volumeSnapshotContentController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	__traceStack()

	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *volumeSnapshotContentController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	__traceStack()

	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *volumeSnapshotContentController) OnChange(ctx context.Context, name string, sync VolumeSnapshotContentHandler) {
	__traceStack()

	c.AddGenericHandler(ctx, name, FromVolumeSnapshotContentHandlerToHandler(sync))
}

func (c *volumeSnapshotContentController) OnRemove(ctx context.Context, name string, sync VolumeSnapshotContentHandler) {
	__traceStack()

	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromVolumeSnapshotContentHandlerToHandler(sync)))
}

func (c *volumeSnapshotContentController) Enqueue(name string) {
	__traceStack()

	c.controller.Enqueue("", name)
}

func (c *volumeSnapshotContentController) EnqueueAfter(name string, duration time.Duration) {
	__traceStack()

	c.controller.EnqueueAfter("", name, duration)
}

func (c *volumeSnapshotContentController) Informer() cache.SharedIndexInformer {
	__traceStack()

	return c.controller.Informer()
}

func (c *volumeSnapshotContentController) GroupVersionKind() schema.GroupVersionKind {
	__traceStack()

	return c.gvk
}

func (c *volumeSnapshotContentController) Cache() VolumeSnapshotContentCache {
	__traceStack()

	return &volumeSnapshotContentCache{
		indexer:	c.Informer().GetIndexer(),
		resource:	c.groupResource,
	}
}

func (c *volumeSnapshotContentController) Create(obj *v1beta1.VolumeSnapshotContent) (*v1beta1.VolumeSnapshotContent, error) {
	__traceStack()

	result := &v1beta1.VolumeSnapshotContent{}
	return result, c.client.Create(context.TODO(), "", obj, result, metav1.CreateOptions{})
}

func (c *volumeSnapshotContentController) Update(obj *v1beta1.VolumeSnapshotContent) (*v1beta1.VolumeSnapshotContent, error) {
	__traceStack()

	result := &v1beta1.VolumeSnapshotContent{}
	return result, c.client.Update(context.TODO(), "", obj, result, metav1.UpdateOptions{})
}

func (c *volumeSnapshotContentController) Delete(name string, options *metav1.DeleteOptions) error {
	__traceStack()

	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), "", name, *options)
}

func (c *volumeSnapshotContentController) Get(name string, options metav1.GetOptions) (*v1beta1.VolumeSnapshotContent, error) {
	__traceStack()

	result := &v1beta1.VolumeSnapshotContent{}
	return result, c.client.Get(context.TODO(), "", name, result, options)
}

func (c *volumeSnapshotContentController) List(opts metav1.ListOptions) (*v1beta1.VolumeSnapshotContentList, error) {
	__traceStack()

	result := &v1beta1.VolumeSnapshotContentList{}
	return result, c.client.List(context.TODO(), "", result, opts)
}

func (c *volumeSnapshotContentController) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.client.Watch(context.TODO(), "", opts)
}

func (c *volumeSnapshotContentController) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (*v1beta1.VolumeSnapshotContent, error) {
	__traceStack()

	result := &v1beta1.VolumeSnapshotContent{}
	return result, c.client.Patch(context.TODO(), "", name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type volumeSnapshotContentCache struct {
	indexer		cache.Indexer
	resource	schema.GroupResource
}

func (c *volumeSnapshotContentCache) Get(name string) (*v1beta1.VolumeSnapshotContent, error) {
	__traceStack()

	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1beta1.VolumeSnapshotContent), nil
}

func (c *volumeSnapshotContentCache) List(selector labels.Selector) (ret []*v1beta1.VolumeSnapshotContent, err error) {
	__traceStack()

	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.VolumeSnapshotContent))
	})

	return ret, err
}

func (c *volumeSnapshotContentCache) AddIndexer(indexName string, indexer VolumeSnapshotContentIndexer) {
	__traceStack()

	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1beta1.VolumeSnapshotContent))
		},
	}))
}

func (c *volumeSnapshotContentCache) GetByIndex(indexName, key string) (result []*v1beta1.VolumeSnapshotContent, err error) {
	__traceStack()

	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1beta1.VolumeSnapshotContent, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1beta1.VolumeSnapshotContent))
	}
	return result, nil
}
