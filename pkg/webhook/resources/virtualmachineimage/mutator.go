package virtualmachineimage

import (
	"encoding/json"
	"fmt"
	"reflect"

	admissionregv1 "k8s.io/api/admissionregistration/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"

	harvesterv1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	ctlharvesterv1 "github.com/harvester/harvester/pkg/generated/controllers/harvesterhci.io/v1beta1"
	"github.com/harvester/harvester/pkg/settings"
	"github.com/harvester/harvester/pkg/webhook/types"
)

func NewMutator(setting ctlharvesterv1.SettingCache) types.Mutator {
	return &virtualMachineImageMutator{setting: setting}
}

type virtualMachineImageMutator struct {
	types.DefaultMutator
	setting ctlharvesterv1.SettingCache
}

func (m *virtualMachineImageMutator) Resource() types.Resource {
	return types.Resource{
		Names:      []string{harvesterv1.VirtualMachineImageResourceName},
		Scope:      admissionregv1.NamespacedScope,
		APIGroup:   harvesterv1.SchemeGroupVersion.Group,
		APIVersion: harvesterv1.SchemeGroupVersion.Version,
		ObjectType: &harvesterv1.VirtualMachineImage{},
		OperationTypes: []admissionregv1.OperationType{
			admissionregv1.Create,
			admissionregv1.Update,
		},
	}
}

func (m *virtualMachineImageMutator) Create(request *types.Request, newObj runtime.Object) (types.PatchOps, error) {
	newImage := newObj.(*harvesterv1.VirtualMachineImage)

	return m.patchImageStorageClassParams(nil, newImage)
}

func (m *virtualMachineImageMutator) Update(request *types.Request, oldObj runtime.Object, newObj runtime.Object) (types.PatchOps, error) {
	newImage := newObj.(*harvesterv1.VirtualMachineImage)
	oldImage := oldObj.(*harvesterv1.VirtualMachineImage)

	return m.patchImageStorageClassParams(oldImage, newImage)
}

func (m *virtualMachineImageMutator) patchImageStorageClassParams(oldImage *harvesterv1.VirtualMachineImage, newImage *harvesterv1.VirtualMachineImage) ([]string, error) {
	var patchOps types.PatchOps
	newParams, err := m.getImageDefaultStorageClassParameters()
	if err != nil {
		return patchOps, err
	}

	for k, v := range newImage.Spec.ExtraStorageClassParameters {
		newParams[k] = v
	}

	if oldImage != nil {
		oldParams := oldImage.Spec.ExtraStorageClassParameters
		if reflect.DeepEqual(oldParams, newParams) {
			return patchOps, nil
		}
	}

	valueBytes, err := json.Marshal(newParams)
	if err != nil {
		return patchOps, err
	}

	verb := "add"
	if newImage.Spec.ExtraStorageClassParameters != nil {
		verb = "replace"
	}

	patchOps = append(patchOps, fmt.Sprintf(`{"op": "%s", "path": "/spec/extraStorageClassParameters", "value": %s}`, verb, string(valueBytes)))
	return patchOps, nil
}

func (m *virtualMachineImageMutator) getImageDefaultStorageClassParameters() (map[string]string, error) {
	s, err := m.setting.Get(settings.ImageDefaultStorageClassParametersSettingName)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	value := s.Value
	if value == "" {
		value = s.Default
	}
	if value == "" {
		return nil, nil
	}

	params := map[string]string{}
	if err = json.Unmarshal([]byte(value), &params); err != nil {
		return params, err
	}
	return params, nil
}
