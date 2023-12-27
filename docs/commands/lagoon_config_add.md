## lagoon config add

Add information about an additional Lagoon instance to use

```
lagoon config add [flags]
```

### Options

```
      --create-config         Create the config file if it is non existent (to be used with --config-file)
  -g, --graphql string        Lagoon GraphQL endpoint (eg, https://api.amazeeio.cloud/graphql)
  -h, --help                  help for add
  -H, --hostname string       Lagoon token endpoint hostname (eg, token.amazeeio.cloud)
      --keycloak-idp string   Optional: Lagoon Keycloak Identity Provider name.
                              	Set this to the name of the separate Identity Provider within keycloak if you use one.
                              	You may need to check with your Lagoon administrator if you use another SSO provider
  -K, --keycloak-url string   Lagoon Keycloak URL (eg, https://keycloak.amazeeio.cloud).
                              	Setting this will use keycloak for authentication instead of SSH based tokens. 
                              	Set 'ssh-token=true' to override.
                              	Note: SSH keys are still required for SSH access.
  -k, --kibana string         Optional: Lagoon Kibana URL (eg, https://logs.amazeeio.cloud)
  -P, --port string           Lagoon token endpoint port (22)
      --ssh-key string        SSH Key to use for this cluster for generating tokens
      --ssh-token             Set this context to only use ssh based tokens
                              	Set this to only use SSH based tokens if you're using the CLI in CI jobs or other automated processes
                              	where logging in via keycloak is not possible.
                              	This is enabled by default, it will be disabled by default in a future release. (default true)
  -t, --token string          Lagoon GraphQL token
  -u, --ui string             Optional: Lagoon UI location (eg, https://dashboard.amazeeio.cloud)
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
```

### SEE ALSO

* [lagoon config](lagoon_config.md)	 - Configure Lagoon CLI

