package projects

import (
	"encoding/json"
	"fmt"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/graphql"
	"github.com/amazeeio/lagoon-cli/output"
)

// ListAllProjectRocketChats will list all rocketchat notifications for a project
func ListAllProjectRocketChats(projectName string) ([]byte, error) {
	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}
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
		Header: []string{"NID", "NotificationName", "Channel", "Webhook"},
		Data:   data,
	}
	returnResult, err := json.Marshal(dataMain)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// ListAllProjectSlacks will list all slack notifications for a project
func ListAllProjectSlacks(projectName string) ([]byte, error) {
	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}
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
		Header: []string{"NID", "NotificationName", "Channel", "Webhook"},
		Data:   data,
	}
	returnResult, err := json.Marshal(dataMain)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// ListAllSlacks will list all slack notifications on all projects
func ListAllSlacks() ([]byte, error) {
	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}
	customReq := api.CustomRequest{
		Query: `query {
			allProjects {
				name
				id
				notifications {
					...Notification
				}
			}
		}
		fragment Notification on NotificationSlack {
			id
			name
			webhook
			channel
		}`,
		Variables:    map[string]interface{}{},
		MappedResult: "allProjects",
	}
	returnResult, err := lagoonAPI.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	var projects []api.Project
	err = json.Unmarshal([]byte(returnResult), &projects)
	if err != nil {
		return []byte(""), err
	}
	// process the data for output
	data := []output.Data{}
	for _, project := range projects {
		for _, notif := range project.Notifications {
			var slack api.NotificationSlack
			returnResult, _ = json.Marshal(notif)
			err := json.Unmarshal([]byte(returnResult), &slack)
			if err != nil {
				return []byte(""), err
			}
			if slack.ID != 0 {
				// fmt.Println(slack.ID)
				data = append(data, []string{
					fmt.Sprintf("%d", slack.ID),
					project.Name,
					slack.Name,
					slack.Channel,
					slack.Webhook,
				})
			}
		}
	}
	dataMain := output.Table{
		Header: []string{"NID", "Project", "NotificationName", "Channel", "Webhook"},
		Data:   data,
	}
	returnResult, err = json.Marshal(dataMain)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// ListAllRocketChats will list all rocketchat notifications on all projects
func ListAllRocketChats() ([]byte, error) {
	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}
	customReq := api.CustomRequest{
		Query: `query {
			allProjects {
				name
				id
				notifications {
					...Notification
				}
			}
		}
		fragment Notification on NotificationRocketChat {
			id
			name
			webhook
			channel
		}`,
		Variables:    map[string]interface{}{},
		MappedResult: "allProjects",
	}
	returnResult, err := lagoonAPI.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	var projects []api.Project
	err = json.Unmarshal([]byte(returnResult), &projects)
	if err != nil {
		return []byte(""), err
	}
	// process the data for output
	data := []output.Data{}
	for _, project := range projects {
		for _, notif := range project.Notifications {
			var rocketchat api.NotificationRocketChat
			returnResult, _ = json.Marshal(notif)
			err := json.Unmarshal([]byte(returnResult), &rocketchat)
			if err != nil {
				return []byte(""), err
			}
			if rocketchat.ID != 0 {
				// fmt.Println(slack.ID)
				data = append(data, []string{
					fmt.Sprintf("%d", rocketchat.ID),
					project.Name,
					rocketchat.Name,
					rocketchat.Channel,
					rocketchat.Webhook,
				})
			}
		}
	}
	dataMain := output.Table{
		Header: []string{"NID", "Project", "NotificationName", "Channel", "Webhook"},
		Data:   data,
	}
	returnResult, err = json.Marshal(dataMain)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// AddSlackNotification will add a slack notification to lagoon to be used by a project
func AddSlackNotification(notificationName string, channel string, webhookURL string) ([]byte, error) {
	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}
	customReq := api.CustomRequest{
		Query: `mutation ($name: String!, $channel: String!, $webhook: String!) {
			addNotificationSlack(input:{name: $name, channel: $channel, webhook: $webhook}
			){
				id
			}
		}`,
		Variables: map[string]interface{}{
			"name":    notificationName,
			"channel": channel,
			"webhook": webhookURL,
		},
		MappedResult: "addNotificationSlack",
	}
	returnResult, err := lagoonAPI.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// AddSlackNotificationToProject will add a notification to a project
func AddSlackNotificationToProject(projectName string, notificationName string) ([]byte, error) {
	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}
	customReq := api.CustomRequest{
		Query: `mutation ($name: String!, $project: String!) {
			addNotificationToProject(input:{notificationName: $name, notificationType: SLACK, project: $project}
			){
				id
			}
		}`,
		Variables: map[string]interface{}{
			"name":    notificationName,
			"project": projectName,
		},
		MappedResult: "addNotificationToProject",
	}
	returnResult, err := lagoonAPI.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// DeleteSlackNotification will delete a slack notification from lagoon
