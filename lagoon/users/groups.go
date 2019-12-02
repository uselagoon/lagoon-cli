package users

import (
	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/graphql"
)

// AddGroup function
func AddGroup(group api.Group) ([]byte, error) {
	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}
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
	returnResult, err := lagoonAPI.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// AddGroupWithParent function
func AddGroupWithParent(group api.Group, parent api.Group) ([]byte, error) {
	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}
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
	returnResult, err := lagoonAPI.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// AddUserToGroup function
func AddUserToGroup(userGroup api.UserGroupRole) ([]byte, error) {
	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}
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
	returnResult, err := lagoonAPI.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// AddProjectToGroup function
func AddProjectToGroup(groups api.ProjectGroups) ([]byte, error) {
	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}
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
	returnResult, err := lagoonAPI.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// DeleteGroup function
func DeleteGroup(group api.Group) ([]byte, error) {
	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}
	customReq := api.CustomRequest{
		Query: `mutation deleteGroup ($name: String!) {
				deleteGroup(input:{group:{name: $name}})
			}`,
		Variables: map[string]interface{}{
			"name": group.Name,
		},
		MappedResult: "deleteGroup",
	}
	returnResult, err := lagoonAPI.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// RemoveUserFromGroup function
func RemoveUserFromGroup(userGroup api.UserGroup) ([]byte, error) {
	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}
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
	returnResult, err := lagoonAPI.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// RemoveGroupsFromProject function
func RemoveGroupsFromProject(groups api.ProjectGroups) ([]byte, error) {
	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}
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
	returnResult, err := lagoonAPI.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}
