# Introduction

Lagoon CLI uses a shared configuration package. You can [read more about it here](https://github.com/uselagoon/machinery/tree/main/utils/config).

It uses [XDG Base Directory Specification](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html) for storing the configuration in known locations.

The main configuration file will be stored in `$XDG_CONFIG_HOME` with the file location being in a `lagoon` directory, with the actual configuration file being named `config.yaml`.

The full path will then be `$XDG_CONFIG_HOME/lagoon/config.yml`. 

If `$XDG_CONFIG_HOME` is not set, then depending on the operating system, this location could be different. Review the specification to understand more about this.

# Usage

See the sub command `lagoon configuration` for information on managing users and contexts.

# Components

Configuration is broken down into two core components, `users` and `contexts`. You need a user, and you need a context, a user must be linked to a context.

Information about the two components are below, but further information can be found in the package that defines configurations if you want further information.

## Users

Users are a way to leverage separate SSH keys if using SSH based authentication. This user does not have to match the name of an account within the Lagoon, but a friendly name for the conifugration owner.

## Contexts

Contexts are a way to leverage multiple Lagoons. A context needs a user linked to it, you can change the user associated to a context, and you can define multiple contexts for the same cluster with different users.

## Features

The Lagoon CLI has some features that can be enabled and disabled. Features may be disabled by default, depending on the feature or features name. Features can be defined globally, or for a specific context if required.

See the help output of `lagoon configuration feature` for how to change features.

### ssh-token

This feature when set to `true` will only use SSH to get a token, instead of using the Lagoon provided Keycloak OAuth mechanism. This is useful if you're using the CLI and are unable to use the OAuth mechanism to authenticate with the API for any reason, for example in a CI job.

The default value of this is `false` (or unset).

### disable-update-check

This feature when set to `true` will inform the CLI to not perform the automatic update checks.

The default value of this is `false` (or unset).

### environment-from-directory

This feature when set to `true` will enable the abilty for the CLI to read some basic information from the directory the command is executed in to try and guess which project or environment you're in.

> Note: The recommendation is to never use this feature. We may deprecate this feature in the future.

The default value of this is `false` (or unset).

# Migrating from the legacy configuration

Previous versions of Lagoon CLI used a `.lagoon.yml` in a home directory. This is no longer the case. 

The CLI offers a way to convert a legacy configuration into the new format. It will not modify the legacy configuration though, only read it.

You can run it in dry run mode first to see what the resulting configuration will prompt you to provide input for.
```
lagoon configuration convert-config
```

If you're happy with the process, and the output looks good, you can run it again with the `--write-config` flag (you will be prompted for input again).
```
lagoon configuration convert-config --write-config
```