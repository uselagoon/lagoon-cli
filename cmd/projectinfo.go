package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var projectInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Details about a project",
	Run: func(cmd *cobra.Command, args []string) {
		apiToken := viper.GetString("lagoon_token")
		if apiToken == "" {
			fmt.Println("Need to run `lagoon login` first")
			os.Exit(1)
		}
		if len(args) == 0 {
			fmt.Println("You must provide a project name.")
			os.Exit(1)
		}
		if len(args) > 1 {
			fmt.Println("Too many arguments.")
			os.Exit(1)
		}
		projectName := args[0]
		getProject(projectName)
	},
}

func init() {
	projectCmd.AddCommand(projectInfoCmd)
}
