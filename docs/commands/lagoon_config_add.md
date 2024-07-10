## lagoon config add

Add information about an additional Lagoon instance to use

```
lagoon config add [flags]
```

### Options

```
      --create-config                    Create the config file if it is non existent (to be used with --config-file)
  -g, --graphql string                   Lagoon GraphQL endpoint
  -h, --help                             help for add
  -H, --hostname string                  Lagoon SSH hostname
  -k, --kibana string                    Lagoon Kibana URL (https://logs.amazeeio.cloud)
  -P, --port string                      Lagoon SSH port
      --publickey-identityfile strings   Specific public key identity files to use when doing ssh-agent checks (support multiple)
      --ssh-key string                   SSH Key to use for this cluster for generating tokens
  -t, --token string                     Lagoon GraphQL token
  -u, --ui string                        Lagoon UI location (https://dashboard.amazeeio.cloud)
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
      --ssh-publickey string   Specify path to a specific SSH public key to use for lagoon authentication using ssh-agent.
                               This will override any public key identities defined in configuration
  -v, --verbose                Enable verbose output to stderr (if supported)
```

### SEE ALSO

* [lagoon config](lagoon_config.md)	 - Configure Lagoon CLI

