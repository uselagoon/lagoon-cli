## lagoon config add

Add information about an additional Lagoon instance to use

### Synopsis

Add information about an additional Lagoon instance to use

```
lagoon config add [flags]
```

### Options

```
  -g, --graphql string    Lagoon GraphQL endpoint
  -h, --help              help for add
  -H, --hostname string   Lagoon SSH hostname
  -k, --kibana string     Lagoon Kibana URL (https://logs-db-ui-lagoon-master.ch.amazee.io)
  -P, --port string       Lagoon SSH port
  -t, --token string      Lagoon GraphQL token
  -u, --ui string         Lagoon UI location (https://ui-lagoon-master.ch.amazee.io)
```

### Options inherited from parent commands

```
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

* [lagoon config](lagoon_config.md)	 - Configure Lagoon CLI

