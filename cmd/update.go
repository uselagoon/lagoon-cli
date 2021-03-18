package cmd

import (
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"u"},
	Short:   "Update a resource",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateToken(lagoonCLIConfig.Current) // get a new token if the current one is invalid
	},
}

func init() {
	updateCmd.AddCommand(updateProjectCmd)
	updateCmd.AddCommand(updateRocketChatNotificationCmd)
	updateCmd.AddCommand(updateSlackNotificationCmd)
	updateCmd.AddCommand(updateUserCmd)
}
