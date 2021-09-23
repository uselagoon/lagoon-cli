# lagoon run activestandby

Run the active/standby switch for a project

## Synopsis

Run the active/standby switch for a project You should only run this once and then check the status of the task that gets created. If the task fails or fails to update, contact your Lagoon administrator for assistance.

```text
lagoon run activestandby [flags]
```

## Options

```text
  -h, --help   help for activestandby
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
  -i, --ssh-key string       Specify path to a specific SSH key to use for lagoon authentication
```

## SEE ALSO

* [lagoon run](lagoon_run.md)     - Run a task against an environment

