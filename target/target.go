/*
 * © 2025 Snyk Limited
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

// Package target contains errors related to the namespace Target
// of the Error Catalog.
package target

import (
  "github.com/snyk/error-catalog-golang-public/snyk_errors"
  "github.com/google/uuid"
)

// NewTargetNotFoundError displays errors with the following description:
// Snyk was unable to resolve the imported target. Ensure that Snyk created the target and try again.
//
// Read more:
// - https://docs.snyk.io/snyk-admin/snyk-projects#target
func NewTargetNotFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-target-0001",
    Title:      "Target not found",
    Description: "Snyk was unable to resolve the imported target. Ensure that Snyk created the target and try again.",
    StatusCode: 404,
    ErrorCode:  "SNYK-TARGET-0001",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/snyk-admin/snyk-projects#target",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewNoUniqueTargetFoundError displays errors with the following description:
// Snyk was unable to resolve a single target. Snyk found multiple targets configured for the same integration and repository URL pair. Ensure a unique target exists.
//
// Read more:
// - https://docs.snyk.io/snyk-admin/snyk-projects#target
func NewNoUniqueTargetFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-target-0002",
    Title:      "No unique target found",
    Description: "Snyk was unable to resolve a single target. Snyk found multiple targets configured for the same integration and repository URL pair. Ensure a unique target exists.",
    StatusCode: 422,
    ErrorCode:  "SNYK-TARGET-0002",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/snyk-admin/snyk-projects#target",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

