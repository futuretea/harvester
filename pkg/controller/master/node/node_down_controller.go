package node

import (
	"context"
	"fmt"
	"time"

	ctlcorev1 "github.com/rancher/wrangler/pkg/generated/controllers/core/v1"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	kv1 "kubevirt.io/client-go/api/v1"

	harvesterv1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	"github.com/harvester/harvester/pkg/config"
	v1 "github.com/harvester/harvester/pkg/generated/controllers/kubevirt.io/v1"
	"github.com/harvester/harvester/pkg/settings"
)

const (
	nodeDownControllerName = "node-down-controller"
)

type nodeDownHandler struct {
	nodes				ctlcorev1.NodeController
	nodeCache			ctlcorev1.NodeCache
	pods				ctlcorev1.PodClient
	virtualMachineInstanceCache	v1.VirtualMachineInstanceCache
}

func NodeDownRegister(ctx context.Context, management *config.Management, options config.Options) error {
	__traceStack()

	nodes := management.CoreFactory.Core().V1().Node()
	pods := management.CoreFactory.Core().V1().Pod()
	setting := management.HarvesterFactory.Harvesterhci().V1beta1().Setting()
	vmis := management.VirtFactory.Kubevirt().V1().VirtualMachineInstance()
	nodeDownHandler := &nodeDownHandler{
		nodes:				nodes,
		nodeCache:			nodes.Cache(),
		pods:				pods,
		virtualMachineInstanceCache:	vmis.Cache(),
	}

	nodes.OnChange(ctx, nodeDownControllerName, nodeDownHandler.OnNodeChanged)
	setting.OnChange(ctx, nodeDownControllerName, nodeDownHandler.OnVMForceResetPolicyChanged)

	return nil
}

func (h *nodeDownHandler) OnNodeChanged(key string, node *corev1.Node) (*corev1.Node, error) {
	__traceStack()

	if node == nil || node.DeletionTimestamp != nil {
		return node, nil
	}

	cond := getNodeCondition(node.Status.Conditions, corev1.NodeReady)
	if cond == nil {
		return node, fmt.Errorf("can't find %s condition in node %s", corev1.NodeReady, node.Name)
	}

	if cond.Status == corev1.ConditionTrue {
		return node, nil
	}

	vmForceResetPolicy, err := settings.DecodeVMForceResetPolicy(settings.VMForceResetPolicySet.Get())
	if err != nil {
		return node, err
	}

	if !vmForceResetPolicy.Enable {
		return node, nil
	}

	if time.Since(cond.LastTransitionTime.Time) < time.Duration(vmForceResetPolicy.Period)*time.Second {
		deadline := cond.LastTransitionTime.Add(time.Duration(vmForceResetPolicy.Period) * time.Second)
		logrus.Debugf("Enqueue node event again at %v", deadline)
		h.nodes.EnqueueAfter(node.Name, time.Until(deadline))
		return node, nil
	}

	pods, err := h.pods.List(corev1.NamespaceAll, metav1.ListOptions{
		LabelSelector: labels.Set{
			kv1.AppLabel: "virt-launcher",
		}.String(),
		FieldSelector:	"spec.nodeName=" + node.Name,
	})
	if err != nil {
		return node, err
	}

	gracePeriod := int64(0)
	for _, pod := range pods.Items {
		logrus.Debugf("force delete pod %s/%s", pod.Namespace, pod.Name)
		if err := h.pods.Delete(
			pod.Namespace,
			pod.Name,
			&metav1.DeleteOptions{
				GracePeriodSeconds: &gracePeriod,
			}); err != nil {
			return node, err
		}

	}
	return node, nil
}

func (h *nodeDownHandler) OnVMForceResetPolicyChanged(key string, setting *harvesterv1.Setting) (*harvesterv1.Setting, error) {
	__traceStack()

	if setting == nil || setting.DeletionTimestamp != nil ||
		setting.Name != settings.VMForceResetPolicySettingName || setting.Value == "" {
		return setting, nil
	}

	vmForceResetPolicy, err := settings.DecodeVMForceResetPolicy(setting.Value)
	if err != nil {
		return setting, err
	}

	if !vmForceResetPolicy.Enable {
		return setting, nil
	}

	nodes, err := h.nodeCache.List(labels.Everything())
	if err != nil {
		return setting, err
	}

	for _, node := range nodes {
		cond := getNodeCondition(node.Status.Conditions, corev1.NodeReady)
		if cond != nil && cond.Status != corev1.ConditionTrue {
			h.nodes.Enqueue(node.Name)
		}
	}
	return setting, nil
}

func getNodeCondition(conditions []corev1.NodeCondition, conditionType corev1.NodeConditionType) *corev1.NodeCondition {
	__traceStack()

	var cond *corev1.NodeCondition
	for _, c := range conditions {
		if c.Type == conditionType {
			cond = &c
			break
		}
	}
	return cond
}
