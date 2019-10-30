package cmd

import (
	"fmt"
	// "os"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/graphql"
	"github.com/amazeeio/lagoon-cli/output"

	"encoding/json"
	"github.com/spf13/cobra"
)

var listProjectCmd = &cobra.Command{
	Use:   "projects",
	Short: "Show your projects",
	Run: func(cmd *cobra.Command, args []string) {

		var jsonBytes []byte

		// set up a lagoonapi client
		lagoonAPI, err := graphql.LagoonAPI()
		if err != nil {
			fmt.Println(errorFormat(err.Error(), JSON))
			return
		}

		allProjects, err := lagoonAPI.GetAllProjects(graphql.AllProjectsFragment)
		if err != nil {
			fmt.Println(errorFormat(err.Error(), JSON))
			return
		}
		jsonBytes, _ = json.Marshal(allProjects)

		// process the result
		reMappedResult := allProjects.(map[string]interface{})
		var projects []api.Project
		jsonBytes, _ = json.Marshal(reMappedResult["allProjects"])
		err = json.Unmarshal([]byte(jsonBytes), &projects)
		if err != nil {
			fmt.Println(errorFormat(err.Error(), JSON))
			return
		}

		// process the data for output
		data := []output.Data{}
		for _, project := range projects {
			// count the current dev environments in a project
			var currentDevEnvironments = 0
			for _, environment := range project.Environments {
				if environment.EnvironmentType == "development" {
					currentDevEnvironments++
				}
			}
			data = append(data, []string{
				fmt.Sprintf("%v", project.ID),
				fmt.Sprintf("%v", project.Name),
				fmt.Sprintf("%v", project.GitURL),
				fmt.Sprintf("%v/%v", currentDevEnvironments, project.DevelopmentEnvironmentsLimit),
			})
		}
		dataMain := output.Table{
			Header: []string{"ID", "Project Name", "Git URL", "Dev Environments"},
			Data:   data,
		}
		output.RenderTable(dataMain)
		// output json
		// jsonBytes, _ = json.Marshal(dataMain)
		// fmt.Println(string(jsonBytes))

		fmt.Println()
		fmt.Println("To view a project's details, run `lagoon info project {name}`.")
	},
}

func init() {
	listCmd.AddCommand(listProjectCmd)
}
