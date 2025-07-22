/*
 * © 2025 Snyk Limited
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-0001",
    Title:      "Unable to parse manifest file",
    Description: "The provided manifest file could not be parsed as it has invalid syntax or does not match the expected schema. Review the manifest file, then try again.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-0002",
    Title:      "Unable to parse lock file",
    Description: "The provided lock file could not be parsed as it has invalid syntax or does not match the expected schema. Review the lock file, then try again.",
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
// - https://support.snyk.io/s/article/Could-not-determine-version-for-dependencies
func NewUnknownDependencyVersionError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-0003",
    Title:      "Unknown dependency version",
    Description: "Dependency version could not be resolved.",
    StatusCode: 404,
    ErrorCode:  "SNYK-OS-0003",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://support.snyk.io/s/article/Could-not-determine-version-for-dependencies",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-0004",
    Title:      "Missing required request header",
    Description: "The server encountered a request that is missing a mandatory request header.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-0005",
    Title:      "Payload missing required elements",
    Description: "The server could not process the request.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-0006",
    Title:      "Files cannot be processed",
    Description: "The dependency service could not process the files.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-0007",
    Title:      "Cannot get file from source",
    Description: "Could not get the file from the source URL.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-0008",
    Title:      "Missing environment variable",
    Description: "The server encountered a critical operation that requires a specific environment variable, but the variable is not set or is not accessible within the current environment.",
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

// NewBrokeredConnectionNotSupportedError displays errors with the following description:
// The service encountered a permissions or credentials error most likely related to an import through a brokered connection for a scanner that does not yet support that.
func NewBrokeredConnectionNotSupportedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-0009",
    Title:      "Brokered connections not currently supported",
    Description: "The service encountered a permissions or credentials error most likely related to an import through a brokered connection for a scanner that does not yet support that.",
    StatusCode: 500,
    ErrorCode:  "SNYK-OS-0009",
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

// NewGitCloneFailedError displays errors with the following description:
// We encountered a fatal error from Git while trying to clone your code using your provided credentials. Please verify that:
// 
// * Your provided credentials are correct or not scoped too narrowly.
// * The branch you've asked us to clone exists.
// * The repository you've provided is accessible from the internet and you are not connected through a broker.
// 
// And try the operation again.
func NewGitCloneFailedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-0010",
    Title:      "Snyk failed to clone your repository",
    Description: "We encountered a fatal error from Git while trying to clone your code using your provided credentials. Please verify that:\n\n* Your provided credentials are correct or not scoped too narrowly.\n* The branch you've asked us to clone exists.\n* The repository you've provided is accessible from the internet and you are not connected through a broker.\n\nAnd try the operation again.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-0010",
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

// NewUnsupportedManifestFileError displays errors with the following description:
// The provided manifest file is not supported by Snyk for .NET.
//
// Read more:
// - https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/.net
func NewUnsupportedManifestFileError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-dotnet-0001",
    Title:      "Unsupported manifest file type for remediation",
    Description: "The provided manifest file is not supported by Snyk for .NET.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-dotnet-0002",
    Title:      "Target framework not supported",
    Description: "The provided manifest file defines a `<TargetFramework>` or `<TargetFrameworks>` that is not currently supported by Snyk's .NET scanning solution.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-dotnet-0003",
    Title:      "Your C# code is missing a static Main function",
    Description: "This error occurs when no static Main method with a correct signature is found in the code that produces an executable file. \nIt also occurs if the entry point function, `Main`, is defined with the wrong case, such as lower-case main.\n\nIn order to fix this issue, ensure that your program has a .cs file that contains a main function, such as\n```c#\nnamespace Example\n{\n    class Program\n    {\n        static void Main(string[] args)\n        {\n            Console.WriteLine(\"hello world\");\n        }\n    }\n}\n```",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-dotnet-0004",
    Title:      "The dotnet CLI is unable to generate a self-contained binary",
    Description: "This error occurs when running `dotnet publish --sc --framework <your-target-framework>` fails to generate a \nself-contained binary. Snyk needs to run this command in order to adequately determine the dependency tree for your project. If this command fails, Snyk cannot continue.\n\nSteps to determine why this happened:\n\n* Checkout a clean version of your project in a temporary folder\n* Run `dotnet publish --sc --framework <your-target-framework> ` on your project, and confirm this step fails.\n\nIf this step is successful locally, it is possible that Snyk is running another version of the .NET SDK. To tell Snyk which version of the .NET SDK to use, consider using the [global.json](https://learn.microsoft.com/en-us/dotnet/core/tools/global-json) solution provided by Microsoft.",
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

// NewFailedToAccessPrivatePackageSourceError displays errors with the following description:
// This error occurs when running `dotnet restore` fails to access dependencies stored in a private package source that Snyk does not have access to. 
// 
// This means that your `.csproj` file or files refer to a dependency hosted on a private package store or Nuget Artifact Registry defined in your `NuGet.config` file, such as:
// 
// ```xml
// <?xml version="1.0" encoding="utf-8"?>
// <configuration>
//   <packageSources>
//     <clear />
//     <add key="AzureFeed" value="https://pkgs.dev.azure.com/your-org/_packaging/your-repo/nuget/v3/index.json" />
//     <add key="nuget.org" value="https://api.nuget.org/v3/index.json" />
//   </packageSources>
// </configuration>
// ```
// 
// In order to allow Snyk to access your private dependency package source, you must supply Snyk with a valid JSON object as a private registry token in the .NET language settings.
// 
// You can set up a connection to your private Nuget repository in your Snyk integration settings.
//
// Read more:
// - https://github.com/microsoft/artifacts-credprovider#environment-variables
func NewFailedToAccessPrivatePackageSourceError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-dotnet-0005",
    Title:      "The dotnet CLI was unable to restore from private package sources",
    Description: "This error occurs when running `dotnet restore` fails to access dependencies stored in a private package source that Snyk does not have access to. \n\nThis means that your `.csproj` file or files refer to a dependency hosted on a private package store or Nuget Artifact Registry defined in your `NuGet.config` file, such as:\n\n```xml\n<?xml version=\"1.0\" encoding=\"utf-8\"?>\n<configuration>\n  <packageSources>\n    <clear />\n    <add key=\"AzureFeed\" value=\"https://pkgs.dev.azure.com/your-org/_packaging/your-repo/nuget/v3/index.json\" />\n    <add key=\"nuget.org\" value=\"https://api.nuget.org/v3/index.json\" />\n  </packageSources>\n</configuration>\n```\n\nIn order to allow Snyk to access your private dependency package source, you must supply Snyk with a valid JSON object as a private registry token in the .NET language settings.\n\nYou can set up a connection to your private Nuget repository in your Snyk integration settings.",
    StatusCode: 401,
    ErrorCode:  "SNYK-OS-DOTNET-0005",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://github.com/microsoft/artifacts-credprovider#environment-variables",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewMissingMSBuildConditionError displays errors with the following description:
// The `dotnet` tool was unable to locate the `.targets`, `.csproj` or `.props` file responsible for one or more MSBuild conditions in your project file.
// 
// The tool encountered an error like 
// ```
// /path/to/file/project.csproj(33,13): error MSB4100: Expected "$(SomeCondition)" to evaluate to a boolean instead of "", in condition "!$(SomeCondition)".
// ```
// 
// This means the condition definition is missing in the project file that is currently being restored and in any project linked to it from there.      
// 
// Snyk can scan only the project files accessible in the current repository or the private dependencies available to Snyk.
// 
// For example, if your code has the following structure:
// 
// ```title=project.targets
// <Project>
//   <PropertyGroup>
//     <SomeCondition Condition="'$(SomeCondition)' == ''">false</SomeCondition>
//   </PropertyGroup>
// </Project>
// ```
// 
// And
// 
// ```title=project.csproj
// <Project Sdk='Microsoft.NET.Sdk'>
//   <Import Project='..\external-libraries\some-library\project.targets' />
//   <PropertyGroup>
//     <TargetFrameworks>net8.0</TargetFrameworks>
//   </PropertyGroup>
//   <ItemGroup Condition='!$(SomeCondition)'>
//     <PackageReference Include='Newtonsoft.Json' Version='13.0.3' />
//   </ItemGroup>
// </Project>
// ```
// 
// And `external-libraries` is not a part of your repository currently being scanned, Snyk is not able to find it.
// 
// This error occurs when your code depends on external libraries that are added to or generated from your source code using external tools unknown to Snyk or as part of a build step in your build or a deployment pipeline.
//
// Read more:
// - https://learn.microsoft.com/en-us/visualstudio/msbuild/msbuild-conditional-constructs
func NewMissingMSBuildConditionError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-dotnet-0006",
    Title:      "Missing MSBuild Condition Construct in project file",
    Description: "The `dotnet` tool was unable to locate the `.targets`, `.csproj` or `.props` file responsible for one or more MSBuild conditions in your project file.\n\nThe tool encountered an error like \n```\n/path/to/file/project.csproj(33,13): error MSB4100: Expected \"$(SomeCondition)\" to evaluate to a boolean instead of \"\", in condition \"!$(SomeCondition)\".\n```\n\nThis means the condition definition is missing in the project file that is currently being restored and in any project linked to it from there.      \n\nSnyk can scan only the project files accessible in the current repository or the private dependencies available to Snyk.\n\nFor example, if your code has the following structure:\n\n```title=project.targets\n<Project>\n  <PropertyGroup>\n    <SomeCondition Condition=\"'$(SomeCondition)' == ''\">false</SomeCondition>\n  </PropertyGroup>\n</Project>\n```\n\nAnd\n\n```title=project.csproj\n<Project Sdk='Microsoft.NET.Sdk'>\n  <Import Project='..\\external-libraries\\some-library\\project.targets' />\n  <PropertyGroup>\n    <TargetFrameworks>net8.0</TargetFrameworks>\n  </PropertyGroup>\n  <ItemGroup Condition='!$(SomeCondition)'>\n    <PackageReference Include='Newtonsoft.Json' Version='13.0.3' />\n  </ItemGroup>\n</Project>\n```\n\nAnd `external-libraries` is not a part of your repository currently being scanned, Snyk is not able to find it.\n\nThis error occurs when your code depends on external libraries that are added to or generated from your source code using external tools unknown to Snyk or as part of a build step in your build or a deployment pipeline.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-DOTNET-0006",
    Classification: "UNSUPPORTED",
    Links: []string{
      "https://learn.microsoft.com/en-us/visualstudio/msbuild/msbuild-conditional-constructs",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewNoTargetFrameworksFoundError displays errors with the following description:
// Snyk was unable to detect any `<TargetFramework>`s in the supplied manifest files. 
// 
// If you are using `Directory.Build.props` files to determine the target framework, ensure that it is named as such. Due to performance considerations on the customer's SCM network, Snyk does not perform case-insensitive searches for `.props` files.
//
// Read more:
// - https://learn.microsoft.com/en-us/visualstudio/msbuild/customize-by-directory?view=vs-2022#directorybuildprops-and-directorybuildtargets
func NewNoTargetFrameworksFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-dotnet-0007",
    Title:      "No target frameworks found in manifest files",
    Description: "Snyk was unable to detect any `<TargetFramework>`s in the supplied manifest files. \n\nIf you are using `Directory.Build.props` files to determine the target framework, ensure that it is named as such. Due to performance considerations on the customer's SCM network, Snyk does not perform case-insensitive searches for `.props` files.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-DOTNET-0007",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://learn.microsoft.com/en-us/visualstudio/msbuild/customize-by-directory?view=vs-2022#directorybuildprops-and-directorybuildtargets",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewOutdatedSDKVersionRequestedError displays errors with the following description:
// Snyk supports the latest channels of .NET which is currently [supported by Microsoft](https://dotnet.microsoft.com/en-us/download/dotnet), but does **not** guarantee to support all SDK versions within each currently supported channel.
// 
// Within the supported channels, Snyk aims to support most, if not all, of the SDK versions currently released under the **newest** of the channels.
// 
// If the channels currently supported by Microsoft are `8.0`, `7.0` and `6.0`, Snyk **will** support all of the *latest* SDKs released for these channels.
// 
// If the SDK versions released under `8.0.3` are: `8.0.203`, `8.0.202` and `8.0.103`, Snyk **cannot** guarantee to support *all* of them, but makes an effort to do so. Snyk **will** support the latest of the SDK versions currently released by Microsoft. 
// 
// If channel `8.0` is the newest channel currently supported, Snyk **cannot** guarantee that multiple, specific SDK versions for older, still supported channels such as .NET 6. 
// 
// ### Example support matrix
// 
// If:
// 
// * .NET channels currently supported by Microsoft are `.NET 8.0`,  `.NET 7.0` and  `.NET 6.0`
// * Newest SDK version under `.NET 8.0` is `8.0.203`
// 
// Then:
// 
// | Channel |              SDK             | End-of-Life |  Supported  |
// |:-------:|:----------------------------:|:-----------:|:-----------:|
// |   8.0   | 8.0.203  (latest in channel) |      No     |     Yes     |
// |   8.0   |            8.0.202           |      No     |     Yes     |
// |   8.0   |            8.0.103           |      No     |     Yes     |
// |         |             (...)            |             |             |
// |   7.0   | 7.0.407  (latest in channel) |      No     |     Yes     |
// |   7.0   |            7.0.314           |      No     |      No     |
// |         |             (...)            |             |             |
// |   6.0   |            6.0.420           |      No     |     Yes     |
// |   6.0   |            6.0.128           |      No     |      No     |
// |         |             (..)             |             |             |
// |   5.0   |  5.0.408 (latest in channel) |     Yes     |      No     |
// |   5.0   |            5.0.214           |     Yes     |      No     |
// |         |             (..)             |             |             |
// 
// ### Workarounds
// 
// This limitation can lead to scan failures for customers that are pinning SDK versions in their `global.json` files without a [rollForward](https://learn.microsoft.com/en-us/dotnet/core/tools/global-json#rollforward) directive, such as:
// ```json
// {
//   "sdk": {
//     "version": "6.0.101"
//   }
// }
// ```
// Since as `6.0` is not the newest .NET channel. 
// 
// To work around this issue, we recommend that customers employ some flexibility in their `global.json` file by employing the `rollFoward` directive to be `latestMajor`, as such:
// ```json
// {
//   "sdk": {
//     "version": "6.0.101",
//     "rollForward": "latestMajor"
//   }
// }
// ```
// 
// Which will allow Snyk to scan your code using a newer version of the SDK, despite your version pinning.
//
// Read more:
// - https://versionsof.net/core/
// - https://dotnet.microsoft.com/en-us/download/dotnet
// - https://learn.microsoft.com/en-us/dotnet/core/tools/global-json#rollforward
func NewOutdatedSDKVersionRequestedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-dotnet-0008",
    Title:      "Your global.json is targeting an outdated SDK version",
    Description: "Snyk supports the latest channels of .NET which is currently [supported by Microsoft](https://dotnet.microsoft.com/en-us/download/dotnet), but does **not** guarantee to support all SDK versions within each currently supported channel.\n\nWithin the supported channels, Snyk aims to support most, if not all, of the SDK versions currently released under the **newest** of the channels.\n\nIf the channels currently supported by Microsoft are `8.0`, `7.0` and `6.0`, Snyk **will** support all of the *latest* SDKs released for these channels.\n\nIf the SDK versions released under `8.0.3` are: `8.0.203`, `8.0.202` and `8.0.103`, Snyk **cannot** guarantee to support *all* of them, but makes an effort to do so. Snyk **will** support the latest of the SDK versions currently released by Microsoft. \n\nIf channel `8.0` is the newest channel currently supported, Snyk **cannot** guarantee that multiple, specific SDK versions for older, still supported channels such as .NET 6. \n\n### Example support matrix\n\nIf:\n\n* .NET channels currently supported by Microsoft are `.NET 8.0`,  `.NET 7.0` and  `.NET 6.0`\n* Newest SDK version under `.NET 8.0` is `8.0.203`\n\nThen:\n\n| Channel |              SDK             | End-of-Life |  Supported  |\n|:-------:|:----------------------------:|:-----------:|:-----------:|\n|   8.0   | 8.0.203  (latest in channel) |      No     |     Yes     |\n|   8.0   |            8.0.202           |      No     |     Yes     |\n|   8.0   |            8.0.103           |      No     |     Yes     |\n|         |             (...)            |             |             |\n|   7.0   | 7.0.407  (latest in channel) |      No     |     Yes     |\n|   7.0   |            7.0.314           |      No     |      No     |\n|         |             (...)            |             |             |\n|   6.0   |            6.0.420           |      No     |     Yes     |\n|   6.0   |            6.0.128           |      No     |      No     |\n|         |             (..)             |             |             |\n|   5.0   |  5.0.408 (latest in channel) |     Yes     |      No     |\n|   5.0   |            5.0.214           |     Yes     |      No     |\n|         |             (..)             |             |             |\n\n### Workarounds\n\nThis limitation can lead to scan failures for customers that are pinning SDK versions in their `global.json` files without a [rollForward](https://learn.microsoft.com/en-us/dotnet/core/tools/global-json#rollforward) directive, such as:\n```json\n{\n  \"sdk\": {\n    \"version\": \"6.0.101\"\n  }\n}\n```\nSince as `6.0` is not the newest .NET channel. \n\nTo work around this issue, we recommend that customers employ some flexibility in their `global.json` file by employing the `rollFoward` directive to be `latestMajor`, as such:\n```json\n{\n  \"sdk\": {\n    \"version\": \"6.0.101\",\n    \"rollForward\": \"latestMajor\"\n  }\n}\n```\n\nWhich will allow Snyk to scan your code using a newer version of the SDK, despite your version pinning.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-DOTNET-0008",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://versionsof.net/core/",
      "https://dotnet.microsoft.com/en-us/download/dotnet",
      "https://learn.microsoft.com/en-us/dotnet/core/tools/global-json#rollforward",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewProjectSkippedAndNotFoundError displays errors with the following description:
// While attempting to build your solution for scanning, the `dotnet` SDK was unable to restore one or more projects referenced in your manifest files.
// 
// Please note that Snyk runs these builds on a **case-sensitive** filesystem, meaning that `<ProjectReference>../src/NS.Project.csproj</ProjectReference>` and `<ProjectReference>../src/ns.project.csproj</ProjectReference>` are not referring to the same thing.
// 
// This can present itself as a problem for customers that are using Mac or Windows build pipeline where file systems are not case-sensitive. In this case, verify you're referring to the right manifest files and check the Snyk import logs for more details.
//
// Read more:
// - https://learn.microsoft.com/en-us/dotnet/csharp/language-reference/compiler-messages/assembly-references#missing-references
func NewProjectSkippedAndNotFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-dotnet-0009",
    Title:      "Project failed to build due to missing type or namespace references",
    Description: "While attempting to build your solution for scanning, the `dotnet` SDK was unable to restore one or more projects referenced in your manifest files.\n\nPlease note that Snyk runs these builds on a **case-sensitive** filesystem, meaning that `<ProjectReference>../src/NS.Project.csproj</ProjectReference>` and `<ProjectReference>../src/ns.project.csproj</ProjectReference>` are not referring to the same thing.\n\nThis can present itself as a problem for customers that are using Mac or Windows build pipeline where file systems are not case-sensitive. In this case, verify you're referring to the right manifest files and check the Snyk import logs for more details.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-DOTNET-0009",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://learn.microsoft.com/en-us/dotnet/csharp/language-reference/compiler-messages/assembly-references#missing-references",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewNugetDependenciesSpaceLimitExceededError displays errors with the following description:
// The total size of the downloaded Nuget dependencies in the manifest files exceeds the 10GB limit.
// This often happens due to the large number or the large size of Nuget dependencies. 
//
// Read more:
// - https://docs.snyk.io/supported-languages-package-managers-and-frameworks/.net/improved-.net-scanning#limitations-on-improved-.net-scanning-for-scm-integrations
func NewNugetDependenciesSpaceLimitExceededError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-dotnet-0010",
    Title:      "The 10 GB space limit for downloaded Nuget dependencies has been exceeded",
    Description: "The total size of the downloaded Nuget dependencies in the manifest files exceeds the 10GB limit.\nThis often happens due to the large number or the large size of Nuget dependencies. ",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-DOTNET-0010",
    Classification: "UNSUPPORTED",
    Links: []string{
      "https://docs.snyk.io/supported-languages-package-managers-and-frameworks/.net/improved-.net-scanning#limitations-on-improved-.net-scanning-for-scm-integrations",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewRestoreFailedError displays errors with the following description:
// This error occurs when running `dotnet restore <path-to-csproj>` fails to generate a 
// manifest of the resolved dependencies. Snyk needs to run this command in order to adequately determine the dependency tree for your project. 
// If this command fails, Snyk cannot continue.
// 
// Steps to determine why this happened:
// 
// * Checkout a clean version of your project in a temporary folder
// * Run `dotnet restore <path-to-csproj> ` on your project, and confirm this step fails.
// 
// If this step is successful locally, it is possible that Snyk is running another version of the .NET SDK. To tell Snyk which version of the .NET SDK to use, consider using the [global.json](https://learn.microsoft.com/en-us/dotnet/core/tools/global-json) solution provided by Microsoft.
//
// Read more:
// - https://learn.microsoft.com/en-us/dotnet/core/tools/sdk-errors/
// - https://learn.microsoft.com/en-us/dotnet/core/tools/global-json
// - https://github.com/snyk/snyk-nuget-plugin/blob/885486aa656c28d3db465c8d22710770d5cc6773/lib/nuget-parser/cli/dotnet.ts#L50
func NewRestoreFailedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-dotnet-0011",
    Title:      "The dotnet CLI is unable to download and install all the required dependencies",
    Description: "This error occurs when running `dotnet restore <path-to-csproj>` fails to generate a \nmanifest of the resolved dependencies. Snyk needs to run this command in order to adequately determine the dependency tree for your project. \nIf this command fails, Snyk cannot continue.\n\nSteps to determine why this happened:\n\n* Checkout a clean version of your project in a temporary folder\n* Run `dotnet restore <path-to-csproj> ` on your project, and confirm this step fails.\n\nIf this step is successful locally, it is possible that Snyk is running another version of the .NET SDK. To tell Snyk which version of the .NET SDK to use, consider using the [global.json](https://learn.microsoft.com/en-us/dotnet/core/tools/global-json) solution provided by Microsoft.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-DOTNET-0011",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://learn.microsoft.com/en-us/dotnet/core/tools/sdk-errors/",
      "https://learn.microsoft.com/en-us/dotnet/core/tools/global-json",
      "https://github.com/snyk/snyk-nuget-plugin/blob/885486aa656c28d3db465c8d22710770d5cc6773/lib/nuget-parser/cli/dotnet.ts#L50",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-go-0001",
    Title:      "Failed to access private module",
    Description: "Snyk could not access the private modules within your go.mod files.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-go-0002",
    Title:      "Go mod file not found",
    Description: "A go.mod file was not found in the current directory or any parent directory.",
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
//
// Deprecated: This error has been moved to a more generalized namespace to avoid repetition.  Use SNYK-OS-8004 instead.
func NewSsoReAuthRequiredError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-go-0003",
    Title:      "OAuth re-authorization required",
    Description: "Your code is cloned on an isolated environment using Git as it is required by Snyk to analyze its dependencies.\n\nYour Organization has enabled or enforced SAML SSO after you authorized Snyk to access your code, and a re-authentication is therefore required.\n\nThe error you're seeing is usually reproducible by attempting to do a `git clone` of your repository with incorrectly configured credentials.\nVerify your authentication configuration with your Git cloud provider and try again.",
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
// This error means that you need some cleanup, (such as `go mod tidy`) or your project deployment process contains a code generation step such as `protobuf` or similar that is not currently supported by Snyk. 
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-go-0004",
    Title:      "Your project repository is missing required files",
    Description: "Generating the dependency graph requires Snyk to run go list `go list -deps -json` inside the project. If the operation fails, creating a full dependency graph cannot continue.  \n\nThis error means that you need some cleanup, (such as `go mod tidy`) or your project deployment process contains a code generation step such as `protobuf` or similar that is not currently supported by Snyk. \n\nTo verify if this is the case, clone your project in a clean environment, run go list `go list -deps -json` and verify whether the operation fails. \n\nIf Snyk cannot process your code successfully, insert the Snyk CLI as part of your deployment pipeline.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-go-0005",
    Title:      "Your project repository has inconsistent vendoring information",
    Description: "Generating the dependency graph requires Snyk to run `go list -deps -json` inside the project. If the operation fails, creating a full dependency graph cannot continue.  \n\nThis error means that there is inconsistency between your `vendor/modules.txt` file and your `go.mod` file. To remediate, you need to:\n\n* `go mod vendor`\n* `go mod tidy`\n\nNext, commit those changes to your repo. Snyk does not manipulate with your code on our end by design, which is why this is not done automatically.\n\nTo verify if this is the case, clone your project in a clean environment, run go list `go list -deps -json` and verify whether the operation fails. \nThen try and run the above mentioned commands and see if your SCM system reports changes in files.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-go-0006",
    Title:      "Unsupported external file generation",
    Description: "Snyk currently does not support external file generation in your project. This limitation is due to Snyk's lack of visibility into the third-party generator tools you may be using and the specific commands required to generate these files.\n\nSnyk can only work with the files available in your repository and does not have insight into the generation process for external files.",
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

// NewUnableToAccessPrivateDepsError displays errors with the following description:
// The Go tool encountered a `DepsError` while trying to download a private dependency. Private repositories that are not accessible to the public internet and are not available on the official Go proxy mirror are cloned with a version control system and built on demand. 
// This requires the VCS to have the correct access rights to that repository.
// 
// Snyk supports private repositories that are hosted in the same Organization and on the same Project that is scanned for vulnerabilities. The authentication to the private repository is the same as the authentication used to integrate that repository with Snyk. 
// 
// This error appears when the authorization credentials do not allow access to the requested private dependency. 
//
// Read more:
// - https://go.dev/ref/mod#vcs
func NewUnableToAccessPrivateDepsError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-go-0007",
    Title:      "Unable to access private dependencies",
    Description: "The Go tool encountered a `DepsError` while trying to download a private dependency. Private repositories that are not accessible to the public internet and are not available on the official Go proxy mirror are cloned with a version control system and built on demand. \nThis requires the VCS to have the correct access rights to that repository.\n\nSnyk supports private repositories that are hosted in the same Organization and on the same Project that is scanned for vulnerabilities. The authentication to the private repository is the same as the authentication used to integrate that repository with Snyk. \n\nThis error appears when the authorization credentials do not allow access to the requested private dependency. ",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-GO-0007",
    Classification: "UNSUPPORTED",
    Links: []string{
      "https://go.dev/ref/mod#vcs",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewUnableToUseCredentialsError displays errors with the following description:
// The Go tool encountered a permissions error while fetching one of the private dependencies. Ensure that the integration token you used to sign in to Snyk is properly configured so that Snyk can access the private dependencies.
// 
// The Snyk Go integration only supports private dependencies that are used inside the same Organization as the Project you are scanning.
// 
// This error appears when Snyk is unable to properly access the authorization credentials for the requested private dependency. 
func NewUnableToUseCredentialsError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-go-0008",
    Title:      "Unable to fetch private dependencies",
    Description: "The Go tool encountered a permissions error while fetching one of the private dependencies. Ensure that the integration token you used to sign in to Snyk is properly configured so that Snyk can access the private dependencies.\n\nThe Snyk Go integration only supports private dependencies that are used inside the same Organization as the Project you are scanning.\n\nThis error appears when Snyk is unable to properly access the authorization credentials for the requested private dependency. ",
    StatusCode: 401,
    ErrorCode:  "SNYK-OS-GO-0008",
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

// NewToolchainNotAvailableError displays errors with the following description:
// Could not download Go toolchain.
func NewToolchainNotAvailableError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-go-0009",
    Title:      "Toolchain not available",
    Description: "Could not download Go toolchain.",
    StatusCode: 500,
    ErrorCode:  "SNYK-OS-GO-0009",
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

// NewGolangSpaceLimitExceededError displays errors with the following description:
// The total size of the downloaded Golang dependencies in the manifest file exceeds the 10GB limit.
// This often happens due to the large size or large number of Golang dependencies.
// 
// Currently this is a product limitation for SCM. As a workaround, use the 'snyk monitor' command via Snyk CLI.
//
// Read more:
// - https://docs.snyk.io/snyk-cli/commands/monitor
func NewGolangSpaceLimitExceededError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-go-0010",
    Title:      "The 10 GB space limit for downloaded Golang dependencies has been exceeded",
    Description: "The total size of the downloaded Golang dependencies in the manifest file exceeds the 10GB limit.\nThis often happens due to the large size or large number of Golang dependencies.\n\nCurrently this is a product limitation for SCM. As a workaround, use the 'snyk monitor' command via Snyk CLI.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-GO-0010",
    Classification: "UNSUPPORTED",
    Links: []string{
      "https://docs.snyk.io/snyk-cli/commands/monitor",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewGolangNoSecureProtocolFoundError displays errors with the following description:
// The Go toolchain could not find a secure protocol (HTTPS) to access the repository.
// 
// This error typically occurs when the repository URL is not configured to use a secure protocol, or the necessary credentials for accessing the repository securely are not provided.
// 
// Ensure that the repository URL uses a secure protocol (https://) and verify that the necessary authentication credentials (such as HTTPS credentials) are correctly configured and accessible by the Go toolchain.
//  
// Note: SSH is not supported.
func NewGolangNoSecureProtocolFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-go-0011",
    Title:      "No secure protocol found for repository",
    Description: "The Go toolchain could not find a secure protocol (HTTPS) to access the repository.\n\nThis error typically occurs when the repository URL is not configured to use a secure protocol, or the necessary credentials for accessing the repository securely are not provided.\n\nEnsure that the repository URL uses a secure protocol (https://) and verify that the necessary authentication credentials (such as HTTPS credentials) are correctly configured and accessible by the Go toolchain.\n \nNote: SSH is not supported.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-GO-0011",
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

// NewGolangConnectionResetByPeerError displays errors with the following description:
// The Go toolchain encountered a connection reset error while trying to access the repository.
// 
// This error typically occurs when the connection to the repository is unexpectedly closed by the remote server. This can be due to network issues, server configuration, or other transient problems.
// 
// Try to reimport the project, if that does not work, please reach out to Snyk support.
func NewGolangConnectionResetByPeerError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-go-0012",
    Title:      "Connection reset by peer",
    Description: "The Go toolchain encountered a connection reset error while trying to access the repository.\n\nThis error typically occurs when the connection to the repository is unexpectedly closed by the remote server. This can be due to network issues, server configuration, or other transient problems.\n\nTry to reimport the project, if that does not work, please reach out to Snyk support.",
    StatusCode: 500,
    ErrorCode:  "SNYK-OS-GO-0012",
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

// NewGolangInvalidZipFileError displays errors with the following description:
// The Go toolchain encountered an error while trying to download and extract a Go package version. The downloaded file is not a valid zip file.
// 
// This error typically occurs when there is an issue with the downloaded file, such as corruption or an incomplete download.
// 
// If the package that fails is a private package that's hosted in a private network, make sure the network is up and running and retry the import. If the issue persists, verify the integrity of the downloaded file and check for any issues with the source of the download.
// 
// If the issue persists, please reach out to Snyk support.
func NewGolangInvalidZipFileError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-go-0013",
    Title:      "Invalid zip file",
    Description: "The Go toolchain encountered an error while trying to download and extract a Go package version. The downloaded file is not a valid zip file.\n\nThis error typically occurs when there is an issue with the downloaded file, such as corruption or an incomplete download.\n\nIf the package that fails is a private package that's hosted in a private network, make sure the network is up and running and retry the import. If the issue persists, verify the integrity of the downloaded file and check for any issues with the source of the download.\n\nIf the issue persists, please reach out to Snyk support.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-GO-0013",
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

// NewGolangVersionMismatchError displays errors with the following description:
// The Go toolchain encountered a version mismatch error while trying to process the module.
// 
// This usually happens when the version of Go that was used in the go.mod file is not yet supported by Snyk.
// 
// We usually try and add support for new Golang versions shortly after a new one was released.
// 
// If the Go version used in the go.mod file is supported based on our Golang documentation, please reach out to Snyk support.
//
// Read more:
// - https://docs.snyk.io/supported-languages-package-managers-and-frameworks/go/go-for-open-source#go-for-snyk-open-source-support
func NewGolangVersionMismatchError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-go-0014",
    Title:      "Go version mismatch",
    Description: "The Go toolchain encountered a version mismatch error while trying to process the module.\n\nThis usually happens when the version of Go that was used in the go.mod file is not yet supported by Snyk.\n\nWe usually try and add support for new Golang versions shortly after a new one was released.\n\nIf the Go version used in the go.mod file is supported based on our Golang documentation, please reach out to Snyk support.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-GO-0014",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/supported-languages-package-managers-and-frameworks/go/go-for-open-source#go-for-snyk-open-source-support",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewGolangInvalidGoVersionError displays errors with the following description:
// The Go toolchain encountered an error while parsing the go.mod file. The specified Go version '__GOLANG_VERSION__' is invalid and must match the format 1.23.4.
// 
// This error typically occurs when the Go version in the go.mod file is not correctly formatted.
// 
// Ensure that the Go version specified in the go.mod file matches the required format (e.g., 1.23.4). Update the go.mod file with the correct Go version and try again.
func NewGolangInvalidGoVersionError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-go-0015",
    Title:      "Invalid Go version in go.mod",
    Description: "The Go toolchain encountered an error while parsing the go.mod file. The specified Go version '__GOLANG_VERSION__' is invalid and must match the format 1.23.4.\n\nThis error typically occurs when the Go version in the go.mod file is not correctly formatted.\n\nEnsure that the Go version specified in the go.mod file matches the required format (e.g., 1.23.4). Update the go.mod file with the correct Go version and try again.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-GO-0015",
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

// NewGolangDialTcpTimeoutError displays errors with the following description:
// The Go toolchain encountered a timeout error while trying to fetch a module/package.
// 
// This error typically occurs when the connection to the server hosting the module/package times out. This can be due to network issues, server configuration, or the server being unreachable.
// 
// Ensure that the server hosting the module is reachable and that there are no network issues. If the issue persists, check if there are any restrictions or firewalls blocking access to the server.
func NewGolangDialTcpTimeoutError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-go-0016",
    Title:      "Dial TCP timeout",
    Description: "The Go toolchain encountered a timeout error while trying to fetch a module/package.\n\nThis error typically occurs when the connection to the server hosting the module/package times out. This can be due to network issues, server configuration, or the server being unreachable.\n\nEnsure that the server hosting the module is reachable and that there are no network issues. If the issue persists, check if there are any restrictions or firewalls blocking access to the server.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-GO-0016",
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

// NewGolangHostKeyVerificationFailedError displays errors with the following description:
// The Go toolchain encountered an error while trying to fetch a module/package. The host key verification failed, which means the key used to connect to the remote repository could not be verified.
// 
// This error typically occurs when the key is not recognized or is incorrect. It can also happen if the remote repository's host key has changed.
// 
// Ensure that the key used to connect to the remote repository is correct and properly configured. You may need to update the known_hosts file to include the correct host key for the remote repository. Verify your repository access configuration and try again. 
//
// Read more:
// - https://docs.github.com/en/authentication/troubleshooting-ssh/error-host-key-verification-failed
func NewGolangHostKeyVerificationFailedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-go-0017",
    Title:      "Host key verification failed",
    Description: "The Go toolchain encountered an error while trying to fetch a module/package. The host key verification failed, which means the key used to connect to the remote repository could not be verified.\n\nThis error typically occurs when the key is not recognized or is incorrect. It can also happen if the remote repository's host key has changed.\n\nEnsure that the key used to connect to the remote repository is correct and properly configured. You may need to update the known_hosts file to include the correct host key for the remote repository. Verify your repository access configuration and try again. ",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-GO-0017",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.github.com/en/authentication/troubleshooting-ssh/error-host-key-verification-failed",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewGolangMissingModuleDeclarationError displays errors with the following description:
// The Go toolchain encountered an error while reading the go.mod file. The error indicates that the module declaration is missing.
// 
// This error typically occurs when the go.mod file does not specify the module path. To resolve this issue, you need to add a module declaration to the go.mod file.
// 
// To specify the module path, run the following command: go mod edit -module=example.com/mod
// 
// Update the go.mod file with the correct module path and try again.
func NewGolangMissingModuleDeclarationError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-go-0018",
    Title:      "Missing module declaration in go.mod",
    Description: "The Go toolchain encountered an error while reading the go.mod file. The error indicates that the module declaration is missing.\n\nThis error typically occurs when the go.mod file does not specify the module path. To resolve this issue, you need to add a module declaration to the go.mod file.\n\nTo specify the module path, run the following command: go mod edit -module=example.com/mod\n\nUpdate the go.mod file with the correct module path and try again.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-GO-0018",
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

// NewGolangModuleVersionConstraintNotMetError displays errors with the following description:
// The Go toolchain encountered a version constraint violation while resolving dependencies.
// 
// This error typically occurs when a module or dependency specifies a version requirement that is not satisfied by the current Go version in use. For example, a module may require "go >= 1.22.0", but your environment is running "go 1.21.0".
// 
// To resolve this, ensure your Go toolchain version meets the module's specified constraint. This might involve upgrading or downgrading your Go version from Settings -> Snyk Open Source -> Edit settings for Go, or finding an alternative version of the module that is compatible with your current Go environment.
func NewGolangModuleVersionConstraintNotMetError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-go-0019",
    Title:      "Go module version constraint not met",
    Description: "The Go toolchain encountered a version constraint violation while resolving dependencies.\n\nThis error typically occurs when a module or dependency specifies a version requirement that is not satisfied by the current Go version in use. For example, a module may require \"go >= 1.22.0\", but your environment is running \"go 1.21.0\".\n\nTo resolve this, ensure your Go toolchain version meets the module's specified constraint. This might involve upgrading or downgrading your Go version from Settings -> Snyk Open Source -> Edit settings for Go, or finding an alternative version of the module that is compatible with your current Go environment.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-GO-0019",
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

// NewMissingRequirementFromPomError displays errors with the following description:
// The required property is missing from the pom object.
func NewMissingRequirementFromPomError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-maven-0001",
    Title:      "Missing property",
    Description: "The required property is missing from the pom object.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-maven-0002",
    Title:      "Unable to resolve value for property",
    Description: "The targeted property could not be resolved with a valid value.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-maven-0003",
    Title:      "Unable to resolve version for property",
    Description: "The targeted property could not be resolved with a valid version.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-maven-0004",
    Title:      "Cyclic property detected in POM file",
    Description: "There is circular dependency among properties in the Maven project's configuration file (POM), preventing proper resolution and causing an error.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-maven-0005",
    Title:      "Error parsing the XML file",
    Description: "There is an error parsing the XML file. This could be referring to either pom.xml or maven-metadata.xml.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-maven-0006",
    Title:      "Invalid coordinates provided",
    Description: "The coordinates provided for a project were invalid.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-maven-0007",
    Title:      "Skipping group",
    Description: "Skipping a specific groupId starting due to remapped coordinates.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-maven-0008",
    Title:      "Pom file not found",
    Description: "The pom file was not found in Maven repository.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-maven-0009",
    Title:      "Missing project from POM",
    Description: "A project element is missing from POM.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-maven-0010",
    Title:      "Cannot resolve the target POM from the input XML",
    Description: "Cannot resolve the targeted POM from the input XML.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-maven-0011",
    Title:      "Cannot resolve the target POM from the repository",
    Description: "Cannot resolve the targeted POM from the repository.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-maven-0012",
    Title:      "Cannot get the build file repository",
    Description: "Cannot get the build file repository.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-maven-0013",
    Title:      "Unable to create hosted git info",
    Description: "Cannot create source URL.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-maven-0014",
    Title:      "No released version for versions range",
    Description: "There was no version released for the specified versions range.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-maven-0015",
    Title:      "Source is not supported",
    Description: "The source used is not supported by fetcher. The supported sources are: github, bitbucket, gitlab.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-maven-0016",
    Title:      "Timeout when processing the dependency tree",
    Description: "There was an timeout when processing the dependency tree.",
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

// NewCannotReachConfiguredRepositoryError displays errors with the following description:
// One or more of the Maven repositories configured under your organisations language settings cannot be reached.
// 
// This error can happen for a variety of reasons:
// 
// * If using broker it could be a misconfiguration in your broker client. Double check the username and password. 
// * It could be network connectivity between the broker client and Snyk or between the broker client and the configured repository, check your firewall rules.
// 
// In order to solve this issue, refer to the specific details of this error message to identify which repository is causing issues. 
//
// Read more:
// - https://docs.snyk.io/integrate-with-snyk/package-repository-integrations
func NewCannotReachConfiguredRepositoryError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-maven-0017",
    Title:      "Cannot reach one or more Maven repositories configured under your Snyk organisations language settings",
    Description: "One or more of the Maven repositories configured under your organisations language settings cannot be reached.\n\nThis error can happen for a variety of reasons:\n\n* If using broker it could be a misconfiguration in your broker client. Double check the username and password. \n* It could be network connectivity between the broker client and Snyk or between the broker client and the configured repository, check your firewall rules.\n\nIn order to solve this issue, refer to the specific details of this error message to identify which repository is causing issues. ",
    StatusCode: 404,
    ErrorCode:  "SNYK-OS-MAVEN-0017",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/integrate-with-snyk/package-repository-integrations",
    },
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-nodejs-0001",
    Title:      "No repository found for A NPM package",
    Description: "No repository found for the NPM package.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-nodejs-0002",
    Title:      "Could not parse NPM registry URL",
    Description: "Could not parse NPM registry URL.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-nodejs-0003",
    Title:      "Could not find a broker resolved URL",
    Description: "Could not find a broker resolved URL.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-nodejs-0004",
    Title:      "Unable to replace broker URL",
    Description: "Unable to replace all broker urls in lock file.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-nodejs-0005",
    Title:      "Bad NPM version",
    Description: "The NPM version is not supported.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-nodejs-0006",
    Title:      "Unknown blob encoding on Github",
    Description: "Unknown blob encoding on Github.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-nodejs-0007",
    Title:      "No result from forked process",
    Description: "No result from forked process.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-nodejs-0008",
    Title:      "Child Process Execution Error",
    Description: "The child process encountered an error during execution.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-nodejs-0009",
    Title:      "No valid package upgrades",
    Description: "The system attempted to find valid upgrades for the packages specified in the lock file, but none were available.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-nodejs-0010",
    Title:      "No dependency updates",
    Description: "There are no available updates for the dependencies.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-nodejs-0011",
    Title:      "Could not parse JSON file",
    Description: "An error occurred while attempting to parse a JSON file.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-nodejs-0012",
    Title:      "Could not Base64 encode",
    Description: "An error occurred while attempting to perform Base64 encoding.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-nodejs-0013",
    Title:      "Could not Base64 decode",
    Description: "An error occurred while attempting to perform Base64 decoding.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-nodejs-0014",
    Title:      "Missing supported file",
    Description: "Could not find supported file.",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-nodejs-0015",
    Title:      "Invalid configuration",
    Description: "The configuration parameter does not meet the expected data type. Please ensure the provided value is of the correct data type.",
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

// NewPnpmOutOfSyncError displays errors with the following description:
// Sometimes a project may become out of sync between the lockfile and the manifest file. This might happen if the package.json is modified or updated but the pnpm-lock.yaml is not. 
// 
// This can be resolved by ensuring the lockfile and manifest file are correctly synced, by executing pnpm install.
// 
// In some cases, it may be necessary to delete the node_modules folder and the pnpm-lock.yaml and run pnpm install again to force a full reinstall. 
//
// Read more:
// - https://support.snyk.io/s/article/Out-of-sync-manifest--lockfile-in-the-project
func NewPnpmOutOfSyncError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-nodejs-0016",
    Title:      "Out of Sync Error",
    Description: "Sometimes a project may become out of sync between the lockfile and the manifest file. This might happen if the package.json is modified or updated but the pnpm-lock.yaml is not. \n\nThis can be resolved by ensuring the lockfile and manifest file are correctly synced, by executing pnpm install.\n\nIn some cases, it may be necessary to delete the node_modules folder and the pnpm-lock.yaml and run pnpm install again to force a full reinstall. ",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-NODEJS-0016",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://support.snyk.io/s/article/Out-of-sync-manifest--lockfile-in-the-project",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewPnpmUnsupportedLockfileVersionError displays errors with the following description:
// The lockfile version is not supported. Supported lockfile versions for pnpm include v5 and v6.
func NewPnpmUnsupportedLockfileVersionError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-nodejs-0017",
    Title:      "Unsupported pnpm lockfile version",
    Description: "The lockfile version is not supported. Supported lockfile versions for pnpm include v5 and v6.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-NODEJS-0017",
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

// NewYarnPackageNotFoundError displays errors with the following description:
// Snyk could not find the package in the Yarn registry.
func NewYarnPackageNotFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-nodejs-0019",
    Title:      "Yarn package not found",
    Description: "Snyk could not find the package in the Yarn registry.",
    StatusCode: 404,
    ErrorCode:  "SNYK-OS-NODEJS-0019",
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

// NewUnableToReachRegistryError displays errors with the following description:
// Snyk could not reach the node package registry.
func NewUnableToReachRegistryError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-nodejs-0020",
    Title:      "Unable to reach package registry",
    Description: "Snyk could not reach the node package registry.",
    StatusCode: 503,
    ErrorCode:  "SNYK-OS-NODEJS-0020",
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

// NewOutdatedYarnLockFileError displays errors with the following description:
// The lock file is outdated. Update the lock file and try again.
func NewOutdatedYarnLockFileError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-nodejs-0021",
    Title:      "Lock file is outdated",
    Description: "The lock file is outdated. Update the lock file and try again.",
    StatusCode: 409,
    ErrorCode:  "SNYK-OS-NODEJS-0021",
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

// NewPermissionDeniedError displays errors with the following description:
// Snyk does not have sufficient permissions to access the repository, or the repository does not exist.
func NewPermissionDeniedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-nodejs-0022",
    Title:      "Unable to read from remote repository",
    Description: "Snyk does not have sufficient permissions to access the repository, or the repository does not exist.",
    StatusCode: 401,
    ErrorCode:  "SNYK-OS-NODEJS-0022",
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

// NewUnsupportedRequirementsFileError displays errors with the following description:
// The provided requirements file is not supported by Snyk for Python.
//
// Read more:
// - https://docs.snyk.io/scan-applications/supported-languages-and-frameworks/python
func NewUnsupportedRequirementsFileError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-python-0001",
    Title:      "Unsupported manifest file type for remediation",
    Description: "The provided requirements file is not supported by Snyk for Python.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-PYTHON-0001",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-python-0002",
    Title:      "Received more manifests than expected",
    Description: "Too many manifest files were provided in the request body.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-PYTHON-0002",
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
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-python-0003",
    Title:      "Failed to apply dependency updates",
    Description: "An error occurred while updating dependencies.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-PYTHON-0003",
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

// NewPythonPackageNotFoundError displays errors with the following description:
// A package listed in the manifest file cannot be found in the Python Package Index(PyPI).
// Make sure all packages included in the manifest file are public existing ones.
func NewPythonPackageNotFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-python-0004",
    Title:      "Python package not found",
    Description: "A package listed in the manifest file cannot be found in the Python Package Index(PyPI).\nMake sure all packages included in the manifest file are public existing ones.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-PYTHON-0004",
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

// NewSyntaxIssuesError displays errors with the following description:
// The manifest file has syntax issues like incorrect package names or unsupported characters.
// Make sure the manifest file follows the syntax stardards and can be installed locally as well.
func NewSyntaxIssuesError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-python-0005",
    Title:      "Syntax errors found in manifest file",
    Description: "The manifest file has syntax issues like incorrect package names or unsupported characters.\nMake sure the manifest file follows the syntax stardards and can be installed locally as well.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-PYTHON-0005",
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

// NewPipUnsupportedPythonVersionError displays errors with the following description:
// At least one of the packages requires a Python version that doesn't match the one used in the project scan.
// Make sure to select a suitable Python version from the organization Python language settings.
// Alternatively, add a `.snyk` file for Python version selection override.
func NewPipUnsupportedPythonVersionError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-python-0006",
    Title:      "Python version not supported",
    Description: "At least one of the packages requires a Python version that doesn't match the one used in the project scan.\nMake sure to select a suitable Python version from the organization Python language settings.\nAlternatively, add a `.snyk` file for Python version selection override.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-PYTHON-0006",
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

// NewPythonVersionConfictError displays errors with the following description:
// Two or more packages have conflicting version requirements that cannot be resolved.
// Make sure no two packages and their requirements cause conflicts and that the manifest file can be installed locally.
func NewPythonVersionConfictError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-python-0007",
    Title:      "Packages versions caused conflicts",
    Description: "Two or more packages have conflicting version requirements that cannot be resolved.\nMake sure no two packages and their requirements cause conflicts and that the manifest file can be installed locally.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-PYTHON-0007",
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

// NewPipNoMatchingPythonDistributionError displays errors with the following description:
// At least one of the packages requires a Python version that doesn't match the one used in the project scan.
// Make sure to select a suitable Python version from the organization Python language settings.
// Alternatively, add a `.snyk` file for Python version selection override.
func NewPipNoMatchingPythonDistributionError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-python-0008",
    Title:      "No matching distribution found for one or more of the packages",
    Description: "At least one of the packages requires a Python version that doesn't match the one used in the project scan.\nMake sure to select a suitable Python version from the organization Python language settings.\nAlternatively, add a `.snyk` file for Python version selection override.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-PYTHON-0008",
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

// NewInstallationFailureError displays errors with the following description:
// Some packages failed during installation due to missing system dependencies, compilation errors, or other package-specific issues.
func NewInstallationFailureError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-python-0009",
    Title:      "Packages installation failed",
    Description: "Some packages failed during installation due to missing system dependencies, compilation errors, or other package-specific issues.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-PYTHON-0009",
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

// NewPipenvUnsupportedPythonVersionError displays errors with the following description:
// At least one of the packages requires a Python version that doesn't match the one used in the project scan.
// Make sure to use the correct python version in the requires section of the Pipfile.
func NewPipenvUnsupportedPythonVersionError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-python-0010",
    Title:      "Python version not supported",
    Description: "At least one of the packages requires a Python version that doesn't match the one used in the project scan.\nMake sure to use the correct python version in the requires section of the Pipfile.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-PYTHON-0010",
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

// NewPipenvNoMatchingPythonDistributionError displays errors with the following description:
// At least one of the packages requires a Python version that doesn't match the one used in the project scan.
// Make sure to use the correct python version in the requires section of the Pipfile.
func NewPipenvNoMatchingPythonDistributionError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-python-0011",
    Title:      "No matching distribution found for one or more of the packages",
    Description: "At least one of the packages requires a Python version that doesn't match the one used in the project scan.\nMake sure to use the correct python version in the requires section of the Pipfile.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-PYTHON-0011",
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

// NewPythonDependenciesSpaceLimitExceededError displays errors with the following description:
// The total size of the downloaded Python dependencies in the manifest file exceeds the 10GB limit.
// This often happens due to the large size of some of the Python dependencies and is usually the case for Python packages that require NVIDIA drivers like PyTorch. 
//
// Read more:
// - https://docs.snyk.io/supported-languages-package-managers-and-frameworks/python/git-repositories-and-python#pip-and-git-repositories
func NewPythonDependenciesSpaceLimitExceededError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-python-0012",
    Title:      "The 10 GB space limit for downloaded Python dependencies has been exceeded",
    Description: "The total size of the downloaded Python dependencies in the manifest file exceeds the 10GB limit.\nThis often happens due to the large size of some of the Python dependencies and is usually the case for Python packages that require NVIDIA drivers like PyTorch. ",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-PYTHON-0012",
    Classification: "UNSUPPORTED",
    Links: []string{
      "https://docs.snyk.io/supported-languages-package-managers-and-frameworks/python/git-repositories-and-python#pip-and-git-repositories",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewCyclicDependencyDetectedError displays errors with the following description:
// Cyclic dependency detected in lockfile.
func NewCyclicDependencyDetectedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-ruby-0001",
    Title:      "Cyclic dependency detected in lockfile",
    Description: "Cyclic dependency detected in lockfile.",
    StatusCode: 422,
    ErrorCode:  "SNYK-OS-RUBY-0001",
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

