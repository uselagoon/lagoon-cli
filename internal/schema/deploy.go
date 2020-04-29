package schema

// DeployEnvironmentLatestInput is based on the Lagoon API type.
type DeployEnvironmentLatestInput struct {
	Environment EnvironmentInput `json:"environment"`
}

// DeployEnvironmentLatest is the response.
type DeployEnvironmentLatest struct {
	DeployEnvironmentLatest string `json:"deployEnvironmentLatest"`
}

// DeployEnvironmentPullrequestInput is based on the Lagoon API type.
type DeployEnvironmentPullrequestInput struct {
	Project        ProjectInput `json:"project"`
	Number         uint         `json:"number"`
	Title          string       `json:"title"`
	BaseBranchName string       `json:"baseBranchName"`
	BaseBranchRef  string       `json:"baseBranchRef"`
	HeadBranchName string       `json:"headBranchName"`
	HeadBranchRef  string       `json:"headBranchRef"`
}

// DeployEnvironmentPullrequest is the response.
type DeployEnvironmentPullrequest struct {
	DeployEnvironmentPullrequest string `json:"deployEnvironmentPullrequest"`
}
