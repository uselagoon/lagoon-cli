## lagoon add group

Add a group to Lagoon, or add a group to an organization

### Synopsis

To add a group to an organization, you'll need to include the `organization` flag and provide the name of the organization. You need to be an owner of this organization to do this.
If you're the organization owner and want to grant yourself ownership to this group to be able to deploy projects that may be added to it, specify the `owner` flag, otherwise you will still be able to add and remove users without being an owner

```
lagoon add group [flags]
```

### Options

```
  -h, --help                       help for group
  -N, --name string                Name of the group
  -O, --organization-name string   Name of the organization
      --owner                      Organization owner only: Flag to grant yourself ownership of this group
```

### Options inherited from parent commands

```
      --config-file string                Path to the config file to use (must be *.yml or *.yaml)
      --debug                             Enable debugging output (if supported)
  -e, --environment string                Specify an environment to use
      --force                             Force yes on prompts (if supported)
  -l, --lagoon string                     The Lagoon instance to interact with
      --no-header                         No header on table (if supported)
      --output-csv                        Output as CSV (if supported)
      --output-json                       Output as JSON (if supported)
      --pretty                            Make JSON pretty (if supported)
  -p, --project string                    Specify a project to use
      --skip-update-check                 Skip checking for updates
  -i, --ssh-key string                    Specify path to a specific SSH key to use for lagoon authentication
      --ssh-publickey string              Specify path to a specific SSH public key to use for lagoon authentication using ssh-agent.
                                          This will override any public key identities defined in configuration
      --strict-host-key-checking string   Similar to SSH StrictHostKeyChecking (accept-new, no, ignore) (default "accept-new")
  -v, --verbose                           Enable verbose output to stderr (if supported)
```

### SEE ALSO

* [lagoon add](lagoon_add.md)	 - Add a project, or add notifications and variables to projects or environments

