package node

import (
	"time"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/rancher/harvester/pkg/config"
)

const (
	appLabelKey      = "app"
	minioName        = "minio"
	timestampAnnoKey = "cattle.io/timestamp"
)

var (
	throttleDelay = 1 * time.Minute
)

// makeMinioBalanced tries to make minio pods balanced if they are not
func (h *Handler) makeMinioBalanced(key string, node *corev1.Node) (*corev1.Node, error) {

	for _, c := range node.Status.Conditions {
		if c.Type == corev1.NodeReady && c.Status != corev1.ConditionTrue {
			// skip unready nodes
			return node, nil
		}

		if c.Type != corev1.NodeReady && c.Status == corev1.ConditionTrue {
			// skip deploy minio to node with conditions like nodeMemoryPressure, nodeDiskPressure, nodePIDPressure
			// and nodeNetworkUnavailable equal to true
			return node, nil
		}
	}

	if len(node.Spec.Taints) > 0 {
		// skip taints nodes
		return node, nil
	}

	if h.isPromoteStatusIn(node, PromoteStatusUnknown, PromoteStatusRunning, PromoteStatusFailed) {
		return node, nil
	}

	sets := labels.Set{
		appLabelKey: minioName,
	}
	pods, err := h.podCache.List(config.Namespace, sets.AsSelector())
	if err != nil {
		return nil, err
	}
	if len(pods) < 4 {
		// only take care of distributed minio
		return node, nil
	}

	var nodeSet = make(map[string]bool)
	for _, p := range pods {
		if p.Status.Phase != corev1.PodRunning {
			// proceed when all pods are running
			return node, nil
		}
		if p.Spec.NodeName == node.Name {
			// The node is already running minio pods
			return node, nil
		}
		nodeSet[p.Spec.NodeName] = true
	}
	if len(nodeSet) < 3 {
		// balance minio pods to tolerate node disruption
		if err := h.redeployMinio(); err != nil {
			return node, err
		}
	}
	return node, nil
}

func (h *Handler) redeployMinio() error {
	ss, err := h.statefulSetCache.Get(config.Namespace, minioName)
	if err != nil && !apierrors.IsNotFound(err) {
		return err
	}

	toUpdate := ss.DeepCopy()
	if toUpdate.Spec.Template.Annotations == nil {
		toUpdate.Spec.Template.Annotations = make(map[string]string, 1)
	}
	prevTimestamp := toUpdate.Spec.Template.Annotations[timestampAnnoKey]
	if prevTimestamp != "" {
		prevTime, err := time.Parse(time.RFC3339, prevTimestamp)
		if err != nil {
			return err
		}
		if prevTime.Add(throttleDelay).After(time.Now()) {
			return nil
		}
	}
	toUpdate.Spec.Template.Annotations[timestampAnnoKey] = time.Now().Format(time.RFC3339)
	_, err = h.statefulSets.Update(toUpdate)
	return err
}
