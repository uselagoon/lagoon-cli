## Lagoon CLI

<img src="./docs/lagoon-cli-logo.png" width=100>

This is a CLI for interacting with a [Lagoon](https://github.com/uselagoon/lagoon) instance. By default, it is configured
to work against [Amazee.io](https://www.amazee.io/) instances.

If you run the CLI in a directory that has a valid `.lagoon.yml` and `docker-compose.yml` that references your project in lagoon, then you don't need to specify your project name on the command line as the CLI can read these files to determine the project. You can still define a project name though if you want to target a different project.

## Install
The preferred method is installation via [Homebrew](https://brew.sh/).
```
brew tap uselagoon/lagoon-cli
brew install lagoon
```

Alternatively, you may install by downloading one of the pre-compiled binaries from the [releases page](https://github.com/uselagoon/lagoon-cli/releases)

If you are building from source, see the Build section below

### Usage
Once installed, to use the Lagoon CLI, run the following command
```
lagoon <command>
```

### Commands
For the full list of commands see the docs for [Lagoon CLI](https://uselagoon.github.io/lagoon-cli/commands/lagoon/)

## Building

### Requirements

Install `Go` - https://go.dev/doc/install

You also need `mockgen`, it can be installed using the following command once `Go` is installed.

```
go install go.uber.org/mock/mockgen@v0.4.0
```

Note: You should make sure you have your `GOPATH` configured and in your path, see https://pkg.go.dev/cmd/go#hdr-GOPATH_environment_variable

### Run tests
```
make test
```

### Build locally

You can compile the binary and load it into your `GOPATH` bin directory using the following.
```
make build
```

Alternatively, these will compile a binary inside a `builds` directory in this repository, you can place them wherever you wish.
```
make build-linux
#macos
make build-darwin
make build-darwin-arm64
```

### Build using Docker
You can build lagoon-cli without installing `go` by running the `docker-build` make command. This will use the `Dockerfile.build` to build the cli inside of a docker container, then copy the binaries into the `builds/` directory once complete
```
make build-docker-darwin
make build-docker-linux
```

### Run all
```
make all #locally
make all-docker-linux
make all-docker-darwin
```

### Install
```
make ARTIFACT_DESTINATION=/usr/local/bin install-linux
make ARTIFACT_DESTINATION=/usr/local/bin install-darwin
```

### Notes
Versions can also be defined, and the binaries will be version tagged
```
make VERSION=v0.0.1 ...
```

### Acknowledgements

[Matt Glaman](https://github.com/mglaman) - Initial conception and development - Thanks Matt!
