package projects

import (
	"encoding/json"
	"fmt"

	"github.com/amazeeio/lagoon-cli/pkg/api"
	"github.com/amazeeio/lagoon-cli/pkg/graphql"
	"github.com/amazeeio/lagoon-cli/pkg/output"
)

// ListProjectRocketChats will list all rocketchat notifications for a project
func (p *Projects) ListProjectRocketChats(projectName string) ([]byte, error) {
	project := api.Project{
		Name: projectName,
	}
	projectRocketChats, err := p.api.GetRocketChatInfoForProject(project, graphql.RocketChatFragment)
	if err != nil {
		return []byte(""), err
	}
	returnResult, err := processProjectRocketChats(projectRocketChats)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// ListAllRocketChats will list all rocketchat notifications on all projects
func (p *Projects) ListAllRocketChats() ([]byte, error) {
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
	allRocketChats, err := p.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	returnResult, err := processAllRocketChats(allRocketChats)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

func processProjectRocketChats(allProjects []byte) ([]byte, error) {
	var rocketChats api.RocketChats
	err := json.Unmarshal([]byte(allProjects), &rocketChats)
	if err != nil {
		return []byte(""), err
	}
	// process the data for output
	data := []output.Data{}
	for _, rocketchat := range rocketChats.RocketChats {
		projectData := processRocketChat(rocketchat)
		data = append(data, projectData)
	}
	dataMain := output.Table{
		Header: []string{"NID", "NotificationName", "Channel", "Webhook"},
		Data:   data,
	}
	return json.Marshal(dataMain)
}

func processRocketChat(rocketchat api.NotificationRocketChat) []string {
	// count the current dev environments in a project
	data := []string{
		fmt.Sprintf("%d", rocketchat.ID),
		rocketchat.Name,
		rocketchat.Channel,
		rocketchat.Webhook,
	}
	return data
}

func processAllRocketChats(allProjects []byte) ([]byte, error) {
	var projects []api.Project
	err := json.Unmarshal([]byte(allProjects), &projects)
	if err != nil {
		return []byte(""), err
	}
	// process the data for output
	data := []output.Data{}
	for _, project := range projects {
		for _, notif := range project.Notifications {
			var rocketchat api.NotificationRocketChat
			rocketNotif, _ := json.Marshal(notif)
			err := json.Unmarshal([]byte(rocketNotif), &rocketchat)
			if err != nil {
				return []byte(""), err
			}
			if rocketchat.ID != 0 {
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
	return json.Marshal(dataMain)
}

// ListProjectSlacks will list all slack notifications for a project
func (p *Projects) ListProjectSlacks(projectName string) ([]byte, error) {
	project := api.Project{
		Name: projectName,
	}
	projectSlacks, err := p.api.GetSlackInfoForProject(project, graphql.SlackFragment)
	if err != nil {
		return []byte(""), err
	}
	returnResult, err := processProjectSlacks(projectSlacks)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// ListAllSlacks will list all slack notifications on all projects
func (p *Projects) ListAllSlacks() ([]byte, error) {
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
	allSlacks, err := p.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	returnResult, err := processAllSlacks(allSlacks)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

func processProjectSlacks(allProjects []byte) ([]byte, error) {
	var rocketChats api.Slacks
	err := json.Unmarshal([]byte(allProjects), &rocketChats)
	if err != nil {
		return []byte(""), err
	}
	// process the data for output
	data := []output.Data{}
	for _, slack := range rocketChats.Slacks {
		projectData := processSlack(slack)
		data = append(data, projectData)
	}
	dataMain := output.Table{
		Header: []string{"NID", "NotificationName", "Channel", "Webhook"},
		Data:   data,
	}
	return json.Marshal(dataMain)
}

func processSlack(rocketchat api.NotificationSlack) []string {
	// count the current dev environments in a project
	data := []string{
		fmt.Sprintf("%d", rocketchat.ID),
		rocketchat.Name,
		rocketchat.Channel,
		rocketchat.Webhook,
	}
	return data
}

func processAllSlacks(allProjects []byte) ([]byte, error) {
	var projects []api.Project
	err := json.Unmarshal([]byte(allProjects), &projects)
	if err != nil {
		return []byte(""), err
	}
	// process the data for output
	data := []output.Data{}
	for _, project := range projects {
		for _, notif := range project.Notifications {
			var slack api.NotificationSlack
			slackNotif, _ := json.Marshal(notif)
			err := json.Unmarshal([]byte(slackNotif), &slack)
			if err != nil {
				return []byte(""), err
			}
			if slack.ID != 0 {
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
	return json.Marshal(dataMain)
}

// AddSlackNotification will add a slack notification to lagoon to be used by a project
func (p *Projects) AddSlackNotification(notificationName string, channel string, webhookURL string) ([]byte, error) {
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
	returnResult, err := p.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// AddSlackNotificationToProject will add a notification to a project
func (p *Projects) AddSlackNotificationToProject(projectName string, notificationName string) ([]byte, error) {
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
	returnResult, err := p.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// DeleteSlackNotification will delete a slack notification from lagoon
func (p *Projects) DeleteSlackNotification(notificationName string) ([]byte, error) {
	customReq := api.CustomRequest{
		Query: `mutation ($name: String!) {
			deleteNotificationSlack(input:{name: $name})
		}`,
		Variables: map[string]interface{}{
			"name": notificationName,
		},
		MappedResult: "deleteNotificationSlack",
	}
	returnResult, err := p.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// DeleteSlackNotificationFromProject will delete a slack notification from a project
func (p *Projects) DeleteSlackNotificationFromProject(projectName string, notificationName string) ([]byte, error) {
	// get project info from lagoon, we need the project ID for later
	project := api.Project{
		Name: projectName,
	}
	projectByName, err := p.api.GetProjectByName(project, graphql.ProjectNameID)
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
	returnResult, err := p.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// AddRocketChatNotification will add a rocketchat notification to lagoon to be used by a project
func (p *Projects) AddRocketChatNotification(notificationName string, channel string, webhookURL string) ([]byte, error) {
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
	returnResult, err := p.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// AddRocketChatNotificationToProject will add a rocketchat notification to a project
func (p *Projects) AddRocketChatNotificationToProject(projectName string, notificationName string) ([]byte, error) {
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
	returnResult, err := p.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// DeleteRocketChatNotification will delete a rocketchat notification from lagoon
func (p *Projects) DeleteRocketChatNotification(notificationName string) ([]byte, error) {
	customReq := api.CustomRequest{
		Query: `mutation ($name: String!) {
			deleteNotificationRocketChat(input:{name: $name})
		}`,
		Variables: map[string]interface{}{
			"name": notificationName,
		},
		MappedResult: "deleteNotificationRocketChat",
	}
	returnResult, err := p.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// DeleteRocketChatNotificationFromProject will delete a rocketchat notification from a project
func (p *Projects) DeleteRocketChatNotificationFromProject(projectName string, notificationName string) ([]byte, error) {
	// get project info from lagoon, we need the project ID for later
	project := api.Project{
		Name: projectName,
	}
	projectByName, err := p.api.GetProjectByName(project, graphql.ProjectNameID)
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
	returnResult, err := p.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// UpdateSlackNotification will update an existing notification
func (p *Projects) UpdateSlackNotification(notificationName string, jsonPatch string) ([]byte, error) {
	var updateSlack api.UpdateNotificationSlackPatch
	err := json.Unmarshal([]byte(jsonPatch), &updateSlack)
	customReq := api.CustomRequest{
		Query: `mutation ($oldname: String!, $patch: UpdateNotificationSlackPatchInput!) {
			updateNotificationSlack(input:{name: $oldname, patch: $patch}
			){
				id
			}
		}`,
		Variables: map[string]interface{}{
			"oldname": notificationName,
			"patch":   updateSlack,
		},
		MappedResult: "updateNotificationSlack",
	}
	returnResult, err := p.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// UpdateRocketChatNotification will update an existing notification
func (p *Projects) UpdateRocketChatNotification(notificationName string, jsonPatch string) ([]byte, error) {
	var updateRocketChat api.UpdateNotificationRocketChatPatch
	err := json.Unmarshal([]byte(jsonPatch), &updateRocketChat)
	customReq := api.CustomRequest{
		Query: `mutation ($oldname: String!, $patch: UpdateNotificationRocketChatPatchInput!) {
			updateNotificationRocketChat(input:{name: $oldname, patch: $patch}
			){
				id
			}
		}`,
		Variables: map[string]interface{}{
			"oldname": notificationName,
			"patch":   updateRocketChat,
		},
		MappedResult: "updateNotificationRocketChat",
	}
	returnResult, err := p.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}
