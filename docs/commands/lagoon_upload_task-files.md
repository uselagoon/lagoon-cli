## lagoon upload task-files

Upload files to a task by its ID

### Synopsis

Upload files to a task by its ID

```
lagoon upload task-files [flags]
```

### Options

```
  -F, --file strings   File to upload (add multiple flags to upload multiple files)
  -h, --help           help for task-files
  -I, --id int         ID of the task
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

* [lagoon upload](lagoon_upload.md)	 - Upload files to tasks

