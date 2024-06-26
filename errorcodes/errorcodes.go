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
  TestLimitReachedError string
  TagsForOrganizationWithoutGroupError string
  ServerError string
}

type OpenSourceEcosystemsCodes struct {
  UnparseableManifestError string
  UnparseableLockFileError string
  UnknownDependencyVersionError string
  MissingHeaderError string
  MissingPayloadError string
  UnprocessableFileError string
  CannotGetFileFromSourceError string
  MissingEnvironmentVariableError string
  BrokeredConnectionNotSupportedError string
  GitCloneFailedError string
  UnsupportedManifestFileError string
  UnsupportedTargetFrameworkError string
  MissingStaticMainFunctionError string
  PublishFailedError string
  FailedToAccessPrivatePackageSourceError string
  MissingMSBuildConditionError string
  NoTargetFrameworksFoundError string
  OutdatedSDKVersionRequestedError string
  ProjectSkippedAndNotFoundError string
  PrivateModuleError string
  GoModFileMissingError string
  SsoReAuthRequiredError string
  IncompleteProjectError string
  InconsistentVendoringError string
  UnsupportedExternalFileGenerationSCMError string
  UnableToAccessPrivateDepsError string
  UnableToUseCredentialsError string
  ToolchainNotAvailableError string
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
  NoReleasedVersionForVersionsRangeError string
  SourceNotSupportedError string
  TimeoutWhenProcessingTheDepTreeError string
  CannotReachConfiguredRepositoryError string
  NoRepoFoundForTheNPMPackageError string
  CouldNotParseNPMRegistryURLError string
  CouldNotFindBrokerURLError string
  UnableToReplaceBrokerURLError string
  BadNPMVersionError string
  UnknownBlobEncodingOnGithubError string
  NoResultsFromForkerProcessesError string
  ChildProcessExecutionError string
  NoValidPackageUpgradesError string
  NoDependencyUpdatesError string
  CouldNotParseJSONFileError string
  Base64EncodeError string
  Base64DecodeError string
  MissingSupportedFileError string
  InvalidConfigurationError string
  PnpmOutOfSyncError string
  PnpmUnsupportedLockfileVersionError string
  YarnPackageNotFoundError string
  UnableToReachRegistryError string
  OutdatedYarnLockFileError string
  PermissionDeniedError string
  UnsupportedRequirementsFileError string
  TooManyManifestFilesError string
  FailedToApplyDependencyUpdatesError string
  PythonPackageNotFoundError string
  SyntaxIssuesError string
  PipUnsupportedPythonVersionError string
  PythonVersionConfictError string
  PipNoMatchingPythonDistributionError string
  InstallationFailureError string
  PipenvUnsupportedPythonVersionError string
  PipenvNoMatchingPythonDistributionError string
}

type PurlVulnerabilityFetchingCodes struct {
  OrganizationNotWhitelistedError string
  AuthorizationRequestFailureError string
  InvalidPurlError string
  NamespaceNotProvidedError string
  UnsupportedEcosystemError string
  MissingComponentError string
  ComponentNotSupportedError string
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
  UnsupportedEcosystemError string
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
  FailedToCreatePRTemplateError string
  FailedToReadPRTemplateError string
  FailedToDeletePRTemplateError string
  PRTemplateInvalidPayloadError string
  FailedToLoadCompiledJSONError string
  FailedToRenderDefaultTemplateError string
}

type CodeCodes struct {
  AnalysisFileCountLimitExceededError string
  AnalysisResultSizeLimitExceededError string
  AnalysisTargetSizeLimitExceededError string
  AnalysisFileNameLengthLimitExceededError string
  FeatureIsNotEnabledError string
  UnsupportedProjectError string
}

type PRChecksCodes struct {
  FailedToReadManifestError string
  ManifestNotFoundError string
  ThirdPartyRateLimitError string
  OutOfSyncError string
  FailedDeterminingProjectTargetError string
  FailedToCompleteTestError string
  FailedToFetchMergeCommitShaError string
  MergeConflictError string
  FailedToDetectIssuesError string
  InvalidThirdPartyCredentialsError string
  FailedToGenerateCommitStatusError string
}

type CLICodes struct {
  ConnectionTimeoutError string
}

type CustomBaseImagesCodes struct {
  VersioningSchemaDoesNotSupportTagError string
  RequiredParameterNotProvidedError string
  ProjectDoesNotExistError string
  ProjectIsNotContainerImageError string
  ProjectDoesNotBelongToGroupError string
  RequestIdsDoNotMatchError string
  RequestBodyAttributesMissingError string
  InvalidPaginationCursorError string
  UnableToSortByVersionError string
  UpdateVersioningSchemaFailError string
  ProjectAlreadyLinkedError string
  VersioningSchemaMissingError string
  VersioningSchemaInapplicableError string
  ImageNotFoundError string
  ImageDoesNotExistError string
  ImageUpdateFailedError string
  PropertiesRetrievalFailedError string
  ImageCollectionRetrievalFailedError string
  CreateVersioningSchemaFailError string
}

