## lagoon get deployment

Get a deployment by name

### Synopsis

Get a deployment by name
This returns information about a deployment, the logs of this build can also be retrieved

```
lagoon get deployment [flags]
```

### Options

```
  -h, --help          help for deployment
  -L, --logs          Show the build logs if available
  -N, --name string   The name of the deployment (eg, lagoon-build-abcdef)
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

