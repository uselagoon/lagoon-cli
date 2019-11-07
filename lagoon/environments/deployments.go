package environments

import (
	"encoding/json"
	"errors"

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
		data = append(data, []string{
			deployment.RemoteID,
			deployment.Name,
			string(deployment.Status),
			string(deployment.Created),
			string(deployment.Started),
			string(deployment.Completed),
		})
	}
	dataMain := output.Table{
		Header: []string{"RemoteID", "Name", "Status", "Created", "Started", "Completed"},
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
