package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	types "k8s.io/apimachinery/pkg/types"
)

func (in *BackupTarget) DeepCopyInto(out *BackupTarget) {
	__traceStack()

	*out = *in
	return
}

func (in *BackupTarget) DeepCopy() *BackupTarget {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(BackupTarget)
	in.DeepCopyInto(out)
	return out
}

func (in *Condition) DeepCopyInto(out *Condition) {
	__traceStack()

	*out = *in
	return
}

func (in *Condition) DeepCopy() *Condition {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

func (in *Error) DeepCopyInto(out *Error) {
	__traceStack()

	*out = *in
	if in.Time != nil {
		in, out := &in.Time, &out.Time
		*out = (*in).DeepCopy()
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	return
}

func (in *Error) DeepCopy() *Error {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(Error)
	in.DeepCopyInto(out)
	return out
}

func (in *ErrorResponse) DeepCopyInto(out *ErrorResponse) {
	__traceStack()

	*out = *in
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

func (in *ErrorResponse) DeepCopy() *ErrorResponse {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(ErrorResponse)
	in.DeepCopyInto(out)
	return out
}

func (in *KeyGenInput) DeepCopyInto(out *KeyGenInput) {
	__traceStack()

	*out = *in
	return
}

func (in *KeyGenInput) DeepCopy() *KeyGenInput {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(KeyGenInput)
	in.DeepCopyInto(out)
	return out
}

func (in *KeyPair) DeepCopyInto(out *KeyPair) {
	__traceStack()

	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

func (in *KeyPair) DeepCopy() *KeyPair {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(KeyPair)
	in.DeepCopyInto(out)
	return out
}

func (in *KeyPair) DeepCopyObject() runtime.Object {
	__traceStack()

	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *KeyPairList) DeepCopyInto(out *KeyPairList) {
	__traceStack()

	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeyPair, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *KeyPairList) DeepCopy() *KeyPairList {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(KeyPairList)
	in.DeepCopyInto(out)
	return out
}

func (in *KeyPairList) DeepCopyObject() runtime.Object {
	__traceStack()

	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *KeyPairSpec) DeepCopyInto(out *KeyPairSpec) {
	__traceStack()

	*out = *in
	return
}

func (in *KeyPairSpec) DeepCopy() *KeyPairSpec {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(KeyPairSpec)
	in.DeepCopyInto(out)
	return out
}

func (in *KeyPairStatus) DeepCopyInto(out *KeyPairStatus) {
	__traceStack()

	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
	return
}

func (in *KeyPairStatus) DeepCopy() *KeyPairStatus {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(KeyPairStatus)
	in.DeepCopyInto(out)
	return out
}

func (in *NodeUpgradeStatus) DeepCopyInto(out *NodeUpgradeStatus) {
	__traceStack()

	*out = *in
	return
}

func (in *NodeUpgradeStatus) DeepCopy() *NodeUpgradeStatus {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(NodeUpgradeStatus)
	in.DeepCopyInto(out)
	return out
}

func (in *PersistentVolumeClaimSourceSpec) DeepCopyInto(out *PersistentVolumeClaimSourceSpec) {
	__traceStack()

	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

func (in *PersistentVolumeClaimSourceSpec) DeepCopy() *PersistentVolumeClaimSourceSpec {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(PersistentVolumeClaimSourceSpec)
	in.DeepCopyInto(out)
	return out
}

func (in *Preference) DeepCopyInto(out *Preference) {
	__traceStack()

	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

func (in *Preference) DeepCopy() *Preference {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(Preference)
	in.DeepCopyInto(out)
	return out
}

func (in *Preference) DeepCopyObject() runtime.Object {
	__traceStack()

	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *PreferenceList) DeepCopyInto(out *PreferenceList) {
	__traceStack()

	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Preference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *PreferenceList) DeepCopy() *PreferenceList {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(PreferenceList)
	in.DeepCopyInto(out)
	return out
}

func (in *PreferenceList) DeepCopyObject() runtime.Object {
	__traceStack()

	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *SecretBackup) DeepCopyInto(out *SecretBackup) {
	__traceStack()

	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make(map[string][]byte, len(*in))
		for key, val := range *in {
			var outVal []byte
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]byte, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	return
}

func (in *SecretBackup) DeepCopy() *SecretBackup {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(SecretBackup)
	in.DeepCopyInto(out)
	return out
}

func (in *Setting) DeepCopyInto(out *Setting) {
	__traceStack()

	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Status.DeepCopyInto(&out.Status)
	return
}

func (in *Setting) DeepCopy() *Setting {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(Setting)
	in.DeepCopyInto(out)
	return out
}

func (in *Setting) DeepCopyObject() runtime.Object {
	__traceStack()

	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *SettingList) DeepCopyInto(out *SettingList) {
	__traceStack()

	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Setting, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *SettingList) DeepCopy() *SettingList {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(SettingList)
	in.DeepCopyInto(out)
	return out
}

func (in *SettingList) DeepCopyObject() runtime.Object {
	__traceStack()

	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *SettingStatus) DeepCopyInto(out *SettingStatus) {
	__traceStack()

	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
	return
}

func (in *SettingStatus) DeepCopy() *SettingStatus {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(SettingStatus)
	in.DeepCopyInto(out)
	return out
}

func (in *SupportBundle) DeepCopyInto(out *SupportBundle) {
	__traceStack()

	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

func (in *SupportBundle) DeepCopy() *SupportBundle {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(SupportBundle)
	in.DeepCopyInto(out)
	return out
}

func (in *SupportBundle) DeepCopyObject() runtime.Object {
	__traceStack()

	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *SupportBundleList) DeepCopyInto(out *SupportBundleList) {
	__traceStack()

	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SupportBundle, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *SupportBundleList) DeepCopy() *SupportBundleList {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(SupportBundleList)
	in.DeepCopyInto(out)
	return out
}

func (in *SupportBundleList) DeepCopyObject() runtime.Object {
	__traceStack()

	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *SupportBundleSpec) DeepCopyInto(out *SupportBundleSpec) {
	__traceStack()

	*out = *in
	return
}

func (in *SupportBundleSpec) DeepCopy() *SupportBundleSpec {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(SupportBundleSpec)
	in.DeepCopyInto(out)
	return out
}

func (in *SupportBundleStatus) DeepCopyInto(out *SupportBundleStatus) {
	__traceStack()

	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
	return
}

func (in *SupportBundleStatus) DeepCopy() *SupportBundleStatus {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(SupportBundleStatus)
	in.DeepCopyInto(out)
	return out
}

func (in *Upgrade) DeepCopyInto(out *Upgrade) {
	__traceStack()

	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

func (in *Upgrade) DeepCopy() *Upgrade {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(Upgrade)
	in.DeepCopyInto(out)
	return out
}

func (in *Upgrade) DeepCopyObject() runtime.Object {
	__traceStack()

	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *UpgradeList) DeepCopyInto(out *UpgradeList) {
	__traceStack()

	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Upgrade, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *UpgradeList) DeepCopy() *UpgradeList {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(UpgradeList)
	in.DeepCopyInto(out)
	return out
}

func (in *UpgradeList) DeepCopyObject() runtime.Object {
	__traceStack()

	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *UpgradeSpec) DeepCopyInto(out *UpgradeSpec) {
	__traceStack()

	*out = *in
	return
}

func (in *UpgradeSpec) DeepCopy() *UpgradeSpec {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(UpgradeSpec)
	in.DeepCopyInto(out)
	return out
}

func (in *UpgradeStatus) DeepCopyInto(out *UpgradeStatus) {
	__traceStack()

	*out = *in
	if in.NodeStatuses != nil {
		in, out := &in.NodeStatuses, &out.NodeStatuses
		*out = make(map[string]NodeUpgradeStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
	return
}

func (in *UpgradeStatus) DeepCopy() *UpgradeStatus {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(UpgradeStatus)
	in.DeepCopyInto(out)
	return out
}

func (in *Version) DeepCopyInto(out *Version) {
	__traceStack()

	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

func (in *Version) DeepCopy() *Version {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(Version)
	in.DeepCopyInto(out)
	return out
}

func (in *Version) DeepCopyObject() runtime.Object {
	__traceStack()

	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *VersionList) DeepCopyInto(out *VersionList) {
	__traceStack()

	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Version, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *VersionList) DeepCopy() *VersionList {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(VersionList)
	in.DeepCopyInto(out)
	return out
}

func (in *VersionList) DeepCopyObject() runtime.Object {
	__traceStack()

	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *VersionSpec) DeepCopyInto(out *VersionSpec) {
	__traceStack()

	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

func (in *VersionSpec) DeepCopy() *VersionSpec {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(VersionSpec)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMachineBackup) DeepCopyInto(out *VirtualMachineBackup) {
	__traceStack()

	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(VirtualMachineBackupStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

func (in *VirtualMachineBackup) DeepCopy() *VirtualMachineBackup {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(VirtualMachineBackup)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMachineBackup) DeepCopyObject() runtime.Object {
	__traceStack()

	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *VirtualMachineBackupList) DeepCopyInto(out *VirtualMachineBackupList) {
	__traceStack()

	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualMachineBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *VirtualMachineBackupList) DeepCopy() *VirtualMachineBackupList {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(VirtualMachineBackupList)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMachineBackupList) DeepCopyObject() runtime.Object {
	__traceStack()

	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *VirtualMachineBackupSpec) DeepCopyInto(out *VirtualMachineBackupSpec) {
	__traceStack()

	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	return
}

func (in *VirtualMachineBackupSpec) DeepCopy() *VirtualMachineBackupSpec {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(VirtualMachineBackupSpec)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMachineBackupStatus) DeepCopyInto(out *VirtualMachineBackupStatus) {
	__traceStack()

	*out = *in
	if in.SourceUID != nil {
		in, out := &in.SourceUID, &out.SourceUID
		*out = new(types.UID)
		**out = **in
	}
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = (*in).DeepCopy()
	}
	if in.BackupTarget != nil {
		in, out := &in.BackupTarget, &out.BackupTarget
		*out = new(BackupTarget)
		**out = **in
	}
	if in.SourceSpec != nil {
		in, out := &in.SourceSpec, &out.SourceSpec
		*out = new(VirtualMachineSourceSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.VolumeBackups != nil {
		in, out := &in.VolumeBackups, &out.VolumeBackups
		*out = make([]VolumeBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecretBackups != nil {
		in, out := &in.SecretBackups, &out.SecretBackups
		*out = make([]SecretBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ReadyToUse != nil {
		in, out := &in.ReadyToUse, &out.ReadyToUse
		*out = new(bool)
		**out = **in
	}
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = new(Error)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
	return
}

func (in *VirtualMachineBackupStatus) DeepCopy() *VirtualMachineBackupStatus {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(VirtualMachineBackupStatus)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMachineImage) DeepCopyInto(out *VirtualMachineImage) {
	__traceStack()

	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

func (in *VirtualMachineImage) DeepCopy() *VirtualMachineImage {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(VirtualMachineImage)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMachineImage) DeepCopyObject() runtime.Object {
	__traceStack()

	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *VirtualMachineImageList) DeepCopyInto(out *VirtualMachineImageList) {
	__traceStack()

	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualMachineImage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *VirtualMachineImageList) DeepCopy() *VirtualMachineImageList {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(VirtualMachineImageList)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMachineImageList) DeepCopyObject() runtime.Object {
	__traceStack()

	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *VirtualMachineImageSpec) DeepCopyInto(out *VirtualMachineImageSpec) {
	__traceStack()

	*out = *in
	return
}

func (in *VirtualMachineImageSpec) DeepCopy() *VirtualMachineImageSpec {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(VirtualMachineImageSpec)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMachineImageStatus) DeepCopyInto(out *VirtualMachineImageStatus) {
	__traceStack()

	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
	return
}

func (in *VirtualMachineImageStatus) DeepCopy() *VirtualMachineImageStatus {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(VirtualMachineImageStatus)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMachineRestore) DeepCopyInto(out *VirtualMachineRestore) {
	__traceStack()

	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(VirtualMachineRestoreStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

func (in *VirtualMachineRestore) DeepCopy() *VirtualMachineRestore {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(VirtualMachineRestore)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMachineRestore) DeepCopyObject() runtime.Object {
	__traceStack()

	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *VirtualMachineRestoreList) DeepCopyInto(out *VirtualMachineRestoreList) {
	__traceStack()

	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualMachineRestore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *VirtualMachineRestoreList) DeepCopy() *VirtualMachineRestoreList {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(VirtualMachineRestoreList)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMachineRestoreList) DeepCopyObject() runtime.Object {
	__traceStack()

	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *VirtualMachineRestoreSpec) DeepCopyInto(out *VirtualMachineRestoreSpec) {
	__traceStack()

	*out = *in
	in.Target.DeepCopyInto(&out.Target)
	return
}

func (in *VirtualMachineRestoreSpec) DeepCopy() *VirtualMachineRestoreSpec {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(VirtualMachineRestoreSpec)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMachineRestoreStatus) DeepCopyInto(out *VirtualMachineRestoreStatus) {
	__traceStack()

	*out = *in
	if in.VolumeRestores != nil {
		in, out := &in.VolumeRestores, &out.VolumeRestores
		*out = make([]VolumeRestore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RestoreTime != nil {
		in, out := &in.RestoreTime, &out.RestoreTime
		*out = (*in).DeepCopy()
	}
	if in.DeletedVolumes != nil {
		in, out := &in.DeletedVolumes, &out.DeletedVolumes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Complete != nil {
		in, out := &in.Complete, &out.Complete
		*out = new(bool)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
	if in.TargetUID != nil {
		in, out := &in.TargetUID, &out.TargetUID
		*out = new(types.UID)
		**out = **in
	}
	return
}

func (in *VirtualMachineRestoreStatus) DeepCopy() *VirtualMachineRestoreStatus {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(VirtualMachineRestoreStatus)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMachineSourceSpec) DeepCopyInto(out *VirtualMachineSourceSpec) {
	__traceStack()

	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

func (in *VirtualMachineSourceSpec) DeepCopy() *VirtualMachineSourceSpec {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(VirtualMachineSourceSpec)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMachineTemplate) DeepCopyInto(out *VirtualMachineTemplate) {
	__traceStack()

	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

func (in *VirtualMachineTemplate) DeepCopy() *VirtualMachineTemplate {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(VirtualMachineTemplate)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMachineTemplate) DeepCopyObject() runtime.Object {
	__traceStack()

	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *VirtualMachineTemplateList) DeepCopyInto(out *VirtualMachineTemplateList) {
	__traceStack()

	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualMachineTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *VirtualMachineTemplateList) DeepCopy() *VirtualMachineTemplateList {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(VirtualMachineTemplateList)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMachineTemplateList) DeepCopyObject() runtime.Object {
	__traceStack()

	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *VirtualMachineTemplateSpec) DeepCopyInto(out *VirtualMachineTemplateSpec) {
	__traceStack()

	*out = *in
	return
}

func (in *VirtualMachineTemplateSpec) DeepCopy() *VirtualMachineTemplateSpec {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(VirtualMachineTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMachineTemplateStatus) DeepCopyInto(out *VirtualMachineTemplateStatus) {
	__traceStack()

	*out = *in
	return
}

func (in *VirtualMachineTemplateStatus) DeepCopy() *VirtualMachineTemplateStatus {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(VirtualMachineTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMachineTemplateVersion) DeepCopyInto(out *VirtualMachineTemplateVersion) {
	__traceStack()

	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

func (in *VirtualMachineTemplateVersion) DeepCopy() *VirtualMachineTemplateVersion {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(VirtualMachineTemplateVersion)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMachineTemplateVersion) DeepCopyObject() runtime.Object {
	__traceStack()

	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *VirtualMachineTemplateVersionList) DeepCopyInto(out *VirtualMachineTemplateVersionList) {
	__traceStack()

	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualMachineTemplateVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *VirtualMachineTemplateVersionList) DeepCopy() *VirtualMachineTemplateVersionList {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(VirtualMachineTemplateVersionList)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMachineTemplateVersionList) DeepCopyObject() runtime.Object {
	__traceStack()

	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *VirtualMachineTemplateVersionSpec) DeepCopyInto(out *VirtualMachineTemplateVersionSpec) {
	__traceStack()

	*out = *in
	if in.KeyPairIDs != nil {
		in, out := &in.KeyPairIDs, &out.KeyPairIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.VM.DeepCopyInto(&out.VM)
	return
}

func (in *VirtualMachineTemplateVersionSpec) DeepCopy() *VirtualMachineTemplateVersionSpec {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(VirtualMachineTemplateVersionSpec)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMachineTemplateVersionStatus) DeepCopyInto(out *VirtualMachineTemplateVersionStatus) {
	__traceStack()

	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
	return
}

func (in *VirtualMachineTemplateVersionStatus) DeepCopy() *VirtualMachineTemplateVersionStatus {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(VirtualMachineTemplateVersionStatus)
	in.DeepCopyInto(out)
	return out
}

func (in *VolumeBackup) DeepCopyInto(out *VolumeBackup) {
	__traceStack()

	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = (*in).DeepCopy()
	}
	in.PersistentVolumeClaim.DeepCopyInto(&out.PersistentVolumeClaim)
	if in.LonghornBackupName != nil {
		in, out := &in.LonghornBackupName, &out.LonghornBackupName
		*out = new(string)
		**out = **in
	}
	if in.ReadyToUse != nil {
		in, out := &in.ReadyToUse, &out.ReadyToUse
		*out = new(bool)
		**out = **in
	}
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = new(Error)
		(*in).DeepCopyInto(*out)
	}
	return
}

func (in *VolumeBackup) DeepCopy() *VolumeBackup {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(VolumeBackup)
	in.DeepCopyInto(out)
	return out
}

func (in *VolumeRestore) DeepCopyInto(out *VolumeRestore) {
	__traceStack()

	*out = *in
	in.PersistentVolumeClaim.DeepCopyInto(&out.PersistentVolumeClaim)
	return
}

func (in *VolumeRestore) DeepCopy() *VolumeRestore {
	__traceStack()

	if in == nil {
		return nil
	}
	out := new(VolumeRestore)
	in.DeepCopyInto(out)
	return out
}
