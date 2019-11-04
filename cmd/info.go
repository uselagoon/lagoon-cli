package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/lagoon/environments"
	"github.com/amazeeio/lagoon-cli/lagoon/projects"
	"github.com/amazeeio/lagoon-cli/output"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info on a project, or deployment",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}

var infoProjectCmd = &cobra.Command{
	Use:   "project [project]",
	Short: "Details about a project",
	Run: func(cmd *cobra.Command, args []string) {
		var projectName string
		if len(args) < 1 {
			if cmdProject.Name != "" {
				projectName = cmdProject.Name
			} else {
				fmt.Println("Not enough arguments. Requires: project name")
				cmd.Help()
				os.Exit(1)
			}
		} else {
			projectName = args[0]
		}

		returnedJSON, err := projects.ListEnvironmentsForProject(projectName)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}

		var dataMain output.Table
		err = json.Unmarshal([]byte(returnedJSON), &dataMain)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		if len(dataMain.Data) == 0 {
			output.RenderError("no data returned", outputOptions)
			os.Exit(1)
		}
		output.RenderOutput(dataMain, outputOptions)

	},
}

var infoDeploymentCmd = &cobra.Command{
	Use:   "deployment [remote id]",
	Short: "Get build log by remote id",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Not enough arguments. Requires: remote id")
			cmd.Help()
			os.Exit(1)
		}
		deploymentID := args[0]

		returnedJSON, err := environments.GetDeploymentLog(deploymentID)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		if string(returnedJSON) == "null" {
			output.RenderError("No data returned from lagoon, remote id might be wrong", outputOptions)
			os.Exit(1)
		}
		var deployment api.Deployment
		err = json.Unmarshal([]byte(returnedJSON), &deployment)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		fmt.Println("START LOG")
		fmt.Println(deployment.BuildLog)
		fmt.Println("END LOG")

	},
}

func init() {
	infoCmd.AddCommand(infoProjectCmd)
	infoCmd.AddCommand(infoDeploymentCmd)
}
