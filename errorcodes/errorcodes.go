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
  BadGatewayError string
  ServiceUnavailableError string
  RequirementsNotMetError string
  MaintenanceWindowError string
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
  UnsupportedPlatformError string
  UnsupportedManifestFileError string
  UnsupportedTargetFrameworkError string
  MissingStaticMainFunctionError string
  PublishFailedError string
  FailedToAccessPrivatePackageSourceError string
  MissingMSBuildConditionError string
  NoTargetFrameworksFoundError string
  OutdatedSDKVersionRequestedError string
  ProjectSkippedAndNotFoundError string
  NugetDependenciesSpaceLimitExceededError string
  RestoreFailedError string
  PrivateModuleError string
  GoModFileMissingError string
  SsoReAuthRequiredError string
  IncompleteProjectError string
  InconsistentVendoringError string
  UnsupportedExternalFileGenerationSCMError string
  UnableToAccessPrivateDepsError string
  UnableToUseCredentialsError string
  ToolchainNotAvailableError string
  GolangSpaceLimitExceededError string
  GolangNoSecureProtocolFoundError string
  GolangConnectionResetByPeerError string
  GolangInvalidZipFileError string
  GolangVersionMismatchError string
  GolangInvalidGoVersionError string
  GolangDialTcpTimeoutError string
  GolangHostKeyVerificationFailedError string
  GolangMissingModuleDeclarationError string
  GolangModuleVersionConstraintNotMetError string
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
  PythonDependenciesSpaceLimitExceededError string
  CyclicDependencyDetectedError string
  GemNotFoundError string
  GemVersionConflictError string
  ReachabilitySettingDisabledError string
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
  UnsupportedAlpineDistroError string
}

