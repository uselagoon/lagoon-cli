# lagoon run invoke

## Synopsis

Invoke a task registered against an environment. 

```text
lagoon run invoke [flags]
```

The following are supported methods to use:

**Direct:** 

```text
lagoon run invoke -p example -e main -N "advanced task name"
```

## Options

```text
  -h, --help          Help for invoke
  -N, --name string   Name of the task that will be invoked
```

## Options inherited from parent commands

```text
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
  -i, --ssh-key string       Specify path to a specific SSH key to use for Lagoon authentication
```

## SEE ALSO

* [lagoon run](lagoon_run.md)     - Run a task against an environment.

