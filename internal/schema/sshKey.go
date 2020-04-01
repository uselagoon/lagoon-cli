package schema

import "github.com/amazeeio/lagoon-cli/pkg/api"

// SSHKey is the basic SSH key information, used by both config and API data.
type SSHKey struct {
	Name     string         `json:"name"`
	KeyValue string         `json:"keyValue"`
	KeyType  api.SSHKeyType `json:"keyType"`
}

// AddSSHKeyInput is based on the Lagoon API type.
type AddSSHKeyInput struct {
	SSHKey
	UserEmail string `json:"userEmail"`
}
