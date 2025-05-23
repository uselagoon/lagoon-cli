## lagoon list

List projects, environments, deployments, variables or notifications

### Options

```
      --all-projects   All projects (if supported)
  -h, --help           help for list
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

* [lagoon](lagoon.md)	 - Command line integration for Lagoon
* [lagoon list all-users](lagoon_list_all-users.md)	 - List all users
* [lagoon list backups](lagoon_list_backups.md)	 - List an environments backups
* [lagoon list deployments](lagoon_list_deployments.md)	 - List deployments for an environment (alias: d)
* [lagoon list deploytarget-configs](lagoon_list_deploytarget-configs.md)	 - List deploytarget configs for a project
* [lagoon list deploytargets](lagoon_list_deploytargets.md)	 - List all DeployTargets in Lagoon
* [lagoon list environment-services](lagoon_list_environment-services.md)	 - Get information about an environments services
* [lagoon list environments](lagoon_list_environments.md)	 - List environments for a project (alias: e)
* [lagoon list group-projects](lagoon_list_group-projects.md)	 - List projects in a group (alias: gp)
* [lagoon list group-users](lagoon_list_group-users.md)	 - List all users in groups
* [lagoon list groups](lagoon_list_groups.md)	 - List groups you have access to (alias: g)
* [lagoon list invokable-tasks](lagoon_list_invokable-tasks.md)	 - Print a list of invokable tasks
* [lagoon list notification](lagoon_list_notification.md)	 - List all notifications or notifications on projects
* [lagoon list organization-admininstrators](lagoon_list_organization-admininstrators.md)	 - List admins in an organization
* [lagoon list organization-deploytargets](lagoon_list_organization-deploytargets.md)	 - List deploy targets in an organization
* [lagoon list organization-groups](lagoon_list_organization-groups.md)	 - List groups in an organization
* [lagoon list organization-projects](lagoon_list_organization-projects.md)	 - List projects in an organization
* [lagoon list organization-users](lagoon_list_organization-users.md)	 - List users in an organization
* [lagoon list organization-variables](lagoon_list_organization-variables.md)	 - List variables for an organization (alias: org-v)
* [lagoon list organizations](lagoon_list_organizations.md)	 - List all organizations
* [lagoon list project-groups](lagoon_list_project-groups.md)	 - List groups in a project (alias: pg)
* [lagoon list projects](lagoon_list_projects.md)	 - List all projects you have access to (alias: p)
* [lagoon list projects-by-metadata](lagoon_list_projects-by-metadata.md)	 - List projects by a given metadata key or key:value
* [lagoon list tasks](lagoon_list_tasks.md)	 - List tasks for an environment (alias: t)
* [lagoon list user-groups](lagoon_list_user-groups.md)	 - List a single users groups and roles
* [lagoon list variables](lagoon_list_variables.md)	 - List variables for a project or environment (alias: v)

