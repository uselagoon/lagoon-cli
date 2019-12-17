package users

import (
	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/graphql"
)

// AddUser function
func AddUser(user api.User) ([]byte, error) {
	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}
	customReq := api.CustomRequest{
		Query: `mutation addUser ($firstname: String, $lastname: String, $email: String!) {
				addUser(input:{firstName: $firstname, lastName: $lastname, email: $email}) {
					id
				}
			}`,
		Variables: map[string]interface{}{
			"firstname": user.FirstName,
			"lastname":  user.LastName,
			"email":     user.Email,
		},
		MappedResult: "addUser",
	}
	returnResult, err := lagoonAPI.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// AddSSHKeyToUser function
func AddSSHKeyToUser(user api.User, sshKey api.SSHKey) ([]byte, error) {
	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}
	customReq := api.CustomRequest{
		Query: `mutation addSshKey ($email: String!, $keyname: String!, $keyvalue: String!, $keytype: SshKeyType!) {
				addSshKey(input:{
					user: {
						email: $email
					}
					name: $keyname
					keyValue: $keyvalue
					keyType: $keytype
				}) {
					id
				}
			}`,
		Variables: map[string]interface{}{
			"email":    user.Email,
			"keyname":  sshKey.Name,
			"keyvalue": sshKey.KeyValue,
			"keytype":  sshKey.KeyType,
		},
		MappedResult: "addSshKey",
	}
	returnResult, err := lagoonAPI.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}

// DeleteUser function
func DeleteUser(user api.User) ([]byte, error) {
	// set up a lagoonapi client
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}
	customReq := api.CustomRequest{
		Query: `mutation deleteUser ($email: String!) {
				deleteUser(input:{user: {email: $email}})
			}`,
		Variables: map[string]interface{}{
			"email": user.Email,
		},
		MappedResult: "deleteUser",
	}
	returnResult, err := lagoonAPI.Request(customReq)
	if err != nil {
		return []byte(""), err
	}
	return returnResult, nil
}
