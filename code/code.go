/*
 * © 2023 Snyk Limited
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
// This error occurs when the analysis target has a supported file count which exceeds current system limits.
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
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-code-0001",
    Title:      "Analysis file count limit exceeded",
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
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-code-0002",
    Title:      "Analysis result size limit exceeded",
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
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-code-0003",
    Title:      "Analysis target size limit exceeded",
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

