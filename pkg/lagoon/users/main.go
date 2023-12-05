package users

import (
	"github.com/uselagoon/lagoon-cli/internal/lagoon"
	"github.com/uselagoon/lagoon-cli/pkg/api"
	"github.com/uselagoon/lagoon-cli/pkg/graphql"
)

// Users .
type Users struct {
	debug bool
	api   api.Client
}

// Client .
type Client interface {
	AddGroup(api.Group) ([]byte, error)
	AddUserToGroup(api.UserGroupRole) ([]byte, error)
	AddProjectToGroup(api.ProjectGroups) ([]byte, error)
	RemoveUserFromGroup(api.UserGroup) ([]byte, error)
	RemoveGroupsFromProject(api.ProjectGroups) ([]byte, error)
	DeleteGroup(api.Group) ([]byte, error)
	ListUsers(string) ([]byte, error)
	AddUser(api.User) ([]byte, error)
	AddSSHKeyToUser(api.User, api.SSHKey) ([]byte, error)
	DeleteUser(api.User) ([]byte, error)
	ModifyUser(api.User, api.User) ([]byte, error)
	ListGroups(string) ([]byte, error)
	ListGroupProjects(string, bool) ([]byte, error)
}

// New .
func New(lc *lagoon.Config, debug bool) (Client, error) {
	lagoonAPI, err := graphql.LagoonAPI(lc, debug)
	if err != nil {
		return &Users{}, err
	}
	return &Users{
		debug: debug,
		api:   lagoonAPI,
	}, nil

}

var noDataError = "no data returned from the lagoon api"

// ExtendedSSHKey .
type ExtendedSSHKey struct {
	*api.SSHKey
	Email string `json:"email"`
}

// UserGroup .
type UserGroup struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Role string `json:"role"`
}

// Data .
type Data struct {
	User []UserData
}

// UserData .
type UserData struct {
	ID        string       `json:"id"`
	Email     string       `json:"email"`
	FirstName string       `json:"firstName"`
	LastName  string       `json:"lastName"`
	SSHKeys   []api.SSHKey `json:"sshKeys"`
	Groups    []UserGroup  `json:"groups"`
}

// GroupMembers .
type GroupMembers []struct {
	ID      string   `json:"id"`
	Members []Member `json:"members"`
	Name    string   `json:"name"`
}

// Member .
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

// AddItem .
func (ud *Data) AddItem(userData UserData) {
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
