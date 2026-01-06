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

// Package scm contains errors related to the namespace SCM
// of the Error Catalog.
package scm

import (
  "github.com/snyk/error-catalog-golang-public/snyk_errors"
  "github.com/google/uuid"
)
// NewUnsupportedIntegrationTypeError displays errors with the following description:
// The integration you provided does not support SCM repository access.
//
// Read more:
// - https://docs.snyk.io/scm-ide-and-ci-cd-workflow-and-integrations/snyk-scm-integrations
func NewUnsupportedIntegrationTypeError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-scm-0001",
    Title:      "Integration type not supported",
    Description: "The integration you provided does not support SCM repository access.",
    StatusCode: 500,
    ErrorCode:  "SNYK-SCM-0001",
    Classification: "UNSUPPORTED",
    Links: []string{
      "https://docs.snyk.io/scm-ide-and-ci-cd-workflow-and-integrations/snyk-scm-integrations",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewRevisionNotResolvedError displays errors with the following description:
// Snyk was unable to resolve the SCM revision you provided. Provide a valid revision, either a full commit ID or an existing commit reference.
func NewRevisionNotResolvedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-scm-0002",
    Title:      "Revision cannot be resolved",
    Description: "Snyk was unable to resolve the SCM revision you provided. Provide a valid revision, either a full commit ID or an existing commit reference.",
    StatusCode: 422,
    ErrorCode:  "SNYK-SCM-0002",
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

// NewIntegrationAuthenticationFailedError displays errors with the following description:
// Snyk was unable to authenticate with your SCM provider. Ensure you are using valid credentials for the Snyk integration.
func NewIntegrationAuthenticationFailedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-scm-0003",
    Title:      "Integration authentication failed",
    Description: "Snyk was unable to authenticate with your SCM provider. Ensure you are using valid credentials for the Snyk integration.",
    StatusCode: 401,
    ErrorCode:  "SNYK-SCM-0003",
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

// NewIntegrationAuthorizationFailedError displays errors with the following description:
// Snyk was unable to authorize with your SCM provider. If your Organization has SAML SSO enabled or enforced, re-authorize the OAuth Application `Snyk`.
func NewIntegrationAuthorizationFailedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-scm-0004",
    Title:      "Integration authorization failed",
    Description: "Snyk was unable to authorize with your SCM provider. If your Organization has SAML SSO enabled or enforced, re-authorize the OAuth Application `Snyk`.",
    StatusCode: 401,
    ErrorCode:  "SNYK-SCM-0004",
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

// NewFilesLimitExceededError displays errors with the following description:
// Snyk was unable to retrieve the repository because the overall file count exceeds the limit of 40000.
// 
// To reduce the file count, use a `.snyk` file to ignore certain directories or files. Alternatively, analyze individual work subdirectories separately.
func NewFilesLimitExceededError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-scm-0005",
    Title:      "Too many files",
    Description: "Snyk was unable to retrieve the repository because the overall file count exceeds the limit of 40000.\n\nTo reduce the file count, use a `.snyk` file to ignore certain directories or files. Alternatively, analyze individual work subdirectories separately.",
    StatusCode: 500,
    ErrorCode:  "SNYK-SCM-0005",
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

// NewSizeLimitExceededError displays errors with the following description:
// Snyk was unable to retrieve the repository because the size of the repository exceeds 15 GB.
// 
// To reduce the overall size of the repository, use a a `.snyk` file to ignore certain directories or files. Alternatively, analyze individual work subdirectories separately.
func NewSizeLimitExceededError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-scm-0006",
    Title:      "Repository size too large",
    Description: "Snyk was unable to retrieve the repository because the size of the repository exceeds 15 GB.\n\nTo reduce the overall size of the repository, use a a `.snyk` file to ignore certain directories or files. Alternatively, analyze individual work subdirectories separately.",
    StatusCode: 500,
    ErrorCode:  "SNYK-SCM-0006",
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

// NewResourceNotFoundError displays errors with the following description:
// Snyk was unable to resolve the SCM resource you provided. Provide a valid resource path and revision.
func NewResourceNotFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-scm-0010",
    Title:      "Specified resource cannot be found",
    Description: "Snyk was unable to resolve the SCM resource you provided. Provide a valid resource path and revision.",
    StatusCode: 422,
    ErrorCode:  "SNYK-SCM-0010",
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
