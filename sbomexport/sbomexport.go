/*
 * © 2026 Snyk Limited
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

// Package sbomexport contains errors related to the namespace SbomExport
// of the Error Catalog.
package sbomexport

import (
  "github.com/snyk/error-catalog-golang-public/snyk_errors"
  "github.com/google/uuid"
)
// NewInternalServerError displays errors with the following description:
// An unexpected error occurred during the SBOM generation. Review the request, then try again. If the error persists, contact Snyk Support.
func NewInternalServerError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-9000",
    Title:      "SBOM generation export server error",
    Description: "An unexpected error occurred during the SBOM generation. Review the request, then try again. If the error persists, contact Snyk Support.",
    StatusCode: 500,
    ErrorCode:  "SNYK-OS-9000",
    Classification: "UNEXPECTED",
    Links: []string{},
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewUnexpectedDepGraphResponseError displays errors with the following description:
// An unexpected dependency graph error occurred. Review the request, then try again. If the error persists, contact Snyk Support.
func NewUnexpectedDepGraphResponseError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-9001",
    Title:      "Dependency graph error",
    Description: "An unexpected dependency graph error occurred. Review the request, then try again. If the error persists, contact Snyk Support.",
    StatusCode: 500,
    ErrorCode:  "SNYK-OS-9001",
    Classification: "UNEXPECTED",
    Links: []string{},
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewUnexpectedParseDepGraphError displays errors with the following description:
// The dependency graph cannot be parsed due to an unexpected error. Review the request, then try again. If the error persists, contact Snyk Support.
func NewUnexpectedParseDepGraphError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-9002",
    Title:      "Error parsing dependency graph",
    Description: "The dependency graph cannot be parsed due to an unexpected error. Review the request, then try again. If the error persists, contact Snyk Support.",
    StatusCode: 500,
    ErrorCode:  "SNYK-OS-9002",
    Classification: "UNEXPECTED",
    Links: []string{},
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewIaCOrSASTProjectError displays errors with the following description:
// Only SBOMs for Snyk Open Source or Snyk Container projects are supported.
func NewIaCOrSASTProjectError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-9003",
    Title:      "SBOM not supported due to project type",
    Description: "Only SBOMs for Snyk Open Source or Snyk Container projects are supported.",
    StatusCode: 404,
    ErrorCode:  "SNYK-OS-9003",
    Classification: "UNSUPPORTED",
    Links: []string{},
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewUnsupportedProjectError displays errors with the following description:
// Only SBOMs for open source projects are supported (Snyk Open Source).
func NewUnsupportedProjectError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-9004",
    Title:      "SBOM not supported",
    Description: "Only SBOMs for open source projects are supported (Snyk Open Source).",
    StatusCode: 404,
    ErrorCode:  "SNYK-OS-9004",
    Classification: "UNSUPPORTED",
    Links: []string{},
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewDepGraphResponseError displays errors with the following description:
// The server cannot process the request due to incomplete data. Review the request, then try again.
func NewDepGraphResponseError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-9005",
    Title:      "Dependency graph request cannot be processed",
    Description: "The server cannot process the request due to incomplete data. Review the request, then try again.",
    StatusCode: 404,
    ErrorCode:  "SNYK-OS-9005",
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

// NewMissingAuthTokenError displays errors with the following description:
// The API token is misconfigured or expired. Configure or generate the API token, then try again.
//
// Read more:
// - https://docs.snyk.io/snyk-api-info/revoking-and-regenerating-snyk-api-tokens
func NewMissingAuthTokenError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-9006",
    Title:      "Authorization failed due to missing API token",
    Description: "The API token is misconfigured or expired. Configure or generate the API token, then try again.",
    StatusCode: 401,
    ErrorCode:  "SNYK-OS-9006",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/snyk-api-info/revoking-and-regenerating-snyk-api-tokens",
    },
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewEmptyRequestBodyError displays errors with the following description:
// The body of the request is empty. Review the request, then try again.
func NewEmptyRequestBodyError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-9007",
    Title:      "Client request cannot be processed",
    Description: "The body of the request is empty. Review the request, then try again.",
    StatusCode: 400,
    ErrorCode:  "SNYK-OS-9007",
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

// NewInvalidDepGraphError displays errors with the following description:
// The supplied dependency graph was not valid. Review the request, then try again.
func NewInvalidDepGraphError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-9008",
    Title:      "Invalid dependency graph",
    Description: "The supplied dependency graph was not valid. Review the request, then try again.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-9008",
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
