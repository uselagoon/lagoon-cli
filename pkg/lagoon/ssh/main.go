// Package ssh implements an SSH client for Lagoon.
package ssh

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"path"
	"strings"

	"github.com/skeema/knownhosts"
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
		return fmt.Errorf("couldn't dial SSH (maybe this service doesn't support logs?): %v", err)
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
		return fmt.Errorf("failed to dial: %s\nCheck that the project or environment you are trying to connect to exists", err.Error())
	}

	// start the session
	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create session: %s", err.Error())
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
		defer func() {
			err = term.Restore(fileDescriptor, originalState)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error restoring ssh terminal:%v\n", err)
			}
		}()
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
		connString = fmt.Sprintf("service=%s", sshService)
	}
	if sshContainer != "" && sshService != "" {
		connString = fmt.Sprintf("%s container=%s", connString, sshContainer)
	}
	err = session.Start(connString)
	if err != nil {
		return fmt.Errorf("failed to start shell: %s", err.Error())
	}
	return session.Wait()
}

// RunSSHCommand .
func RunSSHCommand(lagoon map[string]string, sshService string, sshContainer string, command string, config *ssh.ClientConfig) error {
	client, err := ssh.Dial("tcp", lagoon["hostname"]+":"+lagoon["port"], config)
	if err != nil {
		return fmt.Errorf("failed to dial: %s\nCheck that the project or environment you are trying to connect to exists", err.Error())
	}

	// start the session
	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create session: %s", err.Error())
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
		defer func() {
			err = term.Restore(fileDescriptor, originalState)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error restoring ssh terminal:%v\n", err)
			}
		}()
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

	// if there's anything waiting in the buffer, display it (regardless of error state or not)
	// it may, in error state, contain useful information.
	if b.Len() > 0 {
		fmt.Println(b.String())
	}

	if err != nil {
		return err
	}
	return nil
}

// make the
var hostKeyWarn = `The authenticity of host '%s' can't be established.
%s key fingerprint is %s.
Are you sure you want to continue connecting %s`

var remoteHostChanged = `@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@    WARNING: REMOTE HOST IDENTIFICATION HAS CHANGED!     @
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
IT IS POSSIBLE THAT SOMEONE IS DOING SOMETHING NASTY!
Someone could be eavesdropping on you right now (man-in-the-middle attack)!
It is also possible that a host key has just been changed.
The fingerprint for the %s key sent by the remote host is
%s
Add correct host key in %s to get rid of this message`

// add interactive known hosts to reduce confusion with host key errors
func InteractiveKnownHosts(userPath, host string, ignorehost, accept bool) (ssh.HostKeyCallback, []string, error) {
	if ignorehost {
		// if ignore provided, just skip the hostkey verifications
		return ssh.InsecureIgnoreHostKey(), nil, nil
	}
	kh, err := knownhosts.NewDB(path.Join(userPath, ".ssh/known_hosts"))
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't get ~/.ssh/known_hosts: %v", err)
	}
	// otherwise prompt or accept for the key if required
	return ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		filePath := path.Join(userPath, ".ssh/known_hosts")
		innerCallback := kh.HostKeyCallback()
		err := innerCallback(hostname, remote, key)
		sshPubKey := ssh.MarshalAuthorizedKey(key)
		pub, _, _, _, perr := ssh.ParseAuthorizedKey(sshPubKey)
		if perr != nil {
			return fmt.Errorf("knownhosts: host key verification failed")
		}
		if knownhosts.IsHostKeyChanged(err) {
			os.Stderr.WriteString(fmt.Sprintf(remoteHostChanged, key.Type(), ssh.FingerprintSHA256(pub), filePath))
			return fmt.Errorf("knownhosts: host key verification failed")
		} else if knownhosts.IsHostUnknown(err) {
			f, ferr := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0600)
			if ferr == nil {
				defer f.Close()
				var response string
				if accept {
					response = "y"
				} else {
					os.Stderr.WriteString(fmt.Sprintf(hostKeyWarn, hostname, key.Type(), ssh.FingerprintSHA256(pub), "(yes/no)? "))
					reader := bufio.NewReader(os.Stdin)
					response, err = reader.ReadString('\n')
					if err != nil {
						return fmt.Errorf("knownhosts: host key verification failed %v", err)
					}
					response = strings.ToLower(strings.TrimSpace(response))
				}
				if response == "Yes" || response == "yes" || response == "y" {
					ferr = knownhosts.WriteKnownHost(f, hostname, remote, key)
				} else {
					return fmt.Errorf("knownhosts: host key verification failed")
				}
			}
			if ferr == nil {
				os.Stderr.WriteString(fmt.Sprintf("Warning: Permanently added '%s' to the list of known hosts\n", hostname))
			} else {
				os.Stderr.WriteString(fmt.Sprintf("Failed to add host %s to known_hosts: %v\n", hostname, ferr))
			}
			return nil // permit previously-unknown hosts (warning: may be insecure)
		}
		return err
	}), kh.HostKeyAlgorithms(host), nil
}

func CheckStrictHostKey(v string) (bool, bool) {
	if v == "ignore" || v == "no" {
		// return ignore true, accept false
		return true, false
	}
	if v == "accept-new" {
		// return ignore false, accept true
		return false, true
	}
	// default "accept-new"
	return false, false
}
