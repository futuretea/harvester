package upgrade

import (
	v1 "github.com/harvester/harvester/pkg/generated/controllers/upgrade.cattle.io/v1"
	"github.com/rancher/lasso/pkg/controller"
)

type Interface interface {
	V1() v1.Interface
}

type group struct {
	controllerFactory controller.SharedControllerFactory
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	__traceStack()

	return &group{
		controllerFactory: controllerFactory,
	}
}

func (g *group) V1() v1.Interface {
	__traceStack()

	return v1.New(g.controllerFactory)
}
