package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"os"
	"strings"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log into a Lagoon instance",
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, _ := os.UserHomeDir()
		config := &ssh.ClientConfig{
			User: "lagoon",
			Auth: []ssh.AuthMethod{
				publicKey(fmt.Sprintf("%s/.ssh/id_rsa", homeDir)),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
		var err error

		conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", viper.GetString("lagoon_hostname"), viper.GetString("lagoon_port")), config)
		if err != nil {
			panic(err)
		}
		session, err := conn.NewSession()
		if err != nil {
			_ = conn.Close()
			panic(err)
		}

		out, err := session.CombinedOutput("token")
		if err != nil {
			panic(err)
		}
		err = conn.Close()
		viper.Set("lagoon_token", strings.TrimSpace(string(out)))
		err = viper.WriteConfig()
		if err != nil {
			panic(err)
		}
		fmt.Println("Token fetched and saved.")
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
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
