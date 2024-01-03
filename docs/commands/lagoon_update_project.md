## lagoon update project

Update a project

```
lagoon update project [flags]
```

### Options

```
  -a, --autoIdle int                          Auto idle setting of the project
  -b, --branches string                       Which branches should be deployed
      --deploymentsDisabled int               Admin only flag for disabling deployments on a project, 1 to disable deployments, 0 to enable
  -L, --developmentEnvironmentsLimit int      How many environments can be deployed at one time
      --factsUi int                           Enables the Lagoon insights Facts tab in the UI. Set to 1 to enable, 0 to disable
  -g, --gitUrl string                         GitURL of the project
  -h, --help                                  help for project
  -j, --json string                           JSON string to patch
  -N, --name string                           Change the name of the project by specifying a new name (careful!)
  -S, --openshift int                         Reference to OpenShift Object this Project should be deployed to
  -o, --openshiftProjectPattern string        Pattern of OpenShift Project/Namespace that should be generated
  -I, --privateKey string                     Private key to use for the project
      --problemsUi int                        Enables the Lagoon insights Problems tab in the UI. Set to 1 to enable, 0 to disable
  -E, --productionEnvironment string          Which environment(the name) should be marked as the production environment
  -m, --pullrequests string                   Which Pull Requests should be deployed
  -Z, --routerPattern string                  Router pattern of the project, e.g. '${service}-${environment}-${project}.lagoon.example.com'
      --standbyProductionEnvironment string   Which environment(the name) should be marked as the standby production environment
  -C, --storageCalc int                       Should storage for this environment be calculated
  -s, --subfolder string                      Set if the .lagoon.yml should be found in a subfolder useful if you have multiple Lagoon projects per Git Repository
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

