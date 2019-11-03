package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/lagoon/environments"
	"github.com/amazeeio/lagoon-cli/lagoon/projects"
	"github.com/amazeeio/lagoon-cli/output"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list projects, environment, etc..",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}

var listProjectCmd = &cobra.Command{
	Use:   "projects",
	Short: "Show your projects",
	Run: func(cmd *cobra.Command, args []string) {
		returnedJSON, err := projects.ListAllProjects()
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

var listVariablesCmd = &cobra.Command{
	Use:   "variables",
	Short: "Show your variables for a project",
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

		returnedJSON, err := projects.ListEnvironmentVariables(projectName, revealValue)
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

var listDeploymentsCmd = &cobra.Command{
	Use:   "deployments [project name] [environment name]",
	Short: "Show your deployments for an environment",
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

		returnedJSON, err := environments.GetEnvironmentDeployments(projectName, environmentName)
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

func init() {
	listCmd.AddCommand(listProjectCmd)
	listCmd.AddCommand(listVariablesCmd)
	listCmd.AddCommand(listDeploymentsCmd)
	listCmd.AddCommand(listRocketChatsCmd)
	listCmd.AddCommand(listSlackCmd)
	listVariablesCmd.Flags().BoolVarP(&revealValue, "reveal", "", false, "Reveal the variable values")
}
