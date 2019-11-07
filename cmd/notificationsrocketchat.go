package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/lagoon/projects"
	"github.com/amazeeio/lagoon-cli/output"
	"github.com/spf13/cobra"
)

var listRocketChatsCmd = &cobra.Command{
	Use:     "rocketchat",
	Aliases: []string{"r"},
	Short:   "Rocketchat details about a project (alias: r)",
	Run: func(cmd *cobra.Command, args []string) {
		var returnedJSON []byte
		var err error
		if listAllProjects {
			returnedJSON, err = projects.ListAllRocketChats()
			if err != nil {
				output.RenderError(err.Error(), outputOptions)
				os.Exit(1)
			}
		} else {
			notificationFlags := parseNotificationFlags(*cmd.Flags())
			if notificationFlags.Project == "" {
				fmt.Println("Not enough arguments. Requires: project name")
				cmd.Help()
				os.Exit(1)
			}

			returnedJSON, err = projects.ListProjectRocketChats(notificationFlags.Project)
			if err != nil {
				output.RenderError(err.Error(), outputOptions)
				os.Exit(1)
			}
		}
		var dataMain output.Table
		err = json.Unmarshal([]byte(returnedJSON), &dataMain)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		if len(dataMain.Data) == 0 {
			output.RenderError("no data returned", outputOptions)
			os.Exit(1)
		}
		output.RenderOutput(dataMain, outputOptions)
	},
}

var addRocketChatNotificationCmd = &cobra.Command{
	Use:     "rocketchat",
	Aliases: []string{"r"},
	Short:   "Add a new rocketchat notification",
	Run: func(cmd *cobra.Command, args []string) {
		notificationFlags := parseNotificationFlags(*cmd.Flags())
		if notificationFlags.NotificationName == "" || notificationFlags.NotificationChannel == "" || notificationFlags.NotificationWebhook == "" {
			fmt.Println("Not enough arguments. Requires: notifcation name, channel, and webhook url")
			cmd.Help()
			os.Exit(1)
		}

		addResult, err := projects.AddRocketChatNotification(notificationFlags.NotificationName, notificationFlags.NotificationChannel, notificationFlags.NotificationWebhook)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		var resultMap map[string]interface{}
		err = json.Unmarshal([]byte(addResult), &resultMap)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
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
	Run: func(cmd *cobra.Command, args []string) {
		notificationFlags := parseNotificationFlags(*cmd.Flags())
		if notificationFlags.Project == "" || notificationFlags.NotificationName == "" {
			fmt.Println("Not enough arguments. Requires: project name and notifcation name")
			cmd.Help()
			os.Exit(1)
		}
		addResult, err := projects.AddRocketChatNotificationToProject(notificationFlags.Project, notificationFlags.NotificationName)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		var resultMap map[string]interface{}
		err = json.Unmarshal([]byte(addResult), &resultMap)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
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
			fmt.Println("Not enough arguments. Requires: project name and notifcation name")
			cmd.Help()
			os.Exit(1)
		}
		deleteResult, err := projects.DeleteRocketChatNotificationFromProject(notificationFlags.Project, notificationFlags.NotificationName)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		var addedProject api.NotificationSlack
		err = json.Unmarshal([]byte(deleteResult), &addedProject)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		resultData := output.Result{
			Result: "success",
		}
		output.RenderResult(resultData, outputOptions)
	},
}
var deleteRocketChatNotificationCmd = &cobra.Command{
	Use:     "rocketchat",
	Aliases: []string{"r"},
	Short:   "Delete a rocketchat notification from lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		notificationFlags := parseNotificationFlags(*cmd.Flags())
		if notificationFlags.NotificationName == "" {
			fmt.Println("Not enough arguments. Requires: notifcation name")
			cmd.Help()
			os.Exit(1)
		}
		deleteResult, err := projects.DeleteRocketChatNotification(notificationFlags.NotificationName)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		resultData := output.Result{
			Result: string(deleteResult),
		}
		output.RenderResult(resultData, outputOptions)
	},
}

var updateRocketChatNotificationCmd = &cobra.Command{
	Use:     "rocketchat",
	Aliases: []string{"r"},
	Short:   "Update an existing rocketchat notification",
	Run: func(cmd *cobra.Command, args []string) {
		notificationFlags := parseNotificationFlags(*cmd.Flags())
		if notificationFlags.NotificationName == "" {
			fmt.Println("Not enough arguments. Requires: current notifcation name")
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
			if err != nil {
				output.RenderError(err.Error(), outputOptions)
				os.Exit(1)
			}
			jsonPatch = string(jsonPatchBytes)
		}
		updateResult, err := projects.UpdateRocketChatNotification(notificationFlags.NotificationOldName, jsonPatch)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		var resultMap map[string]interface{}
		err = json.Unmarshal([]byte(updateResult), &resultMap)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
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

	updateRocketChatNotificationCmd.Flags().StringVarP(&notificationName, "name", "n", "", "The current name of the notification")
	updateRocketChatNotificationCmd.Flags().StringVarP(&notificationNewName, "newname", "N", "", "The name of the notification")
	updateRocketChatNotificationCmd.Flags().StringVarP(&notificationWebhook, "webhook", "w", "", "The webhook URL of the notification")
	updateRocketChatNotificationCmd.Flags().StringVarP(&notificationChannel, "channel", "c", "", "The channel for the notification")

	updateRocketChatNotificationCmd.Flags().StringVarP(&jsonPatch, "json", "j", "", "JSON string to patch")
}
