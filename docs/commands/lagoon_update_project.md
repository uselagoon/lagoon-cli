## lagoon update project

Update a project

### Synopsis

Update a project

```
lagoon update project [flags]
```

### Options

```
  -D, --activeSystemsDeploy string         Which internal Lagoon System is responsible for deploying 
  -P, --activeSystemsPromote string        Which internal Lagoon System is responsible for promoting
  -R, --activeSystemsRemove string         Which internal Lagoon System is responsible for promoting
  -T, --activeSystemsTask string           Which internal Lagoon System is responsible for tasks 
  -a, --autoIdle int                       Auto idle setting of the project
  -b, --branches string                    Which branches should be deployed
  -L, --developmentEnvironmentsLimit int   How many environments can be deployed at one time
  -g, --gitUrl string                      GitURL of the project
  -h, --help                               help for project
  -j, --json string                        JSON string to patch
  -S, --openshift int                      Reference to OpenShift Object this Project should be deployed to
  -o, --openshiftProjectPattern string     Pattern of OpenShift Project/Namespace that should be generated
  -I, --privateKey string                  Private key to use for the project
  -E, --productionEnvironment string       Which environment(the name) should be marked as the production environment
  -m, --pullrequests string                Which Pull Requests should be deployed
  -C, --storageCalc int                    Should storage for this environment be calculated
  -s, --subfolder string                   Set if the .lagoon.yml should be found in a subfolder useful if you have multiple Lagoon projects per Git Repository
```

### Options inherited from parent commands

```
      --all-projects         All projects (if supported)
      --debug                Enable debugging output (if supported)
  -e, --environment string   Specify an environment to use
      --force                Force (if supported)
  -l, --lagoon string        The Lagoon instance to interact with
      --no-header            No header on table (if supported)
      --output-csv           Output as CSV (if supported)
      --output-json          Output as JSON (if supported)
      --pretty               Make JSON pretty (if supported)
  -p, --project string       Specify a project to use
  -i, --ssh-key string       Specify a specific SSH key to use
      --version              Version information
```

### SEE ALSO

* [lagoon update](lagoon_update.md)	 - Update project, environment, or notification

