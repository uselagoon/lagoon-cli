package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/pkg/api"
	"github.com/amazeeio/lagoon-cli/pkg/output"
	"github.com/spf13/cobra"
)

var listRocketChatsCmd = &cobra.Command{
	Use:     "rocketchat",
	Aliases: []string{"r"},
	Short:   "List Rocketchat details about a project (alias: r)",
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
				output.RenderInfo("No notifications for RocketChat", outputOptions)
			} else {
				output.RenderInfo(fmt.Sprintf("No notifications for RocketChat in project '%s'", notificationFlags.Project), outputOptions)
			}
		}
		output.RenderOutput(dataMain, outputOptions)
	},
}

var addRocketChatNotificationCmd = &cobra.Command{
	Use:     "rocketchat",
	Aliases: []string{"r"},
	Short:   "Add a new rocketchat notification",
	Long: `Add a new rocketchat notification
This command is used to set up a new rocketchat notification in lagoon. This requires information to talk to rocketchat like the webhook URL and the name of the channel.
It does not configure a project to send notifications to rocketchat though, you need to use project-rocketchat for that.`,
	Run: func(cmd *cobra.Command, args []string) {
		notificationFlags := parseNotificationFlags(*cmd.Flags())
		if notificationFlags.NotificationName == "" || notificationFlags.NotificationChannel == "" || notificationFlags.NotificationWebhook == "" {
			fmt.Println("Missing arguments: Notifcation name, channel, or webhook url are not defined")
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
	Short:   "Add a rocketchat notification to a project",
	Long: `Add a rocketchat notification to a project
This command is used to add an existing rocketchat notification in lagoon to a project.`,
	Run: func(cmd *cobra.Command, args []string) {
		notificationFlags := parseNotificationFlags(*cmd.Flags())
		if notificationFlags.Project == "" || notificationFlags.NotificationName == "" {
			fmt.Println("Missing arguments: Project name or notifcation name are not defined")
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
	Short:   "Delete a rocketchat notification from a project",
	Run: func(cmd *cobra.Command, args []string) {
		notificationFlags := parseNotificationFlags(*cmd.Flags())
		if notificationFlags.Project == "" || notificationFlags.NotificationName == "" {
			fmt.Println("Missing arguments: Project name or notifcation name are not defined")
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
	Short:   "Delete a rocketchat notification from lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		notificationFlags := parseNotificationFlags(*cmd.Flags())
		if notificationFlags.NotificationName == "" {
			fmt.Println("Missing arguments: Notifcation name is not defined")
			cmd.Help()
			os.Exit(1)
		}
		if yesNo(fmt.Sprintf("You are attempting to delete notification '%s' from lagoon, are you sure?", notificationFlags.NotificationName)) {
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
	Short:   "Update an existing rocketchat notification",
	Run: func(cmd *cobra.Command, args []string) {
		notificationFlags := parseNotificationFlags(*cmd.Flags())
		if notificationFlags.NotificationName == "" {
			fmt.Println("Missing arguments: Current notifcation name is not defined")
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
