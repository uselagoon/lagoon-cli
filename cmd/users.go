package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	l "github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	ls "github.com/uselagoon/machinery/api/schema"

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

func parseSSHKeyFile(sshPubKey string, keyName string, keyValue string, userEmail string) (ls.AddSSHKeyInput, error) {
	// if we haven't got a keyvalue
	if keyValue == "" {
		b, err := os.ReadFile(sshPubKey) // just pass the file name
		handleError(err)
		keyValue = string(b)
	}
	splitKey := strings.Split(keyValue, " ")
	var keyType ls.SSHKeyType
	var err error

	// will fail if value is not right
	if strings.EqualFold(string(splitKey[0]), "ssh-rsa") {
		keyType = ls.SSHRsa
	} else if strings.EqualFold(string(splitKey[0]), "ssh-ed25519") {
		keyType = ls.SSHEd25519
	} else if strings.EqualFold(string(splitKey[0]), "ecdsa-sha2-nistp256") {
		keyType = ls.SSHECDSA256
	} else if strings.EqualFold(string(splitKey[0]), "ecdsa-sha2-nistp384") {
		keyType = ls.SSHECDSA384
	} else if strings.EqualFold(string(splitKey[0]), "ecdsa-sha2-nistp521") {
		keyType = ls.SSHECDSA521
	} else {
		// return error stating key type not supported
		keyType = ls.SSHRsa
		err = fmt.Errorf(fmt.Sprintf("SSH key type %s not supported", splitKey[0]))
	}

	// if the sshkey has a comment/name in it, we can use that
	if keyName == "" && len(splitKey) == 3 {
		//strip new line
		keyName = stripNewLines(splitKey[2])
	} else if keyName == "" && len(splitKey) == 2 {
		keyName = userEmail
		output.RenderInfo("no name provided, using email address as key name", outputOptions)
	}
	SSHKeyInput := ls.AddSSHKeyInput{
		SSHKey: ls.SSHKey{
			KeyType:  keyType,
			KeyValue: stripNewLines(splitKey[1]),
			Name:     keyName,
		},
		UserEmail: userEmail,
	}

	return SSHKeyInput, err
}

var addUserCmd = &cobra.Command{
	Use:     "user",
	Aliases: []string{"u"},
	Short:   "Add a user to lagoon",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		firstName, err := cmd.Flags().GetString("first-name")
		if err != nil {
			return err
		}
		LastName, err := cmd.Flags().GetString("last-name")
		if err != nil {
			return err
		}
		email, err := cmd.Flags().GetString("email")
		if err != nil {
			return err
		}
		resetPassword, err := cmd.Flags().GetBool("reset-password")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Email address", email); err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)

		userInput := &ls.AddUserInput{
			FirstName:     firstName,
			LastName:      LastName,
			Email:         email,
			ResetPassword: resetPassword,
		}
		user, err := l.AddUser(context.TODO(), userInput, lc)
		if err != nil {
			return err
		}

		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"id": user.ID,
			},
		}
		output.RenderResult(resultData, outputOptions)
		return nil
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
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		pubKeyFile, err := cmd.Flags().GetString("pubkey")
		if err != nil {
			return err
		}
		sshKeyName, err := cmd.Flags().GetString("keyname")
		if err != nil {
			return err
		}
		pubKeyValue, err := cmd.Flags().GetString("keyvalue")
		if err != nil {
			return err
		}
		email, err := cmd.Flags().GetString("email")
		if err != nil {
			return err
		}

		if err := requiredInputCheck("Email address", email); err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)

		userSSHKey, err := parseSSHKeyFile(pubKeyFile, sshKeyName, pubKeyValue, email)
		if err != nil {
			return err
		}
		result, err := l.AddSSHKey(context.TODO(), &userSSHKey, lc)
		if err != nil {
			return err
		}

		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"ID": result.ID,
			},
		}
		output.RenderResult(resultData, outputOptions)
		return nil
	},
}

var deleteSSHKeyCmd = &cobra.Command{
	Use:     "user-sshkey",
	Aliases: []string{"uk"},
	Short:   "Delete an SSH key from Lagoon",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		sshKeyID, err := cmd.Flags().GetUint("id")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("SSH key ID", strconv.Itoa(int(sshKeyID))); err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)

		if yesNo(fmt.Sprintf("You are attempting to delete SSH key ID:'%d', are you sure?", sshKeyID)) {
			_, err := l.RemoveSSHKey(context.TODO(), sshKeyID, lc)
			if err != nil {
				return err
			}
			resultData := output.Result{
				Result: "success",
			}
			output.RenderResult(resultData, outputOptions)
		}
		return nil
	},
}

