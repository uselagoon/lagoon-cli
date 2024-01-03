## lagoon add deploytarget-config

Add deploytarget config to a project

```
lagoon add deploytarget-config [flags]
```

### Options

```
  -b, --branches string       Branches regex
  -d, --deploytarget uint     Deploytarget id
  -h, --help                  help for deploytarget-config
  -P, --pullrequests string   Pullrequests title regex
  -w, --weight uint           Deploytarget config weighting (default:1) (default 1)
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

* [lagoon add](lagoon_add.md)	 - Add a project, or add notifications and variables to projects or environments

