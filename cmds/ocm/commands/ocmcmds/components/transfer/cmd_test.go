// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

package transfer_test

import (
	"bytes"
	"encoding/json"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/open-component-model/ocm/cmds/ocm/testhelper"
	. "github.com/open-component-model/ocm/pkg/contexts/oci/testhelper"
	. "github.com/open-component-model/ocm/pkg/testutils"

	"github.com/open-component-model/ocm/pkg/common/accessio"
	"github.com/open-component-model/ocm/pkg/common/accessobj"
	"github.com/open-component-model/ocm/pkg/contexts/oci"
	"github.com/open-component-model/ocm/pkg/contexts/oci/artdesc"
	"github.com/open-component-model/ocm/pkg/contexts/oci/repositories/artifactset"
	"github.com/open-component-model/ocm/pkg/contexts/ocm"
	"github.com/open-component-model/ocm/pkg/contexts/ocm/accessmethods/ociartifact"
	metav1 "github.com/open-component-model/ocm/pkg/contexts/ocm/compdesc/meta/v1"
	ctfocm "github.com/open-component-model/ocm/pkg/contexts/ocm/repositories/ctf"
	"github.com/open-component-model/ocm/pkg/contexts/ocm/resourcetypes"
	"github.com/open-component-model/ocm/pkg/mime"
)

const ARCH = "/tmp/ctf"
const ARCH2 = "/tmp/ctf2"
const PROVIDER = "mandelsoft"
const VERSION = "v1"
const COMPONENT = "github.com/mandelsoft/test"
const COMPONENT2 = "github.com/mandelsoft/test2"
const OUT = "/tmp/res"
const OCIPATH = "/tmp/oci"
const OCIHOST = "alias"

func Check(env *TestEnv, ldesc *artdesc.Descriptor, out string) {
	tgt, err := ctfocm.Open(env.OCMContext(), accessobj.ACC_READONLY, out, 0, accessio.PathFileSystem(env.FileSystem()))
	Expect(err).To(Succeed())
	defer Close(tgt, "ctf")

	list, err := tgt.ComponentLister().GetComponents("", true)
	Expect(err).To(Succeed())
	Expect(list).To(Equal([]string{COMPONENT}))
	CheckComponent(env, ldesc, tgt)
}

func CheckComponent(env *TestEnv, ldesc *artdesc.Descriptor, tgt ocm.Repository) {
	comp, err := tgt.LookupComponentVersion(COMPONENT, VERSION)
	Expect(err).To(Succeed())
	defer Close(comp, "comvers")
	Expect(len(comp.GetDescriptor().Resources)).To(Equal(3))

	data, err := json.Marshal(comp.GetDescriptor().Resources[2].Access)
	Expect(err).To(Succeed())
	hash := HashManifest2(artifactset.DefaultArtifactSetDescriptorFileName)
	Expect(string(data)).To(StringEqualWithContext("{\"localReference\":\"" + hash + "\",\"mediaType\":\"application/vnd.oci.image.manifest.v1+tar+gzip\",\"referenceName\":\"ocm/ref:v2.0\",\"type\":\"localBlob\"}"))

	data, err = json.Marshal(comp.GetDescriptor().Resources[1].Access)
	Expect(err).To(Succeed())
	hash = HashManifest1(artifactset.DefaultArtifactSetDescriptorFileName)
	Expect(string(data)).To(StringEqualWithContext("{\"localReference\":\"" + hash + "\",\"mediaType\":\"application/vnd.oci.image.manifest.v1+tar+gzip\",\"referenceName\":\"ocm/value:v2.0\",\"type\":\"localBlob\"}"))

	racc, err := comp.GetResourceByIndex(1)
	Expect(err).To(Succeed())
	reader, err := ocm.ResourceReader(racc)
	Expect(err).To(Succeed())
	defer reader.Close()
	set, err := artifactset.Open(accessobj.ACC_READONLY, "", 0, accessio.Reader(reader))
	Expect(err).To(Succeed())
	defer set.Close()

	blob, err := set.GetBlob(ldesc.Digest)
	Expect(err).To(Succeed())
	defer Close(blob, "blob")
	data, err = blob.Get()
	Expect(err).To(Succeed())
	Expect(string(data)).To(Equal("manifestlayer"))
}

