package projects

import (
	"fmt"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/graphql"
	"github.com/amazeeio/lagoon-cli/output"

	"encoding/json"
)

// ListAllProjectRocketChats will list all rocketchat notifications for a project
func ListAllProjectRocketChats(projectName string) ([]byte, error) {
	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}

	// get rocketchat project info from lagoon
	project := api.Project{
		Name: projectName,
	}
	projectRocketChats, err := lagoonAPI.GetRocketChatInfoForProject(project, graphql.RocketChatFragment)
	if err != nil {
		return []byte(""), err
	}
	var rocketChats api.RocketChats
	err = json.Unmarshal([]byte(projectRocketChats), &rocketChats)
	if err != nil {
		return []byte(""), err
	}
	// process the data for output
	data := []output.Data{}
	for _, rocketchat := range rocketChats.RocketChats {
		data = append(data, []string{
			fmt.Sprintf("%d", rocketchat.ID),
			rocketchat.Name,
			rocketchat.Channel,
			rocketchat.Webhook,
		})
	}
	dataMain := output.Table{
		Header: []string{"ID", "Name", "Channel", "Webhook"},
		Data:   data,
	}
	returnJSON, err := json.Marshal(dataMain)
	if err != nil {
		return []byte(""), err
	}
	return returnJSON, nil
}

// ListAllProjectSlacks will list all rocketchat notifications for a project
func ListAllProjectSlacks(projectName string) ([]byte, error) {
	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}

	// get rocketchat project info from lagoon
	project := api.Project{
		Name: projectName,
	}
	projectSlacks, err := lagoonAPI.GetSlackInfoForProject(project, graphql.SlackFragment)
	if err != nil {
		return []byte(""), err
	}
	var slacks api.Slacks
	err = json.Unmarshal([]byte(projectSlacks), &slacks)
	if err != nil {
		return []byte(""), err
	}
	// process the data for output
	data := []output.Data{}
	for _, slack := range slacks.Slacks {
		data = append(data, []string{
			fmt.Sprintf("%d", slack.ID),
			slack.Name,
			slack.Channel,
			slack.Webhook,
		})
	}
	dataMain := output.Table{
		Header: []string{"ID", "Name", "Channel", "Webhook"},
		Data:   data,
	}
	returnJSON, err := json.Marshal(dataMain)
	if err != nil {
		return []byte(""), err
	}
	return returnJSON, nil
}
