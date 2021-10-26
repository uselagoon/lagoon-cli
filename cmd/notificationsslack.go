package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/pkg/api"
	"github.com/uselagoon/lagoon-cli/pkg/output"
)

var listSlackCmd = &cobra.Command{
	Use:     "slack",
	Aliases: []string{"s"},
	Short:   "List Slack details about a project (alias: s)",
	Run: func(cmd *cobra.Command, args []string) {
		var returnedJSON []byte
		var err error
		var notificationFlags NotificationFlags
		if listAllProjects {
			returnedJSON, err = pClient.ListAllSlacks()
			handleError(err)
		} else {
			notificationFlags = parseNotificationFlags(*cmd.Flags())
			if notificationFlags.Project == "" {
				fmt.Println("Missing arguments: Project name is not defined")
				cmd.Help()
				os.Exit(1)
			}

			returnedJSON, err = pClient.ListProjectSlacks(notificationFlags.Project)
			handleError(err)
		}
		var dataMain output.Table
		err = json.Unmarshal([]byte(returnedJSON), &dataMain)
		handleError(err)
		if len(dataMain.Data) == 0 {
			if listAllProjects {
				output.RenderInfo("No notifications for Slack", outputOptions)
			} else {
				output.RenderInfo(fmt.Sprintf("No notifications for Slack in project '%s'", notificationFlags.Project), outputOptions)
			}
			os.Exit(0)
		}
		output.RenderOutput(dataMain, outputOptions)

	},
}

var addSlackNotificationCmd = &cobra.Command{
	Use:     "slack",
	Aliases: []string{"s"},
	Short:   "Add a new Slack notification",
	Long: `Add a new Slack notification
This command is used to set up a new Slack notification in Lagoon. This requires information to talk to Slack like the webhook URL and the name of the channel.
It does not configure a project to send notifications to Slack though, you need to use project-slack for that.`,
	Run: func(cmd *cobra.Command, args []string) {
		notificationFlags := parseNotificationFlags(*cmd.Flags())
		if notificationFlags.NotificationName == "" || notificationFlags.NotificationChannel == "" || notificationFlags.NotificationWebhook == "" {
			fmt.Println("Missing arguments: Notification name, channel, or webhook url are not defined")
			cmd.Help()
			os.Exit(1)
		}
		addResult, err := pClient.AddSlackNotification(notificationFlags.NotificationName, notificationFlags.NotificationChannel, notificationFlags.NotificationWebhook)
		handleError(err)
		var resultMap map[string]interface{}
		err = json.Unmarshal([]byte(addResult), &resultMap)
		handleError(err)
		resultData := output.Result{
			Result:     "success",
			ResultData: resultMap,
		}
		output.RenderResult(resultData, outputOptions)
	},
}

var addProjectSlackNotificationCmd = &cobra.Command{
	Use:     "project-slack",
	Aliases: []string{"ps"},
	Short:   "Add a Slack notification to a project",
	Long: `Add a Slack notification to a project
This command is used to add an existing Slack notification in Lagoon to a project.`,
	Run: func(cmd *cobra.Command, args []string) {
		notificationFlags := parseNotificationFlags(*cmd.Flags())
		if notificationFlags.Project == "" || notificationFlags.NotificationName == "" {
			fmt.Println("Missing arguments: Project name or notification name are not defined")
			cmd.Help()
			os.Exit(1)
		}
		addResult, err := pClient.AddSlackNotificationToProject(notificationFlags.Project, notificationFlags.NotificationName)
		handleError(err)
		var resultMap map[string]interface{}
		err = json.Unmarshal([]byte(addResult), &resultMap)
		handleError(err)
		resultData := output.Result{
			Result:     "success",
			ResultData: resultMap,
		}
		output.RenderResult(resultData, outputOptions)
	},
}

