## lagoon get backup

Get a backup download link

### Synopsis

Get a backup download link
This returns a direct URL to the backup, this is a signed download link with a limited time to initiate the download (usually 5 minutes).

```
lagoon get backup [flags]
```

### Options

```
  -B, --backup-id string   The backup ID you want to restore
  -h, --help               help for backup
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

