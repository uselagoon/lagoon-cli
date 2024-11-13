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

var addNotificationSlackCmd = &cobra.Command{
	Use:     "slack",
	Aliases: []string{"s"},
	Short:   "Add a new Slack notification",
	Long: `Add a new Slack notification
This command is used to set up a new Slack notification in Lagoon. This requires information to talk to Slack like the webhook URL and the name of the channel.
It does not configure a project to send notifications to Slack though, you need to use project-slack for that.`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
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
		if yesNo(fmt.Sprintf("You are attempting to create an Slack notification '%s' with webhook '%s' channel '%s', are you sure?", name, webhook, channel)) {
			current := lagoonCLIConfig.Current
			token := lagoonCLIConfig.Lagoons[current].Token
			lc := lclient.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIVersion,
				lagoonCLIConfig.Lagoons[current].Version,
				&token,
				debug)

			notification := schema.AddNotificationSlackInput{
				Name:         name,
				Webhook:      webhook,
				Channel:      channel,
				Organization: &organizationID,
			}

			result, err := lagoon.AddNotificationSlack(context.TODO(), &notification, lc)
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

var addProjectNotificationSlackCmd = &cobra.Command{
	Use:     "project-slack",
	Aliases: []string{"ps"},
	Short:   "Add a Slack notification to a project",
	Long: `Add a Slack notification to a project
This command is used to add an existing Slack notification in Lagoon to a project.`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
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
		if yesNo(fmt.Sprintf("You are attempting to add Slack notification '%s' to project '%s', are you sure?", name, cmdProjectName)) {
			current := lagoonCLIConfig.Current
			token := lagoonCLIConfig.Lagoons[current].Token
			lc := lclient.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIVersion,
				lagoonCLIConfig.Lagoons[current].Version,
				&token,
				debug)
			notification := &schema.AddNotificationToProjectInput{
				NotificationType: schema.SlackNotification,
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

var listProjectSlacksCmd = &cobra.Command{
	Use:     "project-slack",
	Aliases: []string{"ps"},
	Short:   "List Slacks details about a project (alias: ps)",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Project name", cmdProjectName); err != nil {
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

		result, err := lagoon.GetProjectNotificationSlack(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		if len(result.Name) == 0 {
			return handleNilResults("No project found for '%s'\n", cmd, cmdProjectName)
		} else if len(result.Notifications.Slack) == 0 {
			return handleNilResults("No slack notificatons found for project: '%s'\n", cmd, cmdProjectName)
		}

		data := []output.Data{}
		if result.Notifications != nil {
			for _, notification := range result.Notifications.Slack {
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

var listAllSlacksCmd = &cobra.Command{
	Use:     "slack",
	Aliases: []string{"s"},
	Short:   "List all Slacks notification details (alias: s)",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
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
		result, err := lagoon.GetAllNotificationSlack(context.TODO(), lc)
		if err != nil {
			return err
		}
		data := []output.Data{}
		for _, res := range *result {
			b, _ := json.Marshal(res.Notifications.Slack)
			if string(b) != "null" {
				for _, notif := range res.Notifications.Slack {
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

var deleteProjectSlackNotificationCmd = &cobra.Command{
	Use:     "project-slack",
	Aliases: []string{"ps"},
	Short:   "Delete a Slack notification from a project",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
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

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)

		project, err := lagoon.GetProjectByName(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		if project.Name == "" {
			return handleNilResults("No project found for '%s'\n", cmd, cmdProjectName)
		}

		if yesNo(fmt.Sprintf("You are attempting to delete Slack notification '%s' from project '%s', are you sure?", name, cmdProjectName)) {
			notification := &schema.RemoveNotificationFromProjectInput{
				NotificationType: schema.SlackNotification,
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

var deleteSlackNotificationCmd = &cobra.Command{
	Use:     "slack",
	Aliases: []string{"s"},
	Short:   "Delete a Slack notification from Lagoon",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
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
		if yesNo(fmt.Sprintf("You are attempting to delete Slack notification '%s', are you sure?", name)) {
			current := lagoonCLIConfig.Current
			token := lagoonCLIConfig.Lagoons[current].Token
			lc := lclient.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIVersion,
				lagoonCLIConfig.Lagoons[current].Version,
				&token,
				debug)
			result, err := lagoon.DeleteNotificationSlack(context.TODO(), name, lc)
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

var updateSlackNotificationCmd = &cobra.Command{
	Use:     "slack",
	Aliases: []string{"s"},
	Short:   "Update an existing Slack notification",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
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
		patch := schema.UpdateNotificationSlackPatchInput{
			Name:    nullStrCheck(newname),
			Webhook: nullStrCheck(webhook),
			Channel: nullStrCheck(channel),
		}
		if patch == (schema.UpdateNotificationSlackPatchInput{}) {
			return fmt.Errorf("missing arguments: either channel, webhook, or newname must be defined")
		}

		if yesNo(fmt.Sprintf("You are attempting to update Slack notification '%s', are you sure?", name)) {
			current := lagoonCLIConfig.Current
			token := lagoonCLIConfig.Lagoons[current].Token
			lc := lclient.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIVersion,
				lagoonCLIConfig.Lagoons[current].Version,
				&token,
				debug)

			notification := &schema.UpdateNotificationSlackInput{
				Name:  name,
				Patch: patch,
			}
			result, err := lagoon.UpdateNotificationSlack(context.TODO(), notification, lc)
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
	addNotificationSlackCmd.Flags().StringP("name", "n", "", "The name of the notification")
	addNotificationSlackCmd.Flags().StringP("webhook", "w", "", "The webhook URL of the notification")
	addNotificationSlackCmd.Flags().StringP("channel", "c", "", "The channel for the notification")
	addNotificationSlackCmd.Flags().Uint("organization-id", 0, "ID of the Organization")
	addProjectNotificationSlackCmd.Flags().StringP("name", "n", "", "The name of the notification")
	deleteProjectSlackNotificationCmd.Flags().StringP("name", "n", "", "The name of the notification")
	deleteSlackNotificationCmd.Flags().StringP("name", "n", "", "The name of the notification")
	updateSlackNotificationCmd.Flags().StringP("name", "n", "", "The current name of the notification")
	updateSlackNotificationCmd.Flags().StringP("newname", "N", "", "The name of the notification")
	updateSlackNotificationCmd.Flags().StringP("webhook", "w", "", "The webhook URL of the notification")
	updateSlackNotificationCmd.Flags().StringP("channel", "c", "", "The channel for the notification")

}
