
/*
 * © 2023 Snyk Limited
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
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
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-pr-template-0001",
    Title:      "Failed to get pull request attributes",
    StatusCode: 500,
    ErrorCode:  "SNYK-PR-TEMPLATE-0001",
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
// Could not find pull request template.
func NewPullRequestTemplateNotFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-pr-template-0002",
    Title:      "Not found",
    StatusCode: 404,
    ErrorCode:  "SNYK-PR-TEMPLATE-0002",
    Links: []string{},
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewFailedToCompilePrTemplateError displays errors with the following description:
// Could not compile your customize pull request template, using Handlebars compilation and Snyk variables in place.
//
// Read more:
// - https://docs.snyk.io/scan-application-code/snyk-open-source/open-source-basics/customize-pr-templates-closed-beta
func NewFailedToCompilePrTemplateError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-pr-template-0003",
    Title:      "Failed to compile pull request template",
    StatusCode: 400,
    ErrorCode:  "SNYK-PR-TEMPLATE-0003",
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
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-pr-template-0004",
    Title:      "Failed to parse pull request attributes",
    StatusCode: 500,
    ErrorCode:  "SNYK-PR-TEMPLATE-0004",
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
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-pr-template-0005",
    Title:      "Failed to load YAML file after substituting Snyk variables",
    StatusCode: 400,
    ErrorCode:  "SNYK-PR-TEMPLATE-0005",
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
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-pr-template-0006",
    Title:      "Failed to generate hash for custom PR template",
    StatusCode: 500,
    ErrorCode:  "SNYK-PR-TEMPLATE-0006",
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

