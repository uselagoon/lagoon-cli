# lagoon add user-group

Add a user to a group in Lagoon.

## Synopsis

Add a user to a group in Lagoon.

```text
lagoon add user-group [flags]
```

## Options

```text
  -E, --email string   Email address of the user
  -h, --help           Help for user-group
  -N, --name string    Name of the group
  -R, --role string    Role in the group [owner, maintainer, developer, reporter, guest]
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

* [lagoon add](lagoon_add.md)     - Add a project, or add notifications and variables to projects or environments.

