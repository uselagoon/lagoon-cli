package cmd

import (
	pygmy "github.com/fubarhouse/pygmy/v1/service/library"
	"github.com/spf13/cobra"
)

var pygmyCmd = &cobra.Command{
	Use:   "pygmy",
	Short: "start, stop or check the status of pygmy",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}

var pygmyCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "clean pygmy",
	Run: func(cmd *cobra.Command, args []string) {
		pygmy.Clean([]string{""})
	},
}

var pygmyRestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "restart pygmy",
	Run: func(cmd *cobra.Command, args []string) {
		pygmy.Restart([]string{""})
	},
}

var pygmyStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display status of pygmy",
	Run: func(cmd *cobra.Command, args []string) {
		pygmy.Status([]string{""})
	},
}

var pygmyStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop pygmy",
	Run: func(cmd *cobra.Command, args []string) {
		pygmy.Stop([]string{""})
	},
}

var pygmyUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Start pygmy",
	Run: func(cmd *cobra.Command, args []string) {
		pygmy.Up([]string{""})
	},
}

var pygmyUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Pull latest pygmy images",
	Run: func(cmd *cobra.Command, args []string) {
		pygmy.Stop([]string{""})
	},
}

var pygmyVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version of pygmy",
	Run: func(cmd *cobra.Command, args []string) {
		pygmy.Stop([]string{""})
	},
}

func init() {
	pygmyCmd.AddCommand(pygmyCleanCmd)
	pygmyCmd.AddCommand(pygmyRestartCmd)
	pygmyCmd.AddCommand(pygmyStatusCmd)
	pygmyCmd.AddCommand(pygmyStopCmd)
	pygmyCmd.AddCommand(pygmyUpCmd)
	pygmyCmd.AddCommand(pygmyUpdateCmd)
	pygmyCmd.AddCommand(pygmyVersionCmd)
}
