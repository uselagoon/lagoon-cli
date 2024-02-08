package api

// SSHKeyType .
type SSHKeyType string

// . .
const (
	SSHRsa      SSHKeyType = "SSH_RSA"
	SSHEd25519  SSHKeyType = "SSH_ED25519"
	SSHECDSA256 SSHKeyType = "ECDSA_SHA2_NISTP256"
	SSHECDSA384 SSHKeyType = "ECDSA_SHA2_NISTP384"
	SSHECDSA521 SSHKeyType = "ECDSA_SHA2_NISTP521"
)

// DeployType .
type DeployType string

// . .
const (
	Branch      DeployType = "BRANCH"
	PullRequest DeployType = "PULLREQUEST"
	Promote     DeployType = "PROMOTE"
)

// EnvType .
type EnvType string

// . .
const (
	ProductionEnv  EnvType = "PRODUCTION"
	DevelopmentEnv EnvType = "DEVELOPMENT"
)

// NotificationType .
type NotificationType string

// . .
const (
	SlackNotification          NotificationType = "SLACK"
	RocketChatNotification     NotificationType = "ROCKETCHAT"
	EmailNotification          NotificationType = "EMAIL"
	MicrosoftTeamsNotification NotificationType = "MICROSOFTTEAMS"
	WebhookNotification        NotificationType = "WEBHOOK"
)

// DeploymentStatusType .
type DeploymentStatusType string

// . .
const (
	NewDeploy       DeploymentStatusType = "NEW"
	PendingDeploy   DeploymentStatusType = "PENDING"
	RunningDeploy   DeploymentStatusType = "RUNNING"
	CancelledDeploy DeploymentStatusType = "CANCELLED"
	ErrorDeploy     DeploymentStatusType = "ERROR"
	FailedDeploy    DeploymentStatusType = "FAILED"
	CompleteDeploy  DeploymentStatusType = "COMPLETE"
)

// EnvVariableType .
type EnvVariableType string

// . .
const (
	ProjectVar     EnvVariableType = "PROJECT"
	EnvironmentVar EnvVariableType = "ENVIRONMENT"
)

// EnvVariableScope .
type EnvVariableScope string

// . .
const (
	BuildVar                     EnvVariableScope = "BUILD"
	RuntimeVar                   EnvVariableScope = "RUNTIME"
	GlobalVar                    EnvVariableScope = "GLOBAL"
	InternalContainerRegistryVar EnvVariableScope = "INTERNAL_CONTAINER_REGISTRY"
	ContainerRegistryVar         EnvVariableScope = "CONTAINER_REGISTRY"
)

// TaskStatusType .
type TaskStatusType string

// . .
const (
	ActiveTask    TaskStatusType = "ACTIVE"
	SucceededTask TaskStatusType = "SUCCEEDED"
	FailedTask    TaskStatusType = "FAILED"
)

// RestoreStatusType .
type RestoreStatusType string

// Pending .
const (
	PendingRestore    RestoreStatusType = "PENDING"
	SuccessfulRestore RestoreStatusType = "SUCCESSFUL"
	FailedRestore     RestoreStatusType = "FAILED"
)

// EnvOrderType .
type EnvOrderType string

// . .
const (
	NameEnvOrder    EnvOrderType = "NAME"
	UpdatedEnvOrder EnvOrderType = "UPDATED"
)

// ProjectOrderType .
type ProjectOrderType string

// . .
const (
	NameProjectOrder    ProjectOrderType = "NAME"
	CreatedProjectOrder ProjectOrderType = "UPCREATEDDATED"
)

// GroupRole .
type GroupRole string

// Guest .
const (
	GuestRole      GroupRole = "GUEST"
	ReporterRole   GroupRole = "REPORTER"
	DeveloperRole  GroupRole = "DEVELOPER"
	MaintainerRole GroupRole = "MAINTAINER"
	OwnerRole      GroupRole = "OWNER"
)

// File .
type File struct {
	ID       int    `json:"id,omitempty"`
	Filename string `json:"filename,omitempty"`
	Download string `json:"download,omitempty"`
	Created  string `json:"created,omitempty"`
}

// SSHKey .
type SSHKey struct {
	ID             int        `json:"id,omitempty"`
	Name           string     `json:"name,omitempty"`
	KeyValue       string     `json:"keyValue,omitempty"`
	KeyType        SSHKeyType `json:"keyType,omitempty"`
	KeyFingerprint string     `json:"keyFingerprint,omitempty"`
}

// SSHKeyValue .
type SSHKeyValue string

