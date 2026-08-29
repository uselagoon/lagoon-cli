## lagoon add project

Add a new project to Lagoon, or add a project to an organization

### Synopsis

To add a project to an organization, you'll need to include the `organization` flag and provide the name of the organization. You need to be an owner of this organization to do this.
If you're the organization owner and want to grant yourself ownership to this project to be able to deploy environments, specify the `owner` flag.

```
lagoon add project [flags]
```

### Options

```
  -a, --auto-idle                               Auto idle setting of the project. Set to enable, --auto-idle=false to disable
  -b, --branches string                         Which branches should be deployed
      --build-image string                      Build Image for the project
  -S, --deploytarget uint                       Reference to Deploytarget(Kubernetes) target this Project should be deployed to
      --deploytarget-project-pattern string     Pattern of Deploytarget(Kubernetes) Project/Namespace that should be generated
  -L, --development-environments-limit uint     How many environments can be deployed at one time
  -g, --git-url string                          GitURL of the project
  -h, --help                                    help for project
      --interactive                             Set Interactive mode for the project creation wizard.
  -j, --json string                             JSON string to patch
      --organization-id uint                    ID of the Organization to add the project to
  -O, --organization-name string                Name of the Organization to add the project to
      --owner                                   Add the user as an owner of the project
  -I, --private-key string                      Private key to use for the project
  -E, --production-environment string           Which environment(the name) should be marked as the production environment
  -m, --pullrequests string                     Which Pull Requests should be deployed
  -Z, --router-pattern string                   Router pattern of the project, e.g. '${service}-${environment}-${project}.lagoon.example.com'
      --standby-production-environment string   Which environment(the name) should be marked as the standby production environment
  -C, --storage-calc                            Should storage for this environment be calculated. Set to enable, --storage-calc=false to disable
  -s, --subfolder string                        Set if the .lagoon.yml should be found in a subfolder useful if you have multiple Lagoon projects per Git Repository
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

* [lagoon add](lagoon_add.md)	 - Add a project, or add notifications and variables to projects or environments

