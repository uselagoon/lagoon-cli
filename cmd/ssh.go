package cmd

import (
	"fmt"
	"os"

	lagoonssh "github.com/amazeeio/lagoon-cli/lagoon/ssh"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
)

var sshConnString bool
var sshService string
var sshContainer string

var sshEnvCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Display the SSH command to access a specific environment in a project",
	Run: func(cmd *cobra.Command, args []string) {
		validateToken(viper.GetString("current")) // get a new token if the current one is invalid

		if cmdProjectName == "" || cmdProjectEnvironment == "" {
			fmt.Println("Not enough arguments. Requires: project name and environment name")
			cmd.Help()
			os.Exit(1)
		}
		sshConfig := map[string]string{
			"hostname": viper.GetString("lagoons." + cmdLagoon + ".hostname"),
			"port":     viper.GetString("lagoons." + cmdLagoon + ".port"),
			"username": cmdProjectName + "-" + cmdProjectEnvironment,
		}
		if sshConnString {
			fmt.Println(lagoonssh.GenerateSSHConnectionString(sshConfig, sshService, sshContainer))
		} else {
			// get private key that the cli is using
			homeDir, _ := os.UserHomeDir()
			skipAgent := false

			privateKey := fmt.Sprintf("%s/.ssh/id_rsa", homeDir)
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
			if sshCommand != "" {
				lagoonssh.RunSSHCommand(sshConfig, sshService, sshContainer, sshCommand, config)
			} else {
				lagoonssh.InteractiveSSH(sshConfig, sshService, sshContainer, config)
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
