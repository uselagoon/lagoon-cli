package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	lagoonssh "github.com/uselagoon/lagoon-cli/pkg/lagoon/ssh"
	"github.com/uselagoon/lagoon-cli/pkg/output"
	"golang.org/x/crypto/ssh"
)

var sshConnString bool
var sshService string
var sshContainer string

var sshEnvCmd = &cobra.Command{
	Use:     "ssh",
	Aliases: []string{"s"},
	Short:   "Display the SSH command to access a specific environment in a project",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requiredInputCheck("Project name", cmdProjectName, "Environment name", cmdProjectEnvironment); err != nil {
			return err
		}
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		ignoreHostKey, acceptNewHostKey := lagoonssh.CheckStrictHostKey(strictHostKeyCheck)

		// allow the use of the `feature/branch` and standard `feature-branch` type environment names to be used
		// since ssh requires the `feature-branch` type name to be used as the ssh username
		// run the environment through the makesafe and shorted functions that lagoon uses
		environmentName := makeSafe(shortenEnvironment(cmdProjectName, cmdProjectEnvironment))
		sshHost, sshPort, username, _, err := getSSHHostPort(environmentName, debug)
		if err != nil {
			return fmt.Errorf("couldn't get SSH endpoint: %v", err)
		}
		// get private key that the cli is using
		skipAgent := false

		privateKey := fmt.Sprintf("%s/.ssh/id_rsa", userPath)
		// if the user has a key defined in their lagoon cli config, use it
		if lagoonCLIConfig.Lagoons[lagoonCLIConfig.Current].SSHKey != "" {
			privateKey = lagoonCLIConfig.Lagoons[lagoonCLIConfig.Current].SSHKey
			skipAgent = true
		}
		// otherwise check if one has been provided by the override flag
		if cmdSSHKey != "" {
			privateKey = cmdSSHKey
			skipAgent = true
		}
		sshConfig := map[string]string{
			"hostname": sshHost,
			"port":     sshPort,
			"username": username,
			"sshkey":   privateKey,
		}
		if sshConnString {
			fmt.Println(generateSSHConnectionString(sshConfig, sshService, sshContainer))
		} else {
			hkcb, hkalgo, err := lagoonssh.InteractiveKnownHosts(userPath, fmt.Sprintf("%s:%s", sshHost, sshPort), ignoreHostKey, acceptNewHostKey)
			if err != nil {
				return fmt.Errorf("couldn't get ~/.ssh/known_hosts: %v", err)
			}
			// start an interactive ssh session
			authMethod, closeSSHAgent := publicKey(privateKey, cmdPubkeyIdentity, lagoonCLIConfig.Lagoons[lagoonCLIConfig.Current].PublicKeyIdentities, skipAgent)
			config := &ssh.ClientConfig{
				User: sshConfig["username"],
				Auth: []ssh.AuthMethod{
					authMethod,
				},
				HostKeyCallback:   hkcb,
				HostKeyAlgorithms: hkalgo,
			}
			defer func() {
				err = closeSSHAgent()
				if err != nil {
					fmt.Fprintf(os.Stderr, "error closing ssh agent:%v\n", err)
				}
			}()
			if sshCommand != "" {
				err = lagoonssh.RunSSHCommand(sshConfig, sshService, sshContainer, sshCommand, config)
			} else {
				err = lagoonssh.InteractiveSSH(sshConfig, sshService, sshContainer, config)
			}
			if err != nil {
				output.RenderError(err.Error(), outputOptions)
				os.Exit(1)
			}
		}
		return nil
	},
}
var (
	sshCommand string
)

func init() {
	sshEnvCmd.Flags().StringVarP(&sshService, "service", "s", "", "specify a specific service name")
	sshEnvCmd.Flags().StringVarP(&sshContainer, "container", "c", "", "specify a specific container name")
	sshEnvCmd.Flags().BoolVarP(&sshConnString, "conn-string", "", false, "Display the full ssh connection string")
	sshEnvCmd.Flags().StringVarP(&sshCommand, "command", "C", "", "Command to run on remote")
}

// generateSSHConnectionString .
func generateSSHConnectionString(lagoon map[string]string, service string, container string) string {
	var connString string
	if service != "" {
		connString = "ssh -t"
	} else {
		connString = "ssh"
	}
	if lagoon["sshKey"] != "" {
		connString = fmt.Sprintf("%s -i %s", connString, lagoon["sshKey"])
	}
	if lagoon["port"] != "22" {
		connString = fmt.Sprintf("%s -p %v", connString, lagoon["port"])
	}
	connString = fmt.Sprintf("%s %s@%s", connString, lagoon["username"], lagoon["hostname"])
	if service != "" {
		connString = fmt.Sprintf("%s service=%s", connString, service)
	}
	if container != "" && service != "" {
		connString = fmt.Sprintf("%s container=%s", connString, container)
	}
	return connString
}
