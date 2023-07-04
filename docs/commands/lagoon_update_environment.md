## lagoon update environment

Update an environment

### Synopsis

Update an environment

```
lagoon update environment [flags]
```

### Options

```
  -a, --auto-idle uint            Auto idle setting of the environment
      --deploy-base-ref string    Updates the deploy base ref
      --deploy-head-ref string    Updates the deploy head ref
  -d, --deploy-target uint        Reference to OpenShift Object this Environment should be deployed to
      --deploy-title string       Updates the deploy title
      --deploy-type string        Update the deploy type - BRANCH | PULLREQUEST | PROMOTE
      --environment-type string   Update the environment type - PRODUCTION | DEVELOPMENT
  -h, --help                      help for environment
      --namespace string          Update the namespace for this environment
      --route string              Update the route
      --routes string             Update the routes
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
