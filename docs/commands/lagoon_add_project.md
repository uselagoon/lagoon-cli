## lagoon add project

Add a new project to Lagoon

```
lagoon add project [flags]
```

### Options

```
  -a, --autoIdle int                          Auto idle setting of the project
  -b, --branches string                       Which branches should be deployed
  -L, --developmentEnvironmentsLimit int      How many environments can be deployed at one time
  -g, --gitUrl string                         GitURL of the project
  -h, --help                                  help for project
  -j, --json string                           JSON string to patch
  -S, --openshift int                         Reference to OpenShift Object this Project should be deployed to
  -o, --openshiftProjectPattern string        Pattern of OpenShift Project/Namespace that should be generated
  -I, --privateKey string                     Private key to use for the project
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

* [lagoon add](lagoon_add.md)	 - Add a project, or add notifications and variables to projects or environments