type IsolatedBuildsCodes struct {
  InvalidRequestError string
  BuildEnvironmentNotFoundError string
  UnsupportedEcosystemError string
  SsoReAuthRequiredError string
  ProjectTooBigError string
  DefaultImageNotFoundError string
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
  ConflictError string
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

type SbomTestCodes struct {
  InternalError string
  OrgIDMismatchError string
  NotFoundError string
  FailedTestError string
  PendingTestError string
  FormatUnknownError string
  UnprocessableInputError string
  FormatNotSupportedError string
  ConversionFailedError string
  NoTestablePackagesError string
}

type FixCodes struct {
  FixScenarioNotSupportedError string
  SCMRateLimitError string
  UnauthorisedAccessError string
  UnsupportedEcosystemError string
  MetadataNotFoundError string
  NoMatureVersionsFoundError string
  VersionNotFoundError string
  AlreadyLatestVersionError string
  DowngradeVersionUnsupportedError string
  VersionParsingError string
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
  RuleExtensionAlreadyExistsForGroupError string
  OrgRelationshipsMustBeUniqueError string
  GroupRelationshipMustBeForAdminGroupError string
  OrgOutsideAdminGroupError string
  RuleExtensionsLimitReachedError string
  TestRuleExtensionAlreadyPublishedForGroupError string
  TestIDNotFoundError string
  TestResultsExpiredError string
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
  GeneralCLIFailureError string
  ConfigEnvironmentFailedError string
  ConfigEnvironmentConsistencyIssueError string
  EmptyFlagOptionError string
  InvalidFlagOptionError string
  GetVulnsFromResourceFailedError string
  AuthConfigError string
  CommandArgsError string
  NoSupportedFilesFoundError string
  TooManyVulnerablePathsError string
  ValidationFailureError string
  GeneralSCAFailureError string
  GeneralIACFailureError string
  GeneralSASTFailureError string
  FeatureUnderDevelopmentError string
  CommandIsExperimentalError string
  FeatureNotEnabledError string
  DNSResolutionError string
  NetworkTimeoutError string
  NetworkUnreachableError string
  TLSCertificateError string
  ConnectionRefusedError string
  GenericNetworkError string
  GeneralSecretsFailureError string
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

type IntegrationCodes struct {
  IntegrationNotFoundError string
}

type TargetCodes struct {
  TargetNotFoundError string
  NoUniqueTargetFoundError string
}

type SCMCodes struct {
  UnsupportedIntegrationTypeError string
  RevisionNotResolvedError string
  IntegrationAuthenticationFailedError string
  IntegrationAuthorizationFailedError string
  FilesLimitExceededError string
  SizeLimitExceededError string
  ResourceNotFoundError string
}

type PoliciesCodes struct {
  InvalidPolicyApplyError string
}

type AiBomCodes struct {
  InternalError string
  ForbiddenError string
  NoSupportedFilesError string
}

type UploadRevisionCodes struct {
  UploadRevisionNotFoundError string
  UploadRevisionSealedError string
  FileTooLargeError string
  TotalFilesSizeLimitExceededError string
  FileCountLimitExceededError string
  FilePathTooLongError string
  PopulateRequestLimitExceededError string
  TotalUploadRevisionFileCountLimitExceededError string
  TotalUploadRevisionSizeLimitExceededError string
  UploadRevisionIdMismatchError string
  MultipartFieldNameMissingError string
}

var Snyk = SnykCodes {
  TooManyRequestsError: "SNYK-0001",
  NotImplementedError: "SNYK-0002",
  BadRequestError: "SNYK-0003",
  TimeoutError: "SNYK-0004",
  UnauthorisedError: "SNYK-0005",
  TestLimitReachedError: "SNYK-0006",
  TagsForOrganizationWithoutGroupError: "SNYK-0007",
  BadGatewayError: "SNYK-0008",
  ServiceUnavailableError: "SNYK-0009",
  RequirementsNotMetError: "SNYK-0010",
  MaintenanceWindowError: "SNYK-0099",
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
  UnsupportedPlatformError: "SNYK-OS-0011",
  UnsupportedManifestFileError: "SNYK-OS-DOTNET-0001",
  UnsupportedTargetFrameworkError: "SNYK-OS-DOTNET-0002",
  MissingStaticMainFunctionError: "SNYK-OS-DOTNET-0003",
  PublishFailedError: "SNYK-OS-DOTNET-0004",
  FailedToAccessPrivatePackageSourceError: "SNYK-OS-DOTNET-0005",
  MissingMSBuildConditionError: "SNYK-OS-DOTNET-0006",
  NoTargetFrameworksFoundError: "SNYK-OS-DOTNET-0007",
  OutdatedSDKVersionRequestedError: "SNYK-OS-DOTNET-0008",
  ProjectSkippedAndNotFoundError: "SNYK-OS-DOTNET-0009",
  NugetDependenciesSpaceLimitExceededError: "SNYK-OS-DOTNET-0010",
  RestoreFailedError: "SNYK-OS-DOTNET-0011",
  PrivateModuleError: "SNYK-OS-GO-0001",
  GoModFileMissingError: "SNYK-OS-GO-0002",
  SsoReAuthRequiredError: "SNYK-OS-GO-0003",
  IncompleteProjectError: "SNYK-OS-GO-0004",
  InconsistentVendoringError: "SNYK-OS-GO-0005",
  UnsupportedExternalFileGenerationSCMError: "SNYK-OS-GO-0006",
  UnableToAccessPrivateDepsError: "SNYK-OS-GO-0007",
  UnableToUseCredentialsError: "SNYK-OS-GO-0008",
  ToolchainNotAvailableError: "SNYK-OS-GO-0009",
  GolangSpaceLimitExceededError: "SNYK-OS-GO-0010",
  GolangNoSecureProtocolFoundError: "SNYK-OS-GO-0011",
  GolangConnectionResetByPeerError: "SNYK-OS-GO-0012",
  GolangInvalidZipFileError: "SNYK-OS-GO-0013",
  GolangVersionMismatchError: "SNYK-OS-GO-0014",
  GolangInvalidGoVersionError: "SNYK-OS-GO-0015",
  GolangDialTcpTimeoutError: "SNYK-OS-GO-0016",
  GolangHostKeyVerificationFailedError: "SNYK-OS-GO-0017",
  GolangMissingModuleDeclarationError: "SNYK-OS-GO-0018",
  GolangModuleVersionConstraintNotMetError: "SNYK-OS-GO-0019",
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
  PythonDependenciesSpaceLimitExceededError: "SNYK-OS-PYTHON-0012",
  CyclicDependencyDetectedError: "SNYK-OS-RUBY-0001",
  GemNotFoundError: "SNYK-OS-RUBY-0002",
  GemVersionConflictError: "SNYK-OS-RUBY-0003",
  ReachabilitySettingDisabledError: "SNYK-OS-SETTINGS-0001",
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
  UnsupportedAlpineDistroError: "SNYK-OSSI-2048",
}

var IsolatedBuilds = IsolatedBuildsCodes {
  InvalidRequestError: "SNYK-OS-8001",
  BuildEnvironmentNotFoundError: "SNYK-OS-8002",
  UnsupportedEcosystemError: "SNYK-OS-8003",
  SsoReAuthRequiredError: "SNYK-OS-8004",
  ProjectTooBigError: "SNYK-OS-8005",
  DefaultImageNotFoundError: "SNYK-OS-8006",
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
  ConflictError: "SNYK-OPENAPI-0009",
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

var SbomTest = SbomTestCodes {
  InternalError: "SNYK-SBOM-0001",
  OrgIDMismatchError: "SNYK-SBOM-0002",
  NotFoundError: "SNYK-SBOM-0003",
  FailedTestError: "SNYK-SBOM-0004",
  PendingTestError: "SNYK-SBOM-0005",
  FormatUnknownError: "SNYK-SBOM-0006",
  UnprocessableInputError: "SNYK-SBOM-0007",
  FormatNotSupportedError: "SNYK-SBOM-0008",
  ConversionFailedError: "SNYK-SBOM-0009",
  NoTestablePackagesError: "SNYK-SBOM-0010",
}

var Fix = FixCodes {
  FixScenarioNotSupportedError: "PR-FAILURES-0001",
  SCMRateLimitError: "PR-FAILURES-0002",
  UnauthorisedAccessError: "PR-FAILURES-0003",
  UnsupportedEcosystemError: "SNYK-PACKAGES-0001",
  MetadataNotFoundError: "SNYK-PACKAGES-0003",
  NoMatureVersionsFoundError: "SNYK-PACKAGES-0005",
  VersionNotFoundError: "SNYK-PACKAGES-0006",
  AlreadyLatestVersionError: "SNYK-PACKAGES-0007",
  DowngradeVersionUnsupportedError: "SNYK-PACKAGES-0008",
  VersionParsingError: "SNYK-PACKAGES-0009",
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
  RuleExtensionAlreadyExistsForGroupError: "SNYK-CODE-0007",
  OrgRelationshipsMustBeUniqueError: "SNYK-CODE-0008",
  GroupRelationshipMustBeForAdminGroupError: "SNYK-CODE-0009",
  OrgOutsideAdminGroupError: "SNYK-CODE-0010",
  RuleExtensionsLimitReachedError: "SNYK-CODE-0011",
  TestRuleExtensionAlreadyPublishedForGroupError: "SNYK-CODE-0012",
  TestIDNotFoundError: "SNYK-CODE-0013",
  TestResultsExpiredError: "SNYK-CODE-0014",
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
  GeneralCLIFailureError: "SNYK-CLI-0000",
  ConfigEnvironmentFailedError: "SNYK-CLI-0001",
  ConfigEnvironmentConsistencyIssueError: "SNYK-CLI-0002",
  EmptyFlagOptionError: "SNYK-CLI-0003",
  InvalidFlagOptionError: "SNYK-CLI-0004",
  GetVulnsFromResourceFailedError: "SNYK-CLI-0005",
  AuthConfigError: "SNYK-CLI-0006",
  CommandArgsError: "SNYK-CLI-0007",
  NoSupportedFilesFoundError: "SNYK-CLI-0008",
  TooManyVulnerablePathsError: "SNYK-CLI-0009",
  ValidationFailureError: "SNYK-CLI-0010",
  GeneralSCAFailureError: "SNYK-CLI-0011",
  GeneralIACFailureError: "SNYK-CLI-0012",
  GeneralSASTFailureError: "SNYK-CLI-0013",
  FeatureUnderDevelopmentError: "SNYK-CLI-0014",
  CommandIsExperimentalError: "SNYK-CLI-0015",
  FeatureNotEnabledError: "SNYK-CLI-0016",
  DNSResolutionError: "SNYK-CLI-0017",
  NetworkTimeoutError: "SNYK-CLI-0018",
  NetworkUnreachableError: "SNYK-CLI-0019",
  TLSCertificateError: "SNYK-CLI-0020",
  ConnectionRefusedError: "SNYK-CLI-0021",
  GenericNetworkError: "SNYK-CLI-0022",
  GeneralSecretsFailureError: "SNYK-CLI-0023",
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

var Integration = IntegrationCodes {
  IntegrationNotFoundError: "SNYK-INTEGRATION-0001",
}

var Target = TargetCodes {
  TargetNotFoundError: "SNYK-TARGET-0001",
  NoUniqueTargetFoundError: "SNYK-TARGET-0002",
}

var SCM = SCMCodes {
  UnsupportedIntegrationTypeError: "SNYK-SCM-0001",
  RevisionNotResolvedError: "SNYK-SCM-0002",
  IntegrationAuthenticationFailedError: "SNYK-SCM-0003",
  IntegrationAuthorizationFailedError: "SNYK-SCM-0004",
  FilesLimitExceededError: "SNYK-SCM-0005",
  SizeLimitExceededError: "SNYK-SCM-0006",
  ResourceNotFoundError: "SNYK-SCM-0010",
}

var Policies = PoliciesCodes {
  InvalidPolicyApplyError: "SNYK-POLICY-0001",
}

var AiBom = AiBomCodes {
  InternalError: "SNYK-AIBOM-0001",
  ForbiddenError: "SNYK-AIBOM-0002",
  NoSupportedFilesError: "SNYK-AIBOM-0003",
}

var UploadRevision = UploadRevisionCodes {
  UploadRevisionNotFoundError: "SNYK-UPLOAD-REVISION-0001",
  UploadRevisionSealedError: "SNYK-UPLOAD-REVISION-0002",
  FileTooLargeError: "SNYK-UPLOAD-REVISION-0003",
  TotalFilesSizeLimitExceededError: "SNYK-UPLOAD-REVISION-0004",
  FileCountLimitExceededError: "SNYK-UPLOAD-REVISION-0005",
  FilePathTooLongError: "SNYK-UPLOAD-REVISION-0006",
  PopulateRequestLimitExceededError: "SNYK-UPLOAD-REVISION-0007",
  TotalUploadRevisionFileCountLimitExceededError: "SNYK-UPLOAD-REVISION-0008",
  TotalUploadRevisionSizeLimitExceededError: "SNYK-UPLOAD-REVISION-0009",
  UploadRevisionIdMismatchError: "SNYK-UPLOAD-REVISION-0010",
  MultipartFieldNameMissingError: "SNYK-UPLOAD-REVISION-0011",
}

