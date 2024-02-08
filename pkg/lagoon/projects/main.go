package projects

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/uselagoon/lagoon-cli/internal/lagoon"
	"github.com/uselagoon/lagoon-cli/pkg/api"
	"github.com/uselagoon/lagoon-cli/pkg/graphql"
	"github.com/uselagoon/lagoon-cli/pkg/output"
)

// Projects .
type Projects struct {
	debug bool
	api   api.Client
}

// Client .
type Client interface {
	ListAllProjects() ([]byte, error)
	ListProjectVariables(string, bool) ([]byte, error)
	GetProjectKey(string, bool) ([]byte, error)
	GetProjectInfo(string) ([]byte, error)
	DeleteProject(string) ([]byte, error)
	AddProject(string, string) ([]byte, error)
	UpdateProject(string, string) ([]byte, error)
	AddEnvironmentVariableToProject(string, api.EnvVariable) ([]byte, error)
	DeleteEnvironmentVariableFromProject(string, api.EnvVariable) ([]byte, error)
}

// New .
func New(lc *lagoon.Config, debug bool) (Client, error) {
	lagoonAPI, err := graphql.LagoonAPI(lc, debug)
	if err != nil {
		return &Projects{}, err
	}
	return &Projects{
		debug: debug,
		api:   lagoonAPI,
	}, nil

}

var noDataError = "no data returned from the lagoon api"

// ListAllProjects will list all projects
func (p *Projects) ListAllProjects() ([]byte, error) {
	allProjects, err := p.api.GetAllProjects(graphql.AllProjectsFragment)
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
		Header: []string{"ID", "ProjectName", "GitURL", "ProductionEnvironment", "DevEnvironments"},
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
		fmt.Sprintf("%v", project.ProductionEnvironment),
		fmt.Sprintf("%v/%v", currentDevEnvironments, project.DevelopmentEnvironmentsLimit),
	}
	return data
}

// GetProjectInfo will get basic info about a project
func (p *Projects) GetProjectInfo(projectName string) ([]byte, error) {
	// get project info from lagoon
	project := api.Project{
		Name: projectName,
	}
	projectByName, err := p.api.GetProjectByName(project, graphql.ProjectByNameFragment)
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
		Header: []string{"ID", "ProjectName", "GitURL", "Branches", "PullRequests", "ProductionRoute", "DevEnvironments", "DevEnvLimit", "ProductionEnv", "RouterPattern", "AutoIdle", "FactsUI", "ProblemsUI"},
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
		fmt.Sprintf("%v", project.DevelopmentEnvironmentsLimit),
		fmt.Sprintf("%v", project.ProductionEnvironment),
		fmt.Sprintf("%s", project.RouterPattern),
		fmt.Sprintf("%v", *project.AutoIdle),
		fmt.Sprintf("%v", *project.FactsUI),
		fmt.Sprintf("%v", *project.ProblemsUI),
	}
	return data
}

// AddProject .
func (p *Projects) AddProject(projectName string, jsonPatch string) ([]byte, error) {
	project := api.ProjectPatch{}
	err := json.Unmarshal([]byte(jsonPatch), &project)
	if err != nil {
		return []byte(""), err
	}
	project.Name = projectName
	projectAddResult, err := p.api.AddProject(project, graphql.ProjectByNameFragment)
	if err != nil {
		return []byte(""), err
	}
	return projectAddResult, nil
}

// DeleteProject .
func (p *Projects) DeleteProject(projectName string) ([]byte, error) {
	project := api.Project{
		Name: projectName,
	}
	returnResult, err := p.api.DeleteProject(project)
	return returnResult, err
}

// UpdateProject .
func (p *Projects) UpdateProject(projectName string, jsonPatch string) ([]byte, error) {
	// get the project id from name
	projectBName := api.Project{
		Name: projectName,
	}
	projectByName, err := p.api.GetProjectByName(projectBName, graphql.ProjectByNameFragment)
	if err != nil {
		return []byte(""), err
	}
	projectUpdate, err := processProjectUpdate(projectByName, jsonPatch)
	if err != nil {
		return []byte(""), err
	}
	returnResult, err := p.api.UpdateProject(projectUpdate, graphql.ProjectByNameFragment)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

func processProjectUpdate(projectByName []byte, jsonPatch string) (api.UpdateProject, error) {
	var projects api.Project
	var projectUpdate api.UpdateProject
	var project api.ProjectPatch
	err := json.Unmarshal(projectByName, &projects)
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

// GetProjectKey will get basic info about a project
func (p *Projects) GetProjectKey(projectName string, revealValue bool) ([]byte, error) {
	// get project info from lagoon
	project := api.Project{
		Name: projectName,
	}
	keyFragment := `fragment Project on Project {
		publicKey
	}`
	if revealValue {
		keyFragment = `fragment Project on Project {
			privateKey
			publicKey
		}`
	}
	projectByName, err := p.api.GetProjectByName(project, keyFragment)
	if err != nil {
		return []byte(""), err
	}
	returnResult, err := processProjectKey(projectByName, revealValue)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

func processProjectKey(projectByName []byte, revealValue bool) ([]byte, error) {
	var project api.Project
	err := json.Unmarshal([]byte(projectByName), &project)
	if err != nil {
		return []byte(""), err
	}
	// get the key, but strip the newlines we don't need
	projectData := []string{
		strings.TrimSuffix(project.PublicKey, "\n"),
	}
	if revealValue {
		projectData = append(projectData, strings.TrimSuffix(project.PrivateKey, "\n"))
	}
	var data []output.Data
	data = append(data, projectData)
	dataMain := output.Table{
		Header: []string{"PublicKey"},
		Data:   data,
	}
	if revealValue {
		dataMain.Header = append(dataMain.Header, "PrivateKey")
	}
	return json.Marshal(dataMain)
}
