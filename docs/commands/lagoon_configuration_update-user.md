## lagoon configuration update-user

Update a Lagoon context user

```
lagoon configuration update-user [flags]
```

### Options

```
  -h, --help             help for update-user
      --name string      The name to reference this user as
      --ssh-key string   The full path to this users ssh-key
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
```

### SEE ALSO

* [lagoon configuration](lagoon_configuration.md)	 - Manage or view the contexts and users for interacting with Lagoon

