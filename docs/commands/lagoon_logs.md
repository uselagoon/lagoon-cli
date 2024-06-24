## lagoon logs

Display logs for a service of an environment and project

```
lagoon logs [flags]
```

### Options

```
  -c, --container string   specify a specific container name
  -f, --follow             continue outputting new lines as they are logged
  -h, --help               help for logs
  -n, --lines uint         the number of lines to return for each container (default 32)
  -s, --service string     specify a specific service name
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

* [lagoon](lagoon.md)	 - Command line integration for Lagoon