var _ = Describe("Test Environment", func() {
	var (
		env   *TestEnv
		ldesc *artdesc.Descriptor
	)

	_ = ldesc
	BeforeEach(func() {
		env = NewTestEnv()

		FakeOCIRepo(env.Builder, OCIPATH, OCIHOST)

		env.OCICommonTransport(OCIPATH, accessio.FormatDirectory, func() {
			ldesc = OCIManifest1(env.Builder)
			OCIManifest2(env.Builder)
		})

		env.OCMCommonTransport(ARCH, accessio.FormatDirectory, func() {
			env.Component(COMPONENT, func() {
				env.Version(VERSION, func() {
					env.Provider(PROVIDER)
					env.Resource("testdata", "", "PlainText", metav1.LocalRelation, func() {
						env.BlobStringData(mime.MIME_TEXT, "testdata")
					})
					env.Resource("value", "", resourcetypes.OCI_IMAGE, metav1.LocalRelation, func() {
						env.Access(
							ociartifact.New(oci.StandardOCIRef(OCIHOST+".alias", OCINAMESPACE, OCIVERSION)),
						)
						env.Label("transportByValue", true)
					})
					env.Resource("ref", "", resourcetypes.OCI_IMAGE, metav1.LocalRelation, func() {
						env.Access(
							ociartifact.New(oci.StandardOCIRef(OCIHOST+".alias", OCINAMESPACE2, OCIVERSION)),
						)
					})
				})
			})
		})

		env.OCMCommonTransport(ARCH2, accessio.FormatDirectory, func() {
			env.Component(COMPONENT2, func() {
				env.Version(VERSION, func() {
					env.Reference("ref", COMPONENT, VERSION)
					env.Provider(PROVIDER)
				})
			})
		})
	})

	AfterEach(func() {
		env.Cleanup()
	})

	It("transfers ctf", func() {
		buf := bytes.NewBuffer(nil)
		Expect(env.CatchOutput(buf).Execute("transfer", "components", "--copy-resources", ARCH, ARCH, OUT)).To(Succeed())
		Expect(buf.String()).To(StringEqualTrimmedWithContext(`
transferring version "github.com/mandelsoft/test:v1"...
...resource 0...
...resource 1(ocm/value:v2.0)...
...resource 2(ocm/ref:v2.0)...
...adding component version...
1 versions transferred
`))

		Expect(env.DirExists(OUT)).To(BeTrue())
		Check(env, ldesc, OUT)
	})

	It("transfers ctf with --closure --lookup", func() {
		buf := bytes.NewBuffer(nil)
		Expect(env.CatchOutput(buf).Execute("transfer", "components", "--copy-resources", "--recursive", "--lookup", ARCH, ARCH2, ARCH2, OUT)).To(Succeed())
		Expect(buf.String()).To(StringEqualTrimmedWithContext(`
transferring version "github.com/mandelsoft/test2:v1"...
  transferring version "github.com/mandelsoft/test:v1"...
  ...resource 0...
  ...resource 1(ocm/value:v2.0)...
  ...resource 2(ocm/ref:v2.0)...
  ...adding component version...
...adding component version...
2 versions transferred
`))

		Expect(env.DirExists(OUT)).To(BeTrue())
		tgt, err := ctfocm.Open(env.OCMContext(), accessobj.ACC_READONLY, OUT, 0, accessio.PathFileSystem(env.FileSystem()))
		Expect(err).To(Succeed())
		defer Close(tgt, "ctf")

		list, err := tgt.ComponentLister().GetComponents("", true)
		Expect(err).To(Succeed())
		Expect(list).To(ContainElements([]string{COMPONENT2, COMPONENT}))

		Expect(tgt.ExistsComponentVersion(COMPONENT2, VERSION)).To(BeTrue())

		CheckComponent(env, ldesc, tgt)
	})

	It("transfers ctf to tgz", func() {
		buf := bytes.NewBuffer(nil)
		Expect(env.CatchOutput(buf).Execute("transfer", "components", "--copy-resources", ARCH, ARCH, accessio.FormatTGZ.String()+"::"+OUT)).To(Succeed())
		Expect(buf.String()).To(StringEqualTrimmedWithContext(`
transferring version "github.com/mandelsoft/test:v1"...
...resource 0...
...resource 1(ocm/value:v2.0)...
...resource 2(ocm/ref:v2.0)...
...adding component version...
1 versions transferred
`))

		Expect(env.FileExists(OUT)).To(BeTrue())
		Check(env, ldesc, OUT)
	})

	It("transfers ctf to ctf+tgz", func() {
		buf := bytes.NewBuffer(nil)
		Expect(env.CatchOutput(buf).Execute("transfer", "components", "--copy-resources", ARCH, ARCH, "ctf+"+accessio.FormatTGZ.String()+"::"+OUT)).To(Succeed())
		Expect(buf.String()).To(StringEqualTrimmedWithContext(`
transferring version "github.com/mandelsoft/test:v1"...
...resource 0...
...resource 1(ocm/value:v2.0)...
...resource 2(ocm/ref:v2.0)...
...adding component version...
1 versions transferred
`))

		Expect(env.FileExists(OUT)).To(BeTrue())
		Check(env, ldesc, OUT)
	})
})
