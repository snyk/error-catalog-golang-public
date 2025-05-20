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

// Package aibom contains errors related to the namespace AiBom
// of the Error Catalog.
package aibom

import (
  "github.com/snyk/error-catalog-golang-public/snyk_errors"
  "github.com/google/uuid"
)

// NewInternalError displays errors with the following description:
// An unexpected error occurred in the AI-BOM request. Review the request while providing the debug command flag `-d`. If the error persists, contact Snyk Support.
func NewInternalError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-ai-bom-0001",
    Title:      "Unexpected error",
    Description: "An unexpected error occurred in the AI-BOM request. Review the request while providing the debug command flag `-d`. If the error persists, contact Snyk Support.",
    StatusCode: 500,
    ErrorCode:  "SNYK-AI-BOM-0001",
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

// NewForbiddenError displays errors with the following description:
// You or your Organization do not have permission to use the AI-BOM feature. Check your user permissions or contact Snyk support.
func NewForbiddenError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-ai-bom-0002",
    Title:      "Forbidden",
    Description: "You or your Organization do not have permission to use the AI-BOM feature. Check your user permissions or contact Snyk support.",
    StatusCode: 403,
    ErrorCode:  "SNYK-AI-BOM-0002",
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

// NewNoSupportedFilesError displays errors with the following description:
// Snyk was unable to find any supported files for the aibom command. Ensure the directory you are scanning contains supported files.
func NewNoSupportedFilesError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-ai-bom-0003",
    Title:      "No supported files",
    Description: "Snyk was unable to find any supported files for the aibom command. Ensure the directory you are scanning contains supported files.",
    StatusCode: 422,
    ErrorCode:  "SNYK-AI-BOM-0003",
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

