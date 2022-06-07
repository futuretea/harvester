package virtualmachineimage

import (
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	harvesterv1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	"github.com/harvester/harvester/pkg/generated/clientset/versioned/fake"
	"github.com/harvester/harvester/pkg/settings"
	"github.com/harvester/harvester/pkg/util/fakeclients"
)

func Test_virtualmachineimage_mutator(t *testing.T) {
	tests := []struct {
		name      string
		oldParams map[string]string
		newParams map[string]string
		setting   string
		patchOps  []string
	}{
		{
			name:      "[create] new params is nil",
			oldParams: nil,
			newParams: nil,
			patchOps: []string{
				`{"op": "add", "path": "/spec/extraStorageClassParameters", "value": {"migratable":"true","numberOfReplicas":"3","staleReplicaTimeout":"30"}}`,
			},
		},
		{
			name:      "[create] new params is empty",
			oldParams: nil,
			newParams: map[string]string{},
			patchOps: []string{
				`{"op": "replace", "path": "/spec/extraStorageClassParameters", "value": {"migratable":"true","numberOfReplicas":"3","staleReplicaTimeout":"30"}}`,
			},
		},
		{
			name:      "[create] only overwrite default setting",
			oldParams: nil,
			newParams: map[string]string{
				"numberOfReplicas": "1",
			},
			patchOps: []string{
				`{"op": "replace", "path": "/spec/extraStorageClassParameters", "value": {"migratable":"true","numberOfReplicas":"1","staleReplicaTimeout":"30"}}`,
			},
		},
		{
			name:      "[create] add new params",
			oldParams: nil,
			newParams: map[string]string{
				"diskSelector": "sata",
			},
			patchOps: []string{
				`{"op": "replace", "path": "/spec/extraStorageClassParameters", "value": {"diskSelector":"sata","migratable":"true","numberOfReplicas":"3","staleReplicaTimeout":"30"}}`,
			},
		},
		{
			name:      "[create] overwrite default setting and add new params",
			oldParams: nil,
			newParams: map[string]string{
				"diskSelector":     "nvme",
				"numberOfReplicas": "1",
			},
			patchOps: []string{
				`{"op": "replace", "path": "/spec/extraStorageClassParameters", "value": {"diskSelector":"nvme","migratable":"true","numberOfReplicas":"1","staleReplicaTimeout":"30"}}`,
			},
		},
		{
			name:      "[update] old params is empty",
			oldParams: map[string]string{},
			newParams: map[string]string{
				"diskSelector":     "nvme",
				"numberOfReplicas": "1",
			},
			patchOps: []string{
				`{"op": "replace", "path": "/spec/extraStorageClassParameters", "value": {"diskSelector":"nvme","migratable":"true","numberOfReplicas":"1","staleReplicaTimeout":"30"}}`,
			},
		},
		{
			name: "[update] new params is nil",
			oldParams: map[string]string{
				"diskSelector":        "nvme",
				"migratable":          "true",
				"numberOfReplicas":    "1",
				"staleReplicaTimeout": "30",
			},
			newParams: nil,
			patchOps: []string{
				`{"op": "add", "path": "/spec/extraStorageClassParameters", "value": {"migratable":"true","numberOfReplicas":"3","staleReplicaTimeout":"30"}}`,
			},
		},
		{
			name: "[update] new params is empty",
			oldParams: map[string]string{
				"diskSelector":        "nvme",
				"migratable":          "true",
				"numberOfReplicas":    "1",
				"staleReplicaTimeout": "30",
			},
			newParams: map[string]string{},
			patchOps: []string{
				`{"op": "replace", "path": "/spec/extraStorageClassParameters", "value": {"migratable":"true","numberOfReplicas":"3","staleReplicaTimeout":"30"}}`,
			},
		},
		{
			name: "[update] only params changed",
			oldParams: map[string]string{
				"diskSelector":        "sata",
				"migratable":          "true",
				"numberOfReplicas":    "3",
				"staleReplicaTimeout": "30",
			},
			newParams: map[string]string{
				"diskSelector":     "nvme",
				"numberOfReplicas": "1",
			},
			patchOps: []string{
				`{"op": "replace", "path": "/spec/extraStorageClassParameters", "value": {"diskSelector":"nvme","migratable":"true","numberOfReplicas":"1","staleReplicaTimeout":"30"}}`,
			},
		},
		{
			name: "[update] both params and setting changed",
			oldParams: map[string]string{
				"diskSelector":        "sata",
				"migratable":          "true",
				"numberOfReplicas":    "3",
				"staleReplicaTimeout": "30",
			},
			newParams: map[string]string{
				"diskSelector":     "nvme",
				"numberOfReplicas": "1",
			},
			setting: `{"numberOfReplicas":"2","staleReplicaTimeout":"30","migratable":"true"}`,
			patchOps: []string{
				`{"op": "replace", "path": "/spec/extraStorageClassParameters", "value": {"diskSelector":"nvme","migratable":"true","numberOfReplicas":"1","staleReplicaTimeout":"30"}}`,
			},
		},
		{
			name: "[update] only setting changed",
			oldParams: map[string]string{
				"diskSelector":        "sata",
				"migratable":          "true",
				"numberOfReplicas":    "3",
				"staleReplicaTimeout": "30",
			},
			newParams: map[string]string{
				"diskSelector":     "sata",
				"numberOfReplicas": "3",
			},
			setting:  `{"numberOfReplicas":"1","staleReplicaTimeout":"30","migratable":"true"}`,
			patchOps: nil,
		},
		{
			name: "[update] nothing changed",
			oldParams: map[string]string{
				"diskSelector":        "nvme",
				"migratable":          "true",
				"numberOfReplicas":    "1",
				"staleReplicaTimeout": "30",
			},
			newParams: map[string]string{
				"diskSelector":        "nvme",
				"migratable":          "true",
				"numberOfReplicas":    "1",
				"staleReplicaTimeout": "30",
			},
			patchOps: nil,
		},
	}

	setting := &harvesterv1.Setting{
		ObjectMeta: metav1.ObjectMeta{
			Name: settings.ImageDefaultStorageClassParametersSettingName,
		},
		Default: `{"numberOfReplicas":"3","staleReplicaTimeout":"30","migratable":"true"}`,
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			clientset := fake.NewSimpleClientset()
			fakeSetting := setting.DeepCopy()
			if tc.setting != "" {
				fakeSetting.Value = tc.setting
			}
			err := clientset.Tracker().Add(fakeSetting)
			assert.Nil(t, err, "Mock resource should add into fake controller tracker")
			mutator := NewMutator(fakeclients.HarvesterSettingCache(clientset.HarvesterhciV1beta1().Settings)).(*virtualMachineImageMutator)

			oldImage := &harvesterv1.VirtualMachineImage{
				Spec: harvesterv1.VirtualMachineImageSpec{
					ExtraStorageClassParameters: tc.oldParams,
				},
			}
			if tc.oldParams == nil {
				oldImage = nil
			}
			newImage := &harvesterv1.VirtualMachineImage{
				Spec: harvesterv1.VirtualMachineImageSpec{
					ExtraStorageClassParameters: tc.newParams,
				},
			}

			actual, err := mutator.patchImageStorageClassParams(oldImage, newImage)

			assert.Nil(t, err, tc.name)
			assert.Equal(t, tc.patchOps, actual)
		})
	}
}
