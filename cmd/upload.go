package cmd

import (
	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:     "upload",
	Aliases: []string{"u"},
	Short:   "Upload files to tasks",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateToken(lContext.Name) // get a new token if the current one is invalid
	},
}

func init() {
	uploadCmd.AddCommand(uploadFilesToTask)
}