var deleteUserCmd = &cobra.Command{
	Use:     "user",
	Aliases: []string{"u"},
	Short:   "Delete a user from Lagoon",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		emailAddress, err := cmd.Flags().GetString("email")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Email address", emailAddress); err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)

		deleteUserInput := &ls.DeleteUserInput{
			User: ls.UserInput{Email: emailAddress},
		}
		if yesNo(fmt.Sprintf("You are attempting to delete user with email address '%s', are you sure?", emailAddress)) {
			_, err := l.DeleteUser(context.TODO(), deleteUserInput, lc)
			if err != nil {
				return err
			}
			resultData := output.Result{
				Result: "success",
			}
			output.RenderResult(resultData, outputOptions)
		}
		return nil
	},
}

var updateUserCmd = &cobra.Command{
	Use:     "user",
	Aliases: []string{"u"},
	Short:   "Update a user in Lagoon",
	Long:    "Update a user in Lagoon (change name, or email address)",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		emailAddress, err := cmd.Flags().GetString("email")
		if err != nil {
			return err
		}
		firstName, err := cmd.Flags().GetString("first-name")
		if err != nil {
			return err
		}
		lastName, err := cmd.Flags().GetString("last-name")
		if err != nil {
			return err
		}
		currentEmail, err := cmd.Flags().GetString("current-email")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Current email address", currentEmail); err != nil {
			return err
		}
		if firstName == "" && lastName == "" && emailAddress == "" {
			cmd.Help()
			output.RenderError("Missing arguments: Nothing to update, please provide a field to update", outputOptions)
			return nil
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)

		currentUser := &ls.UpdateUserInput{
			User: ls.UserInput{
				Email: strings.ToLower(currentEmail),
			},
			Patch: ls.UpdateUserPatchInput{
				Email:     nullStrCheck(strings.ToLower(emailAddress)),
				FirstName: nullStrCheck(firstName),
				LastName:  nullStrCheck(lastName),
			},
		}

		user, err := l.UpdateUser(context.TODO(), currentUser, lc)
		if err != nil {
			return err
		}

		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"ID": user.ID,
			},
		}
		output.RenderResult(resultData, outputOptions)
		return nil
	},
}

var getUserKeysCmd = &cobra.Command{
	//@TODO: once individual user interaction comes in, this will need to be adjusted
	Use:     "user-sshkeys",
	Aliases: []string{"us"},
	Short:   "Get a user's SSH keys",
	Long:    `Get a user's SSH keys. This will only work for users that are part of a group`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		userEmail, err := cmd.Flags().GetString("email")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Email address", userEmail); err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)
		userKeys, err := l.GetUserSSHKeysByEmail(context.TODO(), userEmail, lc)
		if err != nil {
			return err
		}
		if len(userKeys.SSHKeys) == 0 {
			output.RenderInfo(fmt.Sprintf("No SSH keys for user '%s'", strings.ToLower(userEmail)), outputOptions)
			return nil
		}

		data := []output.Data{}
		for _, userkey := range userKeys.SSHKeys {
			data = append(data, []string{
				strconv.Itoa(int(userkey.ID)),
				userKeys.Email,
				userkey.Name,
				string(userkey.KeyType),
				userkey.KeyValue,
			})
		}

		dataMain := output.Table{
			Header: []string{"ID", "Email", "Name", "Type", "Value"},
			Data:   data,
		}

		output.RenderOutput(dataMain, outputOptions)
		return nil
	},
}

var getAllUserKeysCmd = &cobra.Command{
	//@TODO: once individual user interaction comes in, this will need to be adjusted
	Use:     "all-user-sshkeys",
	Aliases: []string{"aus"},
	Short:   "Get all user SSH keys",
	Long:    `Get all user SSH keys. This will only work for users that are part of a group`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		groupName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)
		groupMembers, err := l.ListAllGroupMembersWithKeys(context.TODO(), groupName, lc)
		if err != nil {
			return err
		}

		var userGroups []ls.AddSSHKeyInput
		for _, group := range *groupMembers {
			for _, member := range group.Members {
				for _, key := range member.User.SSHKeys {
					userGroups = append(userGroups, ls.AddSSHKeyInput{SSHKey: key, UserEmail: member.User.Email})
				}
			}
		}
		if len(userGroups) == 0 {
			outputOptions.Error = fmt.Sprintf("No SSH keys for group '%s'\n", groupName)
		}
		var data []output.Data
		for _, userData := range userGroups {
			keyID := strconv.Itoa(int(userData.SSHKey.ID))
			userEmail := returnNonEmptyString(strings.Replace(userData.UserEmail, " ", "_", -1))
			keyName := returnNonEmptyString(strings.Replace(userData.SSHKey.Name, " ", "_", -1))
			keyValue := returnNonEmptyString(strings.Replace(userData.SSHKey.KeyValue, " ", "_", -1))
			keyType := returnNonEmptyString(strings.Replace(string(userData.SSHKey.KeyType), " ", "_", -1))
			data = append(data, []string{
				keyID,
				userEmail,
				keyName,
				keyType,
				keyValue,
			})
		}

		dataMain := output.Table{
			Header: []string{"ID", "Email", "Name", "Type", "Value"},
			Data:   data,
		}

		output.RenderOutput(dataMain, outputOptions)
		return nil
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
		if err != nil {
			return err
		}
		organizationName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Organization name", organizationName); err != nil {
			return err
		}
		userEmail, err := cmd.Flags().GetString("email")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("User email", userEmail); err != nil {
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
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)

		organization, err := l.GetOrganizationByName(context.TODO(), organizationName, lc)
		if err != nil {
			return err
		}

		userInput := ls.AddUserToOrganizationInput{
			User:         ls.UserInput{Email: userEmail},
			Organization: organization.ID,
			Owner:        owner,
		}

		orgUser := ls.Organization{}
		err = lc.AddUserToOrganization(context.TODO(), &userInput, &orgUser)
		if err != nil {
			return err
		}

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

