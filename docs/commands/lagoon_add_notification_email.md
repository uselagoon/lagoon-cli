## lagoon add notification email

Add a new email notification

### Synopsis

Add a new email notification
This command is used to set up a new email notification in Lagoon. This requires information to talk to the email address to send to.
It does not configure a project to send notifications to email though, you need to use project-email for that.

```
lagoon add notification email [flags]
```

### Options

```
  -E, --email string           The email address of the notification
  -h, --help                   help for email
  -n, --name string            The name of the notification
      --organization-id uint   ID of the Organization
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

* [lagoon add notification](lagoon_add_notification.md)	 - Add notifications or add notifications to projects

