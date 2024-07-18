package cmd

import (
	"bytes"
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

func publicKey(path, publicKeyOverride string, publicKeyIdentities []string, skipAgent bool) (ssh.AuthMethod, func() error) {
	noopCloseFunc := func() error { return nil }

	if !skipAgent {
		// Connect to SSH agent to ask for unencrypted private keys
		if sshAgentConn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
			sshAgent := agent.NewClient(sshAgentConn)
			agentSigners, err := sshAgent.Signers()
			handleError(err)
			// There are key(s) in the agent
			if len(agentSigners) > 0 {
				identities := make(map[string]ssh.PublicKey)
				if publicKeyOverride == "" {
					// check for identify files in the current lagoon config context
					for _, identityFile := range publicKeyIdentities {
						// append to identityfiles
						keybytes, err := os.ReadFile(identityFile)
						handleError(err)
						pubkey, _, _, _, err := ssh.ParseAuthorizedKey(keybytes)
						handleError(err)
						identities[identityFile] = pubkey
					}
				} else {
					// append to identityfiles
					keybytes, err := os.ReadFile(publicKeyOverride)
					handleError(err)
					pubkey, _, _, _, err := ssh.ParseAuthorizedKey(keybytes)
					handleError(err)
					identities[publicKeyOverride] = pubkey
				}
				// check all keys in the agent to see if there is a matching identity file
				for _, signer := range agentSigners {
					for file, identity := range identities {
						if bytes.Equal(signer.PublicKey().Marshal(), identity.Marshal()) {
							if verboseOutput {
								fmt.Fprintf(os.Stderr, "ssh: attempting connection using identity file public key: %s\n", file)
							}
							// only provide this matching key back to the ssh client to use
							return ssh.PublicKeys(signer), noopCloseFunc
						}
					}
				}
				if publicKeyOverride != "" {
					handleError(fmt.Errorf("ssh: no key matching %s in agent", publicKeyOverride))
				}
				// if no matching identity files, just return all agent keys like previous behaviour
				if verboseOutput {
					fmt.Fprintf(os.Stderr, "ssh: attempting connection using any keys in ssh-agent\n")
				}
				return ssh.PublicKeysCallback(sshAgent.Signers), noopCloseFunc
			}
		}
	}

	// if no keys in the agent, and a specific private key has been defined, then check the key and use it if possible
	if verboseOutput {
		fmt.Fprintf(os.Stderr, "ssh: attempting connection using private key: %s\n", path)
	}
	key, err := os.ReadFile(path)
	handleError(err)

	// Try to look for an unencrypted private key
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		// if encrypted, prompt for passphrase or error and ask user to add to their agent
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
	}
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
	authMethod, closeSSHAgent := publicKey(privateKey, cmdPubkeyIdentity, lUser.UserConfig.PublicKeyIdentities, skipAgent)
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
