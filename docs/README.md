![lagoon-cli logo](lagoon-cli-logo.png){: width=100}

# Introduction

This is a CLI for interacting with a [Lagoon](https://github.com/uselagoon/lagoon) instance. By default, it is configured to work against [Amazee.io](https://www.amazee.io/) instance.

If you run the CLI in a directory that has a valid `.lagoon.yml` and `docker-compose.yml` that references your project in lagoon, then you don't need to specify your project name on the command line as the CLI can read these files to determine the project. You can still define a project name though if you want to target a different project.

# Requirements
To use this CLI, you need an account in the Lagoon that you wish to communicate with, and your SSH key needs to be associated to your account.

# Installation
The preferred method to install is via [Homebrew](https://brew.sh/).
```
brew tap amazeeio/lagoon-cli
brew install lagoon
```

Alternatively, you may install by downloading one of the pre-compiled binaries from the [releases page](https://github.com/uselagoon/lagoon-cli/releases)

# Running as a Docker Image
In order to use the Lagoon CLI as a docker image (if that's the way you roll) you will need to add your own `.lagoon.yml` and ssh keys as volume mounts. This will use your existing
config files with their defaults etc, and the full range of [Commands](commands/lagoon.md) are available.  Note that it needs read-write access to the .lagoon.yml to store the login token.
```
docker run \
-v ~/.lagoon.yml:/root/.lagoon.yml:rw \
-v ~/.ssh/cli_id_rsa:/root/.ssh/id_rsa:ro \
uselagoon/lagoon-cli:latest \
config list
```

# Usage
See [Commands](commands/lagoon.md)
