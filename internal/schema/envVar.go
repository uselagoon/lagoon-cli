package schema

import "github.com/amazeeio/lagoon-cli/pkg/api"

// EnvKeyValue is the base type of Environment variable.
type EnvKeyValue struct {
	ID    uint                 `json:"id,omitempty"`
	Scope api.EnvVariableScope `json:"scope"`
	Name  string               `json:"name"`
	Value string               `json:"value"`
}

// EnvVariableInput is based on the Lagoon API type.
type EnvVariableInput struct {
	EnvKeyValue
	Type   api.EnvVariableType `json:"type"`
	TypeID uint                `json:"typeId"`
}
