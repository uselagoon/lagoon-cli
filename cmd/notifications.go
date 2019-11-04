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

var listSlackCmd = &cobra.Command{
	Use:   "slack [project]",
	Short: "Slack details about a project",
	Run: func(cmd *cobra.Command, args []string) {
		var projectName string
		var returnedJSON []byte
		var err error
		if listAllProjects {
			returnedJSON, err = projects.ListAllSlacks()
			if err != nil {
				output.RenderError(err.Error(), outputOptions)
				os.Exit(1)
			}
		} else {
			if len(args) < 1 {
				if cmdProject.Name != "" {
					projectName = cmdProject.Name
				} else {
					fmt.Println("Not enough arguments. Requires: project name")
					cmd.Help()
					os.Exit(1)
				}
			} else {
				projectName = args[0]
			}

			returnedJSON, err = projects.ListAllProjectSlacks(projectName)
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

var listRocketChatsCmd = &cobra.Command{
	Use:   "rocketchat [project]",
	Short: "Rocketchat details about a project",
	Run: func(cmd *cobra.Command, args []string) {
		var projectName string
		var returnedJSON []byte
		var err error
		if listAllProjects {
			returnedJSON, err = projects.ListAllRocketChats()
			if err != nil {
				output.RenderError(err.Error(), outputOptions)
				os.Exit(1)
			}
		} else {
			if len(args) < 1 {
				if cmdProject.Name != "" {
					projectName = cmdProject.Name
				} else {
					fmt.Println("Not enough arguments. Requires: project name")
					cmd.Help()
					os.Exit(1)
				}
			} else {
				projectName = args[0]
			}

			returnedJSON, err = projects.ListAllProjectRocketChats(projectName)
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

var addSlackNotificationCmd = &cobra.Command{
	Use:   "slack [notification name] [channel] [webhook url]",
	Short: "Add a new slack notification",
	Run: func(cmd *cobra.Command, args []string) {
		var notificationName string
		var channel string
		var webhookURL string
		if len(args) < 3 {
			fmt.Println("Not enough arguments. Requires: notifcation name, channel, and webhook url")
			cmd.Help()
			os.Exit(1)
		}
		notificationName = args[0]
		channel = args[1]
		webhookURL = args[2]
		addResult, err := projects.AddSlackNotification(notificationName, channel, webhookURL)
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

var addProjectSlackNotificationCmd = &cobra.Command{
	Use:   "project-slack [project name] [notification name]",
	Short: "Add a slack notification to a project",
	Run: func(cmd *cobra.Command, args []string) {
		var projectName string
		var notificationName string
		if len(args) < 2 {
			if cmdProject.Name != "" && len(args) == 1 {
				projectName = cmdProject.Name
				notificationName = args[0]
			} else {
				fmt.Println("Not enough arguments. Requires: project name and notification name")
				cmd.Help()
				os.Exit(1)
			}
		} else {
			projectName = args[0]
			notificationName = args[1]
		}
		addResult, err := projects.AddSlackNotificationToProject(projectName, notificationName)
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

var addRocketChatNotificationCmd = &cobra.Command{
	Use:   "rocketchat [notification name] [channel] [webhook url]",
	Short: "Add a new rocketchat notification",
	Run: func(cmd *cobra.Command, args []string) {
		var notificationName string
		var channel string
		var webhookURL string
		if len(args) < 3 {
			fmt.Println("Not enough arguments. Requires: notifcation name, channel, and webhook url")
			cmd.Help()
			os.Exit(1)
		}
		notificationName = args[0]
		channel = args[1]
		webhookURL = args[2]
		addResult, err := projects.AddRocketChatNotification(notificationName, channel, webhookURL)
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
	Use:   "project-rocketchat [project name] [notification name]",
	Short: "Add a rocketchat notification to a project",
	Run: func(cmd *cobra.Command, args []string) {
		var projectName string
		var notificationName string
		if len(args) < 2 {
			if cmdProject.Name != "" && len(args) == 1 {
				projectName = cmdProject.Name
				notificationName = args[0]
			} else {
				fmt.Println("Not enough arguments. Requires: project name and notification name")
				cmd.Help()
				os.Exit(1)
			}
		} else {
			projectName = args[0]
			notificationName = args[1]
		}
		addResult, err := projects.AddRocketChatNotificationToProject(projectName, notificationName)
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

var deleteProjectSlackNotificationCmd = &cobra.Command{
	Use:   "project-slack [project name] [notification name]",
	Short: "Delete a slack notification from a project",
	Run: func(cmd *cobra.Command, args []string) {
		var projectName string
		var notificationName string
		if len(args) < 2 {
			if cmdProject.Name != "" && len(args) == 1 {
				projectName = cmdProject.Name
				notificationName = args[0]
			} else {
				fmt.Println("Not enough arguments. Requires: project name and notification name")
				cmd.Help()
				os.Exit(1)
			}
		} else {
			projectName = args[0]
			notificationName = args[1]
		}
		deleteResult, err := projects.DeleteSlackNotificationFromProject(projectName, notificationName)
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

var deleteProjectRocketChatNotificationCmd = &cobra.Command{
	Use:   "project-rocketchat [project name] [notification name]",
	Short: "Delete a rocketchat notification from a project",
	Run: func(cmd *cobra.Command, args []string) {
		var projectName string
		var notificationName string
		if len(args) < 2 {
			if cmdProject.Name != "" && len(args) == 1 {
				projectName = cmdProject.Name
				notificationName = args[0]
			} else {
				fmt.Println("Not enough arguments. Requires: project name and notification name")
				cmd.Help()
				os.Exit(1)
			}
		} else {
			projectName = args[0]
			notificationName = args[1]
		}
		deleteResult, err := projects.DeleteRocketChatNotificationFromProject(projectName, notificationName)
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
	Use:   "rocketchat [notification name]",
	Short: "Delete a rocketchat notification from lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		var notificationName string
		if len(args) < 1 {
			fmt.Println("Not enough arguments. Requires: notification name")
			cmd.Help()
			os.Exit(1)
		}
		notificationName = args[0]
		deleteResult, err := projects.DeleteRocketChatNotification(notificationName)
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

var deleteSlackNotificationCmd = &cobra.Command{
	Use:   "slack [notification name]",
	Short: "Delete a slack notification from lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		var notificationName string
		if len(args) < 1 {
			fmt.Println("Not enough arguments. Requires: notification name")
			cmd.Help()
			os.Exit(1)
		}
		notificationName = args[0]
		deleteResult, err := projects.DeleteSlackNotification(notificationName)
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
	Use:   "rocketchat [notification name]",
	Short: "Update an existing rocketchat notification",
	Run: func(cmd *cobra.Command, args []string) {
		var notificationName string
		if len(args) < 1 {
			fmt.Println("Not enough arguments. Requires: project name")
			cmd.Help()
			os.Exit(1)
		}
		notificationName = args[0]
		updateResult, err := projects.UpdateRocketChatNotification(notificationName, jsonPatch)
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

var updateSlackNotificationCmd = &cobra.Command{
	Use:   "slack [notification name]",
	Short: "Update an existing slack notification",
	Run: func(cmd *cobra.Command, args []string) {
		var notificationName string
		if len(args) < 1 {
			fmt.Println("Not enough arguments. Requires: project name")
			cmd.Help()
			os.Exit(1)
		}
		notificationName = args[0]
		updateResult, err := projects.UpdateSlackNotification(notificationName, jsonPatch)
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
