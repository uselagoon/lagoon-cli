package schema

// EnvKeyValue is the base type of Environment variable.
type EnvKeyValue struct {
	ID    uint             `json:"id,omitempty"`
	Scope EnvVariableScope `json:"scope"`
	Name  string           `json:"name"`
	Value string           `json:"value"`
}

// EnvVariableInput is based on the Lagoon API type.
type EnvVariableInput struct {
	EnvKeyValue
	Type   EnvVariableType `json:"type"`
	TypeID uint            `json:"typeId"`
}

type EnvVar struct {
	Scope EnvVariableScope `json:"scope"`
	Name  string           `json:"name"`
	Value string           `json:"value"`
}

type EnvVariableByNameInput struct {
	Environment string           `json:"environment,omitempty"`
	Project     string           `json:"project"`
	Scope       EnvVariableScope `json:"scope,omitempty"`
	Name        string           `json:"name"`
	Value       string           `json:"value"`
}

type DeleteEnvVariableByNameInput struct {
	Environment string `json:"environment,omitempty"`
	Project     string `json:"project"`
	Name        string `json:"name"`
}

type EnvVariableByProjectEnvironmentNameInput struct {
	Environment string `json:"environment,omitempty"`
	Project     string `json:"project"`
}

type UpdateEnvVarResponse struct {
	EnvKeyValue
}

type DeleteEnvVarResponse struct {
	DeleteEnvVar string `json:"deleteEnvVariableByName,omitempty"`
}

// EnvKeyValueInput  is based on the Lagoon API type.
type EnvKeyValueInput struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}
