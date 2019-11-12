## Lagoon CLI

This is a CLI for interacting with a [Lagoon](https://github.com/amazeeio/lagoon) instance. By default, it is configured
to work against [Amazee.io](https://www.amazee.io/) instance.

If you run the CLI in a directory that has a valid `.lagoon.yml` and `docker-compose.yml` that references your project in lagoon, then you don't need to specify your project name on the command line as the CLI can read these files to determine the project. You can still define a project name though if you want to target a different project.

### Usage
Once installed, to use the Lagoon CLI, run the following command
```
lagoon <command>
```

### Commands
For the full list of commands see the docs for [Lagoon CLI](https://amazeeio.github.io/lagoon-cli/commands/lagoon/)

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

### Notes
Versions can also be defined, and the binaries will be version tagged
```
make VERSION=v0.0.1 ...
```

## Install
The best way to install is by downloading one of the pre-compiled binaries from the releases page - [Lagoon CLI releases](https://github.com/amazeeio/lagoon-cli/releases)
```
# MacOS
VERSION=0.1.0 sudo curl -L "https://github.com/amazeeio/lagoon-cli/releases/download/${VERSION}/lagoon-cli-${VERSION}-darwin-amd64" -o /usr/local/bin/lagoon

# Linux
VERSION=0.1.0 sudo curl -L "https://github.com/amazeeio/lagoon-cli/releases/download/${VERSION}/lagoon-cli-${VERSION}-linux-amd64" -o /usr/local/bin/lagoon
```

If you are building from source, you can install with one of the following commands
```
make ARTIFACT_DESTINATION=/usr/local/bin install-linux
make ARTIFACT_DESTINATION=/usr/local/bin install-darwin
```

### Acknowledgements

[Matt Glaman](https://github.com/mglaman) - Initial conception and development - Thanks Matt!
