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

// Package cli contains errors related to the namespace CLI
// of the Error Catalog.
package cli

import (
  "github.com/snyk/error-catalog-golang-public/snyk_errors"
  "github.com/google/uuid"
)
// NewGeneralCLIFailureError displays errors with the following description:
// The encountered error only provides basic information, please take a look at the given details. If they do not help to resolve the issue, consider debugging or consulting support.
//
// Read more:
// - https://docs.snyk.io/snyk-cli/commands
// - https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli
func NewGeneralCLIFailureError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cli-0000",
    Title:      "Unspecified Error",
    Description: "The encountered error only provides basic information, please take a look at the given details. If they do not help to resolve the issue, consider debugging or consulting support.",
    StatusCode: 0,
    ErrorCode:  "SNYK-CLI-0000",
    Classification: "UNEXPECTED",
    Links: []string{
      "https://docs.snyk.io/snyk-cli/commands",
      "https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewConfigEnvironmentFailedError displays errors with the following description:
// The specified environment cannot be used. As a result, the configuration remains unchanged. Provide the correct specifications for the environment and try again.
//
// Read more:
// - https://docs.snyk.io/snyk-cli/commands/config-environment
func NewConfigEnvironmentFailedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cli-0001",
    Title:      "Unable to set environment",
    Description: "The specified environment cannot be used. As a result, the configuration remains unchanged. Provide the correct specifications for the environment and try again.",
    StatusCode: 0,
    ErrorCode:  "SNYK-CLI-0001",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/snyk-cli/commands/config-environment",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewConfigEnvironmentConsistencyIssueError displays errors with the following description:
// You can configure the CLI in different ways, for example via Environment Variables or configuration file.
// If one parameter is configured multiple times, it is probably unintentional and might cause unexpected behavior.
// Review configured environment variables and ensure that everything is intentional. If so, you can skip this check by using --no-check.
//
// Read more:
// - https://docs.snyk.io/snyk-cli/commands/config-environment
func NewConfigEnvironmentConsistencyIssueError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cli-0002",
    Title:      "Possible inconsistent configuration",
    Description: "You can configure the CLI in different ways, for example via Environment Variables or configuration file.\nIf one parameter is configured multiple times, it is probably unintentional and might cause unexpected behavior.\nReview configured environment variables and ensure that everything is intentional. If so, you can skip this check by using --no-check.",
    StatusCode: 0,
    ErrorCode:  "SNYK-CLI-0002",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/snyk-cli/commands/config-environment",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewEmptyFlagOptionError displays errors with the following description:
// A specified flag is missing an option value. Provide a correct option value and try again.
//
// Read more:
// - https://docs.snyk.io/snyk-cli/cli-commands-and-options-summary
func NewEmptyFlagOptionError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cli-0003",
    Title:      "Empty flag option",
    Description: "A specified flag is missing an option value. Provide a correct option value and try again.",
    StatusCode: 0,
    ErrorCode:  "SNYK-CLI-0003",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/snyk-cli/cli-commands-and-options-summary",
    },
    Level:  "fatal",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewInvalidFlagOptionError displays errors with the following description:
// A specified flag option or combination is invalid. Provide a valid flag option or combination and try again.
//
// Read more:
// - https://docs.snyk.io/snyk-cli/cli-commands-and-options-summary
func NewInvalidFlagOptionError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cli-0004",
    Title:      "Invalid flag option",
    Description: "A specified flag option or combination is invalid. Provide a valid flag option or combination and try again.",
    StatusCode: 0,
    ErrorCode:  "SNYK-CLI-0004",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/snyk-cli/cli-commands-and-options-summary",
    },
    Level:  "fatal",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewGetVulnsFromResourceFailedError displays errors with the following description:
