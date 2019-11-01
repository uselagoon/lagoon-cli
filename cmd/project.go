package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/app"
	"github.com/amazeeio/lagoon-cli/graphql"
	"github.com/amazeeio/lagoon-cli/lagoon/projects"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Show your projects, or details about a project",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// get a new token if the current one is invalid
		valid := graphql.VerifyTokenExpiry()
		if valid == false {
			loginErr := loginToken()
			if loginErr != nil {
				fmt.Println("Unable to refresh token, you may need to run `lagoon login` first")
				os.Exit(1)
			}
		}
		// can use this to pick out info from a local project for some operations
		cmdProject, _ = app.GetLocalProject()
	},
}

var deleteProjectCmd = &cobra.Command{
	Use:   "project [project name]",
	Short: "Delete a project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Not enough arguments. Requires: project name.")
			cmd.Help()
			os.Exit(1)
		}
		projectName := args[0]

		fmt.Println(fmt.Sprintf("Deleting %s", projectName))

		if yesNo() {
			deleteResult, err := projects.DeleteProject(projectName)
			if err != nil {
				fmt.Println(err)
				return
			}

			if string(deleteResult) == "success" {
				fmt.Println(fmt.Sprintf("Result: %s", aurora.Green(string(deleteResult))))
			} else {
				fmt.Println(fmt.Sprintf("Result: %s", aurora.Yellow(string(deleteResult))))
			}
		}
	},
}

var addProjectCmd = &cobra.Command{
	Use:   "project [project name]",
	Short: "Add a new project to lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Not enough arguments. Requires: project name")
			cmd.Help()
			os.Exit(1)
		}
		projectName := args[0]

		addResult, err := projects.AddProject(projectName, jsonPatch)
		if err != nil {
			fmt.Println(err)
			return
		}
		var addedProject api.Project
		err = json.Unmarshal([]byte(addResult), &addedProject)

		if err != nil {
			fmt.Println(err)
			return
		}

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(fmt.Sprintf("Result: %s\n", aurora.Green("success")))
			fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Project"), addedProject.Name))
			fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Git"), addedProject.GitURL))
		}
	},
}

var updateProjectCmd = &cobra.Command{
	Use:   "project [project name]",
	Short: "Update a project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Not enough arguments. Requires: project name")
			cmd.Help()
			os.Exit(1)
		}
		projectName := args[0]

		projectUpdateID, err := projects.UpdateProject(projectName, jsonPatch)
		if err != nil {
			fmt.Println(errorFormat(err.Error(), JSON))
			return
		}
		var updatedProject api.Project
		err = json.Unmarshal([]byte(projectUpdateID), &updatedProject)
		if err != nil {
			fmt.Println(errorFormat(err.Error(), JSON))
			return
		}

		fmt.Println(fmt.Sprintf("Result: %s", aurora.Green("success")))
		fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Project"), updatedProject.Name))

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
