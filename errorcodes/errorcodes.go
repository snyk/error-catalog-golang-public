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
// Package errorcodes contains a shortcut to comparing error codes with human-readable error messages.
//
// Example:
//  errorCode := "SNYK-OS-9000"
//  if errorCode == errorcodes.SbomExport.InternalServerError {
//   // Do something
//  }
package errorcodes

type SnykCodes struct {
  TooManyRequestsError string
  NotImplementedError string
  BadRequestError string
  TimeoutError string
  UnauthorisedError string
  ServerError string
}

type OpenSourceEcosystemsCodes struct {
  UnsupportedEcosystemError string
  UnparseableManifestError string
  LockFileOutOfSyncError string
  UnparseableLockFileError string
  UnknownDependencyVersionError string
  MissingPayloadError string
  UnprocessableFileError string
  SourceNotSupportedError string
  CannotGetFileFromSourceError string
  NoReleasedVersionForVersionsRangeError string
  TimeoutWhenProcessingTheDepTreeError string
  TooManyManifestFilesError string
  FailedToApplyDependencyUpdatesError string
  UnknownBlobEncodingOnGithubError string
  NoResultsFromForkerProcessesError string
  ChildProcessExecutionError string
  NoValidPackageUpgradesError string
  NoDependencyUpdatesError string
  MissingRequiredRequestHeaderError string
  CouldNotParseJSONFileError string
  Base64EncodeError string
  Base64DecodeError string
  MissingSupportedFileError string
  MissingEnvironmentVariableError string
  NoOutputFromIsolatedBuildsError string
  FailedToRelockError string
  InvalidConfigurationError string
  UnsupportedManifestFileError string
  UnsupportedTargetFrameworkError string
  PrivateModuleError string
  GoModFileMissingError string
  SsoReAuthRequiredError string
  IncompleteProjectError string
  MissingRequirementFromPomError string
  UnableToResolveValueForPropertyError string
  UnableToResolveVersionForPropertyError string
  CyclicPropertyDetectedInPomFileError string
  UnableToParseXMLError string
  InvalidCoordinatesError string
  SkippedGroupError string
  PomFileNotFoundError string
  MissingProjectFromPomError string
  CannotResolveTargetPomFromXmlError string
  CannotResolveTargetPomFromRepoError string
  CannotGetBuildFileFromRepoError string
  CannotCreateGitHostError string
  NoRepoFoundForTheNPMPackageError string
  CouldNotParseNPMRegistryURLError string
  CouldNotFindBrokerURLError string
  UnableToReplaceBrokerURLError string
  BadNPMVersionError string
  UnsupportedRequirementsFileError string
}

type PurlVulnerabilityFetchingCodes struct {
  OrganizationNotWhitelistedError string
  AuthorizationRequestFailureError string
  InvalidPurlError string
  NamespaceNotProvidedError string
  UnsupportedEcosystemError string
  MissingComponentError string
  ComponentNotSupportedError string
  UnsupportedGoVersionFormatError string
  PackageNotFoundError string
  VulnerabilityServiceUnavailableError string
  VulnDBInvalidResponseError string
  VulndbNextError string
  InternalServerError string
  InvalidPaginationParametersError string
  TooManyPurlsError string
  TooManyIssuesError string
  UndefinedContainerDistroError string
  UnsupportedDebianDistroError string
  UndefinedContainerVendorError string
  UnsupportedContainerVendorError string
}

type IsolatedBuildsCodes struct {
  InvalidRequestError string
  BuildEnvironmentNotFoundError string
}

type OpenSourceProjectSnapshotsCodes struct {
  InvalidRequestError string
  InvalidResponseError string
  DataTransformationError string
  StorageFailureError string
  InternalServerError string
}

type OpenSourceProjectIssuesCodes struct {
  InvalidRequestError string
  InvalidResponseError string
  DataTransformationError string
  StorageFailureError string
  InternalServerError string
}

type OpenAPICodes struct {
  BadRequestError string
  ForbiddenError string
  NotAcceptableError string
  NotFoundError string
  MethodNotAllowedError string
  RequestEntityTooLargeError string
  UnauthorizedError string
  UnsupportedMediaTypeError string
}

type OpenSourceUnmanagedCodes struct {
  MavenSearchServiceUnavailableError string
  Sha1NotFoundError string
}

type SbomExportCodes struct {
  InternalServerError string
  UnexpectedDepGraphResponseError string
  UnexpectedParseDepGraphError string
  IaCOrSASTProjectError string
  UnsupportedProjectError string
  DepGraphResponseError string
  MissingAuthTokenError string
  EmptyRequestBodyError string
  InvalidDepGraphError string
}

