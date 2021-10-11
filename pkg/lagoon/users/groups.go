package users

import (
	"encoding/json"
	"strconv"

	"github.com/uselagoon/lagoon-cli/pkg/api"
	"github.com/uselagoon/lagoon-cli/pkg/output"
)

// AddGroup function
func (u *Users) AddGroup(group api.Group) ([]byte, error) {
	customReq := api.CustomRequest{
		Query: `mutation addGroup ($name: String!) {
			addGroup(input:{name: $name}) {
					id
				}
			}`,
		Variables: map[string]interface{}{
			"name": group.Name,
		},
		MappedResult: "addGroup",
	}
	returnResult, err := u.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

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

// AddUserToGroup function
func (u *Users) AddUserToGroup(userGroup api.UserGroupRole) ([]byte, error) {
	customReq := api.CustomRequest{
		Query: `mutation addUserToGroup($email: String!, $group: String!, $role: GroupRole!) {
				addUserToGroup(input:{
					user: {
						email: $email
					}
					group: {
						name: $group
					}
					role: $role
				}) 
				{
					id
				}
			}`,
		Variables: map[string]interface{}{
			"email": userGroup.User.Email,
			"group": userGroup.Group.Name,
			"role":  userGroup.Role,
		},
		MappedResult: "addUser",
	}
	returnResult, err := u.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// AddProjectToGroup function
func (u *Users) AddProjectToGroup(groups api.ProjectGroups) ([]byte, error) {
	customReq := api.CustomRequest{
		Query: `mutation addGroupsToProject($project: String!, $groups: [GroupInput!]!) {
			addGroupsToProject(input:{
				groups: $groups
				project: {name: $project}
			}) {
					id
				}
			}`,
		Variables: map[string]interface{}{
			"groups":  groups.Groups,
			"project": groups.Project.Name,
		},
		MappedResult: "addGroupsToProject",
	}
	returnResult, err := u.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// DeleteGroup function
func (u *Users) DeleteGroup(group api.Group) ([]byte, error) {
	customReq := api.CustomRequest{
		Query: `mutation deleteGroup ($name: String!) {
				deleteGroup(input:{group:{name: $name}})
			}`,
		Variables: map[string]interface{}{
			"name": group.Name,
		},
		MappedResult: "deleteGroup",
	}
	returnResult, err := u.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// RemoveUserFromGroup function
func (u *Users) RemoveUserFromGroup(userGroup api.UserGroup) ([]byte, error) {
	customReq := api.CustomRequest{
		Query: `mutation removeUserFromGroup ($email: String!, $group: String!) {
				removeUserFromGroup(input:{group:{name: $group} user:{email: $email}}) {
					id
				}
			}`,
		Variables: map[string]interface{}{
			"email": userGroup.User.Email,
			"group": userGroup.Group.Name,
		},
		MappedResult: "removeUserFromGroup",
	}
	returnResult, err := u.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// RemoveGroupsFromProject function
func (u *Users) RemoveGroupsFromProject(groups api.ProjectGroups) ([]byte, error) {
	customReq := api.CustomRequest{
		Query: `mutation removeGroupsFromProject ($project: String!, $groups: [GroupInput!]!) {
				removeGroupsFromProject(input:{groups: $groups project:{name: $project}}) {
					id
				}
			}`,
		Variables: map[string]interface{}{
			"groups":  groups.Groups,
			"project": groups.Project.Name,
		},
		MappedResult: "removeGroupsFromProject",
	}
	returnResult, err := u.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// ListGroups function
func (u *Users) ListGroups(name string) ([]byte, error) {
	customReq := api.CustomRequest{
		Query: `query allGroups ($name: String) {
			allGroups(name: $name) {
				id
				name
			}
		}`,
		Variables: map[string]interface{}{
			"name": name,
		},
		MappedResult: "allGroups",
	}
	reqResult, err := u.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	returnResult, err := processListGroups(reqResult)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

func processListGroups(groupData []byte) ([]byte, error) {
	var data []output.Data
	var groups []struct {
		ID   string
		Name string
	}
	err := json.Unmarshal(groupData, &groups)
	if err != nil {
		return []byte(""), err
	}
	for _, group := range groups {
		data = append(data, []string{group.ID, group.Name})
	}
	dataMain := output.Table{
		Header: []string{"ID", "Name"},
		Data:   data,
	}
	return json.Marshal(dataMain)
}

// ListGroupProjects function
func (u *Users) ListGroupProjects(name string, allProjects bool) ([]byte, error) {
	customReq := api.CustomRequest{
		Query: `query allGroups ($name: String) {
			allGroups(name: $name) {
				id
				name
				projects {
					id
					name
				}
			}
		}`,
		Variables: map[string]interface{}{
			"name": name,
		},
		MappedResult: "allGroups",
	}
	reqResult, err := u.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	returnResult, err := processListGroupProjects(reqResult, allProjects)
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
