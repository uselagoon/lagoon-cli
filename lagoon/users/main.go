package users

import (
	"github.com/amazeeio/lagoon-cli/api"
)

type ExtendedSSHKey struct {
	*api.SSHKey
	Email string `json:"email"`
}
type UserGroup struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Role string `json:"role"`
}

type UsersData struct {
	User []UserData
}

type UserData struct {
	ID        string       `json:"id"`
	Email     string       `json:"email"`
	FirstName string       `json:"firstName"`
	LastName  string       `json:"lastName"`
	SSHKeys   []api.SSHKey `json:"sshKeys"`
	Groups    []UserGroup  `json:"groups"`
}

type GroupMembers []struct {
	ID      string   `json:"id"`
	Members []Member `json:"members"`
	Name    string   `json:"name"`
}

type Member struct {
	Role string `json:"role"`
	User struct {
		ID        string       `json:"id"`
		Email     string       `json:"email"`
		FirstName string       `json:"firstName"`
		SSHKeys   []api.SSHKey `json:"sshKeys"`
		LastName  string       `json:"lastName"`
	} `json:"user"`
}

func returnNonEmptyString(value string) string {
	if len(value) == 0 {
		value = "-"
	}
	return value
}

func (ud *UsersData) AddItem(userData UserData) {
	ud.User = append(ud.User, userData)
}

func distinctObjects(objs []UserData) (distinctedObjs []UserData) {
	var output []UserData
	for i := range objs {
		if output == nil || len(output) == 0 {
			output = append(output, objs[i])
		} else {
			founded := false
			for j := range output {
				if output[j].ID == objs[i].ID {
					founded = true
				}
			}
			if !founded {
				output = append(output, objs[i])
			}
		}
	}
	return output
}

func distinctKeys(objs []ExtendedSSHKey) (distinctedObjs []ExtendedSSHKey) {
	var output []ExtendedSSHKey
	for i := range objs {
		if output == nil || len(output) == 0 {
			output = append(output, objs[i])
		} else {
			founded := false
			for j := range output {
				if output[j].Email == objs[i].Email {
					founded = true
				}
			}
			if !founded {
				output = append(output, objs[i])
			}
		}
	}
	return output
}
