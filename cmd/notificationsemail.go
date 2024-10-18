package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/pkg/output"
	"github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	"github.com/uselagoon/machinery/api/schema"
)

var addNotificationEmailCmd = &cobra.Command{
	Use:   "email",
	Short: "Add a new email notification",
	Long: `Add a new email notification
This command is used to set up a new email notification in Lagoon. This requires information to talk to the email address to send to.
It does not configure a project to send notifications to email though, you need to use project-email for that.`,
	Aliases: []string{"e"},
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lContext.Name)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		email, err := cmd.Flags().GetString("email")
		if err != nil {
			return err
		}
		organizationID, err := cmd.Flags().GetUint("organization-id")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Name", name, "Email", email); err != nil {
			return err
		}
		if yesNo(fmt.Sprintf("You are attempting to create an email notification '%s' with email address '%s', are you sure?", name, email)) {
			utoken := lUser.UserConfig.Grant.AccessToken
			lc := lclient.New(
				fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
				lagoonCLIVersion,
				lContext.ContextConfig.Version,
				&utoken,
				debug)

			notification := schema.AddNotificationEmailInput{
				Name:         name,
				EmailAddress: email,
				Organization: &organizationID,
			}

			result, err := lagoon.AddNotificationEmail(context.TODO(), &notification, lc)
			if err != nil {
				return err
			}
			var data []output.Data
			notificationData := []string{
				returnNonEmptyString(fmt.Sprintf("%v", result.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", result.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", result.EmailAddress)),
			}
			if result.Organization != nil {
				organization, err := lagoon.GetOrganizationByID(context.TODO(), organizationID, lc)
				if err != nil {
					return err
				}
				notificationData = append(notificationData, organization.Name)
			} else {
				notificationData = append(notificationData, "-")
			}
			data = append(data, notificationData)
			r := output.RenderOutput(output.Table{
				Header: []string{
					"ID",
					"Name",
					"EmailAddress",
					"Organization",
				},
				Data: data,
			}, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		}
		return nil
	},
}

var addProjectNotificationEmailCmd = &cobra.Command{
	Use:     "project-email",
	Aliases: []string{"pe"},
	Short:   "Add an email notification to a project",
	Long: `Add an email notification to a project
This command is used to add an existing email notification in Lagoon to a project.`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lContext.Name)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Notification name", name, "Project name", cmdProjectName); err != nil {
			return err
		}
		if yesNo(fmt.Sprintf("You are attempting to add email notification '%s' to project '%s', are you sure?", name, cmdProjectName)) {
			utoken := lUser.UserConfig.Grant.AccessToken
			lc := lclient.New(
				fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
				lagoonCLIVersion,
				lContext.ContextConfig.Version,
				&utoken,
				debug)
			notification := &schema.AddNotificationToProjectInput{
				NotificationType: schema.EmailNotification,
				NotificationName: name,
				Project:          cmdProjectName,
			}
			_, err := lagoon.AddNotificationToProject(context.TODO(), notification, lc)
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

var listProjectEmailsCmd = &cobra.Command{
	Use:     "project-email",
	Aliases: []string{"pe"},
	Short:   "List email details about a project (alias: pe)",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lContext.Name)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Project name", cmdProjectName); err != nil {
			return err
		}
		utoken := lUser.UserConfig.Grant.AccessToken
		lc := lclient.New(
			fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
			lagoonCLIVersion,
			lContext.ContextConfig.Version,
			&utoken,
			debug)

		result, err := lagoon.GetProjectNotificationEmail(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		if len(result.Name) == 0 {
			return handleNilResults("No project found for '%s'\n", cmd, cmdProjectName)
		} else if len(result.Notifications.Email) == 0 {
			return handleNilResults("No email notificatons found for project: '%s'\n", cmd, cmdProjectName)
		}

		data := []output.Data{}
		if result.Notifications != nil {
			for _, notification := range result.Notifications.Email {
				data = append(data, []string{
					returnNonEmptyString(fmt.Sprintf("%v", notification.Name)),
					returnNonEmptyString(fmt.Sprintf("%v", notification.EmailAddress)),
				})
			}
		}
		r := output.RenderOutput(output.Table{
			Header: []string{
				"Name",
				"EmailAddress",
			},
			Data: data,
		}, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		return nil
	},
}

var listAllEmailsCmd = &cobra.Command{
	Use:     "email",
	Aliases: []string{"e"},
	Short:   "List all email notification details (alias: e)",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lContext.Name)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		utoken := lUser.UserConfig.Grant.AccessToken
		lc := lclient.New(
			fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
			lagoonCLIVersion,
			lContext.ContextConfig.Version,
			&utoken,
			debug)
		result, err := lagoon.GetAllNotificationEmail(context.TODO(), lc)
		if err != nil {
			return err
		}
		data := []output.Data{}
		for _, res := range *result {
			b, _ := json.Marshal(res.Notifications.Email)
			if string(b) != "null" {
				for _, notif := range res.Notifications.Email {
					data = append(data, []string{
						returnNonEmptyString(fmt.Sprintf("%v", res.Name)),
						returnNonEmptyString(fmt.Sprintf("%v", notif.Name)),
						returnNonEmptyString(fmt.Sprintf("%v", notif.EmailAddress)),
					})
				}
			}
		}
		r := output.RenderOutput(output.Table{
			Header: []string{
				"Project",
				"Name",
				"EmailAddress",
			},
			Data: data,
		}, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		return nil
	},
}

