## lagoon deploy promote

Promote an environment

### Synopsis

Promote one environment to another

```
lagoon deploy promote [flags]
```

### Options

```
      --buildvar stringArray   Add one or more build variables to deployment (--buildvar KEY1=VALUE1 [--buildvar KEY2=VALUE2])
  -d, --destination string     Destination environment name to create
  -h, --help                   help for promote
      --returndata             Returns the build name instead of success text
  -s, --source string          Source environment name to use as the base to deploy from
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

* [lagoon deploy](lagoon_deploy.md)	 - Actions for deploying or promoting branches or environments in lagoon

