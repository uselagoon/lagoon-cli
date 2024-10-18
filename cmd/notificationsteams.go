package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	"github.com/uselagoon/machinery/api/schema"

	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/pkg/output"
)

var addNotificationMicrosoftTeamsCmd = &cobra.Command{
	Use:   "microsoftteams",
	Short: "Add a new Microsoft Teams notification",
	Long: `Add a new Microsoft Teams notification
This command is used to set up a new Microsoft Teams notification in Lagoon. This requires information to talk to the webhook like the webhook URL.
It does not configure a project to send notifications to Microsoft Teams though, you need to use project-microsoftteams for that.`,
	Aliases: []string{"m"},
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
		webhook, err := cmd.Flags().GetString("webhook")
		if err != nil {
			return err
		}
		organizationID, err := cmd.Flags().GetUint("organization-id")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Notification name", name, "Webhook", webhook); err != nil {
			return err
		}
		if yesNo(fmt.Sprintf("You are attempting to create a Microsoft Teams notification '%s' with webhook url '%s', are you sure?", name, webhook)) {
			utoken := lUser.UserConfig.Grant.AccessToken
			lc := lclient.New(
				fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
				lagoonCLIVersion,
				lContext.ContextConfig.Version,
				&utoken,
				debug)

			notification := schema.AddNotificationMicrosoftTeamsInput{
				Name:         name,
				Webhook:      webhook,
				Organization: &organizationID,
			}

			result, err := lagoon.AddNotificationMicrosoftTeams(context.TODO(), &notification, lc)
			if err != nil {
				return err
			}
			var data []output.Data
			notificationData := []string{
				returnNonEmptyString(fmt.Sprintf("%v", result.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", result.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", result.Webhook)),
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
					"Webhook",
					"Organization",
				},
				Data: data,
			}, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		}
		return nil
	},
}

