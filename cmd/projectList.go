package cmd

import (
	"fmt"
	// "os"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/graphql"
	"github.com/amazeeio/lagoon-cli/output"

	"encoding/json"
	// "github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var projectListCmd = &cobra.Command{
	Use:   "list",
	Short: "Show your projects",
	Run: func(cmd *cobra.Command, args []string) {

		lagoonAPI, err := graphql.LagoonAPI()
		if err != nil {
			fmt.Println(err)
			return
		}

		var jsonBytes []byte
		allProjects, err := lagoonAPI.GetAllProjects(graphql.AllProjectsFragment)
		if err != nil {
			fmt.Println(err)
			return
		}

		jsonBytes, _ = json.Marshal(allProjects)
		reMappedResult := allProjects.(map[string]interface{})
		var projects []api.Project
		jsonBytes, _ = json.Marshal(reMappedResult["allProjects"])
		err = json.Unmarshal([]byte(jsonBytes), &projects)
		if err != nil {
			fmt.Println(err)
			return
		}

		data := []output.Data{}
		for _, project := range projects {
			data = append(data, []string{
				fmt.Sprintf("%v", project.ID),
				fmt.Sprintf("%v", project.Name),
				fmt.Sprintf("%v", project.GitURL),
				fmt.Sprintf("%v/%v", len(project.Environments), project.DevelopmentEnvironmentsLimit),
			})
		}
		dataMain := output.Table{
			Header: []string{"ID", "Project Name", "Git URL", "Dev Environments"},
			Data:   data,
		}
		output.RenderTable(dataMain)

		fmt.Println()
		fmt.Println("To view a project's details, run `lagoon project info {name}`.")
	},
}

func init() {
	projectCmd.AddCommand(projectListCmd)

}
