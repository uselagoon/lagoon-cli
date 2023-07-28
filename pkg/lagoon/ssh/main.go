// Package ssh implements an SSH client for Lagoon.
package ssh

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

// LogStream connects to host:port using the given config, and executes the
// argv command. It does not request a PTY, and instead just streams the
// response to the attached terminal. argv should contain a logs=... argument.
func LogStream(config *ssh.ClientConfig, host, port string, argv []string) error {
	// https://stackoverflow.com/a/37088088
	client, err := ssh.Dial("tcp", host+":"+port, config)
	if err != nil {
		return fmt.Errorf("couldn't dial SSH: %v", err)
	}
	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("couldn't create SSH session: %v", err)
	}
	defer session.Close()
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin
	err = session.Start(strings.Join(argv, " "))
	if err != nil {
		return fmt.Errorf("couldn't start SSH session: %v", err)
	}
	return session.Wait()
}

// InteractiveSSH .
func InteractiveSSH(lagoon map[string]string, sshService string, sshContainer string, config *ssh.ClientConfig) error {
	client, err := ssh.Dial("tcp", lagoon["hostname"]+":"+lagoon["port"], config)
	if err != nil {
		return fmt.Errorf("Failed to dial: " + err.Error() + "\nCheck that the project or environment you are trying to connect to exists")
	}

	// start the session
	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("Failed to create session: " + err.Error())
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
	if term.IsTerminal(fileDescriptor) {
		originalState, err := term.MakeRaw(fileDescriptor)
		if err != nil {
			return err
		}
		defer term.Restore(fileDescriptor, originalState)
		termWidth, termHeight, err := term.GetSize(fileDescriptor)
		if err != nil {
			return err
		}
		err = session.RequestPty("xterm-256color", termHeight, termWidth, modes)
		if err != nil {
			return err
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
		return fmt.Errorf("Failed to start shell: " + err.Error())
	}
	session.Wait()
	return nil
}

// RunSSHCommand .
func RunSSHCommand(lagoon map[string]string, sshService string, sshContainer string, command string, config *ssh.ClientConfig) error {
	client, err := ssh.Dial("tcp", lagoon["hostname"]+":"+lagoon["port"], config)
	if err != nil {
		return fmt.Errorf("Failed to dial: " + err.Error() + "\nCheck that the project or environment you are trying to connect to exists")
	}

	// start the session
	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("Failed to create session: " + err.Error())
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
	if term.IsTerminal(fileDescriptor) {
		originalState, err := term.MakeRaw(fileDescriptor)
		if err != nil {
			return err
		}
		defer term.Restore(fileDescriptor, originalState)
		termWidth, termHeight, err := term.GetSize(fileDescriptor)
		if err != nil {
			return err
		}
		err = session.RequestPty("xterm-256color", termHeight, termWidth, modes)
		if err != nil {
			return err
		}
	}
	var connString string
	if sshService != "" {
		connString = fmt.Sprintf("%s service=%s", connString, sshService)
	}
	if sshContainer != "" && sshService != "" {
		connString = fmt.Sprintf("%s container=%s", connString, sshContainer)
	}
	var b bytes.Buffer
	session.Stdout = &b

	err = session.Run(connString + " " + command)
	if err != nil {
		return err
	}
	fmt.Println(b.String())
	return nil
}
