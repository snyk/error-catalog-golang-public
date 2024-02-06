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

