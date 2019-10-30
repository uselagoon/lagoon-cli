package api

import (
	"github.com/machinebox/graphql"
)

// GetOpenShiftInfoForProject .
func (api *Interface) GetOpenShiftInfoForProject(project Project) (interface{}, error) {
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
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// AddProject .
func (api *Interface) AddProject(project ProjectPatch, fragment string) (interface{}, error) {
	if fragment == "" {
		fragment = projectFragment
	}
	req := graphql.NewRequest(`
	mutation ($name: String!, $gitUrl: String!, $openshift: Int!, $productionEnvironment: String!, $id: Int) {
		addProject(input: {
			name: $name,
			gitUrl: $gitUrl,
			openshift: $openshift,
			productionEnvironment: $productionEnvironment,
			id: $id,
		}) {
			...Project
		}
	}` + fragment)
	generateVars(req, project)
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// UpdateProject .
func (api *Interface) UpdateProject(project UpdateProject, fragment string) (interface{}, error) {
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
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// DeleteProject .
func (api *Interface) DeleteProject(project Project) (interface{}, error) {
	req := graphql.NewRequest(`
	mutation ($name: String!) {
		deleteProject(input: {
		  	project: $name
		})
	}`)
	generateVars(req, project)
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// GetProductionEnvironmentForProject .
func (api *Interface) GetProductionEnvironmentForProject(project Project) (interface{}, error) {
	req := graphql.NewRequest(`
	query ($name: String!) {
		project:projectByName(name: $name){
			productionEnvironment
		}
	}`)
	generateVars(req, project)
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// GetEnvironmentByOpenshiftProjectName .
func (api *Interface) GetEnvironmentByOpenshiftProjectName(environment Environment) (interface{}, error) {
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
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// GetProjectsByGitURL .
func (api *Interface) GetProjectsByGitURL(project Project) (interface{}, error) {
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
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// GetProjectByName .
func (api *Interface) GetProjectByName(project Project, fragment string) (interface{}, error) {
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
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// GetAllProjects .
func (api *Interface) GetAllProjects(fragment string) (interface{}, error) {
	if fragment == "" {
		fragment = projectFragment
	}
	req := graphql.NewRequest(`
	query {
		allProjects {
			...Project
		}
	}` + fragment)
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// GetRocketChatInfoForProject .
func (api *Interface) GetRocketChatInfoForProject(project Project) (interface{}, error) {
	req := graphql.NewRequest(`
	query ($name: String!) {
		project:projectByName(name: $name) {
			rocketchats: notifications(type: ROCKETCHAT) {
				...Notification
			}
		}
	}` + notificationsRocketChatFragment)
	generateVars(req, project)
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// GetSlackinfoForProject .
func (api *Interface) GetSlackinfoForProject(project Project) (interface{}, error) {
	req := graphql.NewRequest(`
	query ($name: String!){
		project:projectByName(name: $name) {
			rocketchats: notifications(type: SLACK) {
				...Notification
			}
		}
	}` + notificationsSlackFragment)
	generateVars(req, project)
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// GetActiveSystemForProject . @TODO
func (api *Interface) GetActiveSystemForProject(project Project, task string) (interface{}, error) {
	req := graphql.NewRequest(`
	query ($name: String!){
		project:projectByName(name: $name){
			${field}
			branches
			pullrequests
		}
	}` + notificationsSlackFragment)
	generateVars(req, project)
	req.Var("task", task)
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// GetEnvironmentsForProject .
func (api *Interface) GetEnvironmentsForProject(project Project) (interface{}, error) {
	req := graphql.NewRequest(`
	query ($name: String!){
		project:projectByName(name: $name){
			developmentEnvironmentsLimit
			productionEnvironment
			environments(includeDeleted:false) { name, environmentType }
		}
	}`)
	generateVars(req, project)
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// GetDeploymentByRemoteID .
func (api *Interface) GetDeploymentByRemoteID(deployment Deployment) (interface{}, error) {
	req := graphql.NewRequest(`
	query deploymentByRemoteId($id: String!) {
		deploymentByRemoteId(id: $id) {
			...Deployment
		}
	}` + deploymentFragment)
	generateVars(req, deployment)
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// AddDeployment .
func (api *Interface) AddDeployment(deployment Deployment) (interface{}, error) {
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
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// UpdateDeployment .
func (api *Interface) UpdateDeployment(deployment UpdateDeployment) (interface{}, error) {
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
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}
