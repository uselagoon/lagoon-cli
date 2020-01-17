package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/amazeeio/lagoon-cli/output"
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

// import vars
var importFile string

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
}

// read stdin or file, if the input is stdin it will try read from it, if a path is provided it will read that instead
// returns the
func readStdInOrFile(filePath string) (string, error) {
	// var err error
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// check if we are getting data froms stdin
		scanner := bufio.NewScanner(os.Stdin)
		taskCommand = ""
		// read the buff into a string
		for scanner.Scan() {
			taskCommand = taskCommand + scanner.Text() + "\n"
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
			// return "", err
		}
		// return string(taskCommand), err
	} else {
		// otherwise we can read from a file
		if filePath != "" {
			taskCommandBytes, err := ioutil.ReadFile(filePath) // just pass the file name
			if err != nil {
				return "", err
			}
			taskCommand = string(taskCommandBytes)
		} else {
			return "", errors.New("no path provided")
		}
	}
	return string(taskCommand), nil
}

var debugEnable bool

var noDataError = "no data returned from the lagoon api"

func handleError(err error) {
	if err != nil {
		output.RenderError(err.Error(), outputOptions)
		os.Exit(1)
	}
}
