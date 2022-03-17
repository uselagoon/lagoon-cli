## lagoon retrieve backup

Retrieve a backup

### Synopsis

Retrieve a backup
Given a backup-id, you can initiate a retrieval for it.
You can check the status of the backup using the list backups or get backup command.

```
lagoon retrieve backup [flags]
```

### Options

```
  -B, --backup-id string   The backup ID you want to commence a retrieval for
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

* [lagoon retrieve](lagoon_retrieve.md)	 - Trigger a retrieval operation on backups

