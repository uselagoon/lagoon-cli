## lagoon delete project-slack

Delete a Slack notification from a project

### Synopsis

Delete a Slack notification from a project

```
lagoon delete project-slack [flags]
```

### Options

```
  -h, --help          help for project-slack
  -n, --name string   The name of the notification
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

* [lagoon delete](lagoon_delete.md)	 - Delete a project, or delete notifications and variables from projects or environments

