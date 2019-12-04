package cmd

import (
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/lagoon/ssh"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
			fmt.Println(ssh.GenerateSSHConnectionString(sshConfig, sshService, sshContainer))
		} else {
			// get private key that the cli is using
			homeDir, _ := os.UserHomeDir()
			privateKey := fmt.Sprintf("%s/.ssh/id_rsa", homeDir)
			if cmdSSHKey != "" {
				privateKey = cmdSSHKey
			}
			// start an interactive ssh session
			ssh.InteractiveSSH(sshConfig, sshService, sshContainer, privateKey)
		}

	},
}

func init() {
	sshEnvCmd.Flags().StringVarP(&sshService, "service", "s", "", "specify a specific service name")
	sshEnvCmd.Flags().StringVarP(&sshContainer, "container", "c", "", "specify a specific container name")
	sshEnvCmd.Flags().BoolVarP(&sshConnString, "conn-string", "", false, "Display the full ssh connection string")
}
