## lagoon add organization group

Add a group to an Organization

```
lagoon add organization group [flags]
```

### Options

```
  -G, --group string   Name of the group
  -h, --help           help for group
  -O, --name string    Name of the organization
      --org-owner      Flag to add the user to the group as an owner
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

* [lagoon add organization](lagoon_add_organization.md)	 - Add an organization, or add a group/project to an organization

