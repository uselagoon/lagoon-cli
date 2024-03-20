## lagoon deploy pullrequest

Deploy a pullrequest

### Synopsis

Deploy a pullrequest
This pullrequest may not already exist as an environment in lagoon.

```
lagoon deploy pullrequest [flags]
```

### Options

```
  -N, --baseBranchName string   Pullrequest base branch name
  -R, --baseBranchRef string    Pullrequest base branch reference hash
      --buildvar strings        Add one or more build variables to deployment (--buildvar KEY1=VALUE1 [--buildvar KEY2=VALUE2])
  -H, --headBranchName string   Pullrequest head branch name
  -M, --headBranchRef string    Pullrequest head branch reference hash
  -h, --help                    help for pullrequest
  -n, --number uint             Pullrequest number
      --returnData              Returns the build name instead of success text
  -t, --title string            Pullrequest title
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

