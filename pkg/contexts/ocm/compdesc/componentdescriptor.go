// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

package compdesc

import (
	"fmt"
	"reflect"
	"strings"

	metav1 "github.com/open-component-model/ocm/pkg/contexts/ocm/compdesc/meta/v1"
	"github.com/open-component-model/ocm/pkg/errors"
	"github.com/open-component-model/ocm/pkg/runtime"
)

var NotFound = errors.ErrNotFound()

const KIND_REFERENCE = "component reference"

type (
	ObjectMeta = metav1.ObjectMeta
	Provider   = metav1.Provider
)

const ComponentDescriptorFileName = "component-descriptor.yaml"

// Metadata defines the configured metadata of the component descriptor.
// It is taken from the original serialization format. It can be set
// to define a default serialization version.
type Metadata struct {
	ConfiguredVersion string `json:"configuredSchemaVersion"`
}

// ComponentDescriptor defines a versioned component with a source and dependencies.
// +k8s:deepcopy-gen=true
// +k8s:openapi-gen=true
type ComponentDescriptor struct {
	// Metadata specifies the schema version of the component.
	Metadata Metadata `json:"meta"`
	// Spec contains the specification of the component.
	ComponentSpec `json:"component"`
	// Signatures contains a list of signatures for the ComponentDescriptor
	Signatures metav1.Signatures `json:"signatures,omitempty"`
}

func New(name, version string) *ComponentDescriptor {
	return DefaultComponent(&ComponentDescriptor{
		Metadata: Metadata{
			ConfiguredVersion: "v2",
		},
		ComponentSpec: ComponentSpec{
			ObjectMeta: metav1.ObjectMeta{
				Name:    name,
				Version: version,
				Provider: metav1.Provider{
					Name: "acme",
				},
			},
		},
	})
}

// SchemaVersion returns the scheme version configured in the representation.
func (cd *ComponentDescriptor) SchemaVersion() string {
	return cd.Metadata.ConfiguredVersion
}

func (cd *ComponentDescriptor) Copy() *ComponentDescriptor {
	out := &ComponentDescriptor{
		Metadata: cd.Metadata,
		ComponentSpec: ComponentSpec{
			ObjectMeta:         *cd.ObjectMeta.Copy(),
			RepositoryContexts: cd.RepositoryContexts.Copy(),
			Sources:            cd.Sources.Copy(),
			References:         cd.References.Copy(),
			Resources:          cd.Resources.Copy(),
		},
		Signatures: cd.Signatures.Copy(),
	}
	return out
}

func (cd *ComponentDescriptor) Reset() {
	cd.Provider.Name = ""
	cd.Provider.Labels = nil
	cd.Resources = nil
	cd.Sources = nil
	cd.References = nil
	cd.RepositoryContexts = nil
	cd.Signatures = nil
	cd.Labels = nil
	DefaultComponent(cd)
}

// ComponentSpec defines a virtual component with
// a repository context, source and dependencies.
// +k8s:deepcopy-gen=true
// +k8s:openapi-gen=true
type ComponentSpec struct {
	metav1.ObjectMeta `json:",inline"`
	// RepositoryContexts defines the previous repositories of the component
	RepositoryContexts runtime.UnstructuredTypedObjectList `json:"repositoryContexts"`
	// Sources defines sources that produced the component
	Sources Sources `json:"sources"`
	// References references component dependencies that can be resolved in the current context.
	References References `json:"componentReferences"`
	// Resources defines all resources that are created by the component and by a third party.
	Resources Resources `json:"resources"`
}

const (
	SystemIdentityName    = "name"
	SystemIdentityVersion = "version"
)

// ElementMeta defines a object that is uniquely identified by its identity.
// +k8s:deepcopy-gen=true
// +k8s:openapi-gen=true
type ElementMeta struct {
	// Name is the context unique name of the object.
	Name string `json:"name"`
	// Version is the semver version of the object.
	Version string `json:"version"`
	// ExtraIdentity is the identity of an object.
	// An additional label with key "name" ist not allowed
	ExtraIdentity metav1.Identity `json:"extraIdentity,omitempty"`
	// Labels defines an optional set of additional labels
	// describing the object.
	// +optional
	Labels metav1.Labels `json:"labels,omitempty"`
}

// GetName returns the name of the object.
func (o *ElementMeta) GetName() string {
	return o.Name
}

// SetName sets the name of the object.
func (o *ElementMeta) SetName(name string) {
	o.Name = name
}

// GetVersion returns the version of the object.
func (o *ElementMeta) GetVersion() string {
	return o.Version
}

