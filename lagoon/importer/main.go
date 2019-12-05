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
	Users      []lagoonUsers                `json:"users"`
	Projects   []lagoonProjects             `json:"projects"`
}

type lagoonProjects struct {
	Project       api.ProjectPatch    `json:"project"`
	Groups        []string            `json:"groups"`
	Notifications lagoonNotifications `json:"notifications"`
}

type lagoonUsers struct {
	User   lagoonUser       `json:"user"`
	Groups []addUserToGroup `json:"groups"`
}

type lagoonNotifications struct {
	Slack      []string `json:"slack"`
	RocketChat []string `json:"rocketchat"`
}

type lagoonUser struct {
	Email      string `json:"email"`
	SSHKey     string `json:"sshkey"`
	KeyName    string `json:"keyname,omitempty"`
	SSHKeyFile string `json:"sshkeyfile,omitempty"`
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
	// start with adding groups
	for _, group := range importedData.Groups {
		addGroup(group, forceAction)
	}
	// next add users and any keys, then add them to any groups they need to be in
	for _, user := range importedData.Users {
		addUser(user.User, forceAction)
		addKeyToUser(user.User, forceAction)
		for _, group := range user.Groups {
			addUserGroup(user.User, group, forceAction)
		}
	}
	// create any notification providers
	for _, slack := range importedData.Slack {
		addSlack(slack, forceAction)
	}
	for _, rocketchat := range importedData.RocketChat {
		addRocketChat(rocketchat, forceAction)
	}
	// now add the projects
	for _, project := range importedData.Projects {
		addProject(project.Project, forceAction)
		// add them to any groups they need to be in
		for _, group := range project.Groups {
			addGroupProject(project.Project.Name, group, forceAction)
		}
		// then add any notification services to the project if required
		addSlacks(project.Notifications.Slack, project.Project.Name, forceAction)
		addRocketChats(project.Notifications.RocketChat, project.Project.Name, forceAction)
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

func addRocketChats(rocketchats []string, name string, action bool) {
	for _, rocketchat := range rocketchats {
		fmt.Println("Adding rocketchat", rocketchat, "to project", name)
		_, err := projects.AddRocketChatNotificationToProject(name, rocketchat)
		if err != nil {
			fmt.Println(err)
			if !yesNo("Continue?", action) {
				os.Exit(1)
			}
		}
	}
}

func addSlacks(slacks []string, name string, action bool) {
	for _, slack := range slacks {
		fmt.Println("Adding slack", slack, "to project", name)
		_, err := projects.AddSlackNotificationToProject(name, slack)
		if err != nil {
			fmt.Println(err)
			if !yesNo("Continue?", action) {
				os.Exit(1)
			}
		}
	}
}

func addGroupProject(name string, group string, action bool) {
	fmt.Println("Adding project", name, "to group", group)
	projectGroup := api.ProjectGroups{
		Project: api.Project{
			Name: name,
		},
		Groups: []api.Group{
			api.Group{
				Name: group,
			},
		},
	}
	_, err := users.AddProjectToGroup(projectGroup)
	if err != nil {
		fmt.Println(err)
		if !yesNo("Continue?", action) {
			os.Exit(1)
		}
	}
}

func addProject(project api.ProjectPatch, action bool) {
	jsonPatch, _ := json.Marshal(project)
	fmt.Println("Adding project", project.Name)
	addResult, err := projects.AddProject(project.Name, string(jsonPatch))
	if err != nil {
		fmt.Println(err)
		if !yesNo("Continue?", action) {
			os.Exit(1)
		}
		//os.Exit(1)
	} else {
		var addedProject api.Project
		err = json.Unmarshal([]byte(addResult), &addedProject)
		if err != nil {
			fmt.Println(err)
			if !yesNo("Continue?", action) {
				os.Exit(1)
			}
			//os.Exit(1)
		}
		// fmt.Println(addedProject)
	}
}

func addRocketChat(rocketchat api.NotificationRocketChat, action bool) {
	fmt.Println("Adding rocketchat", rocketchat.Name)
	_, err := projects.AddRocketChatNotification(rocketchat.Name, rocketchat.Channel, rocketchat.Webhook)
	if err != nil {
		fmt.Println(err)
		if !yesNo("Continue?", action) {
			os.Exit(1)
		}
	}
}

func addSlack(slack api.NotificationSlack, action bool) {
	fmt.Println("Adding slack", slack.Name)
	_, err := projects.AddRocketChatNotification(slack.Name, slack.Channel, slack.Webhook)
	if err != nil {
		fmt.Println(err)
		if !yesNo("Continue?", action) {
			os.Exit(1)
		}
	}
}

func addGroup(group api.Group, action bool) {
	fmt.Println("Adding group", group.Name)
	_, err := users.AddGroup(group)
	if err != nil {
		fmt.Println(err)
		if !yesNo("Continue?", action) {
			os.Exit(1)
		}
	}
}

func addUser(user lagoonUser, action bool) {
	fmt.Println("Adding user", user.Email)
	userData := api.User{
		Email: user.Email,
	}
	_, err := users.AddUser(userData)
	if err != nil {
		fmt.Println(err)
		if !yesNo("Continue?", action) {
			os.Exit(1)
		}
	}
}

func addKeyToUser(user lagoonUser, action bool) {
	userData := api.User{
		Email: user.Email,
	}
	splitKey := strings.Split(user.SSHKey, " ")
	var keyType api.SSHKeyType
	// default to ssh-rsa, otherwise check if ssh-ed25519 as we only support these in lagoon
	// will fail if neither are right
	keyType = api.SSHRsa
	if strings.EqualFold(string(splitKey[0]), "ssh-ed25519") {
		keyType = api.SSHEd25519
	}
	// if the sshkey has a comment/name in it, we can use that, otherwise define one using `keyname`
	keyName := user.KeyName
	if keyName == "" && len(splitKey) == 3 {
		//strip new line
		keyName = strings.TrimSuffix(splitKey[2], "\n")
	} else if keyName == "" && len(splitKey) == 2 {
		fmt.Println("No keyname defined")
		if !yesNo("Continue?", action) {
			os.Exit(1)
		}
	}
	sshKey := api.SSHKey{
		KeyType:  keyType,
		KeyValue: splitKey[1],
		Name:     keyName,
	}
	fmt.Println("Adding key to user", user.Email)
	_, err := users.AddSSHKeyToUser(userData, sshKey)
	if err != nil {
		fmt.Println(err)
		if !yesNo("Continue?", action) {
			os.Exit(1)
		}
	}
}

func addUserGroup(user lagoonUser, group addUserToGroup, action bool) {
	userData := api.User{
		Email: user.Email,
	}
	var roleType api.GroupRole
	switch strings.ToLower(string(group.Role)) {
	case "guest":
		roleType = api.GuestRole
	case "reporter":
		roleType = api.ReporterRole
	case "developer":
		roleType = api.DeveloperRole
	case "maintainer":
		roleType = api.MaintainerRole
	case "owner":
		roleType = api.OwnerRole
	default:
		// default to guest if unable to determine from provided role
		roleType = api.GuestRole
	}
	userGroupRole := api.UserGroupRole{
		User: userData,
		Group: api.Group{
			Name: group.Name,
		},
		Role: roleType,
	}
	var err error
	fmt.Println("Adding user", user.Email, "to group", group.Name, "with role", roleType)
	_, err = users.AddUserToGroup(userGroupRole)
	if err != nil {
		fmt.Println(err)
		if !yesNo("Continue?", action) {
			os.Exit(1)
		}
	}
}
