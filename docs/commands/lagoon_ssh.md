## lagoon ssh

Display the SSH command to access a specific environment in a project

### Synopsis

Display the SSH command to access a specific environment in a project

```
lagoon ssh [flags]
```

### Options

```
  -C, --command string     Command to run on remote
      --conn-string        Display the full ssh connection string
  -c, --container string   specify a specific container name
  -h, --help               help for ssh
  -s, --service string     specify a specific service name
```

### Options inherited from parent commands

```
      --all-projects         All projects (if supported)
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

* [lagoon](lagoon.md)	 - Command line integration for Lagoon

