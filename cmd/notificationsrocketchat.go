package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/pkg/api"
	"github.com/uselagoon/lagoon-cli/pkg/output"
)

var listRocketChatsCmd = &cobra.Command{
	Use:     "rocketchat",
	Aliases: []string{"r"},
	Short:   "List Rocket.Chat details about a project (alias: r)",
	Run: func(cmd *cobra.Command, args []string) {
		var returnedJSON []byte
		var err error
		var notificationFlags NotificationFlags
		if listAllProjects {
			returnedJSON, err = pClient.ListAllRocketChats()
			handleError(err)
		} else {
			notificationFlags = parseNotificationFlags(*cmd.Flags())
			if notificationFlags.Project == "" {
				fmt.Println("Missing arguments: Project name is not defined")
				cmd.Help()
				os.Exit(1)
			}
			returnedJSON, err = pClient.ListProjectRocketChats(notificationFlags.Project)
			handleError(err)
		}
		var dataMain output.Table
		err = json.Unmarshal([]byte(returnedJSON), &dataMain)
		handleError(err)
		if len(dataMain.Data) == 0 {
			if listAllProjects {
				output.RenderInfo("No notifications for Rocket.Chat", outputOptions)
			} else {
				output.RenderInfo(fmt.Sprintf("No notifications for Rocket.Chat in project '%s'", notificationFlags.Project), outputOptions)
			}
		}
		output.RenderOutput(dataMain, outputOptions)
	},
}

var addRocketChatNotificationCmd = &cobra.Command{
	Use:     "rocketchat",
	Aliases: []string{"r"},
	Short:   "Add a new Rocket.Chat notification",
	Long: `Add a new Rocket.Chat notification
This command is used to set up a new Rocket.Chat notification in Lagoon. This requires information to talk to Rocket.Chat like the webhook URL and the name of the channel.
It does not configure a project to send notifications to Rocket.Chat though, you need to use project-rocketchat for that.`,
	Run: func(cmd *cobra.Command, args []string) {
		notificationFlags := parseNotificationFlags(*cmd.Flags())
		if notificationFlags.NotificationName == "" || notificationFlags.NotificationChannel == "" || notificationFlags.NotificationWebhook == "" {
			fmt.Println("Missing arguments: Notification name, channel, or webhook url are not defined")
			cmd.Help()
			os.Exit(1)
		}
		addResult, err := pClient.AddRocketChatNotification(notificationFlags.NotificationName, notificationFlags.NotificationChannel, notificationFlags.NotificationWebhook)
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

var addProjectRocketChatNotificationCmd = &cobra.Command{
	Use:     "project-rocketchat",
	Aliases: []string{"pr"},
	Short:   "Add a Rocket.Chat notification to a project",
	Long: `Add a Rocket.Chat notification to a project
This command is used to add an existing Rocket.Chat notification in Lagoon to a project.`,
	Run: func(cmd *cobra.Command, args []string) {
		notificationFlags := parseNotificationFlags(*cmd.Flags())
		if notificationFlags.Project == "" || notificationFlags.NotificationName == "" {
			fmt.Println("Missing arguments: Project name or notification name are not defined")
			cmd.Help()
			os.Exit(1)
		}
		addResult, err := pClient.AddRocketChatNotificationToProject(notificationFlags.Project, notificationFlags.NotificationName)
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

var deleteProjectRocketChatNotificationCmd = &cobra.Command{
	Use:     "project-rocketchat",
	Aliases: []string{"pr"},
	Short:   "Delete a Rocket.Chat notification from a project",
	Run: func(cmd *cobra.Command, args []string) {
		notificationFlags := parseNotificationFlags(*cmd.Flags())
		if notificationFlags.Project == "" || notificationFlags.NotificationName == "" {
			fmt.Println("Missing arguments: Project name or notification name are not defined")
			cmd.Help()
			os.Exit(1)
		}
		if yesNo(fmt.Sprintf("You are attempting to delete notification '%s' from project '%s', are you sure?", notificationFlags.NotificationName, notificationFlags.Project)) {
			deleteResult, err := pClient.DeleteRocketChatNotificationFromProject(notificationFlags.Project, notificationFlags.NotificationName)
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
var deleteRocketChatNotificationCmd = &cobra.Command{
	Use:     "rocketchat",
	Aliases: []string{"r"},
	Short:   "Delete a Rocket.Chat notification from Lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		notificationFlags := parseNotificationFlags(*cmd.Flags())
		if notificationFlags.NotificationName == "" {
			fmt.Println("Missing arguments: notification name is not defined")
			cmd.Help()
			os.Exit(1)
		}
		if yesNo(fmt.Sprintf("You are attempting to delete notification '%s' from Lagoon, are you sure?", notificationFlags.NotificationName)) {
			deleteResult, err := pClient.DeleteRocketChatNotification(notificationFlags.NotificationName)
			handleError(err)
			resultData := output.Result{
				Result: string(deleteResult),
			}
			output.RenderResult(resultData, outputOptions)
		}
	},
}

var updateRocketChatNotificationCmd = &cobra.Command{
	Use:     "rocketchat",
	Aliases: []string{"r"},
	Short:   "Update an existing Rocket.Chat notification",
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
		updateResult, err := pClient.UpdateRocketChatNotification(notificationFlags.NotificationOldName, jsonPatch)
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
	addRocketChatNotificationCmd.Flags().StringVarP(&notificationName, "name", "n", "", "The name of the notification")
	addRocketChatNotificationCmd.Flags().StringVarP(&notificationWebhook, "webhook", "w", "", "The webhook URL of the notification")
	addRocketChatNotificationCmd.Flags().StringVarP(&notificationChannel, "channel", "c", "", "The channel for the notification")

	addProjectRocketChatNotificationCmd.Flags().StringVarP(&notificationName, "name", "n", "", "The name of the notification")

	deleteProjectRocketChatNotificationCmd.Flags().StringVarP(&notificationName, "name", "n", "", "The name of the notification")
	deleteRocketChatNotificationCmd.Flags().StringVarP(&notificationName, "name", "n", "", "The name of the notification")

	updateRocketChatNotificationCmd.Flags().StringVarP(&notificationName, "name", "n", "", "The current name of the notification")
	updateRocketChatNotificationCmd.Flags().StringVarP(&notificationNewName, "newname", "N", "", "The name of the notification")
	updateRocketChatNotificationCmd.Flags().StringVarP(&notificationWebhook, "webhook", "w", "", "The webhook URL of the notification")
	updateRocketChatNotificationCmd.Flags().StringVarP(&notificationChannel, "channel", "c", "", "The channel for the notification")

	updateRocketChatNotificationCmd.Flags().StringVarP(&jsonPatch, "json", "j", "", "JSON string to patch")
}
