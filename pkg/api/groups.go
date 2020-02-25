package api

import (
	"encoding/json"
	"errors"

	"github.com/machinebox/graphql"
)

// AddGroup .
func (api *Interface) AddGroup(group AddGroup) ([]byte, error) {
	req := graphql.NewRequest(`
	mutation ($name: String!) {
		addGroup(input: {
			name: $name
		}) {
			...Group
		}
	}` + groupFragment)
	generateVars(req, group)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["addGroup"])
	if err != nil {
		return []byte(""), err
	}
	if api.debug {
		debugResponse(jsonBytes)
	}
	if string(jsonBytes) == "null" {
		return []byte(""), errors.New("graphql: returned null")
	}
	return jsonBytes, nil
}

// AddGroupWithParent .
func (api *Interface) AddGroupWithParent(group AddGroup) ([]byte, error) {
	req := graphql.NewRequest(`
	mutation ($name: String!, $parentGroupName: String) {
		addGroup(input: {
			name: $name
			parentGroup: { name: $parentGroupName }
		}) {
		 	...Group
		}
	}` + groupFragment)
	req.Var("name", group.Name)
	req.Var("parentGroupName", group.ParentGroup.Name)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["addGroup"])
	if err != nil {
		return []byte(""), err
	}
	if api.debug {
		debugResponse(jsonBytes)
	}
	if string(jsonBytes) == "null" {
		return []byte(""), errors.New("graphql: returned null")
	}
	return jsonBytes, nil
}

// UpdateGroup .
func (api *Interface) UpdateGroup(group UpdateGroup) ([]byte, error) {
	req := graphql.NewRequest(`
	mutation ($name: String!, $patch: UpdateGroupPatchInput!) {
		updateGroup(input: {
			group: {
				name: $name
			}
			patch: $patch
		}) {
		  	...Group
		}
	}` + groupFragment)
	req.Var("name", group.Group.Name)
	req.Var("patch", group.Patch)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["updateGroup"])
	if err != nil {
		return []byte(""), err
	}
	if api.debug {
		debugResponse(jsonBytes)
	}
	if string(jsonBytes) == "null" {
		return []byte(""), errors.New("graphql: returned null")
	}
	return jsonBytes, nil
}

// DeleteGroup .
func (api *Interface) DeleteGroup(group AddGroup) ([]byte, error) {
	req := graphql.NewRequest(`
	mutation ($name: String!) {
		deleteGroup(input: {
			group: {
				name: $name
			}
		})
	}`)
	generateVars(req, group)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["deleteGroup"])
	if err != nil {
		return []byte(""), err
	}
	if api.debug {
		debugResponse(jsonBytes)
	}
	if string(jsonBytes) == "null" {
		return []byte(""), errors.New("graphql: returned null")
	}
	return jsonBytes, nil
}

// AddUserToGroup .
func (api *Interface) AddUserToGroup(user AddUserToGroup) ([]byte, error) {
	req := graphql.NewRequest(`
	mutation ($userEmail: String!, $groupName: String!, $role: GroupRole!) {
		addUserToGroup(input: {
			user: { email: $userEmail }
			group: { name: $groupName }
			role: $role
		}) {
		 	...Group
		}
	}` + groupFragment)
	req.Var("userEmail", user.User.Email)
	req.Var("groupName", user.Group)
	req.Var("role", user.Role)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["addUserToGroup"])
	if err != nil {
		return []byte(""), err
	}
	if api.debug {
		debugResponse(jsonBytes)
	}
	if string(jsonBytes) == "null" {
		return []byte(""), errors.New("graphql: returned null")
	}
	return jsonBytes, nil
}

// AddGroupToProject .
func (api *Interface) AddGroupToProject(group ProjectToGroup) ([]byte, error) {
	req := graphql.NewRequest(`
	mutation ($project: String!, $group: String!) {
		addUserToGroup(input: {
			project: { name: $project}
			groups: [{name: $group}]
		}) {
		 	...Project
		}
	}` + projectFragment)
	generateVars(req, group)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["addUserToGroup"])
	if err != nil {
		return []byte(""), err
	}
	if api.debug {
		debugResponse(jsonBytes)
	}
	if string(jsonBytes) == "null" {
		return []byte(""), errors.New("graphql: returned null")
	}
	return jsonBytes, nil
}

// RemoveGroupFromProject .
func (api *Interface) RemoveGroupFromProject(group ProjectToGroup) ([]byte, error) {
	req := graphql.NewRequest(`
	mutation ($project: String!, $group: String!) {
		removeGroupsFromProject(input: {
			project: { name: $project}
			groups: [{name: $group}]
		}) {
		 	...Project
		}
	}` + projectFragment)
	generateVars(req, group)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["removeGroupsFromProject"])
	if err != nil {
		return []byte(""), err
	}
	if api.debug {
		debugResponse(jsonBytes)
	}
	if string(jsonBytes) == "null" {
		return []byte(""), errors.New("graphql: returned null")
	}
	return jsonBytes, nil
}

// RemoveUserFromGroup .
func (api *Interface) RemoveUserFromGroup(user UserGroup) ([]byte, error) {
	req := graphql.NewRequest(`
	mutation ($userEmail: String!, $groupName: String!) {
		removeUserFromGroup(input: {
			user: { email: $userEmail }
			group: { name: $groupName }
		}) {
		 	...Group
		}
	}` + groupFragment)
	req.Var("userEmail", user.User.Email)
	req.Var("groupName", user.Group)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["removeUserFromGroup"])
	if err != nil {
		return []byte(""), err
	}
	if api.debug {
		debugResponse(jsonBytes)
	}
	if string(jsonBytes) == "null" {
		return []byte(""), errors.New("graphql: returned null")
	}
	return jsonBytes, nil
}
