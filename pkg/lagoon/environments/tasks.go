package environments

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/uselagoon/lagoon-cli/pkg/api"
	"github.com/uselagoon/lagoon-cli/pkg/graphql"
	"github.com/uselagoon/lagoon-cli/pkg/output"
)

// RunDrushArchiveDump will trigger a drush archive dump task
func (e *Environments) RunDrushArchiveDump(projectName string, environmentName string) ([]byte, error) {
	// get project info from lagoon, we need the project ID for later
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
	environmentByName, err := e.api.GetEnvironmentByName(environment, "")
	if err != nil {
		return []byte(""), err
	}
	var environmentInfo api.Environment
	err = json.Unmarshal([]byte(environmentByName), &environmentInfo)
	if err != nil {
		return []byte(""), err
	}

	// run the query to add the environment variable to lagoon
	customReq := api.CustomRequest{
		Query: `mutation runArdTask ($environment: Int!) {
			taskDrushArchiveDump(environment: $environment) {
				id
			}
		}`,
		Variables: map[string]interface{}{
			"environment": environmentInfo.ID,
		},
		MappedResult: "taskDrushArchiveDump",
	}
	returnResult, err := e.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// RunDrushSQLDump will trigger a drush archive dump task
func (e *Environments) RunDrushSQLDump(projectName string, environmentName string) ([]byte, error) {
	// get project info from lagoon, we need the project ID for later
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
	environmentByName, err := e.api.GetEnvironmentByName(environment, "")
	if err != nil {
		return []byte(""), err
	}
	var environmentInfo api.Environment
	err = json.Unmarshal([]byte(environmentByName), &environmentInfo)
	if err != nil {
		return []byte(""), err
	}

	// run the query to add the environment variable to lagoon
	customReq := api.CustomRequest{
		Query: `mutation runSqlDump ($environment: Int!) {
			taskDrushSqlDump(environment: $environment) {
				id
			}
		}`,
		Variables: map[string]interface{}{
			"environment": environmentInfo.ID,
		},
		MappedResult: "taskDrushSqlDump",
	}
	returnResult, err := e.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// RunDrushCacheClear will trigger a drush archive dump task
func (e *Environments) RunDrushCacheClear(projectName string, environmentName string) ([]byte, error) {
	// get project info from lagoon, we need the project ID for later
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
	environmentByName, err := e.api.GetEnvironmentByName(environment, "")
	if err != nil {
		return []byte(""), err
	}
	var environmentInfo api.Environment
	err = json.Unmarshal([]byte(environmentByName), &environmentInfo)
	if err != nil {
		return []byte(""), err
	}

	// run the query to add the environment variable to lagoon
	customReq := api.CustomRequest{
		Query: `mutation runCacheClear ($environment: Int!) {
			taskDrushCacheClear(environment: $environment) {
				id
			}
		}`,
		Variables: map[string]interface{}{
			"environment": environmentInfo.ID,
		},
		MappedResult: "taskDrushCacheClear",
	}
	returnResult, err := e.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// GetEnvironmentTasks .
func (e *Environments) GetEnvironmentTasks(projectName string, environmentName string) ([]byte, error) {
	// get project info from lagoon, we need the project ID for later
	project := api.Project{
		Name: projectName,
	}
	projectByName, err := e.api.GetProjectByName(project, graphql.ProjectNameID)
	if err != nil {
		return []byte(""), err
	}
	var projectInfo api.Project
	err = json.Unmarshal([]byte(projectByName), &projectInfo)
	if err != nil {
		return []byte(""), err
	}

	customRequest := api.CustomRequest{
		Query: `query ($project: Int!, $name: String!){
			environmentByName(
					project: $project
					name: $name
			){
				tasks{
					name
					id
					remoteId
					status
					created
					started
					completed
				}
			}
		}`,
		Variables: map[string]interface{}{
			"name":    environmentName,
			"project": projectInfo.ID,
		},
		MappedResult: "environmentByName",
	}
	environmentByName, err := e.api.Request(customRequest)
	if err != nil {
		return []byte(""), err
	}
	returnResult, err := processEnvironmentTasks(environmentByName)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

func processEnvironmentTasks(environmentByName []byte) ([]byte, error) {
	var environment api.Environment
	err := json.Unmarshal([]byte(environmentByName), &environment)
	if err != nil {
		return []byte(""), errors.New(noDataError) // @TODO could be a permissions thing when no data is returned
	}
	// process the data for output
	data := []output.Data{}
	for _, task := range environment.Tasks {
		remoteID := returnNonEmptyString(task.RemoteID)
		taskID := returnNonEmptyString(strconv.Itoa(task.ID))
		taskName := returnNonEmptyString(strings.Replace(task.Name, " ", "_", -1)) //remove spaces to make friendly for parsing with awk
		taskStatus := returnNonEmptyString(string(task.Status))
		taskCreated := returnNonEmptyString(string(task.Created))
		taskStarted := returnNonEmptyString(string(task.Started))
		taskComplete := returnNonEmptyString(string(task.Completed))
		taskService := returnNonEmptyString(task.Service)
		data = append(data, []string{
			taskID,
			remoteID,
			taskName,
			taskStatus,
			taskCreated,
			taskStarted,
			taskComplete,
			taskService,
		})
	}
	dataMain := output.Table{
		Header: []string{"ID", "RemoteID", "Name", "Status", "Created", "Started", "Completed", "Service"},
		Data:   data,
	}
	return json.Marshal(dataMain)
}

// RunCustomTask will trigger a drush archive dump task
func (e *Environments) RunCustomTask(projectName string, environmentName string, task api.Task) ([]byte, error) {
	// get project info from lagoon, we need the project ID for later
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
	environmentByName, err := e.api.GetEnvironmentByName(environment, "")
	if err != nil {
		return []byte(""), err
	}
	var environmentInfo api.Environment
	err = json.Unmarshal([]byte(environmentByName), &environmentInfo)
	if err != nil {
		return []byte(""), err
	}

	// run the query to add the environment variable to lagoon
	customReq := api.CustomRequest{
		Query: `mutation addTask ($environment: Int!, $name: String!, $command: String!, $service: String!) {
			addTask(input:{
			environment: $environment
			command: $command
			execute:true
			name: $name
			service: $service
		  }) {
				id
			}
		}`,
		Variables: map[string]interface{}{
			"environment": environmentInfo.ID,
			"name":        task.Name,
			"service":     task.Service,
			"command":     task.Command,
		},
		MappedResult: "addTask",
	}
	returnResult, err := e.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

const environmentWithTasksFragment = `fragment Environment on Environment {
	id
	name
	route
	routes
	deployType
	environmentType
	openshiftProjectName
	updated
	created
	deleted
    advancedTasks {
      ... on AdvancedTaskDefinitionCommand {
        id
        name
        description
      }
      ... on AdvancedTaskDefinitionImage {
        id
        name
        description
      }
    }
}`

// ListInvokableAdvancedTaskDefinitions returns a list of tasks invokable against an environment
func (e *Environments) ListInvokableAdvancedTaskDefinitions(projectName string, environmentName string) ([]byte, error) {
	// get project info from lagoon, we need the project ID for later
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

	environmentByName, err := e.api.GetEnvironmentByName(environment, environmentWithTasksFragment)
	if err != nil {
		return []byte(""), err
	}
	var environmentInfo api.Environment
	err = json.Unmarshal([]byte(environmentByName), &environmentInfo)
	if err != nil {
		return []byte(""), err
	}

	retvar, _ := json.Marshal(environmentInfo.AdvancedTasks)
	return retvar, nil
}

// InvokeAdvancedTaskDefinition will attempt to invoke an advanced task definition on an environment
func (e *Environments) InvokeAdvancedTaskDefinition(projectName string, environmentName string, advancedTaskName string) ([]byte, error) {
	// get project info from lagoon, we need the project ID for later
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

	environmentByName, err := e.api.GetEnvironmentByName(environment, environmentWithTasksFragment)
	if err != nil {
		return []byte(""), err
	}
	var environmentInfo api.Environment
	err = json.Unmarshal([]byte(environmentByName), &environmentInfo)
	if err != nil {
		return []byte(""), err
	}

	var taskId int
	for _, task := range environmentInfo.AdvancedTasks {
		if advancedTaskName == task.Name {
			taskId = task.ID
		}
	}

	if taskId == 0 {
		return nil, errors.New(fmt.Sprintf("Could not find a task `%v` for project/environment %v/%v",
			advancedTaskName, projectName, environmentName))
	}

	// run the query to add the environment variable to lagoon
	customReq := api.CustomRequest{
		Query: `mutation invokeRegisteredTask ($environment: Int!, $advancedTaskDefinition: Int!) {
			invokeRegisteredTask(advancedTaskDefinition: $advancedTaskDefinition, environment: $environment) {
				id
				name
				status
			}
		}`,
		Variables: map[string]interface{}{
			"advancedTaskDefinition": taskId,
			"environment":            environmentInfo.ID,
		},
		MappedResult: "invokeRegisteredTask",
	}
	returnResult, err := e.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}
