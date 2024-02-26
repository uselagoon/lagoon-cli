## lagoon add organization organization

Add a new organization to Lagoon

```
lagoon add organization organization [flags]
```

### Options

```
      --description string       Description of the organization
      --environment-quota int    Environment quota for the organization
      --friendly-name string     Friendly name of the organization
      --group-quota int          Group quota for the organization
  -h, --help                     help for organization
  -O, --name string              Name of the organization
      --notification-quota int   Notification quota for the organization
      --project-quota int        Project quota for the organization
      --route-quota int          Route quota for the organization
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

