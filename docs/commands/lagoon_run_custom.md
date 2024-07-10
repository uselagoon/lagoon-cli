## lagoon run custom

Run a custom command on an environment

### Synopsis

Run a custom command on an environment
The following are supported methods to use
Direct:
  lagoon run custom -p example -e main -N "My Task" -S cli -c "ps -ef"

STDIN:
  cat /path/to/my-script.sh | lagoon run custom -p example -e main -N "My Task" -S cli

Path:
  lagoon run custom -p example -e main -N "My Task" -S cli -s /path/to/my-script.sh


```
lagoon run custom [flags]
```

### Options

```
  -c, --command string   The command to run in the task
  -h, --help             help for custom
  -N, --name string      Name of the task that will show in the UI (default: Custom Task) (default "Custom Task")
  -s, --script string    Path to bash script to run (will use this before command(-c) if both are defined)
  -S, --service string   Name of the service (cli, nginx, other) that should run the task (default: cli) (default "cli")
```

### Options inherited from parent commands

```
      --config-file string     Path to the config file to use (must be *.yml or *.yaml)
      --debug                  Enable debugging output (if supported)
  -e, --environment string     Specify an environment to use
      --force                  Force yes on prompts (if supported)
  -l, --lagoon string          The Lagoon instance to interact with
      --no-header              No header on table (if supported)
      --output-csv             Output as CSV (if supported)
      --output-json            Output as JSON (if supported)
      --pretty                 Make JSON pretty (if supported)
  -p, --project string         Specify a project to use
      --skip-update-check      Skip checking for updates
  -i, --ssh-key string         Specify path to a specific SSH key to use for lagoon authentication
      --ssh-publickey string   Specify path to a specific SSH public key to use for lagoon authentication using ssh-agent.
                               This will override any public key identities defined in configuration
  -v, --verbose                Enable verbose output to stderr (if supported)
```

### SEE ALSO

* [lagoon run](lagoon_run.md)	 - Run a task against an environment

