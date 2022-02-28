package cmd

import (
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add a project, or add notifications and variables to projects or environments",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateToken(lagoonCLIConfig.Current) // get a new token if the current one is invalid
	},
}

func init() {
	addCmd.AddCommand(addGroupCmd)
	addCmd.AddCommand(addProjectCmd)
	addCmd.AddCommand(addProjectToGroupCmd)
	addCmd.AddCommand(addProjectRocketChatNotificationCmd)
	addCmd.AddCommand(addProjectSlackNotificationCmd)
	addCmd.AddCommand(addRocketChatNotificationCmd)
	addCmd.AddCommand(addSlackNotificationCmd)
	addCmd.AddCommand(addUserCmd)
	addCmd.AddCommand(addUserToGroupCmd)
	addCmd.AddCommand(addUserSSHKeyCmd)
	addCmd.AddCommand(addVariableCmd)
}
