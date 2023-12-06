## lagoon list group-users

List all users in groups

### Synopsis

List all users in groups in lagoon, this only shows users that are in groups.
If no group name is provided, all groups are queried.
Without a group name, this query may time out in large Lagoon installs.

```
lagoon list group-users [flags]
```

### Options

```
  -h, --help          help for group-users
  -N, --name string   Name of the group to list users in
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

* [lagoon list](lagoon_list.md)	 - List projects, environments, deployments, variables or notifications

