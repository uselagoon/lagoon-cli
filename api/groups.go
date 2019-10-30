package api

import (
	"github.com/machinebox/graphql"
)

// AddGroup .
func (api *Interface) AddGroup(group AddGroup) (interface{}, error) {
	req := graphql.NewRequest(`
	mutation ($name: String!) {
		addGroup(input: {
			name: $name
		}) {
			...Group
		}
	}` + groupFragment)
	generateVars(req, group)
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// AddGroupWithParent .
func (api *Interface) AddGroupWithParent(group AddGroup) (interface{}, error) {
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
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// UpdateGroup .
func (api *Interface) UpdateGroup(group UpdateGroup) (interface{}, error) {
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
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// DeleteGroup .
func (api *Interface) DeleteGroup(group AddGroup) (interface{}, error) {
	req := graphql.NewRequest(`
	mutation ($name: String!) {
		deleteGroup(input: {
			group: {
				name: $name
			}
		})
	}`)
	generateVars(req, group)
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// AddUserToGroup .
func (api *Interface) AddUserToGroup(user AddUserToGroup) (interface{}, error) {
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
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// AddGroupToProject .
func (api *Interface) AddGroupToProject(group ProjectToGroup) (interface{}, error) {
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
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// RemoveGroupFromProject .
func (api *Interface) RemoveGroupFromProject(group ProjectToGroup) (interface{}, error) {
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
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// RemoveUserFromGroup .
func (api *Interface) RemoveUserFromGroup(user UserGroup) (interface{}, error) {
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
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}
