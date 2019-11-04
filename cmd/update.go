package cmd

import (
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update project, environment, or notification",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	updateCmd.AddCommand(updateSlackNotificationCmd)
	updateSlackNotificationCmd.Flags().StringVarP(&jsonPatch, "json", "j", "", "JSON string to patch")
	updateCmd.AddCommand(updateRocketChatNotificationCmd)
	updateRocketChatNotificationCmd.Flags().StringVarP(&jsonPatch, "json", "j", "", "JSON string to patch")
}
