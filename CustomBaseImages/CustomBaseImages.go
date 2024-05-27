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

// Package CustomBaseImages contains errors related to the namespace CustomBaseImages
// of the Error Catalog.
package CustomBaseImages

import (
  "github.com/snyk/error-catalog-golang-public/snyk_errors"
  "github.com/google/uuid"
)

// NewVersioningSchemaDoesNotSupportTagError displays errors with the following description:
// The versioning schema used does not support the given tag. Update the versioning schema to include the tag.
// 
// Once the tag of the custom base image is correct, the versioning schema must be modified.
// You can use a different versioning schema that supports all tags in the repository or you can update the relevant properties of the versioning schema.
// 
// For example, if the repository currently uses Semver, and a new tag "1.2.5.7" needs to be added, then you can use a Custom versioning schema.
//
// Read more:
// - https://docs.snyk.io/scan-using-snyk/snyk-container/use-snyk-container-from-the-web-ui/use-custom-base-image-recommendations/versioning-schema-for-custom-base-images
func NewVersioningSchemaDoesNotSupportTagError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cbi-0001",
    Title:      "Versioning schema does not support tag",
    StatusCode: 400,
    ErrorCode:  "SNYK-CBI-0001",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/scan-using-snyk/snyk-container/use-snyk-container-from-the-web-ui/use-custom-base-image-recommendations/versioning-schema-for-custom-base-images",
    },
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewRequiredParameterNotProvidedError displays errors with the following description:
// Provide an ORG ID or GROUP ID.
func NewRequiredParameterNotProvidedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cbi-0002",
    Title:      "Missing required parameter",
    StatusCode: 400,
    ErrorCode:  "SNYK-CBI-0002",
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

// NewProjectDoesNotExistError displays errors with the following description:
// The project could not be found. Check that the project exists, that you have access to the project, and also check that the ID you have provided is the project ID and not a CBI ID.
func NewProjectDoesNotExistError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cbi-0003",
    Title:      "Project does not exist",
    StatusCode: 400,
    ErrorCode:  "SNYK-CBI-0003",
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

// NewProjectIsNotContainerImageError displays errors with the following description:
// The project is not a container image.
//
// Read more:
// - https://docs.snyk.io/scan-using-snyk/snyk-container/use-snyk-container-from-the-web-ui/use-custom-base-image-recommendations
func NewProjectIsNotContainerImageError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cbi-0004",
    Title:      "Project is not a container image",
    StatusCode: 400,
    ErrorCode:  "SNYK-CBI-0004",
    Classification: "UNSUPPORTED",
    Links: []string{
      "https://docs.snyk.io/scan-using-snyk/snyk-container/use-snyk-container-from-the-web-ui/use-custom-base-image-recommendations",
    },
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewProjectDoesNotBelongToGroupError displays errors with the following description:
// The project's org does not belong to a group. In order to use a Custom Base Image, recreate the project and add it to a group or add a group to the org. Note that the group feature is not available to free users.
func NewProjectDoesNotBelongToGroupError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cbi-0005",
    Title:      "Unable to retrieve group",
    StatusCode: 400,
    ErrorCode:  "SNYK-CBI-0005",
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

// NewRequestIdsDoNotMatchError displays errors with the following description:
// The request body ID and the request path ID do not match. Ensure that the values are the same and try again.
func NewRequestIdsDoNotMatchError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cbi-0006",
    Title:      "The values in the request do not match",
    StatusCode: 400,
    ErrorCode:  "SNYK-CBI-0006",
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

// NewRequestBodyAttributesMissingError displays errors with the following description:
// The request body does not contain any attributes that can be updated. Provide the necessary attributes and try again.
func NewRequestBodyAttributesMissingError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cbi-0007",
    Title:      "The request body cannot be updated",
    StatusCode: 400,
    ErrorCode:  "SNYK-CBI-0007",
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

// NewInvalidPaginationCursorError displays errors with the following description:
// The provided pagination cursor is invalid.
func NewInvalidPaginationCursorError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cbi-0008",
    Title:      "Invalid pagination cursor",
    StatusCode: 400,
    ErrorCode:  "SNYK-CBI-0008",
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

