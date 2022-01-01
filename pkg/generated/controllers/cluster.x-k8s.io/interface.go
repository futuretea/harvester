package cluster

import (
	v1alpha4 "github.com/harvester/harvester/pkg/generated/controllers/cluster.x-k8s.io/v1alpha4"
	"github.com/rancher/lasso/pkg/controller"
)

type Interface interface {
	V1alpha4() v1alpha4.Interface
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

func (g *group) V1alpha4() v1alpha4.Interface {
	__traceStack()

	return v1alpha4.New(g.controllerFactory)
}
