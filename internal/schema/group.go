package schema

import (
	"github.com/amazeeio/lagoon-cli/pkg/api"
	"github.com/google/uuid"
)

// AddGroupInput is based on the input to addGroup.
type AddGroupInput struct {
	Name        string      `json:"name"`
	ParentGroup *GroupInput `json:"parentGroup,omitempty"`
}

// Group provides for unmarshalling the groups contained with a Project.
type Group struct {
	AddGroupInput
	ID      *uuid.UUID `json:"id,omitempty"`
	Members []struct {
		User User          `json:"user"`
		Role api.GroupRole `json:"role"`
	} `json:"members,omitempty"`
}

// GroupConfig embeds AddGroupInput as well as a list of members.
type GroupConfig struct {
	AddGroupInput
	Users []UserRoleConfig `json:"users,omitempty"`
}

// GroupInput is based on the Lagoon API type.
type GroupInput ProjectInput

// UserGroupRoleInput is based on the Lagoon API type.
type UserGroupRoleInput struct {
	UserEmail string        `json:"userEmail"`
	GroupName string        `json:"groupName"`
	GroupRole api.GroupRole `json:"groupRole"`
}

// UserRoleConfig stores a user/role config within a group.
type UserRoleConfig struct {
	Email string        `json:"email"`
	Role  api.GroupRole `json:"role"`
}
