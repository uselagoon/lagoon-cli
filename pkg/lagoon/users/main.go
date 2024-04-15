package users

import (
	"slices"

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
	ids := []string{}
	for _, e := range objs {
		if !slices.Contains(ids, e.ID) {
			ids = append(ids, e.ID)
			output = append(output, e)
		}
	}
	return output
}

func distinctKeys(objs []ExtendedSSHKey) (distinctedObjs []ExtendedSSHKey) {
	var output []ExtendedSSHKey
	ids := []string{}
	for _, e := range objs {
		if !slices.Contains(ids, e.Email) {
			ids = append(ids, e.Email)
			output = append(output, e)
		}
	}
	return output
}
