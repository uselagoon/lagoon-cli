package importer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/lagoon/projects"
	"github.com/amazeeio/lagoon-cli/lagoon/users"
	"github.com/manifoldco/promptui"
	"gopkg.in/yaml.v2"
)

type lagoonImport struct {
	Groups     []api.Group                  `json:"groups"`
	Slack      []api.NotificationSlack      `json:"slack"`
	RocketChat []api.NotificationRocketChat `json:"rocketchat"`
	Users      []struct {
		User struct {
			Email      string `json:"email"`
			SSHKey     string `json:"sshkey"`
			KeyName    string `json:"keyname,omitempty"`
			SSHKeyFile string `json:"sshkeyfile,omitempty"`
		} `json:"user"`
		Groups []addUserToGroup `json:"groups"`
	} `json:"users"`
	Projects []struct {
		Project api.ProjectPatch `json:"project"`
		// Project struct {
		// 	Name                  string `yaml:"name"`
		// 	GitURL                string `yaml:"gitUrl"`
		// 	Openshift             int    `yaml:"openshift"`
		// 	Branches              string `yaml:"branches"`
		// 	ProductionEnvironment string `yaml:"productionEnvironment"`
		// } `json:"project"`
		Groups        []string `json:"groups"`
		Notifications struct {
			Slack      []string `json:"slack"`
			RocketChat []string `json:"rocketchat"`
		} `json:"notifications"`
	} `json:"projects"`
}

type addUserToGroup struct {
	User api.User `json:"user"`
	Name string   `json:"name"`
	Role string   `json:"role"`
}

// example
// var yamlData = `
// groups:
//   - name: example-com
// users:
//   - user:
//       email: usera@example.com
//       sshkey: ~/usera.pub
//     groups:
//       - name: example-com
//         role: owner
//   - user:
//       email: userb@example.com
//       sshkey: ~/userb.pub
//     groups:
//       - name: example-com
//         role: developere
// projects:
//   - name: example-com
//     giturl: git@github.com:example/example-com.git
//     openshift: 1
//     branches: "master|develop|staging"
//     productionenvironment: master
//     groups:
//       - example-com
// `

