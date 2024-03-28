package environments

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/uselagoon/lagoon-cli/pkg/api"
	"github.com/uselagoon/lagoon-cli/pkg/output"
)

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
