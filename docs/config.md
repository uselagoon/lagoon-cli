# Introduction

By default the CLI is configured to use the `amazeeio` Lagoon. But you can also define additional Lagoon contexts if you need to.

The `.lagoon.yml` file will be installed in your home directory by default

## Layout of the configuration file
The configuration file is laid out like below
```yaml
current: amazeeio
default: amazeeio
lagoons:
  amazeeio:
    graphql: https://api.amazeeio.cloud/graphql
    hostname: token.amazeeio.cloud
    port: 22
    token: ey.....xA
```
There are a few sections to cover off

* `current` is the current Lagoon context that you will be using, if you only have the one, it will be `amazeeio`
* `default` is the default Lagoon context to use, if you always use a particular context then you can set your preference as your default
* `lagoons` is where the actual connection parameters are stored for each Lagoon context, they all follow the same template.
    * `graphql` is the graphql endpoint
    * `hostname` is the ssh token hostname
    * `port` is the ssh token port
    * `token` is the graphql token, this is automatically generate the first time you `lagoon login` and will automatically refresh if it expires via ssh.

# Add a Lagoon context
If you want to add a different Lagoon context to use, then you can use the CLI command to view the flags available
```bash
lagoon config add --lagoon example
```

It is also possible to set an SSH key or public key itendity file to use when adding (or updating) a context. 

The flag `--publickey-identitfyfile /path/to/key.pub` can be used to tell the CLI which public key identity in your SSH agent should be used.

The flag `--ssh-key /path/to/key` can be used to tell the CLI which SSH private key should be used.

For troubleshooting, see [troubleshooting SSH keys](./troubleshooting.md#ssh-keys)

## Example
```bash
lagoon config add --lagoon amazeeio \
    --graphql https://api.amazeeio.cloud/graphql \
    --hostname token.amazeeio.cloud \
    --port 22
```

# Delete a Lagoon context
If you want to remove a Lagoon context, you can use
```bash
lagoon config delete --lagoon example
```
## Example
```bash
lagoon config delete --lagoon amazeeio
```

# Change default Lagoon context
If you add additional Lagoon contexts, you can select which one is the default you want to use by running
```bash
lagoon config default --lagoon example
```
## Example
```bash
lagoon config default --lagoon amazeeio
```

# Use a different Lagoon context
If you want to temporarily use a different Lagoon context, when you run any commands you can specify the flag `--lagoon` or `-l` and then the name of the context
## Example
```bash
lagoon --lagoon example list projects
```

# View Lagoon contexts
You can view all the Lagoon contexts you have configured by running
```bash
lagoon config list
```
Output
```yaml
NAME                      	VERSION	GRAPHQL                            	SSH-HOSTNAME               	SSH-PORT	SSH-KEY
amazeeio(default)(current)	v2.22.0	https://api.amazeeio.cloud/graphql 	token.amazeeio.cloud       	22   	-
example                   	v2.22.0	https://api.example.com/graphql  	token.example.com          	22    	-
```