// NewUnableToSortByVersionError displays errors with the following description:
// Snyk was unable to filter by version. Provide a repository filter and try again.
func NewUnableToSortByVersionError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cbi-0009",
    Title:      "Unable to sort by version",
    StatusCode: 400,
    ErrorCode:  "SNYK-CBI-0009",
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

// NewUpdateVersioningSchemaFailError displays errors with the following description:
// The versioning schema could not be applied to all images in the repository. Therefore, no resources have been updated. Update the provided versioning schema so that all tags in the repository fit the new schema.
func NewUpdateVersioningSchemaFailError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cbi-0010",
    Title:      "Unable to update versioning schema",
    StatusCode: 400,
    ErrorCode:  "SNYK-CBI-0010",
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

// NewProjectAlreadyLinkedError displays errors with the following description:
// The project ID provided is already linked to another Custom Base Image.
func NewProjectAlreadyLinkedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cbi-0011",
    Title:      "Project is already linked to a custom base image",
    StatusCode: 400,
    ErrorCode:  "SNYK-CBI-0011",
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

// NewVersioningSchemaMissingError displays errors with the following description:
// No versioning schema exists for the repository. This image is the first in its repository. Provide a versioning schema that fits the format of current and future images in this repository.
func NewVersioningSchemaMissingError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cbi-0012",
    Title:      "No versioning schema for repository",
    StatusCode: 400,
    ErrorCode:  "SNYK-CBI-0012",
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

// NewVersioningSchemaInapplicableError displays errors with the following description:
// A versioning schema already exists for repository. Remove the "versioning_schema" property or, if you want to update the versioning schema, use the PATCH endpoint.
func NewVersioningSchemaInapplicableError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cbi-0013",
    Title:      "Unable to apply versioning schema",
    StatusCode: 400,
    ErrorCode:  "SNYK-CBI-0013",
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

// NewImageNotFoundError displays errors with the following description:
// Unable to find the requested custom base image. Try again, and if the error persists, contact Snyk support.
func NewImageNotFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cbi-0014",
    Title:      "Unable to find custom base image",
    StatusCode: 404,
    ErrorCode:  "SNYK-CBI-0014",
    Classification: "UNEXPECTED",
    Links: []string{},
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewImageDoesNotExistError displays errors with the following description:
// The requested custom base image does not exist.
//
// Read more:
// - https://status.snyk.io/
func NewImageDoesNotExistError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cbi-0015",
    Title:      "Custom base image does not exist",
    StatusCode: 404,
    ErrorCode:  "SNYK-CBI-0015",
    Classification: "UNEXPECTED",
    Links: []string{
      "https://status.snyk.io/",
    },
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewImageUpdateFailedError displays errors with the following description:
// An internal error occurred while trying to update a custom base image. Try again, and if the error persists, contact Snyk support.
//
// Read more:
// - https://status.snyk.io/
func NewImageUpdateFailedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cbi-0016",
    Title:      "Unable to update custom base image",
    StatusCode: 500,
    ErrorCode:  "SNYK-CBI-0016",
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

// NewPropertiesRetrievalFailedError displays errors with the following description:
// An internal error occurred while trying to retrieve project properties. Try again, and if the error persists, contact Snyk support.
//
// Read more:
// - https://status.snyk.io/
func NewPropertiesRetrievalFailedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cbi-0017",
    Title:      "Unable to retrieve project properties",
    StatusCode: 500,
    ErrorCode:  "SNYK-CBI-0017",
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

// NewImageCollectionRetrievalFailedError displays errors with the following description:
// An internal error occurred while trying to retrieve the image collection. Try again, and if the error persists, contact Snyk support.
//
// Read more:
// - https://status.snyk.io/
func NewImageCollectionRetrievalFailedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cbi-0018",
    Title:      "Unable to retrieve image collection",
    StatusCode: 500,
    ErrorCode:  "SNYK-CBI-0018",
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

