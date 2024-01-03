## lagoon custom

Run a custom command

### Synopsis

Run a custom command.
This command alone does nothing, but you can create custom commands and put them into the custom commands directory,
these commands will then be available to use.
The directory for custom commands is ${HOME}/.lagoon-cli/commands.

```
lagoon custom [flags]
```

### Options

```
  -h, --help   help for custom
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

