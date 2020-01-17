## lagoon add user-group

Add user to a group in lagoon

### Synopsis

Add user to a group in lagoon

```
lagoon add user-group [flags]
```

### Options

```
  -E, --email string   Email address of the user
  -h, --help           help for user-group
  -N, --name string    Name of the group
  -R, --role string    Role in the group [owner, maintainer, developer, reporter, guest]
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

* [lagoon add](lagoon_add.md)	 - Add a project, or add notifications and variables to projects or environments

