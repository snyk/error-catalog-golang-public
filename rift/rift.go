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

// Package rift contains errors related to the namespace Rift
// of the Error Catalog.
package rift

import (
  "github.com/snyk/error-catalog-golang-public/snyk_errors"
  "github.com/google/uuid"
)
// NewScanFailedError displays errors with the following description:
// The vulnerability detection agent could not complete the scan. Try again. If the problem persists, run the command with `-d` and contact Snyk Support with the scan job ID.
func NewScanFailedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-rift-0001",
    Title:      "Vulnerability detection agent scan failed",
    Description: "The vulnerability detection agent could not complete the scan. Try again. If the problem persists, run the command with `-d` and contact Snyk Support with the scan job ID.",
    StatusCode: 500,
    ErrorCode:  "SNYK-RIFT-0001",
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
