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

// Package policies contains errors related to the namespace Policies
// of the Error Catalog.
package policies

import (
  "github.com/snyk/error-catalog-golang-public/snyk_errors"
  "github.com/google/uuid"
)

// NewInvalidPolicyApplyError displays errors with the following description:
// Snyk could not apply a policy whilst executing a test because the configuration for the policy was invalid.
// You may be able to fix the policy and try again.
//
// Read more:
// - https://docs.snyk.io/manage-risk/policies
func NewInvalidPolicyApplyError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-policy-0001",
    Title:      "Unable to apply a policy with an invalid configuration",
    Description: "Snyk could not apply a policy whilst executing a test because the configuration for the policy was invalid.\nYou may be able to fix the policy and try again.",
    StatusCode: 422,
    ErrorCode:  "SNYK-POLICY-0001",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/manage-risk/policies",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

