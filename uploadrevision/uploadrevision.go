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

// Package uploadrevision contains errors related to the namespace UploadRevision
// of the Error Catalog.
package uploadrevision

import (
  "github.com/snyk/error-catalog-golang-public/snyk_errors"
  "github.com/google/uuid"
)
// NewUploadRevisionNotFoundError displays errors with the following description:
// The upload revision was not found. The upload revision may have expired or does not exist. Provide a valid upload revision identifier.
func NewUploadRevisionNotFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-upload-revision-0001",
    Title:      "Upload revision not found",
    Description: "The upload revision was not found. The upload revision may have expired or does not exist. Provide a valid upload revision identifier.",
    StatusCode: 404,
    ErrorCode:  "SNYK-UPLOAD-REVISION-0001",
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

// NewUploadRevisionSealedError displays errors with the following description:
// The upload revision cannot be modified after it has been sealed. Create a new upload revision and retry the operation.
func NewUploadRevisionSealedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-upload-revision-0002",
    Title:      "Upload revision is sealed",
    Description: "The upload revision cannot be modified after it has been sealed. Create a new upload revision and retry the operation.",
    StatusCode: 400,
    ErrorCode:  "SNYK-UPLOAD-REVISION-0002",
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

// NewFileTooLargeError displays errors with the following description:
// One of the uploaded files exceeds the maximum allowed size. Reduce the file size and try again, or contact support for assistance.
func NewFileTooLargeError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-upload-revision-0003",
    Title:      "File too large",
    Description: "One of the uploaded files exceeds the maximum allowed size. Reduce the file size and try again, or contact support for assistance.",
    StatusCode: 400,
    ErrorCode:  "SNYK-UPLOAD-REVISION-0003",
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

// NewTotalFilesSizeLimitExceededError displays errors with the following description:
// The total size of all files uploaded in a single request exceeds the maximum allowed limit. Reduce the total size of files in the request.
func NewTotalFilesSizeLimitExceededError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-upload-revision-0004",
    Title:      "Total files size limit exceeded",
    Description: "The total size of all files uploaded in a single request exceeds the maximum allowed limit. Reduce the total size of files in the request.",
    StatusCode: 400,
    ErrorCode:  "SNYK-UPLOAD-REVISION-0004",
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

// NewFileCountLimitExceededError displays errors with the following description:
// The request includes more files than the maximum allowed file count. Reduce the number of files in the request.
func NewFileCountLimitExceededError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-upload-revision-0005",
    Title:      "File count limit exceeded",
    Description: "The request includes more files than the maximum allowed file count. Reduce the number of files in the request.",
    StatusCode: 400,
    ErrorCode:  "SNYK-UPLOAD-REVISION-0005",
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

// NewFilePathTooLongError displays errors with the following description:
// One or more file paths exceed the maximum allowed length. Reduce the length of the file paths.
func NewFilePathTooLongError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-upload-revision-0006",
    Title:      "File path too long",
    Description: "One or more file paths exceed the maximum allowed length. Reduce the length of the file paths.",
    StatusCode: 400,
    ErrorCode:  "SNYK-UPLOAD-REVISION-0006",
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

// NewPopulateRequestLimitExceededError displays errors with the following description:
// The number of populate requests for this upload revision exceeds the maximum allowed limit. Create a new upload revision and upload a revision using fewer requests.
func NewPopulateRequestLimitExceededError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-upload-revision-0007",
    Title:      "Populate request limit exceeded",
    Description: "The number of populate requests for this upload revision exceeds the maximum allowed limit. Create a new upload revision and upload a revision using fewer requests.",
    StatusCode: 400,
    ErrorCode:  "SNYK-UPLOAD-REVISION-0007",
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

// NewTotalUploadRevisionFileCountLimitExceededError displays errors with the following description:
// The number of files exceeds the limit for a single upload revision. Consider splitting the files across multiple upload revisions.
func NewTotalUploadRevisionFileCountLimitExceededError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-upload-revision-0008",
    Title:      "Upload revision file count limit exceeded",
    Description: "The number of files exceeds the limit for a single upload revision. Consider splitting the files across multiple upload revisions.",
    StatusCode: 400,
    ErrorCode:  "SNYK-UPLOAD-REVISION-0008",
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

// NewTotalUploadRevisionSizeLimitExceededError displays errors with the following description:
// The total size of all files uploaded across the entire upload revision exceeds the maximum allowed limit. Consider splitting the files across multiple upload revisions.
func NewTotalUploadRevisionSizeLimitExceededError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-upload-revision-0009",
    Title:      "Upload revision size limit exceeded",
    Description: "The total size of all files uploaded across the entire upload revision exceeds the maximum allowed limit. Consider splitting the files across multiple upload revisions.",
    StatusCode: 400,
    ErrorCode:  "SNYK-UPLOAD-REVISION-0009",
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

// NewUploadRevisionIdMismatchError displays errors with the following description:
// The upload revision identifier in the request body does not match the upload revision identifier in the request path. Ensure both identifiers match and retry the request.
func NewUploadRevisionIdMismatchError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-upload-revision-0010",
    Title:      "Upload revision identifier mismatch",
    Description: "The upload revision identifier in the request body does not match the upload revision identifier in the request path. Ensure both identifiers match and retry the request.",
    StatusCode: 400,
    ErrorCode:  "SNYK-UPLOAD-REVISION-0010",
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

// NewMultipartFieldNameMissingError displays errors with the following description:
// The multipart request has a missing or empty name parameter in the Content-Disposition header. Ensure each part includes a valid field name.
func NewMultipartFieldNameMissingError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-upload-revision-0011",
    Title:      "Missing field name in multipart request",
    Description: "The multipart request has a missing or empty name parameter in the Content-Disposition header. Ensure each part includes a valid field name.",
    StatusCode: 400,
    ErrorCode:  "SNYK-UPLOAD-REVISION-0011",
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

// NewUploadRevisionUnsealedError displays errors with the following description:
// The upload revision cannot be consumed before it is sealed. Seal the revision or create a new one and retry the operation.
func NewUploadRevisionUnsealedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-upload-revision-0012",
    Title:      "Upload revision is unsealed",
    Description: "The upload revision cannot be consumed before it is sealed. Seal the revision or create a new one and retry the operation.",
    StatusCode: 400,
    ErrorCode:  "SNYK-UPLOAD-REVISION-0012",
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
