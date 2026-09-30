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

// Package supplychain contains errors related to the namespace SupplyChain
// of the Error Catalog.
package supplychain

import (
  "github.com/snyk/error-catalog-golang-public/snyk_errors"
  "github.com/google/uuid"
)
// NewConfigurationDriftError displays errors with the following description:
// Package managers on this machine no longer point at the Registry Proxy, so the packages they install may not be screened. Run `snyk sc install` again.
func NewConfigurationDriftError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-sc-regproxy-0001",
    Title:      "Package manager configuration has drifted",
    Description: "Package managers on this machine no longer point at the Registry Proxy, so the packages they install may not be screened. Run `snyk sc install` again.",
    StatusCode: 0,
    ErrorCode:  "SNYK-SC-REGPROXY-0001",
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

// NewConflictingConfigurationError displays errors with the following description:
// Existing configuration sends package managers to a different registry, bypassing the Registry Proxy. Remove the setting named in the output, then run `snyk sc install` again.
func NewConflictingConfigurationError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-sc-regproxy-0002",
    Title:      "Conflicting package manager configuration",
    Description: "Existing configuration sends package managers to a different registry, bypassing the Registry Proxy. Remove the setting named in the output, then run `snyk sc install` again.",
    StatusCode: 0,
    ErrorCode:  "SNYK-SC-REGPROXY-0002",
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

// NewRegistryCredentialRejectedError displays errors with the following description:
// The Registry Proxy rejected the stored token, so package installs through it fail. Create a new token and run `snyk sc install` again.
func NewRegistryCredentialRejectedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-sc-regproxy-0003",
    Title:      "Registry Credential rejected",
    Description: "The Registry Proxy rejected the stored token, so package installs through it fail. Create a new token and run `snyk sc install` again.",
    StatusCode: 0,
    ErrorCode:  "SNYK-SC-REGPROXY-0003",
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
