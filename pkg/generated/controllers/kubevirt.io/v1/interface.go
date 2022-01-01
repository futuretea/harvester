package v1

import (
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/schemes"
	"k8s.io/apimachinery/pkg/runtime/schema"
	v1 "kubevirt.io/client-go/api/v1"
)

func init() {
	__traceStack()

	schemes.Register(v1.AddToScheme)
}

type Interface interface {
	VirtualMachine() VirtualMachineController
	VirtualMachineInstance() VirtualMachineInstanceController
	VirtualMachineInstanceMigration() VirtualMachineInstanceMigrationController
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

func (c *version) VirtualMachine() VirtualMachineController {
	__traceStack()

	return NewVirtualMachineController(schema.GroupVersionKind{Group: "kubevirt.io", Version: "v1", Kind: "VirtualMachine"}, "virtualmachines", true, c.controllerFactory)
}
func (c *version) VirtualMachineInstance() VirtualMachineInstanceController {
	__traceStack()

	return NewVirtualMachineInstanceController(schema.GroupVersionKind{Group: "kubevirt.io", Version: "v1", Kind: "VirtualMachineInstance"}, "virtualmachineinstances", true, c.controllerFactory)
}
func (c *version) VirtualMachineInstanceMigration() VirtualMachineInstanceMigrationController {
	__traceStack()

	return NewVirtualMachineInstanceMigrationController(schema.GroupVersionKind{Group: "kubevirt.io", Version: "v1", Kind: "VirtualMachineInstanceMigration"}, "virtualmachineinstancemigrations", true, c.controllerFactory)
}
