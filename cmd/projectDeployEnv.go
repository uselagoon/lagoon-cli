package cmd

import (
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/graphql"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var projectDeployEnvCmd = &cobra.Command{
	Use:   "deploy [project name] [environment name]",
	Short: "Deploy an environment",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			fmt.Println("Not enough arguments. Requires: project name and environment.")
			cmd.Help()
			os.Exit(1)
		}
		projectName := args[0]
		projectEnvironment := args[1]

		fmt.Println(fmt.Sprintf("Deploying %s-%s", projectName, projectEnvironment))

		if yesNo() {
			var responseData DeployResult
			err := graphql.GraphQLRequest(fmt.Sprintf(`mutation {
    deployEnvironmentBranch(
      input: {
        project:{name:"%s"}
        branchName:"%s"
      }
    )
  }`, projectName, projectEnvironment), &responseData)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				if responseData.DeployEnvironmentBranch == "success" {
					fmt.Println(fmt.Sprintf("Result: %s", aurora.Green(responseData.DeployEnvironmentBranch)))
				} else {
					fmt.Println(fmt.Sprintf("Result: %s", aurora.Yellow(responseData.DeployEnvironmentBranch)))
				}
			}
		}

	},
}

func init() {
	projectCmd.AddCommand(projectDeployEnvCmd)
}
