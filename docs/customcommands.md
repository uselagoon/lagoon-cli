# Introduction

Lagoon allows users to create simple custom commands that can execute raw graphql queries or mutations. The response of these commands will be the JSON response from the API, so tools like `jq` can be used to parse the response.

These commands are meant for simple tasks, and may not perform complex things very well. In some cases, the defaults of a flag may not work as you intend them to.

> **_NOTE:_** as always, be careful with creating your own commands, especially mutations, as you must be 100% aware of the implications.

## Location

Custom commands must be saved to `${HOME}/.lagoon-cli/commands/${COMMAND_NAME}.yml`

## Layout of a command file

An example of the command file structure is as follows
```yaml
name: project-by-name
description: Query a project by name
query: |
  query projectByName($name: String!) {
    projectByName(name: $name) {
      id
      name
      organization
      openshift{
        name
      }
      environments{
        name
        openshift{
          name
        }
      }
    }
  }
flags:
  - name: name
    description: Project name to check
    variable: name
    type: String
    required: true
```

* `name` is the name of the command that the user must enter, this should be unique
* `description` is some helpful information about this command
* `query` is the query or mutation that is run
* `flags` allows you to define your own flags
    * `name` is the name of the flag, eg `--name`
    * `description` is some helpful information about the flag
    * `variable` is the name of the variable that will be passed to the graphql query of the same name
    * `type` is the type, currently only `String`, `Int`, `Boolean` are supported
    * `required` is if this flag is required or not
    * `default` is the default value of the flag if defined
        * `String` defaults to ""
        * `Int` defaults to 0
        * `Boolean` defaults to false.

# Usage

Once a command file has been created, they will appear as `Available Commands` of the top level `custom` command, similarly to below

```
$ lagoon custom
Usage:
  lagoon custom [flags]
  lagoon custom [command]

Aliases:
  custom, cus, cust

Available Commands:
  project-by-name         Query a project by name

```

You can then call this command like so, and see the output of the command is the API JSON response
```
$ lagoon custom project-by-name --name lagoon-demo-org | jq
{
  "projectByName": {
    "environments": [
      {
        "name": "development",
        "openshift": {
          "name": "ui-kubernetes-2"
        }
      },
      {
        "name": "main",
        "openshift": {
          "name": "ui-kubernetes-2"
        }
      },
      {
        "name": "pr-15",
        "openshift": {
          "name": "ui-kubernetes-2"
        }
      },
      {
        "name": "staging",
        "openshift": {
          "name": "ui-kubernetes-2"
        }
      }
    ],
    "id": 180,
    "name": "lagoon-demo-org",
    "openshift": {
      "name": "ui-kubernetes-2"
    },
    "organization": 1
  }
}

```