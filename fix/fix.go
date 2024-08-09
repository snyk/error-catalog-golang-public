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

// Package fix contains errors related to the namespace Fix
// of the Error Catalog.
package fix

import (
  "github.com/snyk/error-catalog-golang-public/snyk_errors"
  "github.com/google/uuid"
)

// NewFixScenarioNotSupportedError displays errors with the following description:
// Snyk failed to open a fix PR as the scenario is not supported.
func NewFixScenarioNotSupportedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#pr-failures-0001",
    Title:      "Fix scenario not supported",
    StatusCode: 422,
    ErrorCode:  "PR-FAILURES-0001",
    Classification: "UNSUPPORTED",
    Links: []string{},
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewSCMRateLimitError displays errors with the following description:
// SCM rate limit exceeded due to too many requests.
func NewSCMRateLimitError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#pr-failures-0002",
    Title:      "SCM rate limit",
    StatusCode: 429,
    ErrorCode:  "PR-FAILURES-0002",
    Classification: "ACTIONABLE",
    Links: []string{},
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewUnauthorisedAccessError displays errors with the following description:
// Request failed due to unathorised access. Please read documentation around adding users and permitted roles.
//
// Read more:
// - https://docs.snyk.io/snyk-admin/groups-and-organizations/organizations/manage-users-in-organizations
func NewUnauthorisedAccessError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#pr-failures-0003",
    Title:      "Unauthorised access",
    StatusCode: 403,
    ErrorCode:  "PR-FAILURES-0003",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/snyk-admin/groups-and-organizations/organizations/manage-users-in-organizations",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewUnsupportedEcosystemError displays errors with the following description:
