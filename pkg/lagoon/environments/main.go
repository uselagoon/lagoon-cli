package environments

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/uselagoon/lagoon-cli/internal/lagoon"
	"github.com/uselagoon/lagoon-cli/pkg/api"
	"github.com/uselagoon/lagoon-cli/pkg/graphql"
	"github.com/uselagoon/lagoon-cli/pkg/output"
)

// Environments .
type Environments struct {
	debug bool
	api   api.Client
}

// Client .
type Client interface {
	DeployEnvironmentBranch(string, string) ([]byte, error)
	DeleteEnvironment(string, string) ([]byte, error)
	GetDeploymentLog(string) ([]byte, error)
	GetEnvironmentInfo(string, string) ([]byte, error)
	ListEnvironmentVariables(string, string, bool) ([]byte, error)
	GetEnvironmentDeployments(string, string) ([]byte, error)
	GetEnvironmentTasks(string, string) ([]byte, error)
	RunDrushArchiveDump(string, string) ([]byte, error)
	RunDrushSQLDump(string, string) ([]byte, error)
	RunDrushCacheClear(string, string) ([]byte, error)
	RunCustomTask(string, string, api.Task) ([]byte, error)
	AddEnvironmentVariableToEnvironment(string, string, api.EnvVariable) ([]byte, error)
	DeleteEnvironmentVariableFromEnvironment(string, string, api.EnvVariable) ([]byte, error)
	PromoteEnvironment(string, string, string) ([]byte, error)
	InvokeAdvancedTaskDefinition(string, string, string) ([]byte, error)
	ListInvokableAdvancedTaskDefinitions(string, string) ([]byte, error)
}

// New .
func New(lc *lagoon.Config, debug bool) (Client, error) {
	lagoonAPI, err := graphql.LagoonAPI(lc, debug)
	if err != nil {
		return &Environments{}, err
	}
	return &Environments{
		debug: debug,
		api:   lagoonAPI,
	}, nil

}

var noDataError = "no data returned from the lagoon api"

// DeployEnvironmentBranch .
func (e *Environments) DeployEnvironmentBranch(projectName string, branchName string) ([]byte, error) {
	customRequest := api.CustomRequest{
		Query: `mutation ($project: String!, $branch: String!){
			deployEnvironmentBranch(
				input: {
					project:{name: $project}
					branchName: $branch
				}
			)
		}`,
		Variables: map[string]interface{}{
			"project": projectName,
			"branch":  branchName,
		},
		MappedResult: "deployEnvironmentBranch",
	}
	returnResult, err := e.api.Request(customRequest)
	return returnResult, err
}

// DeleteEnvironment .
func (e *Environments) DeleteEnvironment(projectName string, environmentName string) ([]byte, error) {
	evironment := api.DeleteEnvironment{
		Name:    environmentName,
		Project: projectName,
		Execute: true,
	}
	returnResult, err := e.api.DeleteEnvironment(evironment)
	return returnResult, err
}

// GetEnvironmentInfo will get basic info about a project
func (e *Environments) GetEnvironmentInfo(projectName string, environmentName string) ([]byte, error) {
	// get project info from lagoon
	project := api.Project{
		Name: projectName,
	}
	projectByName, err := e.api.GetProjectByName(project, graphql.ProjectByNameFragment)
	if err != nil {
		return []byte(""), err
	}
	var projectInfo api.Project
	err = json.Unmarshal([]byte(projectByName), &projectInfo)
	if err != nil {
		return []byte(""), err
	}
	// get the environment info from lagoon, we need the environment ID for later
	// we consume the project ID here
	environment := api.EnvironmentByName{
		Name:    environmentName,
		Project: projectInfo.ID,
	}
	environmentByName, err := e.api.GetEnvironmentByName(environment, graphql.EnvironmentByNameFragment)
	if err != nil {
		return []byte(""), err
	}
	var environmentInfo api.Environment
	err = json.Unmarshal([]byte(environmentByName), &environmentInfo)
	if err != nil {
		return []byte(""), err
	}

	returnResult, err := processEnvInfo(environmentByName)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

func processEnvInfo(projectByName []byte) ([]byte, error) {
	var environment api.Environment
	err := json.Unmarshal([]byte(projectByName), &environment)
	if err != nil {
		return []byte(""), err
	}
	environmentData := processEnvExtra(environment)
	var data []output.Data
	data = append(data, environmentData)
	dataMain := output.Table{
		Header: []string{"ID", "EnvironmentName", "EnvironmentType", "DeployType", "Created", "OpenshiftProjectName", "Route", "Routes", "AutoIdle", "DeployTitle", "DeployBaseRef", "DeployHeadRef"},
		Data:   data,
	}
	return json.Marshal(dataMain)
}

func processEnvExtra(environment api.Environment) []string {
	envID := returnNonEmptyString(strconv.Itoa(environment.ID))
	envName := returnNonEmptyString(string(environment.Name))
	envEnvironmentType := returnNonEmptyString(string(environment.EnvironmentType))
	envDeployType := returnNonEmptyString(string(environment.DeployType))
	envCreated := returnNonEmptyString(string(environment.Created))
	envOpenshiftProjectName := returnNonEmptyString(string(environment.OpenshiftProjectName))
	envRoute := returnNonEmptyString(string(environment.Route))
	envRoutes := returnNonEmptyString(string(environment.Routes))
	envDeployTitle := returnNonEmptyString(string(environment.DeployTitle))
	envDeployBaseRef := returnNonEmptyString(string(environment.DeployBaseRef))
	envDeployHeadRef := returnNonEmptyString(string(environment.DeployHeadRef))
	envAutoIdle := *environment.AutoIdle
	data := []string{
		fmt.Sprintf("%v", envID),
		fmt.Sprintf("%v", envName),
		fmt.Sprintf("%v", envEnvironmentType),
		fmt.Sprintf("%v", envDeployType),
		fmt.Sprintf("%v", envCreated),
		fmt.Sprintf("%v", envOpenshiftProjectName),
		fmt.Sprintf("%v", envRoute),
		fmt.Sprintf("%v", envRoutes),
		fmt.Sprintf("%v", envAutoIdle),
		fmt.Sprintf("%v", envDeployTitle),
		fmt.Sprintf("%v", envDeployBaseRef),
		fmt.Sprintf("%v", envDeployHeadRef),
	}
	return data
}

func returnNonEmptyString(value string) string {
	if len(value) == 0 {
		value = "-"
	}
	return value
}

// PromoteEnvironment .
func (e *Environments) PromoteEnvironment(projectName string, sourceEnv string, destEnv string) ([]byte, error) {
	customRequest := api.CustomRequest{
		Query: `mutation deployEnvironmentPromote ($project: String!, $sourceEnv: String!, $destEnv: String!){
		deployEnvironmentPromote(input:{
			sourceEnvironment:{
				name: $sourceEnv
				project:{
					name: $project
				}
			}
			project:{
				name: $project
			}
			destinationEnvironment: $destEnv
			})
		}`,
		Variables: map[string]interface{}{
			"project":   projectName,
			"sourceEnv": sourceEnv,
			"destEnv":   destEnv,
		},
		MappedResult: "deployEnvironmentPromote",
	}
	returnResult, err := e.api.Request(customRequest)
	return returnResult, err
}
