// SPDX-FileCopyrightText: 2022 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

package version

import "fmt"

// ReleaseVersion is the version number in semver format "vX.Y.Z", prefixed with "v".
var ReleaseVersion = "v0.2.0"

// ReleaseCandidateNumber is used to auto-increment pre-releases
var ReleaseCandidateNumber = 0

// ReleaseCandidate is the release candidate ID in format "rc.X", which will be appended to the release version.
var ReleaseCandidate = fmt.Sprintf("rc.%d", ReleaseCandidateNumber)