// If you are testing an npm package, check the version and package name and try running `snyk test` again. If you are testing a repository, try testing it at https://snyk.io/test/. For further assistance, run `snyk help` or see the Snyk docs.
func NewGetVulnsFromResourceFailedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cli-0005",
    Title:      "Unable to get vulnerabilities from resource",
    Description: "If you are testing an npm package, check the version and package name and try running `snyk test` again. If you are testing a repository, try testing it at https://snyk.io/test/. For further assistance, run `snyk help` or see the Snyk docs.",
    StatusCode: 0,
    ErrorCode:  "SNYK-CLI-0005",
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

// NewAuthConfigError displays errors with the following description:
// When running your command, Snyk requires an authenticated account. You must include your API token as an environment value, or use `snyk auth` to authenticate.
//
// Read more:
// - https://docs.snyk.io/snyk-cli/authenticate-to-use-the-cli
// - https://docs.snyk.io/snyk-cli/configure-the-snyk-cli/environment-variables-for-snyk-cli
func NewAuthConfigError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cli-0006",
    Title:      "Missing AUTH token",
    Description: "When running your command, Snyk requires an authenticated account. You must include your API token as an environment value, or use `snyk auth` to authenticate.",
    StatusCode: 0,
    ErrorCode:  "SNYK-CLI-0006",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/snyk-cli/authenticate-to-use-the-cli",
      "https://docs.snyk.io/snyk-cli/configure-the-snyk-cli/environment-variables-for-snyk-cli",
    },
    Level:  "fatal",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewCommandArgsError displays errors with the following description:
// The specified CLI command includes missing or misconfigured arguments. Provide the correct arguments and try again.
//
// Read more:
// - https://docs.snyk.io/snyk-cli/cli-commands-and-options-summary
func NewCommandArgsError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cli-0007",
    Title:      "Incomplete command arguments",
    Description: "The specified CLI command includes missing or misconfigured arguments. Provide the correct arguments and try again.",
    StatusCode: 0,
    ErrorCode:  "SNYK-CLI-0007",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/snyk-cli/cli-commands-and-options-summary",
    },
    Level:  "fatal",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewNoSupportedFilesFoundError displays errors with the following description:
// Snyk could not detect any supported target files. Ensure the files you are importing are supported, that you are in the right directory, and try again.
//
// Read more:
// - https://docs.snyk.io/supported-languages-package-managers-and-frameworks
func NewNoSupportedFilesFoundError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cli-0008",
    Title:      "No supported files found",
    Description: "Snyk could not detect any supported target files. Ensure the files you are importing are supported, that you are in the right directory, and try again.",
    StatusCode: 0,
    ErrorCode:  "SNYK-CLI-0008",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/supported-languages-package-managers-and-frameworks",
    },
    Level:  "fatal",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewTooManyVulnerablePathsError displays errors with the following description:
// There are too many vulnerable paths to process the project. If your command supports it, consider the following: pruning repeated sub-dependencies (`snyk test -p`); excluding directories (`snyk test --all-projects --exclude=dir1,file2`); setting a detection depth (`snyk test --all-projects --detection-depth=3`). If the error still occurs, consider debugging or contact Snyk Support.
//
// Read more:
// - https://docs.snyk.io/snyk-cli/cli-commands-and-options-summary#options-for-multiple-commands
// - https://docs.snyk.io/snyk-cli/commands/test#prune-repeated-subdependencies-p
// - https://docs.snyk.io/snyk-cli/commands/test#detection-depth-less-than-depth-greater-than
// - https://docs.snyk.io/snyk-cli/commands/test#exclude-less-than-name-greater-than-less-than-name-greater-than-...greater-than
func NewTooManyVulnerablePathsError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cli-0009",
    Title:      "Too many vulnerable paths to Project",
    Description: "There are too many vulnerable paths to process the project. If your command supports it, consider the following: pruning repeated sub-dependencies (`snyk test -p`); excluding directories (`snyk test --all-projects --exclude=dir1,file2`); setting a detection depth (`snyk test --all-projects --detection-depth=3`). If the error still occurs, consider debugging or contact Snyk Support.",
    StatusCode: 413,
    ErrorCode:  "SNYK-CLI-0009",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/snyk-cli/cli-commands-and-options-summary#options-for-multiple-commands",
      "https://docs.snyk.io/snyk-cli/commands/test#prune-repeated-subdependencies-p",
      "https://docs.snyk.io/snyk-cli/commands/test#detection-depth-less-than-depth-greater-than",
      "https://docs.snyk.io/snyk-cli/commands/test#exclude-less-than-name-greater-than-less-than-name-greater-than-...greater-than",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewValidationFailureError displays errors with the following description:
