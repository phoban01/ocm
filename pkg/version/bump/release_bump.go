// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Masterminds/semver/v3"
	"golang.org/x/tools/go/ast/astutil"

	"github.com/open-component-model/ocm/pkg/version"
)

// This program will auto-increment the release versions the specified "releseVersionFile".
//
// The program can be executed using the following command:
// go run ./pkg/version/bump --type [patch|minor|major] [--pre]
//
// The program will increment the semver part specified in the type flag value:
// Where type is "patch": v2.1.1 becomes v2.1.2
// Where type is "minor": v2.1.1 becomes v2.2.0
// Where type is "major": v2.1.1 becomes v3.0.0
//
// When the "pre" flag is set the program will use the following rules:
//
// 1: if this is the first pre release then increment the semver part specified in the type flag
// and set the suffix to rc.1
// 2: if this is a subsequent pre-release then increment only the suffix part
// 3: when pre-releases are complete, remove the suffix and don't increment the semver, regardless of type flag

const releaseVersionFile = "./pkg/version/release.go"

func main() {
	var (
		releaseType        string
		isReleaseCandidate bool
	)

	flag.StringVar(&releaseType, "type", "patch", "set the release type to: patch, minor or major")
	flag.BoolVar(&isReleaseCandidate, "pre", false, "denotes whether this is a pre-release")
	flag.Parse()

	vers := semver.MustParse(version.ReleaseVersion)

	var newVers semver.Version

	//nolint:forbidigo // Logger not needed for this command.
	switch releaseType {
	case "patch":
		newVers = vers.IncPatch()
	case "minor":
		newVers = vers.IncMinor()
	case "major":
		newVers = vers.IncMajor()
	default:
		log.Println("type must be one of: patch, minor, major")
		os.Exit(1)
		return
	}

	fset := token.NewFileSet()

	path, err := filepath.Abs(releaseVersionFile)
	if err != nil {
		log.Fatal(err)
	}

	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	astutil.Apply(file, nil, func(c *astutil.Cursor) bool {
		n := c.Node()
		switch x := n.(type) {
		case *ast.ValueSpec:
			for i := len(x.Values); i > 0; i-- {
				name := x.Names[i-1].Name
				switch name {
				case "ReleaseVersion":
					if isReleaseCandidate && version.ReleaseCandidateNumber > 0 {
						return true
					} else if version.ReleaseCandidateNumber > 0 {
						return true
					}
					x.Values[i-1] = &ast.BasicLit{
						Kind:  token.STRING,
						Value: strconv.Quote(newVers.Original()),
					}
				case "ReleaseCandidateNumber":
					rc, err := strconv.Atoi(x.Values[i-1].(*ast.BasicLit).Value)
					if err != nil {
						os.Exit(1)
						return false
					}

					if !isReleaseCandidate {
						rc = 0
					} else if rc == 0 {
						rc = 1
					} else {
						rc++
					}

					x.Values[i-1] = &ast.BasicLit{
						Kind:  token.INT,
						Value: strconv.Itoa(rc),
					}
				}
			}
		}

		return true
	})

	printer.Fprint(os.Stdout, fset, file)

	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	if err := printer.Fprint(f, fset, file); err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
