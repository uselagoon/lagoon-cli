package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var projectUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
			fmt.Println("Not enough arguments. Requires: project name, property, and value.")
			os.Exit(1)
		}
		projectName := args[0]
		projectProperty := args[1]
		projectValue := args[2]
		fmt.Println(fmt.Sprintf("Updating %s property %s with %s", projectName, projectProperty, projectValue))
	},
}

func init() {
	projectCmd.AddCommand(projectUpdateCmd)
}