var Snyk = SnykCodes {
  TooManyRequestsError: "SNYK-0001",
  NotImplementedError: "SNYK-0002",
  BadRequestError: "SNYK-0003",
  TimeoutError: "SNYK-0004",
  UnauthorisedError: "SNYK-0005",
  TestLimitReachedError: "SNYK-0006",
  TagsForOrganizationWithoutGroupError: "SNYK-0007",
  ServerError: "SNYK-9999",
}

var OpenSourceEcosystems = OpenSourceEcosystemsCodes {
  UnparseableManifestError: "SNYK-OS-0001",
  UnparseableLockFileError: "SNYK-OS-0002",
  UnknownDependencyVersionError: "SNYK-OS-0003",
  MissingHeaderError: "SNYK-OS-0004",
  MissingPayloadError: "SNYK-OS-0005",
  UnprocessableFileError: "SNYK-OS-0006",
  CannotGetFileFromSourceError: "SNYK-OS-0007",
  MissingEnvironmentVariableError: "SNYK-OS-0008",
  BrokeredConnectionNotSupportedError: "SNYK-OS-0009",
  GitCloneFailedError: "SNYK-OS-0010",
  UnsupportedManifestFileError: "SNYK-OS-DOTNET-0001",
  UnsupportedTargetFrameworkError: "SNYK-OS-DOTNET-0002",
  MissingStaticMainFunctionError: "SNYK-OS-DOTNET-0003",
  PublishFailedError: "SNYK-OS-DOTNET-0004",
  FailedToAccessPrivatePackageSourceError: "SNYK-OS-DOTNET-0005",
  MissingMSBuildConditionError: "SNYK-OS-DOTNET-0006",
  NoTargetFrameworksFoundError: "SNYK-OS-DOTNET-0007",
  OutdatedSDKVersionRequestedError: "SNYK-OS-DOTNET-0008",
  ProjectSkippedAndNotFoundError: "SNYK-OS-DOTNET-0009",
  PrivateModuleError: "SNYK-OS-GO-0001",
  GoModFileMissingError: "SNYK-OS-GO-0002",
  SsoReAuthRequiredError: "SNYK-OS-GO-0003",
  IncompleteProjectError: "SNYK-OS-GO-0004",
  InconsistentVendoringError: "SNYK-OS-GO-0005",
  UnsupportedExternalFileGenerationSCMError: "SNYK-OS-GO-0006",
  UnableToAccessPrivateDepsError: "SNYK-OS-GO-0007",
  UnableToUseCredentialsError: "SNYK-OS-GO-0008",
  ToolchainNotAvailableError: "SNYK-OS-GO-0009",
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
  NoReleasedVersionForVersionsRangeError: "SNYK-OS-MAVEN-0014",
  SourceNotSupportedError: "SNYK-OS-MAVEN-0015",
  TimeoutWhenProcessingTheDepTreeError: "SNYK-OS-MAVEN-0016",
  CannotReachConfiguredRepositoryError: "SNYK-OS-MAVEN-0017",
  NoRepoFoundForTheNPMPackageError: "SNYK-OS-NODEJS-0001",
  CouldNotParseNPMRegistryURLError: "SNYK-OS-NODEJS-0002",
  CouldNotFindBrokerURLError: "SNYK-OS-NODEJS-0003",
  UnableToReplaceBrokerURLError: "SNYK-OS-NODEJS-0004",
  BadNPMVersionError: "SNYK-OS-NODEJS-0005",
  UnknownBlobEncodingOnGithubError: "SNYK-OS-NODEJS-0006",
  NoResultsFromForkerProcessesError: "SNYK-OS-NODEJS-0007",
  ChildProcessExecutionError: "SNYK-OS-NODEJS-0008",
  NoValidPackageUpgradesError: "SNYK-OS-NODEJS-0009",
  NoDependencyUpdatesError: "SNYK-OS-NODEJS-0010",
  CouldNotParseJSONFileError: "SNYK-OS-NODEJS-0011",
  Base64EncodeError: "SNYK-OS-NODEJS-0012",
  Base64DecodeError: "SNYK-OS-NODEJS-0013",
  MissingSupportedFileError: "SNYK-OS-NODEJS-0014",
  InvalidConfigurationError: "SNYK-OS-NODEJS-0015",
  PnpmOutOfSyncError: "SNYK-OS-NODEJS-0016",
  PnpmUnsupportedLockfileVersionError: "SNYK-OS-NODEJS-0017",
  YarnPackageNotFoundError: "SNYK-OS-NODEJS-0019",
  UnableToReachRegistryError: "SNYK-OS-NODEJS-0020",
  OutdatedYarnLockFileError: "SNYK-OS-NODEJS-0021",
  PermissionDeniedError: "SNYK-OS-NODEJS-0022",
  UnsupportedRequirementsFileError: "SNYK-OS-PYTHON-0001",
  TooManyManifestFilesError: "SNYK-OS-PYTHON-0002",
  FailedToApplyDependencyUpdatesError: "SNYK-OS-PYTHON-0003",
  PythonPackageNotFoundError: "SNYK-OS-PYTHON-0004",
  SyntaxIssuesError: "SNYK-OS-PYTHON-0005",
  PipUnsupportedPythonVersionError: "SNYK-OS-PYTHON-0006",
  PythonVersionConfictError: "SNYK-OS-PYTHON-0007",
  PipNoMatchingPythonDistributionError: "SNYK-OS-PYTHON-0008",
  InstallationFailureError: "SNYK-OS-PYTHON-0009",
  PipenvUnsupportedPythonVersionError: "SNYK-OS-PYTHON-0010",
  PipenvNoMatchingPythonDistributionError: "SNYK-OS-PYTHON-0011",
}

