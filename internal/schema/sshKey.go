package schema

import (
	api "github.com/uselagoon/machinery/api/schema"
)

// SSHKey is the basic SSH key information, used by both config and API data.
// @TODO: once Lagoon API returns proper TZ, fix up `Created` to time.Time.
type SSHKey struct {
	Name           string         `json:"name"`
	KeyValue       string         `json:"keyValue"`
	Created        string         `json:"created"`
	KeyType        api.SSHKeyType `json:"keyType"`
	KeyFingerprint string         `json:"keyFingerprint"`
}

// AddSSHKeyInput is based on the Lagoon API type.
type AddSSHKeyInput struct {
	SSHKey
	UserEmail string `json:"userEmail"`
}