// SetVersion sets the version of the object.
func (o *ElementMeta) SetVersion(version string) {
	o.Version = version
}

// GetLabels returns the label of the object.
func (o *ElementMeta) GetLabels() metav1.Labels {
	return o.Labels
}

// SetLabels sets the labels of the object.
func (o *ElementMeta) SetLabels(labels []metav1.Label) {
	o.Labels = labels
}

// SetLabel sets a single label to an effective value.
// If the value is no byte slice, it is marshaled.
func (o *ElementMeta) SetLabel(name string, value interface{}) error {
	return o.Labels.Set(name, value)
}

// RemoveLabel removes a single label.
func (o *ElementMeta) RemoveLabel(name string) bool {
	return o.Labels.Remove(name)
}

// SetExtraIdentity sets the identity of the object.
func (o *ElementMeta) SetExtraIdentity(identity metav1.Identity) {
	o.ExtraIdentity = identity
}

// GetIdentity returns the identity of the object.
func (o *ElementMeta) GetIdentity(accessor ElementAccessor) metav1.Identity {
	identity := o.ExtraIdentity.Copy()
	if identity == nil {
		identity = metav1.Identity{}
	}
	identity[SystemIdentityName] = o.Name
	if accessor != nil {
		found := false
		l := accessor.Len()
		for i := 0; i < l; i++ {
			m := accessor.Get(i).GetMeta()
			if m.Name == o.Name && m.ExtraIdentity.Equals(o.ExtraIdentity) {
				if found {
					identity[SystemIdentityVersion] = o.Version

					break
				}
				found = true
			}
		}
	}
	return identity
}

// GetRawIdentity returns the identity plus version.
func (o *ElementMeta) GetRawIdentity() metav1.Identity {
	identity := o.ExtraIdentity.Copy()
	if identity == nil {
		identity = metav1.Identity{}
	}
	identity[SystemIdentityName] = o.Name
	if o.Version != "" {
		identity[SystemIdentityVersion] = o.Version
	}
	return identity
}

// GetMatchBaseIdentity returns all possible identity attributes for resource matching.
func (o *ElementMeta) GetMatchBaseIdentity() metav1.Identity {
	identity := o.ExtraIdentity.Copy()
	if identity == nil {
		identity = metav1.Identity{}
	}
	identity[SystemIdentityName] = o.Name
	identity[SystemIdentityVersion] = o.Version

	return identity
}

// GetIdentityDigest returns the digest of the object's identity.
func (o *ElementMeta) GetIdentityDigest(accessor ElementAccessor) []byte {
	return o.GetIdentity(accessor).Digest()
}

func (o *ElementMeta) Copy() *ElementMeta {
	if o == nil {
		return nil
	}
	return &ElementMeta{
		Name:          o.Name,
		Version:       o.Version,
		ExtraIdentity: o.ExtraIdentity.Copy(),
		Labels:        o.Labels.Copy(),
	}
}

// NameAccessor describes a accessor for a named object.
type NameAccessor interface {
	// GetName returns the name of the object.
	GetName() string
	// SetName sets the name of the object.
	SetName(name string)
}

// VersionAccessor describes a accessor for a versioned object.
type VersionAccessor interface {
	// GetVersion returns the version of the object.
	GetVersion() string
	// SetVersion sets the version of the object.
	SetVersion(version string)
}

// LabelsAccessor describes a accessor for a labeled object.
type LabelsAccessor interface {
	// GetLabels returns the labels of the object.
	GetLabels() metav1.Labels
	// SetLabels sets the labels of the object.
	SetLabels(labels []metav1.Label)
}

// ObjectMetaAccessor describes a accessor for named and versioned object.
type ObjectMetaAccessor interface {
	NameAccessor
	VersionAccessor
	LabelsAccessor
}

// ElementMetaAccessor provides generic access an elements meta information.
type ElementMetaAccessor interface {
	GetMeta() *ElementMeta
}

// ElementAccessor provides generic access to list of elements.
type ElementAccessor interface {
	Len() int
	Get(i int) ElementMetaAccessor
}

// ElementArtifactAccessor provides access to generic artifact information of an element.
type ElementArtifactAccessor interface {
	ElementMetaAccessor
	GetAccess() AccessSpec
}

// ArtifactAccessor provides generic access to list of artifacts.
// There are resources or sources.
type ArtifactAccessor interface {
	Len() int
	GetArtifact(i int) ElementArtifactAccessor
}

