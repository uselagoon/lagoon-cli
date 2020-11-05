package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var runCmd = &cobra.Command{
	Use:     "run",
	Aliases: []string{"r"},
	Short:   "Run a task against an environment",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateToken(viper.GetString("current")) // get a new token if the current one is invalid
	},
}

func init() {
	runCmd.AddCommand(runCustomTask)
	runCmd.AddCommand(runDrushArchiveDump)
	runCmd.AddCommand(runDrushCacheClear)
	runCmd.AddCommand(runDrushSQLDump)
	runCmd.AddCommand(runActiveStandbySwitch)
}
