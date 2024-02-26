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

var addNotificationCmd = &cobra.Command{
	Use:     "notification",
	Aliases: []string{"n"},
	Short:   "Add notifications or add notifications to projects",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateToken(lagoonCLIConfig.Current) // get a new token if the current one is invalid
	},
}

var addOrganizationCmd = &cobra.Command{
	Use:     "organization",
	Aliases: []string{"o"},
	Short:   "Add an organization, or add a group/project to an organization",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateToken(lagoonCLIConfig.Current) // get a new token if the current one is invalid
	},
}

func init() {
	addCmd.AddCommand(addDeployTargetCmd)
	addCmd.AddCommand(addGroupCmd)
	addCmd.AddCommand(addProjectCmd)
	addCmd.AddCommand(addProjectToGroupCmd)
	addCmd.AddCommand(addNotificationCmd)
	addCmd.AddCommand(addUserCmd)
	addCmd.AddCommand(addOrganizationCmd)
	addCmd.AddCommand(addUserToGroupCmd)
	addCmd.AddCommand(addUserSSHKeyCmd)
	addCmd.AddCommand(addVariableCmd)
	addCmd.AddCommand(addDeployTargetConfigCmd)
}
