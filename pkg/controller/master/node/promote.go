package node

import (
	"fmt"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
)

const (
	KubeNodeRoleLabelPrefix = "node-role.kubernetes.io/"
	KubeMasterNodeLabelKey  = KubeNodeRoleLabelPrefix + "master"

	KubeNodeSVCLabelPrefix      = "svccontroller.k3s.cattle.io/"
	KubeNodeSVCEnableLBLabelKey = KubeNodeSVCLabelPrefix + "enablelb"

	HarvesterLabelAnnotationPrefix      = "harvester.cattle.io/"
	HarvesterVersionLabelKey            = HarvesterLabelAnnotationPrefix + "version"
	HarvesterPromoteNodeLabelKey        = HarvesterLabelAnnotationPrefix + "promote-node"
	HarvesterPromoteStatusAnnotationKey = HarvesterLabelAnnotationPrefix + "promote-status"

	PromoteStatusComplete = "complete"
	PromoteStatusRunning  = "running"
	PromoteStatusUnknown  = "unknown"
	PromoteStatusFailed   = "failed"

	defaultSpecMasterNumber = 3
)

func (h *Handler) promote(node *corev1.Node) (*corev1.Node, error) {
	// first, mark node into promote status
	startedNode, err := h.setPromoteStart(node)
	if err != nil {
		return nil, err
	}

	// then, create a promote job on the node
	if _, err := h.createPromoteJob(node); err != nil {
		return nil, err
	}

	return startedNode, nil
}

// setPromoteStart set node unschedulable and set promote status running.
func (h *Handler) setPromoteStart(node *corev1.Node) (*corev1.Node, error) {
	toUpdate := node.DeepCopy()
	toUpdate.Labels[KubeNodeSVCEnableLBLabelKey] = "true"
	toUpdate.Annotations[HarvesterPromoteStatusAnnotationKey] = PromoteStatusRunning
	toUpdate.Spec.Unschedulable = true
	updatedNode, err := h.nodes.Update(toUpdate)
	return updatedNode, checkError(fmt.Sprintf("setPromoteStart on node %s", node.Name), err)
}

// setPromoteResult set node schedulable and update promote status if the promote is successful
func (h *Handler) setPromoteResult(job *batchv1.Job, node *corev1.Node, status string) (*batchv1.Job, error) {
	toUpdate := node.DeepCopy()
	toUpdate.Annotations[HarvesterPromoteStatusAnnotationKey] = status
	if status == PromoteStatusComplete {
		toUpdate.Spec.Unschedulable = false
	}
	_, err := h.nodes.Update(toUpdate)
	return job, checkError(fmt.Sprintf("setPromoteResult to %s on node %s", status, node.Name), err)
}
