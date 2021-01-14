package node

import (
	"sort"

	"github.com/sirupsen/logrus"

	corev1 "k8s.io/api/core/v1"
)

type SortByAgeNodeList []*corev1.Node

func (s SortByAgeNodeList) Len() int      { return len(s) }
func (s SortByAgeNodeList) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s SortByAgeNodeList) Less(i, j int) bool {
	// reverse order
	return !s[i].CreationTimestamp.Before(&s[j].CreationTimestamp)
}

// selectPromoteNodeList return a list contain nodes need to be promoted
// If the cluster doesn't need to be promoted, return nil
func (h *Handler) selectPromoteNodeList(nodeList []*corev1.Node) ([]*corev1.Node, error) {
	masterNumber := 0
	ignoreNodeNumber := 0
	canBePromoteNodeList := make([]*corev1.Node, 0, len(nodeList))
	for _, node := range nodeList {
		if h.shouldBeMaster(node) {
			masterNumber++
		} else {
			if h.isHarvesterNode(node) {
				if h.isHealthyNode(node) {
					canBePromoteNodeList = append(canBePromoteNodeList, node)
				}
			} else {
				ignoreNodeNumber++
			}
		}
	}

	if len(canBePromoteNodeList) == 0 {
		return nil, nil
	}

	// there is no need to promote if the spec number has been reached
	specMasterNumber := h.getSpecMasterNumber(len(nodeList) - ignoreNodeNumber)
	promoteNodeNumber := specMasterNumber - masterNumber
	if promoteNodeNumber <= 0 {
		return nil, nil
	}

	// sort by creation time from largest to smallest to find the youngest node that can be promoted
	promoteNodeList := make([]*corev1.Node, 0, promoteNodeNumber)
	var sortByAgeNodeList SortByAgeNodeList = canBePromoteNodeList
	sort.Sort(sortByAgeNodeList)
	for _, node := range sortByAgeNodeList {
		if promoteNodeNumber > 0 {
			promoteNodeList = append(promoteNodeList, node)
			promoteNodeNumber--
		}
	}

	return promoteNodeList, nil
}

// getSpecMasterNumber get spec master number by all node number
func (h *Handler) getSpecMasterNumber(nodeNumber int) int {
	if nodeNumber < 3 {
		return 1
	}
	return defaultSpecMasterNumber
}

// isHealthyNode determine whether it's an healthy node
func (h *Handler) isHealthyNode(node *corev1.Node) bool {
	for _, c := range node.Status.Conditions {
		if c.Type == corev1.NodeReady && c.Status != corev1.ConditionTrue {
			// skip unready nodes
			return false
		}

		if c.Type != corev1.NodeReady && c.Status == corev1.ConditionTrue {
			// skip node with conditions like nodeMemoryPressure, nodeDiskPressure, nodePIDPressure
			// and nodeNetworkUnavailable equal to true
			return false
		}
	}
	return true
}

// isHarvesterNode determine whether it's an Harvester node based on the node's label
func (h *Handler) isHarvesterNode(node *corev1.Node) bool {
	_, ok := node.Labels[HarvesterVersionLabelKey]
	return ok
}

// shouldBeMaster determine whether it should be an master node
func (h *Handler) shouldBeMaster(node *corev1.Node) bool {
	return h.isMasterRole(node) || h.hasBeenPromoted(node)
}

// isMasterRole determine whether it's an master node based on the node's label
func (h *Handler) isMasterRole(node *corev1.Node) bool {
	if value, ok := node.Labels[KubeMasterNodeLabelKey]; ok {
		return value == "true"
	}

	return false
}

func (h *Handler) hasBeenPromoted(node *corev1.Node) bool {
	return h.hasPromoteJob(node) || h.hasPromoteStatus(node)
}

func (h *Handler) hasPromoteJob(node *corev1.Node) bool {
	jobName := buildPromoteJobName(node.Name)
	_, err := h.jobCache.Get(promoteJobNamespaceName, jobName)
	return err == nil
}

func (h *Handler) hasPromoteStatus(node *corev1.Node) bool {
	_, ok := node.Annotations[HarvesterPromoteStatusAnnotationKey]
	return ok
}

func (h *Handler) isPromoteStatusIn(node *corev1.Node, statuses ...string) bool {
	status, ok := node.Annotations[HarvesterPromoteStatusAnnotationKey]
	if !ok {
		return false
	}

	for _, s := range statuses {
		if status == s {
			return true
		}
	}

	return false
}

func checkError(action string, err error) error {
	if err != nil {
		logrus.Errorf("%s failed, %s", action, err)
		return err
	}
	logrus.Infof("%s", action)
	return nil
}
