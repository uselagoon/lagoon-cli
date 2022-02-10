## lagoon apply

Apply the configuration of workflows or tasks from a given yaml configuration file

### Synopsis

Apply the configuration of workflows or tasks from a given yaml configuration file.
Workflows or advanced task definitions will be created if they do not already exist.

```
lagoon apply [flags]
```

### Options

```
  -t, --advanced-tasks   Target advanced tasks only
  -f, --file string      lagoon apply (-f FILENAME) [options]
  -h, --help             help for apply
  -w, --workflows        Target workflows only
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
* [lagoon apply set-last-applied](lagoon_apply_set-last-applied.md)	 - Set the latest applied workflows or advanced task definitions for project/environment.
* [lagoon apply view-last-applied](lagoon_apply_view-last-applied.md)	 - View the latest applied workflows or advanced task definitions for project/environment.

