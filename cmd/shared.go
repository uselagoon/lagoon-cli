package cmd

import (
	"os"
	"regexp"
	"strings"

	"github.com/amazeeio/lagoon-cli/pkg/output"
)

// config vars
var lagoonHostname string
var lagoonPort string
var lagoonGraphQL string
var lagoonToken string
var lagoonUI string
var lagoonKibana string

// variable vars
var variableValue string
var variableName string
var variableScope string

// user vars
var userFirstName string
var userLastName string
var userEmail string
var pubKeyFile string
var nameInPubKeyFile bool
var sshKeyName string

// group vars
var groupName string
var groupRole string

var jsonPatch string
var revealValue bool
var listAllProjects bool
var noHeader bool

// These are available to all cmds and are set either by flags (-p and -e) or via `lagoon-cli/app` when entering a directory that has a valid lagoon project
var cmdProjectName string
var cmdProjectEnvironment string

var remoteID string

var notificationName string
var notificationNewName string
var notificationWebhook string
var notificationChannel string

var deployBranchName string

var outputOptions = output.Options{
	Header: false,
	CSV:    false,
	JSON:   false,
	Pretty: false,
	Debug:  false,
}

var debugEnable bool

var noDataError = "no data returned from the lagoon api"

func handleError(err error) {
	if err != nil {
		output.RenderError(err.Error(), outputOptions)
		os.Exit(1)
	}
}

func returnNonEmptyString(value string) string {
	if len(value) == 0 {
		return "-"
	}
	return value
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func stripNewLines(stripString string) string {
	return strings.TrimSuffix(stripString, "\n")
}

// regex the name to make it match what lagoon is expecting
func sanitizeName(name string) string {
	var re = regexp.MustCompile(`[^a-zA-Z0-9-]`)
	sanitizedName := re.ReplaceAllString(name, `$1-$2`)
	return sanitizedName
}