var addProjectNotificationMicrosoftTeamsCmd = &cobra.Command{
	Use:     "project-microsoftteams",
	Aliases: []string{"pm"},
	Short:   "Add a Microsoft Teams notification to a project",
	Long: `Add a Microsoft Teams notification to a project
This command is used to add an existing Microsoft Teams notification in Lagoon to a project.`,
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
		if yesNo(fmt.Sprintf("You are attempting to add Microsoft Teams notification '%s' to project '%s', are you sure?", name, cmdProjectName)) {
			utoken := lUser.UserConfig.Grant.AccessToken
			lc := lclient.New(
				fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
				lagoonCLIVersion,
				lContext.ContextConfig.Version,
				&utoken,
				debug)
			notification := &schema.AddNotificationToProjectInput{
				NotificationType: schema.MicrosoftTeamsNotification,
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

var listProjectMicrosoftTeamsCmd = &cobra.Command{
	Use:     "project-microsoftteams",
	Aliases: []string{"pm"},
	Short:   "List Microsoft Teams details about a project (alias: pm)",
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

		result, err := lagoon.GetProjectNotificationMicrosoftTeams(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		if len(result.Name) == 0 {
			return handleNilResults("No project found for '%s'\n", cmd, cmdProjectName)
		} else if len(result.Notifications.MicrosoftTeams) == 0 {
			return handleNilResults("No microsoft teams notificatons found for project: '%s'\n", cmd, cmdProjectName)
		}

		data := []output.Data{}
		if result.Notifications != nil {
			for _, notification := range result.Notifications.MicrosoftTeams {
				data = append(data, []string{
					returnNonEmptyString(fmt.Sprintf("%v", notification.Name)),
					returnNonEmptyString(fmt.Sprintf("%v", notification.Webhook)),
				})
			}
		}
		r := output.RenderOutput(output.Table{
			Header: []string{
				"Name",
				"Webhook",
			},
			Data: data,
		}, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		return nil
	},
}

var listAllMicrosoftTeamsCmd = &cobra.Command{
	Use:     "microsoftteams",
	Aliases: []string{"m"},
	Short:   "List all Microsoft Teams notification details (alias: m)",
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
		result, err := lagoon.GetAllNotificationMicrosoftTeams(context.TODO(), lc)
		if err != nil {
			return err
		}
		data := []output.Data{}
		for _, res := range *result {
			b, _ := json.Marshal(res.Notifications.MicrosoftTeams)
			if string(b) != "null" {
				for _, notif := range res.Notifications.MicrosoftTeams {
					data = append(data, []string{
						returnNonEmptyString(fmt.Sprintf("%v", res.Name)),
						returnNonEmptyString(fmt.Sprintf("%v", notif.Name)),
						returnNonEmptyString(fmt.Sprintf("%v", notif.Webhook)),
					})
				}
			}
		}
		r := output.RenderOutput(output.Table{
			Header: []string{
				"Project",
				"Name",
				"Webhook",
			},
			Data: data,
		}, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		return nil
	},
}

var deleteProjectMicrosoftTeamsNotificationCmd = &cobra.Command{
	Use:     "project-microsoftteams",
	Aliases: []string{"pm"},
	Short:   "Delete a Microsoft Teams notification from a project",
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

		if yesNo(fmt.Sprintf("You are attempting to delete Microsoft Teams notification '%s' from project '%s', are you sure?", name, cmdProjectName)) {
			notification := &schema.RemoveNotificationFromProjectInput{
				NotificationType: schema.MicrosoftTeamsNotification,
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

var deleteMicrosoftTeamsNotificationCmd = &cobra.Command{
	Use:     "microsoftteams",
	Aliases: []string{"m"},
	Short:   "Delete a Microsoft Teams notification from Lagoon",
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
		if yesNo(fmt.Sprintf("You are attempting to delete Microsoft Teams notification '%s', are you sure?", name)) {
			utoken := lUser.UserConfig.Grant.AccessToken
			lc := lclient.New(
				fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
				lagoonCLIVersion,
				lContext.ContextConfig.Version,
				&utoken,
				debug)
			result, err := lagoon.DeleteNotificationMicrosoftTeams(context.TODO(), name, lc)
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

var updateMicrosoftTeamsNotificationCmd = &cobra.Command{
	Use:     "microsoftteams",
	Aliases: []string{"m"},
	Short:   "Update an existing Microsoft Teams notification",
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
		webhook, err := cmd.Flags().GetString("webhook")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Notification name", name); err != nil {
			return err
		}
		patch := schema.UpdateNotificationMicrosoftTeamsPatchInput{
			Name:    nullStrCheck(newname),
			Webhook: nullStrCheck(webhook),
		}
		if patch == (schema.UpdateNotificationMicrosoftTeamsPatchInput{}) {
			return fmt.Errorf("missing arguments: either webhook or newname must be defined")
		}

		if yesNo(fmt.Sprintf("You are attempting to update Microsoft Teams notification '%s', are you sure?", name)) {
			utoken := lUser.UserConfig.Grant.AccessToken
			lc := lclient.New(
				fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
				lagoonCLIVersion,
				lContext.ContextConfig.Version,
				&utoken,
				debug)

			notification := &schema.UpdateNotificationMicrosoftTeamsInput{
				Name:  name,
				Patch: patch,
			}
			result, err := lagoon.UpdateNotificationMicrosoftTeams(context.TODO(), notification, lc)
			if err != nil {
				return err
			}
			data := []output.Data{
				[]string{
					returnNonEmptyString(fmt.Sprintf("%v", result.ID)),
					returnNonEmptyString(fmt.Sprintf("%v", result.Name)),
					returnNonEmptyString(fmt.Sprintf("%v", result.Webhook)),
				},
			}
			r := output.RenderOutput(output.Table{
				Header: []string{
					"ID",
					"Name",
					"Webhook",
				},
				Data: data,
			}, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		}
		return nil
	},
}

func init() {
	addNotificationMicrosoftTeamsCmd.Flags().StringP("name", "n", "", "The name of the notification")
	addNotificationMicrosoftTeamsCmd.Flags().StringP("webhook", "w", "", "The webhook URL of the notification")
	addNotificationMicrosoftTeamsCmd.Flags().Uint("organization-id", 0, "ID of the Organization")
	addProjectNotificationMicrosoftTeamsCmd.Flags().StringP("name", "n", "", "The name of the notification")
	deleteProjectMicrosoftTeamsNotificationCmd.Flags().StringP("name", "n", "", "The name of the notification")
	deleteMicrosoftTeamsNotificationCmd.Flags().StringP("name", "n", "", "The name of the notification")
	updateMicrosoftTeamsNotificationCmd.Flags().StringP("name", "n", "", "The name of the notification")
	updateMicrosoftTeamsNotificationCmd.Flags().StringP("newname", "N", "", "The name of the notification")
	updateMicrosoftTeamsNotificationCmd.Flags().StringP("webhook", "w", "", "The webhook URL of the notification")
}
