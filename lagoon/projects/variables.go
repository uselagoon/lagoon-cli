package projects

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/graphql"
	"github.com/amazeeio/lagoon-cli/output"
)

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
	returnResult, err := json.Marshal(dataMain)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// AddEnvironmentVariableToProject will list all environments for a project
func AddEnvironmentVariableToProject(projectName string, envVar api.EnvVariable) ([]byte, error) {

	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}

	// get project info from lagoon, we need the project ID for later
	project := api.Project{
		Name: projectName,
	}
	projectByName, err := lagoonAPI.GetProjectByName(project, graphql.ProjectByNameFragment)
	if err != nil {
		return []byte(""), err
	}
	var projectInfo api.Project
	err = json.Unmarshal([]byte(projectByName), &projectInfo)
	if err != nil {
		return []byte(""), err
	}

	// run the query to add the environment variable to lagoon
	// we consume the project ID here
	customReq := api.CustomRequest{
		Query: `mutation addEnvironmentVariableToProject ($type: EnvVariableType!, $typeId: Int!, $scope: EnvVariableScope!, $name: String!, $value: String!) {
			addEnvVariable(input:{type: $type, typeId: $typeId, scope: $scope, name: $name, value: $value}) {
				id
			}
		}`,
		Variables: map[string]interface{}{
			"type":   api.ProjectVar,
			"typeId": projectInfo.ID,
			"scope":  envVar.Scope,
			"name":   envVar.Name,
			"value":  envVar.Value,
		},
		MappedResult: "addEnvVariable",
	}
	returnResult, err := lagoonAPI.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// DeleteEnvironmentVariableFromProject will list all environments for a project
func DeleteEnvironmentVariableFromProject(projectName string, envVar api.EnvVariable) ([]byte, error) {

	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}

	// get project info from lagoon, we need the project ID for later
	project := api.Project{
		Name: projectName,
	}
	projectByName, err := lagoonAPI.GetProjectByName(project, graphql.ProjectEnvVars)
	if err != nil {
		return []byte(""), err
	}
	var envVars api.Project
	err = json.Unmarshal([]byte(projectByName), &envVars)
	if err != nil {
		return []byte(""), err
	}
	for _, v := range envVars.EnvVariables {
		if envVar.Name == v.Name {
			envVar.ID = v.ID
			break
		}
	}
	if envVar.ID == 0 {
		return []byte(""), errors.New("no matching var found")
	}
	customReq := api.CustomRequest{
		Query: `mutation deleteEnvironmentVariableFromProject ($id: Int!) {
			deleteEnvVariable(input:{id: $id})
		}`,
		Variables: map[string]interface{}{
			"id": envVar.ID,
		},
		MappedResult: "deleteEnvVariable",
	}
	returnResult, err := lagoonAPI.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}
