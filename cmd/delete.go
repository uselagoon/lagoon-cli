package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	Short:   "Delete a project, or delete notifications and variables from projects or environments",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateToken(viper.GetString("current")) // get a new token if the current one is invalid
	},
}

func init() {
	deleteCmd.AddCommand(deleteVariableCmd)
	deleteCmd.AddCommand(delUserFromGroupCmd)
	deleteCmd.AddCommand(delProjectFromGroupCmd)
	deleteCmd.AddCommand(delUserCmd)
	deleteCmd.AddCommand(delGroupCmd)
	deleteCmd.AddCommand(deleteProjectSlackNotificationCmd)
	deleteCmd.AddCommand(deleteSlackNotificationCmd)
	deleteCmd.AddCommand(deleteProjectRocketChatNotificationCmd)
	deleteCmd.AddCommand(deleteRocketChatNotificationCmd)
}
