## lagoon add deploytarget

Add a DeployTarget to lagoon

### Synopsis

Add a Deploytarget(Kubernetes) to lagoon, this requires admin level permissions

```
lagoon add deploytarget [flags]
```

### Options

```
      --build-image string      DeployTarget build image to use (if different to the default)
      --cloud-provider string   DeployTarget cloud provider
      --cloud-region string     DeployTarget cloud region
      --console-url string      DeployTarget console URL
      --friendly-name string    DeployTarget friendly name
  -h, --help                    help for deploytarget
      --id uint                 ID of the DeployTarget
      --name string             Name of DeployTarget
      --router-pattern string   Router pattern of the DeployTarget, e.g. '${service}-${environment}-${project}.lagoon.example.com'
      --ssh-host string         DeployTarget ssh host
      --ssh-port string         DeployTarget ssh port
      --token string            DeployTarget token
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

* [lagoon add](lagoon_add.md)	 - Add a project, or add notifications and variables to projects or environments

