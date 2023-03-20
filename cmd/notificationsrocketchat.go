package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/internal/lagoon"
	"github.com/uselagoon/lagoon-cli/internal/lagoon/client"
	"github.com/uselagoon/lagoon-cli/internal/schema"
	"github.com/uselagoon/lagoon-cli/pkg/api"
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
		if name == "" || channel == "" || webhook == "" {
			return fmt.Errorf("Missing arguments: name, webhook, or email is not defined")
		}
		if yesNo(fmt.Sprintf("You are attempting to create an RocketChat notification '%s' with webhook '%s' channel '%s', are you sure?", name, webhook, channel)) {
			current := lagoonCLIConfig.Current
			lc := client.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIConfig.Lagoons[current].Token,
				lagoonCLIConfig.Lagoons[current].Version,
				lagoonCLIVersion,
				debug)
			notification := &schema.AddNotificationRocketChatInput{
				Name:    name,
				Webhook: webhook,
			}
			result, err := lagoon.AddNotificationRocketChat(context.TODO(), notification, lc)
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
			output.RenderOutput(output.Table{
				Header: []string{
					"ID",
					"Name",
					"Webhook",
					"Channel",
				},
				Data: data,
			}, outputOptions)
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
		if yesNo(fmt.Sprintf("You are attempting to add RocketChat notification '%s' to project '%s', are you sure?", name, cmdProjectName)) {
			current := lagoonCLIConfig.Current
			lc := client.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIConfig.Lagoons[current].Token,
				lagoonCLIConfig.Lagoons[current].Version,
				lagoonCLIVersion,
				debug)
			notification := &schema.AddNotificationToProjectInput{
				NotificationType: api.RocketChatNotification,
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

var listProjectRocketChatsCmd = &cobra.Command{
	Use:     "project-rocketchat",
	Aliases: []string{"pr"},
	Short:   "List RocketChats details about a project (alias: pr)",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		if cmdProjectName == "" {
			return fmt.Errorf("Missing arguments: project name is not defined")
		}
		current := lagoonCLIConfig.Current
		lc := client.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIConfig.Lagoons[current].Token,
			lagoonCLIConfig.Lagoons[current].Version,
			lagoonCLIVersion,
			debug)
		result, err := lagoon.GetProjectNotificationRocketChat(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		data := []output.Data{}
		for _, notification := range result.Notifications.RocketChat {
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%v", notification.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", notification.Webhook)),
				returnNonEmptyString(fmt.Sprintf("%v", notification.Channel)),
			})
		}
		output.RenderOutput(output.Table{
			Header: []string{
				"Name",
				"Webhook",
				"Channel",
			},
			Data: data,
		}, outputOptions)
		return nil
	},
}

var listAllRocketChatsCmd = &cobra.Command{
	Use:     "rocketchat",
	Aliases: []string{"r"},
	Short:   "List all RocketChats notification details (alias: r)",
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
		output.RenderOutput(output.Table{
			Header: []string{
				"Project",
				"Name",
				"Webhook",
				"Channel",
			},
			Data: data,
		}, outputOptions)
		return nil
	},
}

var deleteProjectRocketChatNotificationCmd = &cobra.Command{
	Use:     "project-rocketchat",
	Aliases: []string{"pr"},
	Short:   "Delete a RocketChat notification from a project",
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
		if yesNo(fmt.Sprintf("You are attempting to delete RocketChat notification '%s' from project '%s', are you sure?", name, cmdProjectName)) {
			current := lagoonCLIConfig.Current
			lc := client.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIConfig.Lagoons[current].Token,
				lagoonCLIConfig.Lagoons[current].Version,
				lagoonCLIVersion,
				debug)
			notification := &schema.RemoveNotificationFromProjectInput{
				NotificationType: api.RocketChatNotification,
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

var deleteRocketChatNotificationCmd = &cobra.Command{
	Use:     "rocketchat",
	Aliases: []string{"r"},
	Short:   "Delete a RocketChat notification from Lagoon",
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
		if yesNo(fmt.Sprintf("You are attempting to delete RocketChat notification '%s', are you sure?", name)) {
			current := lagoonCLIConfig.Current
			lc := client.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIConfig.Lagoons[current].Token,
				lagoonCLIConfig.Lagoons[current].Version,
				lagoonCLIVersion,
				debug)
			result, err := lagoon.DeleteNotificationRocketChat(context.TODO(), name, lc)
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

var updateRocketChatNotificationCmd = &cobra.Command{
	Use:     "rocketchat",
	Aliases: []string{"r"},
	Short:   "Update an existing RocketChat notification",
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
		if name == "" {
			return fmt.Errorf("Missing arguments: notification name is not defined")
		}
		patch := schema.AddNotificationRocketChatInput{
			Name:    newname,
			Webhook: webhook,
			Channel: channel,
		}
		b1, _ := json.Marshal(patch)
		if bytes.Equal(b1, []byte("{}")) {
			return fmt.Errorf("Missing arguments: either channel, webhook, or newname must be defined")
		}

		if yesNo(fmt.Sprintf("You are attempting to update RocketChat notification '%s', are you sure?", name)) {
			current := lagoonCLIConfig.Current
			lc := client.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIConfig.Lagoons[current].Token,
				lagoonCLIConfig.Lagoons[current].Version,
				lagoonCLIVersion,
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
			output.RenderOutput(output.Table{
				Header: []string{
					"ID",
					"Name",
					"Webhook",
					"Channel",
				},
				Data: data,
			}, outputOptions)
		}
		return nil
	},
}

func init() {
	addNotificationRocketchatCmd.Flags().StringP("name", "n", "", "The name of the notification")
	addNotificationRocketchatCmd.Flags().StringP("webhook", "w", "", "The webhook URL of the notification")
	addNotificationRocketchatCmd.Flags().StringP("channel", "c", "", "The channel for the notification")
	addProjectNotificationRocketChatCmd.Flags().StringP("name", "n", "", "The name of the notification")
	deleteProjectRocketChatNotificationCmd.Flags().StringP("name", "n", "", "The name of the notification")
	deleteRocketChatNotificationCmd.Flags().StringP("name", "n", "", "The name of the notification")
	updateRocketChatNotificationCmd.Flags().StringP("name", "n", "", "The current name of the notification")
	updateRocketChatNotificationCmd.Flags().StringP("newname", "N", "", "The new name of the notification (if required)")
	updateRocketChatNotificationCmd.Flags().StringP("webhook", "w", "", "The webhook URL of the notification")
	updateRocketChatNotificationCmd.Flags().StringP("channel", "c", "", "The channel for the notification")
}
