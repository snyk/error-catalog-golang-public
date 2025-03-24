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

// Package code contains errors related to the namespace Code
// of the Error Catalog.
package code

import (
  "github.com/snyk/error-catalog-golang-public/snyk_errors"
  "github.com/google/uuid"
)

// NewAnalysisFileCountLimitExceededError displays errors with the following description:
// This error occurs when the analysis target has a supported file count that exceeds current system limits.
// 
// To reduce the file count, use a `.snyk` file to ignore specified directories or files. Alternatively, use the Snyk CLI to analyze individual subdirectories separately.
//
// Read more:
// - https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/supported-languages-frameworks-and-feature-availability-overview#code-analysis-snyk-code
// - https://docs.snyk.io/scan-applications/start-scanning-using-the-cli-web-ui-or-api/snyk-code-and-your-repositories/excluding-directories-and-files-from-the-import-process
// - https://docs.snyk.io/snyk-cli/using-snyk-code-from-the-cli
func NewAnalysisFileCountLimitExceededError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-code-0001",
    Title:      "Analysis file count limit exceeded",
    Description: "This error occurs when the analysis target has a supported file count that exceeds current system limits.\n\nTo reduce the file count, use a `.snyk` file to ignore specified directories or files. Alternatively, use the Snyk CLI to analyze individual subdirectories separately.",
    StatusCode: 422,
    ErrorCode:  "SNYK-CODE-0001",
    Classification: "UNSUPPORTED",
    Links: []string{
      "https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/supported-languages-frameworks-and-feature-availability-overview#code-analysis-snyk-code",
      "https://docs.snyk.io/scan-applications/start-scanning-using-the-cli-web-ui-or-api/snyk-code-and-your-repositories/excluding-directories-and-files-from-the-import-process",
      "https://docs.snyk.io/snyk-cli/using-snyk-code-from-the-cli",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewAnalysisResultSizeLimitExceededError displays errors with the following description:
// This error occurs when the analysis target generates a result with a byte size that exceeds current system limits.
// 
// To reduce the overall result size, use a `.snyk` file to ignore specified directories or files. Alternatively, use the Snyk CLI to analyze individual subdirectories separately.
//
// Read more:
// - https://docs.snyk.io/scan-applications/start-scanning-using-the-cli-web-ui-or-api/snyk-code-and-your-repositories/excluding-directories-and-files-from-the-import-process
// - https://docs.snyk.io/snyk-cli/using-snyk-code-from-the-cli
func NewAnalysisResultSizeLimitExceededError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-code-0002",
    Title:      "Analysis result size limit exceeded",
    Description: "This error occurs when the analysis target generates a result with a byte size that exceeds current system limits.\n\nTo reduce the overall result size, use a `.snyk` file to ignore specified directories or files. Alternatively, use the Snyk CLI to analyze individual subdirectories separately.",
    StatusCode: 422,
    ErrorCode:  "SNYK-CODE-0002",
    Classification: "UNSUPPORTED",
    Links: []string{
      "https://docs.snyk.io/scan-applications/start-scanning-using-the-cli-web-ui-or-api/snyk-code-and-your-repositories/excluding-directories-and-files-from-the-import-process",
      "https://docs.snyk.io/snyk-cli/using-snyk-code-from-the-cli",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewAnalysisTargetSizeLimitExceededError displays errors with the following description:
// This error occurs when the analysis target byte size exceeds current system limits.
// 
// To reduce the overall result size, use a `.snyk` file to ignore specified directories or files. Alternatively, use the Snyk CLI to analyze individual subdirectories separately.
//
// Read more:
// - https://docs.snyk.io/snyk-cli/using-snyk-code-from-the-cli
func NewAnalysisTargetSizeLimitExceededError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-code-0003",
    Title:      "Analysis target size limit exceeded",
    Description: "This error occurs when the analysis target byte size exceeds current system limits.\n\nTo reduce the overall result size, use a `.snyk` file to ignore specified directories or files. Alternatively, use the Snyk CLI to analyze individual subdirectories separately.",
    StatusCode: 422,
    ErrorCode:  "SNYK-CODE-0003",
    Classification: "UNSUPPORTED",
    Links: []string{
      "https://docs.snyk.io/snyk-cli/using-snyk-code-from-the-cli",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewAnalysisFileNameLengthLimitExceededError displays errors with the following description:
// This error occurs when the analysis target has a file name length that exceeds 255 bytes.
// 
// To be able to scan the analysis target, rename the file to a name that is 255 bytes or less.
//
// Read more:
// - https://docs.snyk.io/scan-with-snyk/supported-languages-and-frameworks/introduction-to-snyk-supported-languages-and-frameworks#filename-length-limitation
func NewAnalysisFileNameLengthLimitExceededError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-code-0004",
    Title:      "Analysis target includes a file with a name longer than 255 bytes",
    Description: "This error occurs when the analysis target has a file name length that exceeds 255 bytes.\n\nTo be able to scan the analysis target, rename the file to a name that is 255 bytes or less.",
    StatusCode: 422,
    ErrorCode:  "SNYK-CODE-0004",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/scan-with-snyk/supported-languages-and-frameworks/introduction-to-snyk-supported-languages-and-frameworks#filename-length-limitation",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewFeatureIsNotEnabledError displays errors with the following description:
// This error occurs when Snyk Code is not enabled for the current Organization. Activate Snyk Code and try again..
//
// Read more:
// - https://docs.snyk.io/scan-using-snyk/snyk-code/configure-snyk-code#enable-snyk-code-in-snyk-web-ui
func NewFeatureIsNotEnabledError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-code-0005",
    Title:      "Snyk Code is not enabled",
    Description: "This error occurs when Snyk Code is not enabled for the current Organization. Activate Snyk Code and try again..",
    StatusCode: 403,
    ErrorCode:  "SNYK-CODE-0005",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/scan-using-snyk/snyk-code/configure-snyk-code#enable-snyk-code-in-snyk-web-ui",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewUnsupportedProjectError displays errors with the following description:
// Snyk was unable to find supported files.
//
// Read more:
// - https://docs.snyk.io/getting-started/supported-languages-frameworks-and-feature-availability-overview#code-analysis-snyk-code
func NewUnsupportedProjectError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-code-0006",
    Title:      "Project not supported",
    Description: "Snyk was unable to find supported files.",
    StatusCode: 422,
    ErrorCode:  "SNYK-CODE-0006",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/getting-started/supported-languages-frameworks-and-feature-availability-overview#code-analysis-snyk-code",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewRuleExtensionAlreadyExistsForGroupError displays errors with the following description:
// A Sast Rule extension with the same type and attributes already exists
// for the given Group.
// 
// Either modify the existing Sast Rule extension or create a new
// Sast Rule extension with a different type or attributes.
func NewRuleExtensionAlreadyExistsForGroupError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-code-0007",
    Title:      "Sast Rule extension already exists for the Group",
    Description: "A Sast Rule extension with the same type and attributes already exists\nfor the given Group.\n\nEither modify the existing Sast Rule extension or create a new\nSast Rule extension with a different type or attributes.",
    StatusCode: 409,
    ErrorCode:  "SNYK-CODE-0007",
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

// NewOrgRelationshipsMustBeUniqueError displays errors with the following description:
// Each Org relationship to a Snyk Sast Rule extension must be unique.
// 
// Make sure each Org in relationships has a different ID.
func NewOrgRelationshipsMustBeUniqueError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-code-0008",
    Title:      "Organization relationships must be unique",
    Description: "Each Org relationship to a Snyk Sast Rule extension must be unique.\n\nMake sure each Org in relationships has a different ID.",
    StatusCode: 400,
    ErrorCode:  "SNYK-CODE-0008",
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

// NewGroupRelationshipMustBeForAdminGroupError displays errors with the following description:
// You cannot associate a Snyk Sast Rule extension to any other Group.
// 
// Make sure the Group ID under relationships matches the Group ID in the request path.
func NewGroupRelationshipMustBeForAdminGroupError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-code-0009",
    Title:      "Group relationship must match the Group in the requested URL",
    Description: "You cannot associate a Snyk Sast Rule extension to any other Group.\n\nMake sure the Group ID under relationships matches the Group ID in the request path.",
    StatusCode: 400,
    ErrorCode:  "SNYK-CODE-0009",
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

// NewOrgOutsideAdminGroupError displays errors with the following description:
// You cannot associate a Snyk Sast Rule extension to an Org outside of the administrating Group.
// 
// Make sure each Org under relationships is within the Group in the requested URL.
func NewOrgOutsideAdminGroupError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-code-0010",
    Title:      "Organization outside of the administrating Group",
    Description: "You cannot associate a Snyk Sast Rule extension to an Org outside of the administrating Group.\n\nMake sure each Org under relationships is within the Group in the requested URL.",
    StatusCode: 400,
    ErrorCode:  "SNYK-CODE-0010",
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

// NewRuleExtensionsLimitReachedError displays errors with the following description:
// You have hit the maximum number of published Snyk Sast Rule extensions allowed for a Group.
// 
// To create a new Sast Rule extension you will have to remove an existing one.
func NewRuleExtensionsLimitReachedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-code-0011",
    Title:      "Sast Rule extension limit reached",
    Description: "You have hit the maximum number of published Snyk Sast Rule extensions allowed for a Group.\n\nTo create a new Sast Rule extension you will have to remove an existing one.",
    StatusCode: 400,
    ErrorCode:  "SNYK-CODE-0011",
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

// NewUnsupportedOrgError displays errors with the following description:
// The Snyk Sast Rule extensions feature is not enabled for this Organization ID.
// 
// Please reach out to the account team to get access.
func NewUnsupportedOrgError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-code-0012",
    Title:      "Sast Rule extensions feature is not enabled for this Organization ID",
    Description: "The Snyk Sast Rule extensions feature is not enabled for this Organization ID.\n\nPlease reach out to the account team to get access.",
    StatusCode: 421,
    ErrorCode:  "SNYK-CODE-0012",
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

// NewRuleExtensionsDecryptionTimeoutError displays errors with the following description:
// Decrypting the Sast Rule extensions for the requested Organization ID process timed out.
// You have created too many published rules for the Sast Rule extensions beta. 
// 
// Please remove one or more to try again.
// If the issue persists, please open a customer support ticket.
func NewRuleExtensionsDecryptionTimeoutError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-code-0013",
    Title:      "Decryption time out for the requested Sast Rule extensions",
    Description: "Decrypting the Sast Rule extensions for the requested Organization ID process timed out.\nYou have created too many published rules for the Sast Rule extensions beta. \n\nPlease remove one or more to try again.\nIf the issue persists, please open a customer support ticket.",
    StatusCode: 408,
    ErrorCode:  "SNYK-CODE-0013",
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

// NewRulesRelationshipsMustBeUniqueError displays errors with the following description:
// Each Rule relationship to a Snyk Sast Rule extension must be unique.
// 
// Make sure each Rule in relationships has a different name.
func NewRulesRelationshipsMustBeUniqueError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-code-0014",
    Title:      "Rules relationships must be unique",
    Description: "Each Rule relationship to a Snyk Sast Rule extension must be unique.\n\nMake sure each Rule in relationships has a different name.",
    StatusCode: 400,
    ErrorCode:  "SNYK-CODE-0014",
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

