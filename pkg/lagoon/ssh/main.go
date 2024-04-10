package ssh

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

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
