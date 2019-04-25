package cmd

import (
	"github.com/spf13/cobra"
)

var projectListCmd = &cobra.Command{
	Use:   "list",
	Short: "Show your projects",
	Run: func(cmd *cobra.Command, args []string) {
		listProjects()
	},
}

func init() {
	projectCmd.AddCommand(projectListCmd)
}
