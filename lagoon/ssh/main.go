package ssh

import (
	"fmt"

	"github.com/helloyi/go-sshclient"
	"golang.org/x/crypto/ssh"
)

// InteractiveSSH .
func InteractiveSSH(lagoon map[string]string, sshService string, sshContainer string, privKey string) {
	client, err := sshclient.DialWithKey(lagoon["hostname"]+":"+lagoon["port"], lagoon["username"], privKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	// with a terminal config
	config := &sshclient.TerminalConfig{
		Term:   "xterm",
		Height: 40,
		Weight: 80,
		Modes: ssh.TerminalModes{
			ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
			ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
		},
	}
	if err := client.Terminal(config).Start(); err != nil {
		fmt.Println(err)
		return
	}
}

// GenerateSSHConnectionString .
func GenerateSSHConnectionString(lagoon map[string]string, service string, container string) string {
	connString := fmt.Sprintf("ssh -t -o \"UserKnownHostsFile=/dev/null\" -o \"StrictHostKeyChecking=no\" -p %v %s@%s", lagoon["port"], lagoon["username"], lagoon["hostname"])
	if service != "" {
		connString = fmt.Sprintf("%s service=%s", connString, service)
	}
	if container != "" && service != "" {
		connString = fmt.Sprintf("%s container=%s", connString, container)
	}
	return connString
}
