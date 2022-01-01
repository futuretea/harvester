package virtualmachine

import (
	"context"

	"github.com/harvester/harvester/pkg/config"
)

const (
	vmControllerCreatePVCsFromAnnotationControllerName	= "VMController.CreatePVCsFromAnnotation"
	vmControllerSetOwnerOfPVCsControllerName		= "VMController.SetOwnerOfPVCs"
	vmControllerUnsetOwnerOfPVCsControllerName		= "VMController.UnsetOwnerOfPVCs"
	vmiControllerUnsetOwnerOfPVCsControllerName		= "VMIController.UnsetOwnerOfPVCs"
	vmControllerSetDefaultManagementNetworkMac		= "VMController.SetDefaultManagementNetworkMacAddress"
)

func Register(ctx context.Context, management *config.Management, options config.Options) error {
	__traceStack()

	var pvcClient = management.CoreFactory.Core().V1().PersistentVolumeClaim()
	var pvcCache = pvcClient.Cache()

	var vmCtrl = &VMController{
		pvcClient:	pvcClient,
		pvcCache:	pvcCache,
	}
	var virtualMachineClient = management.VirtFactory.Kubevirt().V1().VirtualMachine()
	virtualMachineClient.OnChange(ctx, vmControllerCreatePVCsFromAnnotationControllerName, vmCtrl.createPVCsFromAnnotation)
	virtualMachineClient.OnChange(ctx, vmControllerSetOwnerOfPVCsControllerName, vmCtrl.SetOwnerOfPVCs)
	virtualMachineClient.OnRemove(ctx, vmControllerUnsetOwnerOfPVCsControllerName, vmCtrl.UnsetOwnerOfPVCs)

	var virtualMachineCache = virtualMachineClient.Cache()
	var vmiCtrl = &VMIController{
		virtualMachineCache:	virtualMachineCache,
		pvcClient:		pvcClient,
		pvcCache:		pvcCache,
	}
	var virtualMachineInstanceClient = management.VirtFactory.Kubevirt().V1().VirtualMachineInstance()
	virtualMachineInstanceClient.OnRemove(ctx, vmiControllerUnsetOwnerOfPVCsControllerName, vmiCtrl.UnsetOwnerOfPVCs)

	var (
		vmClient	= management.VirtFactory.Kubevirt().V1().VirtualMachine()
		vmCache		= management.VirtFactory.Kubevirt().V1().VirtualMachine().Cache()
		vmiClient	= management.VirtFactory.Kubevirt().V1().VirtualMachineInstance()
	)
	var vmNetworkCtl = &VMNetworkController{
		vmClient:	vmClient,
		vmCache:	vmCache,
		vmiClient:	vmiClient,
	}
	virtualMachineInstanceClient.OnChange(ctx, vmControllerSetDefaultManagementNetworkMac, vmNetworkCtl.SetDefaultNetworkMacAddress)

	return nil
}
