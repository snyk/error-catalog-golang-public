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

// Package isolatedbuilds contains errors related to the namespace IsolatedBuilds
// of the Error Catalog.
package isolatedbuilds

import (
  "github.com/snyk/error-catalog-golang-public/snyk_errors"
  "github.com/google/uuid"
)

// NewInvalidRequestError displays errors with the following description:
// The provided request payload is not valid for the selected ecosystem. Please review the API documentation.
//
// Read more:
// - https://apidocs.snyk.io/
func NewInvalidRequestError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-8001",
    Title:      "Invalid request",
    StatusCode: 400,
    ErrorCode:  "SNYK-OS-8001",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://apidocs.snyk.io/",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewBuildEnvironmentNotFoundError displays errors with the following description:
// The build environment for the provided context could not be found. Please ensure you have created the build environment first.
func NewBuildEnvironmentNotFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-8002",
    Title:      "Build environment not found",
    StatusCode: 404,
    ErrorCode:  "SNYK-OS-8002",
    Classification: "ACTIONABLE",
    Links: []string{},
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewUnsupportedEcosystemError displays errors with the following description:
// The language or package manager is not supported. Please refer to the supported package managers in the links.
//
// Read more:
// - https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/supported-languages-frameworks-and-feature-availability-overview#open-source-and-licensing-snyk-open-source
func NewUnsupportedEcosystemError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-8003",
    Title:      "Unsupported Ecosystem",
    StatusCode: 400,
    ErrorCode:  "SNYK-OS-8003",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/supported-languages-frameworks-and-feature-availability-overview#open-source-and-licensing-snyk-open-source",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

