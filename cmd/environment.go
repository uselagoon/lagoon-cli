package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/graphql"
	"github.com/amazeeio/lagoon-cli/output"
	"github.com/spf13/cobra"
)

var deleteEnvCmd = &cobra.Command{
	Use:   "environment [project name] [environment name]",
	Short: "Delete an environment",
	Run: func(cmd *cobra.Command, args []string) {
		var projectName string
		var projectEnvironment string
		if len(args) != 0 || cmdProject.Name == "" {
			if len(args) < 2 {
				fmt.Println("Not enough arguments. Requires: project name and environment")
				cmd.Help()
				os.Exit(1)
			}
			projectName = args[0]
			projectEnvironment = args[1]
		} else {
			if len(args) < 1 {
				fmt.Println("Not enough arguments. Requires: environment")
				cmd.Help()
				os.Exit(1)
			}
			projectName = strings.TrimSpace(cmdProject.Name)
			projectEnvironment = args[0]
		}

		lagoonAPI, err := graphql.LagoonAPI()
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
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
				output.RenderError(err.Error(), outputOptions)
				os.Exit(1)
			}
			resultData := output.Result{
				Result: string(projectByName),
			}
			output.RenderResult(resultData, outputOptions)
		}

	},
}

func init() {
	deleteCmd.AddCommand(deleteEnvCmd)
}