// The language or package manager is not supported. Please refer to the supported package managers in the documentation.
//
// Read more:
// - https://docs.snyk.io/scan-using-snyk/pull-requests/snyk-fix-pull-or-merge-requests/upgrade-dependencies-with-automatic-prs/upgrade-open-source-dependencies-with-automatic-prs#supported-languages-and-scms
func NewUnsupportedEcosystemError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-packages-0001",
    Title:      "Unsupported ecosystem",
    StatusCode: 400,
    ErrorCode:  "SNYK-PACKAGES-0001",
    Classification: "UNSUPPORTED",
    Links: []string{
      "https://docs.snyk.io/scan-using-snyk/pull-requests/snyk-fix-pull-or-merge-requests/upgrade-dependencies-with-automatic-prs/upgrade-open-source-dependencies-with-automatic-prs#supported-languages-and-scms",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewMetadataNotFoundError displays errors with the following description:
// Package metadata not or found or missing.
//
// Read more:
// - https://docs.snyk.io/scan-using-snyk/pull-requests/snyk-fix-pull-or-merge-requests/upgrade-dependencies-with-automatic-prs/upgrade-private-dependencies-with-automatic-prs#private-packages-api
func NewMetadataNotFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-packages-0003",
    Title:      "Metadata not found",
    StatusCode: 404,
    ErrorCode:  "SNYK-PACKAGES-0003",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/scan-using-snyk/pull-requests/snyk-fix-pull-or-merge-requests/upgrade-dependencies-with-automatic-prs/upgrade-private-dependencies-with-automatic-prs#private-packages-api",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewNoMatureVersionsFoundError displays errors with the following description:
// Unable to provide a recommended version as no mature versions were found.
func NewNoMatureVersionsFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-packages-0005",
    Title:      "No mature versions found for package",
    StatusCode: 404,
    ErrorCode:  "SNYK-PACKAGES-0005",
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

// NewVersionNotFoundError displays errors with the following description:
// Unable to provide a recommended version for package using this policy.
func NewVersionNotFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-packages-0006",
    Title:      "No recommended version found",
    StatusCode: 404,
    ErrorCode:  "SNYK-PACKAGES-0006",
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

// NewAlreadyLatestVersionError displays errors with the following description:
// No newer version found for this package, as it is already to latest version.
func NewAlreadyLatestVersionError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-packages-0007",
    Title:      "Package is already at latest version",
    StatusCode: 404,
    ErrorCode:  "SNYK-PACKAGES-0007",
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

// NewDowngradeVersionUnsupportedError displays errors with the following description:
// Unable to suggest a downgrade for a package version.
func NewDowngradeVersionUnsupportedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-packages-0008",
    Title:      "Version downgrade is not supported",
    StatusCode: 400,
    ErrorCode:  "SNYK-PACKAGES-0008",
    Classification: "ACTIONABLE",
    Links: []string{},
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewVersionParsingError displays errors with the following description:
// Not a valid version for semver format.
//
// Read more:
// - https://semver.org/
func NewVersionParsingError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-packages-0009",
    Title:      "Invalid version",
    StatusCode: 400,
    ErrorCode:  "SNYK-PACKAGES-0009",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://semver.org/",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewFailedToGetPullRequestAttributesError displays errors with the following description:
// Snyk could not get the custom pull request template attributes, using the given variables and the fetched pr template.
//
// Read more:
// - https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta
func NewFailedToGetPullRequestAttributesError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-pr-template-0001",
    Title:      "Failed to get pull request attributes",
    StatusCode: 500,
    ErrorCode:  "SNYK-PR-TEMPLATE-0001",
    Classification: "UNEXPECTED",
    Links: []string{
      "https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewPullRequestTemplateNotFoundError displays errors with the following description:
// We could not find your pull request template, have you created one yet? Please check the attached link for instructions on how to setup your pull request template.
//
// Read more:
// - https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta
func NewPullRequestTemplateNotFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-pr-template-0002",
    Title:      "Not found",
    StatusCode: 404,
    ErrorCode:  "SNYK-PR-TEMPLATE-0002",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewFailedToCompilePrTemplateError displays errors with the following description:
// Could not compile your customize pull request template. Please check for syntax errors using the Snyk variables inside the template.
//
// Read more:
// - https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta
func NewFailedToCompilePrTemplateError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-pr-template-0003",
    Title:      "Failed to compile pull request template",
    StatusCode: 400,
    ErrorCode:  "SNYK-PR-TEMPLATE-0003",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewFailedToParsePullRequestAttributesError displays errors with the following description:
// Snyk could not parse the custom pull request template, using the given variables and assigning them to the fetched pr template.
//
// Read more:
// - https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta
func NewFailedToParsePullRequestAttributesError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-pr-template-0004",
    Title:      "Failed to parse pull request attributes",
    StatusCode: 500,
    ErrorCode:  "SNYK-PR-TEMPLATE-0004",
    Classification: "UNEXPECTED",
    Links: []string{
      "https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewFailedToLoadCompiledYamlError displays errors with the following description:
// Could not load YAML file after substituting Snyk variables into the custom PR template.
//
// Read more:
// - https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta
func NewFailedToLoadCompiledYamlError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-pr-template-0005",
    Title:      "Failed to load YAML file after substituting Snyk variables",
    StatusCode: 500,
    ErrorCode:  "SNYK-PR-TEMPLATE-0005",
    Classification: "UNEXPECTED",
    Links: []string{
      "https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewFailedToGenerateHashError displays errors with the following description:
// Snyk could not generate hash using the customer PR files and projects vulnIds.
//
// Read more:
// - https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta
func NewFailedToGenerateHashError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-pr-template-0006",
    Title:      "Failed to generate hash for custom PR template",
    StatusCode: 500,
    ErrorCode:  "SNYK-PR-TEMPLATE-0006",
    Classification: "UNEXPECTED",
    Links: []string{
      "https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewFailedToCreatePRTemplateError displays errors with the following description:
// Snyk could not create pull request template.
//
// Read more:
// - https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta
func NewFailedToCreatePRTemplateError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-pr-template-0007",
    Title:      "Unable to create pull request template",
    StatusCode: 500,
    ErrorCode:  "SNYK-PR-TEMPLATE-0007",
    Classification: "UNEXPECTED",
    Links: []string{
      "https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewFailedToReadPRTemplateError displays errors with the following description:
// Snyk could not get pull request template.
//
// Read more:
// - https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta
func NewFailedToReadPRTemplateError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-pr-template-0008",
    Title:      "Unable to get pull request template",
    StatusCode: 500,
    ErrorCode:  "SNYK-PR-TEMPLATE-0008",
    Classification: "UNEXPECTED",
    Links: []string{
      "https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewFailedToDeletePRTemplateError displays errors with the following description:
// Snyk could not delete pull request template.
//
// Read more:
// - https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta
func NewFailedToDeletePRTemplateError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-pr-template-0009",
    Title:      "Unable to delete pull request template",
    StatusCode: 500,
    ErrorCode:  "SNYK-PR-TEMPLATE-0009",
    Classification: "UNEXPECTED",
    Links: []string{
      "https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewPRTemplateInvalidPayloadError displays errors with the following description:
// The pull request template payload is invalid.
func NewPRTemplateInvalidPayloadError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-pr-template-0010",
    Title:      "Invalid payload",
    StatusCode: 500,
    ErrorCode:  "SNYK-PR-TEMPLATE-0010",
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

// NewFailedToLoadCompiledJSONError displays errors with the following description:
// Could not load JSON file after substituting Snyk variables into the custom PR template.
//
// Read more:
// - https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta
func NewFailedToLoadCompiledJSONError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-pr-template-0011",
    Title:      "Failed to load JSON file after substituting Snyk variables",
    StatusCode: 500,
    ErrorCode:  "SNYK-PR-TEMPLATE-0011",
    Classification: "UNEXPECTED",
    Links: []string{
      "https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewFailedToRenderDefaultTemplateError displays errors with the following description:
// Could not render default PR template.
//
// Read more:
// - https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta
func NewFailedToRenderDefaultTemplateError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-pr-template-0012",
    Title:      "Failed to render default PR template",
    StatusCode: 500,
    ErrorCode:  "SNYK-PR-TEMPLATE-0012",
    Classification: "UNEXPECTED",
    Links: []string{
      "https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

