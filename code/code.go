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
// A SAST Rule extension with the same type and attributes already exists
// for the given Group.
// 
// Either modify the existing SAST Rule extension or create a new
// SAST Rule extension with a different type or attributes.
func NewRuleExtensionAlreadyExistsForGroupError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-code-0007",
    Title:      "SAST Rule extension already exists for the Group",
    Description: "A SAST Rule extension with the same type and attributes already exists\nfor the given Group.\n\nEither modify the existing SAST Rule extension or create a new\nSAST Rule extension with a different type or attributes.",
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
// Each Org relationship to a Snyk SAST Rule extension must be unique.
// 
// Make sure each Org in relationships has a different ID.
func NewOrgRelationshipsMustBeUniqueError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-code-0008",
    Title:      "Organization relationships must be unique",
    Description: "Each Org relationship to a Snyk SAST Rule extension must be unique.\n\nMake sure each Org in relationships has a different ID.",
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
// You cannot associate a Snyk SAST Rule extension to any other Group.
// 
// Make sure the Group ID under relationships matches the Group ID in the request path.
func NewGroupRelationshipMustBeForAdminGroupError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-code-0009",
    Title:      "Group relationship must match the Group in the requested URL",
    Description: "You cannot associate a Snyk SAST Rule extension to any other Group.\n\nMake sure the Group ID under relationships matches the Group ID in the request path.",
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
// You cannot use the SAST Rule extensions feature with an Org outside of the administrating Group.
// 
// Make sure each Org in the request is within the requested Group.
func NewOrgOutsideAdminGroupError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-code-0010",
    Title:      "Organization outside of the administrating Group",
    Description: "You cannot use the SAST Rule extensions feature with an Org outside of the administrating Group.\n\nMake sure each Org in the request is within the requested Group.",
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
// You have hit the maximum number of published Snyk SAST Rule extensions allowed for a Group.
// 
// To create a new SAST Rule extension you will have to remove an existing one.
func NewRuleExtensionsLimitReachedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-code-0011",
    Title:      "SAST Rule extension limit reached",
    Description: "You have hit the maximum number of published Snyk SAST Rule extensions allowed for a Group.\n\nTo create a new SAST Rule extension you will have to remove an existing one.",
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

// NewTestRuleExtensionAlreadyPublishedForGroupError displays errors with the following description:
// The Rule Extension under test conflicts with an already published SAST Rule Extension.
// 
// A test cannot be performed if a SAST Rule Extension with the same fully qualified name
// and type is already published for the Group. Either delete the already published SAST Rule Extension
// or perform a test with a different fully qualified name or type.
func NewTestRuleExtensionAlreadyPublishedForGroupError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-code-0012",
    Title:      "SAST Rule Extension already published for the Group",
    Description: "The Rule Extension under test conflicts with an already published SAST Rule Extension.\n\nA test cannot be performed if a SAST Rule Extension with the same fully qualified name\nand type is already published for the Group. Either delete the already published SAST Rule Extension\nor perform a test with a different fully qualified name or type.",
    StatusCode: 400,
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

// NewTestIDNotFoundError displays errors with the following description:
// The requested Test ID for testing SAST Rule Extension was not found.
// 
// Make sure to provide a valid Test ID.
func NewTestIDNotFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-code-0013",
    Title:      "Requested test ID not found",
    Description: "The requested Test ID for testing SAST Rule Extension was not found.\n\nMake sure to provide a valid Test ID.",
    StatusCode: 404,
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

// NewTestResultsExpiredError displays errors with the following description:
// The results for testing SAST Rule Extensions have expired and are no longer available.
// 
// Please trigger a new test.
func NewTestResultsExpiredError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-code-0014",
    Title:      "Test results have expired",
    Description: "The results for testing SAST Rule Extensions have expired and are no longer available.\n\nPlease trigger a new test.",
    StatusCode: 404,
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

// NewTestIDNotAssociatedWithGroupError displays errors with the following description:
// There is no SAST Rule Extension impact test associated with the given Test ID for the group.
// 
// Please verify you are using the correct Group ID associated with the test.
func NewTestIDNotAssociatedWithGroupError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-code-0015",
    Title:      "Test ID not associated with group",
    Description: "There is no SAST Rule Extension impact test associated with the given Test ID for the group.\n\nPlease verify you are using the correct Group ID associated with the test.",
    StatusCode: 404,
    ErrorCode:  "SNYK-CODE-0015",
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
