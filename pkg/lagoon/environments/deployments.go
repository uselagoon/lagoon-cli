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

// GetEnvironmentDeployments .
func (e *Environments) GetEnvironmentDeployments(projectName string, environmentName string) ([]byte, error) {
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
				deployments{
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
	returnResult, err := processEnvironmentDeployments(environmentByName)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

func processEnvironmentDeployments(environmentByName []byte) ([]byte, error) {
	var projects api.Project
	err := json.Unmarshal([]byte(environmentByName), &projects)
	if err != nil {
		return []byte(""), errors.New(noDataError) // @TODO could be a permissions thing when no data is returned
	}
	// process the data for output
	data := []output.Data{}
	for _, deployment := range projects.Deployments {
		deploymentID := returnNonEmptyString(strconv.Itoa(deployment.ID))
		remoteID := returnNonEmptyString(deployment.RemoteID)
		deploymentName := returnNonEmptyString(strings.Replace(deployment.Name, " ", "_", -1)) //remove spaces to make friendly for parsing with awk
		deploymentStatus := returnNonEmptyString(string(deployment.Status))
		deploymentCreated := returnNonEmptyString(string(deployment.Created))
		deploymentStarted := returnNonEmptyString(string(deployment.Started))
		deploymentComplete := returnNonEmptyString(string(deployment.Completed))
		data = append(data, []string{
			deploymentID,
			remoteID,
			deploymentName,
			deploymentStatus,
			deploymentCreated,
			deploymentStarted,
			deploymentComplete,
		})
	}
	dataMain := output.Table{
		Header: []string{"ID", "RemoteID", "Name", "Status", "Created", "Started", "Completed"},
		Data:   data,
	}
	return json.Marshal(dataMain)
}

// GetDeploymentLog .
func (e *Environments) GetDeploymentLog(deploymentID string) ([]byte, error) {
	customRequest := api.CustomRequest{
		Query: `query ($id: String!){
			deploymentByRemoteId(
					id: $id
			){
				id
				buildLog
			}
		}`,
		Variables: map[string]interface{}{
			"id": deploymentID,
		},
		MappedResult: "deploymentByRemoteId",
	}
	deploymentByRemoteID, err := e.api.Request(customRequest)
	return deploymentByRemoteID, err
}
