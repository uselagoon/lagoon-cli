package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/internal/lagoon"
	"github.com/uselagoon/lagoon-cli/internal/lagoon/client"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		validateToken(lagoonCLIConfig.Current) // get a new token if the current one is invalid

		if cmdProjectName == "" || cmdProjectEnvironment == "" {
			return fmt.Errorf("Missing arguments: Project name or environment name are not defined")
		}
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}

		// allow the use of the `feature/branch` and standard `feature-branch` type environment names to be used
		// since ssh requires the `feature-branch` type name to be used as the ssh username
		// run the environment through the makesafe and shorted functions that lagoon uses
		environmentName := makeSafe(shortenEnvironment(cmdProjectName, cmdProjectEnvironment))

		current := lagoonCLIConfig.Current
		// set the default ssh host and port to the core ssh endpoint
		sshHost := lagoonCLIConfig.Lagoons[current].HostName
		sshPort := lagoonCLIConfig.Lagoons[current].Port
		isPortal := false

		// if the config for this lagoon is set to use ssh portal support, handle that here
		lc := client.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIConfig.Lagoons[current].Token,
			lagoonCLIConfig.Lagoons[current].Version,
			lagoonCLIVersion,
			debug)
		project, err := lagoon.GetSSHEndpointsByProject(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		// check all the environments for this project
		for _, env := range project.Environments {
			// if the env name matches the requested environment then check if the deploytarget supports regional ssh endpoints
			if env.Name == environmentName {
				// if the deploytarget supports regional endpoints, then set these as the host and port for ssh
				if env.DeployTarget.SSHHost != "" && env.DeployTarget.SSHPort != "" {
					sshHost = env.DeployTarget.SSHHost
					sshPort = env.DeployTarget.SSHPort
					isPortal = true
				}
			}
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
			"username": cmdProjectName + "-" + environmentName,
			"sshkey":   privateKey,
		}
		if sshConnString {
			fmt.Println(generateSSHConnectionString(sshConfig, sshService, sshContainer, isPortal))
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
func generateSSHConnectionString(lagoon map[string]string, service string, container string, isPortal bool) string {
	connString := fmt.Sprintf("ssh -t")
	if lagoon["sshKey"] != "" {
		connString = fmt.Sprintf("%s -i %s", connString, lagoon["sshKey"])
	}
	if !isPortal {
		connString = fmt.Sprintf("%s -o \"UserKnownHostsFile=/dev/null\" -o \"StrictHostKeyChecking=no\"", connString)
	}
	connString = fmt.Sprintf("%s -p %v %s@%s", connString, lagoon["port"], lagoon["username"], lagoon["hostname"])
	if service != "" {
		connString = fmt.Sprintf("%s service=%s", connString, service)
	}
	if container != "" && service != "" {
		connString = fmt.Sprintf("%s container=%s", connString, container)
	}
	return connString
}
