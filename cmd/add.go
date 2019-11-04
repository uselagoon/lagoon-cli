package cmd

import (
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a project, or add notifications and variables to projects or environments",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}

var addVariableCmd = &cobra.Command{
	Use:   "variable",
	Short: "Add variables on environments or projects",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	addCmd.AddCommand(addVariableCmd)
	addCmd.AddCommand(addSlackNotificationCmd)
	addCmd.AddCommand(addRocketChatNotificationCmd)
}
