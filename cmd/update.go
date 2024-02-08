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

var updateNotificationCmd = &cobra.Command{
	Use:     "notification",
	Aliases: []string{"n"},
	Short:   "List all notifications or notifications on projects",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateToken(lagoonCLIConfig.Current) // get a new token if the current one is invalid
	},
}

func init() {
	updateCmd.AddCommand(updateProjectCmd)
	updateCmd.AddCommand(updateEnvironmentCmd)
	updateCmd.AddCommand(updateNotificationCmd)
	updateCmd.AddCommand(updateUserCmd)
	updateCmd.AddCommand(updateDeployTargetConfigCmd)
	updateCmd.AddCommand(updateDeployTargetCmd)
	updateCmd.AddCommand(updateOrganizationCmd)
}
