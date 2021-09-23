# lagoon get task-by-id

Get information about a task by its ID.

## Synopsis

Get information about a task by its ID.

```text
lagoon get task-by-id [flags]
```

## Options

```text
  -h, --help     Help for task-by-id
  -I, --id int   ID of the task
  -L, --logs     Show the task logs if available
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

* [lagoon get](lagoon_get.md)     - Get info on a resource.

