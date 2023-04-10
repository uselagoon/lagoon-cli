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

var addNotificationMicrosoftTeamsCmd = &cobra.Command{
	Use:   "microsoftteams",
	Short: "Add a new Microsoft Teams notification",
	Long: `Add a new Microsoft Teams notification
This command is used to set up a new Microsoft Teams notification in Lagoon. This requires information to talk to the webhook like the webhook URL.
It does not configure a project to send notifications to Microsoft Teams though, you need to use project-microsoftteams for that.`,
	Aliases: []string{"m"},
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
		if name == "" || webhook == "" {
			return fmt.Errorf("Missing arguments: name or webhook is not defined")
		}
		if yesNo(fmt.Sprintf("You are attempting to create a Microsoft Teams notification '%s' with webhook url '%s', are you sure?", name, webhook)) {
			current := lagoonCLIConfig.Current
			lc := client.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIConfig.Lagoons[current].Token,
				lagoonCLIConfig.Lagoons[current].Version,
				lagoonCLIVersion,
				debug)
			notification := &schema.AddNotificationMicrosoftTeamsInput{
				Name:    name,
				Webhook: webhook,
			}
			result, err := lagoon.AddNotificationMicrosoftTeams(context.TODO(), notification, lc)
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

var addProjectNotificationMicrosoftTeamsCmd = &cobra.Command{
	Use:     "project-microsoftteams",
	Aliases: []string{"pm"},
	Short:   "Add a Microsoft Teams notification to a project",
	Long: `Add a Microsoft Teams notification to a project
This command is used to add an existing Microsoft Teams notification in Lagoon to a project.`,
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
		if yesNo(fmt.Sprintf("You are attempting to add Microsoft Teams notification '%s' to project '%s', are you sure?", name, cmdProjectName)) {
			current := lagoonCLIConfig.Current
			lc := client.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIConfig.Lagoons[current].Token,
				lagoonCLIConfig.Lagoons[current].Version,
				lagoonCLIVersion,
				debug)
			notification := &schema.AddNotificationToProjectInput{
				NotificationType: api.MicrosoftTeamsNotification,
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

var listProjectMicrosoftTeamsCmd = &cobra.Command{
	Use:     "project-microsoftteams",
	Aliases: []string{"pm"},
	Short:   "List Microsoft Teams details about a project (alias: pm)",
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
		result, err := lagoon.GetProjectNotificationMicrosoftTeams(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		data := []output.Data{}
		for _, notification := range result.Notifications.MicrosoftTeams {
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%v", notification.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", notification.Webhook)),
			})
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

var listAllMicrosoftTeamsCmd = &cobra.Command{
	Use:     "microsoftteams",
	Aliases: []string{"m"},
	Short:   "List all Microsoft Teams notification details (alias: m)",
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

var deleteProjectMicrosoftTeamsNotificationCmd = &cobra.Command{
	Use:     "project-microsoftteams",
	Aliases: []string{"pr"},
	Short:   "Delete a Microsoft Teams notification from a project",
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
		if yesNo(fmt.Sprintf("You are attempting to delete Microsoft Teams notification '%s' from project '%s', are you sure?", name, cmdProjectName)) {
			current := lagoonCLIConfig.Current
			lc := client.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIConfig.Lagoons[current].Token,
				lagoonCLIConfig.Lagoons[current].Version,
				lagoonCLIVersion,
				debug)
			notification := &schema.RemoveNotificationFromProjectInput{
				NotificationType: api.MicrosoftTeamsNotification,
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

var deleteMicrosoftTeamsNotificationCmd = &cobra.Command{
	Use:     "microsoftteams",
	Aliases: []string{"m"},
	Short:   "Delete a Microsoft Teams notification from Lagoon",
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
		if yesNo(fmt.Sprintf("You are attempting to delete Microsoft Teams notification '%s', are you sure?", name)) {
			current := lagoonCLIConfig.Current
			lc := client.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIConfig.Lagoons[current].Token,
				lagoonCLIConfig.Lagoons[current].Version,
				lagoonCLIVersion,
				debug)
			result, err := lagoon.DeleteNotificationMicrosoftTeams(context.TODO(), name, lc)
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

var updateMicrosoftTeamsNotificationCmd = &cobra.Command{
	Use:     "microsoftteams",
	Aliases: []string{"m"},
	Short:   "Update an existing Microsoft Teams notification",
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
		patch := schema.AddNotificationMicrosoftTeamsInput{
			Name:    newname,
			Webhook: webhook,
		}
		b1, _ := json.Marshal(patch)
		if bytes.Equal(b1, []byte("{}")) {
			return fmt.Errorf("Missing arguments: either webhook or newname must be defined")
		}

		if yesNo(fmt.Sprintf("You are attempting to update Microsoft Teams notification '%s', are you sure?", name)) {
			current := lagoonCLIConfig.Current
			lc := client.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIConfig.Lagoons[current].Token,
				lagoonCLIConfig.Lagoons[current].Version,
				lagoonCLIVersion,
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
	addNotificationMicrosoftTeamsCmd.Flags().StringP("name", "n", "", "The name of the notification")
	addNotificationMicrosoftTeamsCmd.Flags().StringP("webhook", "w", "", "The webhook URL of the notification")
	addProjectNotificationMicrosoftTeamsCmd.Flags().StringP("name", "n", "", "The name of the notification")
	deleteProjectMicrosoftTeamsNotificationCmd.Flags().StringP("name", "n", "", "The name of the notification")
	deleteMicrosoftTeamsNotificationCmd.Flags().StringP("name", "n", "", "The name of the notification")
	updateMicrosoftTeamsNotificationCmd.Flags().StringP("name", "n", "", "The name of the notification")
	updateMicrosoftTeamsNotificationCmd.Flags().StringP("newname", "N", "", "The name of the notification")
	updateMicrosoftTeamsNotificationCmd.Flags().StringP("webhook", "w", "", "The webhook URL of the notification")
}
