/*
 * © 2024 Snyk Limited
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-0001",
    Title:      "Service temporarily throttled",
    Description: "The request rate limit has been exceeded. Wait a few minutes, then try again.",
    StatusCode: 429,
    ErrorCode:  "SNYK-0001",
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

// NewNotImplementedError displays errors with the following description:
// The server doesn't recognize the request method, or it cannot fulfill it. Review the request and try again.
//
// Read more:
// - https://docs.snyk.io/snyk-api-info
func NewNotImplementedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-0002",
    Title:      "Server error response",
    Description: "The server doesn't recognize the request method, or it cannot fulfill it. Review the request and try again.",
    StatusCode: 501,
    ErrorCode:  "SNYK-0002",
    Classification: "UNSUPPORTED",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-0003",
    Title:      "Client request cannot be processed",
    Description: "The server cannot process the request due to a client error, such as malformed request syntax, size too large, invalid request message framing, or deceptive request routing. Review the request and try again.",
    StatusCode: 400,
    ErrorCode:  "SNYK-0003",
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

// NewTimeoutError displays errors with the following description:
// The server timed out during the request. Check Snyk status, then try again.
//
// Read more:
// - https://status.snyk.io/
func NewTimeoutError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-0004",
    Title:      "Server communication error",
    Description: "The server timed out during the request. Check Snyk status, then try again.",
    StatusCode: 504,
    ErrorCode:  "SNYK-0004",
    Classification: "UNEXPECTED",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-0005",
    Title:      "Authentication error",
    Description: "Authentication credentials not recognized, or user access is not provisioned. Revise credentials and try again, or request access from your Snyk administrator.",
    StatusCode: 401,
    ErrorCode:  "SNYK-0005",
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

// NewTestLimitReachedError displays errors with the following description:
// You have reached the maximum number of tests in your Snyk plan. This causes Snyk tests on PRs and CLI to fail. Deactivate Snyk Test on your Project or upgrade your Snyk plan.
//
// Read more:
// - https://support.snyk.io/hc/en-us/articles/4409805538833-Rate-limit-hit-while-testing-the-project
// - https://docs.snyk.io/scan-using-snyk/working-with-snyk-in-your-environment/what-counts-as-a-test
// - https://support.snyk.io/hc/en-us/articles/360001945297-Snyk-Test-of-PR-failing-due-to-test-limit
func NewTestLimitReachedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-0006",
    Title:      "Test limit reached",
    Description: "You have reached the maximum number of tests in your Snyk plan. This causes Snyk tests on PRs and CLI to fail. Deactivate Snyk Test on your Project or upgrade your Snyk plan.",
    StatusCode: 429,
    ErrorCode:  "SNYK-0006",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://support.snyk.io/hc/en-us/articles/4409805538833-Rate-limit-hit-while-testing-the-project",
      "https://docs.snyk.io/scan-using-snyk/working-with-snyk-in-your-environment/what-counts-as-a-test",
      "https://support.snyk.io/hc/en-us/articles/360001945297-Snyk-Test-of-PR-failing-due-to-test-limit",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewTagsForOrganizationWithoutGroupError displays errors with the following description:
// This error occures when trying to add tags to Organizations that that are part of a Group.
// 
// Verify with your Group Admin if the Organization should be in a Group.
// 
// If you have more than one Organization, you can set the Organization with which new Projects should be associated by running `snyk config set org=ORG_ID`.
// 
// If you want to override this global configuration for individual runs of snyk monitor, `run snyk test --org=ORG_ID` or `snyk monitor --org=ORG_ID`.
//
// Read more:
// - https://docs.snyk.io/snyk-admin/snyk-projects/project-tags
func NewTagsForOrganizationWithoutGroupError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-0007",
    Title:      "Organization is not part of a group",
    Description: "This error occures when trying to add tags to Organizations that that are part of a Group.\n\nVerify with your Group Admin if the Organization should be in a Group.\n\nIf you have more than one Organization, you can set the Organization with which new Projects should be associated by running `snyk config set org=ORG_ID`.\n\nIf you want to override this global configuration for individual runs of snyk monitor, `run snyk test --org=ORG_ID` or `snyk monitor --org=ORG_ID`.",
    StatusCode: 422,
    ErrorCode:  "SNYK-0007",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/snyk-admin/snyk-projects/project-tags",
    },
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-9999",
    Title:      "Request not fulfilled due to server error",
    Description: "The server cannot process the request due to an unexpected error. Check Snyk status, then try again.",
    StatusCode: 500,
    ErrorCode:  "SNYK-9999",
    Classification: "UNEXPECTED",
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

