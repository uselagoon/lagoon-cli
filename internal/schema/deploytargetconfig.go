package schema

// DeployTargetConfig .
type DeployTargetConfig struct {
	ID                         uint         `json:"id"`
	Project                    Project      `json:"project,omitempty"`
	Weight                     uint         `json:"weight,omitempty"`
	Branches                   string       `json:"branches,omitempty"`
	Pullrequests               string       `json:"pullrequests,omitempty"`
	DeployTarget               DeployTarget `json:"deployTarget,omitempty"`
	DeployTargetProjectPattern string       `json:"deployTargetProjectPattern,omitempty"`
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
