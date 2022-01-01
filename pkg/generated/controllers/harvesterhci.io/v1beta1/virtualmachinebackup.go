package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
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

type VirtualMachineBackupHandler func(string, *v1beta1.VirtualMachineBackup) (*v1beta1.VirtualMachineBackup, error)

type VirtualMachineBackupController interface {
	generic.ControllerMeta
	VirtualMachineBackupClient

	OnChange(ctx context.Context, name string, sync VirtualMachineBackupHandler)
	OnRemove(ctx context.Context, name string, sync VirtualMachineBackupHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() VirtualMachineBackupCache
}

type VirtualMachineBackupClient interface {
	Create(*v1beta1.VirtualMachineBackup) (*v1beta1.VirtualMachineBackup, error)
	Update(*v1beta1.VirtualMachineBackup) (*v1beta1.VirtualMachineBackup, error)

	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1beta1.VirtualMachineBackup, error)
	List(namespace string, opts metav1.ListOptions) (*v1beta1.VirtualMachineBackupList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.VirtualMachineBackup, err error)
}

type VirtualMachineBackupCache interface {
	Get(namespace, name string) (*v1beta1.VirtualMachineBackup, error)
	List(namespace string, selector labels.Selector) ([]*v1beta1.VirtualMachineBackup, error)

	AddIndexer(indexName string, indexer VirtualMachineBackupIndexer)
	GetByIndex(indexName, key string) ([]*v1beta1.VirtualMachineBackup, error)
}

type VirtualMachineBackupIndexer func(obj *v1beta1.VirtualMachineBackup) ([]string, error)

type virtualMachineBackupController struct {
	controller	controller.SharedController
	client		*client.Client
	gvk		schema.GroupVersionKind
	groupResource	schema.GroupResource
}

func NewVirtualMachineBackupController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) VirtualMachineBackupController {
	__traceStack()

	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &virtualMachineBackupController{
		controller:	c,
		client:		c.Client(),
		gvk:		gvk,
		groupResource: schema.GroupResource{
			Group:		gvk.Group,
			Resource:	resource,
		},
	}
}

func FromVirtualMachineBackupHandlerToHandler(sync VirtualMachineBackupHandler) generic.Handler {
	__traceStack()

	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1beta1.VirtualMachineBackup
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1beta1.VirtualMachineBackup))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *virtualMachineBackupController) Updater() generic.Updater {
	__traceStack()

	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1beta1.VirtualMachineBackup))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateVirtualMachineBackupDeepCopyOnChange(client VirtualMachineBackupClient, obj *v1beta1.VirtualMachineBackup, handler func(obj *v1beta1.VirtualMachineBackup) (*v1beta1.VirtualMachineBackup, error)) (*v1beta1.VirtualMachineBackup, error) {
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

func (c *virtualMachineBackupController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	__traceStack()

	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *virtualMachineBackupController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	__traceStack()

	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *virtualMachineBackupController) OnChange(ctx context.Context, name string, sync VirtualMachineBackupHandler) {
	__traceStack()

	c.AddGenericHandler(ctx, name, FromVirtualMachineBackupHandlerToHandler(sync))
}

func (c *virtualMachineBackupController) OnRemove(ctx context.Context, name string, sync VirtualMachineBackupHandler) {
	__traceStack()

	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromVirtualMachineBackupHandlerToHandler(sync)))
}

func (c *virtualMachineBackupController) Enqueue(namespace, name string) {
	__traceStack()

	c.controller.Enqueue(namespace, name)
}

func (c *virtualMachineBackupController) EnqueueAfter(namespace, name string, duration time.Duration) {
	__traceStack()

	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *virtualMachineBackupController) Informer() cache.SharedIndexInformer {
	__traceStack()

	return c.controller.Informer()
}

func (c *virtualMachineBackupController) GroupVersionKind() schema.GroupVersionKind {
	__traceStack()

	return c.gvk
}

func (c *virtualMachineBackupController) Cache() VirtualMachineBackupCache {
	__traceStack()

	return &virtualMachineBackupCache{
		indexer:	c.Informer().GetIndexer(),
		resource:	c.groupResource,
	}
}

func (c *virtualMachineBackupController) Create(obj *v1beta1.VirtualMachineBackup) (*v1beta1.VirtualMachineBackup, error) {
	__traceStack()

	result := &v1beta1.VirtualMachineBackup{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *virtualMachineBackupController) Update(obj *v1beta1.VirtualMachineBackup) (*v1beta1.VirtualMachineBackup, error) {
	__traceStack()

	result := &v1beta1.VirtualMachineBackup{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *virtualMachineBackupController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	__traceStack()

	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *virtualMachineBackupController) Get(namespace, name string, options metav1.GetOptions) (*v1beta1.VirtualMachineBackup, error) {
	__traceStack()

	result := &v1beta1.VirtualMachineBackup{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *virtualMachineBackupController) List(namespace string, opts metav1.ListOptions) (*v1beta1.VirtualMachineBackupList, error) {
	__traceStack()

	result := &v1beta1.VirtualMachineBackupList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *virtualMachineBackupController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *virtualMachineBackupController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1beta1.VirtualMachineBackup, error) {
	__traceStack()

	result := &v1beta1.VirtualMachineBackup{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type virtualMachineBackupCache struct {
	indexer		cache.Indexer
	resource	schema.GroupResource
}

func (c *virtualMachineBackupCache) Get(namespace, name string) (*v1beta1.VirtualMachineBackup, error) {
	__traceStack()

	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1beta1.VirtualMachineBackup), nil
}

func (c *virtualMachineBackupCache) List(namespace string, selector labels.Selector) (ret []*v1beta1.VirtualMachineBackup, err error) {
	__traceStack()

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.VirtualMachineBackup))
	})

	return ret, err
}

func (c *virtualMachineBackupCache) AddIndexer(indexName string, indexer VirtualMachineBackupIndexer) {
	__traceStack()

	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1beta1.VirtualMachineBackup))
		},
	}))
}

func (c *virtualMachineBackupCache) GetByIndex(indexName, key string) (result []*v1beta1.VirtualMachineBackup, err error) {
	__traceStack()

	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1beta1.VirtualMachineBackup, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1beta1.VirtualMachineBackup))
	}
	return result, nil
}
