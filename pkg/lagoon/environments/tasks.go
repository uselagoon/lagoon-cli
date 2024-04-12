package environments

import (
	"encoding/json"
	"errors"
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
