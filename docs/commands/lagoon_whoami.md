## lagoon whoami

Whoami will return your user information for lagoon

### Synopsis

Whoami will return your user information for lagoon. 
This is useful if you have multiple keys or accounts in multiple lagoons and need to check which you are using.

```
lagoon whoami [flags]
```

### Options

```
  -h, --help                help for whoami
      --show-keys strings   Select which fields to display when showing SSH keys. Valid options (others are ignored): type,created,key,fingerprint
```

### Options inherited from parent commands

```
      --config-file string     Path to the config file to use (must be *.yml or *.yaml)
      --debug                  Enable debugging output (if supported)
  -e, --environment string     Specify an environment to use
      --force                  Force yes on prompts (if supported)
  -l, --lagoon string          The Lagoon instance to interact with
      --no-header              No header on table (if supported)
      --output-csv             Output as CSV (if supported)
      --output-json            Output as JSON (if supported)
      --pretty                 Make JSON pretty (if supported)
  -p, --project string         Specify a project to use
      --skip-update-check      Skip checking for updates
  -i, --ssh-key string         Specify path to a specific SSH key to use for lagoon authentication
      --ssh-publickey string   Specify path to a specific SSH public key to use for lagoon authentication using ssh-agent.
                               This will override any public key identities defined in configuration
  -v, --verbose                Enable verbose output to stderr (if supported)
```

### SEE ALSO

* [lagoon](lagoon.md)	 - Command line integration for Lagoon

