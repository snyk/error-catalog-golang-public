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

// Package prchecks contains errors related to the namespace PRChecks
// of the Error Catalog.
package prchecks

import (
  "github.com/snyk/error-catalog-golang-public/snyk_errors"
  "github.com/google/uuid"
)

// NewFailedToReadManifestError displays errors with the following description:
// Snyk failed to read 1 or more manifest files.
// Sometimes things go wrong: a flaky connection, 3rd party services go down and Snyk is unable to read the files needed in order to test your project. 
// 
// If this happens, you could try:
// 
// - Opening and re-opening your Pull Request / Merge Request, to kick off a new test
// - Removing and re-adding the repo to Snyk
// 
// Ultimately, you should contact support@snyk.io if the issue persists
//
// Read more:
// - https://support.snyk.io/hc/en-us/articles/360000910517-Failed-to-read-manifest-file
func NewFailedToReadManifestError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-pr-check-0001",
    Title:      "Error reading manifest",
    StatusCode: 500,
    ErrorCode:  "SNYK-PR-CHECK-0001",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://support.snyk.io/hc/en-us/articles/360000910517-Failed-to-read-manifest-file",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewManifestNotFoundError displays errors with the following description:
// Snyk uses your project manifest file to analyze your projects for vulnerabilities. When you import a project for monitoring, Snyk scans the project to locate the manifest file and then remembers where that file is. 
// When a project manifest file is moved or deleted, we still try to look for in it in the last known location in order to run tests on commit statuses. If we can't find the file, this error can occur.
// 
// If this happens, you could try the following:
// 1. Delete the matching project from your account in the Snyk app (UI or CLI).
// 2. Now import the same project from scratch.
// 
// As during the original import, Snyk scans the project and locates the manifest file.
//
// Read more:
// - https://support.snyk.io/hc/en-us/articles/360000910537-Manifest-not-found
func NewManifestNotFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-pr-check-0002",
    Title:      "Manifest not found",
    StatusCode: 404,
    ErrorCode:  "SNYK-PR-CHECK-0002",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://support.snyk.io/hc/en-us/articles/360000910537-Manifest-not-found",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewThirdPartyRateLimitError displays errors with the following description:
// Snyk makes requests to your SCM when testing a project, in order to analyze your projects for vulnerabilities. If we need to make a lot of requests in a short time period, we may encounter third party rate limits, and this error can occur.
// 
// If you receive any of these errors, try re-running the tests, by closing and reopening the pull request.
func NewThirdPartyRateLimitError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-pr-check-0003",
    Title:      "Rate limit hit while testing project",
    StatusCode: 429,
    ErrorCode:  "SNYK-PR-CHECK-0003",
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

// NewOutOfSyncError displays errors with the following description:
// Sometimes a project may become out of sync between the lockfile and the manifest file. This might happen if the package.json is modified or updated but the lockfile is not. 
// 
// This can be resolved by ensuring the lockfile and manifest file are correctly synced, by executing npm install or yarn install.
// 
// In some cases, it may be necessary to delete the node_modules folder and the package-lock.json and run npm install again to force a full reinstall. 
//
// Read more:
// - https://support.snyk.io/hc/en-us/articles/360000912457-Out-of-sync-manifest-lockfile-in-the-project
func NewOutOfSyncError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-pr-check-0004",
    Title:      "Out of Sync Error",
    StatusCode: 422,
    ErrorCode:  "SNYK-PR-CHECK-0004",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://support.snyk.io/hc/en-us/articles/360000912457-Out-of-sync-manifest-lockfile-in-the-project",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewFailedDeterminingProjectTargetError displays errors with the following description:
