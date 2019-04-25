package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var projectListCmd = &cobra.Command{
	Use:   "list",
	Short: "Show your projects",
	Run: func(cmd *cobra.Command, args []string) {
		apiToken := viper.GetString("lagoon_token")
		if apiToken == "" {
			fmt.Println("Need to run `lagoon login` first")
			os.Exit(1)
		}
		listProjects()
	},
}

func init() {
	projectCmd.AddCommand(projectListCmd)
}
