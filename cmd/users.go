package cmd

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	"github.com/uselagoon/machinery/api/schema"

	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/pkg/output"
)

func parseSSHKeyFile(sshPubKey string, keyName string, keyValue string, userEmail string) (schema.AddSSHKeyInput, error) {
	// if we haven't got a keyvalue
	if keyValue == "" {
		b, err := os.ReadFile(sshPubKey) // just pass the file name
		handleError(err)
		keyValue = string(b)
	}
	splitKey := strings.Split(keyValue, " ")
	var keyType schema.SSHKeyType
	var err error

	// will fail if value is not right
	if strings.EqualFold(string(splitKey[0]), "ssh-rsa") {
		keyType = schema.SSHRsa
	} else if strings.EqualFold(string(splitKey[0]), "ssh-ed25519") {
		keyType = schema.SSHEd25519
	} else if strings.EqualFold(string(splitKey[0]), "ecdsa-sha2-nistp256") {
		keyType = schema.SSHECDSA256
	} else if strings.EqualFold(string(splitKey[0]), "ecdsa-sha2-nistp384") {
		keyType = schema.SSHECDSA384
	} else if strings.EqualFold(string(splitKey[0]), "ecdsa-sha2-nistp521") {
		keyType = schema.SSHECDSA521
	} else {
		// return error stating key type not supported
		keyType = schema.SSHRsa
		err = fmt.Errorf(fmt.Sprintf("SSH key type %s not supported", splitKey[0]))
	}

	// if the sshkey has a comment/name in it, we can use that
	if keyName == "" && len(splitKey) == 3 {
		//strip new line
		keyName = stripNewLines(splitKey[2])
	} else if keyName == "" && len(splitKey) == 2 {
		keyName = userEmail
		output.RenderInfo("no name provided, using email address as key name\n", outputOptions)
	}
	SSHKeyInput := schema.AddSSHKeyInput{
		SSHKey: schema.SSHKey{
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

		userInput := &schema.AddUserInput{
			FirstName:     firstName,
			LastName:      LastName,
			Email:         email,
			ResetPassword: resetPassword,
		}
		user, err := lagoon.AddUser(context.TODO(), userInput, lc)
		if err != nil {
			return err
		}

		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"id": user.ID,
			},
		}
		r := output.RenderResult(resultData, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
		result, err := lagoon.AddSSHKey(context.TODO(), &userSSHKey, lc)
		if err != nil {
			return err
		}

		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"ID": result.ID,
			},
		}
		r := output.RenderResult(resultData, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
			_, err := lagoon.RemoveSSHKey(context.TODO(), sshKeyID, lc)
			if err != nil {
				return err
			}
			resultData := output.Result{
				Result: "success",
			}
			r := output.RenderResult(resultData, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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

		deleteUserInput := &schema.DeleteUserInput{
			User: schema.UserInput{Email: emailAddress},
		}
		if yesNo(fmt.Sprintf("You are attempting to delete user with email address '%s', are you sure?", emailAddress)) {
			_, err := lagoon.DeleteUser(context.TODO(), deleteUserInput, lc)
			if err != nil {
				return err
			}
			resultData := output.Result{
				Result: "success",
			}
			r := output.RenderResult(resultData, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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

		currentUser := &schema.UpdateUserInput{
			User: schema.UserInput{
				Email: strings.ToLower(currentEmail),
			},
			Patch: schema.UpdateUserPatchInput{
				Email:     nullStrCheck(strings.ToLower(emailAddress)),
				FirstName: nullStrCheck(firstName),
				LastName:  nullStrCheck(lastName),
			},
		}

		user, err := lagoon.UpdateUser(context.TODO(), currentUser, lc)
		if err != nil {
			return err
		}

		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"ID": user.ID,
			},
		}
		r := output.RenderResult(resultData, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
		userKeys, err := lagoon.GetUserSSHKeysByEmail(context.TODO(), userEmail, lc)
		if err != nil {
			return err
		}
		if len(userKeys.SSHKeys) == 0 {
			output.RenderInfo(fmt.Sprintf("There are no SSH keys for user '%s'\n", strings.ToLower(userEmail)), outputOptions)
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

		outputOptions.MultiLine = true
		r := output.RenderOutput(dataMain, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
		groupMembers, err := lagoon.ListAllGroupMembersWithKeys(context.TODO(), groupName, lc)
		if err != nil {
			return err
		}

		var userGroups []schema.AddSSHKeyInput
		for _, group := range *groupMembers {
			for _, member := range group.Members {
				for _, key := range member.User.SSHKeys {
					userGroups = append(userGroups, schema.AddSSHKeyInput{SSHKey: key, UserEmail: member.User.Email})
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
		outputOptions.MultiLine = true
		r := output.RenderOutput(dataMain, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		return nil
	},
}

var addAdministratorToOrganizationCmd = &cobra.Command{
	Use:     "organization-administrator",
	Aliases: []string{"organization-admin", "org-admin", "org-a"},
	Short:   "Add an administrator to an Organization",
	Long:    "Add an administrator to an Organization. If the role flag is not provided users will be added as viewers",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}

		organizationName, err := cmd.Flags().GetString("organization-name")
		if err != nil {
			return err
		}
		userEmail, err := cmd.Flags().GetString("email")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Organization name", organizationName, "User email", userEmail); err != nil {
			return err
		}
		role, err := cmd.Flags().GetString("role")
		if err != nil {
			return err
		}
		userInput := schema.AddUserToOrganizationInput{
			User: schema.UserInput{Email: userEmail},
		}
		switch strings.ToLower(role) {
		case "viewer":
		case "admin":
			userInput.Admin = true
		case "owner":
			userInput.Owner = true
		default:
			return fmt.Errorf(`role '%s' is not valid - valid roles include "viewer", "admin", or "owner"`, role)
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)

		organization, err := lagoon.GetOrganizationByName(context.TODO(), organizationName, lc)
		if err != nil {
			return err
		}
		if organization.Name == "" {
			return fmt.Errorf("error querying organization by name")
		}
		userInput.Organization = organization.ID

		orgUser := schema.Organization{}
		err = lc.AddAdminToOrganization(context.TODO(), &userInput, &orgUser)
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
		r := output.RenderResult(resultData, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		return nil
	},
}

var removeAdministratorFromOrganizationCmd = &cobra.Command{
	Use:     "organization-administrator",
	Aliases: []string{"organization-admin", "org-admin", "org-a"},
	Short:   "Remove an administrator from an Organization",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}

		organizationName, err := cmd.Flags().GetString("organization-name")
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

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)

		organization, err := lagoon.GetOrganizationByName(context.TODO(), organizationName, lc)
		if err != nil {
			return err
		}
		if organization.Name == "" {
			return fmt.Errorf("error querying organization by name")
		}

		userInput := schema.AddUserToOrganizationInput{
			User:         schema.UserInput{Email: userEmail},
			Organization: organization.ID,
		}

		orgUser := schema.Organization{}

		if yesNo(fmt.Sprintf("You are attempting to remove user '%s' from organization '%s'. This removes the users ability to view or manage the organizations groups, projects, & notifications, are you sure?", userEmail, organization.Name)) {
			err = lc.RemoveAdminFromOrganization(context.TODO(), &userInput, &orgUser)
			handleError(err)
			resultData := output.Result{
				Result: "success",
				ResultData: map[string]interface{}{
					"User":              userEmail,
					"Organization Name": organizationName,
				},
			}
			r := output.RenderResult(resultData, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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

		resetPasswordInput := schema.ResetUserPasswordInput{
			User: schema.UserInput{Email: userEmail},
		}

		if yesNo(fmt.Sprintf("You are attempting to send a password reset email to '%s', are you sure?", userEmail)) {
			_, err := lagoon.ResetUserPassword(context.TODO(), &resetPasswordInput, lc)
			if err != nil {
				return err
			}
			resultData := output.Result{
				Result: "success",
			}
			r := output.RenderResult(resultData, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
	addAdministratorToOrganizationCmd.Flags().StringP("organization-name", "O", "", "Name of the organization")
	addAdministratorToOrganizationCmd.Flags().StringP("email", "E", "", "Email address of the user")
	addAdministratorToOrganizationCmd.Flags().StringP("role", "R", "", "Role in the organization [owner, admin, viewer]")
	removeAdministratorFromOrganizationCmd.Flags().StringP("organization-name", "O", "", "Name of the organization")
	removeAdministratorFromOrganizationCmd.Flags().StringP("email", "E", "", "Email address of the user")
	resetPasswordCmd.Flags().StringP("email", "E", "", "Email address of the user")
}
