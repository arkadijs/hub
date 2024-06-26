# Hub CLI

[Hub CLI](https://superhub.io) is stack composition and lifecycle tool.

## Example of usage

```shell
hub elaborate hub.yaml params.yaml -o hub.yaml.elaborate
hub deploy hub.yaml.elaborate -e NAME=stage
hub version
hub help
```

## Installation

### Pre-build binary

To download the latest release, run:

```shell
curl -LJ "https://github.com/agilestacks/hub/releases/latest/download/hub_$(uname -s)_$(uname -m).tar.gz" | tar xz -C /tmp
sudo mv /tmp/hub /usr/local/bin
```

There are [macOS amd64](https://github.com/agilestacks/hub/releases/latest/download/hub_Darwin_arm64.tar.gz), [macOS arm64](https://github.com/agilestacks/hub/releases/latest/download/hub_Darwin_x86_64.tar.gz), [Linux amd64](https://github.com/agilestacks/hub/releases/latest/download/hub_Linux_arm64.tar.gz), [Linux arm64](https://github.com/agilestacks/hub/releases/latest/download/hub_Linux_x86_64.tar.gz) and [Windows x64](https://github.com/agilestacks/hub/releases/latest/download/hub_Windows_x86_64.zip) binaries.

### [Homebrew](https://brew.sh/)

```shell
brew tap agilestacks/tap
brew install agilestacks/tap/hub
```

Sorry for the name conflict with [GitHub hub](https://hub.github.com).

### Extensions

Optionally, install extensions:

```shell
hub extensions install
```

Hub CLI Extensions requires: [jq] and [yq v4].
Optionally install [Node.js] and NPM for `hub pull` extension, [AWS CLI], [Azure CLI], [GCP SDK] [kubectl], [eksctl] for `hub ext eks` extension.

### macOS users

Depending on your's machine Security & Privacy settings and macOS version (10.15+), you may get an error _cannot be opened because the developer cannot be verified_. Please [read on](https://github.com/hashicorp/terraform/issues/23033#issuecomment-542302933) for a simple workaround:

```shell
xattr -d com.apple.quarantine hub.darwin_amd64
```

Alternatively, to set a global preference to _Allow apps downloaded from: Anywhere_, execute:

```shell
sudo spctl --master-disable
```

## Development

### Setup

Before make any changes you should configure git hooks for this repository

```bash
git config core.hooksPath .githooks
```

### Build

Use `make` to build Hub CLI:

```shell
make
```

Or directly with `go`:

```shell
go build -o bin/$(go env GOOS)/hub github.com/agilestacks/hub/cmd/hub
```

### Clean up

```shell
make clean
```

## Usage metrics

When you use a pre-built binary from the releases page, it will send usage metrics to SuperHub and Datadog.

We value your privacy and only send anonymized usage metrics for the following commands:

- elaborate
- deploy
- undeploy
- backup create
- api *

A usage metric sample contains:

- Hub CLI command invoked without arguments, ie. 'deploy' or 'backup create', or 'api instance get'
- synthetic machine id - an UUID generated in first interactive session (stdout is a TTY)
- usage counter - 1 per invocation

Edit `$HOME/.hub-cache.yaml` to change settings:

```yaml
    metrics:
      disabled: false
      host: 68af657e-6a51-4d4b-890c-4b548852724d
```

Set `disabled: true` to skip usage metrics reporting.
Set `host: ""` to send the counter but not the UUID.

You could always review an up-to-date help via `hub util metrics -h`.

## What's next?

Deploy [App Stack](https://github.com/agilestacks/stack-app-eks) or [Machine Learning Stack](https://github.com/agilestacks/stack-ml-eks) on AWS EKS.

More information in the [wiki](https://github.com/agilestacks/hub/wiki).

[AWS CLI]: https://aws.amazon.com/cli/
[Azure CLI]: https://docs.microsoft.com/en-us/cli/azure/
[GCP SDK CLI]: https://cloud.google.com/sdk/docs/install
[kubectl]: https://kubernetes.io/docs/reference/kubectl/overview/
[eksctl]: https://eksctl.io
[jq]: https://stedolan.github.io/jq/
[yq v4]: https://github.com/mikefarah/yq
[Node.js]: https://nodejs.org
