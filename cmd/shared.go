package cmd

var projectBranches string
var projectProductionEnvironment string
var projectDevelopmentEnvironmentsLimit int
var projectPullRequests string
var projectAutoIdle *int
var projectGitURL string
var projectProductionEnv string
var projectOpenshift int
var lagoonHostname string
var lagoonPort string
var lagoonGraphQL string
var lagoonToken string

var jsonPatch string

// ProjectByName struct.
type ProjectByName struct {
	ProjectByName Project `json:"projectByName"`
}

// WhatIsThere struct.
type WhatIsThere struct {
	AllProjects []Project `json:"allProjects"`
}

// Environments struct.
type Environments struct {
	ID                   int            `json:"id,omitempty"`
	Name                 string         `json:"name,omitempty"`
	DeployType           string         `json:"deployType,omitempty"`
	DeployTitle          string         `json:"deployTitle,omitempty"`
	DeployBaseRef        string         `json:"deployBaseRef,omitempty"`
	DeployHeadRef        string         `json:"deployHeadRef,omitempty"`
	AutoIdle             int            `json:"autoIdle,omitempty"`
	EnvironmentType      string         `json:"environmentType,omitempty"`
	OpenshiftProjectName string         `json:"openshiftProjectName,omitempty"`
	Created              string         `json:"created,omitempty"`
	Deleted              string         `json:"deleted,omitempty"`
	Route                string         `json:"route,omitempty"`
	Routes               string         `json:"routes,omitempty"`
	MonitoringUrls       string         `json:"monitoringUrls,omitempty"`
	EnvVariables         []EnvVariables `json:"envVariables,omitempty"`
}

// Project struct.
type Project struct {
	ID                           int            `json:"id,omitempty"`
	Name                         string         `json:"name,omitempty"`
	GitURL                       string         `json:"gitUrl,omitempty"`
	PrivateKey                   string         `json:"privateKey,omitempty"`
	Subfolder                    string         `json:"subfolder,omitempty"`
	ActiveSystemsTask            string         `json:"activeSystemsTask,omitempty"`
	ActiveSystemsDeploy          string         `json:"activeSystemsDeploy,omitempty"`
	ActiveSystemsRemove          string         `json:"activeSystemsRemove,omitempty"`
	ActiveSystemsPromote         string         `json:"activeSystemsPromote,omitempty"`
	Branches                     string         `json:"branches,omitempty"`
	Pullrequests                 string         `json:"pullrequests,omitempty"`
	ProductionEnvironment        string         `json:"productionEnvironment,omitempty"`
	AutoIdle                     int            `json:"autoIdle,omitempty"`
	StorageCalc                  int            `json:"storageCalc,omitempty"`
	OpenshiftProjectPattern      string         `json:"openshiftProjectPattern,omitempty"`
	DevelopmentEnvironmentsLimit int            `json:"developmentEnvironmentsLimit,omitempty"`
	Created                      string         `json:"created,omitempty"`
	Openshift                    Openshift      `json:"openshift,omitempty"`
	EnvVariables                 []EnvVariables `json:"envVariables,omitempty"`
	Environments                 []Environments `json:"environments,omitempty"`
	Deployments                  []Deployments  `json:"deployments,omitempty"`
}

type EnvVariables struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Scope string `json:"scope,omitempty"`
	Value string `json:"value,omitempty"`
}

type Deployments struct {
	ID        int    `json:"id,omitempty"`
	Status    string `json:"status,omitempty"`
	Started   string `json:"started,omitempty"`
	Completed string `json:"completed,omitempty"`
	RemoteID  string `json:"remoteId,omitempty"`
	Name      string `json:"name,omitempty"`
	BuildLog  string `json:"buildLog,omitempty"`
}

type Backup struct {
	ID       int     `json:"id,omitempty"`
	Source   string  `json:"source,omitempty"`
	BackupID string  `json:"backupId,omitempty"`
	Created  string  `json:"created,omitempty"`
	Deleted  string  `json:"deleted,omitempty"`
	Restore  Restore `json:"restore,omitempty"`
}

type Restore struct {
	ID              int    `json:"id,omitempty"`
	BackupID        string `json:"backupId,omitempty"`
	Status          string `json:"status,omitempty"`
	RestoreLocation string `json:"restoreLocation,omitempty"`
	Created         string `json:"created,omitempty"`
}

type Openshift struct {
	ID            int    `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Created       string `json:"created,omitempty"`
	ConsoleURL    string `json:"consoleUrl,omitempty"`
	ProjectUser   string `json:"projectUser,omitempty"`
	RouterPattern string `json:"routerPattern,omitempty"`
	SSHHost       string `json:"sshHost,omitempty"`
	SSHPort       string `json:"sshPort,omitempty"`
	Token         string `json:"token,omitempty"`
}

type Environment struct {
	EnvironmentByOpenshiftProjectName Environments `json:"environmentByOpenshiftProjectName"`
}

type DeployResult struct {
	DeployEnvironmentBranch string `json:"deployEnvironmentBranch"`
}

type DeleteEnvironmentResult struct {
	DeleteEnvironment string `json:"deleteEnvironment"`
}

type DeleteProjectResult struct {
	DeleteProject string `json:"deleteProject"`
}

type AddProject struct {
	AddProject Project `json:"addProject"`
}

type UpdateProject struct {
	UpdateProject Project `json:"updateProject"`
}
