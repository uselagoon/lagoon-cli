package cmd

import (
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "Delete a project, or delete notifications and variables from projects or environments",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateToken(lagoonCLIConfig.Current) // get a new token if the current one is invalid
	},
}

func init() {
	deleteCmd.AddCommand(deleteEnvCmd)
	deleteCmd.AddCommand(deleteGroupCmd)
	deleteCmd.AddCommand(deleteDeployTargetCmd)
	deleteCmd.AddCommand(deleteProjectCmd)
	deleteCmd.AddCommand(deleteProjectFromGroupCmd)
	deleteCmd.AddCommand(deleteProjectRocketChatNotificationCmd)
	deleteCmd.AddCommand(deleteProjectSlackNotificationCmd)
	deleteCmd.AddCommand(deleteRocketChatNotificationCmd)
	deleteCmd.AddCommand(deleteSlackNotificationCmd)
	deleteCmd.AddCommand(deleteUserCmd)
	deleteCmd.AddCommand(deleteSSHKeyCmd)
	deleteCmd.AddCommand(deleteUserFromGroupCmd)
	deleteCmd.AddCommand(deleteVariableCmd)
	deleteCmd.AddCommand(deleteDeployTargetConfigCmd)
}
