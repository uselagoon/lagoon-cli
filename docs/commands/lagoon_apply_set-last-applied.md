## lagoon apply set-last-applied

Set the latest applied workflows or advanced task definitions for project/environment.

### Synopsis

Finds latest configuration match by workflow/task definition 'Name' and sets the latest applied workflow or advanced task definition for project/environment with the contents of file.

```
lagoon apply set-last-applied -f FILENAME [flags]
```

### Options

```
  -h, --help   help for set-last-applied
```

### Options inherited from parent commands

```
      --config-file string   Path to the config file to use (must be *.yml or *.yaml)
      --debug                Enable debugging output (if supported)
  -e, --environment string   Specify an environment to use
  -f, --file string          lagoon apply (-f FILENAME) [options]
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

* [lagoon apply](lagoon_apply.md)	 - Apply the configuration of workflows or tasks from a given yaml configuration file

