# Lagoon CLI

This is a CLI for interacting with a [Lagoon](https://github.com/amazeeio/lagoon) instance. By default, it is configured
to work against [Amazee.io](https://www.amazee.io/) instance.

## Usage

### `config`

Allows you to configure Lagoon CLI to specify endpoints, such as your own Lagoon instance.

### `add`

### `delete`

### `list`

### `update`

### `info`

### `deploy`

# Build
## Build using make
```
make deps
make build-linux
make build-darwin #macos
```

## Build using Docker
You can build lagoon-cli without installing `go` by running the `docker-build` make command. This will use the `Dockerfile.build` to build the cli inside of a docker container, then copy the binaries into the `builds/` directory once complete
```
make docker-build
```
Versions can also be defined, and the binaries will be version tagged
```
make VERSION=v0.0.1 docker-build
```

# Install
```
make install-linux
make install-darwin #macos
```

## Acknowledgements

[Matt Glaman](https://github.com/mglaman) - Initial conception and development - Thanks Matt!
