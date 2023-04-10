## lagoon add notification project-email

Add an email notification to a project

### Synopsis

Add an email notification to a project
This command is used to add an existing email notification in Lagoon to a project.

```
lagoon add notification project-email [flags]
```

### Options

```
  -h, --help          help for project-email
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

* [lagoon add notification](lagoon_add_notification.md)	 - Add notifications or add notifications to projects

