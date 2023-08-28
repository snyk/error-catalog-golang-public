
/*
 * © 2023 Snyk Limited
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package snyk contains errors related to the namespace Snyk
// of the Error Catalog.
package snyk

import (
  "github.com/snyk/error-catalog-golang-public/snyk_errors"
  "github.com/google/uuid"
)

// NewTooManyRequestsError displays errors with the following description:
// The request rate limit has been exceeded. Wait a few minutes, then try again.
func NewTooManyRequestsError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-0001",
    Title:      "Service temporarily throttled",
    StatusCode: 429,
    ErrorCode:  "SNYK-0001",
    Links: []string{},
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewNotImplementedError displays errors with the following description:
// The server doesn’t recognize the request method, or it cannot fulfill it. Review the request and try again.
//
// Read more:
// - https://docs.snyk.io/snyk-api-info
func NewNotImplementedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-0002",
    Title:      "Server error response",
    StatusCode: 501,
    ErrorCode:  "SNYK-0002",
    Links: []string{
      "https://docs.snyk.io/snyk-api-info",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewBadRequestError displays errors with the following description:
// The server cannot process the request due to a client error, such as malformed request syntax, size too large, invalid request message framing, or deceptive request routing. Review the request and try again.
func NewBadRequestError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-0003",
    Title:      "Client request cannot be processed",
    StatusCode: 400,
    ErrorCode:  "SNYK-0003",
    Links: []string{},
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewTimeoutError displays errors with the following description:
// The server timed out during the request. Check Snyk status, then try again.
//
// Read more:
// - https://status.snyk.io/
func NewTimeoutError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-0004",
    Title:      "Server communication error",
    StatusCode: 504,
    ErrorCode:  "SNYK-0004",
    Links: []string{
      "https://status.snyk.io/",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewUnauthorisedError displays errors with the following description:
// Authentication credentials not recognized, or user access is not provisioned. Revise credentials and try again, or request access from your Snyk administrator.
func NewUnauthorisedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-0005",
    Title:      "Authentication error",
    StatusCode: 401,
    ErrorCode:  "SNYK-0005",
    Links: []string{},
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewServerError displays errors with the following description:
// The server cannot process the request due to an unexpected error. Check Snyk status, then try again.
//
// Read more:
// - https://status.snyk.io/
func NewServerError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-9999",
    Title:      "Request not fulfilled due to server error ",
    StatusCode: 500,
    ErrorCode:  "SNYK-9999",
    Links: []string{
      "https://status.snyk.io/",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