// ArtifactAccess provides access to a dedicated kind of artifact set
// in the component descriptor (resources or sources).
type ArtifactAccess func(cd *ComponentDescriptor) ArtifactAccessor

// AccessSpec is an abstract specification of an access method
// The outbound object is typicall a runtime.UnstructuredTypedObject.
// Inbound any serializable AccessSpec implementation is possible.
type AccessSpec interface {
	runtime.VersionedTypedObject
}

// GenericAccessSpec returns a generic AccessSpec implementation for an unstructured object.
// It can always be used instead of a dedicated access spec implementation. The core
// methods will map these spec into effective ones before an access is returned to the caller.
func GenericAccessSpec(un *runtime.UnstructuredTypedObject) AccessSpec {
	return &runtime.UnstructuredVersionedTypedObject{
		*un.DeepCopy(),
	}
}

// Sources describes a set of source specifications.
type Sources []Source

var _ ElementAccessor = Sources{}

func SourceArtifacts(cd *ComponentDescriptor) ArtifactAccessor {
	return cd.Sources
}

func (s Sources) Len() int {
	return len(s)
}

func (s Sources) Get(i int) ElementMetaAccessor {
	return &s[i]
}

func (s Sources) GetArtifact(i int) ElementArtifactAccessor {
	return &s[i]
}

func (s Sources) Copy() Sources {
	if s == nil {
		return nil
	}
	out := make(Sources, len(s))
	for i, v := range s {
		out[i] = *v.Copy()
	}
	return out
}

// Source is the definition of a component's source.
// +k8s:deepcopy-gen=true
// +k8s:openapi-gen=true
type Source struct {
	SourceMeta `json:",inline"`
	Access     AccessSpec `json:"access"`
}

func (s *Source) GetMeta() *ElementMeta {
	return &s.ElementMeta
}

func (s *Source) GetAccess() AccessSpec {
	return s.Access
}

func (s *Source) Copy() *Source {
	return &Source{
		SourceMeta: *s.SourceMeta.Copy(),
		Access:     s.Access,
	}
}

// SourceMeta is the definition of the meta data of a source.
// +k8s:deepcopy-gen=true
// +k8s:openapi-gen=true
type SourceMeta struct {
	ElementMeta
	// Type describes the type of the object.
	Type string `json:"type"`
}

// GetType returns the type of the object.
func (o *SourceMeta) GetType() string {
	return o.Type
}

// SetType sets the type of the object.
func (o *SourceMeta) SetType(ttype string) {
	o.Type = ttype
}

// Copy copies a source meta.
func (o *SourceMeta) Copy() *SourceMeta {
	if o == nil {
		return nil
	}
	return &SourceMeta{
		ElementMeta: *o.ElementMeta.Copy(),
		Type:        o.Type,
	}
}

// SourceRef defines a reference to a source
// +k8s:deepcopy-gen=true
// +k8s:openapi-gen=true
type SourceRef struct {
	// IdentitySelector defines the identity that is used to match a source.
	IdentitySelector metav1.StringMap `json:"identitySelector,omitempty"`
	// Labels defines an optional set of additional labels
	// describing the object.
	// +optional
	Labels metav1.Labels `json:"labels,omitempty"`
}

// Copy copy a source ref.
func (r *SourceRef) Copy() *SourceRef {
	if r == nil {
		return nil
	}
	return &SourceRef{
		IdentitySelector: r.IdentitySelector.Copy(),
		Labels:           r.Labels.Copy(),
	}
}

type SourceRefs []SourceRef

// Copy copies a list of source refs.
func (r SourceRefs) Copy() SourceRefs {
	if r == nil {
		return nil
	}

	result := make(SourceRefs, len(r))
	for i, v := range r {
		result[i] = *v.Copy()
	}
	return result
}

// Resources describes a set of resource specifications.
type Resources []Resource

var _ ElementAccessor = Resources{}

func ResourceArtifacts(cd *ComponentDescriptor) ArtifactAccessor {
	return cd.Resources
}

func (r Resources) Len() int {
	return len(r)
}

func (r Resources) Get(i int) ElementMetaAccessor {
	return &r[i]
}

func (r Resources) GetArtifact(i int) ElementArtifactAccessor {
	return &r[i]
}

func (r Resources) Copy() Resources {
	if r == nil {
		return nil
	}
	out := make(Resources, len(r))
	for i, v := range r {
		out[i] = *v.Copy()
	}
	return out
}

