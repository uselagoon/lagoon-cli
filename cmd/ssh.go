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
	Run: func(cmd *cobra.Command, args []string) {
		validateToken(lagoonCLIConfig.Current) // get a new token if the current one is invalid

		if cmdProjectName == "" || cmdProjectEnvironment == "" {
			fmt.Println("Missing arguments: Project name or environment name are not defined")
			cmd.Help()
			os.Exit(1)
		}

		// allow the use of the `feature/branch` and standard `feature-branch` type environment names to be used
		// since ssh requires the `feature-branch` type name to be used as the ssh username
		// run the environment through the makesafe and shorted functions that lagoon uses
		environmentName := makeSafe(shortenEnvironment(cmdProjectName, cmdProjectEnvironment))

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
			"hostname": lagoonCLIConfig.Lagoons[lagoonCLIConfig.Current].HostName,
			"port":     lagoonCLIConfig.Lagoons[lagoonCLIConfig.Current].Port,
			"username": cmdProjectName + "-" + environmentName,
			"sshkey":   privateKey,
		}
		if sshConnString {
			fmt.Println(generateSSHConnectionString(sshConfig, sshService, sshContainer))
		} else {

			// start an interactive ssh session
			authMethod, closeSSHAgent := publicKey(privateKey, skipAgent)
			config := &ssh.ClientConfig{
				User: sshConfig["username"],
				Auth: []ssh.AuthMethod{
					authMethod,
				},
				HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			}
			defer closeSSHAgent()
			var err error
			if sshCommand != "" {
				err = lagoonssh.RunSSHCommand(sshConfig, sshService, sshContainer, sshCommand, config)
			} else {
				err = lagoonssh.InteractiveSSH(sshConfig, sshService, sshContainer, config)
			}
			if err != nil {
				output.RenderError(err.Error(), outputOptions)
			}
		}

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
	connString := fmt.Sprintf("ssh -t %s-o \"UserKnownHostsFile=/dev/null\" -o \"StrictHostKeyChecking=no\" -p %v %s@%s", lagoon["sshKey"], lagoon["port"], lagoon["username"], lagoon["hostname"])
	if service != "" {
		connString = fmt.Sprintf("%s service=%s", connString, service)
	}
	if container != "" && service != "" {
		connString = fmt.Sprintf("%s container=%s", connString, container)
	}
	return connString
}
