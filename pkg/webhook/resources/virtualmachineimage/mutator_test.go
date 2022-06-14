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
		name     string
		params   map[string]string
		setting  string
		patchOps []string
	}{
		{
			name:   "params is nil",
			params: nil,
			patchOps: []string{
				`{"op": "add", "path": "/spec/extraStorageClassParameters", "value": {"migratable":"true","numberOfReplicas":"3","staleReplicaTimeout":"30"}}`,
			},
		},
		{
			name:   "params is empty",
			params: map[string]string{},
			patchOps: []string{
				`{"op": "replace", "path": "/spec/extraStorageClassParameters", "value": {"migratable":"true","numberOfReplicas":"3","staleReplicaTimeout":"30"}}`,
			},
		},
		{
			name: "only overwrite default setting",
			params: map[string]string{
				"numberOfReplicas": "1",
			},
			patchOps: []string{
				`{"op": "replace", "path": "/spec/extraStorageClassParameters", "value": {"migratable":"true","numberOfReplicas":"1","staleReplicaTimeout":"30"}}`,
			},
		},
		{
			name: "add new params",
			params: map[string]string{
				"diskSelector": "sata",
			},
			patchOps: []string{
				`{"op": "replace", "path": "/spec/extraStorageClassParameters", "value": {"diskSelector":"sata","migratable":"true","numberOfReplicas":"3","staleReplicaTimeout":"30"}}`,
			},
		},
		{
			name: "overwrite default setting and add new params",
			params: map[string]string{
				"diskSelector":     "nvme",
				"numberOfReplicas": "1",
			},
			patchOps: []string{
				`{"op": "replace", "path": "/spec/extraStorageClassParameters", "value": {"diskSelector":"nvme","migratable":"true","numberOfReplicas":"1","staleReplicaTimeout":"30"}}`,
			},
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

			image := &harvesterv1.VirtualMachineImage{
				Spec: harvesterv1.VirtualMachineImageSpec{
					ExtraStorageClassParameters: tc.params,
				},
			}

			actual, err := mutator.patchImageStorageClassParams(image)

			assert.Nil(t, err, tc.name)
			assert.Equal(t, tc.patchOps, actual)
		})
	}
}
