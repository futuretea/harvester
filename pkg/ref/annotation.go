package ref

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/gengo/examples/set-gen/sets"
)

const (
	AnnotationSchemaOwnerKeyName = "harvesterhci.io/owned-by"
)

type AnnotationSchemaReferences struct {
	sets.String
}

func (s AnnotationSchemaReferences) MarshalJSON() ([]byte, error) {
	__traceStack()

	return json.Marshal(s.List())
}

func (s *AnnotationSchemaReferences) UnmarshalJSON(bytes []byte) error {
	__traceStack()

	var arr []string
	if err := json.Unmarshal(bytes, &arr); err != nil {
		return err
	}
	s.String = sets.NewString(arr...)
	return nil
}

func NewAnnotationSchemaOwnerReferences(refs ...string) AnnotationSchemaReferences {
	__traceStack()

	return AnnotationSchemaReferences{String: sets.NewString(refs...)}
}

type AnnotationSchemaReference struct {
	SchemaID	string	`json:"schema"`

	References	AnnotationSchemaReferences	`json:"refs,omitempty"`
}

type AnnotationSchemaOwners map[string]AnnotationSchemaReference

func (o AnnotationSchemaOwners) String() string {
	__traceStack()

	var bytes, _ = o.MarshalJSON()
	return string(bytes)
}

func (o AnnotationSchemaOwners) MarshalJSON() ([]byte, error) {
	__traceStack()

	if o == nil {
		return []byte(`[]`), nil
	}
	var refs = make(sortableSliceOfAnnotationSchemaReference, 0, len(o))
	for _, ref := range o {
		refs = append(refs, ref)
	}
	if len(refs) > 1 {
		sort.Sort(refs)
	}
	return json.Marshal([]AnnotationSchemaReference(refs))
}

func (o *AnnotationSchemaOwners) UnmarshalJSON(bytes []byte) error {
	__traceStack()

	var refs []AnnotationSchemaReference
	if err := json.Unmarshal(bytes, &refs); err != nil {
		return err
	}

	var owners = make(AnnotationSchemaOwners, len(refs))
	for _, ref := range refs {
		if ref.SchemaID == "" {
			continue
		}
		if _, existed := owners[ref.SchemaID]; !existed {

			owners[ref.SchemaID] = ref
			continue
		}

		owners[ref.SchemaID].References.Insert(ref.References.List()...)
	}
	*o = owners
	return nil
}

func (o AnnotationSchemaOwners) List(ownerGK schema.GroupKind) []string {
	__traceStack()

	var schemaID = GroupKindToSchemaID(ownerGK)
	var schemaRef, existed = o[schemaID]
	if !existed || schemaRef.SchemaID != schemaID {
		return []string{}
	}
	return schemaRef.References.UnsortedList()
}

func (o AnnotationSchemaOwners) Has(ownerGK schema.GroupKind, owner metav1.Object) bool {
	__traceStack()

	var schemaID = GroupKindToSchemaID(ownerGK)
	var ownerRef = Construct(owner.GetNamespace(), owner.GetName())

	var schemaRef, existed = o[schemaID]
	if !existed {
		return false
	}
	return schemaRef.SchemaID == schemaID && schemaRef.References.Has(ownerRef)
}

func (o AnnotationSchemaOwners) Add(ownerGK schema.GroupKind, owner metav1.Object) bool {
	__traceStack()

	if o.Has(ownerGK, owner) {
		return false
	}

	var schemaID = GroupKindToSchemaID(ownerGK)
	var ownerRef = Construct(owner.GetNamespace(), owner.GetName())
	var schemaRef, existed = o[schemaID]
	if !existed {
		schemaRef = AnnotationSchemaReference{SchemaID: schemaID, References: NewAnnotationSchemaOwnerReferences()}
	}
	schemaRef.References.Insert(ownerRef)
	o[schemaID] = schemaRef
	return true
}

func (o AnnotationSchemaOwners) Remove(ownerGK schema.GroupKind, owner metav1.Object) bool {
	__traceStack()

	if !o.Has(ownerGK, owner) {
		return false
	}

	var schemaID = GroupKindToSchemaID(ownerGK)
	var ownerRef = Construct(owner.GetNamespace(), owner.GetName())
	var schemaRef = o[schemaID]
	if schemaRef.References.Delete(ownerRef).Len() == 0 {
		delete(o, schemaID)
	}
	return true
}

func (o AnnotationSchemaOwners) Bind(obj metav1.Object) error {
	__traceStack()

	var annotations = obj.GetAnnotations()
	if annotations == nil {
		annotations = map[string]string{}
	}

	if len(o) == 0 {
		delete(annotations, AnnotationSchemaOwnerKeyName)
	} else {
		var ownersBytes, err = json.Marshal(o)
		if err != nil {
			return fmt.Errorf("failed to marshal annotation schema owners: %w", err)
		}
		annotations[AnnotationSchemaOwnerKeyName] = string(ownersBytes)
	}

	if len(annotations) == 0 {
		obj.SetAnnotations(nil)
	} else {
		obj.SetAnnotations(annotations)
	}
	return nil
}

type sortableSliceOfAnnotationSchemaReference []AnnotationSchemaReference

func (s sortableSliceOfAnnotationSchemaReference) Len() int {
	__traceStack()

	return len(s)
}

func (s sortableSliceOfAnnotationSchemaReference) Less(i, j int) bool {
	__traceStack()

	return s[i].SchemaID < s[j].SchemaID
}

func (s sortableSliceOfAnnotationSchemaReference) Swap(i, j int) {
	__traceStack()

	s[i], s[j] = s[j], s[i]
}

func GroupKindToSchemaID(kind schema.GroupKind) string {
	__traceStack()

	return strings.ToLower(fmt.Sprintf("%s.%s", kind.Group, kind.Kind))
}

func GetSchemaOwnersFromAnnotation(obj metav1.Object) (AnnotationSchemaOwners, error) {
	__traceStack()

	var annotations = obj.GetAnnotations()
	var ownedByAnnotation, ok = annotations[AnnotationSchemaOwnerKeyName]
	if !ok {
		return AnnotationSchemaOwners{}, nil
	}

	var owner AnnotationSchemaOwners
	if err := json.Unmarshal([]byte(ownedByAnnotation), &owner); err != nil {
		return owner, fmt.Errorf("failed to unmarshal annotation schema owners: %w", err)
	}
	return owner, nil
}
