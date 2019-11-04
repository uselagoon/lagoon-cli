package api

import (
	"encoding/json"
	"github.com/machinebox/graphql"
)

// AddUser .
func (api *Interface) AddUser(user User) ([]byte, error) {
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
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["addUser"])
	if err != nil {
		return []byte(""), err
	}
	return jsonBytes, nil
}

// UpdateUser .
func (api *Interface) UpdateUser(user UpdateUser) ([]byte, error) {
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
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["updateUser"])
	if err != nil {
		return []byte(""), err
	}
	return jsonBytes, nil
}

// DeleteUser .
func (api *Interface) DeleteUser(user User) ([]byte, error) {
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
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["deleteUser"])
	if err != nil {
		return []byte(""), err
	}
	return jsonBytes, nil
}

// GetUserBySSHKey .
func (api *Interface) GetUserBySSHKey(sshKey SSHKeyValue) ([]byte, error) {
	req := graphql.NewRequest(`
	query userBySshKey($sshKey: String!) {
		userBySshKey(sshKey: $sshKey) {
			...User
		}
	}` + userFragment)
	req.Var("sshKey", sshKey)
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["userBySshKey"])
	if err != nil {
		return []byte(""), err
	}
	return jsonBytes, nil
}

// AddSSHKey .
func (api *Interface) AddSSHKey(sshKey AddSSHKey) ([]byte, error) {
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
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["addSshKey"])
	if err != nil {
		return []byte(""), err
	}
	return jsonBytes, nil
}

// DeleteSSHKey .
func (api *Interface) DeleteSSHKey(sshKey DeleteSSHKey) ([]byte, error) {
	req := graphql.NewRequest(`
	mutation ($name: String!) {
		deleteSshKey(input: {
		 	name: $name
		})
	}`)
	req.Var("name", sshKey.Name)
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["deleteSshKey"])
	if err != nil {
		return []byte(""), err
	}
	return jsonBytes, nil
}
