// Copyright 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package accessmethods

import (
	"encoding/json"

	"github.com/gardener/ocm/pkg/errors"
	"github.com/gardener/ocm/pkg/ocm/cpi"
	"github.com/gardener/ocm/pkg/runtime"
)

// LocalBlobType is the access type of a blob local to a component.
const LocalBlobType = "localBlob"
const LocalBlobTypeV1 = LocalBlobType + runtime.VersionSeparator + "v1"

func init() {
	cpi.RegisterAccessType(cpi.NewConvertedAccessSpecType(LocalBlobType, LocalBlobV1))
	cpi.RegisterAccessType(cpi.NewConvertedAccessSpecType(LocalBlobTypeV1, LocalBlobV1))
}

// NewLocalBlobAccessSpecV1 creates a new localFilesystemBlob accessor.
func NewLocalBlobAccessSpecV1(path string, mediaType string) *LocalBlobAccessSpec {
	return &LocalBlobAccessSpec{
		ObjectVersionedType: runtime.NewVersionedObjectType(LocalBlobType),
		Filename:            path,
		MediaType:           mediaType,
	}
}

// LocalBlobAccessSpec describes the access for a blob on the filesystem.
type LocalBlobAccessSpec struct {
	runtime.ObjectVersionedType
	// Filename is the name of the blob in the local filesystem.
	// The blob is expected to be at <fs-root>/blobs/<name>
	Filename string
	// MediaType is the media type of the object this filename refers to.
	MediaType string
}

var _ json.Marshaler = &LocalBlobAccessSpec{}

func (s *LocalBlobAccessSpec) MarshalJSON() ([]byte, error) {
	return cpi.MarshalConvertedAccessSpec(s)
}

func (a *LocalBlobAccessSpec) ValidFor(repo cpi.Repository) bool {
	return repo.LocalSupportForAccessSpec(a)
}

func (a *LocalBlobAccessSpec) AccessMethod(c cpi.ComponentAccess) (cpi.AccessMethod, error) {
	if a.ValidFor(c.GetRepository()) {
		return c.AccessMethod(a)
	}
	return nil, errors.ErrNotImplemented(errors.KIND_ACCESSMETHOD, LocalBlobType, c.GetRepository().GetSpecification().GetKind())
}

////////////////////////////////////////////////////////////////////////////////

type LocalBlobAccessSpecV1 struct {
	runtime.ObjectVersionedType `json:",inline"`
	// Filename is the name of the blob in the local filesystem.
	// The blob is expected to be at <fs-root>/blobs/<name>
	Filename string `json:"filename"`
	// MediaType is the media type of the object this filename refers to.
	MediaType string `json:"mediaType,omitempty"`
}

type localblobConverterV1 struct{}

var LocalBlobV1 = cpi.NewAccessSpecVersion(&LocalBlobAccessSpecV1{}, localblobConverterV1{})

func (_ localblobConverterV1) ConvertFrom(object cpi.AccessSpec) (runtime.TypedObject, error) {
	in := object.(*LocalBlobAccessSpec)
	return &LocalBlobAccessSpecV1{
		ObjectVersionedType: runtime.NewVersionedObjectType(in.Type),
		Filename:            in.Filename,
		MediaType:           in.MediaType,
	}, nil
}

func (_ localblobConverterV1) ConvertTo(object interface{}) (cpi.AccessSpec, error) {
	in := object.(*LocalBlobAccessSpecV1)
	return &LocalBlobAccessSpec{
		ObjectVersionedType: runtime.NewVersionedObjectType(in.Type),
		Filename:            in.Filename,
		MediaType:           in.MediaType,
	}, nil
}
