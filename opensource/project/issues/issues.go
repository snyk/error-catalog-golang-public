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

// Package issues contains errors related to the namespace OpenSourceProjectIssues
// of the Error Catalog.
package issues

import (
  "github.com/snyk/error-catalog-golang-public/snyk_errors"
  "github.com/google/uuid"
)

// NewInvalidRequestError displays errors with the following description:
// Check the body of your request and try again.
func NewInvalidRequestError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-ossi-ospi-1001",
    Title:      "Invalid request",
    Description: "Check the body of your request and try again.",
    StatusCode: 400,
    ErrorCode:  "SNYK-OSSI-OSPI-1001",
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

// NewInvalidResponseError displays errors with the following description:
// This issue is unexpected, and the service will recover shortly. If the error still occurs, contact support.
func NewInvalidResponseError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-ossi-ospi-1002",
    Title:      "Unable to return valid API response",
    Description: "This issue is unexpected, and the service will recover shortly. If the error still occurs, contact support.",
    StatusCode: 500,
    ErrorCode:  "SNYK-OSSI-OSPI-1002",
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

// NewDataTransformationError displays errors with the following description:
// This issue is unexpected, and the service will recover shortly. If the error still occurs, contact support.
func NewDataTransformationError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-ossi-ospi-2001",
    Title:      "Failed to process data",
    Description: "This issue is unexpected, and the service will recover shortly. If the error still occurs, contact support.",
    StatusCode: 500,
    ErrorCode:  "SNYK-OSSI-OSPI-2001",
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

// NewStorageFailureError displays errors with the following description:
// Check inputs and then try again. If the error still occurs, contact support.
func NewStorageFailureError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-ossi-ospi-3001",
    Title:      "Failed to store issue data",
    Description: "Check inputs and then try again. If the error still occurs, contact support.",
    StatusCode: 500,
    ErrorCode:  "SNYK-OSSI-OSPI-3001",
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

// NewInternalServerError displays errors with the following description:
// This issue is unexpected, and the service will recover shortly. If the error still occurs, contact support.
func NewInternalServerError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-ossi-ospi-4001",
    Title:      "Internal server error",
    Description: "This issue is unexpected, and the service will recover shortly. If the error still occurs, contact support.",
    StatusCode: 500,
    ErrorCode:  "SNYK-OSSI-OSPI-4001",
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

