package ssh

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

// InteractiveSSH .
func InteractiveSSH(lagoon map[string]string, sshService string, sshContainer string, privKey string) {
	pk, _ := ioutil.ReadFile(privKey)
	signer, err := ssh.ParsePrivateKey(pk)
	if err != nil {
		panic(err)
	}
	// ignore insecure hostkey, changes in lagoon
	config := &ssh.ClientConfig{
		User:            lagoon["username"],
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
	}
	client, err := ssh.Dial("tcp", lagoon["hostname"]+":"+lagoon["port"], config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}

	// start the session
	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // enable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	fileDescriptor := int(os.Stdin.Fd())
	if terminal.IsTerminal(fileDescriptor) {
		originalState, err := terminal.MakeRaw(fileDescriptor)
		if err != nil {
			log.Fatalf("%s", err)
		}
		defer terminal.Restore(fileDescriptor, originalState)
		termWidth, termHeight, err := terminal.GetSize(fileDescriptor)
		if err != nil {
			log.Fatalf("%s", err)
		}
		err = session.RequestPty("xterm-256color", termHeight, termWidth, modes)
		if err != nil {
			log.Fatalf("%s", err)
		}
	}
	var connString string
	if sshService != "" {
		connString = fmt.Sprintf("%s service=%s", connString, sshService)
	}
	if sshContainer != "" && sshService != "" {
		connString = fmt.Sprintf("%s container=%s", connString, sshContainer)
	}
	err = session.Start(connString)
	if err != nil {
		log.Fatalf("failed to start shell: %s", err)
	}
	session.Wait()

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
