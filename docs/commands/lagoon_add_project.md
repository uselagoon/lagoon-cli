## lagoon add project

Add a new project to lagoon

### Synopsis

Add a new project to lagoon

```
lagoon add project [flags]
```

### Options

```
  -a, --autoIdle int                   Auto idle setting of the project
  -b, --branches string                Which branches should be deployed
  -g, --gitUrl string                  GitURL of the project
  -h, --help                           help for project
  -j, --json string                    JSON string to patch
  -I, --privateKey string              Private key to use for the project
  -E, --productionEnvironment string   Which environment(the name) should be marked as the production environment
  -m, --pullrequests string            Which Pull Requests should be deployed
  -C, --storageCalc int                Should storage for this environment be calculated
  -s, --subfolder string               Set if the .lagoon.yml should be found in a subfolder Usefull if you have multiple Lagoon projects per Git Repository
```

### Options inherited from parent commands

```
      --all-projects         All projects (if supported)
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

* [lagoon add](lagoon_add.md)	 - Add a project, or add notifications and variables to projects or environments

