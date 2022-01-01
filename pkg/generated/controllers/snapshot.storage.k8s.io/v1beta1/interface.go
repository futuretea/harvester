package v1beta1

import (
	v1beta1 "github.com/kubernetes-csi/external-snapshotter/v2/pkg/apis/volumesnapshot/v1beta1"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/schemes"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	__traceStack()

	schemes.Register(v1beta1.AddToScheme)
}

type Interface interface {
	VolumeSnapshot() VolumeSnapshotController
	VolumeSnapshotClass() VolumeSnapshotClassController
	VolumeSnapshotContent() VolumeSnapshotContentController
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

func (c *version) VolumeSnapshot() VolumeSnapshotController {
	__traceStack()

	return NewVolumeSnapshotController(schema.GroupVersionKind{Group: "snapshot.storage.k8s.io", Version: "v1beta1", Kind: "VolumeSnapshot"}, "volumesnapshots", true, c.controllerFactory)
}
func (c *version) VolumeSnapshotClass() VolumeSnapshotClassController {
	__traceStack()

	return NewVolumeSnapshotClassController(schema.GroupVersionKind{Group: "snapshot.storage.k8s.io", Version: "v1beta1", Kind: "VolumeSnapshotClass"}, "volumesnapshotclasses", false, c.controllerFactory)
}
func (c *version) VolumeSnapshotContent() VolumeSnapshotContentController {
	__traceStack()

	return NewVolumeSnapshotContentController(schema.GroupVersionKind{Group: "snapshot.storage.k8s.io", Version: "v1beta1", Kind: "VolumeSnapshotContent"}, "volumesnapshotcontents", false, c.controllerFactory)
}
