package cmd

import (
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a project or environment",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}

var deleteVariableCmd = &cobra.Command{
	Use:   "variable",
	Short: "Delete variables from environments or projects",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	deleteCmd.AddCommand(deleteVariableCmd)
	deleteCmd.AddCommand(deleteProjectSlackNotificationCmd)
	deleteCmd.AddCommand(deleteSlackNotificationCmd)
	deleteCmd.AddCommand(deleteProjectRocketChatNotificationCmd)
	deleteCmd.AddCommand(deleteRocketChatNotificationCmd)
}
