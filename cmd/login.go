package cmd

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"golang.org/x/crypto/ssh/terminal"
)

var loginCmd = &cobra.Command{
	Use:     "login",
	Short:   "Log into a Lagoon instance",
	Aliases: []string{"l"},
	Run: func(cmd *cobra.Command, args []string) {
		validateToken(lagoonCLIConfig.Current) // get a new token if the current one is invalid
		fmt.Println("Token fetched and saved.")
	},
}

func publicKey(path string, skipAgent bool) (ssh.AuthMethod, func() error) {
	noopCloseFunc := func() error { return nil }

	if skipAgent != true {
		// Connect to SSH agent to ask for unencrypted private keys
		if sshAgentConn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
			sshAgent := agent.NewClient(sshAgentConn)

			keys, _ := sshAgent.List()
			if len(keys) > 0 {
				// There are key(s) in the agent
				//defer sshAgentConn.Close()
				return ssh.PublicKeysCallback(sshAgent.Signers), sshAgentConn.Close
			}
		}
	}

	key, err := os.ReadFile(path)
	handleError(err)

	// Try to look for an unencrypted private key
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		fmt.Println("Error was:", err.Error())
		fmt.Println("Lagoon CLI cannot decode private keys, you will need to add your private key to your ssh-agent.")
		os.Exit(1)
	} else if err == nil {
		// return unencrypted private key
		return ssh.PublicKeys(signer), noopCloseFunc
	}

	// Handle encrypted private keys
	fmt.Println("Found an encrypted private key!")
	fmt.Printf("Enter passphrase for '%s': ", path)
	bytePassword, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	signer, err = ssh.ParsePrivateKeyWithPassphrase(key, bytePassword)
	return ssh.PublicKeys(signer), noopCloseFunc
}

func loginToken() error {
	out, err := retrieveTokenViaSsh()
	if err != nil {
		return err
	}

	lc := lagoonCLIConfig.Lagoons[lagoonCLIConfig.Current]
	lc.Token = out
	lagoonCLIConfig.Lagoons[lagoonCLIConfig.Current] = lc
	if err = writeLagoonConfig(&lagoonCLIConfig, filepath.Join(configFilePath, configName+configExtension)); err != nil {
		return fmt.Errorf("couldn't write config: %v", err)
	}

	return nil
}

func retrieveTokenViaSsh() (string, error) {
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
	authMethod, closeSSHAgent := publicKey(privateKey, skipAgent)
	config := &ssh.ClientConfig{
		User: "lagoon",
		Auth: []ssh.AuthMethod{
			authMethod,
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	defer closeSSHAgent()

	sshHost := fmt.Sprintf("%s:%s",
		lagoonCLIConfig.Lagoons[lagoonCLIConfig.Current].HostName,
		lagoonCLIConfig.Lagoons[lagoonCLIConfig.Current].Port)
	conn, err := ssh.Dial("tcp", sshHost, config)
	if err != nil {
		return "", fmt.Errorf("couldn't connect to %s: %v", sshHost, err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		return "", fmt.Errorf("couldn't open session: %v", err)
	}

	out, err := session.CombinedOutput("token")
	if err != nil {
		return "", fmt.Errorf("couldn't get token: %v", err)
	}
	return strings.TrimSpace(string(out)), err
}
