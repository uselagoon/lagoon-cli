## lagoon list projects-by-metadata

List projects by a given metadata key or key:value

```
lagoon list projects-by-metadata [flags]
```

### Options

```
  -h, --help            help for projects-by-metadata
  -K, --key string      The key name of the metadata value you are querying on
      --show-metadata   Show the metadata for each project as another field (this could be a lot of data)
  -V, --value string    The value for the key you are querying on
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

* [lagoon list](lagoon_list.md)	 - List projects, environments, deployments, variables or notifications

