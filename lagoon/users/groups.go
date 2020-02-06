package users

import (
	"github.com/amazeeio/lagoon-cli/api"
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
	returnResult, err := u.api.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}