// User struct.
type User struct {
	ID        string   `json:"id,omitempty"`
	Email     string   `json:"email,omitempty"`
	FirstName string   `json:"firstName,omitempty"`
	LastName  string   `json:"lastName,omitempty"`
	Comment   string   `json:"comment,omitempty"`
	GitlabID  int      `json:"gitlabId,omitempty"`
	SSHKeys   []SSHKey `json:"sshKeys,omitempty"`
	Groups    []Group  `json:"groups,omitempty"`
}

// GroupMembership .
type GroupMembership struct {
	User User      `json:"user"`
	Role GroupRole `json:"role"`
}

// Group struct.
type Group struct {
	ID      string  `json:"id,omitempty"`
	Name    string  `json:"name,omitempty"`
	Groups  []Group `json:"groups,omitempty"`
	Members string  `json:"members,omitempty"`
}

// Openshift struct.
type Openshift struct {
	ID            int    `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	ConsoleURL    string `json:"consoleUrl,omitempty"`
	Token         string `json:"token,omitempty"`
	RouterPattern string `json:"routerPattern,omitempty"`
	ProjectUser   string `json:"projectUser,omitempty"`
	SSHHost       string `json:"sshHost,omitempty"`
	SSHPort       string `json:"sshPort,omitempty"`
	Created       string `json:"created,omitempty"`
}

// NotificationRocketChat struct.
type NotificationRocketChat struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Webhook string `json:"webhook,omitempty"`
	Channel string `json:"channel,omitempty"`
}

// NotificationSlack struct.
type NotificationSlack struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Webhook string `json:"webhook,omitempty"`
	Channel string `json:"channel,omitempty"`
}

// UnassignedNotification struct.
type UnassignedNotification struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

// UpdateProject struct.
type UpdateProject struct {
	ID    int          `json:"id"`
	Patch ProjectPatch `json:"patch"`
}

// DeleteProjectResult struct.
type DeleteProjectResult struct {
	DeleteProject string `json:"deleteProject"`
}

// Projects struct.
type Projects struct {
	Projects []Project //`json:"allProjects"`
}

// Project struct.
type Project struct {
	ID                           int                   `json:"id,omitempty"`
	Name                         string                `json:"name,omitempty"`
	GitURL                       string                `json:"gitUrl,omitempty"`
	PrivateKey                   string                `json:"privateKey,omitempty"`
	PublicKey                    string                `json:"publicKey,omitempty"`
	Subfolder                    string                `json:"subfolder,omitempty"`
	RouterPattern                string                `json:"routerPattern,omitempty"`
	Branches                     string                `json:"branches,omitempty"`
	Pullrequests                 string                `json:"pullrequests,omitempty"`
	ProductionEnvironment        string                `json:"productionEnvironment,omitempty"`
	StandbyProductionEnvironment string                `json:"standbyProductionEnvironment,omitempty"`
	AutoIdle                     *int                  `json:"autoIdle,omitempty"`
	StorageCalc                  *int                  `json:"storageCalc,omitempty"`
	OpenshiftProjectPattern      string                `json:"openshiftProjectPattern,omitempty"`
	DevelopmentEnvironmentsLimit int                   `json:"developmentEnvironmentsLimit,omitempty"`
	Created                      string                `json:"created,omitempty"`
	Openshift                    Openshift             `json:"openshift,omitempty"`
	EnvVariables                 []EnvironmentVariable `json:"envVariables,omitempty"`
	Environments                 []Environment         `json:"environments,omitempty"`
	Deployments                  []Deployment          `json:"deployments,omitempty"`
	Notifications                []interface{}         `json:"notifications,omitempty"`
	FactsUI                      *int                  `json:"factsUi,omitempty"`
	ProblemsUI                   *int                  `json:"problemsUi,omitempty"`
}

// ProjectPatch struct.
type ProjectPatch struct {
	ID                           int    `json:"id,omitempty"`
	Name                         string `json:"name,omitempty"`
	GitURL                       string `json:"gitUrl,omitempty"`
	PrivateKey                   string `json:"privateKey,omitempty"`
	Subfolder                    string `json:"subfolder,omitempty"`
	RouterPattern                string `json:"routerPattern,omitempty"`
	Branches                     string `json:"branches,omitempty"`
	Pullrequests                 string `json:"pullrequests,omitempty"`
	ProductionEnvironment        string `json:"productionEnvironment,omitempty"`
	StandbyProductionEnvironment string `json:"standbyProductionEnvironment,omitempty"`
	AutoIdle                     *int   `json:"autoIdle,omitempty"`
	StorageCalc                  *int   `json:"storageCalc,omitempty"`
	OpenshiftProjectPattern      string `json:"openshiftProjectPattern,omitempty"`
	DevelopmentEnvironmentsLimit *int   `json:"developmentEnvironmentsLimit,omitempty"`
	Openshift                    *int   `json:"openshift,omitempty"`
	FactsUI                      *int   `json:"factsUi,omitempty"`
	ProblemsUI                   *int   `json:"problemsUi,omitempty"`
	DeploymentsDisabled          *int   `json:"deploymentsDisabled,omitempty"`
}

// AddSSHKey .
type AddSSHKey struct {
	ID       string     `json:"id,omitempty"`
	Name     string     `json:"name"`
	KeyValue string     `json:"keyValue"`
	KeyType  SSHKeyType `json:"keyType"`
	User     User       `json:"user"`
}

// DeleteSSHKey .
type DeleteSSHKey struct {
	Name string `json:"name"`
}

// EnvironmentByName struct.
type EnvironmentByName struct {
	Name    string `json:"name"`
	Project int    `json:"project"`
}

// AddOrUpdateEnvironmentStorage struct.
type AddOrUpdateEnvironmentStorage struct {
	Environment            int    `json:"environment"`
	PersistentStorageClaim string `json:"persistentStorageClaim"`
	BytesUsed              int    `json:"bytesUsed"`
}

// AddBackup struct.
type AddBackup struct {
	ID          int    `json:"id,omitempty"`
	Environment string `json:"environment"`
	Source      string `json:"source"`
	BackupID    string `json:"backupId"`
	Created     string `json:"created"`
}

// DeleteBackup struct.
type DeleteBackup struct {
	BackupID string `json:"backupId"`
}

// AddRestore struct.
type AddRestore struct {
	ID              int               `json:"id,omitempty"`
	Status          RestoreStatusType `json:"status"`
	RestoreLocation string            `json:"restoreLocation"`
	Created         string            `json:"created"`
	Execute         bool              `json:"execute"`
	BackupID        string            `json:"backupId,omitempty"`
}

// UpdateRestore struct.
type UpdateRestore struct {
	BackupID string             `json:"backupId"`
	Patch    UpdateRestorePatch `json:"patch"`
}

// UpdateRestorePatch struct.
type UpdateRestorePatch struct {
	Status          string `json:"status,omitempty"`
	Created         string `json:"created,omitempty"`
	RestoreLocation string `json:"restoreLocation,omitempty"`
}

// Deployment struct.
type Deployment struct {
	ID          int                  `json:"id,omitempty"`
	Name        string               `json:"name,omitempty"`
	Status      DeploymentStatusType `json:"status,omitempty"`
	Created     string               `json:"created,omitempty"`
	Started     string               `json:"started,omitempty"`
	Completed   string               `json:"completed,omitempty"`
	Environment int                  `json:"environment,omitempty"`
	RemoteID    string               `json:"remoteId,omitempty"`
	BuildLog    string               `json:"buildLog,omitempty"`
}

// DeploymentEnv struct.
type DeploymentEnv struct {
	ID          int                  `json:"id,omitempty"`
	Name        string               `json:"name,omitempty"`
	Status      DeploymentStatusType `json:"status,omitempty"`
	Created     string               `json:"created,omitempty"`
	Started     string               `json:"started,omitempty"`
	Completed   string               `json:"completed,omitempty"`
	Environment Environment          `json:"environment,omitempty"`
	RemoteID    string               `json:"remoteId,omitempty"`
	BuildLog    string               `json:"buildLog,omitempty"`
}

// DeleteDeployment struct.
type DeleteDeployment struct {
	ID int `json:"id"`
}

// UpdateDeployment struct.
type UpdateDeployment struct {
	ID    int        `json:"id"`
	Patch Deployment `json:"patch"`
}

// Task struct.
type Task struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	Status      TaskStatusType `json:"status,omitempty"`
	Created     string         `json:"created,omitempty"`
	Started     string         `json:"started,omitempty"`
	Completed   string         `json:"completed,omitempty"`
	Environment int            `json:"environment"`
	Service     string         `json:"service,omitempty"`
	Command     string         `json:"command,omitempty"`
	RemoteID    string         `json:"remoteId,omitempty"`
	Execute     bool           `json:"execute,omitempty"`
}

// DeleteTask struct.
type DeleteTask struct {
	ID int `json:"id"`
}

// UpdateTask struct.
type UpdateTask struct {
	ID    int  `json:"id"`
	Patch Task `json:"patch"`
}

// DeleteOpenshift struct.
type DeleteOpenshift struct {
	Name string `json:"name"`
}

// AddNotificationRocketChat struct.
type AddNotificationRocketChat struct {
	Name    string `json:"name"`
	Webhook string `json:"webhook"`
	Channel string `json:"channel"`
}

// AddNotificationSlack struct.
type AddNotificationSlack struct {
	Name    string `json:"name"`
	Webhook string `json:"webhook"`
	Channel string `json:"channel"`
}

// DeleteNotificationRocketChat struct.
type DeleteNotificationRocketChat struct {
	Name string `json:"name"`
}

// DeleteNotificationSlack struct.
type DeleteNotificationSlack struct {
	Name string `json:"name"`
}

// AddNotificationToProject struct.
type AddNotificationToProject struct {
	Project          string           `json:"project"`
	NotificationType NotificationType `json:"notificationType"`
	NotificationName string           `json:"notificationName"`
}

// RemoveNotificationFromProject struct.
type RemoveNotificationFromProject struct {
	Project          string           `json:"project"`
	NotificationType NotificationType `json:"notificationType"`
	NotificationName string           `json:"notificationName"`
}

// UpdateUser struct.
type UpdateUser struct {
	User  User `json:"user"`
	Patch User `json:"patch"`
}

// DeleteUser struct.
type DeleteUser struct {
	User User `json:"user"`
}

// DeleteProject struct.
type DeleteProject struct {
	Project string `json:"project"`
}

// UpdateOpenshift struct.
type UpdateOpenshift struct {
	ID    int       `json:"id"`
	Patch Openshift `json:"patch"`
}

// UpdateNotificationRocketChatPatch struct.
type UpdateNotificationRocketChatPatch struct {
	Name    string `json:"name,omitempty"`
	Webhook string `json:"webhook,omitempty"`
	Channel string `json:"channel,omitempty"`
}

// UpdateNotificationSlackPatch struct.
type UpdateNotificationSlackPatch struct {
	Name    string `json:"name,omitempty"`
	Webhook string `json:"webhook,omitempty"`
	Channel string `json:"channel,omitempty"`
}

// UpdateNotificationSlack struct.
type UpdateNotificationSlack struct {
	ID    int                          `json:"id"`
	Patch UpdateNotificationSlackPatch `json:"patch"`
}

// UpdateNotificationRocketChatInput struct.
type UpdateNotificationRocketChatInput struct {
	ID    int                               `json:"id"`
	Patch UpdateNotificationRocketChatPatch `json:"patch"`
}

// UpdateSSHKeyPatch struct.
type UpdateSSHKeyPatch struct {
	Name     string     `json:"name,omitempty"`
	KeyValue string     `json:"keyValue,omitempty"`
	KeyType  SSHKeyType `json:"keyType,omitempty"`
}

// UpdateSSHKey struct.
type UpdateSSHKey struct {
	ID    int               `json:"id"`
	Patch UpdateSSHKeyPatch `json:"patch"`
}

// UpdateEnvironment struct.
type UpdateEnvironment struct {
	ID    int         `json:"id"`
	Patch Environment `json:"patch"`
}

// AddUpdateEnvironment struct.
type AddUpdateEnvironment struct {
	Name  string      `json:"name"`
	Patch Environment `json:"patch"`
}

// DeleteEnvironment struct.
type DeleteEnvironment struct {
	Name    string `json:"name"`
	Project string `json:"project"`
	Execute bool   `json:"execute"`
}

// EnvVariable struct.
type EnvVariable struct {
	ID     int              `json:"id,omitempty"`
	Type   EnvVariableType  `json:"type,omitempty"`
	TypeID int              `json:"typeId"`
	Name   string           `json:"name"`
	Scope  EnvVariableScope `json:"scope"`
	Value  string           `json:"value"`
}

// DeleteEnvVariable struct.
type DeleteEnvVariable struct {
	ID int `json:"id"`
}

// SetEnvironmentServices struct.
type SetEnvironmentServices struct {
	Environment int      `json:"environment"`
	Services    []string `json:"services"`
}

// UploadFilesForTask struct.
type UploadFilesForTask struct {
	Task int `json:"task"`
	// @TODO
	//	Files []Upload `json:"files"`
}

// DeleteFilesForTask struct.
type DeleteFilesForTask struct {
	ID int `json:"id"`
}

// DeployEnvironmentLatest struct.
type DeployEnvironmentLatest struct {
	Environment Environment `json:"environment"`
}

// DeployEnvironmentBranch struct.
type DeployEnvironmentBranch struct {
	Project    Project `json:"project"`
	BranchName string  `json:"branchName"`
	BranchRef  string  `json:"branchRef,omitempty"`
}

// DeployEnvironmentPullrequest struct.
type DeployEnvironmentPullrequest struct {
	Project        Project `json:"project"`
	Number         int     `json:"number"`
	Title          string  `json:"title"`
	BaseBranchName string  `json:"baseBranchName"`
	BaseBranchRef  string  `json:"baseBranchRef"`
	HeadBranchName string  `json:"headBranchName"`
	HeadBranchRef  string  `json:"headBranchRef"`
}

// DeployEnvironmentPromote struct.
type DeployEnvironmentPromote struct {
	SourceEnvironment      Environment `json:"sourceEnvironment"`
	Project                Project     `json:"project"`
	DestinationEnvironment string      `json:"destinationEnvironment"`
}

// AddGroup struct.
type AddGroup struct {
	Name        string `json:"name"`
	ParentGroup Group  `json:"parentGroup,omitempty"`
}

// AddUserToGroup struct.
type AddUserToGroup struct {
	User  User      `json:"user"`
	Group string    `json:"group"`
	Role  GroupRole `json:"role"`
}

// ProjectToGroup struct.
type ProjectToGroup struct {
	Project string `json:"project"`
	Group   string `json:"group"`
}

// UpdateGroup struct.
type UpdateGroup struct {
	Group Group `json:"group"`
	Patch Group `json:"patch"`
}

// DeleteGroup struct.
type DeleteGroup struct {
	Group Group `json:"group"`
}

// UserGroup struct.
type UserGroup struct {
	User  User  `json:"user"`
	Group Group `json:"group"`
}

// UserGroupRole struct.
type UserGroupRole struct {
	User  User      `json:"user"`
	Group Group     `json:"group"`
	Role  GroupRole `json:"role"`
}

// ProjectGroups struct.
type ProjectGroups struct {
	Project Project `json:"project"`
	Groups  []Group `json:"groups"`
}

// Environment struct.
type Environment struct {
	ID                   int                   `json:"id,omitempty"`
	Name                 string                `json:"name,omitempty"`
	DeployType           DeployType            `json:"deployType,omitempty"`
	DeployTitle          string                `json:"deployTitle,omitempty"`
	DeployBaseRef        string                `json:"deployBaseRef,omitempty"`
	DeployHeadRef        string                `json:"deployHeadRef,omitempty"`
	AutoIdle             *int                  `json:"autoIdle,omitempty"`
	EnvironmentType      EnvType               `json:"environmentType,omitempty"`
	OpenshiftProjectName string                `json:"openshiftProjectName,omitempty"`
	Created              string                `json:"created,omitempty"`
	Deleted              string                `json:"deleted,omitempty"`
	Route                string                `json:"route,omitempty"`
	Routes               string                `json:"routes,omitempty"`
	EnvVariables         []EnvironmentVariable `json:"envVariables,omitempty"`
	Backups              []Backup              `json:"backups,omitempty"`
	Tasks                []Task                `json:"tasks,omitempty"`
	Project              int                   `json:"project,omitempty"`
	AdvancedTasks        []AdvancedTask        `json:"advancedTasks"`
}

// EnvironmentBackups struct.
type EnvironmentBackups struct {
	OpenshiftProjectName string `json:"openshiftProjectName"`
}

// Data .
type Data struct {
	Data  interface{}
	Error interface{}
}

// EnvironmentVariable struct.
type EnvironmentVariable struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Scope string `json:"scope,omitempty"`
	Value string `json:"value,omitempty"`
}

// AdvancedTask task def struct
type AdvancedTask struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description"`
}

// Backup struct.
type Backup struct {
	ID          int     `json:"id,omitempty"`
	Source      string  `json:"source,omitempty"`
	BackupID    string  `json:"backupId,omitempty"`
	Created     string  `json:"created,omitempty"`
	Deleted     string  `json:"deleted,omitempty"`
	Environment int     `json:"environment,omitempty"`
	Restore     Restore `json:"restore,omitempty"`
}

// Restore struct.
type Restore struct {
	ID              int    `json:"id,omitempty"`
	BackupID        string `json:"backupId,omitempty"`
	Status          string `json:"status,omitempty"`
	RestoreLocation string `json:"restoreLocation,omitempty"`
	Created         string `json:"created,omitempty"`
}

// RocketChats struct.
type RocketChats struct {
	RocketChats []NotificationRocketChat `json:"rocketchats"`
}

// Slacks struct.
type Slacks struct {
	Slacks []NotificationSlack `json:"slacks"`
}