// Resource describes a resource dependency of a component.
// +k8s:deepcopy-gen=true
// +k8s:openapi-gen=true
type Resource struct {
	ResourceMeta `json:",inline"`
	// Access describes the type specific method to
	// access the defined resource.
	Access AccessSpec `json:"access"`
}

func (r *Resource) GetMeta() *ElementMeta {
	return &r.ElementMeta
}

func (r *Resource) GetAccess() AccessSpec {
	return r.Access
}

func (r *Resource) Copy() *Resource {
	return &Resource{
		ResourceMeta: *r.ResourceMeta.Copy(),
		Access:       r.Access,
	}
}

// ResourceMeta describes the meta data of a resource.
// +k8s:deepcopy-gen=true
// +k8s:openapi-gen=true
type ResourceMeta struct {
	ElementMeta `json:",inline"`

	// Type describes the type of the object.
	Type string `json:"type"`

	// Relation describes the relation of the resource to the component.
	// Can be a local or external resource
	Relation metav1.ResourceRelation `json:"relation,omitempty"`

	// SourceRef defines a list of source names.
	// These names reference the sources defines in `component.sources`.
	SourceRef SourceRefs `json:"srcRef,omitempty"`

	// Digest is the optional digest of the referenced resource.
	// +optional
	Digest *metav1.DigestSpec `json:"digest,omitempty"`
}

// HashEqual indicates whether the digest hash would be equal.
// Excluded: Labels
// Adapt together with version specific hash excludes.
func (o *ResourceMeta) HashEqual(r *ResourceMeta) bool {
	if o.Type != r.Type {
		return false
	}
	if o.Relation != r.Relation {
		return false
	}
	if o.ElementMeta.Name != r.ElementMeta.Name {
		return false
	}
	if o.ElementMeta.Version != r.ElementMeta.Version {
		return false
	}
	if o.Digest != nil && !reflect.DeepEqual(o.Digest, r.Digest) {
		return false
	}
	if !reflect.DeepEqual(o.ElementMeta.ExtraIdentity, r.ElementMeta.ExtraIdentity) {
		return false
	}
	return true
}

// GetType returns the type of the object.
func (o *ResourceMeta) GetType() string {
	return o.Type
}

// SetType sets the type of the object.
func (o *ResourceMeta) SetType(ttype string) {
	o.Type = ttype
}

// Copy copies a resource meta.
func (o *ResourceMeta) Copy() *ResourceMeta {
	if o == nil {
		return nil
	}
	r := &ResourceMeta{
		ElementMeta: *o.ElementMeta.Copy(),
		Type:        o.Type,
		Relation:    o.Relation,
		SourceRef:   o.SourceRef.Copy(),
		Digest:      o.Digest.Copy(),
	}
	return r
}

type References []ComponentReference

func (r References) Len() int {
	return len(r)
}

func (r References) Get(i int) ElementMetaAccessor {
	return &r[i]
}

func (r References) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r References) Less(i, j int) bool {
	c := strings.Compare(r[i].Name, r[j].Name)
	if c != 0 {
		return c < 0
	}
	return strings.Compare(r[i].Version, r[j].Version) < 0
}

func (r References) Copy() References {
	if r == nil {
		return nil
	}
	out := make(References, len(r))
	for i, v := range r {
		out[i] = *v.Copy()
	}
	return out
}

// ComponentReference describes the reference to another component in the registry.
// +k8s:deepcopy-gen=true
// +k8s:openapi-gen=true
type ComponentReference struct {
	ElementMeta `json:",inline"`
	// ComponentName describes the remote name of the referenced object
	ComponentName string `json:"componentName"`
	// Digest is the optional digest of the referenced component.
	// +optional
	Digest *metav1.DigestSpec `json:"digest,omitempty"`
}

func NewComponentReference(name, componentName, version string, extraIdentity metav1.Identity) *ComponentReference {
	return &ComponentReference{
		ElementMeta: ElementMeta{
			Name:          name,
			Version:       version,
			ExtraIdentity: extraIdentity,
		},
		ComponentName: componentName,
	}
}

func (r ComponentReference) String() string {
	return fmt.Sprintf("%s[%s:%s]", r.Name, r.ComponentName, r.Version)
}

func (r *ComponentReference) GetMeta() *ElementMeta {
	return &r.ElementMeta
}

func (r *ComponentReference) GetComponentName() string {
	return r.ComponentName
}

func (r *ComponentReference) Copy() *ComponentReference {
	return &ComponentReference{
		ElementMeta:   *r.ElementMeta.Copy(),
		ComponentName: r.ComponentName,
		Digest:        r.Digest.Copy(),
	}
}
