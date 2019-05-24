package cmd

// ProjectByName struct.
type ProjectByName struct {
	ProjectByName Project `json:"projectByName"`
}

// WhatIsThere struct.
type WhatIsThere struct {
	AllProjects []Project `json:"allProjects"`
}

type EnvironmentByOpenshiftProjectName struct {
	Environment Environment `json:"environmentByOpenshiftProjectName"`
}

type DeploymentByRemoteId struct {
	Deployment Deployment `json:"deploymentByRemoteId"`
}

// Environment struct.
type Environment struct {
	Name                 string             `json:"name"`
	EnvironmentType      string             `json:"environmentType"`
	DeployType           string             `json:"deployType"`
	Route                string             `json:"route"`
	Routes               string             `json:"routes"`
	OpenshiftProjectName string             `json:"openshiftProjectName"`
	HitsMonth            EnvironmentHits    `json:"hitsMonth"`
	StorageMonth         EnvironmentStorage `json:"storageMonth"`
	Deployments          []Deployment       `json:"deployments"`
}

// Project struct.
type Project struct {
	ID                           int           `json:"id"`
	GitURL                       string        `json:"gitUrl"`
	Subfolder                    string        `json:"subfolder"`
	Name                         string        `json:"name"`
	Branches                     string        `json:"branches"`
	Pullrequests                 string        `json:"pullrequests"`
	ProductionEnvironment        string        `json:"productionEnvironment"`
	Environments                 []Environment `json:"environments"`
	AutoIdle                     int           `json:"autoIdle"`
	DevelopmentEnvironmentsLimit int           `json:"developmentEnvironmentsLimit"`
	Customer                     Customer      `json:"customer"`
}

// Customer struct.
type Customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type EnvironmentHits struct {
	Total int `json:"total"`
}
type EnvironmentStorage struct {
	BytesUsed int `json:"bytesUsed"`
}

type Deployment struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	Created   string `json:"created"`
	Started   string `json:"started"`
	Completed string `json:"completed"`
	RemoteID  string `json:"remoteId"`
	BuildLog  string `json:"buildLog"`
}
