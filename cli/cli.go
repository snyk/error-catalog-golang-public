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

// Package cli contains errors related to the namespace CLI
// of the Error Catalog.
package cli

import (
  "github.com/snyk/error-catalog-golang-public/snyk_errors"
  "github.com/google/uuid"
)

// NewUnableToCreateMonitorError displays errors with the following description:
// There was an unexpected error when attempting to monitor specified project.
//
// Read more:
// - https://docs.snyk.io/snyk-cli
func NewUnableToCreateMonitorError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-7001",
    Title:      "Unable to create monitor",
    StatusCode: 500,
    ErrorCode:  "SNYK-OS-7001",
    Classification: "UNEXPECTED",
    Links: []string{
      "https://docs.snyk.io/snyk-cli",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewConnectionTimeoutError displays errors with the following description:
// A request to the Snyk CLI has unexpectedly timeout.
//
// Read more:
// - https://docs.snyk.io/snyk-cli
func NewConnectionTimeoutError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-7002",
    Title:      "Request to Snyk API timeout",
    StatusCode: 504,
    ErrorCode:  "SNYK-OS-7002",
    Classification: "UNEXPECTED",
    Links: []string{
      "https://docs.snyk.io/snyk-cli",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

