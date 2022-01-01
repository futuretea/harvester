package backup

import (
	"fmt"
	"reflect"
	"strings"

	snapshotv1 "github.com/kubernetes-csi/external-snapshotter/v2/pkg/apis/volumesnapshot/v1beta1"
	lhv1beta1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
	kubevirtv1 "kubevirt.io/client-go/api/v1"

	harvesterv1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	"github.com/harvester/harvester/pkg/util"
)

func (h *Handler) updateConditions(vmBackup *harvesterv1.VirtualMachineBackup) error {
	__traceStack()

	var vmBackupCpy = vmBackup.DeepCopy()
	if isBackupProgressing(vmBackupCpy) {
		updateBackupCondition(vmBackupCpy, newProgressingCondition(corev1.ConditionTrue, "Operation in progress"))
		updateBackupCondition(vmBackupCpy, newReadyCondition(corev1.ConditionFalse, "Not ready"))
	}

	ready := true
	errorMessage := ""
	for _, vb := range vmBackup.Status.VolumeBackups {
		if vb.ReadyToUse == nil || !*vb.ReadyToUse {
			ready = false
		}

		if vb.Error != nil {
			errorMessage = fmt.Sprintf("VolumeSnapshot %s in error state", *vb.Name)
			break
		}
	}

	if ready && (vmBackupCpy.Status.ReadyToUse == nil || !*vmBackupCpy.Status.ReadyToUse) {
		vmBackupCpy.Status.CreationTime = currentTime()
		vmBackupCpy.Status.Error = nil
		updateBackupCondition(vmBackupCpy, newProgressingCondition(corev1.ConditionFalse, "Operation complete"))
		updateBackupCondition(vmBackupCpy, newReadyCondition(corev1.ConditionTrue, "Operation complete"))
	}

	if errorMessage != "" && (vmBackupCpy.Status.Error == nil || vmBackupCpy.Status.Error.Message == nil || *vmBackupCpy.Status.Error.Message != errorMessage) {
		vmBackupCpy.Status.Error = &harvesterv1.Error{
			Time:		currentTime(),
			Message:	pointer.StringPtr(errorMessage),
		}
	}

	vmBackupCpy.Status.ReadyToUse = pointer.BoolPtr(ready)

	if !reflect.DeepEqual(vmBackup.Status, vmBackupCpy.Status) {
		if _, err := h.vmBackups.Update(vmBackupCpy); err != nil {
			return err
		}
	}
	return nil
}

func (h *Handler) updateVolumeSnapshotChanged(key string, snapshot *snapshotv1.VolumeSnapshot) (*snapshotv1.VolumeSnapshot, error) {
	__traceStack()

	if snapshot == nil || snapshot.DeletionTimestamp != nil {
		return nil, nil
	}

	controllerRef := metav1.GetControllerOf(snapshot)

	if controllerRef != nil {
		ref := h.resolveVolSnapshotRef(snapshot.Namespace, controllerRef)
		if ref == nil {
			return nil, nil
		}
		h.vmBackupController.Enqueue(ref.Namespace, ref.Name)
	}
	return nil, nil
}

func (h *Handler) resolveVolSnapshotRef(namespace string, controllerRef *metav1.OwnerReference) *harvesterv1.VirtualMachineBackup {
	__traceStack()

	if controllerRef.Kind != vmBackupKind.Kind {
		return nil
	}
	backup, err := h.vmBackupCache.Get(namespace, controllerRef.Name)
	if err != nil {
		return nil
	}
	if backup.UID != controllerRef.UID {

		return nil
	}
	return backup
}

func (h *Handler) mountLonghornVolumes(vm *kubevirtv1.VirtualMachine) error {
	__traceStack()

	for _, vol := range vm.Spec.Template.Spec.Volumes {
		if vol.PersistentVolumeClaim == nil {
			continue
		}
		name := vol.PersistentVolumeClaim.ClaimName

		pvc, err := h.pvcCache.Get(vm.Namespace, name)
		if err != nil {
			return fmt.Errorf("failed to get pvc %s/%s, error: %s", name, vm.Namespace, err.Error())
		}

		volume, err := h.volumeCache.Get(util.LonghornSystemNamespaceName, pvc.Spec.VolumeName)
		if err != nil {
			return fmt.Errorf("failed to get volume %s/%s, error: %s", name, vm.Namespace, err.Error())
		}

		volCpy := volume.DeepCopy()
		if volume.Status.State == lhv1beta1.VolumeStateDetached || volume.Status.State == lhv1beta1.VolumeStateDetaching {
			volCpy.Spec.NodeID = volume.Status.OwnerID
		}

		if !reflect.DeepEqual(volCpy, volume) {
			logrus.Infof("mount detached volume %s to the node %s", volCpy.Name, volCpy.Spec.NodeID)
			if _, err = h.volumes.Update(volCpy); err != nil {
				return err
			}
		}
	}
	return nil
}

func getVolumeSnapshotContentName(volumeBackup harvesterv1.VolumeBackup) string {
	__traceStack()

	return fmt.Sprintf("%s-vsc", *volumeBackup.Name)
}

func (h *Handler) OnLHBackupChanged(key string, lhBackup *lhv1beta1.Backup) (*lhv1beta1.Backup, error) {
	__traceStack()

	if lhBackup == nil || lhBackup.DeletionTimestamp != nil || lhBackup.Status.SnapshotName == "" {
		return nil, nil
	}

	snapshotContent, err := h.snapshotContentCache.Get(strings.Replace(lhBackup.Status.SnapshotName, "snapshot", "snapcontent", 1))
	if err != nil {
		if !apierrors.IsNotFound(err) {
			return nil, err
		}
		return nil, nil
	}

	snapshot, err := h.snapshotCache.Get(snapshotContent.Spec.VolumeSnapshotRef.Namespace, snapshotContent.Spec.VolumeSnapshotRef.Name)
	if err != nil {
		return nil, err
	}

	controllerRef := metav1.GetControllerOf(snapshot)

	if controllerRef != nil {
		vmBackup := h.resolveVolSnapshotRef(snapshot.Namespace, controllerRef)
		if vmBackup == nil || vmBackup.Status == nil || vmBackup.Status.BackupTarget == nil {
			return nil, nil
		}

		vmBackupCpy := vmBackup.DeepCopy()
		for i, volumeBackup := range vmBackupCpy.Status.VolumeBackups {
			if *volumeBackup.Name == snapshot.Name {
				vmBackupCpy.Status.VolumeBackups[i].LonghornBackupName = pointer.StringPtr(lhBackup.Name)
			}
		}

		if !reflect.DeepEqual(vmBackup.Status, vmBackupCpy.Status) {
			if _, err := h.vmBackups.Update(vmBackupCpy); err != nil {
				return nil, err
			}
		}
	}
	return nil, nil
}
