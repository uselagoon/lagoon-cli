package cmd

import (
	"context"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/spf13/cobra"
	lagoonssh "github.com/uselagoon/lagoon-cli/pkg/lagoon/ssh"
	"github.com/uselagoon/lagoon-cli/pkg/output"
	"github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

var (
	// connTimeout is the network connection timeout used for SSH connections and
	// calls to the Lagoon API.
	connTimeout = 8 * time.Second
	// these variables are assigned in init() to flag values
	logsService   string
	logsContainer string
	logsTailLines uint
	logsFollow    bool
)

func init() {
	logsCmd.Flags().StringVarP(&logsService, "service", "s", "", "specify a specific service name")
	logsCmd.Flags().StringVarP(&logsContainer, "container", "c", "", "specify a specific container name")
	logsCmd.Flags().UintVarP(&logsTailLines, "lines", "n", 32, "the number of lines to return for each container")
	logsCmd.Flags().BoolVarP(&logsFollow, "follow", "f", false, "continue outputting new lines as they are logged")
}

func generateLogsCommand(service, container string, lines uint,
	follow bool) ([]string, error) {
	var argv []string
	if service == "" {
		return nil, fmt.Errorf("empty service name")
	}
	if unsafeRegex.MatchString(service) {
		return nil, fmt.Errorf("service name contains invalid characters")
	}
	argv = append(argv, "service="+service)
	if container != "" {
		if unsafeRegex.MatchString(container) {
			return nil, fmt.Errorf("container name contains invalid characters")
		}
		argv = append(argv, "container="+container)
	}
	logsCmd := fmt.Sprintf("logs=tailLines=%d", lines)
	if follow {
		logsCmd += ",follow"
	}
	argv = append(argv, logsCmd)
	return argv, nil
}

func getSSHHostPort(environmentName string, debug bool) (string, string, error) {
	sshHost := lContext.ContextConfig.TokenHost
	sshPort := fmt.Sprintf("%d", lContext.ContextConfig.TokenPort)

	// if the config for this lagoon is set to use ssh portal support, handle that here
	utoken := lUser.UserConfig.Grant.AccessToken
	lc := lclient.New(
		fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
		lagoonCLIVersion,
		lContext.ContextConfig.Version,
		&utoken,
		debug)
	ctx, cancel := context.WithTimeout(context.Background(), connTimeout)
	defer cancel()
	project, err := lagoon.GetSSHEndpointsByProject(ctx, cmdProjectName, lc)
	if err != nil {
		return "", "", fmt.Errorf("couldn't get SSH endpoint by project: %v", err)
	}
	// check all the environments for this project
	for _, env := range project.Environments {
		// if the env name matches the requested environment then check if the deploytarget supports regional ssh endpoints
		if env.Name == environmentName {
			// if the deploytarget supports regional endpoints, then set these as the host and port for ssh
			if env.DeployTarget.SSHHost != "" && env.DeployTarget.SSHPort != "" {
				sshHost = env.DeployTarget.SSHHost
				sshPort = env.DeployTarget.SSHPort
			}
		}
	}
	return sshHost, sshPort, nil
}

func getSSHClientConfig(environmentName string) (*ssh.ClientConfig,
	func() error, error) {
	skipAgent := false
	privateKey := fmt.Sprintf("%s/.ssh/id_rsa", userPath)
	// check for user-defined key
	if lUser.UserConfig.SSHKey != "" {
		privateKey = lUser.UserConfig.SSHKey
		skipAgent = true
	}
	// check for specified key
	if cmdSSHKey != "" {
		privateKey = cmdSSHKey
		skipAgent = true
	}
	// parse known_hosts
	kh, err := knownhosts.New(path.Join(userPath, ".ssh/known_hosts"))
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't get ~/.ssh/known_hosts: %v", err)
	}
	// configure an SSH client session
	authMethod, closeSSHAgent := publicKey(privateKey, cmdPubkeyIdentity, lUser.UserConfig.PublicKeyIdentities, skipAgent)
	return &ssh.ClientConfig{
		User:            cmdProjectName + "-" + environmentName,
		Auth:            []ssh.AuthMethod{authMethod},
		HostKeyCallback: kh,
		Timeout:         connTimeout,
	}, closeSSHAgent, nil
}

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Display logs for a service of an environment and project",
	RunE: func(cmd *cobra.Command, args []string) error {
		// validate/refresh token
		validateToken(lContext.Name)
		// validate and parse arguments
		if cmdProjectName == "" || cmdProjectEnvironment == "" {
			return fmt.Errorf(
				"missing arguments: Project name or environment name are not defined")
		}
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return fmt.Errorf("couldn't get debug value: %v", err)
		}
		argv, err := generateLogsCommand(logsService, logsContainer, logsTailLines,
			logsFollow)
		if err != nil {
			return fmt.Errorf("couldn't generate logs command: %v", err)
		}
		// replace characters in environment name to allow flexible referencing
		environmentName := makeSafe(
			shortenEnvironment(cmdProjectName, cmdProjectEnvironment))
		// query the Lagoon API for the environment's SSH endpoint
		sshHost, sshPort, err := getSSHHostPort(environmentName, debug)
		if err != nil {
			return fmt.Errorf("couldn't get SSH endpoint: %v", err)
		}
		// configure SSH client session
		sshConfig, closeSSHAgent, err := getSSHClientConfig(environmentName)
		if err != nil {
			return fmt.Errorf("couldn't get SSH client config: %v", err)
		}
		defer closeSSHAgent()
		// start SSH log streaming session
		err = lagoonssh.LogStream(sshConfig, sshHost, sshPort, argv)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			switch e := err.(type) {
			case *ssh.ExitMissingError:
				// https://github.com/openssh/openssh-portable/blob/
				// 	6958f00acf3b9e0b3730f7287e69996bcf3ceda4/fatal.c#L45
				os.Exit(255)
			case *ssh.ExitError:
				os.Exit(e.ExitStatus())
			default:
				os.Exit(254) // internal error
			}
		}
		return nil
	},
}
