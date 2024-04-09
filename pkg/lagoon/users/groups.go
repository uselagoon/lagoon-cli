package users

import (
	"encoding/json"
	"strconv"

	"github.com/uselagoon/lagoon-cli/pkg/api"
	"github.com/uselagoon/lagoon-cli/pkg/output"
)

// AddGroupWithParent function
func (u *Users) AddGroupWithParent(group api.Group, parent api.Group) ([]byte, error) {
	customReq := api.CustomRequest{
		Query: `mutation addGroup ($name: String!, $parent: GroupInput) {
			addGroup(input:{name: $name}) {
					id
				}
			}`,
		Variables: map[string]interface{}{
			"name":  group.Name,
			"group": parent,
		},
		MappedResult: "addGroup",
	}
	returnResult, err := u.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

func processListGroupProjects(groupData []byte, allProjects bool) ([]byte, error) {
	var data []output.Data
	var groups []struct {
		Name     string
		ID       string
		Projects []struct {
			ID   int
			Name string
		}
	}
	err := json.Unmarshal(groupData, &groups)
	if err != nil {
		return []byte(""), err
	}
	for _, group := range groups {
		for _, project := range group.Projects {
			projectData := []string{strconv.Itoa(project.ID), project.Name}
			if allProjects {
				projectData = append(projectData, group.Name)
			}
			data = append(data, projectData)
		}
	}
	dataMain := output.Table{
		Header: []string{"ID", "ProjectName"},
		Data:   data,
	}
	if allProjects {
		dataMain.Header = append(dataMain.Header, "GroupName")
	}
	return json.Marshal(dataMain)
}
