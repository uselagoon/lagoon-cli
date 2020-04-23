package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/amazeeio/lagoon-cli/pkg/api"
	"github.com/amazeeio/lagoon-cli/pkg/output"
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

func parseSSHKeyFile(sshPubKey string, keyName string, keyValue string, userEmail string) api.SSHKey {
	// if we haven't got a keyvalue
	if keyValue == "" {
		b, err := ioutil.ReadFile(sshPubKey) // just pass the file name
		handleError(err)
		keyValue = string(b)
	}
	splitKey := strings.Split(keyValue, " ")
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
		keyName = stripNewLines(splitKey[2])
	} else if keyName == "" && len(splitKey) == 2 {
		keyName = userEmail
		output.RenderInfo("no name provided, using email address as key name", outputOptions)
	}
	parsedFlags := api.SSHKey{
		KeyType:  keyType,
		KeyValue: stripNewLines(splitKey[1]),
		Name:     keyName,
	}
	return parsedFlags
}

var addUserCmd = &cobra.Command{
	Use:     "user",
	Aliases: []string{"u"},
	Short:   "Add a user to lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		userFlags := parseUser(*cmd.Flags())
		if userFlags.Email == "" {
			fmt.Println("Missing arguments: Email address is not defined")
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
	Short:   "Add an sshkey to a user",
	Long: `Add an sshkey to a user

Examples:
Add key from public key file:
  lagoon add user-sshkey --email test@example.com --pubkey /path/to/id_rsa.pub

Add key by defining full key value:
  lagoon add user-sshkey --email test@example.com --keyvalue "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAINA0ITV2gbDc6noYeWaqfxTYpaEKq7HzU3+F71XGhSL/ my-computer@example"

Add key by defining full key value, but a specific key name:
  lagoon add user-sshkey --email test@example.com --keyname my-computer@example --keyvalue "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAINA0ITV2gbDc6noYeWaqfxTYpaEKq7HzU3+F71XGhSL/"

Add key by defining key value, but not specifying a key name (will default to try using the email address as key name):
  lagoon add user-sshkey --email test@example.com --keyvalue "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAINA0ITV2gbDc6noYeWaqfxTYpaEKq7HzU3+F71XGhSL/"

`,
	Run: func(cmd *cobra.Command, args []string) {
		userFlags := parseUser(*cmd.Flags())
		if userFlags.Email == "" {
			fmt.Println("Missing arguments: Email address is not defined")
			cmd.Help()
			os.Exit(1)
		}
		userSSHKey := parseSSHKeyFile(pubKeyFile, sshKeyName, pubKeyValue, userFlags.Email)
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

var deleteSSHKeyCmd = &cobra.Command{
	Use:     "user-sshkey",
	Aliases: []string{"u"},
	Short:   "Delete an sshkey from lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		if sshKeyName == "" {
			fmt.Println("Missing arguments: SSH key name is not defined")
			cmd.Help()
			os.Exit(1)
		}
		var customReqResult []byte
		var err error
		if yesNo(fmt.Sprintf("You are attempting to delete ssh key named '%s', are you sure?", sshKeyName)) {
			customReqResult, err = uClient.DeleteSSHKey(sshKeyName)
			handleError(err)
			resultData := output.Result{
				Result: string(customReqResult),
			}
			output.RenderResult(resultData, outputOptions)
		}
	},
}

var deleteUserCmd = &cobra.Command{
	Use:     "user",
	Aliases: []string{"u"},
	Short:   "Delete a user from lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		userFlags := parseUser(*cmd.Flags())
		if userFlags.Email == "" {
			fmt.Println("Missing arguments: Email address is not defined")
			cmd.Help()
			os.Exit(1)
		}
		var customReqResult []byte
		var err error
		if yesNo(fmt.Sprintf("You are attempting to delete user with email address '%s', are you sure?", userFlags.Email)) {
			customReqResult, err = uClient.DeleteUser(userFlags)
			handleError(err)
			resultData := output.Result{
				Result: string(customReqResult),
			}
			output.RenderResult(resultData, outputOptions)
		}
	},
}

var updateUserCmd = &cobra.Command{
	Use:     "user",
	Aliases: []string{"u"},
	Short:   "Update a user in lagoon",
	Long:    "Update a user in lagoon (change name, or email address)",
	Run: func(cmd *cobra.Command, args []string) {
		userFlags := parseUser(*cmd.Flags())
		if userFlags.Email == "" {
			fmt.Println("Missing arguments: Email address is not defined")
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
	Use:     "user-sshkeys",
	Aliases: []string{"us"},
	Short:   "Get a users SSH keys",
	Long:    `Get a users SSH keys. This will only work for users that are part of a group`,
	Run: func(cmd *cobra.Command, args []string) {
		if userEmail == "" {
			fmt.Println("Missing arguments: Email address is not defined")
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
	Use:     "all-user-sshkeys",
	Aliases: []string{"aus"},
	Short:   "Get all user SSH keys",
	Long:    `Get all user SSH keys. This will only work for users that are part of a group`,
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
	pubKeyValue      string
)

func init() {
	addUserCmd.Flags().StringVarP(&userFirstName, "firstName", "F", "", "Firstname of the user")
	addUserCmd.Flags().StringVarP(&userLastName, "lastName", "L", "", "Lastname of the user")
	addUserCmd.Flags().StringVarP(&userEmail, "email", "E", "", "Email address of the user")
	addUserSSHKeyCmd.Flags().StringVarP(&userEmail, "email", "E", "", "Email address of the user")
	addUserSSHKeyCmd.Flags().StringVarP(&sshKeyName, "keyname", "N", "", "Name of the sshkey (optional, if not provided will try use what is in the pubkey file)")
	addUserSSHKeyCmd.Flags().StringVarP(&pubKeyFile, "pubkey", "K", "", "Specify path to the public key to add")
	addUserSSHKeyCmd.Flags().StringVarP(&pubKeyValue, "keyvalue", "V", "", "Value of the public key to add (ssh-ed25519 AAA..)")
	deleteUserCmd.Flags().StringVarP(&userEmail, "email", "E", "", "Email address of the user")
	deleteSSHKeyCmd.Flags().StringVarP(&sshKeyName, "keyname", "N", "", "Name of the sshkey")
	updateUserCmd.Flags().StringVarP(&userFirstName, "firstName", "F", "", "New firstname of the user")
	updateUserCmd.Flags().StringVarP(&userLastName, "lastName", "L", "", "New lastname of the user")
	updateUserCmd.Flags().StringVarP(&userEmail, "email", "E", "", "New email address of the user")
	updateUserCmd.Flags().StringVarP(&currentUserEmail, "current-email", "C", "", "Current email address of the user")
	getUserKeysCmd.Flags().StringVarP(&userEmail, "email", "E", "", "New email address of the user")
	getUserKeysCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group to check users in (if not specified, will default to all groups)")
	getAllUserKeysCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group to list users in (if not specified, will default to all groups)")
}
