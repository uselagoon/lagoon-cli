package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Show your projects, or details about a project",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if !ValidateToken() {
			fmt.Println("Need to run `lagoon login` first")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)
}