var deleteProjectSlackNotificationCmd = &cobra.Command{
	Use:     "project-slack",
	Aliases: []string{"ps"},
	Short:   "Delete a Slack notification from a project",
	Run: func(cmd *cobra.Command, args []string) {
		notificationFlags := parseNotificationFlags(*cmd.Flags())
		if notificationFlags.Project == "" || notificationFlags.NotificationName == "" {
			fmt.Println("Missing arguments: Project name or notification name are not defined")
			cmd.Help()
			os.Exit(1)
		}
		if yesNo(fmt.Sprintf("You are attempting to delete notification '%s' from project '%s', are you sure?", notificationFlags.NotificationName, notificationFlags.Project)) {
			deleteResult, err := pClient.DeleteSlackNotificationFromProject(notificationFlags.Project, notificationFlags.NotificationName)
			handleError(err)
			var addedProject api.NotificationSlack
			err = json.Unmarshal([]byte(deleteResult), &addedProject)
			handleError(err)
			resultData := output.Result{
				Result: "success",
			}
			output.RenderResult(resultData, outputOptions)
		}
	},
}

var deleteSlackNotificationCmd = &cobra.Command{
	Use:     "slack",
	Aliases: []string{"s"},
	Short:   "Delete a Slack notification from Lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		notificationFlags := parseNotificationFlags(*cmd.Flags())
		if notificationFlags.NotificationName == "" {
			fmt.Println("Missing arguments: Notification name is not defined")
			cmd.Help()
			os.Exit(1)
		}
		fmt.Println(fmt.Sprintf("Deleting notification %s", notificationFlags.NotificationName))

		if yesNo(fmt.Sprintf("You are attempting to delete notification '%s' from Lagoon, are you sure?", notificationFlags.NotificationName)) {
			deleteResult, err := pClient.DeleteSlackNotification(notificationFlags.NotificationName)
			handleError(err)
			resultData := output.Result{
				Result: string(deleteResult),
			}
			output.RenderResult(resultData, outputOptions)
		}
	},
}

var updateSlackNotificationCmd = &cobra.Command{
	Use:     "slack",
	Aliases: []string{"s"},
	Short:   "Update an existing Slack notification",
	Run: func(cmd *cobra.Command, args []string) {
		notificationFlags := parseNotificationFlags(*cmd.Flags())
		if notificationFlags.NotificationName == "" {
			fmt.Println("Missing arguments: Current notification name is not defined")
			cmd.Help()
			os.Exit(1)
		}
		oldName := notificationFlags.NotificationName
		// if we have a new name, shuffle around the name
		if notificationFlags.NotificationNewName != "" {
			newName := notificationFlags.NotificationNewName
			notificationFlags.NotificationName = newName
		}
		notificationFlags.NotificationOldName = oldName
		if jsonPatch == "" {
			jsonPatchBytes, err := json.Marshal(notificationFlags)
			handleError(err)
			jsonPatch = string(jsonPatchBytes)
		}
		updateResult, err := pClient.UpdateSlackNotification(notificationFlags.NotificationOldName, jsonPatch)
		handleError(err)
		var resultMap map[string]interface{}
		err = json.Unmarshal([]byte(updateResult), &resultMap)
		handleError(err)
		resultData := output.Result{
			Result:     "success",
			ResultData: resultMap,
		}
		output.RenderResult(resultData, outputOptions)
	},
}

func init() {
	addSlackNotificationCmd.Flags().StringVarP(&notificationName, "name", "n", "", "The name of the notification")
	addSlackNotificationCmd.Flags().StringVarP(&notificationWebhook, "webhook", "w", "", "The webhook URL of the notification")
	addSlackNotificationCmd.Flags().StringVarP(&notificationChannel, "channel", "c", "", "The channel for the notification")

	addProjectSlackNotificationCmd.Flags().StringVarP(&notificationName, "name", "n", "", "The name of the notification")

	deleteProjectSlackNotificationCmd.Flags().StringVarP(&notificationName, "name", "n", "", "The name of the notification")
	deleteSlackNotificationCmd.Flags().StringVarP(&notificationName, "name", "n", "", "The name of the notification")

	updateSlackNotificationCmd.Flags().StringVarP(&notificationName, "name", "n", "", "The current name of the notification")
	updateSlackNotificationCmd.Flags().StringVarP(&notificationNewName, "newname", "N", "", "The name of the notification")
	updateSlackNotificationCmd.Flags().StringVarP(&notificationWebhook, "webhook", "w", "", "The webhook URL of the notification")
	updateSlackNotificationCmd.Flags().StringVarP(&notificationChannel, "channel", "c", "", "The channel for the notification")

	updateSlackNotificationCmd.Flags().StringVarP(&jsonPatch, "json", "j", "", "JSON string to patch")
}