// An internal error occurred, whereby Snyk was unable to determine the correct target for a given project in your PR Check.
// 
// If you receive this error, try re-running the tests, by closing and reopening the pull request.
// 
// Ultimately, you should contact support@snyk.io if the issue persists.
func NewFailedDeterminingProjectTargetError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-pr-check-0005",
    Title:      "Failed determining project target",
    StatusCode: 500,
    ErrorCode:  "SNYK-PR-CHECK-0005",
    Classification: "UNEXPECTED",
    Links: []string{},
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewFailedToCompleteTestError displays errors with the following description:
// A "Failed to complete testing check status" appears in your commit checks when an unknown error occurs while Snyk was trying to test your projects for vulnerabilities or license issues.
// 
// If you receive this error, try re-running the tests, by closing and reopening the pull request.
// 
// Ultimately, you should contact support@snyk.io if the issue persists.
//
// Read more:
// - https://support.snyk.io/hc/en-us/articles/360004358517-Unknown-PR-test-error
func NewFailedToCompleteTestError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-pr-check-0006",
    Title:      "Failed to complete the test",
    StatusCode: 500,
    ErrorCode:  "SNYK-PR-CHECK-0006",
    Classification: "UNEXPECTED",
    Links: []string{
      "https://support.snyk.io/hc/en-us/articles/360004358517-Unknown-PR-test-error",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewFailedToFetchMergeCommitShaError displays errors with the following description:
// In order for snyk test to run, we need the merge commit SHA from the GitHub. For some reason, we couldn’t get it.
// 
// Try closing and then reopening the pull request, or you can Skip the Pull Request Check if it is consistent.
//
// Read more:
// - https://support.snyk.io/hc/en-us/articles/360005281837
func NewFailedToFetchMergeCommitShaError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-pr-check-0007",
    Title:      "Failed to fetch merge commit SHA",
    StatusCode: 500,
    ErrorCode:  "SNYK-PR-CHECK-0007",
    Classification: "UNEXPECTED",
    Links: []string{
      "https://support.snyk.io/hc/en-us/articles/360005281837",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewMergeConflictError displays errors with the following description:
// Merge Conflict Error is not a Snyk specific issue but rather some issues on your SCM environment. As an example, merge conflicts could happen when people make different changes to the same line of the same file, or when one person edits a file and another person deletes the same file.
// 
// To resolve this, you might need to figure out all the merge conflicts on your SCM environment and resolve them to fully remediate these types of errors on Snyk. As a note, this cannot be modified/changed on Snyk's side.
//
// Read more:
// - https://support.snyk.io/hc/en-us/articles/360005281098
func NewMergeConflictError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-pr-check-0008",
    Title:      "Merge conflict error",
    StatusCode: 422,
    ErrorCode:  "SNYK-PR-CHECK-0008",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://support.snyk.io/hc/en-us/articles/360005281098",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewFailedToDetectIssuesError displays errors with the following description:
// Snyk is always trying to check for new issues and vulnerabilities to keep you safe. We do so by testing on your code on webhook Pull Request events and Push events.
// 
// Occasionally you might see a "Failed to detect issues" commit status which may block your PR. This means that we tried to run a test against your changes but unfortunately something went wrong / we encountered an internal problem. If this happens to you try recreating the pull request and if it still occurs reach out and let us know which user, organization and project and commit sha you experienced the issue with on support@snyk.io
//
// Read more:
// - https://support.snyk.io/hc/en-us/articles/360000920678-Failed-to-detect-issues
func NewFailedToDetectIssuesError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-pr-check-0009",
    Title:      "Failed to detect issues",
    StatusCode: 500,
    ErrorCode:  "SNYK-PR-CHECK-0009",
    Classification: "UNEXPECTED",
    Links: []string{
      "https://support.snyk.io/hc/en-us/articles/360000920678-Failed-to-detect-issues",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewInvalidThirdPartyCredentialsError displays errors with the following description:
// Snyk uses credentials configured on your integration to test your code and to update your PR Check.
// 
// If this error occurs, please ensure your integration and credentials are correctly set up, by following the instructions for your SCM here: https://docs.snyk.io/integrate-with-snyk/git-repository-scm-integrations
func NewInvalidThirdPartyCredentialsError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-pr-check-0010",
    Title:      "No valid credentials to process PR check",
    StatusCode: 401,
    ErrorCode:  "SNYK-PR-CHECK-0010",
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

// NewFailedToGenerateCommitStatusError displays errors with the following description:
// Snyk is always trying to check for new issues and vulnerabilities to keep you safe. We do so by testing on your code on webhook Pull Request events and Push events.
// 
// Occasionally you might see a "Failed to generate a commit status" which may block your PR. This means that we tried to run a test against your changes but unfortunately something went wrong / we encountered an internal problem. If this happens to you try recreating the pull request and if it still occurs reach out and let us know which user, organization and project and commit sha you experienced the issue with on support@snyk.io
func NewFailedToGenerateCommitStatusError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-pr-check-0011",
    Title:      "Failed to generate a commit status",
    StatusCode: 500,
    ErrorCode:  "SNYK-PR-CHECK-0011",
    Classification: "UNEXPECTED",
    Links: []string{},
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

