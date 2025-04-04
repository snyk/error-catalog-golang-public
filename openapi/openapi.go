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

// Package openapi contains errors related to the namespace OpenAPI
// of the Error Catalog.
package openapi

import (
  "github.com/snyk/error-catalog-golang-public/snyk_errors"
  "github.com/google/uuid"
)

// NewBadRequestError displays errors with the following description:
// The server cannot process the request due to invalid or corrupt data. Review the request, then try again.
//
// Read more:
// - https://docs.snyk.io/snyk-api-info/getting-started-using-snyk-rest-api 
func NewBadRequestError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-openapi-0001",
    Title:      "Bad request",
    Description: "The server cannot process the request due to invalid or corrupt data. Review the request, then try again.",
    StatusCode: 400,
    ErrorCode:  "SNYK-OPENAPI-0001",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/snyk-api-info/getting-started-using-snyk-rest-api ",
    },
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewForbiddenError displays errors with the following description:
// Access to the requested resource is forbidden. Review the request, then try again.
func NewForbiddenError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-openapi-0002",
    Title:      "Forbidden",
    Description: "Access to the requested resource is forbidden. Review the request, then try again.",
    StatusCode: 403,
    ErrorCode:  "SNYK-OPENAPI-0002",
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

// NewNotAcceptableError displays errors with the following description:
// The server cannot provide a response that matches the provided accept headers. Review the request, then try again.
func NewNotAcceptableError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-openapi-0003",
    Title:      "Not acceptable",
    Description: "The server cannot provide a response that matches the provided accept headers. Review the request, then try again.",
    StatusCode: 406,
    ErrorCode:  "SNYK-OPENAPI-0003",
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

// NewNotFoundError displays errors with the following description:
// The server cannot find the requested resource. Review the request, then try again.
func NewNotFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-openapi-0004",
    Title:      "Not found",
    Description: "The server cannot find the requested resource. Review the request, then try again.",
    StatusCode: 404,
    ErrorCode:  "SNYK-OPENAPI-0004",
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

// NewMethodNotAllowedError displays errors with the following description:
// The target endpoint does not support your request method. Review the request, then try again.
func NewMethodNotAllowedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-openapi-0005",
    Title:      "Method not allowed",
    Description: "The target endpoint does not support your request method. Review the request, then try again.",
    StatusCode: 405,
    ErrorCode:  "SNYK-OPENAPI-0005",
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

// NewRequestEntityTooLargeError displays errors with the following description:
// The request entity exceeds server limitations. Reduce the size of the request entity, then try again.
func NewRequestEntityTooLargeError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-openapi-0006",
    Title:      "Request entity too large",
    Description: "The request entity exceeds server limitations. Reduce the size of the request entity, then try again.",
    StatusCode: 413,
    ErrorCode:  "SNYK-OPENAPI-0006",
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

// NewUnauthorizedError displays errors with the following description:
// The request lacks authentication credentials for the requested resource. Ensure you are sending valid credentials, then try again.
//
// Read more:
// - https://docs.snyk.io/snyk-api-info/authentication-for-api
func NewUnauthorizedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-openapi-0007",
    Title:      "Unauthorized",
    Description: "The request lacks authentication credentials for the requested resource. Ensure you are sending valid credentials, then try again.",
    StatusCode: 401,
    ErrorCode:  "SNYK-OPENAPI-0007",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/snyk-api-info/authentication-for-api",
    },
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewUnsupportedMediaTypeError displays errors with the following description:
// The media format of the request is not supported. Change media format, then try again.
func NewUnsupportedMediaTypeError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-openapi-0008",
    Title:      "Unsupported media type",
    Description: "The media format of the request is not supported. Change media format, then try again.",
    StatusCode: 415,
    ErrorCode:  "SNYK-OPENAPI-0008",
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

// NewConflictError displays errors with the following description:
// The request could not be completed due to a conflict with the current state of the target resource. Review the request, then try again.
func NewConflictError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-openapi-0009",
    Title:      "Conflict",
    Description: "The request could not be completed due to a conflict with the current state of the target resource. Review the request, then try again.",
    StatusCode: 409,
    ErrorCode:  "SNYK-OPENAPI-0009",
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

