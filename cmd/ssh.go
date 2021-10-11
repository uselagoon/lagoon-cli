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
		sshConfig := map[string]string{
			"hostname": lagoonCLIConfig.Lagoons[lagoonCLIConfig.Current].HostName,
			"port":     lagoonCLIConfig.Lagoons[lagoonCLIConfig.Current].Port,
			"username": cmdProjectName + "-" + cmdProjectEnvironment,
		}
		if sshConnString {
			fmt.Println(lagoonssh.GenerateSSHConnectionString(sshConfig, sshService, sshContainer))
		} else {
			// get private key that the cli is using
			skipAgent := false

			privateKey := fmt.Sprintf("%s/.ssh/id_rsa", userPath)
			if cmdSSHKey != "" {
				privateKey = cmdSSHKey
				skipAgent = true
			}
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
