## lagoon add organization user

Add a user to an Organization

```
lagoon add organization user [flags]
```

### Options

```
  -E, --email string   Email address of the user
  -h, --help           help for user
  -O, --name string    Name of the organization
      --owner          Set the user as an owner of the organization
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

