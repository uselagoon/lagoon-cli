package api

import (
	"encoding/json"
	"errors"

	"github.com/machinebox/graphql"
)

// GetOpenShiftInfoForProject .
func (api *Interface) GetOpenShiftInfoForProject(project Project) ([]byte, error) {
	req := graphql.NewRequest(`
	query ($project: String!) {
		project:projectByName(name: $project) {
			id
			openshift  {
				name
				consoleUrl
				token
				projectUser
				routerPattern
			}
			gitUrl
			deployTargetConfigs{
				deployTarget{
				  id
				  name
				  token
				}
			  }
			privateKey
			subfolder
			openshiftProjectPattern
			productionEnvironment
			envVariables {
				name
				value
				scope
			}
		}
	}`)
	generateVars(req, project)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["project"])
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

// AddProject .
func (api *Interface) AddProject(project ProjectPatch, fragment string) ([]byte, error) {
	if fragment == "" {
		fragment = projectFragment
	}
	//@TODO: Make this use the actual AddProjectInput type instead of defining everything here
	req := graphql.NewRequest(`
	mutation ($name: String!, $gitUrl: String!, $openshift: Int!, $productionEnvironment: String!, 
		$id: Int, $privateKey: String, $subfolder: String, $openshiftProjectPattern: String, 
		$branches: String, $pullrequests: String, $availability: ProjectAvailability, $autoIdle: Int, $developmentEnvironmentsLimit: Int) {
		addProject(input: {
			name: $name,
			gitUrl: $gitUrl,
			openshift: $openshift,
			productionEnvironment: $productionEnvironment,
			id: $id,
			privateKey: $privateKey,
			subfolder: $subfolder,
			openshiftProjectPattern: $openshiftProjectPattern,
			branches: $branches,
			pullrequests: $pullrequests,
			availability: $availability,
			autoIdle: $autoIdle,
			developmentEnvironmentsLimit: $developmentEnvironmentsLimit
		}) {
			...Project
		}
	}` + fragment)
	generateVars(req, project)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["addProject"])
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

// UpdateProject .
func (api *Interface) UpdateProject(project UpdateProject, fragment string) ([]byte, error) {
	if fragment == "" {
		fragment = projectFragment
	}
	req := graphql.NewRequest(`
	mutation ($id: Int!, $patch: UpdateProjectPatchInput!) {
		updateProject(input: {
		  	id: $id
			patch: $patch
		}) {
			...Project
		}
	}` + fragment)
	generateVars(req, project)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["updateProject"])
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

// DeleteProject .
func (api *Interface) DeleteProject(project Project) ([]byte, error) {
	req := graphql.NewRequest(`
	mutation ($name: String!) {
		deleteProject(input: {
		  	project: $name
		})
	}`)
	generateVars(req, project)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["deleteProject"])
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

// GetProductionEnvironmentForProject .
func (api *Interface) GetProductionEnvironmentForProject(project Project) ([]byte, error) {
	req := graphql.NewRequest(`
	query ($name: String!) {
		project:projectByName(name: $name){
			productionEnvironment
		}
	}`)
	generateVars(req, project)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["project"])
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

// GetEnvironmentByOpenshiftProjectName .
func (api *Interface) GetEnvironmentByOpenshiftProjectName(environment Environment) ([]byte, error) {
	req := graphql.NewRequest(`
	query {
		environmentByOpenshiftProjectName(openshiftProjectName: "${openshiftProjectName}") {
			id,
			name,
			project {
				name
			}
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
	jsonBytes, err := json.Marshal(reMappedResult["environmentByOpenshiftProjectName"])
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

// GetProjectsByGitURL .
func (api *Interface) GetProjectsByGitURL(project Project) ([]byte, error) {
	req := graphql.NewRequest(`
	query {
		allProjects(gitUrl: "${gitUrl}") {
			name
			productionEnvironment
			openshift {
				consoleUrl
				token
				projectUser
				routerPattern
			}
		}
	}`)
	generateVars(req, project)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["allProjects"])
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

// GetProjectByName .
func (api *Interface) GetProjectByName(project Project, fragment string) ([]byte, error) {
	if fragment == "" {
		fragment = projectFragment
	}
	req := graphql.NewRequest(`
	query ($name: String!){
		project:projectByName(name: $name) {
		  	...Project
		}
	}` + fragment)
	generateVars(req, project)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["project"])
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

// GetAllProjects .
func (api *Interface) GetAllProjects(fragment string) ([]byte, error) {
	if fragment == "" {
		fragment = projectFragment
	}
	req := graphql.NewRequest(`
	query {
		allProjects {
			...Project
		}
	}` + fragment)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["allProjects"])
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

// GetRocketChatInfoForProject .
func (api *Interface) GetRocketChatInfoForProject(project Project, fragment string) ([]byte, error) {
	if fragment == "" {
		fragment = notificationsRocketChatFragment
	}
	req := graphql.NewRequest(`
	query ($name: String!) {
		project:projectByName(name: $name) {
			rocketchats: notifications(type: ROCKETCHAT) {
				...Notification
			}
		}
	}` + fragment)
	generateVars(req, project)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["project"])
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

// GetSlackInfoForProject .
func (api *Interface) GetSlackInfoForProject(project Project, fragment string) ([]byte, error) {
	if fragment == "" {
		fragment = notificationsSlackFragment
	}
	req := graphql.NewRequest(`
	query ($name: String!){
		project:projectByName(name: $name) {
			slacks: notifications(type: SLACK) {
				...Notification
			}
		}
	}` + fragment)
	generateVars(req, project)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["project"])
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

// GetEnvironmentsForProject .
func (api *Interface) GetEnvironmentsForProject(project Project) ([]byte, error) {
	req := graphql.NewRequest(`
	query ($name: String!){
		project:projectByName(name: $name){
			developmentEnvironmentsLimit
			productionEnvironment
			environments(includeDeleted:false) { name, environmentType }
		}
	}`)
	generateVars(req, project)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["project"])
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

// GetDeploymentByRemoteID .
func (api *Interface) GetDeploymentByRemoteID(deployment Deployment) ([]byte, error) {
	req := graphql.NewRequest(`
	query deploymentByRemoteId($id: String!) {
		deploymentByRemoteId(id: $id) {
			...Deployment
		}
	}` + deploymentFragment)
	generateVars(req, deployment)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["deploymentByRemoteId"])
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

// AddDeployment .
func (api *Interface) AddDeployment(deployment Deployment) ([]byte, error) {
	req := graphql.NewRequest(`
	mutation ($name: String!, $status: DeploymentStatusType!, $created: String!, $environment: Int!, $id: Int, $remoteId: String, $started: String, $completed: String) {
		addDeployment(input: {
			name: $name
			status: $status
			created: $created
			environment: $environment
			id: $id
			remoteId: $remoteId
			started: $started
			completed: $completed
		}) {
		  	...Deployment
		}
	}` + deploymentFragment)
	generateVars(req, deployment)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["addDeployment"])
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

// UpdateDeployment .
func (api *Interface) UpdateDeployment(deployment UpdateDeployment) ([]byte, error) {
	req := graphql.NewRequest(`
	mutation ($id: Int!, $patch: UpdateDeploymentPatchInput!) {
		updateDeployment(input: {
			id: $id
			patch: $patch
		}) {
		  	...Deployment
		}
	}` + deploymentFragment)
	generateVars(req, deployment)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["updateDeployment"])
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