func DeleteSlackNotification(notificationName string) ([]byte, error) {
	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}
	customReq := api.CustomRequest{
		Query: `mutation ($name: String!) {
			deleteNotificationSlack(input:{name: $name})
		}`,
		Variables: map[string]interface{}{
			"name": notificationName,
		},
		MappedResult: "deleteNotificationSlack",
	}
	returnResult, err := lagoonAPI.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// DeleteSlackNotificationFromProject will delete a slack notification from a project
func DeleteSlackNotificationFromProject(projectName string, notificationName string) ([]byte, error) {
	// set up a lagoonapi client
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
	customReq := api.CustomRequest{
		Query: `mutation ($name: String!, $project: String!) {
			removeNotificationFromProject(input:{notificationName: $name, project: $project, notificationType: SLACK})
			{
				id
			}
		}`,
		Variables: map[string]interface{}{
			"name":    notificationName,
			"project": projectName,
		},
		MappedResult: "removeNotificationFromProject",
	}
	returnResult, err := lagoonAPI.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// AddRocketChatNotification will add a rocketchat notification to lagoon to be used by a project
func AddRocketChatNotification(notificationName string, channel string, webhookURL string) ([]byte, error) {
	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}
	customReq := api.CustomRequest{
		Query: `mutation ($name: String!, $channel: String!, $webhook: String!) {
			addNotificationRocketChat(input:{name: $name, channel: $channel, webhook: $webhook}
			){
				id
			}
		}`,
		Variables: map[string]interface{}{
			"name":    notificationName,
			"channel": channel,
			"webhook": webhookURL,
		},
		MappedResult: "addNotificationRocketChat",
	}
	returnResult, err := lagoonAPI.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// AddRocketChatNotificationToProject will add a rocketchat notification to a project
func AddRocketChatNotificationToProject(projectName string, notificationName string) ([]byte, error) {
	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}
	customReq := api.CustomRequest{
		Query: `mutation ($name: String!, $project: String!) {
			addNotificationToProject(input:{notificationName: $name, notificationType: ROCKETCHAT, project: $project}
			){
				id
			}
		}`,
		Variables: map[string]interface{}{
			"name":    notificationName,
			"project": projectName,
		},
		MappedResult: "addNotificationToProject",
	}
	returnResult, err := lagoonAPI.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// DeleteRocketChatNotification will delete a rocketchat notification from lagoon
func DeleteRocketChatNotification(notificationName string) ([]byte, error) {
	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}
	customReq := api.CustomRequest{
		Query: `mutation ($name: String!) {
			deleteNotificationRocketChat(input:{name: $name})
		}`,
		Variables: map[string]interface{}{
			"name": notificationName,
		},
		MappedResult: "deleteNotificationRocketChat",
	}
	returnResult, err := lagoonAPI.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// DeleteRocketChatNotificationFromProject will delete a rocketchat notification from a project
func DeleteRocketChatNotificationFromProject(projectName string, notificationName string) ([]byte, error) {
	// set up a lagoonapi client
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
	customReq := api.CustomRequest{
		Query: `mutation ($name: String!, $project: String!) {
			removeNotificationFromProject(input:{notificationName: $name, project: $project, notificationType: ROCKETCHAT})
			{
				id
			}
		}`,
		Variables: map[string]interface{}{
			"name":    notificationName,
			"project": projectName,
		},
		MappedResult: "removeNotificationFromProject",
	}
	returnResult, err := lagoonAPI.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// UpdateSlackNotification will update an existing notification
func UpdateSlackNotification(notificationName string, jsonPatch string) ([]byte, error) {
	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}
	var updateSlack api.UpdateNotificationSlackPatch

	err = json.Unmarshal([]byte(jsonPatch), &updateSlack)
	customReq := api.CustomRequest{
		Query: `mutation ($name: String!, $patch: UpdateNotificationSlackPatchInput!) {
			updateNotificationSlack(input:{name: $name, patch: $patch}
			){
				id
			}
		}`,
		Variables: map[string]interface{}{
			"name":  notificationName,
			"patch": updateSlack,
		},
		MappedResult: "updateNotificationSlack",
	}
	returnResult, err := lagoonAPI.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// UpdateRocketChatNotification will update an existing notification
func UpdateRocketChatNotification(notificationName string, jsonPatch string) ([]byte, error) {
	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}
	var updateRocketChat api.UpdateNotificationRocketChatPatch
	err = json.Unmarshal([]byte(jsonPatch), &updateRocketChat)
	customReq := api.CustomRequest{
		Query: `mutation ($name: String!, $patch: UpdateNotificationRocketChatPatchInput!) {
			updateNotificationRocketChat(input:{name: $name, patch: $patch}
			){
				id
			}
		}`,
		Variables: map[string]interface{}{
			"name":  notificationName,
			"patch": updateRocketChat,
		},
		MappedResult: "updateNotificationRocketChat",
	}
	returnResult, err := lagoonAPI.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}
