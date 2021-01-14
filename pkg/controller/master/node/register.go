package node

import (
	"context"

	"github.com/rancher/harvester/pkg/config"
)

const (
	controllerAgentName = "node-controller"
)

// Register registers the node controller
func Register(ctx context.Context, management *config.Management) error {
	nodes := management.CoreFactory.Core().V1().Node()
	pods := management.CoreFactory.Core().V1().Pod()
	jobs := management.BatchFactory.Batch().V1().Job()
	statefulsets := management.AppsFactory.Apps().V1().StatefulSet()

	controller := &Handler{
		podCache:         pods.Cache(),
		nodes:            nodes,
		nodeCache:        nodes.Cache(),
		jobs:             jobs,
		jobCache:         jobs.Cache(),
		statefulSets:     statefulsets,
		statefulSetCache: statefulsets.Cache(),
		recorder:         management.NewRecorder("harvester-"+controllerAgentName, "", ""),
	}

	nodes.OnChange(ctx, controllerAgentName, controller.OnNodeChanged)

	jobs.OnChange(ctx, controllerAgentName, controller.OnJobChanged)
	jobs.OnRemove(ctx, controllerAgentName, controller.OnJobRemove)

	return nil
}
