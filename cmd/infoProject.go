package cmd

import (
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/graphql"
	"github.com/amazeeio/lagoon-cli/output"

	"encoding/json"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var infoProjectCmd = &cobra.Command{
	Use:   "project [project]",
	Short: "Details about a project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Not enough arguments. Requires: project name")
			cmd.Help()
			os.Exit(1)
		}
		projectName := args[0]

		var jsonBytes []byte

		// set up a lagoonapi client
		lagoonAPI, err := graphql.LagoonAPI()
		if err != nil {
			fmt.Println(errorFormat(err.Error(), JSON))
			return
		}

		// get project info from lagoon
		project := api.Project{
			Name: projectName,
		}
		projectByName, err := lagoonAPI.GetProjectByName(project, graphql.ProjectByNameFragment)
		if err != nil {
			fmt.Println(errorFormat(err.Error(), JSON))
			return
		}
		jsonBytes, _ = json.Marshal(projectByName)

		// remap the returned data for processing
		reMappedResult := projectByName.(map[string]interface{})
		var projects api.Project
		jsonBytes, _ = json.Marshal(reMappedResult["project"])
		err = json.Unmarshal([]byte(jsonBytes), &projects)
		if err != nil {
			fmt.Println(errorFormat(err.Error(), JSON))
			return
		}

		// count the current dev environments in a project
		var currentDevEnvironments = 0
		for _, environment := range projects.Environments {
			if environment.EnvironmentType == "development" {
				currentDevEnvironments++
			}
		}

		fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Project Name"), projects.Name))
		fmt.Println(fmt.Sprintf("%s: %d", aurora.Yellow("Project ID"), projects.ID))
		fmt.Println()
		fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Git"), projects.GitURL))
		fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Branches"), projects.Branches))
		fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Pull Requests"), projects.Pullrequests))
		fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Production Environment"), projects.ProductionEnvironment))
		fmt.Println(fmt.Sprintf("%s: %d / %d", aurora.Yellow("Development Environments"), currentDevEnvironments, projects.DevelopmentEnvironmentsLimit))
		fmt.Println()

		// process the data for output
		data := []output.Data{}
		for _, environment := range projects.Environments {
			data = append(data, []string{
				fmt.Sprintf("%d", environment.ID),
				environment.Name,
				string(environment.DeployType),
				string(environment.EnvironmentType),
				environment.Route,
				fmt.Sprintf("ssh -p %s -t %s@%s", viper.GetString("lagoons."+cmdLagoon+".port"), environment.OpenshiftProjectName, viper.GetString("lagoons."+cmdLagoon+".hostname")),
			})
		}
		dataMain := output.Table{
			Header: []string{"ID", "Name", "Deploy Type", "Environment", "Route", "SSH"},
			Data:   data,
		}
		output.RenderTable(dataMain)

	},
}

func init() {
	infoCmd.AddCommand(infoProjectCmd)
}
