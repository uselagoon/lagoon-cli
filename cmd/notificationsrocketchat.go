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

var addNotificationRocketchatCmd = &cobra.Command{
	Use:     "rocketchat",
	Aliases: []string{"r"},
	Short:   "Add a new RocketChat notification",
	Long: `Add a new RocketChat notification
This command is used to set up a new RocketChat notification in Lagoon. This requires information to talk to RocketChat like the webhook URL and the name of the channel.
It does not configure a project to send notifications to RocketChat though, you need to use project-rocketchat for that.`,
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
		channel, err := cmd.Flags().GetString("channel")
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
		if err := requiredInputCheck("Notification name", name, "Channel", channel, "Webhook", webhook); err != nil {
			return err
		}
		if yesNo(fmt.Sprintf("You are attempting to create an RocketChat notification '%s' with webhook '%s' channel '%s', are you sure?", name, webhook, channel)) {
			utoken := lUser.UserConfig.Grant.AccessToken
			lc := lclient.New(
				fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
				lagoonCLIVersion,
				lContext.ContextConfig.Version,
				&utoken,
				debug)

			notification := schema.AddNotificationRocketChatInput{
				Name:         name,
				Webhook:      webhook,
				Channel:      channel,
				Organization: &organizationID,
			}

			result, err := lagoon.AddNotificationRocketChat(context.TODO(), &notification, lc)
			if err != nil {
				return err
			}
			var data []output.Data
			notificationData := []string{
				returnNonEmptyString(fmt.Sprintf("%v", result.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", result.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", result.Webhook)),
				returnNonEmptyString(fmt.Sprintf("%v", result.Channel)),
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
					"Channel",
					"Organization",
				},
				Data: data,
			}, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		}
		return nil
	},
}

