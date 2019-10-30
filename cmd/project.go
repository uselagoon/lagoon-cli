package cmd

import (
	"fmt"
	"os"

	"encoding/json"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/app"
	"github.com/amazeeio/lagoon-cli/graphql"
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

		lagoonAPI, err := graphql.LagoonAPI()
		if err != nil {
			fmt.Println(err)
			return
		}

		var jsonBytes []byte
		project := api.Project{
			Name: projectName,
		}

		fmt.Println(fmt.Sprintf("Deleting %s", projectName))

		if yesNo() {
			projectByName, err := lagoonAPI.DeleteProject(project)
			if err != nil {
				fmt.Println(err)
				return
			}
			jsonBytes, _ = json.Marshal(projectByName)
			reMappedResult := projectByName.(map[string]interface{})
			jsonBytes, _ = json.Marshal(reMappedResult["deleteProject"])

			if string(jsonBytes) == "success" {
				fmt.Println(fmt.Sprintf("Result: %s", aurora.Green(string(jsonBytes))))
			} else {
				fmt.Println(fmt.Sprintf("Result: %s", aurora.Yellow(string(jsonBytes))))
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

		var jsonBytes []byte

		lagoonAPI, err := graphql.LagoonAPI()
		if err != nil {
			fmt.Println(err)
			return
		}
		project := api.ProjectPatch{}
		err = json.Unmarshal([]byte(jsonPatch), &project)
		if err != nil {
			fmt.Println(err)
			return
		}
		project.Name = projectName

		projectAddResult, err := lagoonAPI.AddProject(project, graphql.ProjectByNameFragment)
		if err != nil {
			fmt.Println(err)
			return
		}
		jsonBytes, _ = json.Marshal(projectAddResult)

		reMappedResult := projectAddResult.(map[string]interface{})
		var addedProject api.Project
		jsonBytes, _ = json.Marshal(reMappedResult["addProject"])
		err = json.Unmarshal([]byte(jsonBytes), &addedProject)

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

func init() {
	addCmd.AddCommand(addProjectCmd)
	addProjectCmd.Flags().StringVarP(&jsonPatch, "json", "j", "", "JSON string to patch")
	deleteCmd.AddCommand(deleteProjectCmd)
}
