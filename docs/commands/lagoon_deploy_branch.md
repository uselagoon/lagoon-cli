## lagoon deploy branch

Deploy a new branch

### Synopsis

Deploy a new branch
This branch may or may not already exist in lagoon, if it already exists you may want to
use 'lagoon deploy latest' instead

```
lagoon deploy branch [flags]
```

### Options

```
  -b, --branch string      Branch name to deploy
  -r, --branchRef string   Branch ref to deploy
      --buildvar strings   Add one or more build variables to deployment (--buildvar KEY1=VALUE1 [--buildvar KEY2=VALUE2])
  -h, --help               help for branch
      --returnData         Returns the build name instead of success text
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

* [lagoon deploy](lagoon_deploy.md)	 - Actions for deploying or promoting branches or environments in lagoon

