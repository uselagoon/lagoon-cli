## lagoon get user-sshkeys

Get a user's SSH keys

### Synopsis

Get a user's SSH keys. This will only work for users that are part of a group

```
lagoon get user-sshkeys [flags]
```

### Options

```
  -E, --email string   New email address of the user
  -h, --help           help for user-sshkeys
```

### Options inherited from parent commands

```
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

### SEE ALSO

* [lagoon get](lagoon_get.md)	 - Get info on a resource

