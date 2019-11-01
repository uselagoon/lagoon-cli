package cmd

import (
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/graphql"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var deleteEnvCmd = &cobra.Command{
	Use:   "environment [project name] [environment name]",
	Short: "Delete an environment",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			fmt.Println("Not enough arguments. Requires: project name and environment.")
			cmd.Help()
			os.Exit(1)
		}
		projectName := args[0]
		projectEnvironment := args[1]

		lagoonAPI, err := graphql.LagoonAPI()
		if err != nil {
			fmt.Println(err)
			return
		}

		evironment := api.DeleteEnvironment{
			Name:    projectEnvironment,
			Project: projectName,
			Execute: true,
		}

		fmt.Println(fmt.Sprintf("Deleting %s-%s", projectName, projectEnvironment))

		if yesNo() {
			projectByName, err := lagoonAPI.DeleteEnvironment(evironment)
			if err != nil {
				fmt.Println(err)
				return
			}
			if string(projectByName) == "success" {
				fmt.Println(fmt.Sprintf("Result: %s", aurora.Green(string(projectByName))))
			} else {
				fmt.Println(fmt.Sprintf("Result: %s", aurora.Yellow(string(projectByName))))
			}
		}

	},
}

func init() {
	deleteCmd.AddCommand(deleteEnvCmd)
}
