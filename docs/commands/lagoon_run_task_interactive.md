## lagoon run task interactive

Interactively run a custom task against an environment

### Synopsis

Interactively run a custom task against an environment
Provides prompts for arguments
example:
 lagoon run invoke interactive -p example -e main


```
lagoon run task interactive [flags]
```

### Options

```
  -h, --help   help for interactive
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

* [lagoon run task](lagoon_run_task.md)	 - Run a custom task registered against an environment

