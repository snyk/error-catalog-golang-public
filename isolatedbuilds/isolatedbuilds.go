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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-8001",
    Title:      "Invalid request",
    Description: "The provided request payload is not valid for the selected ecosystem. Please review the API documentation.",
    StatusCode: 400,
    ErrorCode:  "SNYK-OS-8001",
    Classification: "ACTIONABLE",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-8002",
    Title:      "Build environment not found",
    Description: "The build environment for the provided context could not be found. Please ensure you have created the build environment first.",
    StatusCode: 404,
    ErrorCode:  "SNYK-OS-8002",
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

// NewUnsupportedEcosystemError displays errors with the following description:
// The language or package manager is not supported. Please refer to the supported package managers in the links.
//
// Read more:
// - https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/supported-languages-frameworks-and-feature-availability-overview#open-source-and-licensing-snyk-open-source
func NewUnsupportedEcosystemError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-8003",
    Title:      "Unsupported Ecosystem",
    Description: "The language or package manager is not supported. Please refer to the supported package managers in the links.",
    StatusCode: 400,
    ErrorCode:  "SNYK-OS-8003",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/supported-languages-frameworks-and-feature-availability-overview#open-source-and-licensing-snyk-open-source",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewSsoReAuthRequiredError displays errors with the following description:
// Your code is cloned on an isolated environment using Git as it is required by Snyk to analyze its dependencies.
// 
// Your Organization has enabled or enforced SAML SSO after you authorized Snyk to access your code, and a re-authentication is therefore required.
// 
// The error you're seeing is usually reproducible by attempting to do a `git clone` of your repository with incorrectly configured credentials.
// Verify your authentication configuration with your Git cloud provider and try again.
//
// Read more:
// - https://docs.github.com/en/enterprise-cloud@latest/authentication/authenticating-with-saml-single-sign-on/about-authentication-with-saml-single-sign-on#about-oauth-apps-github-apps-and-saml-sso
func NewSsoReAuthRequiredError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-8004",
    Title:      "OAuth re-authorization required",
    Description: "Your code is cloned on an isolated environment using Git as it is required by Snyk to analyze its dependencies.\n\nYour Organization has enabled or enforced SAML SSO after you authorized Snyk to access your code, and a re-authentication is therefore required.\n\nThe error you're seeing is usually reproducible by attempting to do a `git clone` of your repository with incorrectly configured credentials.\nVerify your authentication configuration with your Git cloud provider and try again.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-8004",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.github.com/en/enterprise-cloud@latest/authentication/authenticating-with-saml-single-sign-on/about-authentication-with-saml-single-sign-on#about-oauth-apps-github-apps-and-saml-sso",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewProjectTooBigError displays errors with the following description:
// The project cannot be built or processed due to requiring more memory than available. 
// For node projects, please try again after removing requirement to generate a lockfile when opening a fix PR.
//
// Read more:
// - https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/supported-languages-frameworks-and-feature-availability-overview#howtonotrequestalockfiletobegenerated
func NewProjectTooBigError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-8005",
    Title:      "Project too large to be processed",
    Description: "The project cannot be built or processed due to requiring more memory than available. \nFor node projects, please try again after removing requirement to generate a lockfile when opening a fix PR.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-8005",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/supported-languages-frameworks-and-feature-availability-overview#howtonotrequestalockfiletobegenerated",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewDefaultImageNotFoundError displays errors with the following description:
// Unable to find the default image. Please try again, and contact Snyk support if the error persists.
func NewDefaultImageNotFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-8006",
    Title:      "No default image found in repository",
    Description: "Unable to find the default image. Please try again, and contact Snyk support if the error persists.",
    StatusCode: 500,
    ErrorCode:  "SNYK-OS-8006",
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
