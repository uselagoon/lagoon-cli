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

## Build

Note: You should make sure you have your `GOPATH` configured and in your path if you are going to build the lagoon CLI. If you haven't got `go` installed and are using the docker method, you can export `GOPATH` to be somewhere else in your `PATH` for binaries.

### Run tests
```
make test
make test-docker
```

### Build locally
```
make build-linux
make build-darwin #macos
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

### Releasing
New releases can be created by running one of the following, this will create the version bump and update the `gh-pages` branch
```
make release-patch
make release-minor
make release-major
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
