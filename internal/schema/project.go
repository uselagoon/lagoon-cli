package schema

import "github.com/uselagoon/lagoon-cli/pkg/api"

// AddProjectInput is based on the Lagoon API type.
type AddProjectInput struct {
	ID                           uint                `json:"id,omitempty"`
	Name                         string              `json:"name"`
	GitURL                       string              `json:"gitUrl"`
	Subfolder                    string              `json:"subfolder,omitempty"`
	Openshift                    uint                `json:"openshift"`
	OpenshiftProjectPattern      string              `json:"openshiftProjectPattern,omitempty"`
	Branches                     string              `json:"branches,omitempty"`
	PullRequests                 string              `json:"pullrequests,omitempty"`
	ProductionEnvironment        string              `json:"productionEnvironment"`
	StandbyProductionEnvironment string              `json:"standbyProductionEnvironment,omitempty"`
	Availability                 ProjectAvailability `json:"availability,omitempty"`
	// AutoIdle and StorageCalc can't be omitempty because their zero-values are
	// significant - Lagoon uses it as a boolean (0/1).
	AutoIdle                     uint   `json:"autoIdle"`
	StorageCalc                  uint   `json:"storageCalc"`
	DevelopmentEnvironmentsLimit uint   `json:"developmentEnvironmentsLimit,omitempty"`
	PrivateKey                   string `json:"privateKey,omitempty"`
}

// Project is the Lagoon API Project object.
type Project struct {
	AddProjectInput
	Environments []EnvironmentConfig `json:"environments,omitempty"`
	EnvVariables []EnvKeyValue       `json:"envVariables,omitempty"`
	// Notifications is unmarshalled during export.
	Notifications *Notifications `json:"notifications,omitempty"`
	// Openshift is unmarshalled during export.
	OpenshiftID *OpenshiftID `json:"openshift,omitempty"`
	// Groups are unmarshalled during export.
	Groups             *Groups              `json:"groups,omitempty"`
	DeployTargetConfig []DeployTargetConfig `json:"deployTargetConfigs,omitempty"`
}

// ProjectConfig contains project configuration.
type ProjectConfig struct {
	Project
	// ProjectNotifications are (un)marshalled during import.
	Notifications *ProjectNotifications `json:"notifications,omitempty"`
	// Group are (un)marshalled during import.
	Groups []string `json:"groups,omitempty"`
	// Users are members of the project.
	// Note that in Lagoon this is implemented as being a member of the
	// project-<projectname> group.
	Users []UserRoleConfig `json:"users,omitempty"`
}

// ProjectNotifications lists the notifications for a project within a
// ProjectConfig.
type ProjectNotifications struct {
	Slack          []string `json:"slack,omitempty"`
	RocketChat     []string `json:"rocketChat,omitempty"`
	Email          []string `json:"email,omitempty"`
	MicrosoftTeams []string `json:"microsoftTeams,omitempty"`
	Webhook        []string `json:"webhook,omitempty"`
}

// OpenshiftID is unmarshalled during export.
type OpenshiftID struct {
	ID uint `json:"id,omitempty"`
}

// ProjectGroupsInput is based on the input to
// addGroupsToProject.
type ProjectGroupsInput struct {
	Project ProjectInput `json:"project"`
	Groups  []GroupInput `json:"groups"`
}

// ProjectInput is based on the Lagoon API type.
type ProjectInput struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// AddNotificationToProjectInput is based on the input to
// addNotificationToProject.
type AddNotificationToProjectInput struct {
	Project          string               `json:"project"`
	NotificationType api.NotificationType `json:"notificationType"`
	NotificationName string               `json:"notificationName"`
}

// RemoveNotificationFromProjectInput is based on the input to
// removeNotificationFromProject.
type RemoveNotificationFromProjectInput struct {
	Project          string               `json:"project"`
	NotificationType api.NotificationType `json:"notificationType"`
	NotificationName string               `json:"notificationName"`
}

// ProjectMetadata .
type ProjectMetadata struct {
	Project
	Metadata map[string]string `json:"metadata"`
}
