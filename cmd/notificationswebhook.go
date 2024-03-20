package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	l "github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	s "github.com/uselagoon/machinery/api/schema"

	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/internal/lagoon"
	"github.com/uselagoon/lagoon-cli/internal/lagoon/client"
	"github.com/uselagoon/lagoon-cli/internal/schema"
	"github.com/uselagoon/lagoon-cli/pkg/api"
	"github.com/uselagoon/lagoon-cli/pkg/output"
)

var addNotificationWebhookCmd = &cobra.Command{
	Use:   "webhook",
	Short: "Add a new webhook notification",
	Long: `Add a new webhook notification
This command is used to set up a new webhook notification in Lagoon. This requires information to talk to the webhook like the webhook URL.
It does not configure a project to send notifications to webhook though, you need to use project-webhook for that.`,
	Aliases: []string{"w"},
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
		webhook, err := cmd.Flags().GetString("webhook")
		if err != nil {
			return err
		}
		organizationID, err := cmd.Flags().GetUint("organization-id")
		if err != nil {
			return err
		}
		if name == "" || webhook == "" {
			return fmt.Errorf("Missing arguments: name or webhook is not defined")
		}
		if yesNo(fmt.Sprintf("You are attempting to create a webhook notification '%s' with webhook url '%s', are you sure?", name, webhook)) {
			current := lagoonCLIConfig.Current
			token := lagoonCLIConfig.Lagoons[current].Token
			lc := lclient.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIVersion,
				&token,
				debug)

			notification := s.AddNotificationWebhookInput{
				Name:         name,
				Webhook:      webhook,
				Organization: &organizationID,
			}

			result, err := l.AddNotificationWebhook(context.TODO(), &notification, lc)
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
				organization, err := l.GetOrganizationByID(context.TODO(), organizationID, lc)
				if err != nil {
					return err
				}
				notificationData = append(notificationData, fmt.Sprintf("%s", organization.Name))
			} else {
				notificationData = append(notificationData, "-")
			}
			data = append(data, notificationData)
			output.RenderOutput(output.Table{
				Header: []string{
					"ID",
					"Name",
					"Webhook",
					"Organization",
				},
				Data: data,
			}, outputOptions)
		}
		return nil
	},
}

var addProjectNotificationWebhookCmd = &cobra.Command{
	Use:     "project-webhook",
	Aliases: []string{"pe"},
	Short:   "Add a webhook notification to a project",
	Long: `Add a webhook notification to a project
This command is used to add an existing webhook notification in Lagoon to a project.`,
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
		if name == "" || cmdProjectName == "" {
			return fmt.Errorf("Missing arguments: project name or notification name is not defined")
		}
		if yesNo(fmt.Sprintf("You are attempting to add webhook notification '%s' to project '%s', are you sure?", name, cmdProjectName)) {
			current := lagoonCLIConfig.Current
			lc := client.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIConfig.Lagoons[current].Token,
				lagoonCLIConfig.Lagoons[current].Version,
				lagoonCLIVersion,
				debug)
			notification := &schema.AddNotificationToProjectInput{
				NotificationType: api.WebhookNotification,
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
			output.RenderResult(resultData, outputOptions)
		}
		return nil
	},
}

var listProjectWebhooksCmd = &cobra.Command{
	Use:     "project-webhook",
	Aliases: []string{"pw"},
	Short:   "List webhook details about a project (alias: pw)",
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
			&token,
			debug)

		result, err := l.GetProjectNotificationWebhook(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		if len(result.Name) == 0 {
			outputOptions.Error = fmt.Sprintf("No project found for '%s'\n", cmdProjectName)
		} else if len(result.Notifications.Webhook) == 0 {
			outputOptions.Error = fmt.Sprintf("No webhook notificatons found for project: '%s'\n", cmdProjectName)
		}

		data := []output.Data{}
		if result.Notifications != nil {
			for _, notification := range result.Notifications.Webhook {
				data = append(data, []string{
					returnNonEmptyString(fmt.Sprintf("%v", notification.Name)),
					returnNonEmptyString(fmt.Sprintf("%v", notification.Webhook)),
				})
			}
		}
		output.RenderOutput(output.Table{
			Header: []string{
				"Name",
				"Webhook",
			},
			Data: data,
		}, outputOptions)
		return nil
	},
}

