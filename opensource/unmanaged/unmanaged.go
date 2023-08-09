
// Package unmanaged contains errors related to the namespace OpenSourceUnmanaged
// of the Error Catalog.
package unmanaged

import (
  "github.com/snyk/error-catalog-golang-public/snyk_errors"
  "github.com/google/uuid"
)

// NewMavenSearchServiceUnavailableError displays errors with the following description:
// The upstream Maven search service is not available.
//
// Read more:
// - https://search.maven.org
// - https://status.maven.org
func NewMavenSearchServiceUnavailableError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-osjvm-001",
    Title:      "Maven search service unavailable",
    StatusCode: 503,
    ErrorCode:  "SNYK-OSJVM-001",
    Links: []string{
      "https://search.maven.org",
      "https://status.maven.org",
    },
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewSha1NotFoundError displays errors with the following description:
// Unable to find the coordinates for the provided SHA1.
//
// Read more:
// - https://docs.snyk.io/snyk-cli/test-for-vulnerabilities/scan-all-unmanaged-jar-files
func NewSha1NotFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-osjvm-002",
    Title:      "SHA1 not found",
    StatusCode: 404,
    ErrorCode:  "SNYK-OSJVM-002",
    Links: []string{
      "https://docs.snyk.io/snyk-cli/test-for-vulnerabilities/scan-all-unmanaged-jar-files",
    },
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

