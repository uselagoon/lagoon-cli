# Lagoon CLI

This is a CLI for interacting with a [Lagoon](https://github.com/amazeeio/lagoon) instance. By default, it is configured
to work against [Amazee.io](https://www.amazee.io/) instance.

If you run the CLI in a directory that has a valid `.lagoon.yml` and `docker-compose.yml` that references your project in lagoon, then you don't need to specify your project name on the command line as the CLI can read these files to determine the project. You can still define a project name though if you want to target a different project.

## Usage

### `config`

Allows you to configure Lagoon CLI to specify endpoints, such as your own Lagoon instance.

### `add`

* Projects
* Variables
    * To Projects
    * To Environments
* Notifications
    * Slack
    * RocketChat

### `delete`

* Projects
* Environments
* Variables
    * From Projects
    * From Environments
* Notifications
    * Slack
    * RocketChat


### `list`

* Projects
* Deployments
* Variables
* Notifications
    * Slack
    * RocketChat

### `update`

* Projects

### `info`

* Projects
* Deployments

### `deploy`

* Environments

# Build
## Run tests
```
make test
make test-docker
```

## Build locally
```
make build-linux
make build-darwin #macos
```

## Build using Docker
You can build lagoon-cli without installing `go` by running the `docker-build` make command. This will use the `Dockerfile.build` to build the cli inside of a docker container, then copy the binaries into the `builds/` directory once complete
```
make build-docker-darwin
make build-docker-linux
```

## Run all
```
make all #locally
make all-docker-linux
make all-docker-darwin
```

## Notes
Versions can also be defined, and the binaries will be version tagged
```
make VERSION=v0.0.1 ...
```

# Install
```
make install-linux
make install-darwin
```

## Acknowledgements

[Matt Glaman](https://github.com/mglaman) - Initial conception and development - Thanks Matt!
