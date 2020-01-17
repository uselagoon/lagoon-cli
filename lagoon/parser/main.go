package parser

/*
Usage:

lagoon export -p $projectname --force
lagoon export --force

lagoon parse -I /path/to/file.json

*/

import (
	"encoding/json"
	"fmt"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/graphql"
	"github.com/amazeeio/lagoon-cli/lagoon/importer"
	"github.com/ghodss/yaml"
)

// Parser .
type Parser struct {
	debug bool
	api   api.Client
}

// SkipExport .
type SkipExport struct {
	Users         bool
	Groups        bool
	Notifications bool
	Slack         bool
	RocketChat    bool
}

// Client .
type Client interface {
	ParseJSONImport(string) importer.LagoonImport
	ParseProject(string, SkipExport) ([]byte, error)
	ParseAllProjects(SkipExport) ([]byte, error)
}

// New .
func New(debug bool) (Client, error) {
	lagoonAPI, err := graphql.LagoonAPI(debug)
	if err != nil {
		return &Parser{}, err
	}
	return &Parser{
		debug: debug,
		api:   lagoonAPI,
	}, nil
}

type lagoonImport struct {
	Data map[string]interface{} `json:"data"`
}

// ParseJSONImport given a file path that contains a full all projects data dump from lagoon, parse it into something that the importer can use
/*
{
	"data": {
		"allProjects": [
			{
				"name": "credentialstest-project1",
			},
			{
				"name": "credentialstest-project1",
			}
		]
	}
}
*/
func (p *Parser) ParseJSONImport(jsonData string) importer.LagoonImport {
	// jsonStr, err := ioutil.ReadFile(jsonFile) // just pass the file name
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	var lagoonImporter lagoonImport
	var returnLagoonImport importer.LagoonImport
	json.Unmarshal([]byte(jsonData), &lagoonImporter)
	var lagoonData []importer.ExtendedProject
	lagoonDataBytes, _ := json.Marshal(lagoonImporter)
	json.Unmarshal(lagoonDataBytes, &lagoonData)
	for _, projects := range lagoonImporter.Data {
		lagoonDataBytes, _ := json.Marshal(projects)
		skip := SkipExport{
			Users:         false,
			Groups:        false,
			Notifications: false,
			Slack:         false,
			RocketChat:    false,
		}
		yamlBytes := processParser(lagoonDataBytes, skip)
		fmt.Println(string(yamlBytes))
		return returnLagoonImport
	}
	return returnLagoonImport
}

