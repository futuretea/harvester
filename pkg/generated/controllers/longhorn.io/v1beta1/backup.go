package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/apply"
	"github.com/rancher/wrangler/pkg/condition"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/kv"
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

type BackupHandler func(string, *v1beta1.Backup) (*v1beta1.Backup, error)

type BackupController interface {
	generic.ControllerMeta
	BackupClient

	OnChange(ctx context.Context, name string, sync BackupHandler)
	OnRemove(ctx context.Context, name string, sync BackupHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() BackupCache
}

type BackupClient interface {
	Create(*v1beta1.Backup) (*v1beta1.Backup, error)
	Update(*v1beta1.Backup) (*v1beta1.Backup, error)
	UpdateStatus(*v1beta1.Backup) (*v1beta1.Backup, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1beta1.Backup, error)
	List(namespace string, opts metav1.ListOptions) (*v1beta1.BackupList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Backup, err error)
}

type BackupCache interface {
	Get(namespace, name string) (*v1beta1.Backup, error)
	List(namespace string, selector labels.Selector) ([]*v1beta1.Backup, error)

	AddIndexer(indexName string, indexer BackupIndexer)
	GetByIndex(indexName, key string) ([]*v1beta1.Backup, error)
}

type BackupIndexer func(obj *v1beta1.Backup) ([]string, error)

type backupController struct {
	controller	controller.SharedController
	client		*client.Client
	gvk		schema.GroupVersionKind
	groupResource	schema.GroupResource
}

func NewBackupController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) BackupController {
	__traceStack()

	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &backupController{
		controller:	c,
		client:		c.Client(),
		gvk:		gvk,
		groupResource: schema.GroupResource{
			Group:		gvk.Group,
			Resource:	resource,
		},
	}
}

func FromBackupHandlerToHandler(sync BackupHandler) generic.Handler {
	__traceStack()

	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1beta1.Backup
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1beta1.Backup))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *backupController) Updater() generic.Updater {
	__traceStack()

	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1beta1.Backup))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateBackupDeepCopyOnChange(client BackupClient, obj *v1beta1.Backup, handler func(obj *v1beta1.Backup) (*v1beta1.Backup, error)) (*v1beta1.Backup, error) {
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

func (c *backupController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	__traceStack()

	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *backupController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	__traceStack()

	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *backupController) OnChange(ctx context.Context, name string, sync BackupHandler) {
	__traceStack()

	c.AddGenericHandler(ctx, name, FromBackupHandlerToHandler(sync))
}

func (c *backupController) OnRemove(ctx context.Context, name string, sync BackupHandler) {
	__traceStack()

	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromBackupHandlerToHandler(sync)))
}

func (c *backupController) Enqueue(namespace, name string) {
	__traceStack()

	c.controller.Enqueue(namespace, name)
}

func (c *backupController) EnqueueAfter(namespace, name string, duration time.Duration) {
	__traceStack()

	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *backupController) Informer() cache.SharedIndexInformer {
	__traceStack()

	return c.controller.Informer()
}

func (c *backupController) GroupVersionKind() schema.GroupVersionKind {
	__traceStack()

	return c.gvk
}

func (c *backupController) Cache() BackupCache {
	__traceStack()

	return &backupCache{
		indexer:	c.Informer().GetIndexer(),
		resource:	c.groupResource,
	}
}

