package schema

type WorkflowInput struct {
	ID                     uint      `json:"id,omitempty"`
	Name                   string    `json:"name,omitempty"`
	Event                  EventType `json:"event,omitempty"`
	Project                int       `json:"project,omitempty"`
	AdvancedTaskDefinition int       `json:"advancedTaskDefinition,omitempty"`
}

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

type EventType string
type AdvancedTaskDefinitionType string

const (
	APIDeployEnvironmentLatest EventType = "api:deployEnvironmentLatest"
	APIDeployEnvironmentBranch EventType = "api:deployEnvironmentBranch"
	APIDeleteEnvironment       EventType = "api:deleteEnvironment"

	DeployOSFinished         EventType = "task:builddeploy-openshift:complete"
	DeployKubernetesFinished EventType = "task:builddeploy-kubernetes:complete"
	RemoveOSFinished         EventType = "task:remove-openshift:finished"
	RemoveKubernetesFinished EventType = "task:remove-kubernetes:finished"

	DeployErrorRemoveKubernetes      EventType = "task:remove-kubernetes:error"
	DeployErrorRemoveOS              EventType = "task:remove-openshift:error"
	DeployErrorBuildDeployKubernetes EventType = "task:builddeploy-kubernetes:failed"
	DeployErrorBuildDeployOS         EventType = "task:builddeploy-openshift:failed"

	GithubPush              EventType = "github:push:handled"
	GithubPushSkip          EventType = "github:push:skipped"
	GithubPROpened          EventType = "github:pull_request:opened:handled"
	GithubPRUpdated         EventType = "github:pull_request:synchronize:handled"
	GithubPRClosed          EventType = "github:pull_request:closed:handled"
	GithubDeleteEnvironment EventType = "github:delete:handled"
	GithubPRNotDeleted      EventType = "github:pull_request:closed:CannotDeleteProductionEnvironment"
	GithubPushNotDeleted    EventType = "github:push:CannotDeleteProductionEnvironment"

	GitlabPush              EventType = "gitlab:push:handled"
	GitlabPushSkip          EventType = "gitlab:push:skipped"
	GitlabPROpened          EventType = "gitlab:pull_request:opened:handled"
	GitlabPRUpdated         EventType = "gitlab:pull_request:synchronize:handled"
	GitlabPRClosed          EventType = "gitlab:pull_request:closed:handled"
	GitlabDeleteEnvironment EventType = "gitlab:delete:handled"
	GitlabPushNotDeleted    EventType = "gitlab:push:CannotDeleteProductionEnvironment"

	BitbucketPush              EventType = "bitbucket:repo:push:handled"
	BitbucketPushSkip          EventType = "bitbucket:push:skipped"
	BitbucketPROpened          EventType = "bitbucket:pullrequest:created:handled"
	BitbucketPRCreatedOpened   EventType = "bitbucket:pullrequest:created:opened:handled"
	BitbucketPRUpdated         EventType = "bitbucket:pullrequest:updated:handled"
	BitbucketPRUpdatedOpened   EventType = "bitbucket:pullrequest:updated:opened:handled"
	BitbucketPRFulfilled       EventType = "bitbucket:pullrequest:fulfilled:handled"
	BitbucketPRRejected        EventType = "bitbucket:pullrequest:rejected:handled"
	BitbucketDeleteEnvironment EventType = "bitbucket:delete:handled"
	BitbucketPushNotDeleted    EventType = "bitbucket:repo:push:CannotDeleteProductionEnvironment"
)

// WorkflowResponse A workflow response is based on the Lagoon API GraphQL type.
type WorkflowResponse struct {
	ID                     uint                           `json:"id,omitempty" yaml:"id,omitempty"`
	Name                   string                         `json:"name,omitempty" yaml:"name,omitempty"`
	Event                  string                         `json:"event,omitempty" yaml:"event,omitempty"`
	Project                int                            `json:"project,omitempty" yaml:"project,omitempty"`
	AdvancedTaskDefinition AdvancedTaskDefinitionResponse `json:"advancedTaskDefinition,omitempty" yaml:"advancedTaskDefinition,omitempty"`
}

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

type File struct {
	ID       uint `json:"id,omitempty"`
	Filename uint `json:"filename,omitempty"`
	Download uint `json:"download,omitempty"`
	Created  uint `json:"created,omitempty"`
}

type AdvancedTaskDefinitionArgument struct {
	ID                     uint   `json:"id,omitempty" yaml:"id,omitempty"`
	Name                   string `json:"name,omitempty" yaml:"name,omitempty"`
	Type                   string `json:"type,omitempty" yaml:"type,omitempty"`
	AdvancedTaskDefinition int    `json:"advancedTaskDefinition,omitempty" yaml:"advancedTaskDefinition,omitempty"`
}
