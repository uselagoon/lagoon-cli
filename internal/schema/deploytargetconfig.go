package schema

// DeployTargetConfig .
type DeployTargetConfig struct {
	ID                         uint         `json:"id"`
	Project                    Project      `json:"project"`
	Weight                     uint         `json:"weight"`
	Branches                   string       `json:"branches"`
	Pullrequests               string       `json:"pullrequests"`
	DeployTarget               DeployTarget `json:"deployTarget"`
	DeployTargetProjectPattern string       `json:"deployTargetProjectPattern"`
}

// DeployTarget .
type DeployTarget struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	FriendlyName  string `json:"friendlyName"`
	CloudProvider string `json:"cloudProvider"`
	CloudRegion   string `json:"cloudRegion"`
	SSHHost       string `json:"sshHost"`
	SSHPort       string `json:"sshPort"`
}

// AddDeployTargetConfigInput .
type AddDeployTargetConfigInput struct {
	ID                         uint   `json:"id,omitempty"`
	Project                    uint   `json:"project,omitempty"`
	Weight                     uint   `json:"weight,omitempty"`
	Branches                   string `json:"branches,omitempty"`
	Pullrequests               string `json:"pullrequests,omitempty"`
	DeployTarget               uint   `json:"deployTarget,omitempty"`
	DeployTargetProjectPattern string `json:"deployTargetProjectPattern,omitempty"`
}

// UpdateDeployTargetConfigInput .
type UpdateDeployTargetConfigInput struct {
	ID                         uint   `json:"id,omitempty"`
	Weight                     uint   `json:"weight,omitempty"`
	Branches                   string `json:"branches,omitempty"`
	Pullrequests               string `json:"pullrequests,omitempty"`
	DeployTarget               uint   `json:"deployTarget,omitempty"`
	DeployTargetProjectPattern string `json:"deployTargetProjectPattern,omitempty"`
}

// DeleteDeployTargetConfig is the response.
type DeleteDeployTargetConfig struct {
	DeleteDeployTargetConfig string `json:"deleteDeployTargetConfig"`
}
