package api

import (
	"encoding/json"
	"errors"

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
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["addUser"])
	if err != nil {
		return []byte(""), err
	}
	if api.debug {
		debugResponse(jsonBytes)
	}
	if string(jsonBytes) == "null" {
		return []byte(""), errors.New("GraphQL API returned a null response, the requested resource may not exist, or there was an error. Use `--debug` to check what was returned")
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
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["userBySshKey"])
	if err != nil {
		return []byte(""), err
	}
	if api.debug {
		debugResponse(jsonBytes)
	}
	if string(jsonBytes) == "null" {
		return []byte(""), errors.New("GraphQL API returned a null response, the requested resource may not exist, or there was an error. Use `--debug` to check what was returned")
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
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["addSshKey"])
	if err != nil {
		return []byte(""), err
	}
	if api.debug {
		debugResponse(jsonBytes)
	}
	if string(jsonBytes) == "null" {
		return []byte(""), errors.New("GraphQL API returned a null response, the requested resource may not exist, or there was an error. Use `--debug` to check what was returned")
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
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["deleteSshKey"])
	if err != nil {
		return []byte(""), err
	}
	if api.debug {
		debugResponse(jsonBytes)
	}
	if string(jsonBytes) == "null" {
		return []byte(""), errors.New("GraphQL API returned a null response, the requested resource may not exist, or there was an error. Use `--debug` to check what was returned")
	}
	return jsonBytes, nil
}
