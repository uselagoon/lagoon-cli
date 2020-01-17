package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/amazeeio/lagoon-cli/api"
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
	handleError(err)
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
		//strip new line
		keyName = strings.TrimSuffix(splitKey[2], "\n")
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
		customReqResult, err = uClient.AddUser(userFlags)
		handleError(err)
		returnResultData := map[string]interface{}{}
		err = json.Unmarshal([]byte(customReqResult), &returnResultData)
		handleError(err)
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
		customReqResult, err = uClient.AddSSHKeyToUser(userFlags, userSSHKey)
		handleError(err)
		returnResultData := map[string]interface{}{}
		err = json.Unmarshal([]byte(customReqResult), &returnResultData)
		handleError(err)
		resultData := output.Result{
			Result:     "success",
			ResultData: returnResultData,
		}
		output.RenderResult(resultData, outputOptions)
	},
}

var delSSHKeyCmd = &cobra.Command{
	Use:     "user-sshkey",
	Aliases: []string{"u"},
	Short:   "Delete sshkey from lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		if sshKeyName == "" {
			fmt.Println("Not enough arguments. Requires: ssh key name")
			cmd.Help()
			os.Exit(1)
		}
		var customReqResult []byte
		var err error
		customReqResult, err = uClient.DeleteSSHKey(sshKeyName)
		handleError(err)
		resultData := output.Result{
			Result: string(customReqResult),
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
		customReqResult, err = uClient.DeleteUser(userFlags)
		handleError(err)
		resultData := output.Result{
			Result: string(customReqResult),
		}
		output.RenderResult(resultData, outputOptions)
	},
}

var updateUserCmd = &cobra.Command{
	Use:     "user",
	Aliases: []string{"u"},
	Short:   "Modify a user in lagoon (change name, or email address)",
	Run: func(cmd *cobra.Command, args []string) {
		userFlags := parseUser(*cmd.Flags())
		if userFlags.Email == "" {
			fmt.Println("Not enough arguments. Requires: email address")
			cmd.Help()
			os.Exit(1)
		}
		var customReqResult []byte
		var err error
		currentUser := api.User{
			Email: currentUserEmail,
		}
		customReqResult, err = uClient.ModifyUser(currentUser, userFlags)
		handleError(err)
		returnResultData := map[string]interface{}{}
		err = json.Unmarshal([]byte(customReqResult), &returnResultData)
		handleError(err)
		resultData := output.Result{
			Result:     "success",
			ResultData: returnResultData,
		}
		output.RenderResult(resultData, outputOptions)
	},
}

var getUserKeysCmd = &cobra.Command{
	//@TODO: once individual user interaction comes in, this will need to be adjusted
	Use:   "user-sshkeys",
	Short: "Get a users SSH keys",
	Long:  `Get a users SSH keys. This will only work for users that are part of a group`,
	Run: func(cmd *cobra.Command, args []string) {
		if userEmail == "" {
			fmt.Println("Not enough arguments. Requires: email address")
			cmd.Help()
			os.Exit(1)
		}
		returnedJSON, err := uClient.ListUserSSHKeys(groupName, userEmail, false)
		handleError(err)
		var dataMain output.Table
		err = json.Unmarshal([]byte(returnedJSON), &dataMain)
		handleError(err)
		if len(dataMain.Data) == 0 {
			output.RenderError(noDataError, outputOptions)
			os.Exit(1)
		}
		output.RenderOutput(dataMain, outputOptions)

	},
}

var getAllUserKeysCmd = &cobra.Command{
	//@TODO: once individual user interaction comes in, this will need to be adjusted
	Use:   "all-user-sshkeys",
	Short: "Get all user SSH keys",
	Long:  `Get all user SSH keys. This will only work for users that are part of a group`,
	Run: func(cmd *cobra.Command, args []string) {
		returnedJSON, err := uClient.ListUserSSHKeys(groupName, userEmail, true)
		handleError(err)
		var dataMain output.Table
		err = json.Unmarshal([]byte(returnedJSON), &dataMain)
		handleError(err)
		if len(dataMain.Data) == 0 {
			output.RenderError(noDataError, outputOptions)
			os.Exit(1)
		}
		output.RenderOutput(dataMain, outputOptions)

	},
}

var (
	currentUserEmail string
)

func init() {
	addUserCmd.Flags().StringVarP(&userFirstName, "firstName", "F", "", "Firstname of the user")
	addUserCmd.Flags().StringVarP(&userLastName, "lastName", "L", "", "Lastname of the user")
	addUserCmd.Flags().StringVarP(&userEmail, "email", "E", "", "Email address of the user")
	addUserSSHKeyCmd.Flags().StringVarP(&userEmail, "email", "E", "", "Email address of the user")
	addUserSSHKeyCmd.Flags().StringVarP(&sshKeyName, "keyname", "N", "", "Name of the sshkey (optional, if not provided will try use what is in the pubkey file)")
	addUserSSHKeyCmd.Flags().StringVarP(&pubKeyFile, "pubkey", "K", "", "file location to the public key to add")
	delUserCmd.Flags().StringVarP(&userEmail, "email", "E", "", "Email address of the user")
	delSSHKeyCmd.Flags().StringVarP(&sshKeyName, "keyname", "N", "", "Name of the sshkey")
	updateUserCmd.Flags().StringVarP(&userFirstName, "firstName", "F", "", "New firstname of the user")
	updateUserCmd.Flags().StringVarP(&userLastName, "lastName", "L", "", "New lastname of the user")
	updateUserCmd.Flags().StringVarP(&userEmail, "email", "E", "", "New email address of the user")
	updateUserCmd.Flags().StringVarP(&currentUserEmail, "current-email", "C", "", "Current email address of the user")
	getUserKeysCmd.Flags().StringVarP(&userEmail, "email", "E", "", "New email address of the user")
	getUserKeysCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group to check users in (if not specified, will default to all groups)")
	getAllUserKeysCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group to list users in (if not specified, will default to all groups)")
}
