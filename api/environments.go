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
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["environmentByName"])
	if err != nil {
		return []byte(""), err
	}
	if string(jsonBytes) == "null" {
		return []byte(""), errors.New("graphql: returned null")
	}
	return jsonBytes, nil
}

// AddOrUpdateEnvironment .
func (api *Interface) AddOrUpdateEnvironment(project int, environment AddUpdateEnvironment) ([]byte, error) {
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
	generateVars(req, environment.Patch)
	req.Var("name", environment.Name)
	req.Var("project", project)
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["addOrUpdateEnvironment"])
	if err != nil {
		return []byte(""), err
	}
	if string(jsonBytes) == "null" {
		return []byte(""), errors.New("graphql: returned null")
	}
	return jsonBytes, nil
}

// UpdateEnvironment .
func (api *Interface) UpdateEnvironment(environment UpdateEnvironment) ([]byte, error) {
	req := graphql.NewRequest(`
	mutation {
		updateEnvironment(input: {
			id: ${environmentId},
			patch: ${patch}
		}) {
			id
			name
		}
	}`)
	generateVars(req, environment)
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["updateEnvironment"])
	if err != nil {
		return []byte(""), err
	}
	if string(jsonBytes) == "null" {
		return []byte(""), errors.New("graphql: returned null")
	}
	return jsonBytes, nil
}

// DeleteEnvironment .
func (api *Interface) DeleteEnvironment(environment DeleteEnvironment) ([]byte, error) {
	req := graphql.NewRequest(`
	mutation ($name: String!, $project: String!, $execute: Boolean) {
		deleteEnvironment(input: {
			name: $name
			project: $project
			execute: $execute
		})
	}`)
	generateVars(req, environment)
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["deleteEnvironment"])
	if err != nil {
		return []byte(""), err
	}
	if string(jsonBytes) == "null" {
		return []byte(""), errors.New("graphql: returned null")
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
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["setEnvironmentServices"])
	if err != nil {
		return []byte(""), err
	}
	if string(jsonBytes) == "null" {
		return []byte(""), errors.New("graphql: returned null")
	}
	return jsonBytes, nil
}