var listAllWebhooksCmd = &cobra.Command{
	Use:     "webhook",
	Aliases: []string{"w"},
	Short:   "List all webhook notification details (alias: w)",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		current := lagoonCLIConfig.Current
		lc := client.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIConfig.Lagoons[current].Token,
			lagoonCLIConfig.Lagoons[current].Version,
			lagoonCLIVersion,
			debug)
		result, err := lagoon.GetAllNotificationWebhook(context.TODO(), lc)
		if err != nil {
			return err
		}
		data := []output.Data{}
		for _, res := range *result {
			b, _ := json.Marshal(res.Notifications.Webhook)
			if string(b) != "null" {
				for _, notif := range res.Notifications.Webhook {
					data = append(data, []string{
						returnNonEmptyString(fmt.Sprintf("%v", res.Name)),
						returnNonEmptyString(fmt.Sprintf("%v", notif.Name)),
						returnNonEmptyString(fmt.Sprintf("%v", notif.Webhook)),
					})
				}
			}
		}
		output.RenderOutput(output.Table{
			Header: []string{
				"Project",
				"Name",
				"Webhook",
			},
			Data: data,
		}, outputOptions)
		return nil
	},
}

var deleteProjectWebhookNotificationCmd = &cobra.Command{
	Use:     "project-webhook",
	Aliases: []string{"pr"},
	Short:   "Delete a webhook notification from a project",
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
		if name == "" || cmdProjectName == "" {
			return fmt.Errorf("Missing arguments: project name or notification name is not defined")
		}
		if yesNo(fmt.Sprintf("You are attempting to delete webhook notification '%s' from project '%s', are you sure?", name, cmdProjectName)) {
			current := lagoonCLIConfig.Current
			lc := client.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIConfig.Lagoons[current].Token,
				lagoonCLIConfig.Lagoons[current].Version,
				lagoonCLIVersion,
				debug)
			notification := &schema.RemoveNotificationFromProjectInput{
				NotificationType: api.WebhookNotification,
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
			output.RenderResult(resultData, outputOptions)
		}
		return nil
	},
}

var deleteWebhookNotificationCmd = &cobra.Command{
	Use:     "webhook",
	Aliases: []string{"w"},
	Short:   "Delete a webhook notification from Lagoon",
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
		if name == "" {
			return fmt.Errorf("Missing arguments: notification name is not defined")
		}
		if yesNo(fmt.Sprintf("You are attempting to delete webhook notification '%s', are you sure?", name)) {
			current := lagoonCLIConfig.Current
			lc := client.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIConfig.Lagoons[current].Token,
				lagoonCLIConfig.Lagoons[current].Version,
				lagoonCLIVersion,
				debug)
			result, err := lagoon.DeleteNotificationWebhook(context.TODO(), name, lc)
			if err != nil {
				return err
			}
			resultData := output.Result{
				Result: result.DeleteNotification,
			}
			output.RenderResult(resultData, outputOptions)
		}
		return nil
	},
}

var updateWebhookNotificationCmd = &cobra.Command{
	Use:     "webhook",
	Aliases: []string{"w"},
	Short:   "Update an existing webhook notification",
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
		if name == "" {
			return fmt.Errorf("Missing arguments: notification name is not defined")
		}
		patch := schema.AddNotificationWebhookInput{
			Name:    newname,
			Webhook: webhook,
		}
		b1, _ := json.Marshal(patch)
		if bytes.Equal(b1, []byte("{}")) {
			return fmt.Errorf("Missing arguments: either webhook or newname must be defined")
		}

		if yesNo(fmt.Sprintf("You are attempting to update webhook notification '%s', are you sure?", name)) {
			current := lagoonCLIConfig.Current
			lc := client.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIConfig.Lagoons[current].Token,
				lagoonCLIConfig.Lagoons[current].Version,
				lagoonCLIVersion,
				debug)

			notification := &schema.UpdateNotificationWebhookInput{
				Name:  name,
				Patch: patch,
			}
			result, err := lagoon.UpdateNotificationWebhook(context.TODO(), notification, lc)
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
			output.RenderOutput(output.Table{
				Header: []string{
					"ID",
					"Name",
					"Webhook",
				},
				Data: data,
			}, outputOptions)
		}
		return nil
	},
}

func init() {
	addNotificationWebhookCmd.Flags().StringP("name", "n", "", "The name of the notification")
	addNotificationWebhookCmd.Flags().StringP("webhook", "w", "", "The webhook URL of the notification")
	addNotificationWebhookCmd.Flags().Uint("organization-id", 0, "ID of the Organization")
	addProjectNotificationWebhookCmd.Flags().StringP("name", "n", "", "The name of the notification")
	deleteProjectWebhookNotificationCmd.Flags().StringP("name", "n", "", "The name of the notification")
	deleteWebhookNotificationCmd.Flags().StringP("name", "n", "", "The name of the notification")
	updateWebhookNotificationCmd.Flags().StringP("name", "n", "", "The name of the notification")
	updateWebhookNotificationCmd.Flags().StringP("newname", "N", "", "The name of the notification")
	updateWebhookNotificationCmd.Flags().StringP("webhook", "w", "", "The webhook URL of the notification")
}
