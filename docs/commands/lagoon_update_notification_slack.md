## lagoon update notification slack

Update an existing Slack notification

```
lagoon update notification slack [flags]
```

### Options

```
  -c, --channel string   The channel for the notification
  -h, --help             help for slack
  -n, --name string      The current name of the notification
  -N, --newname string   The name of the notification
  -w, --webhook string   The webhook URL of the notification
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

* [lagoon update notification](lagoon_update_notification.md)	 - List all notifications or notifications on projects

