## lagoon add user-sshkey

Add an sshkey to a user

### Synopsis

Add an sshkey to a user

```
lagoon add user-sshkey [flags]
```

### Options

```
  -E, --email string     Email address of the user
  -h, --help             help for user-sshkey
  -N, --keyname string   Name of the sshkey (optional, if not provided will try use what is in the pubkey file)
  -K, --pubkey string    file location to the public key to add
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
      --skip-update-check    Skip checking for updates
  -i, --ssh-key string       Specify path to a specific SSH key to use for lagoon authentication
```

### SEE ALSO

* [lagoon add](lagoon_add.md)	 - Add a project, or add notifications and variables to projects or environments

