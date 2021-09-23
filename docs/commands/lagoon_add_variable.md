# lagoon add variable

Add a variable to an environment or project.

## Synopsis

Add a variable to an environment or project.

```text
lagoon add variable [flags]
```

## Options

```text
  -h, --help           Help for variable
  -j, --json string    JSON string to patch
  -N, --name string    Name of the variable to add
  -S, --scope string   Scope of the variable[global, build, runtime, container_registry, internal_container_registry]
  -V, --value string   Value of the variable to add
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

* [lagoon add](lagoon_add.md)     - Add a project, or add notifications and variables to projects or environments.

