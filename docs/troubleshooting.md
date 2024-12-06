This page attempts to describe some common problems and provide basic troubleshooting guides on what the issue could be and how to attempt a resolution.

Since every issue can be slightly different, there may be cases where some or none of the steps outlined can help. In these cases, you may have encountered a bug with the CLI or the Lagoon API. Before raising an issue, it is important to try and determine if the issue is related to the API or the CLI.

# Raising Issues
If raising issues against the CLI, please ensure you include which [version](commands/lagoon_version.md) of the CLI you're using in your report. 

If your issues appears to be permissions or API related, you may need to consult with your hosting provider before raising any issues.

# Common Problems

## Refresh token
Sometimes simply refreshing your auth token can help.

Use [lagoon login](commands/lagoon_login.md) to attempt to retrieve a new token. 

You can also use the flag `--verbose` to see information about the authentication attempt being made and which keys may be being used.

```bash title="verbose login example"
$ lagoon login --verbose
ssh: attempting connection using any keys in ssh-agent

$ lagoon login --verbose --ssh-publickey /home/user/.ssh/id_ed25519.pub
ssh: attempting connection using identity file public key: /home/user/.ssh/id_ed25519.pub

$ lagoon login --verbose --ssh-key /home/user/.ssh/id_ed25519
ssh: attempting connection using private key: /home/user/.ssh/id_ed25519
```

## SSH keys
The CLI will prefer to use your ssh-agent when providing keys to the authentication endpoint. If you haven't got any keys in your ssh-agent, the CLI will fall back to checking `~/.ssh/id_rsa`.

!!! Note "Which keys are in my agent?"
    You can check the keys in your ssh-agent using `ssh-add -L`

Some flags and configurations exist that allow you to force which key in your ssh-agent is used, or which SSH key file you want to use.

### specific key in agent
If you know which public key you want to use when interacting Lagoon, you can use the following flag

```bash
--ssh-publickey /path/to/privatekey.pub
```

This will attempt to use this key if it is present in your ssh-agent. Alternatively, you can add `publickeyidentities` to your context in your Lagoon CLI configuration file.

```yaml title="~/.lagoon.yml CLI configuration "
lagoons:
    contextname:
        publickeyidentities:
            - /path/to/privatekey.pub
```

### specific private key
If you know which private key you want to use when interacting Lagoon, you can use the following flag

```bash
--ssh-key /path/to/privatekey
```

Alternatively, you can add `sshkey` to your context in your Lagoon CLI configuration file.

```yaml title="~/.lagoon.yml CLI configuration "
lagoons:
    contextname:
        sshkey: /path/to/privatekey
```

## Unauthorized/permission errors
In the event you experience an unauthorized or permission error, it is important to confirm which user has authenticated against the Lagoon API.

The CLI provides a command [whoami](commands/lagoon_whoami.md) which will query the Lagoon API to determine which user the API is seeing.

```title="lagoon whoami"
ID                  	EMAIL                	FIRSTNAME	LASTNAME	SSHKEYS
95101742-07f373b0bfe3	bob.johnson@example.com	Bob      	Johnson 	3
```

### correct user
If the user returned is the correct user, then you need to ensure that commands you're executing you have permission to do so with. Some things to check:

* That your user actually has access to the resource (project, environment, group, etc..) that you're interacting with
    * check your user with the [list user-groups](commands/lagoon_list_user-groups.md) command (you will only be able to run this for your own user).
    * some options are only available to `platform-owners` or `organization` administrators, and you'll get permission errors you can check against [Lagoon RBAC](https://docs.lagoon.sh/interacting/rbac/).
* Ensure you've typed the command and all the flags required correctly. Sometimes a typo can present as a permission error.
* Check that you're using the right command for what you're trying to do, the `--help` flag on any command will provide information that may be helpful.

### incorrect user
If the user is not the user you're expecting, then it is likely you have multiple user accounts, or multiple SSH keys. To help resolve this, its important to know:

* How many users you've got
* How many SSH keys you've got
* Which SSH keys are attached to which user
    * check your user with the [get user-sshkeys](commands/lagoon_get_user-sshkeys.md) command (you will only be able to run this for your own user). 
* If you're using an SSH agent, [which keys are loaded](#ssh-keys) into it
* [Login with verbosity](#refresh-token) enabled to see how the CLI is attempting to authenticate

## Multiple user accounts or contexts
The CLI supports multiple contexts, think of these as a way to interact with multiple Lagoon instances, or a way to select a user if you have multiple accounts.

An example of how a configuration file looks for a user with multiple accounts is shown below. See [configuration](config.md) on how to add more contexts, or modify your config file manually.
```yaml title="~/.lagoon.yml Multiple accounts"
current: user1
default: user1
lagoons:
    user1:
        graphql: https://api.example.com/graphql
        hostname: ssh.example.com
        port: "22"
        publickeyidentities:
          - /path/to/user1/public.key
    user2:
        graphql: https://api.example.com/graphql
        hostname: ssh.example.com
        port: "22"
        publickeyidentities:
          - /path/to/user2/public.key
```

You'll see that there are 2 `lagoons` (or contexts), one for `user1` which the publickey override, and one for `user2` with a different publickey override.


### Selecting a context
When you want to use the CLI to perform actions as `user1`, you don't need to provide any flags, as the `default` in this example is `user1`.

If you need to perform actions as `user2`, you will have to [select a context](./config.md#use-a-different-lagoon-context) or [change the default](./config.md#change-default-lagoon-context).

!!! Note "Changing context"
    Sometimes changing your context could result in a token not refreshing correctly. Firstly attempt a [token refresh](#refresh-token) with your selected context.