package api

import (
	"github.com/machinebox/graphql"
)

// AddUser .
func (api *Interface) AddUser(user User) (interface{}, error) {
	req := graphql.NewRequest(`
	mutation ($email: String!, $firstName: String, $lastName: String, $comment: String, $gitlabId: Int) {
		addUser(input: {
			email: $email
			firstName: $firstName
			lastName: $lastName
			comment: $comment
			gitlabId: $gitlabId
		}) {
			...User
		}
	}` + userFragment)
	req.Var("email", user.Email)
	req.Var("firstName", user.FirstName)
	req.Var("lastName", user.LastName)
	req.Var("comment", user.Comment)
	req.Var("gitlabId", user.GitlabID)
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// UpdateUser .
func (api *Interface) UpdateUser(user UpdateUser) (interface{}, error) {
	req := graphql.NewRequest(`
	mutation ($email: String!, $patch: UpdateUserPatchInput!) {
		updateUser(input: {
			user: {
				email: $email
			}
			patch: $patch
		}) {
			...User
		}
	}` + userFragment)
	req.Var("email", user.User.Email)
	req.Var("patch", user.Patch)
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// DeleteUser .
func (api *Interface) DeleteUser(user User) (interface{}, error) {
	req := graphql.NewRequest(`
	mutation ($email: String!) {
		deleteUser(input: {
			user: {
				email: $email
			}
		})
	}`)
	req.Var("email", user.Email)
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// GetUserBySSHKey .
func (api *Interface) GetUserBySSHKey(sshKey SSHKeyValue) (interface{}, error) {
	req := graphql.NewRequest(`
	query userBySshKey($sshKey: String!) {
		userBySshKey(sshKey: $sshKey) {
			...User
		}
	}` + userFragment)
	req.Var("sshKey", sshKey)
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// AddSSHKey .
func (api *Interface) AddSSHKey(sshKey AddSSHKey) (interface{}, error) {
	req := graphql.NewRequest(`
	mutation ($id: Int, $name: String!, $keyValue: String!, $keyType: SshKeyType!, $userEmail: String!) {
		addSshKey(input: {
			id: $id
			name: $name
			keyValue: $keyValue
			keyType: $keyType
			user: {
				email: $userEmail
			}
		}) {
		 	...SshKey
		}
	}` + sshKeyFragment)
	req.Var("name", sshKey.Name)
	req.Var("keyValue", sshKey.KeyValue)
	req.Var("keyType", sshKey.KeyType)
	req.Var("userEmail", sshKey.User.Email)
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// DeleteSSHKey .
func (api *Interface) DeleteSSHKey(sshKey DeleteSSHKey) (interface{}, error) {
	req := graphql.NewRequest(`
	mutation ($name: String!) {
		deleteSshKey(input: {
		 	name: $name
		})
	}`)
	req.Var("name", sshKey.Name)
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}
