## lagoon update project

Update a project

```
lagoon update project [flags]
```

### Options

```
  -a, --auto-idle                               Auto idle setting of the project. Set to enable, --auto-idle=false to disable
      --availability string                     Availability of the project
  -b, --branches string                         Which branches should be deployed
      --build-image string                      Build Image for the project. Set to 'null' to remove the build image
      --deployments-disabled                    Admin only flag for disabling deployments on a project. Set to disable deployments, --deployments-disabled=false to enable
  -S, --deploytarget uint                       Reference to Deploytarget(Kubernetes) this Project should be deployed to
  -o, --deploytarget-project-pattern string     Pattern of Deploytarget(Kubernetes) Project/Namespace that should be generated
      --development-build-priority uint         Set the priority of the development build
  -L, --development-environments-limit uint     How many environments can be deployed at one time
      --facts-ui                                Enables the Lagoon insights Facts tab in the UI. Set to enable, --facts-ui=false to disable
  -g, --git-url string                          GitURL of the project
  -h, --help                                    help for project
  -j, --json string                             JSON string to patch
  -N, --name string                             Change the name of the project by specifying a new name (careful!)
  -I, --private-key string                      Private key to use for the project
      --problems-ui                             Enables the Lagoon insights Problems tab in the UI. Set to enable, --problems-ui=false to disable
      --production-build-priority uint          Set the priority of the production build
  -E, --production-environment string           Which environment(the name) should be marked as the production environment
  -m, --pullrequests string                     Which Pull Requests should be deployed
  -Z, --router-pattern string                   Router pattern of the project, e.g. '${service}-${environment}-${project}.lagoon.example.com'
      --standby-production-environment string   Which environment(the name) should be marked as the standby production environment
  -C, --storage-calc                            Should storage for this environment be calculated. Set to enable, --storage-calc=false to disable
  -s, --subfolder string                        Set if the .lagoon.yml should be found in a subfolder useful if you have multiple Lagoon projects per Git Repository
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

* [lagoon update](lagoon_update.md)	 - Update a resource

