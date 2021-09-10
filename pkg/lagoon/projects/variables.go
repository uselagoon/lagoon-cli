package projects

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/uselagoon/lagoon-cli/pkg/api"
	"github.com/uselagoon/lagoon-cli/pkg/graphql"
	"github.com/uselagoon/lagoon-cli/pkg/output"
)

// ListProjectVariables will list the environment variables for a project and all environments attached
func (p *Projects) ListProjectVariables(projectName string, revealValue bool) ([]byte, error) {
	// get project info from lagoon, we need the project ID for later
	project := api.Project{
		Name: projectName,
	}
	queryFragment := graphql.ProjectEnvVars
	if revealValue {
		queryFragment = graphql.ProjectEnvVarsRevealed
	}
	projectByName, err := p.api.GetProjectByName(project, queryFragment)
	if err != nil {
		return []byte(""), err
	}
	returnResult, err := processProjectVariables(projectByName, revealValue)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

func processProjectVariables(projectInfo []byte, revealValue bool) ([]byte, error) {
	var project api.Project
	err := json.Unmarshal([]byte(projectInfo), &project)
	if err != nil {
		return []byte(""), err
	}
	data := []output.Data{}
	if len(project.EnvVariables) != 0 {
		for _, projectEnvVar := range project.EnvVariables {
			envVarRow := []string{
				fmt.Sprintf("%v", projectEnvVar.ID),
				project.Name,
				projectEnvVar.Scope,
				projectEnvVar.Name,
			}
			if revealValue {
				envVarRow = append(envVarRow, projectEnvVar.Value)
			}
			data = append(data, envVarRow)
		}
	}
	dataMain := output.Table{
		Header: []string{"ID", "Project", "Scope", "VariableName"},
		Data:   data,
	}
	if revealValue {
		dataMain.Header = append(dataMain.Header, "VariableValue")
	}
	return json.Marshal(dataMain)
}

// AddEnvironmentVariableToProject will list all environments for a project
func (p *Projects) AddEnvironmentVariableToProject(projectName string, envVar api.EnvVariable) ([]byte, error) {
	// get project info from lagoon, we need the project ID for later
	project := api.Project{
		Name: projectName,
	}
	projectByName, err := p.api.GetProjectByName(project, graphql.ProjectByNameFragment)
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
	returnResult, err := p.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// DeleteEnvironmentVariableFromProject will list all environments for a project
func (p *Projects) DeleteEnvironmentVariableFromProject(projectName string, envVar api.EnvVariable) ([]byte, error) {
	// get project info from lagoon, we need the project ID for later
	project := api.Project{
		Name: projectName,
	}
	projectByName, err := p.api.GetProjectByName(project, graphql.ProjectEnvVars)
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
	returnResult, err := p.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}
