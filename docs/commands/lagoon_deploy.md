## lagoon deploy

Actions for deploying or promoting branches or environments in lagoon

### Options

```
  -h, --help   help for deploy
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

* [lagoon](lagoon.md)	 - Command line integration for Lagoon
* [lagoon deploy branch](lagoon_deploy_branch.md)	 - Deploy a new branch
* [lagoon deploy latest](lagoon_deploy_latest.md)	 - Deploy latest environment
* [lagoon deploy promote](lagoon_deploy_promote.md)	 - Promote an environment
* [lagoon deploy pullrequest](lagoon_deploy_pullrequest.md)	 - Deploy a pullrequest