func processParser(lagoonDataBytes []byte, skip SkipExport) []byte {
	var returnLagoonImport importer.LagoonImport
	var lagoonUsers []importer.LagoonUsers
	var lagoonData []importer.ExtendedProject
	json.Unmarshal(lagoonDataBytes, &lagoonData)
	for ind, project := range lagoonData {
		var lagoonProject importer.ExtendedProject
		projectBytes, _ := json.Marshal(project)
		json.Unmarshal(projectBytes, &lagoonProject)
		var projectPatch api.ProjectPatch
		json.Unmarshal(projectBytes, &projectPatch)
		returnLagoonImport.Projects = append(returnLagoonImport.Projects, importer.LagoonProjects{Project: projectPatch})
		if !skip.Users {
			for _, k := range lagoonProject.Groups {
				returnLagoonImport.Projects[ind].Groups = appendIfMissingGroup(returnLagoonImport.Projects[ind].Groups, k.Name)
				for _, m := range k.Members {
					var userKeys []importer.LagoonUserSSHKeys
					for _, key := range m.User.SSHKeys {
						userKeys = append(userKeys, importer.LagoonUserSSHKeys{SSHKey: string(key.KeyType) + " " + key.KeyValue, KeyName: key.Name})
					}
					lagoonUser := importer.LagoonUsers{
						User: importer.LagoonUser{
							Email:   m.User.Email,
							SSHKeys: userKeys,
						},
					}
					lagoonUserGroupRole := importer.AddUserToGroup{
						Name: k.Name,
						Role: m.Role,
					}
					lagoonUser.Groups = appendIfMissingGroups(lagoonUser.Groups, lagoonUserGroupRole)
					lagoonUsers = appendIfMissingUsers(lagoonUsers, lagoonUser)
				}
			}
		}
		returnLagoonImport.Users = lagoonUsers
		if !skip.Groups {
			for _, k := range lagoonProject.Groups {
				returnLagoonImport.Groups = appendIfMissingGroups2(returnLagoonImport.Groups, api.Group{Name: k.Name})
			}
		}
		if !skip.Notifications {
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
					if !skip.RocketChat {
						var rocketNotification api.NotificationRocketChat
						notifBytes, _ := json.Marshal(notification)
						json.Unmarshal(notifBytes, &rocketNotification)
						returnLagoonImport.Notifications.RocketChat = appendIfMissingRocket(returnLagoonImport.Notifications.RocketChat, rocketNotification)
						returnLagoonImport.Projects[ind].Notifications.RocketChat = append(returnLagoonImport.Projects[ind].Notifications.RocketChat, notification.Name)
					}
				case "NotificationSlack":
					if !skip.Slack {
						var slackNotification api.NotificationSlack
						notifBytes, _ := json.Marshal(notification)
						json.Unmarshal(notifBytes, &slackNotification)
						returnLagoonImport.Notifications.Slack = appendIfMissingSlack(returnLagoonImport.Notifications.Slack, slackNotification)
						returnLagoonImport.Projects[ind].Notifications.Slack = append(returnLagoonImport.Projects[ind].Notifications.Slack, notification.Name)
					}
					// @TODO: enable once 1.2.0+ lagoon is more widespread
					// case "NotificationEmail":
					// 	var emailNotification api.NotificationEmail
					// 	notifBytes, _ := json.Marshal(notification)
					// 	json.Unmarshal(notifBytes, &emailNotification)
					// 	returnLagoonImport.Notifications.Email = appendIfMissingEmail(returnLagoonImport.Notifications.Email, emailNotification)
					// 	returnLagoonImport.Projects[ind].Notifications.Email = append(returnLagoonImport.Projects[ind].Notifications.Email, notification.Name)
					// case "NotificationMicrosoftTeams":
					// 	var teamsNotification api.NotificationMicrosoftTeams
					// 	notifBytes, _ := json.Marshal(notification)
					// 	json.Unmarshal(notifBytes, &teamsNotification)
					// 	returnLagoonImport.Notifications.MicrosoftTeams = appendIfMissingTeams(returnLagoonImport.Notifications.MicrosoftTeams, teamsNotification)
					// 	returnLagoonImport.Projects[ind].Notifications.MicrosoftTeams = append(returnLagoonImport.Projects[ind].Notifications.MicrosoftTeams, notification.Name)
				}
			}
		}
	}
	yamlBytes, _ := yaml.Marshal(returnLagoonImport)
	return yamlBytes
}

// A bunch of append if missing funcs for different types
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

func appendIfMissingGroups2(slice []api.Group, i api.Group) []api.Group {
	for _, ele := range slice {
		if ele.Name == i.Name {
			return slice
		}
	}
	return append(slice, i)
}

func appendIfMissingUser(slice []importer.LagoonUser, i importer.LagoonUser) []importer.LagoonUser {
	for _, ele := range slice {
		if ele.Email == i.Email {
			return slice
		}
	}
	return append(slice, i)
}

func appendIfMissingUsers(slice []importer.LagoonUsers, i importer.LagoonUsers) []importer.LagoonUsers {
	for _, ele := range slice {
		if ele.User.Email == i.User.Email {
			return slice
		}
	}
	return append(slice, i)
}

