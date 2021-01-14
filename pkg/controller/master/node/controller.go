package node

import (
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	ctlappsv1 "github.com/rancher/wrangler-api/pkg/generated/controllers/apps/v1"
	ctlbatchv1 "github.com/rancher/wrangler-api/pkg/generated/controllers/batch/v1"
	ctlcorev1 "github.com/rancher/wrangler-api/pkg/generated/controllers/core/v1"
)

// Handler
type Handler struct {
	podCache         ctlcorev1.PodCache
	nodes            ctlcorev1.NodeClient
	nodeCache        ctlcorev1.NodeCache
	jobs             ctlbatchv1.JobClient
	jobCache         ctlbatchv1.JobCache
	statefulSets     ctlappsv1.StatefulSetClient
	statefulSetCache ctlappsv1.StatefulSetCache
}

// OnNodeChanged automate the upgrade of node roles and balance of minio pods
// If the number of masters in the cluster is less than spec number,
// the harvester node with the smallest age is automatically promoted to be master.
// If minio is deployed in distribute mode, make minio pods balanced.
// Minio balancing only takes place when the node is ready and not in the upgrade state.
func (h *Handler) OnNodeChanged(key string, node *corev1.Node) (*corev1.Node, error) {
	if node == nil || node.DeletionTimestamp != nil {
		return node, nil
	}

	nodeList, err := h.nodeCache.List(labels.Everything())
	if err != nil {
		return nil, err
	}

	promoteNodeList, err := h.selectPromoteNodeList(nodeList)
	if err != nil {
		return nil, err
	}

	switch len(promoteNodeList) {
	case 0:
		return h.makeMinioBalanced(key, node)
	case 1:
		return h.promote(promoteNodeList[0])
	default:
		for _, promoteNode := range promoteNodeList {
			if _, err = h.promote(promoteNode); err != nil {
				return nil, err
			}
		}
		return nil, nil
	}
}

// OnJobChanged
// If the node corresponding to the promote job has been removed, delete the job.
// If the promote job executes successfully, the node's promote status will be marked as complete and schedulable
// If the promote job fails, the node's promote status will be marked as failed.
func (h *Handler) OnJobChanged(key string, job *batchv1.Job) (*batchv1.Job, error) {
	if job == nil || job.DeletionTimestamp != nil {
		return job, nil
	}

	nodeName, ok := job.Labels[HarvesterPromoteNodeLabelKey]
	if !ok {
		return job, nil
	}

	node, err := h.nodeCache.Get(nodeName)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return job, h.deleteJob(job, metav1.DeletePropagationBackground)
		}
		return job, err
	}

	if ConditionJobComplete.IsTrue(job) {
		return h.setPromoteResult(job, node, PromoteStatusComplete)
	}

	if ConditionJobFailed.IsTrue(job) {
		return h.setPromoteResult(job, node, PromoteStatusFailed)
	}

	return job, nil
}

// OnJobRemove
// If the running promote job is deleted, the node's promote status will be marked as unknown
func (h *Handler) OnJobRemove(key string, job *batchv1.Job) (*batchv1.Job, error) {
	if job == nil {
		return job, nil
	}

	nodeName, ok := job.Labels[HarvesterPromoteNodeLabelKey]
	if !ok {
		return job, nil
	}
	if ConditionJobFailed.IsTrue(job) || ConditionJobComplete.IsTrue(job) {
		return job, nil
	}

	node, err := h.nodeCache.Get(nodeName)
	switch {
	case apierrors.IsNotFound(err):
		return job, nil
	case err != nil:
		return job, err
	}

	if h.isPromoteStatusIn(node, PromoteStatusRunning) {
		return h.setPromoteResult(job, node, PromoteStatusUnknown)
	}

	return job, nil
}
