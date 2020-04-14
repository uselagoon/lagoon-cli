package schema

import "github.com/google/uuid"

// AddUserInput is based on the Lagoon API type.
type AddUserInput struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Comment   string `json:"comment,omitempty"`
	GitlabID  uint   `json:"gitlabId,omitempty"`
}

// User provides for unmarshalling the users contained withing a Group.
type User struct {
	AddUserInput
	ID      *uuid.UUID `json:"id,omitempty"`
	SSHKeys []SSHKey   `json:"sshKeys,omitempty"`
}
