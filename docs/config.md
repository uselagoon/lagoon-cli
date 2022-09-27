# Introduction

By default the CLI is configured to use the `amazeeio` Lagoon. But you can also define additional Lagoons if you need to.

The `.lagoon.yml` file will be installed in your home directory by default

## Layout of the configuration file
The configuration file is laid out like below
```yaml
current: amazeeio
default: amazeeio
lagoons:
  amazeeio:
    graphql: https://api.lagoon.amazeeio.cloud/graphql
    hostname: ssh.lagoon.amazeeio.cloud
    port: 32222
    token: ey.....xA
```
There are a few sections to cover off

* `current` is the current Lagoon that you will be using, if you only have the one, it will be `amazeeio`
* `default` is the default Lagoon to use, if you always use a particular Lagoon then you can set your preference as your default
* `lagoons` is where the actual connection parameters are stored for each Lagoon, they all follow the same template.
    * `graphql` is the graphql endpoint
    * `hostname` is the ssh hostname
    * `port` is the ssh port
    * `token` is the graphql token, this is automatically generate the first time you `lagoon login` and will automatically refresh if it expires via ssh.

# Add a Lagoon
If you want to add a different Lagoon to use, then you can use the CLI command to view the flags available
```bash
lagoon config add --lagoon LagoonName
```
## Example
```bash
lagoon config add --lagoon amazeeio \
    --graphql https://api.lagoon.amazeeio.cloud/graphql \
    --hostname ssh.lagoon.amazeeio.cloud \
    --port 32222
```

# Delete a Lagoon
If you want to remove a Lagoon, you can use
```bash
lagoon config delete --lagoon LagoonName
```
## Example
```bash
lagoon config delete --lagoon amazeeio
```

# Change default Lagoon
If you add additional Lagoons, you can select which one is the default you want to use by running
```bash
lagoon config default --lagoon LagoonName
```
## Example
```bash
lagoon config default --lagoon amazeeio
```

# Use a different Lagoon
If you want to temporarily use a different Lagoon, when you run any commands you can specify the flag `--lagoon` or `-l` and then the name of the Lagoon
## Example
```bash
lagoon --lagoon mylagoon list projects
```

# View Lagoons
You can view all the Lagoons you have configured by running
```bash
lagoon config list
```
Output
```yaml
You have the following lagoons configured:
Name: amazeeio
 - Hostname: ssh.lagoon.amazeeio.cloud
 - GraphQL: https://api.lagoon.amazeeio.cloud/graphql
 - Port: 32222
Name: mylagoon
 - Hostname: ssh.mylagoon.example
 - GraphQL: https://api.mylagoon.example/graphql
 - Port: 32000
Name: local
 - Hostname: localhost
 - GraphQL: http://localhost:3000/graphql
 - Port: 2020

Your default lagoon is:
Name: local

Your current lagoon is:
Name: local
```
