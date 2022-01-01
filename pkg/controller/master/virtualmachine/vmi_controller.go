package virtualmachine

import (
	"fmt"

	v1 "github.com/rancher/wrangler/pkg/generated/controllers/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	kubevirtapis "kubevirt.io/client-go/api/v1"

	kubevirtctrl "github.com/harvester/harvester/pkg/generated/controllers/kubevirt.io/v1"
)

type VMIController struct {
	virtualMachineCache	kubevirtctrl.VirtualMachineCache
	pvcClient		v1.PersistentVolumeClaimClient
	pvcCache		v1.PersistentVolumeClaimCache
}

func (h *VMIController) UnsetOwnerOfPVCs(_ string, vmi *kubevirtapis.VirtualMachineInstance) (*kubevirtapis.VirtualMachineInstance, error) {
	__traceStack()

	if vmi == nil || vmi.DeletionTimestamp == nil {
		return vmi, nil
	}

	var vmReferred = metav1.GetControllerOfNoCopy(vmi)
	if vmReferred == nil {

		return vmi, nil
	}
	var vmGVK = kubevirtapis.VirtualMachineGroupVersionKind
	if vmReferred.APIVersion != vmGVK.GroupVersion().String() ||
		vmReferred.Kind != vmGVK.Kind {

		return vmi, nil
	}

	var vm, err = h.virtualMachineCache.Get(vmi.Namespace, vmReferred.Name)
	if err != nil {
		if apierrors.IsNotFound(err) {

			return vmi, nil
		}
		return vmi, fmt.Errorf("failed to get VirtualMachine referred by VirtualMachineInstance(%s/%s): %w", vmi.Namespace, vmi.Name, err)
	}
	if vm.DeletionTimestamp != nil {

		return vmi, nil
	}

	var pvcNames = sets.String{}
	if vmiDesired := vm.Spec.Template; vmiDesired != nil {
		pvcNames = getPVCNames(&vmiDesired.Spec)
	}
	var pvcNameObserved = getPVCNames(&vmi.Spec)

	var pvcNamespace = vmi.Namespace
	var ownerlessPVCNames = pvcNameObserved.Difference(pvcNames).List()
	for _, pvcName := range ownerlessPVCNames {
		var pvc, err = h.pvcCache.Get(pvcNamespace, pvcName)
		if err != nil {
			if apierrors.IsNotFound(err) {

				continue
			}
			return vmi, fmt.Errorf("failed to get PVC(%s/%s): %w", pvcNamespace, pvcName, err)
		}

		err = unsetBoundedPVCReference(h.pvcClient, pvc, vm)
		if err != nil {
			return vmi, fmt.Errorf("failed to revoke VitrualMachine(%s/%s) as PVC(%s/%s)'s owner: %w",
				vm.Namespace, vm.Name, pvcNamespace, pvcName, err)
		}
	}

	return vmi, nil
}
