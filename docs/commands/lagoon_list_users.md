## lagoon list users

List all users (alias: u)

### Synopsis

List all users in groups in lagoon, this only shows users that are in groups.

```
lagoon list users [flags]
```

### Options

```
  -h, --help          help for users
  -N, --name string   Name of the group to list users in (if not specified, will default to all groups)
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

* [lagoon list](lagoon_list.md)	 - List projects, deployments, variables or notifications

