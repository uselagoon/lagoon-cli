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
	Short: "rocketchat details about a project",
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
	Use:   "slack [project name] [webhook url] [channel] [notification name]",
	Short: "Add a new slack notification to project",
	Run: func(cmd *cobra.Command, args []string) {
		var projectName string
		var webhookURL string
		var channel string
		var notificationName string
		if len(args) < 4 {
			if cmdProject.Name != "" && len(args) == 3 {
				projectName = cmdProject.Name
				webhookURL = args[0]
				channel = args[1]
				notificationName = args[2]
			} else {
				fmt.Println("Not enough arguments. Requires: project name and environment name")
				cmd.Help()
				os.Exit(1)
			}
		} else {
			projectName = args[0]
			webhookURL = args[1]
			channel = args[2]
			notificationName = args[3]
		}

		addResult, err := projects.AddSlackNotificationToProject(projectName, webhookURL, channel, notificationName)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		var addedProject api.NotificationSlack
		err = json.Unmarshal([]byte(addResult), &addedProject)

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

var addRocketChatNotificationCmd = &cobra.Command{
	Use:   "rocketchat [project name] [webhook url] [channel] [notification name]",
	Short: "Add a new rocketchat notification to project",
	Run: func(cmd *cobra.Command, args []string) {
		var projectName string
		var webhookURL string
		var channel string
		var notificationName string
		if len(args) < 4 {
			if cmdProject.Name != "" && len(args) == 3 {
				projectName = cmdProject.Name
				webhookURL = args[0]
				channel = args[1]
				notificationName = args[2]
			} else {
				fmt.Println("Not enough arguments. Requires: project name and environment name")
				cmd.Help()
				os.Exit(1)
			}
		} else {
			projectName = args[0]
			webhookURL = args[1]
			channel = args[2]
			notificationName = args[3]
		}

		addResult, err := projects.AddRocketChatNotificationToProject(projectName, webhookURL, channel, notificationName)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		var addedProject api.NotificationSlack
		err = json.Unmarshal([]byte(addResult), &addedProject)

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

var deleteSlackNotificationCmd = &cobra.Command{
	Use:   "slack [project name] [notification name]",
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

var deleteRocketChatNotificationCmd = &cobra.Command{
	Use:   "rocketchat [project name] [notification name]",
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
