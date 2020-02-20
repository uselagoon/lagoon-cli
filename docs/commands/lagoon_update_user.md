## lagoon update user

Update a user in lagoon

### Synopsis

Update a user in lagoon (change name, or email address)

```
lagoon update user [flags]
```

### Options

```
  -C, --current-email string   Current email address of the user
  -E, --email string           New email address of the user
  -F, --firstName string       New firstname of the user
  -h, --help                   help for user
  -L, --lastName string        New lastname of the user
```

### Options inherited from parent commands

```
      --debug                Enable debugging output (if supported)
  -e, --environment string   Specify an environment to use
      --force                Force yes on prompts (if supported)
  -l, --lagoon string        The Lagoon instance to interact with
      --no-header            No header on table (if supported)
      --output-csv           Output as CSV (if supported)
      --output-json          Output as JSON (if supported)
      --pretty               Make JSON pretty (if supported)
  -p, --project string       Specify a project to use
  -i, --ssh-key string       Specify path to a specific SSH key to use for lagoon authentication
```

### SEE ALSO

* [lagoon update](lagoon_update.md)	 - Update a resource

