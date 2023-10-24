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

// Package ecosystems contains errors related to the namespace OpenSourceEcosystems
// of the Error Catalog.
package ecosystems

import (
  "github.com/snyk/error-catalog-golang-public/snyk_errors"
  "github.com/google/uuid"
)

// NewUnparseableManifestError displays errors with the following description:
// The provided manifest file could not be parsed as it has invalid syntax or does not match the expected schema. Review the manifest file, then try again.
func NewUnparseableManifestError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-0001",
    Title:      "Unable to parse manifest file",
    StatusCode: 400,
    ErrorCode:  "SNYK-OS-0001",
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

// NewUnparseableLockFileError displays errors with the following description:
// The provided lock file could not be parsed as it has invalid syntax or does not match the expected schema. Review the lock file, then try again.
func NewUnparseableLockFileError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-0002",
    Title:      "Unable to parse lock file",
    StatusCode: 400,
    ErrorCode:  "SNYK-OS-0002",
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

// NewUnknownDependencyVersionError displays errors with the following description:
// Dependency version could not be resolved.
//
// Read more:
// - https://support.snyk.io/hc/en-us/articles/360001373178-Could-not-determine-version-for-dependencies
func NewUnknownDependencyVersionError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-0003",
    Title:      "Unknown dependency version",
    StatusCode: 404,
    ErrorCode:  "SNYK-OS-0003",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://support.snyk.io/hc/en-us/articles/360001373178-Could-not-determine-version-for-dependencies",
    },
    Level:  "warn",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewMissingHeaderError displays errors with the following description:
// The server encountered a request that is missing a mandatory request header.
func NewMissingHeaderError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-0004",
    Title:      "Missing required request header",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-0004",
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

// NewMissingPayloadError displays errors with the following description:
// The server could not process the request.
func NewMissingPayloadError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-0005",
    Title:      "Payload missing required elements",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-0005",
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

// NewUnprocessableFileError displays errors with the following description:
// The dependency service could not process the files.
func NewUnprocessableFileError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-0006",
    Title:      "Files cannot be processed",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-0006",
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

// NewCannotGetFileFromSourceError displays errors with the following description:
// Could not get the file from the source URL.
func NewCannotGetFileFromSourceError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-0007",
    Title:      "Cannot get file from source",
    StatusCode: 500,
    ErrorCode:  "SNYK-OS-0007",
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

// NewMissingEnvironmentVariableError displays errors with the following description:
// The server encountered a critical operation that requires a specific environment variable, but the variable is not set or is not accessible within the current environment.
func NewMissingEnvironmentVariableError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-0008",
    Title:      "Missing environment variable",
    StatusCode: 500,
    ErrorCode:  "SNYK-OS-0008",
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