// ParseProject given a specific project name, get the json dump, then parse it to the import format
func (p *Parser) ParseProject(projectName string, skip SkipExport) ([]byte, error) {
	customReq := api.CustomRequest{
		Query: `fragment NotificationSlack on NotificationSlack {
			webhook
		  name
			channel
		}
		fragment NotificationRocket on NotificationRocketChat {
			webhook
		  name
			channel
		}
		# @TODO: enable once 1.2.0+ lagoon is more widespread
		# fragment NotificationEmail on NotificationEmail {
		# 	emailAddress
		#   name
		# }
		# fragment NotificationTeams on NotificationMicrosoftTeams {
		# 	webhook
		#   name
		# }
		query Projects ($name: String!) {
			projectByName(name: $name) {
			name
			autoIdle
			branches
			pullrequests
			privateKey
			productionEnvironment
			activeSystemsDeploy
			activeSystemsTask
			activeSystemsRemove
			activeSystemsPromote
			storageCalc
			openshiftProjectPattern
			developmentEnvironmentsLimit
			gitUrl
			autoIdle
			groups{
			  name
			  id
			  members{
				user{
				  email
				  sshKeys{
					name
					keyType
					keyValue
				  }
				  firstName
				  lastName
				}
				role
			  }
			}
			notifications {
			  __typename
			  ...NotificationRocket
			  ...NotificationSlack
			  # @TODO: enable once 1.2.0+ lagoon is more widespread
			  # ...NotificationEmail
			  # ...NotificationTeams
			}
			openshift{
			  id
			}
			envVariables {
			  name
			  scope
			  value
			}
			environments {
			  id
			  name
			  openshiftProjectName
			  autoIdle
			  envVariables {
				name
				scope
				value
			  }
			}
		  }
		}`,
		Variables: map[string]interface{}{
			"name": projectName,
		},
		MappedResult: "projectByName",
	}
	returnResult, err := p.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	retData := string("[" + string(returnResult) + "]")
	yamlBytes := processParser([]byte(retData), skip)
	fmt.Println(string(yamlBytes))
	return returnResult, nil
}

// ParseAllProjects export all projects from lagoon and parse them to the import format
func (p *Parser) ParseAllProjects(skip SkipExport) ([]byte, error) {
	customReq := api.CustomRequest{
		Query: `fragment NotificationSlack on NotificationSlack {
			webhook
		  name
			channel
		}
		fragment NotificationRocket on NotificationRocketChat {
			webhook
		  name
			channel
		}
		# @TODO: enable once 1.2.0+ lagoon is more widespread
		# fragment NotificationEmail on NotificationEmail {
		# 	emailAddress
		#   name
		# }
		# fragment NotificationTeams on NotificationMicrosoftTeams {
		# 	webhook
		#   name
		# }
		query Projects {
			allProjects {
			name
			autoIdle
			branches
			pullrequests
			privateKey
			productionEnvironment
			activeSystemsDeploy
			activeSystemsTask
			activeSystemsRemove
			activeSystemsPromote
			storageCalc
			openshiftProjectPattern
			developmentEnvironmentsLimit
			gitUrl
			autoIdle
			groups{
			  name
			  id
			  members{
				user{
				  email
				  sshKeys{
					name
					keyType
					keyValue
				  }
				  firstName
				  lastName
				}
				role
			  }
			}
			notifications {
			  __typename
			  ...NotificationRocket
			  ...NotificationSlack
			  # @TODO: enable once 1.2.0+ lagoon is more widespread
			  # ...NotificationEmail
			  # ...NotificationTeams
			}
			openshift{
			  id
			}
			envVariables {
			  name
			  scope
			  value
			}
			environments {
			  id
			  name
			  openshiftProjectName
			  autoIdle
			  envVariables {
				name
				scope
				value
			  }
			}
		  }
		}`,
		Variables:    map[string]interface{}{},
		MappedResult: "allProjects",
	}
	returnResult, err := p.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	// fmt.Println(string(returnResult))
	// _ = processParser(returnResult)
	yamlBytes := processParser([]byte(returnResult), skip)
	fmt.Println(string(yamlBytes))
	return returnResult, nil
}
