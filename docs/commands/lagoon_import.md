# lagoon import

Import a config from a yaml file

## Synopsis

Import a config from a yaml file. By default this command will exit on encountering an error \(such as an existing object\). You can get it to continue anyway with --keep-going. To disable any prompts, use --force.

```text
lagoon import [flags]
```

## Options

```text
  -h, --help                 help for import
  -I, --import-file string   path to the file to import
      --keep-going           on error, just log and continue instead of aborting
      --openshiftID uint     ID of the openshift to target for import
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

* [lagoon](lagoon.md)     - Command line integration for Lagoon