// NewUnsupportedManifestFileError displays errors with the following description:
// The provided manifest file is not supported by Snyk for .NET.
//
// Read more:
// - https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/.net
func NewUnsupportedManifestFileError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-dotnet-0001",
    Title:      "Unsupported manifest file type for remediation",
    StatusCode: 400,
    ErrorCode:  "SNYK-OS-DOTNET-0001",
    Classification: "UNSUPPORTED",
    Links: []string{
      "https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/.net",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewUnsupportedTargetFrameworkError displays errors with the following description:
// The provided manifest file defines a `<TargetFramework>` or `<TargetFrameworks>` that is not currently supported by Snyk's .NET scanning solution.
func NewUnsupportedTargetFrameworkError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-dotnet-0002",
    Title:      "Target framework not supported",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-DOTNET-0002",
    Classification: "UNSUPPORTED",
    Links: []string{},
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewMissingStaticMainFunctionError displays errors with the following description:
// This error occurs when no static Main method with a correct signature is found in the code that produces an executable file. 
// It also occurs if the entry point function, `Main`, is defined with the wrong case, such as lower-case main.
// 
// In order to fix this issue, ensure that your program has a .cs file that contains a main function, such as
// ```c#
// namespace Example
// {
//     class Program
//     {
//         static void Main(string[] args)
//         {
//             Console.WriteLine("hello world");
//         }
//     }
// }
// ```
//
// Read more:
// - https://learn.microsoft.com/en-us/dotnet/csharp/misc/cs5001
func NewMissingStaticMainFunctionError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-dotnet-0003",
    Title:      "Your C# code is missing a static Main function",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-DOTNET-0003",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://learn.microsoft.com/en-us/dotnet/csharp/misc/cs5001",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewPublishFailedError displays errors with the following description:
// This error occurs when running `dotnet publish --sc --framework <your-target-framework>` fails to generate a 
// self-contained binary. Snyk needs to run this command in order to adequately determine the dependency tree for your project. If this command fails, Snyk cannot continue.
// 
// Steps to determine why this happened:
// 
// * Checkout a clean version of your project in a temporary folder
// * Run `dotnet publish --sc --framework <your-target-framework> ` on your project, and confirm this step fails.
// 
// If this step is successful locally, it is possible that Snyk is running another version of the .NET SDK. To tell Snyk which version of the .NET SDK to use, consider using the [global.json](https://learn.microsoft.com/en-us/dotnet/core/tools/global-json) solution provided by Microsoft.
//
// Read more:
// - https://learn.microsoft.com/en-us/dotnet/core/tools/sdk-errors/
// - https://learn.microsoft.com/en-us/dotnet/core/tools/global-json
// - https://github.com/snyk/snyk-nuget-plugin/blob/885486aa656c28d3db465c8d22710770d5cc6773/lib/nuget-parser/cli/dotnet.ts#L67
func NewPublishFailedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-dotnet-0004",
    Title:      "The dotnet CLI is unable to generate a self-contained binary",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-DOTNET-0004",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://learn.microsoft.com/en-us/dotnet/core/tools/sdk-errors/",
      "https://learn.microsoft.com/en-us/dotnet/core/tools/global-json",
      "https://github.com/snyk/snyk-nuget-plugin/blob/885486aa656c28d3db465c8d22710770d5cc6773/lib/nuget-parser/cli/dotnet.ts#L67",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewPrivateModuleError displays errors with the following description:
// Snyk could not access the private modules within your go.mod files.
//
// Read more:
// - https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/go
func NewPrivateModuleError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-go-0001",
    Title:      "Failed to access private module",
    StatusCode: 400,
    ErrorCode:  "SNYK-OS-GO-0001",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/go",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewGoModFileMissingError displays errors with the following description:
// A go.mod file was not found in the current directory or any parent directory.
//
// Read more:
// - https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/go
func NewGoModFileMissingError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-go-0002",
    Title:      "Go mod file not found",
    StatusCode: 400,
    ErrorCode:  "SNYK-OS-GO-0002",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/go",
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
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-go-0003",
    Title:      "OAuth re-authorization required",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-GO-0003",
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

// NewIncompleteProjectError displays errors with the following description:
// Generating the dependency graph requires Snyk to run go list `go list -deps -json` inside the project. If the operation fails, creating a full dependency graph cannot continue.  
// 
// This error means that you need some cleanup, such as `go mod tidy`) or your project deployment process contains a code generation step such as `protobuf` or similar that is not currently supported by Snyk. 
// 
// To verify if this is the case, clone your project in a clean environment, run go list `go list -deps -json` and verify whether the operation fails. 
// 
// If Snyk cannot process your code successfully, insert the Snyk CLI as part of your deployment pipeline.
//
// Read more:
// - https://docs.snyk.io/snyk-cli
// - https://github.com/snyk/snyk-go-plugin
// - https://github.com/golang/go/blob/master/src/cmd/go/internal/list/list.go
func NewIncompleteProjectError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-go-0004",
    Title:      "Your project repository is missing required files",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-GO-0004",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/snyk-cli",
      "https://github.com/snyk/snyk-go-plugin",
      "https://github.com/golang/go/blob/master/src/cmd/go/internal/list/list.go",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewInconsistentVendoringError displays errors with the following description:
// Generating the dependency graph requires Snyk to run `go list -deps -json` inside the project. If the operation fails, creating a full dependency graph cannot continue.  
// 
// This error means that there is inconsistency between your `vendor/modules.txt` file and your `go.mod` file. To remediate, you need to:
// 
// * `go mod vendor`
// * `go mod tidy`
// 
// Next, commit those changes to your repo. Snyk does not manipulate with your code on our end by design, which is why this is not done automatically.
// 
// To verify if this is the case, clone your project in a clean environment, run go list `go list -deps -json` and verify whether the operation fails. 
// Then try and run the above mentioned commands and see if your SCM system reports changes in files.
//
// Read more:
// - https://go.dev/ref/mod#go-mod-vendor
func NewInconsistentVendoringError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-go-0005",
    Title:      "Your project repository has inconsistent vendoring information",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-GO-0005",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://go.dev/ref/mod#go-mod-vendor",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewUnsupportedExternalFileGenerationSCMError displays errors with the following description:
// Snyk currently does not support external file generation in your project. This limitation is due to Snyk's lack of visibility into the third-party generator tools you may be using and the specific commands required to generate these files.
// 
// Snyk can only work with the files available in your repository and does not have insight into the generation process for external files.
func NewUnsupportedExternalFileGenerationSCMError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-go-0006",
    Title:      "Unsupported external file generation",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-GO-0006",
    Classification: "UNSUPPORTED",
    Links: []string{},
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewMissingRequirementFromPomError displays errors with the following description:
// The required property is missing from the pom object.
func NewMissingRequirementFromPomError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-maven-0001",
    Title:      "Missing property",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-MAVEN-0001",
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

// NewUnableToResolveValueForPropertyError displays errors with the following description:
// The targeted property could not be resolved with a valid value.
func NewUnableToResolveValueForPropertyError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-maven-0002",
    Title:      "Unable to resolve value for property",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-MAVEN-0002",
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

// NewUnableToResolveVersionForPropertyError displays errors with the following description:
// The targeted property could not be resolved with a valid version.
func NewUnableToResolveVersionForPropertyError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-maven-0003",
    Title:      "Unable to resolve version for property",
    StatusCode: 500,
    ErrorCode:  "SNYK-OS-MAVEN-0003",
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

// NewCyclicPropertyDetectedInPomFileError displays errors with the following description:
// There is circular dependency among properties in the Maven project's configuration file (POM), preventing proper resolution and causing an error.
func NewCyclicPropertyDetectedInPomFileError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-maven-0004",
    Title:      "Cyclic property detected in POM file",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-MAVEN-0004",
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

// NewUnableToParseXMLError displays errors with the following description:
// There is an error parsing the XML file. This could be referring to either pom.xml or maven-metadata.xml.
func NewUnableToParseXMLError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-maven-0005",
    Title:      "Error parsing the XML file",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-MAVEN-0005",
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

// NewInvalidCoordinatesError displays errors with the following description:
// The coordinates provided for a project were invalid.
func NewInvalidCoordinatesError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-maven-0006",
    Title:      "Invalid coordinates provided",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-MAVEN-0006",
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

// NewSkippedGroupError displays errors with the following description:
// Skipping a specific groupId starting due to remapped coordinates.
func NewSkippedGroupError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-maven-0007",
    Title:      "Skipping group",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-MAVEN-0007",
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

// NewPomFileNotFoundError displays errors with the following description:
// The pom file was not found in Maven repository.
func NewPomFileNotFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-maven-0008",
    Title:      "Pom file not found",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-MAVEN-0008",
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

// NewMissingProjectFromPomError displays errors with the following description:
// A project element is missing from POM.
func NewMissingProjectFromPomError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-maven-0009",
    Title:      "Missing project from POM",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-MAVEN-0009",
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

// NewCannotResolveTargetPomFromXmlError displays errors with the following description:
// Cannot resolve the targeted POM from the input XML.
func NewCannotResolveTargetPomFromXmlError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-maven-0010",
    Title:      "Cannot resolve the target POM from the input XML",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-MAVEN-0010",
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

// NewCannotResolveTargetPomFromRepoError displays errors with the following description:
// Cannot resolve the targeted POM from the repository.
func NewCannotResolveTargetPomFromRepoError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-maven-0011",
    Title:      "Cannot resolve the target POM from the repository",
    StatusCode: 404,
    ErrorCode:  "SNYK-OS-MAVEN-0011",
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

// NewCannotGetBuildFileFromRepoError displays errors with the following description:
// Cannot get the build file repository.
func NewCannotGetBuildFileFromRepoError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-maven-0012",
    Title:      "Cannot get the build file repository",
    StatusCode: 404,
    ErrorCode:  "SNYK-OS-MAVEN-0012",
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

// NewCannotCreateGitHostError displays errors with the following description:
// Cannot create source URL.
func NewCannotCreateGitHostError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-maven-0013",
    Title:      "Unable to create hosted git info",
    StatusCode: 500,
    ErrorCode:  "SNYK-OS-MAVEN-0013",
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

// NewNoReleasedVersionForVersionsRangeError displays errors with the following description:
// There was no version released for the specified versions range.
func NewNoReleasedVersionForVersionsRangeError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-maven-0014",
    Title:      "No released version for versions range",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-MAVEN-0014",
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

// NewSourceNotSupportedError displays errors with the following description:
// The source used is not supported by fetcher. The supported sources are: github, bitbucket, gitlab.
func NewSourceNotSupportedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-maven-0015",
    Title:      "Source is not supported",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-MAVEN-0015",
    Classification: "UNSUPPORTED",
    Links: []string{},
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewTimeoutWhenProcessingTheDepTreeError displays errors with the following description:
// There was an timeout when processing the dependency tree.
func NewTimeoutWhenProcessingTheDepTreeError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-maven-0016",
    Title:      "Timeout when processing the dependency tree",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-MAVEN-0016",
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

// NewNoRepoFoundForTheNPMPackageError displays errors with the following description:
// No repository found for the NPM package.
func NewNoRepoFoundForTheNPMPackageError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-nodejs-0001",
    Title:      "No repository found for A NPM package",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-NODEJS-0001",
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

// NewCouldNotParseNPMRegistryURLError displays errors with the following description:
// Could not parse NPM registry URL.
func NewCouldNotParseNPMRegistryURLError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-nodejs-0002",
    Title:      "Could not parse NPM registry URL",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-NODEJS-0002",
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

// NewCouldNotFindBrokerURLError displays errors with the following description:
// Could not find a broker resolved URL.
func NewCouldNotFindBrokerURLError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-nodejs-0003",
    Title:      "Could not find a broker resolved URL",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-NODEJS-0003",
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

// NewUnableToReplaceBrokerURLError displays errors with the following description:
// Unable to replace all broker urls in lock file.
func NewUnableToReplaceBrokerURLError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-nodejs-0004",
    Title:      "Unable to replace broker URL",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-NODEJS-0004",
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

// NewBadNPMVersionError displays errors with the following description:
// The NPM version is not supported.
func NewBadNPMVersionError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-nodejs-0005",
    Title:      "Bad NPM version",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-NODEJS-0005",
    Classification: "UNSUPPORTED",
    Links: []string{},
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewUnknownBlobEncodingOnGithubError displays errors with the following description:
// Unknown blob encoding on Github.
func NewUnknownBlobEncodingOnGithubError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-nodejs-0006",
    Title:      "Unknown blob encoding on Github",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-NODEJS-0006",
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

// NewNoResultsFromForkerProcessesError displays errors with the following description:
// No result from forked process.
func NewNoResultsFromForkerProcessesError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-nodejs-0007",
    Title:      "No result from forked process",
    StatusCode: 500,
    ErrorCode:  "SNYK-OS-NODEJS-0007",
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

// NewChildProcessExecutionError displays errors with the following description:
// The child process encountered an error during execution.
func NewChildProcessExecutionError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-nodejs-0008",
    Title:      "Child Process Execution Error",
    StatusCode: 500,
    ErrorCode:  "SNYK-OS-NODEJS-0008",
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

// NewNoValidPackageUpgradesError displays errors with the following description:
// The system attempted to find valid upgrades for the packages specified in the lock file, but none were available.
func NewNoValidPackageUpgradesError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-nodejs-0009",
    Title:      "No valid package upgrades",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-NODEJS-0009",
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

// NewNoDependencyUpdatesError displays errors with the following description:
// There are no available updates for the dependencies.
func NewNoDependencyUpdatesError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-nodejs-0010",
    Title:      "No dependency updates",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-NODEJS-0010",
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

// NewCouldNotParseJSONFileError displays errors with the following description:
// An error occurred while attempting to parse a JSON file.
func NewCouldNotParseJSONFileError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-nodejs-0011",
    Title:      "Could not parse JSON file",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-NODEJS-0011",
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

// NewBase64EncodeError displays errors with the following description:
// An error occurred while attempting to perform Base64 encoding.
func NewBase64EncodeError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-nodejs-0012",
    Title:      "Could not Base64 encode",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-NODEJS-0012",
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

// NewBase64DecodeError displays errors with the following description:
// An error occurred while attempting to perform Base64 decoding.
func NewBase64DecodeError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-nodejs-0013",
    Title:      "Could not Base64 decode",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-NODEJS-0013",
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

// NewMissingSupportedFileError displays errors with the following description:
// Could not find supported file.
func NewMissingSupportedFileError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-nodejs-0014",
    Title:      "Missing supported file",
    StatusCode: 400,
    ErrorCode:  "SNYK-OS-NODEJS-0014",
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

// NewInvalidConfigurationError displays errors with the following description:
// The configuration parameter does not meet the expected data type. Please ensure the provided value is of the correct data type.
func NewInvalidConfigurationError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-nodejs-0015",
    Title:      "Invalid configuration",
    StatusCode: 400,
    ErrorCode:  "SNYK-OS-NODEJS-0015",
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

// NewUnsupportedRequirementsFileError displays errors with the following description:
// The provided requirements file is not supported by Snyk for Python.
//
// Read more:
// - https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/python
func NewUnsupportedRequirementsFileError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-pip-0001",
    Title:      "Unsupported manifest file type for remediation",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-PIP-0001",
    Classification: "UNSUPPORTED",
    Links: []string{
      "https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/python",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewTooManyManifestFilesError displays errors with the following description:
// Too many manifest files were provided in the request body.
func NewTooManyManifestFilesError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-pip-0002",
    Title:      "Received more manifests than expected",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-PIP-0002",
    Classification: "UNSUPPORTED",
    Links: []string{},
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewFailedToApplyDependencyUpdatesError displays errors with the following description:
// An error occurred while updating dependencies.
func NewFailedToApplyDependencyUpdatesError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-pip-0003",
    Title:      "Failed to apply dependency updates",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-PIP-0003",
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

// NewNoMatchingDistributionError displays errors with the following description:
// One or more of the packages require a different Python version.
func NewNoMatchingDistributionError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/more-info/error-catalog#snyk-os-pip-0004",
    Title:      "No matching distribution found for one or more of the packages",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-PIP-0004",
    Classification: "UNSUPPORTED",
    Links: []string{},
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

