package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/lagoon/projects"
	"github.com/amazeeio/lagoon-cli/output"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// NotificationFlags .
type NotificationFlags struct {
	Project             string `json:"project,omitempty"`
	NotificationName    string `json:"name,omitempty"`
	NotificationNewName string `json:"newname,omitempty"`
	NotificationOldName string `json:"old,omitempty"`
	NotificationWebhook string `json:"webhook,omitempty"`
	NotificationChannel string `json:"channel,omitempty"`
}

func parseNotificationFlags(flags pflag.FlagSet) NotificationFlags {
	configMap := make(map[string]interface{})
	flags.VisitAll(func(f *pflag.Flag) {
		if flags.Changed(f.Name) {
			configMap[f.Name] = f.Value
		}
	})
	jsonStr, _ := json.Marshal(configMap)
	parsedFlags := NotificationFlags{}
	json.Unmarshal(jsonStr, &parsedFlags)
	return parsedFlags
}

var listSlackCmd = &cobra.Command{
	Use:   "slack",
	Short: "Slack details about a project",
	Run: func(cmd *cobra.Command, args []string) {
		var returnedJSON []byte
		var err error
		if listAllProjects {
			returnedJSON, err = projects.ListAllSlacks()
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

			returnedJSON, err = projects.ListAllProjectSlacks(notificationFlags.Project)
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
	Use:   "rocketchat",
	Short: "Rocketchat details about a project",
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

			returnedJSON, err = projects.ListAllProjectRocketChats(notificationFlags.Project)
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
	Use:   "slack",
	Short: "Add a new slack notification",
	Run: func(cmd *cobra.Command, args []string) {
		notificationFlags := parseNotificationFlags(*cmd.Flags())
		if notificationFlags.NotificationName == "" || notificationFlags.NotificationChannel == "" || notificationFlags.NotificationWebhook == "" {
			fmt.Println("Not enough arguments. Requires: notifcation name, channel, and webhook url")
			cmd.Help()
			os.Exit(1)
		}

		addResult, err := projects.AddSlackNotification(notificationFlags.NotificationName, notificationFlags.NotificationChannel, notificationFlags.NotificationWebhook)
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
	Use:   "project-slack",
	Short: "Add a slack notification to a project",
	Run: func(cmd *cobra.Command, args []string) {
		notificationFlags := parseNotificationFlags(*cmd.Flags())
		if notificationFlags.Project == "" || notificationFlags.NotificationName == "" {
			fmt.Println("Not enough arguments. Requires: project name and notifcation name")
			cmd.Help()
			os.Exit(1)
		}
		addResult, err := projects.AddSlackNotificationToProject(notificationFlags.Project, notificationFlags.NotificationName)
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
	Use:   "rocketchat",
	Short: "Add a new rocketchat notification",
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
	Use:   "project-rocketchat",
	Short: "Add a rocketchat notification to a project",
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

var deleteProjectSlackNotificationCmd = &cobra.Command{
	Use:   "project-slack",
	Short: "Delete a slack notification from a project",
	Run: func(cmd *cobra.Command, args []string) {
		notificationFlags := parseNotificationFlags(*cmd.Flags())
		if notificationFlags.Project == "" || notificationFlags.NotificationName == "" {
			fmt.Println("Not enough arguments. Requires: project name and notifcation name")
			cmd.Help()
			os.Exit(1)
		}
		deleteResult, err := projects.DeleteSlackNotificationFromProject(notificationFlags.Project, notificationFlags.NotificationName)
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
	Use:   "project-rocketchat",
	Short: "Delete a rocketchat notification from a project",
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
	Use:   "rocketchat",
	Short: "Delete a rocketchat notification from lagoon",
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

var deleteSlackNotificationCmd = &cobra.Command{
	Use:   "slack",
	Short: "Delete a slack notification from lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		notificationFlags := parseNotificationFlags(*cmd.Flags())
		if notificationFlags.NotificationName == "" {
			fmt.Println("Not enough arguments. Requires: notifcation name")
			cmd.Help()
			os.Exit(1)
		}
		deleteResult, err := projects.DeleteSlackNotification(notificationFlags.NotificationName)
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
	Use:   "rocketchat",
	Short: "Update an existing rocketchat notification",
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

var updateSlackNotificationCmd = &cobra.Command{
	Use:   "slack",
	Short: "Update an existing slack notification",
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
		updateResult, err := projects.UpdateSlackNotification(notificationFlags.NotificationOldName, jsonPatch)
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
	addSlackNotificationCmd.Flags().StringVarP(&notificationName, "name", "n", "", "The name of the notification")
	addSlackNotificationCmd.Flags().StringVarP(&notificationWebhook, "webhook", "w", "", "The webhook URL of the notification")
	addSlackNotificationCmd.Flags().StringVarP(&notificationChannel, "channel", "c", "", "The channel for the notification")

	addRocketChatNotificationCmd.Flags().StringVarP(&notificationName, "name", "n", "", "The name of the notification")
	addRocketChatNotificationCmd.Flags().StringVarP(&notificationWebhook, "webhook", "w", "", "The webhook URL of the notification")
	addRocketChatNotificationCmd.Flags().StringVarP(&notificationChannel, "channel", "c", "", "The channel for the notification")

	addProjectRocketChatNotificationCmd.Flags().StringVarP(&notificationName, "name", "n", "", "The name of the notification")
	addProjectSlackNotificationCmd.Flags().StringVarP(&notificationName, "name", "n", "", "The name of the notification")

	deleteProjectSlackNotificationCmd.Flags().StringVarP(&notificationName, "name", "n", "", "The name of the notification")
	deleteProjectRocketChatNotificationCmd.Flags().StringVarP(&notificationName, "name", "n", "", "The name of the notification")

	updateRocketChatNotificationCmd.Flags().StringVarP(&notificationName, "name", "n", "", "The current name of the notification")
	updateRocketChatNotificationCmd.Flags().StringVarP(&notificationNewName, "newname", "N", "", "The name of the notification")
	updateRocketChatNotificationCmd.Flags().StringVarP(&notificationWebhook, "webhook", "w", "", "The webhook URL of the notification")
	updateRocketChatNotificationCmd.Flags().StringVarP(&notificationChannel, "channel", "c", "", "The channel for the notification")

	updateSlackNotificationCmd.Flags().StringVarP(&notificationName, "name", "n", "", "The current name of the notification")
	updateSlackNotificationCmd.Flags().StringVarP(&notificationNewName, "newname", "N", "", "The name of the notification")
	updateSlackNotificationCmd.Flags().StringVarP(&notificationWebhook, "webhook", "w", "", "The webhook URL of the notification")
	updateSlackNotificationCmd.Flags().StringVarP(&notificationChannel, "channel", "c", "", "The channel for the notification")
}
