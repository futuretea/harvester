package v1beta1

import (
	v1beta1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/schemes"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	__traceStack()

	schemes.Register(v1beta1.AddToScheme)
}

type Interface interface {
	BackingImage() BackingImageController
	BackingImageDataSource() BackingImageDataSourceController
	Backup() BackupController
	Setting() SettingController
	Volume() VolumeController
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	__traceStack()

	return &version{
		controllerFactory: controllerFactory,
	}
}

type version struct {
	controllerFactory controller.SharedControllerFactory
}

func (c *version) BackingImage() BackingImageController {
	__traceStack()

	return NewBackingImageController(schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta1", Kind: "BackingImage"}, "backingimages", true, c.controllerFactory)
}
func (c *version) BackingImageDataSource() BackingImageDataSourceController {
	__traceStack()

	return NewBackingImageDataSourceController(schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta1", Kind: "BackingImageDataSource"}, "backingimagedatasources", true, c.controllerFactory)
}
func (c *version) Backup() BackupController {
	__traceStack()

	return NewBackupController(schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta1", Kind: "Backup"}, "backups", true, c.controllerFactory)
}
func (c *version) Setting() SettingController {
	__traceStack()

	return NewSettingController(schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta1", Kind: "Setting"}, "settings", true, c.controllerFactory)
}
func (c *version) Volume() VolumeController {
	__traceStack()

	return NewVolumeController(schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta1", Kind: "Volume"}, "volumes", true, c.controllerFactory)
}
