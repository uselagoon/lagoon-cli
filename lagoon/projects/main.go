package projects

import (
	"fmt"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/graphql"
	"github.com/amazeeio/lagoon-cli/output"

	"encoding/json"
)

// ListAllProjects will list all projects
func ListAllProjects() ([]byte, error) {
	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}

	allProjects, err := lagoonAPI.GetAllProjects(graphql.AllProjectsFragment)
	if err != nil {
		return []byte(""), err
	}
	var projects []api.Project
	err = json.Unmarshal([]byte(allProjects), &projects)
	if err != nil {
		return []byte(""), err
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
	returnJSON, err := json.Marshal(dataMain)
	if err != nil {
		return []byte(""), err
	}
	return returnJSON, nil
}

// ListEnvironmentVariables will list the environment variables for a project and all environments attached
func ListEnvironmentVariables(projectName string, revealValue bool) ([]byte, error) {
	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}

	// get project info from lagoon, we need the project ID for later
	project := api.Project{
		Name: projectName,
	}
	queryFragment := graphql.ProjectAndEnvironmentEnvVars
	if revealValue {
		queryFragment = graphql.ProjectAndEnvironmentEnvVarsRevealed
	}
	projectByName, err := lagoonAPI.GetProjectByName(project, queryFragment)
	if err != nil {
		return []byte(""), err
	}
	var envVars api.Project
	err = json.Unmarshal([]byte(projectByName), &envVars)
	if err != nil {
		return []byte(""), err
	}
	data := []output.Data{}
	if len(envVars.EnvVariables) != 0 {
		for _, projectEnvVar := range envVars.EnvVariables {
			envVarRow := []string{
				fmt.Sprintf("%v", projectEnvVar.ID),
				project.Name,
				"",
				projectEnvVar.Scope,
				projectEnvVar.Name,
			}
			if revealValue {
				envVarRow = append(envVarRow, projectEnvVar.Value)
			}
			data = append(data, envVarRow)
		}
	}
	for _, v := range envVars.Environments {
		if len(v.EnvVariables) != 0 {
			for _, environmentEnvVar := range v.EnvVariables {
				envVarRow := []string{
					fmt.Sprintf("%v", environmentEnvVar.ID),
					project.Name,
					v.Name,
					environmentEnvVar.Scope,
					environmentEnvVar.Name,
				}
				if revealValue {
					envVarRow = append(envVarRow, environmentEnvVar.Value)
				}
				data = append(data, envVarRow)
			}
		}
	}
	dataMain := output.Table{
		Header: []string{"ID", "Project", "Environment", "Scope", "Variable Name"},
		Data:   data,
	}
	if revealValue {
		dataMain.Header = append(dataMain.Header, "Variable Value")
	}
	returnJSON, err := json.Marshal(dataMain)
	if err != nil {
		return []byte(""), err
	}
	return returnJSON, nil
}
