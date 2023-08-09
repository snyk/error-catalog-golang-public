
// Package isolatedbuilds contains errors related to the namespace IsolatedBuilds
// of the Error Catalog.
package isolatedbuilds

import (
  "github.com/snyk/error-catalog-golang-public/snyk_errors"
  "github.com/google/uuid"
)

// NewInvalidRequestError displays errors with the following description:
// The provided request payload is not valid for the selected ecosystem. Please review the API documentation.
//
// Read more:
// - https://apidocs.snyk.io/
func NewInvalidRequestError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-8001",
    Title:      "Invalid request",
    StatusCode: 400,
    ErrorCode:  "SNYK-OS-8001",
    Links: []string{
      "https://apidocs.snyk.io/",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewBuildEnvironmentNotFoundError displays errors with the following description:
// The build environment for the provided context could not be found. Please ensure you have created the build environment first.
func NewBuildEnvironmentNotFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-8002",
    Title:      "Build environment not found",
    StatusCode: 404,
    ErrorCode:  "SNYK-OS-8002",
    Links: []string{},
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

