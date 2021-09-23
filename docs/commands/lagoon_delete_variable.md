# lagoon delete variable

Delete a variable from an environment.

## Synopsis

This allows you to delete an environment variable from a project.

```text
lagoon delete variable [flags]
```

## Options

```text
  -h, --help          Help for variable
  -N, --name string   Name of the variable to delete
```

## Options inherited from parent commands

```text
      --config-file string   Path to the config file to use (must be *.yml or *.yaml)
      --debug                Enable debugging output (if supported)
  -e, --environment string   Specify an environment to use
      --force                Force yes on prompts (if supported)
  -l, --lagoon string        The Lagoon instance to interact with
      --no-header            No header on table (if supported)
      --output-csv           Output as CSV (if supported)
      --output-json          Output as JSON (if supported)
      --pretty               Make JSON pretty (if supported)
  -p, --project string       Specify a project to use
      --skip-update-check    Skip checking for updates
  -i, --ssh-key string       Specify path to a specific SSH key to use for Lagoon authentication
```

## SEE ALSO

* [lagoon delete](lagoon_delete.md)     - Delete a project, or delete notifications and variables from projects or environments.

