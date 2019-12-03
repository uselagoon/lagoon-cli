package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/lagoon/users"
	"github.com/amazeeio/lagoon-cli/output"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func parseUser(flags pflag.FlagSet) api.User {
	configMap := make(map[string]interface{})
	flags.VisitAll(func(f *pflag.Flag) {
		if flags.Changed(f.Name) {
			configMap[f.Name] = f.Value
		}
	})
	jsonStr, _ := json.Marshal(configMap)
	parsedFlags := api.User{}
	json.Unmarshal(jsonStr, &parsedFlags)
	return parsedFlags
}

func parseSSHKeyFile(sshPubKey string, keyName string) api.SSHKey {
	b, err := ioutil.ReadFile(sshPubKey) // just pass the file name
	if err != nil {
		output.RenderError(err.Error(), outputOptions)
		os.Exit(1)
	}
	splitKey := strings.Split(string(b), " ")
	var keyType api.SSHKeyType
	// default to ssh-rsa, otherwise check if ssh-ed25519
	// will fail if neither are right
	keyType = api.SSHRsa
	if strings.EqualFold(string(splitKey[0]), "ssh-ed25519") {
		keyType = api.SSHEd25519
	}
	// if the sshkey has a comment/name in it, we can use that
	if keyName == "" && len(splitKey) == 3 {
		keyName = splitKey[2]
	} else if keyName == "" && len(splitKey) == 2 {
		output.RenderError("no name provided", outputOptions)
		os.Exit(1)
	}
	parsedFlags := api.SSHKey{
		KeyType:  keyType,
		KeyValue: splitKey[1],
		Name:     keyName,
	}
	return parsedFlags
}

var addUserCmd = &cobra.Command{
	Use:     "user",
	Aliases: []string{"u"},
	Short:   "Add user to lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		userFlags := parseUser(*cmd.Flags())
		if userFlags.Email == "" {
			fmt.Println("Not enough arguments. Requires: email address")
			cmd.Help()
			os.Exit(1)
		}
		var customReqResult []byte
		var err error
		customReqResult, err = users.AddUser(userFlags)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		returnResultData := map[string]interface{}{}
		err = json.Unmarshal([]byte(customReqResult), &returnResultData)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		resultData := output.Result{
			Result:     "success",
			ResultData: returnResultData,
		}
		output.RenderResult(resultData, outputOptions)
	},
}

var addUserSSHKeyCmd = &cobra.Command{
	Use:     "user-sshkey",
	Aliases: []string{"uk"},
	Short:   "Add sshkey to a user",
	Run: func(cmd *cobra.Command, args []string) {
		userFlags := parseUser(*cmd.Flags())
		if userFlags.Email == "" {
			fmt.Println("Not enough arguments. Requires: email address")
			cmd.Help()
			os.Exit(1)
		}
		userSSHKey := parseSSHKeyFile(pubKeyFile, sshKeyName)
		var customReqResult []byte
		var err error
		customReqResult, err = users.AddSSHKeyToUser(userFlags, userSSHKey)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		returnResultData := map[string]interface{}{}
		err = json.Unmarshal([]byte(customReqResult), &returnResultData)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		resultData := output.Result{
			Result:     "success",
			ResultData: returnResultData,
		}
		output.RenderResult(resultData, outputOptions)
	},
}

var delUserCmd = &cobra.Command{
	Use:     "user",
	Aliases: []string{"u"},
	Short:   "Delete user from lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		userFlags := parseUser(*cmd.Flags())
		if userFlags.Email == "" {
			fmt.Println("Not enough arguments. Requires: email address")
			cmd.Help()
			os.Exit(1)
		}
		var customReqResult []byte
		var err error
		customReqResult, err = users.DeleteUser(userFlags)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		resultData := output.Result{
			Result: string(customReqResult),
		}
		output.RenderResult(resultData, outputOptions)
	},
}

func init() {
	addUserCmd.Flags().StringVarP(&userFirstName, "firstName", "F", "", "Firstname of the user")
	addUserCmd.Flags().StringVarP(&userLastName, "lastName", "L", "", "Lastname of the user")
	addUserCmd.Flags().StringVarP(&userEmail, "email", "E", "", "Email address of the user")
	addUserSSHKeyCmd.Flags().StringVarP(&userEmail, "email", "E", "", "Email address of the user")
	addUserSSHKeyCmd.Flags().StringVarP(&sshKeyName, "keyname", "N", "", "Name of the sshkey (optional, if not provided will try use what is in the pubkey file)")
	addUserSSHKeyCmd.Flags().StringVarP(&pubKeyFile, "pubkey", "K", "", "file location to the public key to add")
	delUserCmd.Flags().StringVarP(&userEmail, "email", "E", "", "Email address of the user")
}
