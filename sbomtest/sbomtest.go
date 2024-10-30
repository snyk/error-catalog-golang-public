/*
 * © 2024 Snyk Limited
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

// Package sbomtest contains errors related to the namespace SbomTest
// of the Error Catalog.
package sbomtest

import (
  "github.com/snyk/error-catalog-golang-public/snyk_errors"
  "github.com/google/uuid"
)

// NewInternalError displays errors with the following description:
// An unexpected error occurred. Review the request, then try again. If the error persists, contact Snyk Support.
func NewInternalError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-sbom-0001",
    Title:      "SBOM test error",
    Description: "An unexpected error occurred. Review the request, then try again. If the error persists, contact Snyk Support.",
    StatusCode: 500,
    ErrorCode:  "SNYK-SBOM-0001",
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

// NewOrgIDMismatchError displays errors with the following description:
// The requested organization ID does not match the owner of the SBOM test ID.
// 
// This error occurs when the supplied organization ID is different to the one used when creating an SBOM test run.
// Ensure the organization ID used to make the request is the same as the the organization ID used to create the SBOM test.
func NewOrgIDMismatchError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-sbom-0002",
    Title:      "Organization ID mismatch",
    Description: "The requested organization ID does not match the owner of the SBOM test ID.\n\nThis error occurs when the supplied organization ID is different to the one used when creating an SBOM test run.\nEnsure the organization ID used to make the request is the same as the the organization ID used to create the SBOM test.",
    StatusCode: 404,
    ErrorCode:  "SNYK-SBOM-0002",
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

// NewNotFoundError displays errors with the following description:
// Snyk was unable to find the requested SBOM test.
func NewNotFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-sbom-0003",
    Title:      "Unable to find SBOM test",
    Description: "Snyk was unable to find the requested SBOM test.",
    StatusCode: 404,
    ErrorCode:  "SNYK-SBOM-0003",
    Classification: "UNEXPECTED",
    Links: []string{},
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewFailedTestError displays errors with the following description:
// The SBOM test failed and does not have any results.
// 
// This error occurs when results for a failed SBOM test are being requested.
func NewFailedTestError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-sbom-0004",
    Title:      "SBOM test failed",
    Description: "The SBOM test failed and does not have any results.\n\nThis error occurs when results for a failed SBOM test are being requested.",
    StatusCode: 404,
    ErrorCode:  "SNYK-SBOM-0004",
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

// NewPendingTestError displays errors with the following description:
// The SBOM test is still processing and does not have any results yet.
// 
// This error occurs when the results for an SBOM test have been requested, but the SBOM test is still being processed.
func NewPendingTestError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-sbom-0005",
    Title:      "SBOM test results still pending",
    Description: "The SBOM test is still processing and does not have any results yet.\n\nThis error occurs when the results for an SBOM test have been requested, but the SBOM test is still being processed.",
    StatusCode: 404,
    ErrorCode:  "SNYK-SBOM-0005",
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

// NewFormatUnknownError displays errors with the following description:
// Snyk does not recognize the SBOM file format.
// 
// Provide an SBOM document with a supported format.
//
// Read more:
// - https://docs.snyk.io/snyk-cli/commands/sbom-test
func NewFormatUnknownError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-sbom-0006",
    Title:      "Unknown SBOM format",
    Description: "Snyk does not recognize the SBOM file format.\n\nProvide an SBOM document with a supported format.",
    StatusCode: 422,
    ErrorCode:  "SNYK-SBOM-0006",
    Classification: "UNSUPPORTED",
    Links: []string{
      "https://docs.snyk.io/snyk-cli/commands/sbom-test",
    },
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewUnprocessableInputError displays errors with the following description:
// Snyk is unable to decode the SBOM file. Provide a valid SBOM document and try again.
//
// Read more:
// - https://docs.snyk.io/snyk-cli/commands/sbom-test
func NewUnprocessableInputError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-sbom-0007",
    Title:      "Unable to process SBOM input",
    Description: "Snyk is unable to decode the SBOM file. Provide a valid SBOM document and try again.",
    StatusCode: 422,
    ErrorCode:  "SNYK-SBOM-0007",
    Classification: "UNSUPPORTED",
    Links: []string{
      "https://docs.snyk.io/snyk-cli/commands/sbom-test",
    },
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewFormatNotSupportedError displays errors with the following description:
// Provide a supported format of the SBOM document and try again.
//
// Read more:
// - https://docs.snyk.io/snyk-cli/commands/sbom-test
func NewFormatNotSupportedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-sbom-0008",
    Title:      "SBOM format not supported",
    Description: "Provide a supported format of the SBOM document and try again.",
    StatusCode: 422,
    ErrorCode:  "SNYK-SBOM-0008",
    Classification: "UNSUPPORTED",
    Links: []string{
      "https://docs.snyk.io/snyk-cli/commands/sbom-test",
    },
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewConversionFailedError displays errors with the following description:
// Snyk was unable to process the provided SBOM input and is unable to scan it for issues.
func NewConversionFailedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-sbom-0009",
    Title:      "SBOM analysis failed",
    Description: "Snyk was unable to process the provided SBOM input and is unable to scan it for issues.",
    StatusCode: 422,
    ErrorCode:  "SNYK-SBOM-0009",
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

// NewNoTestablePackagesError displays errors with the following description:
// The SBOM document you provided does not contain any packages supported by Snyk vulnerability analysis, or it does not contain any package.
func NewNoTestablePackagesError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-sbom-0010",
    Title:      "No testable packages found",
    Description: "The SBOM document you provided does not contain any packages supported by Snyk vulnerability analysis, or it does not contain any package.",
    StatusCode: 422,
    ErrorCode:  "SNYK-SBOM-0010",
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

