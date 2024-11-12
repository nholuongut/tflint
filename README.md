# TFLint

![](https://i.imgur.com/waxVImv.png)
### [View all Roadmaps](https://github.com/nholuongut/all-roadmaps) &nbsp;&middot;&nbsp; [Best Practices](https://github.com/nholuongut/all-roadmaps/blob/main/public/best-practices/) &nbsp;&middot;&nbsp; [Questions](https://www.linkedin.com/in/nholuong/)
<br/>

A Pluggable [Terraform](https://www.terraform.io/) Linter

## Features

TFLint is a framework and each feature is provided by plugins, the key features are as follows:

- Find possible errors (like invalid instance types) for Major Cloud providers (AWS/Azure/GCP).
- Warn about deprecated syntax, unused declarations.
- Enforce best practices, naming conventions.

## Installation

Bash script (Linux):

```console
curl -s https://raw.githubusercontent.com/nholuongut/tflint/master/install_linux.sh | bash
```

Homebrew (macOS):

```console
brew install tflint
```

Chocolatey (Windows):

```cmd
choco install tflint
```

NOTE: The Chocolatey package is NOT directly maintained by the TFLint maintainers. The latest version is always available by manual installation.

### Verification

#### GitHub CLI (Recommended)

[Artifact Attestations](https://docs.github.com/en/actions/security-guides/using-artifact-attestations-to-establish-provenance-for-builds) are available that can be verified using the GitHub CLI.

```console
gh attestation verify checksums.txt -R nholuongut/tflint
sha256sum --ignore-missing -c checksums.txt
```

#### Cosign

[Cosign](https://github.com/sigstore/cosign) `verify-blob` command ensures that the release was built with GitHub Actions in this repository.

```console
cosign verify-blob --certificate=checksums.txt.pem --signature=checksums.txt.keyless.sig --certificate-identity-regexp="^https://github.com/nholuongut/tflint" --certificate-oidc-issuer=https://token.actions.githubusercontent.com checksums.txt
sha256sum --ignore-missing -c checksums.txt
```

### Docker

Instead of installing directly, you can use the Docker image:

```console
docker run --rm -v $(pwd):/data -t ghcr.io/nholuongut/tflint
```

### GitHub Actions

If you want to run on GitHub Actions, [setup-tflint](https://github.com/nholuongut/setup-tflint) action is available.

## Getting Started

First, enable rules for [Terraform Language](https://www.terraform.io/language) (e.g. warn about deprecated syntax, unused declarations). [TFLint Ruleset for Terraform Language](https://github.com/nholuongut/tflint-ruleset-terraform) is bundled with TFLint, so you can use it without installing it separately.

The bundled plugin enables the "recommended" preset by default, but you can disable the plugin or use a different preset. Declare the plugin block in `.tflint.hcl` like this:

```hcl
plugin "terraform" {
  enabled = true
  preset  = "recommended"
}
```

See the [tflint-ruleset-terraform documentation](https://github.com/nholuongut/tflint-ruleset-terraform/blob/main/docs/configuration.md) for more information.

Next, If you are using an AWS/Azure/GCP provider, it is a good idea to install the plugin and try it according to each usage:

- [Amazon Web Services](https://github.com/nholuongut/tflint-ruleset-aws)
- [Microsoft Azure](https://github.com/nholuongut/tflint-ruleset-azurerm)
- [Google Cloud Platform](https://github.com/nholuongut/tflint-ruleset-google)

If you want to extend TFLint with other plugins, you can declare the plugins in the config file and easily install them with `tflint --init`.

```hcl
plugin "foo" {
  enabled = true
  version = "0.1.0"
  source  = "github.com/org/tflint-ruleset-foo"

  signing_key = <<-KEY
  -----BEGIN PGP PUBLIC KEY BLOCK-----

  xxxxxx
  ...
  KEY
}
```

See also [Configuring Plugins](docs/user-guide/plugins.md).

If you want to add custom rules that are not in existing plugins, you can build your own plugin or write your own policy in Rego. See [Writing Plugins](docs/developer-guide/plugins.md) or [OPA Ruleset](https://github.com/nholuongut/tflint-ruleset-opa).

## Usage

TFLint inspects files under the current directory by default. You can change the behavior with the following options/arguments:

```
$ tflint --help
Usage:
  tflint --chdir=DIR/--recursive [OPTIONS]

Application Options:
  -v, --version                                                 Print TFLint version
      --init                                                    Install plugins
      --langserver                                              Start language server
  -f, --format=[default|json|checkstyle|junit|compact|sarif]    Output format
  -c, --config=FILE                                             Config file name (default: .tflint.hcl)
      --ignore-module=SOURCE                                    Ignore module sources
      --enable-rule=RULE_NAME                                   Enable rules from the command line
      --disable-rule=RULE_NAME                                  Disable rules from the command line
      --only=RULE_NAME                                          Enable only this rule, disabling all other defaults. Can be specified multiple times
      --enable-plugin=PLUGIN_NAME                               Enable plugins from the command line
      --var-file=FILE                                           Terraform variable file name
      --var='foo=bar'                                           Set a Terraform variable
      --call-module-type=[all|local|none]                       Types of module to call (default: local)
      --chdir=DIR                                               Switch to a different working directory before executing the command
      --recursive                                               Run command in each directory recursively
      --filter=FILE                                             Filter issues by file names or globs
      --force                                                   Return zero exit status even if issues found
      --minimum-failure-severity=[error|warning|notice]         Sets minimum severity level for exiting with a non-zero error code
      --color                                                   Enable colorized output
      --no-color                                                Disable colorized output
      --fix                                                     Fix issues automatically
      --no-parallel-runners                                     Disable per-runner parallelism
      --max-workers=N                                           Set maximum number of workers in recursive inspection (default: number of CPUs)

Help Options:
  -h, --help                                                    Show this help message
```

See [User Guide](docs/user-guide) for details.

## Debugging

If you don't get the expected behavior, you can see the detailed logs when running with `TFLINT_LOG` environment variable.

```console
$ TFLINT_LOG=debug tflint
```

## Developing

See [Developer Guide](docs/developer-guide).

## Security

If you find a security vulnerability, please refer our [security policy](SECURITY.md).


## 🚀 I'm are always open to your feedback.  Please contact as bellow information
![](https://i.imgur.com/waxVImv.png)
# **[Contact Me]**
* [Name: Nho Luong]
* [Skype](luongutnho_skype)
* [Github](https://github.com/nholuongut/)
* [Linkedin](https://www.linkedin.com/in/nholuong/)
* [Email Address](luongutnho@hotmail.com)
* [PayPal.Me](https://www.paypal.com/paypalme/nholuongut)

![](Donate.png)
[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/nholuong)

![](https://i.imgur.com/waxVImv.png)
# License
* Nho Luong (c). All Rights Reserved.🌟