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

package artefactset

import (
	"github.com/gardener/ocm/pkg/common/accessio"
	"github.com/gardener/ocm/pkg/common/accessobj"
	"github.com/gardener/ocm/pkg/errors"
	"github.com/gardener/ocm/pkg/oci/artdesc"
	"github.com/gardener/ocm/pkg/oci/core"
	"github.com/gardener/ocm/pkg/oci/cpi"
	"github.com/opencontainers/go-digest"
)

var ErrNoIndex = errors.New("manifest does not support access to subsequent artefacts")

type Artefact struct {
	artefactBase
}

var _ cpi.ArtefactAccess = (*Artefact)(nil)

func NewArtefactForBlob(access ArtefactSetContainer, blob accessio.BlobAccess) (*Artefact, error) {
	mode := accessobj.ACC_WRITABLE
	if access.IsReadOnly() {
		mode = accessobj.ACC_READONLY
	}
	state, err := accessobj.NewBlobStateForBlob(mode, blob, NewArtefactStateHandler())
	if err != nil {
		return nil, err
	}
	a := &Artefact{
		artefactBase: artefactBase{
			access: access,
			state:  state,
		},
	}
	return a, nil
}

func NewArtefact(access ArtefactSetContainer, defs ...*artdesc.Artefact) *Artefact {
	var def *artdesc.Artefact
	if len(defs) != 0 && defs[0] != nil {
		def = defs[0]
	}
	mode := accessobj.ACC_WRITABLE
	if access.IsReadOnly() {
		mode = accessobj.ACC_READONLY
	}
	state, err := accessobj.NewBlobStateForObject(mode, def, NewArtefactStateHandler())
	if err != nil {
		panic("oops: " + err.Error())
	}

	a := &Artefact{
		artefactBase: artefactBase{
			access: access,
			state:  state,
		},
	}
	return a
}

////////////////////////////////////////////////////////////////////////////////
// forward

func (a *Artefact) AddBlob(access cpi.BlobAccess) error {
	return a.addBlob(access)
}

func (a *Artefact) Blob() (accessio.BlobAccess, error) {
	d := a.state.GetState().(*artdesc.Artefact)
	if !d.IsValid() {
		return nil, errors.ErrUnknown("artefact type")
	}
	blob, err := a.blob()
	if err != nil {
		return nil, err
	}
	return accessio.BlobWithMimeType(d.MimeType(), blob), nil
}

func (a *Artefact) NewArtefact(art ...*artdesc.Artefact) (cpi.ArtefactAccess, error) {
	if !a.IsIndex() {
		return nil, ErrNoIndex
	}
	return a.newArtefact(art...)
}

////////////////////////////////////////////////////////////////////////////////

func (a *Artefact) Artefact() *artdesc.Artefact {
	return a.GetDescriptor()
}

func (a *Artefact) GetDescriptor() *artdesc.Artefact {
	d := a.state.GetState().(*artdesc.Artefact)
	if d.IsValid() {
		return d
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// from artdesc.Artefact

func (a *Artefact) GetBlobDescriptor(digest digest.Digest) *cpi.Descriptor {
	d := a.GetDescriptor().GetBlobDescriptor(digest)
	if d != nil {
		return d
	}
	return a.access.GetBlobDescriptor(digest)
}

func (a *Artefact) Index() (*artdesc.Index, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	d := a.state.GetState().(*artdesc.Artefact)
	idx := d.Index()
	if idx == nil {
		idx = artdesc.NewIndex()
		if err := d.SetIndex(idx); err != nil {
			return nil, err
		}
	}
	return idx, nil
}

func (a *Artefact) Manifest() (*artdesc.Manifest, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	d := a.state.GetState().(*artdesc.Artefact)
	m := d.Manifest()
	if m == nil {
		m = artdesc.NewManifest()
		if err := d.SetManifest(m); err != nil {
			return nil, err
		}
	}
	return m, nil
}

func (a *Artefact) ManifestAccess() core.ManifestAccess {
	a.lock.Lock()
	defer a.lock.Unlock()
	d := a.state.GetState().(*artdesc.Artefact)
	m := d.Manifest()
	if m == nil {
		m = artdesc.NewManifest()
		if err := d.SetManifest(m); err != nil {
			return nil
		}
	}
	return NewManifestForArtefact(a)
}

func (a *Artefact) IndexAccess() core.IndexAccess {
	a.lock.Lock()
	defer a.lock.Unlock()
	d := a.state.GetState().(*artdesc.Artefact)
	m := d.Manifest()
	if m == nil {
		m = artdesc.NewManifest()
		if err := d.SetManifest(m); err != nil {
			return nil
		}
	}
	return NewIndexForArtefact(a)
}

func (a *Artefact) GetArtefact(digest digest.Digest) (cpi.ArtefactAccess, error) {
	if !a.IsIndex() {
		return nil, ErrNoIndex
	}
	return a.getArtefact(digest)
}

func (a *Artefact) GetManifest(digest digest.Digest) (cpi.ManifestAccess, error) {
	if !a.IsIndex() {
		return nil, ErrNoIndex
	}
	return a.IndexAccess().GetManifest(digest)
}

func (a *Artefact) GetIndex(digest digest.Digest) (cpi.IndexAccess, error) {
	if !a.IsIndex() {
		return nil, ErrNoIndex
	}
	return a.IndexAccess().GetIndex(digest)
}

func (a *Artefact) GetBlob(digest digest.Digest) (cpi.BlobAccess, error) {
	d := a.GetBlobDescriptor(digest)
	if d != nil {
		data, err := a.access.GetBlobData(digest)
		if err != nil {
			return nil, err
		}
		return accessio.BlobAccessForDataAccess(d.Digest, d.Size, d.MediaType, data), nil
	}
	return nil, cpi.ErrBlobNotFound(digest)
}

func (a *Artefact) AddArtefact(art cpi.Artefact, platform *artdesc.Platform) (accessio.BlobAccess, error) {
	if a.IsClosed() {
		return nil, accessio.ErrClosed
	}
	if a.IsReadOnly() {
		return nil, accessio.ErrReadOnly
	}
	_, err := a.Index()
	if err != nil {
		return nil, err
	}
	return NewIndexForArtefact(a).AddArtefact(art, platform)
}

func (a *Artefact) AddLayer(blob cpi.BlobAccess, d *cpi.Descriptor) (int, error) {
	if a.IsClosed() {
		return -1, accessio.ErrClosed
	}
	if a.IsReadOnly() {
		return -1, accessio.ErrReadOnly
	}
	_, err := a.Manifest()
	if err != nil {
		return -1, err
	}
	return NewManifestForArtefact(a).AddLayer(blob, d)
}