var PurlVulnerabilityFetching = PurlVulnerabilityFetchingCodes {
  OrganizationNotWhitelistedError: "SNYK-OSSI-1040",
  AuthorizationRequestFailureError: "SNYK-OSSI-1050",
  InvalidPurlError: "SNYK-OSSI-2010",
  NamespaceNotProvidedError: "SNYK-OSSI-2011",
  UnsupportedEcosystemError: "SNYK-OSSI-2020",
  MissingComponentError: "SNYK-OSSI-2021",
  ComponentNotSupportedError: "SNYK-OSSI-2022",
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
  UnsupportedEcosystemError: "SNYK-OS-8003",
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
  FailedToCreatePRTemplateError: "SNYK-PR-TEMPLATE-0007",
  FailedToReadPRTemplateError: "SNYK-PR-TEMPLATE-0008",
  FailedToDeletePRTemplateError: "SNYK-PR-TEMPLATE-0009",
  PRTemplateInvalidPayloadError: "SNYK-PR-TEMPLATE-0010",
  FailedToLoadCompiledJSONError: "SNYK-PR-TEMPLATE-0011",
  FailedToRenderDefaultTemplateError: "SNYK-PR-TEMPLATE-0012",
}

var Code = CodeCodes {
  AnalysisFileCountLimitExceededError: "SNYK-CODE-0001",
  AnalysisResultSizeLimitExceededError: "SNYK-CODE-0002",
  AnalysisTargetSizeLimitExceededError: "SNYK-CODE-0003",
  AnalysisFileNameLengthLimitExceededError: "SNYK-CODE-0004",
  FeatureIsNotEnabledError: "SNYK-CODE-0005",
  UnsupportedProjectError: "SNYK-CODE-0006",
}

var PRChecks = PRChecksCodes {
  FailedToReadManifestError: "SNYK-PR-CHECK-0001",
  ManifestNotFoundError: "SNYK-PR-CHECK-0002",
  ThirdPartyRateLimitError: "SNYK-PR-CHECK-0003",
  OutOfSyncError: "SNYK-PR-CHECK-0004",
  FailedDeterminingProjectTargetError: "SNYK-PR-CHECK-0005",
  FailedToCompleteTestError: "SNYK-PR-CHECK-0006",
  FailedToFetchMergeCommitShaError: "SNYK-PR-CHECK-0007",
  MergeConflictError: "SNYK-PR-CHECK-0008",
  FailedToDetectIssuesError: "SNYK-PR-CHECK-0009",
  InvalidThirdPartyCredentialsError: "SNYK-PR-CHECK-0010",
  FailedToGenerateCommitStatusError: "SNYK-PR-CHECK-0011",
}

var CLI = CLICodes {
  ConnectionTimeoutError: "SNYK-OS-7001",
}

var CustomBaseImages = CustomBaseImagesCodes {
  VersioningSchemaDoesNotSupportTagError: "SNYK-CBI-0001",
  RequiredParameterNotProvidedError: "SNYK-CBI-0002",
  ProjectDoesNotExistError: "SNYK-CBI-0003",
  ProjectIsNotContainerImageError: "SNYK-CBI-0004",
  ProjectDoesNotBelongToGroupError: "SNYK-CBI-0005",
  RequestIdsDoNotMatchError: "SNYK-CBI-0006",
  RequestBodyAttributesMissingError: "SNYK-CBI-0007",
  InvalidPaginationCursorError: "SNYK-CBI-0008",
  UnableToSortByVersionError: "SNYK-CBI-0009",
  UpdateVersioningSchemaFailError: "SNYK-CBI-0010",
  ProjectAlreadyLinkedError: "SNYK-CBI-0011",
  VersioningSchemaMissingError: "SNYK-CBI-0012",
  VersioningSchemaInapplicableError: "SNYK-CBI-0013",
  ImageNotFoundError: "SNYK-CBI-0014",
  ImageDoesNotExistError: "SNYK-CBI-0015",
  ImageUpdateFailedError: "SNYK-CBI-0016",
  PropertiesRetrievalFailedError: "SNYK-CBI-0017",
  ImageCollectionRetrievalFailedError: "SNYK-CBI-0018",
  CreateVersioningSchemaFailError: "SNYK-CBI-0019",
}

