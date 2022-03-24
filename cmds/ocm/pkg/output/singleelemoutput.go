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

package output

import (
	"github.com/gardener/ocm/pkg/errors"
)

type SingleElementOutput struct {
	Elem interface{}
}

var _ Output = &SingleElementOutput{}

func NewSingleElementOutput() *SingleElementOutput {
	return &SingleElementOutput{}
}

func (this *SingleElementOutput) Add(e interface{}) error {
	if this.Elem == nil {
		this.Elem = e
		return nil
	}
	return errors.Newf("only one element can be selected, but multiple elements selected/found")
}

func (this *SingleElementOutput) Close() error {
	return nil
}

func (this *SingleElementOutput) Out() error {
	return nil
}
