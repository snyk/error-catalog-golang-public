/*
 * © 2023 Snyk Limited
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

// Package prchecks contains errors related to the namespace PRChecks
// of the Error Catalog.
package prchecks

import (
  "github.com/snyk/error-catalog-golang-public/snyk_errors"
  "github.com/google/uuid"
)

// NewFailedToReadManifestError displays errors with the following description:
// Snyk failed to read 1 or more manifest files.
// Sometimes things go wrong: a flaky connection, 3rd party services go down and Snyk is unable to read the files needed in order to test your project. 
// 
// If this happens, you could try:
// 
// - Opening and re-opening your Pull Request / Merge Request, to kick off a new test
// - Removing and re-adding the repo to Snyk
// 
// Ultimately, you should contact support@snyk.io if the issue persists
//
// Read more:
// - https://support.snyk.io/hc/en-us/articles/360000910517-Failed-to-read-manifest-file
func NewFailedToReadManifestError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-pr-check-0001",
    Title:      "Error reading manifest",
    StatusCode: 500,
    ErrorCode:  "SNYK-PR-CHECK-0001",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://support.snyk.io/hc/en-us/articles/360000910517-Failed-to-read-manifest-file",
    },
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

