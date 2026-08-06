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

// Package secrets contains errors related to the namespace Secrets
// of the Error Catalog.
package secrets

import (
  "github.com/snyk/error-catalog-golang-public/snyk_errors"
  "github.com/google/uuid"
)
// NewTestError displays errors with the following description:
// An unexpected error occurred. Review the request, then try again. If the error persists, contact Snyk Support.
func NewTestError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-secrets-0001",
    Title:      "Secrets test error",
    Description: "An unexpected error occurred. Review the request, then try again. If the error persists, contact Snyk Support.",
    StatusCode: 500,
    ErrorCode:  "SNYK-SECRETS-0001",
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

// NewNotEnabledError displays errors with the following description:
// This error occurs when Snyk Secrets is not enabled for the current Organization. Activate Snyk Secrets and try again.
func NewNotEnabledError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-secrets-0002",
    Title:      "Snyk Secrets is not enabled",
    Description: "This error occurs when Snyk Secrets is not enabled for the current Organization. Activate Snyk Secrets and try again.",
    StatusCode: 403,
    ErrorCode:  "SNYK-SECRETS-0002",
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

// NewAnalysisError displays errors with the following description:
// An error occurred while analyzing the files for secrets. Review the request, then try again. If the error persists, contact Snyk Support.
func NewAnalysisError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-secrets-0003",
    Title:      "Secrets analysis error",
    Description: "An error occurred while analyzing the files for secrets. Review the request, then try again. If the error persists, contact Snyk Support.",
    StatusCode: 500,
    ErrorCode:  "SNYK-SECRETS-0003",
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
