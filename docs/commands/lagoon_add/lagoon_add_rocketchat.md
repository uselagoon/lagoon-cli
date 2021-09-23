# lagoon add rocketchat

Add a new rocketchat notification

## Synopsis

Add a new rocketchat notification This command is used to set up a new rocketchat notification in lagoon. This requires information to talk to rocketchat like the webhook URL and the name of the channel. It does not configure a project to send notifications to rocketchat though, you need to use project-rocketchat for that.

```text
lagoon add rocketchat [flags]
```

## Options

```text
  -c, --channel string   The channel for the notification
  -h, --help             help for rocketchat
  -n, --name string      The name of the notification
  -w, --webhook string   The webhook URL of the notification
```

## Options inherited from parent commands

```text
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

## SEE ALSO

* [lagoon add](./)     - Add a project, or add notifications and variables to projects or environments

