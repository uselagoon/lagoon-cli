package schema

import (
	"github.com/guregu/null"
)

type DeployTarget struct {
	AddDeployTargetInput
}

// AddDeployTargetInput is based on the input to addDeployTarget.
type AddDeployTargetInput struct {
	ID               uint   `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	ConsoleURL       string `json:"consoleUrl,omitempty"`
	Token            string `json:"token,omitempty"`
	RouterPattern    string `json:"routerPattern,omitempty"`
	SSHHost          string `json:"sshHost,omitempty"`
	SSHPort          string `json:"sshPort,omitempty"`
	BuildImage       string `json:"buildImage,omitempty"`
	Created          string `json:"created,omitempty"`
	MonitoringConfig string `json:"monitoringConfig,omitempty"`
	FriendlyName     string `json:"friendlyName,omitempty"`
	CloudProvider    string `json:"cloudProvider,omitempty"`
	CloudRegion      string `json:"cloudRegion,omitempty"`
}

// UpdateDeployTargetInput is based on the input to addDeployTarget.
type UpdateDeployTargetInput struct {
	AddDeployTargetInput
	// `null` is valid graphql, use a pointer to allow `nil` to be empty
	BuildImage *null.String `json:"buildImage,omitempty"`
}

type UpdateDeployTargetResponse struct {
	DeployTarget
}

type AddDeployTargetResponse struct {
	DeployTarget
}

type DeleteDeployTargetInput struct {
	Name string `json:"name,omitempty"`
	ID   uint   `json:"id,omitempty"`
}

type DeleteDeployTargetResponse struct {
	DeleteDeployTarget string `json:"deleteDeployTarget,omitempty"`
}
