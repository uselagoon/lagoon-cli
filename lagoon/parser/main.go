package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/lagoon/importer"
	"github.com/ghodss/yaml"
)

type lagoonImport struct {
	Data struct {
		AllProjects []importer.ExtendedProject `json:"allProjects"`
	} `json:"data"`
}

// ParseJSONToImport .
func ParseJSONToImport(jsonFile string) importer.LagoonImport {
	jsonStr, err := ioutil.ReadFile(jsonFile) // just pass the file name
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var lagoonImport lagoonImport
	var returnLagoonImport importer.LagoonImport
	json.Unmarshal([]byte(jsonStr), &lagoonImport)
	var lagoonUsers []importer.LagoonUsers
	for ind, project := range lagoonImport.Data.AllProjects {
		var lagoonProject importer.ExtendedProject
		projectBytes, _ := json.Marshal(project)
		json.Unmarshal(projectBytes, &lagoonProject)
		var projectPatch importer.ExtendedProject
		json.Unmarshal(projectBytes, &projectPatch)
		projectPatch.Groups = nil
		projectPatch.Notifications = nil
		returnLagoonImport.Projects = append(returnLagoonImport.Projects, importer.LagoonProjects{Project: projectPatch})
		for _, k := range lagoonProject.Groups {
			returnLagoonImport.Projects[ind].Groups = appendIfMissingGroup(returnLagoonImport.Projects[ind].Groups, k.Name)
			for _, m := range k.Members {
				lagoonUser := importer.LagoonUsers{
					User: importer.LagoonUser{
						Email: m.User.Email,
					},
				}
				lagoonUserGroupRole := importer.AddUserToGroup{
					Name: k.Name,
					Role: k.Role,
				}
				lagoonUser.Groups = appendIfMissingGroups(lagoonUser.Groups, lagoonUserGroupRole)
				lagoonUsers = appendIfMissingUsers(lagoonUsers, lagoonUser)
			}
		}
		// returnLagoonImport.Users = lagoonUser
		// for _, u := range lagoonUsers {
		returnLagoonImport.Users = lagoonUsers //appendIfMissingUsers(returnLagoonImport.Users, u)
		// }
		for _, k := range lagoonProject.Groups {
			// returnLagoonImport.Groups = appendIfMissingGroups(returnLagoonImport.Groups, k)
			returnLagoonImport.Groups = append(returnLagoonImport.Groups, api.Group{Name: k.Name})
		}
		for _, k := range project.Notifications {
			// fmt.Println(k)
			var notification struct {
				TypeName     string `json:"__typename"`
				Name         string `json:"name"`
				Webhook      string `json:"webhook,omitempty"`
				Channel      string `json:"channel,omitempty"`
				EmailAddress string `json:"emailAddress,omitempty"`
			}
			notifBytes, _ := json.Marshal(k)
			json.Unmarshal(notifBytes, &notification)
			switch notification.TypeName {
			case "NotificationRocketChat":
				var rocketNotification api.NotificationRocketChat
				notifBytes, _ := json.Marshal(notification)
				json.Unmarshal(notifBytes, &rocketNotification)
				returnLagoonImport.RocketChat = appendIfMissingRocket(returnLagoonImport.RocketChat, rocketNotification)
				returnLagoonImport.Projects[ind].Notifications.RocketChat = append(returnLagoonImport.Projects[ind].Notifications.RocketChat, notification.Name)
			case "NotificationSlack":
				var slackNotification api.NotificationSlack
				notifBytes, _ := json.Marshal(notification)
				json.Unmarshal(notifBytes, &slackNotification)
				returnLagoonImport.Slack = appendIfMissingSlack(returnLagoonImport.Slack, slackNotification)
				returnLagoonImport.Projects[ind].Notifications.Slack = append(returnLagoonImport.Projects[ind].Notifications.Slack, notification.Name)
			case "NotificationEmail":
				var emailNotification api.NotificationEmail
				notifBytes, _ := json.Marshal(notification)
				json.Unmarshal(notifBytes, &emailNotification)
				returnLagoonImport.Email = appendIfMissingEmail(returnLagoonImport.Email, emailNotification)
				returnLagoonImport.Projects[ind].Notifications.Email = append(returnLagoonImport.Projects[ind].Notifications.Email, notification.Name)
			case "NotificationMicrosoftTeams":
				var teamsNotification api.NotificationMicrosoftTeams
				notifBytes, _ := json.Marshal(notification)
				json.Unmarshal(notifBytes, &teamsNotification)
				returnLagoonImport.MicrosoftTeams = appendIfMissingTeams(returnLagoonImport.MicrosoftTeams, teamsNotification)
				returnLagoonImport.Projects[ind].Notifications.MicrosoftTeams = append(returnLagoonImport.Projects[ind].Notifications.MicrosoftTeams, notification.Name)
			}
		}
	}
	for _, k := range returnLagoonImport.Projects {
		fmt.Println(k.Notifications, k.Project.Name, k.Project.ID, k.Project.Name, k.Groups)
		// fmt.Println(k.Project.Name)
	}
	// for _, k := range returnLagoonImport.RocketChat {
	// 	fmt.Println(k)
	// }
	yamlBytes, _ := yaml.Marshal(returnLagoonImport)
	fmt.Println(string(yamlBytes))
	// fmt.Println(returnLagoonImport.Users)
	return returnLagoonImport
}

func appendIfMissingRocket(slice []api.NotificationRocketChat, i api.NotificationRocketChat) []api.NotificationRocketChat {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

func appendIfMissingSlack(slice []api.NotificationSlack, i api.NotificationSlack) []api.NotificationSlack {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

func appendIfMissingEmail(slice []api.NotificationEmail, i api.NotificationEmail) []api.NotificationEmail {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

func appendIfMissingTeams(slice []api.NotificationMicrosoftTeams, i api.NotificationMicrosoftTeams) []api.NotificationMicrosoftTeams {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

func appendIfMissingGroup(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

func appendIfMissingGroups(slice []importer.AddUserToGroup, i importer.AddUserToGroup) []importer.AddUserToGroup {
	for _, ele := range slice {
		if ele.Name == i.Name {
			return slice
		}
	}
	return append(slice, i)
}

func appendIfMissingUser(slice []importer.LagoonUser, i importer.LagoonUser) []importer.LagoonUser {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

func appendIfMissingUsers(slice []importer.LagoonUsers, i importer.LagoonUsers) []importer.LagoonUsers {
	for _, ele := range slice {
		if ele.User == i.User {
			return slice
		}
	}
	return append(slice, i)
}
