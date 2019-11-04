package cmd

import (
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update project, environment, etc..",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}
