package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	l "github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	s "github.com/uselagoon/machinery/api/schema"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/uselagoon/lagoon-cli/pkg/api"
	"github.com/uselagoon/lagoon-cli/pkg/output"
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
	// lowercase user email address
	parsedFlags.Email = strings.ToLower(parsedFlags.Email)
	return parsedFlags
}

func parseSSHKeyFile(sshPubKey string, keyName string, keyValue string, userEmail string) (api.SSHKey, error) {
	// if we haven't got a keyvalue
	if keyValue == "" {
		b, err := ioutil.ReadFile(sshPubKey) // just pass the file name
		handleError(err)
		keyValue = string(b)
	}
	splitKey := strings.Split(keyValue, " ")
	var keyType api.SSHKeyType
	var err error

	// will fail if value is not right
	if strings.EqualFold(string(splitKey[0]), "ssh-rsa") {
		keyType = api.SSHRsa
	} else if strings.EqualFold(string(splitKey[0]), "ssh-ed25519") {
		keyType = api.SSHEd25519
	} else if strings.EqualFold(string(splitKey[0]), "ecdsa-sha2-nistp256") {
		keyType = api.SSHECDSA256
	} else if strings.EqualFold(string(splitKey[0]), "ecdsa-sha2-nistp384") {
		keyType = api.SSHECDSA384
	} else if strings.EqualFold(string(splitKey[0]), "ecdsa-sha2-nistp521") {
		keyType = api.SSHECDSA521
	} else {
		// return error stating key type not supported
		keyType = api.SSHRsa
		err = errors.New(fmt.Sprintf("SSH key type %s not supported", string(splitKey[0])))
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
	return parsedFlags, err
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
	Short:   "Add an SSH key to a user",
	Long: `Add an SSH key to a user

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
		var err error
		userSSHKey, err := parseSSHKeyFile(pubKeyFile, sshKeyName, pubKeyValue, userFlags.Email)
		handleError(err)
		var customReqResult []byte
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
	Short:   "Delete an SSH key from Lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		if sshKeyName == "" {
			fmt.Println("Missing arguments: SSH key name is not defined")
			cmd.Help()
			os.Exit(1)
		}
		var customReqResult []byte
		var err error
		if yesNo(fmt.Sprintf("You are attempting to delete SSH key named '%s', are you sure?", sshKeyName)) {
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
	Short:   "Delete a user from Lagoon",
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
	Short:   "Update a user in Lagoon",
	Long:    "Update a user in Lagoon (change name, or email address)",
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
			Email: strings.ToLower(currentUserEmail),
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
	Short:   "Get a user's SSH keys",
	Long:    `Get a user's SSH keys. This will only work for users that are part of a group`,
	Run: func(cmd *cobra.Command, args []string) {
		if userEmail == "" {
			fmt.Println("Missing arguments: Email address is not defined")
			cmd.Help()
			os.Exit(1)
		}
		returnedJSON, err := uClient.ListUserSSHKeys(groupName, strings.ToLower(userEmail), false)
		handleError(err)
		var dataMain output.Table
		err = json.Unmarshal([]byte(returnedJSON), &dataMain)
		handleError(err)
		if len(dataMain.Data) == 0 {
			output.RenderInfo(fmt.Sprintf("No SSH keys for user '%s'", strings.ToLower(userEmail)), outputOptions)
			os.Exit(0)
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
		returnedJSON, err := uClient.ListUserSSHKeys(groupName, strings.ToLower(userEmail), true)
		handleError(err)
		var dataMain output.Table
		err = json.Unmarshal([]byte(returnedJSON), &dataMain)
		handleError(err)
		if len(dataMain.Data) == 0 {
			output.RenderInfo("No SSH keys for any users", outputOptions)
			os.Exit(0)
		}
		output.RenderOutput(dataMain, outputOptions)

	},
}

var addUserToOrganizationCmd = &cobra.Command{
	Use:     "user",
	Aliases: []string{"u"},
	Short:   "Add a user to an Organization",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		handleError(err)

		organizationName, err := cmd.Flags().GetString("organization")
		requiredInputCheck("Organization name", organizationName)
		if err != nil {
			return err
		}
		userEmail, err := cmd.Flags().GetString("email")
		requiredInputCheck("User email", userEmail)
		if err != nil {
			return err
		}
		owner, err := cmd.Flags().GetBool("owner")
		if err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)

		organization, err := l.GetOrganizationByName(context.TODO(), organizationName, lc)
		handleError(err)

		userInput := s.AddUserToOrganizationInput{
			User:         s.UserInput{Email: userEmail},
			Organization: organization.ID,
			Owner:        owner,
		}

		orgUser := s.Organization{}
		err = lc.AddUserToOrganization(context.TODO(), &userInput, &orgUser)
		handleError(err)

		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"User":              userEmail,
				"Organization Name": organizationName,
			},
		}
		output.RenderResult(resultData, outputOptions)
		return nil
	},
}

var (
	currentUserEmail string
	pubKeyValue      string
)

func init() {
	addUserCmd.Flags().StringVarP(&userFirstName, "firstName", "F", "", "First name of the user")
	addUserCmd.Flags().StringVarP(&userLastName, "lastName", "L", "", "Last name of the user")
	addUserCmd.Flags().StringVarP(&userEmail, "email", "E", "", "Email address of the user")
	addUserSSHKeyCmd.Flags().StringVarP(&userEmail, "email", "E", "", "Email address of the user")
	addUserSSHKeyCmd.Flags().StringVarP(&sshKeyName, "keyname", "N", "", "Name of the SSH key (optional, if not provided will try use what is in the pubkey file)")
	addUserSSHKeyCmd.Flags().StringVarP(&pubKeyFile, "pubkey", "K", "", "Specify path to the public key to add")
	addUserSSHKeyCmd.Flags().StringVarP(&pubKeyValue, "keyvalue", "V", "", "Value of the public key to add (ssh-ed25519 AAA..)")
	deleteUserCmd.Flags().StringVarP(&userEmail, "email", "E", "", "Email address of the user")
	deleteSSHKeyCmd.Flags().StringVarP(&sshKeyName, "keyname", "N", "", "Name of the SSH key")
	updateUserCmd.Flags().StringVarP(&userFirstName, "firstName", "F", "", "New first name of the user")
	updateUserCmd.Flags().StringVarP(&userLastName, "lastName", "L", "", "New last name of the user")
	updateUserCmd.Flags().StringVarP(&userEmail, "email", "E", "", "New email address of the user")
	updateUserCmd.Flags().StringVarP(&currentUserEmail, "current-email", "C", "", "Current email address of the user")
	getUserKeysCmd.Flags().StringVarP(&userEmail, "email", "E", "", "New email address of the user")
	getUserKeysCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group to check users in (if not specified, will default to all groups)")
	getAllUserKeysCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group to list users in (if not specified, will default to all groups)")

	addUserToOrganizationCmd.Flags().StringP("organization", "O", "", "Name of the organization")
	addUserToOrganizationCmd.Flags().StringP("email", "E", "", "Email address of the user")
	addUserToOrganizationCmd.Flags().Bool("owner", false, "Set the user as an owner of the organization")
}
