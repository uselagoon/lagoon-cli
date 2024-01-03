## lagoon add organization project

Add a project to an Organization

```
lagoon add organization project [flags]
```

### Options

```
      --auto-idle uint                          Auto idle setting of the project
      --availability string                     Availability of the project
      --branches string                         branches
      --build-image string                      Build Image for the project
      --deployments-disabled uint               Admin only flag for disabling deployments on a project, 1 to disable deployments, 0 to enable
      --development-build-priority uint         Set the priority of the development build
      --development-environments-limit uint     How many environments can be deployed at one time
      --facts-ui uint                           Enables the Lagoon insights Facts tab in the UI. Set to 1 to enable, 0 to disable
      --git-url string                          GitURL of the project
  -h, --help                                    help for project
  -O, --name string                             Name of the Organization to add the project to
      --openshift uint                          Reference to OpenShift Object this Project should be deployed to
      --openshift-project-pattern string        Pattern of OpenShift Project/Namespace that should be generated
      --org-owner                               Add the user as an owner of the project
      --private-key string                      Private key to use for the project
      --problems-ui uint                        Enables the Lagoon insights Problems tab in the UI. Set to 1 to enable, 0 to disable
      --production-build-priority uint          Set the priority of the production build
      --production-environment string           Production Environment for the project
      --pullrequests string                     Which Pull Requests should be deployed
      --router-pattern string                   Router pattern of the project, e.g. '${service}-${environment}-${project}.lagoon.example.com'
      --standby-production-environment string   Standby Production Environment for the project
      --storage-calc uint                       Should storage for this environment be calculated
      --subfolder string                        Set if the .lagoon.yml should be found in a subfolder useful if you have multiple Lagoon projects per Git Repository
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

* [lagoon add organization](lagoon_add_organization.md)	 - Add an organization, or add a group/project to an organization

