package users

import (
	"github.com/uselagoon/lagoon-cli/pkg/api"
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
