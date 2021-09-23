# lagoon export

Export lagoon output to YAML.

## Synopsis

Export lagoon output to YAML. You must specify to export a specific project by using the `-p` flag.

```text
lagoon export [flags]
```

## Options

```text
      --exclude strings   Exclude data from the export. Valid options (others are ignored): users, project-users, groups, notifications, project-private-keys (default [project-private-keys])
  -h, --help              Help for export
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

* [lagoon](lagoon.md)     - Command line integration for Lagoon.

