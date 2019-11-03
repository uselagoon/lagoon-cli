# Introduction
This is a CLI for interacting with a [Lagoon](https://github.com/amazeeio/lagoon) instance. By default, it is configured
to work against [Amazee.io](https://www.amazee.io/) instance.

If you run the CLI in a directory that has a valid `.lagoon.yml` and `docker-compose.yml` that references your project in lagoon, then you don't need to specify your project name on the command line as the CLI can read these files to determine the project. You can still define a project name though if you want to target a different project.

# Requirements
To use this CLI, you need an account in the Lagoon that you wish to communicate with, and your SSH key needs to be associated to your account.

## Usage
@TODO