// CLI was unable to validate the required parameter. Provide the correct parameter and try again. If the error still occurs, consider debugging or contact Snyk Support.
func NewValidationFailureError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cli-0010",
    Title:      "CLI validation failure",
    Description: "CLI was unable to validate the required parameter. Provide the correct parameter and try again. If the error still occurs, consider debugging or contact Snyk Support.",
    StatusCode: 400,
    ErrorCode:  "SNYK-CLI-0010",
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

// NewGeneralSCAFailureError displays errors with the following description:
// CLI was unable to execute your SCA command, please take a look at the given details. If they do not help to resolve the issue, consider debugging or consulting support.
//
// Read more:
// - https://docs.snyk.io/snyk-cli/commands/test
func NewGeneralSCAFailureError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cli-0011",
    Title:      "SCA failure",
    Description: "CLI was unable to execute your SCA command, please take a look at the given details. If they do not help to resolve the issue, consider debugging or consulting support.",
    StatusCode: 0,
    ErrorCode:  "SNYK-CLI-0011",
    Classification: "UNEXPECTED",
    Links: []string{
      "https://docs.snyk.io/snyk-cli/commands/test",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewGeneralIACFailureError displays errors with the following description:
// CLI was unable to execute your IAC command, please take a look at the given details. If they do not help to resolve the issue, consider debugging or consulting support.
//
// Read more:
// - https://docs.snyk.io/snyk-cli/commands/iac
func NewGeneralIACFailureError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cli-0012",
    Title:      "IAC failure",
    Description: "CLI was unable to execute your IAC command, please take a look at the given details. If they do not help to resolve the issue, consider debugging or consulting support.",
    StatusCode: 0,
    ErrorCode:  "SNYK-CLI-0012",
    Classification: "UNEXPECTED",
    Links: []string{
      "https://docs.snyk.io/snyk-cli/commands/iac",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewGeneralSASTFailureError displays errors with the following description:
// CLI was unable to execute your SAST command, please take a look at the given details. If they do not help to resolve the issue, consider debugging or consulting support.
//
// Read more:
// - https://docs.snyk.io/snyk-cli/commands/code
func NewGeneralSASTFailureError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cli-0013",
    Title:      "SAST failure",
    Description: "CLI was unable to execute your SAST command, please take a look at the given details. If they do not help to resolve the issue, consider debugging or consulting support.",
    StatusCode: 0,
    ErrorCode:  "SNYK-CLI-0013",
    Classification: "UNEXPECTED",
    Links: []string{
      "https://docs.snyk.io/snyk-cli/commands/code",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewFeatureUnderDevelopmentError displays errors with the following description:
// This feature is under development and is not yet available for public use.
func NewFeatureUnderDevelopmentError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cli-0014",
    Title:      "Feature under development",
    Description: "This feature is under development and is not yet available for public use.",
    StatusCode: 0,
    ErrorCode:  "SNYK-CLI-0014",
    Classification: "UNSUPPORTED",
    Links: []string{},
    Level:  "fatal",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewCommandIsExperimentalError displays errors with the following description:
// This CLI command is experimental, which means it is provided "as-is" without warranty of any kind.
// You must acknowledge this by specifying the --experimental flag to run the command.
func NewCommandIsExperimentalError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cli-0015",
    Title:      "Command is experimental",
    Description: "This CLI command is experimental, which means it is provided \"as-is\" without warranty of any kind.\nYou must acknowledge this by specifying the --experimental flag to run the command.",
    StatusCode: 0,
    ErrorCode:  "SNYK-CLI-0015",
    Classification: "ACTIONABLE",
    Links: []string{},
    Level:  "fatal",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewFeatureNotEnabledError displays errors with the following description:
// This feature is disabled for your current organization. You can enable it in the settings or switch to an organization where it's already enabled.
func NewFeatureNotEnabledError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cli-0016",
    Title:      "Feature not enabled",
    Description: "This feature is disabled for your current organization. You can enable it in the settings or switch to an organization where it's already enabled.",
    StatusCode: 403,
    ErrorCode:  "SNYK-CLI-0016",
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

// NewDNSResolutionError displays errors with the following description:
// Unable to resolve the hostname to an IP address. Troubleshooting steps.
// 1) Test DNS resolution: nslookup api.<instance>.snyk.io.
// 2) Try different DNS servers: Change DNS to 8.8.8.8 or 1.1.1.1.
// 3) Check corporate proxy/firewall DNS blocking.
// 4) Verify hostname spelling in your Snyk configuration.
//
// Read more:
// - https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli
// - https://status.snyk.io/
func NewDNSResolutionError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cli-0017",
    Title:      "DNS resolution failed",
    Description: "Unable to resolve the hostname to an IP address. Troubleshooting steps.\n1) Test DNS resolution: nslookup api.<instance>.snyk.io.\n2) Try different DNS servers: Change DNS to 8.8.8.8 or 1.1.1.1.\n3) Check corporate proxy/firewall DNS blocking.\n4) Verify hostname spelling in your Snyk configuration.",
    StatusCode: 0,
    ErrorCode:  "SNYK-CLI-0017",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli",
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

// NewNetworkTimeoutError displays errors with the following description:
// The network request timed out. Troubleshooting steps.
// 1) Test connectivity: ping api.<instance>.snyk.io.
// 2) Check corporate proxy timeout settings.
// 3) Try different network: Mobile hotspot or different WiFi.
// 4) Check if firewall is blocking or throttling connections.
//
// Read more:
// - https://status.snyk.io/
// - https://privatecloudstatus.snyk.io
// - https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli
func NewNetworkTimeoutError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cli-0018",
    Title:      "Network request timeout",
    Description: "The network request timed out. Troubleshooting steps.\n1) Test connectivity: ping api.<instance>.snyk.io.\n2) Check corporate proxy timeout settings.\n3) Try different network: Mobile hotspot or different WiFi.\n4) Check if firewall is blocking or throttling connections.",
    StatusCode: 0,
    ErrorCode:  "SNYK-CLI-0018",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://status.snyk.io/",
      "https://privatecloudstatus.snyk.io",
      "https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewNetworkUnreachableError displays errors with the following description:
// Unable to reach the target network or host. Troubleshooting steps.
// 1) Test Snyk connectivity: ping api.<instance>.snyk.io.
// 2) Check corporate firewall blocks Snyk domains.
// 3) Check if VPN routing is blocking Snyk domains.
// 4) Try mobile hotspot to isolate network issues.
//
// Read more:
// - https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli
// - https://status.snyk.io/
func NewNetworkUnreachableError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cli-0019",
    Title:      "Network unreachable",
    Description: "Unable to reach the target network or host. Troubleshooting steps.\n1) Test Snyk connectivity: ping api.<instance>.snyk.io.\n2) Check corporate firewall blocks Snyk domains.\n3) Check if VPN routing is blocking Snyk domains.\n4) Try mobile hotspot to isolate network issues.",
    StatusCode: 0,
    ErrorCode:  "SNYK-CLI-0019",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli",
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

// NewTLSCertificateError displays errors with the following description:
// There was an issue with the TLS/SSL certificate during the secure connection. Troubleshooting steps.
// 1) Check system time: Ensure your system clock is correct (certificates are time-sensitive).
// 2) Update system certificates: Windows Update or macOS Software Update.
// 3) Corporate firewall: Check if corporate firewall intercepts SSL traffic.
// 4) Custom certificates: Set NODE_EXTRA_CA_CERTS environment variable to path of your CA certificate file.
//
// Read more:
// - https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli
// - https://docs.snyk.io/snyk-cli/configure-the-snyk-cli/environment-variables-for-snyk-cli
func NewTLSCertificateError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cli-0020",
    Title:      "TLS certificate error",
    Description: "There was an issue with the TLS/SSL certificate during the secure connection. Troubleshooting steps.\n1) Check system time: Ensure your system clock is correct (certificates are time-sensitive).\n2) Update system certificates: Windows Update or macOS Software Update.\n3) Corporate firewall: Check if corporate firewall intercepts SSL traffic.\n4) Custom certificates: Set NODE_EXTRA_CA_CERTS environment variable to path of your CA certificate file.",
    StatusCode: 0,
    ErrorCode:  "SNYK-CLI-0020",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli",
      "https://docs.snyk.io/snyk-cli/configure-the-snyk-cli/environment-variables-for-snyk-cli",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewConnectionRefusedError displays errors with the following description:
// The connection to the server was refused. Troubleshooting steps.
// 1) Check Snyk status: Visit status.snyk.io for service outages.
// 2) Test connectivity: ping api.<instance>.snyk.io.
// 3) Check corporate proxy blocks HTTPS connections to Snyk.
// 4) Try mobile hotspot or different network.
//
// Read more:
// - https://status.snyk.io/
// - https://privatecloudstatus.snyk.io
// - https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli
func NewConnectionRefusedError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cli-0021",
    Title:      "Connection refused",
    Description: "The connection to the server was refused. Troubleshooting steps.\n1) Check Snyk status: Visit status.snyk.io for service outages.\n2) Test connectivity: ping api.<instance>.snyk.io.\n3) Check corporate proxy blocks HTTPS connections to Snyk.\n4) Try mobile hotspot or different network.",
    StatusCode: 0,
    ErrorCode:  "SNYK-CLI-0021",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://status.snyk.io/",
      "https://privatecloudstatus.snyk.io",
      "https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}

// NewGenericNetworkError displays errors with the following description:
// An unexpected network error occurred during communication. Troubleshooting steps.
// 1) Test connectivity: ping api.<instance>.snyk.io.
// 2) Check proxy settings: HTTP_PROXY and HTTPS_PROXY environment variables.
// 3) Run with verbose logging: snyk command --debug.
// 4) Try mobile hotspot to isolate network issues.
//
// Read more:
// - https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli
// - https://status.snyk.io/
func NewGenericNetworkError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-cli-0022",
    Title:      "Network communication error",
    Description: "An unexpected network error occurred during communication. Troubleshooting steps.\n1) Test connectivity: ping api.<instance>.snyk.io.\n2) Check proxy settings: HTTP_PROXY and HTTPS_PROXY environment variables.\n3) Run with verbose logging: snyk command --debug.\n4) Try mobile hotspot to isolate network issues.",
    StatusCode: 0,
    ErrorCode:  "SNYK-CLI-0022",
    Classification: "ACTIONABLE",
    Links: []string{
      "https://docs.snyk.io/snyk-cli/debugging-the-snyk-cli",
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

// NewConnectionTimeoutError displays errors with the following description:
// A request to the Snyk API has unexpectedly timeout. Check Snyk status, then try again.
//
// Read more:
// - https://status.snyk.io/
// - https://privatecloudstatus.snyk.io
func NewConnectionTimeoutError(detail string, options ...snyk_errors.Option) snyk_errors.Error {
  err := snyk_errors.Error{
    ID:         uuid.NewString(),
    Type:       "https://docs.snyk.io/scan-with-snyk/error-catalog#snyk-os-7001",
    Title:      "Request to Snyk API timeout",
    Description: "A request to the Snyk API has unexpectedly timeout. Check Snyk status, then try again.",
    StatusCode: 504,
    ErrorCode:  "SNYK-OS-7001",
    Classification: "UNEXPECTED",
    Links: []string{
      "https://status.snyk.io/",
      "https://privatecloudstatus.snyk.io",
    },
    Level:  "error",
    Detail: detail,
  }

  for _, option := range options {
    option(&err)
  }

  return err
}