var addProjectNotificationRocketChatCmd = &cobra.Command{
	Use:     "project-rocketchat",
	Aliases: []string{"pr"},
	Short:   "Add a RocketChat notification to a project",
	Long: `Add a RocketChat notification to a project
This command is used to add an existing RocketChat notification in Lagoon to a project.`,
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
		if yesNo(fmt.Sprintf("You are attempting to add RocketChat notification '%s' to project '%s', are you sure?", name, cmdProjectName)) {
			utoken := lUser.UserConfig.Grant.AccessToken
			lc := lclient.New(
				fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
				lagoonCLIVersion,
				lContext.ContextConfig.Version,
				&utoken,
				debug)

			notification := &schema.AddNotificationToProjectInput{
				NotificationType: schema.RocketChatNotification,
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

var listProjectRocketChatsCmd = &cobra.Command{
	Use:     "project-rocketchat",
	Aliases: []string{"pr"},
	Short:   "List RocketChats details about a project (alias: pr)",
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

		result, err := lagoon.GetProjectNotificationRocketChat(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		if len(result.Name) == 0 {
			return handleNilResults("No project found for '%s'\n", cmd, cmdProjectName)
		} else if len(result.Notifications.RocketChat) == 0 {
			return handleNilResults("No RocketChat notifications found for project '%s'\n", cmd, cmdProjectName)
		}

		data := []output.Data{}
		if result.Notifications != nil {
			for _, notification := range result.Notifications.RocketChat {
				data = append(data, []string{
					returnNonEmptyString(fmt.Sprintf("%v", notification.Name)),
					returnNonEmptyString(fmt.Sprintf("%v", notification.Webhook)),
					returnNonEmptyString(fmt.Sprintf("%v", notification.Channel)),
				})
			}
		}
		r := output.RenderOutput(output.Table{
			Header: []string{
				"Name",
				"Webhook",
				"Channel",
			},
			Data: data,
		}, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		return nil
	},
}

var listAllRocketChatsCmd = &cobra.Command{
	Use:     "rocketchat",
	Aliases: []string{"r"},
	Short:   "List all RocketChats notification details (alias: r)",
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
		result, err := lagoon.GetAllNotificationRocketChat(context.TODO(), lc)
		if err != nil {
			return err
		}
		data := []output.Data{}
		for _, res := range *result {
			b, _ := json.Marshal(res.Notifications.RocketChat)
			if string(b) != "null" {
				for _, notif := range res.Notifications.RocketChat {
					data = append(data, []string{
						returnNonEmptyString(fmt.Sprintf("%v", res.Name)),
						returnNonEmptyString(fmt.Sprintf("%v", notif.Name)),
						returnNonEmptyString(fmt.Sprintf("%v", notif.Webhook)),
						returnNonEmptyString(fmt.Sprintf("%v", notif.Channel)),
					})
				}
			}
		}
		r := output.RenderOutput(output.Table{
			Header: []string{
				"Project",
				"Name",
				"Webhook",
				"Channel",
			},
			Data: data,
		}, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		return nil
	},
}

var deleteProjectRocketChatNotificationCmd = &cobra.Command{
	Use:     "project-rocketchat",
	Aliases: []string{"pr"},
	Short:   "Delete a RocketChat notification from a project",
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

		if yesNo(fmt.Sprintf("You are attempting to delete RocketChat notification '%s' from project '%s', are you sure?", name, cmdProjectName)) {
			notification := &schema.RemoveNotificationFromProjectInput{
				NotificationType: schema.RocketChatNotification,
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

var deleteRocketChatNotificationCmd = &cobra.Command{
	Use:     "rocketchat",
	Aliases: []string{"r"},
	Short:   "Delete a RocketChat notification from Lagoon",
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
		if yesNo(fmt.Sprintf("You are attempting to delete RocketChat notification '%s', are you sure?", name)) {
			utoken := lUser.UserConfig.Grant.AccessToken
			lc := lclient.New(
				fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
				lagoonCLIVersion,
				lContext.ContextConfig.Version,
				&utoken,
				debug)
			result, err := lagoon.DeleteNotificationRocketChat(context.TODO(), name, lc)
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

var updateRocketChatNotificationCmd = &cobra.Command{
	Use:     "rocketchat",
	Aliases: []string{"r"},
	Short:   "Update an existing RocketChat notification",
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
		channel, err := cmd.Flags().GetString("channel")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Notification name", name); err != nil {
			return err
		}
		patch := schema.UpdateNotificationRocketChatPatchInput{
			Name:    nullStrCheck(newname),
			Webhook: nullStrCheck(webhook),
			Channel: nullStrCheck(channel),
		}
		if patch == (schema.UpdateNotificationRocketChatPatchInput{}) {
			return fmt.Errorf("missing arguments: either channel, webhook, or newname must be defined")
		}

		if yesNo(fmt.Sprintf("You are attempting to update RocketChat notification '%s', are you sure?", name)) {
			utoken := lUser.UserConfig.Grant.AccessToken
			lc := lclient.New(
				fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
				lagoonCLIVersion,
				lContext.ContextConfig.Version,
				&utoken,
				debug)

			notification := &schema.UpdateNotificationRocketChatInput{
				Name:  name,
				Patch: patch,
			}
			result, err := lagoon.UpdateNotificationRocketChat(context.TODO(), notification, lc)
			if err != nil {
				return err
			}
			data := []output.Data{
				[]string{
					returnNonEmptyString(fmt.Sprintf("%v", result.ID)),
					returnNonEmptyString(fmt.Sprintf("%v", result.Name)),
					returnNonEmptyString(fmt.Sprintf("%v", result.Webhook)),
					returnNonEmptyString(fmt.Sprintf("%v", result.Channel)),
				},
			}
			r := output.RenderOutput(output.Table{
				Header: []string{
					"ID",
					"Name",
					"Webhook",
					"Channel",
				},
				Data: data,
			}, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		}
		return nil
	},
}

func init() {
	addNotificationRocketchatCmd.Flags().StringP("name", "n", "", "The name of the notification")
	addNotificationRocketchatCmd.Flags().StringP("webhook", "w", "", "The webhook URL of the notification")
	addNotificationRocketchatCmd.Flags().StringP("channel", "c", "", "The channel for the notification")
	addNotificationRocketchatCmd.Flags().Uint("organization-id", 0, "ID of the Organization")
	addProjectNotificationRocketChatCmd.Flags().StringP("name", "n", "", "The name of the notification")
	deleteProjectRocketChatNotificationCmd.Flags().StringP("name", "n", "", "The name of the notification")
	deleteRocketChatNotificationCmd.Flags().StringP("name", "n", "", "The name of the notification")
	updateRocketChatNotificationCmd.Flags().StringP("name", "n", "", "The current name of the notification")
	updateRocketChatNotificationCmd.Flags().StringP("newname", "N", "", "The new name of the notification (if required)")
	updateRocketChatNotificationCmd.Flags().StringP("webhook", "w", "", "The webhook URL of the notification")
	updateRocketChatNotificationCmd.Flags().StringP("channel", "c", "", "The channel for the notification")
}
