# lagoon deploy pullrequest

Deploy a pull request.

## Synopsis

Deploy a pull request This pull request may not already exist as an environment in Lagoon.

```text
lagoon deploy pullrequest [flags]
```

## Options

```text
  -N, --baseBranchName string   Pull request base branch name
  -R, --baseBranchRef string    Pull request base branch reference hash
  -H, --headBranchName string   Pull request head branch name
  -M, --headBranchRef string    Pull request head branch reference hash
  -h, --help                    Help for pull request
  -n, --number uint             Pull request number
  -t, --title string            Pull request title
```

## Options inherited from parent commands

```text
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
  -i, --ssh-key string       Specify path to a specific SSH key to use for Lagoon authentication
```

## SEE ALSO

* [lagoon deploy](lagoon_deploy.md)     - Actions for deploying or promoting branches or environments in Lagoon.

