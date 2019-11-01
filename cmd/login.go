package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log into a Lagoon instance",
	Run: func(cmd *cobra.Command, args []string) {
		loginErr := loginToken()
		if loginErr != nil {
			fmt.Println("Unable to login, error was: ", loginErr.Error())
			os.Exit(1)
		}
		fmt.Println("Token fetched and saved.")
	},
}

func publicKey(path string) ssh.AuthMethod {
	key, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		panic(err)
	}
	return ssh.PublicKeys(signer)
}

func loginToken() error {
	homeDir, _ := os.UserHomeDir()
	privateKey := fmt.Sprintf("%s/.ssh/id_rsa", homeDir)
	if cmdSSHKey != "" {
		privateKey = cmdSSHKey
		fmt.Println(privateKey)
	}
	config := &ssh.ClientConfig{
		User: "lagoon",
		Auth: []ssh.AuthMethod{
			publicKey(privateKey),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	var err error

	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", viper.GetString("lagoons."+cmdLagoon+".hostname"), viper.GetString("lagoons."+cmdLagoon+".port")), config)
	if err != nil {
		//panic(err)
		return err
	}
	session, err := conn.NewSession()
	if err != nil {
		_ = conn.Close()
		//panic(err)
		return err
	}

	out, err := session.CombinedOutput("token")
	if err != nil {
		//panic(err)
		return err
	}
	err = conn.Close()
	viper.Set("lagoons."+cmdLagoon+".token", strings.TrimSpace(string(out)))
	err = viper.WriteConfig()
	if err != nil {
		//panic(err)
		return err
	}
	return nil
}
