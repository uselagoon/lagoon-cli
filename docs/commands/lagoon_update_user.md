# lagoon update user

Update a user in Lagoon.

## Synopsis

Update a user in Lagoon \(change name, or email address\).

```text
lagoon update user [flags]
```

## Options

```text
  -C, --current-email string   Current email address of the user
  -E, --email string           New email address of the user
  -F, --firstName string       New firstname of the user
  -h, --help                   Help for user
  -L, --lastName string        New lastname of the user
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

* [lagoon update](lagoon_update.md)     - Update a resource.

