package projects

import (
	"encoding/json"
	"fmt"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/graphql"
	"github.com/amazeeio/lagoon-cli/output"
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
	returnResult, err := processAllProjects(allProjects)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

func processAllProjects(allProjects []byte) ([]byte, error) {
	var projects []api.Project
	err := json.Unmarshal([]byte(allProjects), &projects)
	if err != nil {
		return []byte(""), err
	}
	// process the data for output
	data := []output.Data{}
	for _, project := range projects {
		projectData := processProject(project)
		data = append(data, projectData)
	}
	dataMain := output.Table{
		Header: []string{"ID", "ProjectName", "GitURL", "DevEnvironments"},
		Data:   data,
	}
	return json.Marshal(dataMain)
}

func processProject(project api.Project) []string {
	// count the current dev environments in a project
	var currentDevEnvironments = 0
	for _, environment := range project.Environments {
		if environment.EnvironmentType == "development" {
			currentDevEnvironments++
		}
	}
	data := []string{
		fmt.Sprintf("%v", project.ID),
		fmt.Sprintf("%v", project.Name),
		fmt.Sprintf("%v", project.GitURL),
		fmt.Sprintf("%v/%v", currentDevEnvironments, project.DevelopmentEnvironmentsLimit),
	}
	return data
}

// GetProjectInfo will get basic info about a project
func GetProjectInfo(projectName string) ([]byte, error) {
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
	returnResult, err := processProjectInfo(projectByName)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

func processProjectInfo(projectByName []byte) ([]byte, error) {
	var project api.Project
	err := json.Unmarshal([]byte(projectByName), &project)
	if err != nil {
		return []byte(""), err
	}
	projectData := processProjectExtra(project)
	var data []output.Data
	data = append(data, projectData)
	dataMain := output.Table{
		Header: []string{"ID", "ProjectName", "GitURL", "Branches", "PullRequests", "ProductionRoute", "DevEnvironments"},
		Data:   data,
	}
	return json.Marshal(dataMain)
}

func processProjectExtra(project api.Project) []string {
	// count the current dev environments in a project
	var currentDevEnvironments = 0
	var projectRoute = "none"
	for _, environment := range project.Environments {
		if environment.EnvironmentType == "development" {
			currentDevEnvironments++
		}
		if environment.EnvironmentType == "production" {
			projectRoute = environment.Route
		}
	}
	data := []string{
		fmt.Sprintf("%v", project.ID),
		fmt.Sprintf("%v", project.Name),
		fmt.Sprintf("%v", project.GitURL),
		fmt.Sprintf("%v", project.Branches),
		fmt.Sprintf("%v", project.Pullrequests),
		fmt.Sprintf("%v", projectRoute),
		fmt.Sprintf("%v/%v", currentDevEnvironments, project.DevelopmentEnvironmentsLimit),
	}
	return data
}

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
	returnResult, err := processEnvironmentsList(projectByName)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

func processEnvironmentsList(projectByName []byte) ([]byte, error) {
	var projects api.Project
	err := json.Unmarshal([]byte(projectByName), &projects)
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
	// process the data for output
	data := []output.Data{}
	for _, environment := range projects.Environments {
		var envRoute = "none"
		if environment.Route != "" {
			envRoute = environment.Route
		}
		data = append(data, []string{
			fmt.Sprintf("%d", environment.ID),
			environment.Name,
			string(environment.DeployType),
			string(environment.EnvironmentType),
			envRoute,
		})
	}
	dataMain := output.Table{
		Header: []string{"ID", "Name", "DeployType", "Environment", "Route"}, //, "SSH"},
		Data:   data,
	}
	return json.Marshal(dataMain)
}

// AddProject .
func AddProject(projectName string, jsonPatch string) ([]byte, error) {
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}
	project := api.ProjectPatch{}
	err = json.Unmarshal([]byte(jsonPatch), &project)
	if err != nil {
		return []byte(""), err
	}
	project.Name = projectName
	projectAddResult, err := lagoonAPI.AddProject(project, graphql.ProjectByNameFragment)
	if err != nil {
		return []byte(""), err
	}
	return projectAddResult, nil
}

// DeleteProject .
func DeleteProject(projectName string) ([]byte, error) {
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}
	project := api.Project{
		Name: projectName,
	}
	returnResult, err := lagoonAPI.DeleteProject(project)
	return returnResult, err
}

// UpdateProject .
func UpdateProject(projectName string, jsonPatch string) ([]byte, error) {
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}
	// get the project id from name
	projectBName := api.Project{
		Name: projectName,
	}
	projectByName, err := lagoonAPI.GetProjectByName(projectBName, graphql.ProjectByNameFragment)
	if err != nil {
		return []byte(""), err
	}
	projectUpdate, err := processProjectUpdate(projectByName, jsonPatch)
	if err != nil {
		return []byte(""), err
	}
	returnResult, err := lagoonAPI.UpdateProject(projectUpdate, graphql.ProjectByNameFragment)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

func processProjectUpdate(projectByName []byte, jsonPatch string) (api.UpdateProject, error) {
	var projects api.Project
	var projectUpdate api.UpdateProject
	var project api.ProjectPatch
	err := json.Unmarshal([]byte(projectByName), &projects)
	if err != nil {
		return projectUpdate, err
	}
	projectID := projects.ID

	// patch the project by id
	err = json.Unmarshal([]byte(jsonPatch), &project)
	if err != nil {
		return projectUpdate, err
	}
	projectUpdate = api.UpdateProject{
		ID:    projectID,
		Patch: project,
	}
	return projectUpdate, nil
}
