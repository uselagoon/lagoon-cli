package cmd

import (
	"github.com/amazeeio/lagoon-cli/output"
)

var projectBranches string
var projectProductionEnvironment string
var projectDevelopmentEnvironmentsLimit int
var projectPullRequests string
var projectAutoIdle *int
var projectGitURL string
var projectProductionEnv string
var projectOpenshift int

// config vars
var lagoonHostname string
var lagoonPort string
var lagoonGraphQL string
var lagoonToken string

// variable vars
var variableValue string
var variableName string
var variableScope string

var jsonPatch string
var revealValue bool
var listAllProjects bool
var noHeader bool

var outputOptions = output.Options{
	Header: false,
	CSV:    false,
	JSON:   false,
	Pretty: false,
}

// LagoonConfig .
type LagoonConfig struct {
	Hostname string
	Port     string
}
