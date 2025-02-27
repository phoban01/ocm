// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

package dockerconfig_test

import (
	"encoding/json"
	"reflect"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/open-component-model/ocm/pkg/common"
	"github.com/open-component-model/ocm/pkg/contexts/credentials"
	"github.com/open-component-model/ocm/pkg/contexts/credentials/cpi"
	local "github.com/open-component-model/ocm/pkg/contexts/credentials/repositories/dockerconfig"
	"github.com/open-component-model/ocm/pkg/contexts/oci/identity"
)

var _ = Describe("docker config", func() {

	props := common.Properties{
		"username":      "mandelsoft",
		"password":      "password",
		"serverAddress": "https://index.docker.io/v1/",
	}

	props2 := common.Properties{
		"username":      "mandelsoft",
		"password":      "token",
		"serverAddress": "https://ghcr.io",
	}

	specdata := "{\"type\":\"DockerConfig\",\"dockerConfigFile\":\"testdata/dockerconfig.json\"}"
	specdata2 := "{\"type\":\"DockerConfig\",\"dockerConfigFile\":\"testdata/dockerconfig.json\",\"propagateConsumerIdentity\":true}"

	var DefaultContext credentials.Context

	BeforeEach(func() {
		DefaultContext = credentials.New()
	})

	It("serializes repo spec", func() {
		spec := local.NewRepositorySpec("testdata/dockerconfig.json")
		data, err := json.Marshal(spec)
		Expect(err).To(Succeed())
		Expect(data).To(Equal([]byte(specdata)))

		spec = local.NewRepositorySpec("testdata/dockerconfig.json").WithConsumerPropagation(true)
		data, err = json.Marshal(spec)
		Expect(err).To(Succeed())
		Expect(data).To(Equal([]byte(specdata2)))
	})
	It("deserializes repo spec", func() {
		spec, err := DefaultContext.RepositorySpecForConfig([]byte(specdata), nil)
		Expect(err).To(Succeed())
		Expect(reflect.TypeOf(spec).String()).To(Equal("*dockerconfig.RepositorySpec"))
		Expect(spec.(*local.RepositorySpec).DockerConfigFile).To(Equal("testdata/dockerconfig.json"))
	})

	It("resolves repository", func() {
		repo, err := DefaultContext.RepositoryForConfig([]byte(specdata), nil)
		Expect(err).To(Succeed())
		Expect(reflect.TypeOf(repo).String()).To(Equal("*dockerconfig.Repository"))
	})

	It("retrieves credentials", func() {
		repo, err := DefaultContext.RepositoryForConfig([]byte(specdata), nil)
		Expect(err).To(Succeed())

		creds, err := repo.LookupCredentials("index.docker.io")
		Expect(err).To(Succeed())
		Expect(creds.Properties()).To(Equal(props))

		creds, err = repo.LookupCredentials("ghcr.io")
		Expect(err).To(Succeed())
		Expect(creds.Properties()).To(Equal(props2))
	})

	It("propagates credentials to consumer identity", func() {
		_, err := DefaultContext.RepositoryForConfig([]byte(specdata2), nil)
		Expect(err).To(Succeed())

		creds, err := credentials.CredentialsForConsumer(DefaultContext, credentials.ConsumerIdentity{
			cpi.ATTR_TYPE:        identity.CONSUMER_TYPE,
			identity.ID_HOSTNAME: "ghcr.io",
		})
		Expect(err).To(Succeed())
		Expect(creds.Properties()).To(Equal(props2))
	})

})
