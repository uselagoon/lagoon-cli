## lagoon add organization deploytarget

Add a deploy target to an Organization

```
lagoon add organization deploytarget [flags]
```

### Options

```
  -D, --deploy-target uint   ID of DeployTarget
  -h, --help                 help for deploytarget
  -O, --name string          Name of Organization
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

