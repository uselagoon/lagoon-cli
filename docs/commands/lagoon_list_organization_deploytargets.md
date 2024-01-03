## lagoon list organization deploytargets

List deploy targets in an organization

```
lagoon list organization deploytargets [flags]
```

### Options

```
  -h, --help          help for deploytargets
      --id uint       ID of the organization to list associated deploy targets for
  -O, --name string   Name of the organization to list associated deploy targets for
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

* [lagoon list organization](lagoon_list_organization.md)	 - List all organizations projects, groups, deploy targets or users

