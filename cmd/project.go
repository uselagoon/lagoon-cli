package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/app"
	"github.com/amazeeio/lagoon-cli/lagoon/projects"
	"github.com/amazeeio/lagoon-cli/output"
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Show your projects, or details about a project",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// can use this to pick out info from a local project for some operations
		cmdProject, _ = app.GetLocalProject()
		fmt.Println(cmdProject)
	},
}

var deleteProjectCmd = &cobra.Command{
	Use:   "project [project name]",
	Short: "Delete a project",
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

		if !outputOptions.JSON {
			fmt.Println(fmt.Sprintf("Deleting %s", projectName))
		}

		if yesNo() {
			deleteResult, err := projects.DeleteProject(projectName)
			if err != nil {
				output.RenderError(err.Error(), outputOptions)
				os.Exit(1)
			}
			resultData := output.Result{
				Result: string(deleteResult),
			}
			output.RenderResult(resultData, outputOptions)
		}
	},
}

var addProjectCmd = &cobra.Command{
	Use:   "project [project name]",
	Short: "Add a new project to lagoon",
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

		addResult, err := projects.AddProject(projectName, jsonPatch)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		var addedProject api.Project
		err = json.Unmarshal([]byte(addResult), &addedProject)

		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}

		if err != nil {
			output.RenderError(err.Error(), outputOptions)
		} else {
			resultData := output.Result{
				Result: "success",
				ResultData: map[string]interface{}{
					"Project Name": addedProject.Name,
					"GitURL":       addedProject.GitURL,
				},
			}
			output.RenderResult(resultData, outputOptions)
		}
	},
}

var updateProjectCmd = &cobra.Command{
	Use:   "project [project name]",
	Short: "Update a project",
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

		projectUpdateID, err := projects.UpdateProject(projectName, jsonPatch)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		var updatedProject api.Project
		err = json.Unmarshal([]byte(projectUpdateID), &updatedProject)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"Project Name": updatedProject.Name,
			},
		}
		output.RenderResult(resultData, outputOptions)

	},
}

func init() {
	addCmd.AddCommand(addProjectCmd)
	addProjectCmd.Flags().StringVarP(&jsonPatch, "json", "j", "", "JSON string to patch")
	deleteCmd.AddCommand(deleteProjectCmd)
	updateCmd.AddCommand(updateProjectCmd)
	updateProjectCmd.Flags().StringVarP(&jsonPatch, "json", "j", "", "JSON string to patch")
	// use json to patch, maybe re-introduce these better later on @TODO
	// updateProjectCmd.Flags().StringVarP(&projectGitURL, "giturl", "g", "", "GitURL of the project")
	// updateProjectCmd.Flags().StringVarP(&projectBranches, "branches", "b", "", "Branches of the project")
	// updateProjectCmd.Flags().StringVarP(&projectProductionEnvironment, "prod-env", "P", "", "Production environment of the project")
	// updateProjectCmd.Flags().StringVarP(&projectPullRequests, "pull-requests", "r", "", "Pull requests of the project")
	// updateProjectCmd.Flags().IntVarP(projectAutoIdle, "auto-idle", "a", 1, "Auto idle setting of the project")
	// updateProjectCmd.Flags().IntVarP(&projectDevelopmentEnvironmentsLimit, "dev-env-limit", "d", 5, "Auto idle setting of the project")
}
