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

var deleteNotificationCmd = &cobra.Command{
	Use:     "notification",
	Aliases: []string{"n"},
	Short:   "Delete notifications or delete notifications from projects",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateToken(lagoonCLIConfig.Current) // get a new token if the current one is invalid
	},
}

var deleteOrganizationCmd = &cobra.Command{
	Use:     "organization",
	Aliases: []string{"o"},
	Short:   "Add an organization, or add a group/project to an organization",
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
	deleteCmd.AddCommand(deleteNotificationCmd)
	deleteCmd.AddCommand(deleteUserCmd)
	deleteCmd.AddCommand(deleteSSHKeyCmd)
	deleteCmd.AddCommand(deleteUserFromGroupCmd)
	deleteCmd.AddCommand(deleteVariableCmd)
	deleteCmd.AddCommand(deleteDeployTargetConfigCmd)
	deleteCmd.AddCommand(deleteOrganizationCmd)
}
