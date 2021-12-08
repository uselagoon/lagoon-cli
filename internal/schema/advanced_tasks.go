package schema

type AdvancedTaskDefinitionInput struct {
	ID                              uint                             `json:"id,omitempty"`
	Name                            string                           `json:"name,omitempty"`
	Description                     string                           `json:"description,omitempty"`
	Image                           string                           `json:"image,omitempty"`
	Type                            AdvancedTaskDefinitionType       `json:"type,omitempty"`
	Service                         string                           `json:"service,omitempty"`
	Environment                     int                              `json:"environment,omitempty"`
	Project                         int                              `json:"project,omitempty"`
	GroupName                       string                           `json:"groupName,omitempty"`
	Permission                      string                           `json:"permission,omitempty"`
	AdvancedTaskDefinitionArguments []AdvancedTaskDefinitionArgument `json:"arguments,omitempty"`
	Command                         string                           `json:"command,omitempty"`
}

type AdvancedTaskDefinitionType string

// AdvancedTaskDefinition An Advanced Task Definition is based on the Lagoon API GraphQL type.
type AdvancedTaskDefinition struct {
	ID          uint                       `json:"id,omitempty" yaml:"id,omitempty"`
	Name        string                     `json:"name,omitempty" yaml:"name,omitempty"`
	Description string                     `json:"description,omitempty" yaml:"description,omitempty"`
	Type        AdvancedTaskDefinitionType `json:"type,omitempty" yaml:"type,omitempty"`
	Command     string                     `json:"command,omitempty" yaml:"command,omitempty"`
	Image       string                     `json:"image,omitempty" yaml:"image,omitempty"`
	Service     string                     `json:"service,omitempty" yaml:"service,omitempty"`
	Environment int                        `json:"environment,omitempty" yaml:"environment,omitempty"`
	Project     int                        `json:"project,omitempty" yaml:"project,omitempty"`
	GroupName   string                     `json:"groupName,omitempty" yaml:"groupName,omitempty"`
	Permission  string                     `json:"permission,omitempty" yaml:"permission,omitempty"`
	Created     string                     `json:"created,omitempty" yaml:"created,omitempty"`
	Deleted     string                     `json:"deleted,omitempty" yaml:"deleted,omitempty"`
	//AdvancedTaskDefinitionArguments []AdvancedTaskDefinitionArgument `json:"advancedTaskDefinitionArguments,omitempty" yaml:"advancedTaskDefinitionArguments,omitempty"`
}

// AdvancedTask An Advanced Task is based on the Lagoon API GraphQL type.
type AdvancedTask struct {
	ID           uint        `json:"id,omitempty"`
	Name         string      `json:"name,omitempty"`
	Status       string      `json:"status,omitempty"`
	Created      string      `json:"created,omitempty"`
	Started      string      `json:"started,omitempty"`
	Completed    string      `json:"completed,omitempty"`
	Environment  Environment `json:"environment,omitempty"`
	Service      string      `json:"service,omitempty"`
	AdvancedTask string      `json:"string,omitempty"`
	RemoteID     string      `json:"remoteId,omitempty"`
	Logs         string      `json:"logs,omitempty"`
	Files        []File      `json:"files,omitempty"`
}

type File struct {
	ID       uint `json:"id,omitempty"`
	Filename uint `json:"filename,omitempty"`
	Download uint `json:"download,omitempty"`
	Created  uint `json:"created,omitempty"`
}

type AdvancedTaskDefinitionArgument struct {
	ID                     uint   `json:"id,omitempty"`
	Name                   string `json:"name,omitempty"`
	Type                   string `json:"type,omitempty"`
	AdvancedTaskDefinition int    `json:"advancedTaskDefinition,omitempty"`
}