var RemoveUserFromOrganization = &cobra.Command{
	Use:     "user",
	Aliases: []string{"u"},
	Short:   "Remove a user to an Organization",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		organizationName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Organization name", organizationName); err != nil {
			return err
		}
		userEmail, err := cmd.Flags().GetString("email")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("User email", userEmail); err != nil {
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
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)

		organization, err := l.GetOrganizationByName(context.TODO(), organizationName, lc)
		if err != nil {
			return err
		}

		userInput := ls.AddUserToOrganizationInput{
			User:         ls.UserInput{Email: userEmail},
			Organization: organization.ID,
			Owner:        owner,
		}

		orgUser := ls.Organization{}

		if yesNo(fmt.Sprintf("You are attempting to remove user '%s' from organization '%s'. This removes the users ability to view or manage the organizations groups, projects, & notifications, are you sure?", userEmail, organization.Name)) {
			err = lc.RemoveUserFromOrganization(context.TODO(), &userInput, &orgUser)
			if err != nil {
				return err
			}
			resultData := output.Result{
				Result: "success",
				ResultData: map[string]interface{}{
					"User":              userEmail,
					"Organization Name": organizationName,
				},
			}
			output.RenderResult(resultData, outputOptions)
		}
		return nil
	},
}

var resetPasswordCmd = &cobra.Command{
	Use:     "reset-password",
	Aliases: []string{"reset-pass", "rp"},
	Short:   "Send a password reset email",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		userEmail, err := cmd.Flags().GetString("email")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Email address", userEmail); err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)

		resetPasswordInput := ls.ResetUserPasswordInput{
			User: ls.UserInput{Email: userEmail},
		}

		if yesNo(fmt.Sprintf("You are attempting to send a password reset email to '%s', are you sure?", userEmail)) {
			_, err := l.ResetUserPassword(context.TODO(), &resetPasswordInput, lc)
			if err != nil {
				return err
			}
			resultData := output.Result{
				Result: "success",
			}
			output.RenderResult(resultData, outputOptions)
		}
		return nil
	},
}

func init() {
	addUserCmd.Flags().StringP("first-name", "F", "", "First name of the user")
	addUserCmd.Flags().StringP("last-name", "L", "", "Last name of the user")
	addUserCmd.Flags().StringP("email", "E", "", "Email address of the user")
	addUserCmd.Flags().BoolP("reset-password", "", false, "Send a password reset email")
	addUserSSHKeyCmd.Flags().StringP("email", "E", "", "Email address of the user")
	addUserSSHKeyCmd.Flags().StringP("keyname", "N", "", "Name of the SSH key (optional, if not provided will try use what is in the pubkey file)")
	addUserSSHKeyCmd.Flags().StringP("pubkey", "K", "", "Specify path to the public key to add")
	addUserSSHKeyCmd.Flags().StringP("keyvalue", "V", "", "Value of the public key to add (ssh-ed25519 AAA..)")
	deleteUserCmd.Flags().StringP("email", "E", "", "Email address of the user")
	deleteSSHKeyCmd.Flags().Uint("id", 0, "ID of the SSH key")
	updateUserCmd.Flags().StringP("first-name", "F", "", "New first name of the user")
	updateUserCmd.Flags().StringP("last-name", "L", "", "New last name of the user")
	updateUserCmd.Flags().StringP("email", "E", "", "New email address of the user")
	updateUserCmd.Flags().StringP("current-email", "C", "", "Current email address of the user")
	getUserKeysCmd.Flags().StringP("email", "E", "", "New email address of the user")
	getAllUserKeysCmd.Flags().StringP("name", "N", "", "Name of the group to list users in (if not specified, will default to all groups)")
	addUserToOrganizationCmd.Flags().StringP("name", "O", "", "Name of the organization")
	addUserToOrganizationCmd.Flags().StringP("email", "E", "", "Email address of the user")
	addUserToOrganizationCmd.Flags().Bool("owner", false, "Set the user as an owner of the organization")
	RemoveUserFromOrganization.Flags().StringP("name", "O", "", "Name of the organization")
	RemoveUserFromOrganization.Flags().StringP("email", "E", "", "Email address of the user")
	RemoveUserFromOrganization.Flags().Bool("owner", false, "Set the user as an owner of the organization")
	resetPasswordCmd.Flags().StringP("email", "E", "", "Email address of the user")
}
