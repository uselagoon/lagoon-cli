## lagoon add variable

Add variables on environments or projects

### Synopsis

Add variables on environments or projects

```
lagoon add variable [flags]
```

### Options

```
  -h, --help           help for variable
  -j, --json string    JSON string to patch
  -N, --name string    Name of the variable to add
  -S, --scope string   Scope of the variable[global, build, runtime]
  -V, --value string   Value of the variable to add
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

* [lagoon add](lagoon_add.md)	 - Add a project, or add notifications and variables to projects or environments

