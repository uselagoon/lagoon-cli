package api

import (
	"github.com/machinebox/graphql"
)

// GetEnvironmentByName .
func (api *Interface) GetEnvironmentByName(environment EnvironmentByName) (interface{}, error) {
	req := graphql.NewRequest(`
	query {
		environmentByName(name: "${name}", project:${projectId}) {
			id,
			name,
			route,
			routes,
			deployType,
			environmentType,
			openshiftProjectName,
			updated,
			created,
			deleted,
		}
	}`)
	generateVars(req, environment)
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// AddOrUpdateEnvironment .
func (api *Interface) AddOrUpdateEnvironment(environment AddUpdateEnvironment) (interface{}, error) {
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
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// UpdateEnvironment .
func (api *Interface) UpdateEnvironment(environment UpdateEnvironment) (interface{}, error) {
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
	return returnType, err
}

// DeleteEnvironment .
func (api *Interface) DeleteEnvironment(environment DeleteEnvironment) (interface{}, error) {
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
	return returnType, err
}

// SetEnvironmentServices .
func (api *Interface) SetEnvironmentServices(environment SetEnvironmentServices) (interface{}, error) {
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
	return returnType, err
}
