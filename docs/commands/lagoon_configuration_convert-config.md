## lagoon configuration convert-config

Convert legacy .lagoon.yml config to the new configuration format

### Synopsis

Convert legacy .lagoon.yml config to the new configuration format.
This will prompt you to provide any required information if it is missing from your legacy configuration.
Running this command initially will run in dry-run mode, if you're happy with the result you can run it again
with the --write-config flag to save the new configuration.

```
lagoon configuration convert-config [flags]
```

### Options

```
  -h, --help           help for convert-config
      --write-config   Whether the config should be written to the config file or not
```

### Options inherited from parent commands

```
      --config-file string                Path to the config file to use (must be *.yml or *.yaml)
      --debug                             Enable debugging output (if supported)
  -e, --environment string                Specify an environment to use
      --force                             Force yes on prompts (if supported)
  -l, --lagoon string                     The Lagoon instance to interact with
      --no-header                         No header on table (if supported)
      --output-csv                        Output as CSV (if supported)
      --output-json                       Output as JSON (if supported)
      --pretty                            Make JSON pretty (if supported)
  -p, --project string                    Specify a project to use
      --skip-update-check                 Skip checking for updates
  -i, --ssh-key string                    Specify path to a specific SSH key to use for lagoon authentication
      --ssh-publickey string              Specify path to a specific SSH public key to use for lagoon authentication using ssh-agent.
                                          This will override any public key identities defined in configuration
      --strict-host-key-checking string   Similar to SSH StrictHostKeyChecking (accept-new, no, ignore) (default "accept-new")
  -v, --verbose                           Enable verbose output to stderr (if supported)
```

### SEE ALSO

* [lagoon configuration](lagoon_configuration.md)	 - Manage or view the contexts and users for interacting with Lagoon

