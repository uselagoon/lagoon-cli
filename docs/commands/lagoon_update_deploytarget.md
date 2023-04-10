## lagoon update deploytarget

Update a DeployTarget in lagoon

### Synopsis

Update a DeployTarget (kubernetes or openshift) in lagoon, this requires admin level permissions

```
lagoon update deploytarget [flags]
```

### Options

```
      --build-image string      DeployTarget build image to use (if different to the default, use "" to clear)
      --cloud-provider string   DeployTarget cloud provider
      --cloud-region string     DeployTarget cloud region
      --console-url string      DeployTarget console URL
      --friendly-name string    DeployTarget friendly name
  -h, --help                    help for deploytarget
      --id uint                 ID of the DeployTarget
      --router-pattern string   DeployTarget router-pattern
      --ssh-host string         DeployTarget ssh host
      --ssh-port string         DeployTarget ssh port
      --token string            DeployTarget token
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

* [lagoon update](lagoon_update.md)	 - Update a resource

