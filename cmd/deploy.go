package cmd

import (
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/lagoon/environments"
	"github.com/amazeeio/lagoon-cli/output"
	"github.com/spf13/cobra"
)

var deployEnvCmd = &cobra.Command{
	Use:   "deploy [project name] [branch name]",
	Short: "Deploy a latest branch",
	Long:  "Deploy a latest branch",
	Run: func(cmd *cobra.Command, args []string) {
		var projectName string
		var environmentName string
		if len(args) < 2 {
			if cmdProject.Name != "" && len(args) == 1 {
				projectName = cmdProject.Name
				environmentName = args[0]
			} else {
				fmt.Println("Not enough arguments. Requires: project name and environment name")
				cmd.Help()
				os.Exit(1)
			}
		} else {
			projectName = args[0]
			environmentName = args[1]
		}

		if !outputOptions.JSON {
			fmt.Println(fmt.Sprintf("Deploying %s %s", projectName, environmentName))
		}

		if yesNo() {
			deployResult, err := environments.DeployEnvironmentBranch(projectName, environmentName)
			if err != nil {
				output.RenderError(err.Error(), outputOptions)
				os.Exit(1)
			}
			resultData := output.Result{
				Result: string(deployResult),
			}
			output.RenderResult(resultData, outputOptions)
		}

	},
}

/* @TODO
Need to be able to support more than just deploying the latest branch, like deploying pull requests?
*/
