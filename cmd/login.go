package cmd

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
	auth "github.com/uselagoon/machinery/utils/auth"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"golang.org/x/oauth2"
	terminal "golang.org/x/term"
)

var loginCmd = &cobra.Command{
	Use:     "login",
	Short:   "Log into a Lagoon instance",
	Aliases: []string{"l"},
	Run: func(cmd *cobra.Command, args []string) {
		validateToken(lContext.Name) // get a new token if the current one is invalid
		fmt.Println("Token fetched and saved.")
	},
}

func publicKey(path string, skipAgent bool) (ssh.AuthMethod, func() error) {
	noopCloseFunc := func() error { return nil }

	if !skipAgent {
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
		fmt.Printf("Enter passphrase for %s:", path)
		bytePassword, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Error was:", err.Error())
			fmt.Println("Lagoon CLI could not decode private key, you will need to add your private key to your ssh-agent.")
			os.Exit(1)
		}
		fmt.Println()
		signer, err = ssh.ParsePrivateKeyWithPassphrase(key, bytePassword)
		if err != nil {
			fmt.Println("Error was:", err.Error())
			fmt.Println("Lagoon CLI could not decode private key, you will need to add your private key to your ssh-agent.")
			os.Exit(1)
		}
		return ssh.PublicKeys(signer), noopCloseFunc
	}
	// return unencrypted private key
	return ssh.PublicKeys(signer), noopCloseFunc
}

func loginToken() error {
	// check if the ssh-token only feature is enabled for this context for cli generally
	sshTokenOnly, err := lConfig.GetFeature(lContext.Name, configFeaturePrefix, "ssh-token")
	if err != nil {
		return err
	}
	if lContext.ContextConfig.AuthenticationEndpoint == "" || sshTokenOnly {
		// if no keycloak url is found in the config, perform a token request via ssh
		// or the ssh-token override is set to enforce tokens via ssh (accounts in CI jobs)
		out, err := retrieveTokenViaSsh()
		if err != nil {
			return err
		}
		lUser.UserConfig.Grant = out
	} else {
		// otherwise get a token via keycloak
		token := &oauth2.Token{}
		if lUser.UserConfig.Grant != nil {
			token = lUser.UserConfig.Grant
		}
		_ = auth.TokenRequest(lContext.ContextConfig.AuthenticationEndpoint, "lagoon", "", token)
		lUser.UserConfig.Grant = token
	}
	return lConfig.WriteConfig()
}

func retrieveTokenViaSsh() (*oauth2.Token, error) {
	skipAgent := false
	privateKey := fmt.Sprintf("%s/.ssh/id_rsa", userPath)
	// if the user has a key defined in their lagoon cli config, use it
	if lUser.UserConfig.SSHKey != "" {
		privateKey = lUser.UserConfig.SSHKey
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

	sshHost := fmt.Sprintf("%s:%d",
		lContext.ContextConfig.TokenHost,
		lContext.ContextConfig.TokenPort)
	conn, err := ssh.Dial("tcp", sshHost, config)
	if err != nil {
		return nil, fmt.Errorf("unable to authenticate or connect to host %s\nthere may be an issue determining which ssh-key to use, or there may be an issue establishing a connection to the host\nthe error returned was: %v", sshHost, err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		return nil, fmt.Errorf("unable to establish ssh session, error from attempt is: %v", err)
	}

	out, err := session.CombinedOutput("grant")
	if err != nil {
		return nil, fmt.Errorf("unable to get token: %v", err)
	}
	token := &oauth2.Token{}
	json.Unmarshal(out, token)
	return token, err
}