var deleteProjectEmailNotificationCmd = &cobra.Command{
	Use:     "project-email",
	Aliases: []string{"pe"},
	Short:   "Delete a email notification from a project",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lContext.Name)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Project name", cmdProjectName, "Notification name", name); err != nil {
			return err
		}

		utoken := lUser.UserConfig.Grant.AccessToken
		lc := lclient.New(
			fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
			lagoonCLIVersion,
			lContext.ContextConfig.Version,
			&utoken,
			debug)

		project, err := lagoon.GetProjectByName(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		if project.Name == "" {
			return handleNilResults("No project found for '%s'\n", cmd, cmdProjectName)
		}
		if yesNo(fmt.Sprintf("You are attempting to delete email notification '%s' from project '%s', are you sure?", name, cmdProjectName)) {
			notification := &schema.RemoveNotificationFromProjectInput{
				NotificationType: schema.EmailNotification,
				NotificationName: name,
				Project:          cmdProjectName,
			}
			_, err := lagoon.RemoveNotificationFromProject(context.TODO(), notification, lc)
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

var deleteEmailNotificationCmd = &cobra.Command{
	Use:     "email",
	Aliases: []string{"e"},
	Short:   "Delete an email notification from Lagoon",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lContext.Name)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Notification name", name); err != nil {
			return err
		}
		// Todo: Verify notifcation name exists - requires #PR https://github.com/uselagoon/lagoon/pull/3740
		if yesNo(fmt.Sprintf("You are attempting to delete email notification '%s', are you sure?", name)) {
			utoken := lUser.UserConfig.Grant.AccessToken
			lc := lclient.New(
				fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
				lagoonCLIVersion,
				lContext.ContextConfig.Version,
				&utoken,
				debug)
			result, err := lagoon.DeleteNotificationEmail(context.TODO(), name, lc)
			if err != nil {
				return err
			}
			resultData := output.Result{
				Result: result.DeleteNotification,
			}
			r := output.RenderResult(resultData, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		}
		return nil
	},
}

var updateEmailNotificationCmd = &cobra.Command{
	Use:     "email",
	Aliases: []string{"e"},
	Short:   "Update an existing email notification",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lContext.Name)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		newname, err := cmd.Flags().GetString("newname")
		if err != nil {
			return err
		}
		email, err := cmd.Flags().GetString("email")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Notification name", name); err != nil {
			return err
		}
		patch := schema.UpdateNotificationEmailPatchInput{
			Name:         nullStrCheck(newname),
			EmailAddress: nullStrCheck(email),
		}
		if patch == (schema.UpdateNotificationEmailPatchInput{}) {
			return fmt.Errorf("missing arguments: either email or newname must be defined")
		}

		if yesNo(fmt.Sprintf("You are attempting to update email notification '%s', are you sure?", name)) {
			utoken := lUser.UserConfig.Grant.AccessToken
			lc := lclient.New(
				fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
				lagoonCLIVersion,
				lContext.ContextConfig.Version,
				&utoken,
				debug)

			notification := &schema.UpdateNotificationEmailInput{
				Name:  name,
				Patch: patch,
			}
			result, err := lagoon.UpdateNotificationEmail(context.TODO(), notification, lc)
			if err != nil {
				return err
			}
			data := []output.Data{
				[]string{
					returnNonEmptyString(fmt.Sprintf("%v", result.ID)),
					returnNonEmptyString(fmt.Sprintf("%v", result.Name)),
					returnNonEmptyString(fmt.Sprintf("%v", result.EmailAddress)),
				},
			}
			r := output.RenderOutput(output.Table{
				Header: []string{
					"ID",
					"Name",
					"EmailAddress",
				},
				Data: data,
			}, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		}
		return nil
	},
}

func init() {
	addNotificationEmailCmd.Flags().StringP("name", "n", "", "The name of the notification")
	addNotificationEmailCmd.Flags().StringP("email", "E", "", "The email address of the notification")
	addNotificationEmailCmd.Flags().Uint("organization-id", 0, "ID of the Organization")
	addProjectNotificationEmailCmd.Flags().StringP("name", "n", "", "The name of the notification")
	deleteProjectEmailNotificationCmd.Flags().StringP("name", "n", "", "The name of the notification")
	deleteEmailNotificationCmd.Flags().StringP("name", "n", "", "The name of the notification")
	updateEmailNotificationCmd.Flags().StringP("name", "n", "", "The current name of the notification")
	updateEmailNotificationCmd.Flags().StringP("newname", "N", "", "The name of the notification")
	updateEmailNotificationCmd.Flags().StringP("email", "E", "", "The email address of the notification")
}
