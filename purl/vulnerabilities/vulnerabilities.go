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

// Package vulnerabilities contains errors related to the namespace PurlVulnerabilityFetching
// of the Error Catalog.
package vulnerabilities

import (
  "github.com/snyk/error-catalog-golang-public/snyk_errors"
  "github.com/google/uuid"
)

// NewOrganizationNotWhitelistedError displays errors with the following description:
// You likely don’t have access to the features in Beta. To get access, you can request access to features in Beta through your account manager or team.
func NewOrganizationNotWhitelistedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-ossi-1040",
    Title:      "Your Organisation is not authorized to perform this action",
    StatusCode: 403,
    ErrorCode:  "SNYK-OSSI-1040",
    Links: []string{},
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewAuthorizationRequestFailureError displays errors with the following description:
// Unexpected error when authenticating. Try again, and if the error still occurs, contact support.
func NewAuthorizationRequestFailureError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-ossi-1050",
    Title:      "Authorization request failure",
    StatusCode: 500,
    ErrorCode:  "SNYK-OSSI-1050",
    Links: []string{},
    Level:  "fatal",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewInvalidPurlError displays errors with the following description:
// Make sure that the purl is valid. See the Package URL specification link for further information.
//
// Read more:
// - https://github.com/package-url/purl-spec/blob/master/PURL-SPECIFICATION.rst
func NewInvalidPurlError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-ossi-2010",
    Title:      "Invalid purl",
    StatusCode: 400,
    ErrorCode:  "SNYK-OSSI-2010",
    Links: []string{
      "https://github.com/package-url/purl-spec/blob/master/PURL-SPECIFICATION.rst",
    },
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewNamespaceNotProvidedError displays errors with the following description:
// You have requested a package type that requires a namespace (e.g. maven group id). Provide the namespace to retrieve the package.
//
// Read more:
// - https://github.com/package-url/purl-spec/blob/master/PURL-SPECIFICATION.rst
func NewNamespaceNotProvidedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-ossi-2011",
    Title:      "Namespace not specified",
    StatusCode: 400,
    ErrorCode:  "SNYK-OSSI-2011",
    Links: []string{
      "https://github.com/package-url/purl-spec/blob/master/PURL-SPECIFICATION.rst",
    },
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewUnsupportedEcosystemError displays errors with the following description:
// The package type is not supported. Check the List issues for a package in Snyk API.
func NewUnsupportedEcosystemError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-ossi-2020",
    Title:      "Unsupported ecosystem",
    StatusCode: 400,
    ErrorCode:  "SNYK-OSSI-2020",
    Links: []string{},
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewMissingComponentError displays errors with the following description:
// A list of components of the purl spec is required. The purl did not specify all the required components.
func NewMissingComponentError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-ossi-2021",
    Title:      "Purl components required",
    StatusCode: 400,
    ErrorCode:  "SNYK-OSSI-2021",
    Links: []string{},
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewComponentNotSupportedError displays errors with the following description:
// Remove the unsupported component and retry the request.
func NewComponentNotSupportedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-ossi-2022",
    Title:      "Unsupported purl components",
    StatusCode: 400,
    ErrorCode:  "SNYK-OSSI-2022",
    Links: []string{},
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewUnsupportedGoVersionFormatError displays errors with the following description:
// Go pseudo version not supported.
func NewUnsupportedGoVersionFormatError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-ossi-2023",
    Title:      "Go version format not supported",
    StatusCode: 400,
    ErrorCode:  "SNYK-OSSI-2023",
    Links: []string{},
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewPackageNotFoundError displays errors with the following description:
// The package you specified in the purl cannot be found in the vulnerability database. Check the package name, ecosystem, and version, then try again.
func NewPackageNotFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-ossi-2030",
    Title:      "Requested package not found",
    StatusCode: 404,
    ErrorCode:  "SNYK-OSSI-2030",
    Links: []string{},
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewVulnerabilityServiceUnavailableError displays errors with the following description:
// This issue is unexpected, and the service will recover shortly. If the error still occurs, contact support.
func NewVulnerabilityServiceUnavailableError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-ossi-2031",
    Title:      "Vulnerability service not available",
    StatusCode: 503,
    ErrorCode:  "SNYK-OSSI-2031",
    Links: []string{},
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewVulnDBInvalidResponseError displays errors with the following description:
// An unexpected error occurred. Please try again, and if you continue to experience issues please contact support.
func NewVulnDBInvalidResponseError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-ossi-2032",
    Title:      "This issue is unexpected and the service should recover quickly if not please contact support",
    StatusCode: 500,
    ErrorCode:  "SNYK-OSSI-2032",
    Links: []string{},
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewVulndbNextError displays errors with the following description:
// An unexpected error occurred with the vulnerability service. Please try again, and if you continue to experience issues please contact support.
func NewVulndbNextError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-ossi-2033",
    Title:      "This issue is unexpected and the service should recover quickly if not please contact support",
    StatusCode: 500,
    ErrorCode:  "SNYK-OSSI-2033",
    Links: []string{},
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewInternalServerError displays errors with the following description:
// This issue is unexpected, and the service will recover shortly. If the error still occurs, contact support.
func NewInternalServerError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-ossi-2040",
    Title:      "Request not processed due to unexpected error",
    StatusCode: 500,
    ErrorCode:  "SNYK-OSSI-2040",
    Links: []string{},
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewInvalidPaginationParametersError displays errors with the following description:
// The pagination limit is > 1 and ≤ 1000, and the offset is ≥0.
func NewInvalidPaginationParametersError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-ossi-2041",
    Title:      "Invalid pagination parameters",
    StatusCode: 400,
    ErrorCode:  "SNYK-OSSI-2041",
    Links: []string{},
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewTooManyPurlsError displays errors with the following description:
// The number of purls sent in the request exceeds the limit of 1000 set by the service.
func NewTooManyPurlsError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-ossi-2042",
    Title:      "purls exceed limit",
    StatusCode: 400,
    ErrorCode:  "SNYK-OSSI-2042",
    Links: []string{},
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewTooManyIssuesError displays errors with the following description:
// The number of issues found for the provided purls exceeds the limit defined by the API. Reduce the number of purls sent in a single request.
func NewTooManyIssuesError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-ossi-2043",
    Title:      "Number of issues exceeds limit",
    StatusCode: 400,
    ErrorCode:  "SNYK-OSSI-2043",
    Links: []string{},
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewUndefinedContainerDistroError displays errors with the following description:
// The given Package URL does not have a required distro qualifier.
//
// Read more:
// - https://docs.snyk.io/scan-containers/how-snyk-container-works/supported-operating-system-distributions#debian
func NewUndefinedContainerDistroError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-ossi-2044",
    Title:      "Expected distro to be present",
    StatusCode: 400,
    ErrorCode:  "SNYK-OSSI-2044",
    Links: []string{
      "https://docs.snyk.io/scan-containers/how-snyk-container-works/supported-operating-system-distributions#debian",
    },
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewUnsupportedDebianDistroError displays errors with the following description:
// This Debian distro is currently not supported.
func NewUnsupportedDebianDistroError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-ossi-2045",
    Title:      "Unsupported Debian distro",
    StatusCode: 400,
    ErrorCode:  "SNYK-OSSI-2045",
    Links: []string{},
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewUndefinedContainerVendorError displays errors with the following description:
// The given Package URL does not have a required namespace.
func NewUndefinedContainerVendorError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-ossi-2046",
    Title:      "Expected namespace to be present",
    StatusCode: 400,
    ErrorCode:  "SNYK-OSSI-2046",
    Links: []string{},
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewUnsupportedContainerVendorError displays errors with the following description:
// The given Package URL does not contain a supported vendor.
func NewUnsupportedContainerVendorError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-ossi-2047",
    Title:      "Unsupported vendor",
    StatusCode: 400,
    ErrorCode:  "SNYK-OSSI-2047",
    Links: []string{},
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

