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
	"gopkg.in/yaml.v2"
)

type lagoonImport struct {
	Groups []api.Group `json:"groups"`
	Users  []struct {
		User   api.User         `json:"user"`
		Groups []addUserToGroup `json:"groups"`
	} `json:"users"`
	Projects []api.ProjectPatch `json:"projects"`
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
func ImportData(importFile string) {
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
		fmt.Println(group.Name)
		customReqResult, err := users.AddGroup(group)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(customReqResult))
	}
	for _, user := range importedData.Users {
		fmt.Println(user.User.Email)
		customReqResult, err := users.AddUser(user.User)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(customReqResult))
		for _, group := range user.Groups {
			fmt.Println(group.Role)
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
				User: user.User,
				Group: api.Group{
					Name: group.Name,
				},
				Role: roleType,
			}
			var customReqResult []byte
			var err error
			customReqResult, err = users.AddUserToGroup(userGroupRole)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			returnResultData := map[string]interface{}{}
			err = json.Unmarshal([]byte(customReqResult), &returnResultData)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(returnResultData)
		}
	}
	for _, v := range importedData.Projects {
		fmt.Println(v.Name, v.GitURL, v.ProductionEnvironment, v.Branches)
		jsonPatch, _ := json.Marshal(v)
		addResult, err := projects.AddProject(v.Name, string(jsonPatch))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		var addedProject api.Project
		err = json.Unmarshal([]byte(addResult), &addedProject)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(addedProject)
	}
}
