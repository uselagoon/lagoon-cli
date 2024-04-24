package api

import (
	"encoding/json"
	"errors"

	"github.com/machinebox/graphql"
)

// GetEnvironmentByName .
func (api *Interface) GetEnvironmentByName(environment EnvironmentByName, fragment string) ([]byte, error) {
	if fragment == "" {
		fragment = environmentByNameFragment
	}
	req := graphql.NewRequest(`
	query ($name: String!, $project: Int!) {
		environmentByName(name: $name, project: $project) {
			...Environment
		}
	}` + fragment)
	generateVars(req, environment)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["environmentByName"])
	if err != nil {
		return []byte(""), err
	}
	if api.debug {
		debugResponse(jsonBytes)
	}
	if string(jsonBytes) == "null" {
		return []byte(""), errors.New("GraphQL API returned a null response, the requested resource may not exist, or there was an error. Use `--debug` to check what was returned")
	}
	return jsonBytes, nil
}

// AddOrUpdateEnvironment .
func (api *Interface) AddOrUpdateEnvironment(environment AddUpdateEnvironment) ([]byte, error) {
	req := graphql.NewRequest(`
	mutation ($name: String!, $project: Int!, $deployType: DeployType!, $deployBaseRef: String!, $deployHeadRef: String, $deployTitle: String, $environmentType: EnvType!, $openshiftProjectName: String!) {
		addOrUpdateEnvironment(input: {
			name: $name,
			project: $project,
			deployType: $deployType,
			deployBaseRef: $deployBaseRef,
			deployHeadRef: $deployHeadRef,
			deployTitle: $deployTitle,
			environmentType: $environmentType,
			openshiftProjectName: $openshiftProjectName
		}) {
			id
			name
			project {
				name
			}
			deployType
			environmentType
			openshiftProjectName
			envVariables {
				name
				value
				scope
			}
		}
	}`)
	req.Var("name", environment.Name)
	generateVars(req, environment.Patch)
	// req.Var("project", environment.Patch.Project)
	// req.Var("deployType", environment.Patch.DeployType)
	// req.Var("deployBaseRef", environment.Patch.DeployBaseRef)
	// req.Var("deployHeadRef", environment.Patch.DeployHeadRef)
	// req.Var("deployTitle", environment.Patch.DeployTitle)
	// req.Var("environmentType", environment.Patch.EnvironmentType)
	// req.Var("openshiftProjectName", environment.Patch.OpenshiftProjectName)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["addOrUpdateEnvironment"])
	if err != nil {
		return []byte(""), err
	}
	if api.debug {
		debugResponse(jsonBytes)
	}
	if string(jsonBytes) == "null" {
		return []byte(""), errors.New("GraphQL API returned a null response, the requested resource may not exist, or there was an error. Use `--debug` to check what was returned")
	}
	return jsonBytes, nil
}

// SetEnvironmentServices .
func (api *Interface) SetEnvironmentServices(environment SetEnvironmentServices) ([]byte, error) {
	req := graphql.NewRequest(`
	mutation ($environment: Int!, $services: [String]!) {
		setEnvironmentServices(input: {
			environment: $environment
			services: $services
		}) {
			id
			name
		}
	}`)
	generateVars(req, environment)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["setEnvironmentServices"])
	if err != nil {
		return []byte(""), err
	}
	if api.debug {
		debugResponse(jsonBytes)
	}
	if string(jsonBytes) == "null" {
		return []byte(""), errors.New("GraphQL API returned a null response, the requested resource may not exist, or there was an error. Use `--debug` to check what was returned")
	}
	return jsonBytes, nil
}