type FixCodes struct {
  FailedToGetPullRequestAttributesError string
  PullRequestTemplateNotFoundError string
  FailedToCompilePrTemplateError string
  FailedToParsePullRequestAttributesError string
  FailedToLoadCompiledYamlError string
  FailedToGenerateHashError string
}

var Snyk = SnykCodes {
  TooManyRequestsError: "SNYK-0001",
  NotImplementedError: "SNYK-0002",
  BadRequestError: "SNYK-0003",
  TimeoutError: "SNYK-0004",
  UnauthorisedError: "SNYK-0005",
  ServerError: "SNYK-9999",
}

var OpenSourceEcosystems = OpenSourceEcosystemsCodes {
  UnsupportedEcosystemError: "SNYK-OS-0001",
  UnparseableManifestError: "SNYK-OS-0002",
  LockFileOutOfSyncError: "SNYK-OS-0003",
  UnparseableLockFileError: "SNYK-OS-0004",
  UnknownDependencyVersionError: "SNYK-OS-0005",
  MissingPayloadError: "SNYK-OS-0006",
  UnprocessableFileError: "SNYK-OS-0007",
  SourceNotSupportedError: "SNYK-OS-0008",
  CannotGetFileFromSourceError: "SNYK-OS-0009",
  NoReleasedVersionForVersionsRangeError: "SNYK-OS-0010",
  TimeoutWhenProcessingTheDepTreeError: "SNYK-OS-0011",
  TooManyManifestFilesError: "SNYK-OS-0012",
  FailedToApplyDependencyUpdatesError: "SNYK-OS-0013",
  UnknownBlobEncodingOnGithubError: "SNYK-OS-0014",
  NoResultsFromForkerProcessesError: "SNYK-OS-0015",
  ChildProcessExecutionError: "SNYK-OS-0016",
  NoValidPackageUpgradesError: "SNYK-OS-0017",
  NoDependencyUpdatesError: "SNYK-OS-0018",
  MissingRequiredRequestHeaderError: "SNYK-OS-0019",
  CouldNotParseJSONFileError: "SNYK-OS-0020",
  Base64EncodeError: "SNYK-OS-0021",
  Base64DecodeError: "SNYK-OS-0022",
  MissingSupportedFileError: "SNYK-OS-0023",
  MissingEnvironmentVariableError: "SNYK-OS-0024",
  NoOutputFromIsolatedBuildsError: "SNYK-OS-0025",
  FailedToRelockError: "SNYK-OS-0026",
  InvalidConfigurationError: "SNYK-OS-0027",
  UnsupportedManifestFileError: "SNYK-OS-DOTNET-0001",
  UnsupportedTargetFrameworkError: "SNYK-OS-DOTNET-0002",
  PrivateModuleError: "SNYK-OS-GO-0001",
  GoModFileMissingError: "SNYK-OS-GO-0002",
  SsoReAuthRequiredError: "SNYK-OS-GO-0003",
  IncompleteProjectError: "SNYK-OS-GO-0004",
  MissingRequirementFromPomError: "SNYK-OS-MAVEN-0001",
  UnableToResolveValueForPropertyError: "SNYK-OS-MAVEN-0002",
  UnableToResolveVersionForPropertyError: "SNYK-OS-MAVEN-0003",
  CyclicPropertyDetectedInPomFileError: "SNYK-OS-MAVEN-0004",
  UnableToParseXMLError: "SNYK-OS-MAVEN-0005",
  InvalidCoordinatesError: "SNYK-OS-MAVEN-0006",
  SkippedGroupError: "SNYK-OS-MAVEN-0007",
  PomFileNotFoundError: "SNYK-OS-MAVEN-0008",
  MissingProjectFromPomError: "SNYK-OS-MAVEN-0009",
  CannotResolveTargetPomFromXmlError: "SNYK-OS-MAVEN-0010",
  CannotResolveTargetPomFromRepoError: "SNYK-OS-MAVEN-0011",
  CannotGetBuildFileFromRepoError: "SNYK-OS-MAVEN-0012",
  CannotCreateGitHostError: "SNYK-OS-MAVEN-0013",
  NoRepoFoundForTheNPMPackageError: "SNYK-OS-NODEJS-0001",
  CouldNotParseNPMRegistryURLError: "SNYK-OS-NODEJS-0002",
  CouldNotFindBrokerURLError: "SNYK-OS-NODEJS-0003",
  UnableToReplaceBrokerURLError: "SNYK-OS-NODEJS-0004",
  BadNPMVersionError: "SNYK-OS-NODEJS-0005",
  UnsupportedRequirementsFileError: "SNYK-OS-PIP-0001",
}

