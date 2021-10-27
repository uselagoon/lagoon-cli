## lagoon add user-sshkey

Add an SSH key to a user

### Synopsis

Add an SSH key to a user

Examples:
Add key from public key file:
  lagoon add user-sshkey --email test@example.com --pubkey /path/to/id_rsa.pub

Add key by defining full key value:
  lagoon add user-sshkey --email test@example.com --keyvalue "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAINA0ITV2gbDc6noYeWaqfxTYpaEKq7HzU3+F71XGhSL/ my-computer@example"

Add key by defining full key value, but a specific key name:
  lagoon add user-sshkey --email test@example.com --keyname my-computer@example --keyvalue "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAINA0ITV2gbDc6noYeWaqfxTYpaEKq7HzU3+F71XGhSL/"

Add key by defining key value, but not specifying a key name (will default to try using the email address as key name):
  lagoon add user-sshkey --email test@example.com --keyvalue "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAINA0ITV2gbDc6noYeWaqfxTYpaEKq7HzU3+F71XGhSL/"



```
lagoon add user-sshkey [flags]
```

### Options

```
  -E, --email string      Email address of the user
  -h, --help              help for user-sshkey
  -N, --keyname string    Name of the SSH key (optional, if not provided will try use what is in the pubkey file)
  -V, --keyvalue string   Value of the public key to add (ssh-ed25519 AAA..)
  -K, --pubkey string     Specify path to the public key to add
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

* [lagoon add](lagoon_add.md)	 - Add a project, or add notifications and variables to projects or environments

