package environments

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/graphql"
	"github.com/amazeeio/lagoon-cli/output"
)

// GetEnvironmentDeployments .
func GetEnvironmentDeployments(projectName string, environmentName string) ([]byte, error) {
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}

	// get project info from lagoon, we need the project ID for later
	project := api.Project{
		Name: projectName,
	}
	projectByName, err := lagoonAPI.GetProjectByName(project, graphql.ProjectNameID)
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
	environmentByName, err := lagoonAPI.Request(customRequest)
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
		return []byte(""), errors.New("no data returned from lagoon") // @TODO could be a permissions thing when no data is returned
	}
	// process the data for output
	data := []output.Data{}
	for _, deployment := range projects.Deployments {
		deploymentID := strconv.Itoa(deployment.ID)
		remoteID := deployment.RemoteID
		deploymentName := strings.Replace(deployment.Name, " ", "_", -1) //remove spaces to make friendly for parsing with awk
		deploymentStatus := string(deployment.Status)
		deploymentCreated := string(deployment.Created)
		deploymentStarted := string(deployment.Started)
		deploymentComplete := string(deployment.Completed)
		if len(remoteID) == 0 {
			remoteID = "-"
		}
		if len(deploymentID) == 0 {
			deploymentID = "-"
		}
		if len(deploymentStatus) == 0 {
			deploymentStatus = "-"
		}
		if len(deploymentCreated) == 0 {
			deploymentCreated = "-"
		}
		if len(deploymentStarted) == 0 {
			deploymentStarted = "-"
		}
		if len(deploymentComplete) == 0 {
			deploymentComplete = "-"
		}
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
func GetDeploymentLog(deploymentID string) ([]byte, error) {
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}

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
	deploymentByRemoteID, err := lagoonAPI.Request(customRequest)
	return deploymentByRemoteID, err
}
