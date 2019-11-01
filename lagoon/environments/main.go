package environments

import (
	"fmt"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/graphql"
	"github.com/amazeeio/lagoon-cli/output"

	"encoding/json"
)

// ListEnvironmentsForProject will list all environments for a project
func ListEnvironmentsForProject(projectName string) ([]byte, error) {
	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}

	// get project info from lagoon
	project := api.Project{
		Name: projectName,
	}
	projectByName, err := lagoonAPI.GetProjectByName(project, graphql.ProjectByNameFragment)
	if err != nil {
		return []byte(""), err
	}
	var projects api.Project
	err = json.Unmarshal([]byte(projectByName), &projects)
	if err != nil {
		return []byte(""), err
	}

	// count the current dev environments in a project
	var currentDevEnvironments = 0
	for _, environment := range projects.Environments {
		if environment.EnvironmentType == "development" {
			currentDevEnvironments++
		}
	}

	// fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Project Name"), projects.Name))
	// fmt.Println(fmt.Sprintf("%s: %d", aurora.Yellow("Project ID"), projects.ID))
	// fmt.Println()
	// fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Git"), projects.GitURL))
	// fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Branches"), projects.Branches))
	// fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Pull Requests"), projects.Pullrequests))
	// fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Production Environment"), projects.ProductionEnvironment))
	// fmt.Println(fmt.Sprintf("%s: %d / %d", aurora.Yellow("Development Environments"), currentDevEnvironments, projects.DevelopmentEnvironmentsLimit))
	// fmt.Println()

	// process the data for output
	data := []output.Data{}
	for _, environment := range projects.Environments {
		data = append(data, []string{
			fmt.Sprintf("%d", environment.ID),
			environment.Name,
			string(environment.DeployType),
			string(environment.EnvironmentType),
			environment.Route,
			//fmt.Sprintf("ssh -p %s -t %s@%s", viper.GetString("lagoons."+cmdLagoon+".port"), environment.OpenshiftProjectName, viper.GetString("lagoons."+cmdLagoon+".hostname")),
		})
	}
	dataMain := output.Table{
		Header: []string{"ID", "Name", "Deploy Type", "Environment", "Route"}, //, "SSH"},
		Data:   data,
	}
	returnJSON, err := json.Marshal(dataMain)
	if err != nil {
		return []byte(""), err
	}
	return returnJSON, nil
}
