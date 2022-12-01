# Contributing

## Workflow
Pull Requests are welcome. In general, we follow the "fork-and-pull" Git workflow

 1. **Fork** the repo on GitHub
 2. **Clone** the project to your own machine
 3. **Commit** changes to your own branch
 4. **Push** your work back up to your fork
 5. Submit a **Pull request** so that we can review your changes

NOTE: Be sure to merge the latest from "upstream" before making a pull request!.

# Extending

## Code Generation

Some parts of the Lagoon CLI use code generation. Ensure you run `make gen` or `make test`.

## Documentation

When implementing new commands or updating flags, ensure to update the documentation. This can be done with the following

```
make docs
# or
GO111MODULE=on go run main.go --docs
```

# Testing

## Running Tests

Run tests locally

```
make test
```