// ImportData func
func ImportData(importFile string, forceAction bool) {
	yamlData, err := ioutil.ReadFile(importFile) // just pass the file name
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	importedData := lagoonImport{}

	err = yaml.Unmarshal([]byte(yamlData), &importedData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// @TODO: do a platform-owner only check and fail sooner, tell user import is only for platform-owners currently
	for _, group := range importedData.Groups {
		fmt.Println("Adding group", group.Name)
		_, err := users.AddGroup(group)
		if err != nil {
			fmt.Println(err)
			if !yesNo("Continue?", forceAction) {
				os.Exit(1)
			}
		}
		// fmt.Println(string(customReqResult))
	}
	for _, user := range importedData.Users {
		fmt.Println("Adding user", user.User.Email)
		addUser := api.User{
			Email: user.User.Email,
		}
		_, err := users.AddUser(addUser)
		if err != nil {
			fmt.Println(err)
			if !yesNo("Continue?", forceAction) {
				os.Exit(1)
			}
		}
		splitKey := strings.Split(user.User.SSHKey, " ")
		var keyType api.SSHKeyType
		// default to ssh-rsa, otherwise check if ssh-ed25519
		// will fail if neither are right
		keyType = api.SSHRsa
		if strings.EqualFold(string(splitKey[0]), "ssh-ed25519") {
			keyType = api.SSHEd25519
		}
		// if the sshkey has a comment/name in it, we can use that, otherwise define one using `keyname`
		keyName := user.User.KeyName
		if keyName == "" && len(splitKey) == 3 {
			//strip new line
			keyName = strings.TrimSuffix(splitKey[2], "\n")
		} else if keyName == "" && len(splitKey) == 2 {
			fmt.Println(err)
			if !yesNo("Continue?", forceAction) {
				os.Exit(1)
			}
		}
		sshKey := api.SSHKey{
			KeyType:  keyType,
			KeyValue: splitKey[1],
			Name:     keyName,
		}
		_, err = users.AddSSHKeyToUser(addUser, sshKey)
		if err != nil {
			fmt.Println(err)
			if !yesNo("Continue?", forceAction) {
				os.Exit(1)
			}
		}
		for _, group := range user.Groups {
			var roleType api.GroupRole
			roleType = api.GuestRole
			if strings.EqualFold(string(group.Role), "guest") {
				roleType = api.GuestRole
			} else if strings.EqualFold(string(group.Role), "reporter") {
				roleType = api.ReporterRole
			} else if strings.EqualFold(string(group.Role), "developer") {
				roleType = api.DeveloperRole
			} else if strings.EqualFold(string(group.Role), "maintainer") {
				roleType = api.MaintainerRole
			} else if strings.EqualFold(string(group.Role), "owner") {
				roleType = api.OwnerRole
			}
			userGroupRole := api.UserGroupRole{
				User: addUser,
				Group: api.Group{
					Name: group.Name,
				},
				Role: roleType,
			}
			var err error
			fmt.Println("Adding user", user.User.Email, "to group", group.Name)
			_, err = users.AddUserToGroup(userGroupRole)
			if err != nil {
				fmt.Println(err)
				if !yesNo("Continue?", forceAction) {
					os.Exit(1)
				}
			}
		}
	}
	for _, slack := range importedData.Slack {
		fmt.Println("Adding slack", slack.Name)
		_, err := projects.AddSlackNotification(slack.Name, slack.Channel, slack.Webhook)
		if err != nil {
			fmt.Println(err)
			if !yesNo("Continue?", forceAction) {
				os.Exit(1)
			}
		}
		// fmt.Println(string(customReqResult))
	}
	for _, rocketchat := range importedData.RocketChat {
		fmt.Println("Adding rocketchat", rocketchat.Name)
		_, err := projects.AddRocketChatNotification(rocketchat.Name, rocketchat.Channel, rocketchat.Webhook)
		if err != nil {
			fmt.Println(err)
			if !yesNo("Continue?", forceAction) {
				os.Exit(1)
			}
		}
		// fmt.Println(string(customReqResult))
	}
	for _, project := range importedData.Projects {
		jsonPatch, _ := json.Marshal(project.Project)
		fmt.Println("Adding project", project.Project.Name)
		addResult, err := projects.AddProject(project.Project.Name, string(jsonPatch))
		if err != nil {
			fmt.Println(err)
			if !yesNo("Continue?", forceAction) {
				os.Exit(1)
			}
			//os.Exit(1)
		} else {
			var addedProject api.Project
			err = json.Unmarshal([]byte(addResult), &addedProject)
			if err != nil {
				fmt.Println(err)
				if !yesNo("Continue?", forceAction) {
					os.Exit(1)
				}
				//os.Exit(1)
			}
			// fmt.Println(addedProject)
		}
		for _, group := range project.Groups {
			fmt.Println("Adding project", project.Project.Name, "to group", group)
			projectGroup := api.ProjectGroups{
				Project: api.Project{
					Name: project.Project.Name,
				},
				Groups: []api.Group{
					api.Group{
						Name: group,
					},
				},
			}
			_, err = users.AddProjectToGroup(projectGroup)
			if err != nil {
				fmt.Println(err)
				if !yesNo("Continue?", forceAction) {
					os.Exit(1)
				}
			}
		}
		for _, slack := range project.Notifications.Slack {
			fmt.Println("Adding slack", slack, "to project", project.Project.Name)
			_, err = projects.AddSlackNotificationToProject(project.Project.Name, slack)
			if err != nil {
				fmt.Println(err)
				if !yesNo("Continue?", forceAction) {
					os.Exit(1)
				}
			}
		}
		for _, rocketchat := range project.Notifications.RocketChat {
			fmt.Println("Adding rocketchat", rocketchat, "to project", project.Project.Name)
			_, err = projects.AddRocketChatNotificationToProject(project.Project.Name, rocketchat)
			if err != nil {
				fmt.Println(err)
				if !yesNo("Continue?", forceAction) {
					os.Exit(1)
				}
			}
		}
	}
}

func yesNo(message string, forceAction bool) bool {
	if forceAction != true {
		prompt := promptui.Select{
			Label: message + "; Select[Yes/No]",
			Items: []string{"No", "Yes"},
		}
		_, result, err := prompt.Run()
		if err != nil {
			os.Exit(1)
		}
		return result == "Yes"
	}
	return true
}
