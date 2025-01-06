/*
 * © 2025 Snyk Limited
 *
 * Licensed under the Apache License, Version 2.0 (the 'License');
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an 'AS IS' BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package unmanaged contains errors related to the namespace OpenSourceUnmanaged
// of the Error Catalog.
package unmanaged

import (
  "github.com/snyk/error-catalog-golang-public/snyk_errors"
  "github.com/google/uuid"
)

// NewMavenSearchServiceUnavailableError displays errors with the following description:
// The upstream Maven search service is not available.
//
// Read more:
// - https://search.maven.org
// - https://status.maven.org
func NewMavenSearchServiceUnavailableError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-osjvm-001",
    Title:      "Maven search service unavailable",
    Description: "The upstream Maven search service is not available.",
    StatusCode: 503,
    ErrorCode:  "SNYK-OSJVM-001",
    Classification: "UNEXPECTED",
    Links: []string{
      "https://search.maven.org",
      "https://status.maven.org",
    },
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewSha1NotFoundError displays errors with the following description:
// Unable to find the coordinates for the provided SHA1. Please verify the data you are sending and try again.
//
// Read more:
// - https://docs.snyk.io/snyk-cli/test-for-vulnerabilities/scan-all-unmanaged-jar-files
func NewSha1NotFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-osjvm-002",
    Title:      "SHA1 not found",
    Description: "Unable to find the coordinates for the provided SHA1. Please verify the data you are sending and try again.",
    StatusCode: 404,
    ErrorCode:  "SNYK-OSJVM-002",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/snyk-cli/test-for-vulnerabilities/scan-all-unmanaged-jar-files",
    },
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

