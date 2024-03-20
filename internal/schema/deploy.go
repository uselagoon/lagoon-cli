package schema

// DeployEnvironmentLatestInput is used as the input for deploying an environment.
type DeployEnvironmentLatestInput struct {
	Environment    EnvironmentInput   `json:"environment"`
	ReturnData     bool               `json:"returnData"`
	BuildVariables []EnvKeyValueInput `json:"buildVariables,omitempty"`
}

// DeployEnvironmentLatest is the response.
type DeployEnvironmentLatest struct {
	DeployEnvironmentLatest string `json:"deployEnvironmentLatest"`
}

// DeployEnvironmentPullrequestInput is used as the input for deploying a pull request.
type DeployEnvironmentPullrequestInput struct {
	Project        ProjectInput       `json:"project"`
	Number         uint               `json:"number"`
	Title          string             `json:"title"`
	BaseBranchName string             `json:"baseBranchName"`
	BaseBranchRef  string             `json:"baseBranchRef"`
	HeadBranchName string             `json:"headBranchName"`
	HeadBranchRef  string             `json:"headBranchRef"`
	ReturnData     bool               `json:"returnData"`
	BuildVariables []EnvKeyValueInput `json:"buildVariables,omitempty"`
}

// DeployEnvironmentPullrequest is the response.
type DeployEnvironmentPullrequest struct {
	DeployEnvironmentPullrequest string `json:"deployEnvironmentPullrequest"`
}

// DeployEnvironmentBranchInput is used as the input for deploying a branch.
type DeployEnvironmentBranchInput struct {
	Project        string             `json:"project"`
	Branch         string             `json:"branch"`
	BranchRef      string             `json:"branchRef"`
	ReturnData     bool               `json:"returnData"`
	BuildVariables []EnvKeyValueInput `json:"buildVariables,omitempty"`
}

// DeployEnvironmentBranch is the response.
type DeployEnvironmentBranch struct {
	DeployEnvironmentBranch string `json:"deployEnvironmentBranch"`
}

// DeployEnvironmentPromoteInput is used as the input for promoting one environment to another.
type DeployEnvironmentPromoteInput struct {
	Project                string             `json:"project"`
	SourceEnvironment      string             `json:"sourceEnvironment"`
	DestinationEnvironment string             `json:"destinationEnvironment"`
	BuildVariables         []EnvKeyValueInput `json:"buildVariables,omitempty"`
	ReturnData             bool               `json:"returnData"`
}

// DeployEnvironmentPromote is the response.
type DeployEnvironmentPromote struct {
	DeployEnvironmentPromote string `json:"deployEnvironmentPromote"`
}
