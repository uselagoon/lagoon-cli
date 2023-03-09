package schema

type AdvancedTaskDefinitionInput struct {
	ID          uint                             `json:"id,omitempty"`
	Name        string                           `json:"name,omitempty"`
	Description string                           `json:"description,omitempty"`
	Type        AdvancedTaskDefinitionType       `json:"type,omitempty"`
	Command     string                           `json:"command,omitempty"`
	Image       string                           `json:"image,omitempty"`
	Arguments   []AdvancedTaskDefinitionArgument `json:"advancedTaskDefinitionArguments,omitempty" yaml:"advancedTaskDefinitionArguments,omitempty"`
	Service     string                           `json:"service,omitempty"`
	GroupName   string                           `json:"groupName,omitempty"`
	Project     int                              `json:"project,omitempty"`
	Environment int                              `json:"environment,omitempty"`
	Permission  string                           `json:"permission,omitempty"`
}

type AdvancedTaskDefinitionType string

// AdvancedTaskDefinitionResponse An Advanced Task Definition is based on the Lagoon API GraphQL type.
type AdvancedTaskDefinitionResponse struct {
	ID          uint                             `json:"id,omitempty" yaml:"id,omitempty"`
	Name        string                           `json:"name,omitempty" yaml:"name,omitempty"`
	Description string                           `json:"description,omitempty" yaml:"description,omitempty"`
	Type        AdvancedTaskDefinitionType       `json:"type,omitempty" yaml:"type,omitempty"`
	Command     string                           `json:"command,omitempty" yaml:"command,omitempty"`
	Image       string                           `json:"image,omitempty" yaml:"image,omitempty"`
	Arguments   []AdvancedTaskDefinitionArgument `json:"advancedTaskDefinitionArguments,omitempty" yaml:"advancedTaskDefinitionArguments,omitempty"`
	Service     string                           `json:"service,omitempty" yaml:"service,omitempty"`
	Environment int                              `json:"environment,omitempty" yaml:"environment,omitempty"`
	Project     int                              `json:"project,omitempty" yaml:"project,omitempty"`
	GroupName   string                           `json:"groupName,omitempty" yaml:"groupName,omitempty"`
	Permission  string                           `json:"permission,omitempty" yaml:"permission,omitempty"`
	Created     string                           `json:"created,omitempty" yaml:"created,omitempty"`
	Deleted     string                           `json:"deleted,omitempty" yaml:"deleted,omitempty"`
}

type AdvancedTaskDefinitionArgument struct {
	ID                     uint   `json:"id,omitempty" yaml:"id,omitempty"`
	Name                   string `json:"name,omitempty" yaml:"name,omitempty"`
	Type                   string `json:"type,omitempty" yaml:"type,omitempty"`
	AdvancedTaskDefinition int    `json:"advancedTaskDefinition,omitempty" yaml:"advancedTaskDefinition,omitempty"`
}
