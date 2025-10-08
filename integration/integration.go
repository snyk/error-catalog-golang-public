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

// Package integration contains errors related to the namespace Integration
// of the Error Catalog.
package integration

import (
  "github.com/snyk/error-catalog-golang-public/snyk_errors"
  "github.com/google/uuid"
)
// NewIntegrationNotFoundError displays errors with the following description:
// Ensure your SCM integration exists and that it is correctly set up.
//
// Read more:
// - https://docs.snyk.io/scm-ide-and-ci-cd-workflow-and-integrations/snyk-scm-integrations
func NewIntegrationNotFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-integration-0001",
    Title:      "SCM integration not found",
    Description: "Ensure your SCM integration exists and that it is correctly set up.",
    StatusCode: 404,
    ErrorCode:  "SNYK-INTEGRATION-0001",
    Classification: "ACTIONABLE",
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