func (c *backupController) Create(obj *v1beta1.Backup) (*v1beta1.Backup, error) {
	__traceStack()

	result := &v1beta1.Backup{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *backupController) Update(obj *v1beta1.Backup) (*v1beta1.Backup, error) {
	__traceStack()

	result := &v1beta1.Backup{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *backupController) UpdateStatus(obj *v1beta1.Backup) (*v1beta1.Backup, error) {
	__traceStack()

	result := &v1beta1.Backup{}
	return result, c.client.UpdateStatus(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *backupController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	__traceStack()

	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *backupController) Get(namespace, name string, options metav1.GetOptions) (*v1beta1.Backup, error) {
	__traceStack()

	result := &v1beta1.Backup{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *backupController) List(namespace string, opts metav1.ListOptions) (*v1beta1.BackupList, error) {
	__traceStack()

	result := &v1beta1.BackupList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *backupController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	__traceStack()

	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *backupController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1beta1.Backup, error) {
	__traceStack()

	result := &v1beta1.Backup{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type backupCache struct {
	indexer		cache.Indexer
	resource	schema.GroupResource
}

func (c *backupCache) Get(namespace, name string) (*v1beta1.Backup, error) {
	__traceStack()

	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1beta1.Backup), nil
}

func (c *backupCache) List(namespace string, selector labels.Selector) (ret []*v1beta1.Backup, err error) {
	__traceStack()

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Backup))
	})

	return ret, err
}

func (c *backupCache) AddIndexer(indexName string, indexer BackupIndexer) {
	__traceStack()

	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1beta1.Backup))
		},
	}))
}

func (c *backupCache) GetByIndex(indexName, key string) (result []*v1beta1.Backup, err error) {
	__traceStack()

	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1beta1.Backup, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1beta1.Backup))
	}
	return result, nil
}

type BackupStatusHandler func(obj *v1beta1.Backup, status v1beta1.BackupStatus) (v1beta1.BackupStatus, error)

type BackupGeneratingHandler func(obj *v1beta1.Backup, status v1beta1.BackupStatus) ([]runtime.Object, v1beta1.BackupStatus, error)

func RegisterBackupStatusHandler(ctx context.Context, controller BackupController, condition condition.Cond, name string, handler BackupStatusHandler) {
	__traceStack()

	statusHandler := &backupStatusHandler{
		client:		controller,
		condition:	condition,
		handler:	handler,
	}
	controller.AddGenericHandler(ctx, name, FromBackupHandlerToHandler(statusHandler.sync))
}

func RegisterBackupGeneratingHandler(ctx context.Context, controller BackupController, apply apply.Apply,
	condition condition.Cond, name string, handler BackupGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	__traceStack()

	statusHandler := &backupGeneratingHandler{
		BackupGeneratingHandler:	handler,
		apply:				apply,
		name:				name,
		gvk:				controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterBackupStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type backupStatusHandler struct {
	client		BackupClient
	condition	condition.Cond
	handler		BackupStatusHandler
}

func (a *backupStatusHandler) sync(key string, obj *v1beta1.Backup) (*v1beta1.Backup, error) {
	__traceStack()

	if obj == nil {
		return obj, nil
	}

	origStatus := obj.Status.DeepCopy()
	obj = obj.DeepCopy()
	newStatus, err := a.handler(obj, obj.Status)
	if err != nil {

		newStatus = *origStatus.DeepCopy()
	}

	if a.condition != "" {
		if errors.IsConflict(err) {
			a.condition.SetError(&newStatus, "", nil)
		} else {
			a.condition.SetError(&newStatus, "", err)
		}
	}
	if !equality.Semantic.DeepEqual(origStatus, &newStatus) {
		if a.condition != "" {

			a.condition.LastUpdated(&newStatus, time.Now().UTC().Format(time.RFC3339))
		}

		var newErr error
		obj.Status = newStatus
		newObj, newErr := a.client.UpdateStatus(obj)
		if err == nil {
			err = newErr
		}
		if newErr == nil {
			obj = newObj
		}
	}
	return obj, err
}

type backupGeneratingHandler struct {
	BackupGeneratingHandler
	apply	apply.Apply
	opts	generic.GeneratingHandlerOptions
	gvk	schema.GroupVersionKind
	name	string
}

func (a *backupGeneratingHandler) Remove(key string, obj *v1beta1.Backup) (*v1beta1.Backup, error) {
	__traceStack()

	if obj != nil {
		return obj, nil
	}

	obj = &v1beta1.Backup{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *backupGeneratingHandler) Handle(obj *v1beta1.Backup, status v1beta1.BackupStatus) (v1beta1.BackupStatus, error) {
	__traceStack()

	if !obj.DeletionTimestamp.IsZero() {
		return status, nil
	}

	objs, newStatus, err := a.BackupGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