var PurlVulnerabilityFetching = PurlVulnerabilityFetchingCodes {
  OrganizationNotWhitelistedError: "SNYK-OSSI-1040",
  AuthorizationRequestFailureError: "SNYK-OSSI-1050",
  InvalidPurlError: "SNYK-OSSI-2010",
  NamespaceNotProvidedError: "SNYK-OSSI-2011",
  UnsupportedEcosystemError: "SNYK-OSSI-2020",
  MissingComponentError: "SNYK-OSSI-2021",
  ComponentNotSupportedError: "SNYK-OSSI-2022",
  UnsupportedGoVersionFormatError: "SNYK-OSSI-2023",
  PackageNotFoundError: "SNYK-OSSI-2030",
  VulnerabilityServiceUnavailableError: "SNYK-OSSI-2031",
  VulnDBInvalidResponseError: "SNYK-OSSI-2032",
  VulndbNextError: "SNYK-OSSI-2033",
  InternalServerError: "SNYK-OSSI-2040",
  InvalidPaginationParametersError: "SNYK-OSSI-2041",
  TooManyPurlsError: "SNYK-OSSI-2042",
  TooManyIssuesError: "SNYK-OSSI-2043",
  UndefinedContainerDistroError: "SNYK-OSSI-2044",
  UnsupportedDebianDistroError: "SNYK-OSSI-2045",
  UndefinedContainerVendorError: "SNYK-OSSI-2046",
  UnsupportedContainerVendorError: "SNYK-OSSI-2047",
}

var IsolatedBuilds = IsolatedBuildsCodes {
  InvalidRequestError: "SNYK-OS-8001",
  BuildEnvironmentNotFoundError: "SNYK-OS-8002",
}

var OpenSourceProjectSnapshots = OpenSourceProjectSnapshotsCodes {
  InvalidRequestError: "SNYK-OSSI-OSPSS-1001",
  InvalidResponseError: "SNYK-OSSI-OSPSS-1002",
  DataTransformationError: "SNYK-OSSI-OSPSS-2001",
  StorageFailureError: "SNYK-OSSI-OSPSS-3001",
  InternalServerError: "SNYK-OSSI-OSPSS-4001",
}

var OpenSourceProjectIssues = OpenSourceProjectIssuesCodes {
  InvalidRequestError: "SNYK-OSSI-OSPI-1001",
  InvalidResponseError: "SNYK-OSSI-OSPI-1002",
  DataTransformationError: "SNYK-OSSI-OSPI-2001",
  StorageFailureError: "SNYK-OSSI-OSPI-3001",
  InternalServerError: "SNYK-OSSI-OSPI-4001",
}

var OpenAPI = OpenAPICodes {
  BadRequestError: "SNYK-OPENAPI-0001",
  ForbiddenError: "SNYK-OPENAPI-0002",
  NotAcceptableError: "SNYK-OPENAPI-0003",
  NotFoundError: "SNYK-OPENAPI-0004",
  MethodNotAllowedError: "SNYK-OPENAPI-0005",
  RequestEntityTooLargeError: "SNYK-OPENAPI-0006",
  UnauthorizedError: "SNYK-OPENAPI-0007",
  UnsupportedMediaTypeError: "SNYK-OPENAPI-0008",
}

var OpenSourceUnmanaged = OpenSourceUnmanagedCodes {
  MavenSearchServiceUnavailableError: "SNYK-OSJVM-001",
  Sha1NotFoundError: "SNYK-OSJVM-002",
}

var SbomExport = SbomExportCodes {
  InternalServerError: "SNYK-OS-9000",
  UnexpectedDepGraphResponseError: "SNYK-OS-9001",
  UnexpectedParseDepGraphError: "SNYK-OS-9002",
  IaCOrSASTProjectError: "SNYK-OS-9003",
  UnsupportedProjectError: "SNYK-OS-9004",
  DepGraphResponseError: "SNYK-OS-9005",
  MissingAuthTokenError: "SNYK-OS-9006",
  EmptyRequestBodyError: "SNYK-OS-9007",
  InvalidDepGraphError: "SNYK-OS-9008",
}

var Fix = FixCodes {
  FailedToGetPullRequestAttributesError: "SNYK-PR-TEMPLATE-0001",
  PullRequestTemplateNotFoundError: "SNYK-PR-TEMPLATE-0002",
  FailedToCompilePrTemplateError: "SNYK-PR-TEMPLATE-0003",
  FailedToParsePullRequestAttributesError: "SNYK-PR-TEMPLATE-0004",
  FailedToLoadCompiledYamlError: "SNYK-PR-TEMPLATE-0005",
  FailedToGenerateHashError: "SNYK-PR-TEMPLATE-0006",
}

