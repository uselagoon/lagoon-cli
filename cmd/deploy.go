package cmd

import (
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/lagoon/environments"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var deployEnvCmd = &cobra.Command{
	Use:   "deploy [project name] [branch name]",
	Short: "Deploy a branch environment",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			fmt.Println("Not enough arguments. Requires: project name and environment.")
			cmd.Help()
			os.Exit(1)
		}
		projectName := args[0]
		projectEnvironment := args[1]

		fmt.Println(fmt.Sprintf("Deploying %s %s", projectName, projectEnvironment))

		if yesNo() {
			deployResult, err := environments.DeployEnvironmentBranch(projectName, projectEnvironment)
			if err != nil {
				fmt.Println(errorFormat(err.Error(), JSON))
				return
			}
			fmt.Println(string(deployResult))

			if string(deployResult) == "success" {
				fmt.Println(fmt.Sprintf("Result: %s", aurora.Green(string(deployResult))))
			} else {
				fmt.Println(fmt.Sprintf("Result: %s", aurora.Yellow(string(deployResult))))
			}
		}

	},
}
