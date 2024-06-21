package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"os"
	"strings"

	"github.com/uselagoon/lagoon-cli/pkg/output"
)

// config vars
var lagoonHostname string
var lagoonPort string
var lagoonGraphQL string
var lagoonToken string
var lagoonUI string
var lagoonKibana string
var lagoonSSHKey string

// group vars
var groupName string

var jsonPatch string
var revealValue bool
var listAllProjects bool

// These are available to all cmds and are set either by flags (-p and -e) or via `lagoon-cli/app` when entering a directory that has a valid lagoon project
var cmdProjectName string
var cmdProjectEnvironment string

var outputOptions = output.Options{
	Header: false,
	CSV:    false,
	JSON:   false,
	Pretty: false,
	Debug:  false,
}

var groupRoles = []string{"guest", "reporter", "developer", "maintainer", "owner"}

var debugEnable bool

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

func nullStrCheck(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func nullUintCheck(i uint) *uint {
	if i == 0 {
		return nil
	}
	return &i
}

func nullIntCheck(i int) *int {
	if i == 0 {
		return nil
	}
	return &i
}

// Specify the fields and values to check for required input e.g. requiredInputCheck("field1", value1, "field2", value2)
func requiredInputCheck(fieldsAndValues ...string) error {
	for i := 0; i < len(fieldsAndValues); i += 2 {
		field := fieldsAndValues[i]
		value := fieldsAndValues[i+1]

		if value == "" || value == "0" {
			return fmt.Errorf("missing argument: %s is not defined", field)
		}
	}
	return nil
}

// SetUpRootCmdFlags sets up the flags for the root command
func SetUpRootCmdFlags() {
	rootCmd.Flags().StringP("config-file", "", "", "Path to the config file to use (must be *.yml or *.yaml)")
	rootCmd.Flags().StringVarP(&cmdLagoon, "lagoon", "l", "", "The Lagoon instance to interact with")
}

// AddGenericFlags adds the generic flags to the command being executed. --debug, --output-json, --project, --environment, --force
// Instantiates an explicit flagset for each command to avoid 'flag redefined' errors on multiple tests containing the same command
func AddGenericFlags(cmd *cobra.Command) {
	flags := pflag.FlagSet{}
	flags.BoolVarP(&debugEnable, "debug", "", false, "Enable debugging output (if supported)")
	flags.BoolVarP(&outputOptions.JSON, "output-json", "", false, "Output as JSON (if supported)")
	flags.StringVarP(&cmdProjectName, "project", "p", "", "Specify a project to use")
	flags.StringVarP(&cmdProjectEnvironment, "environment", "e", "", "Specify an environment to use")
	flags.BoolVarP(&forceAction, "force", "", false, "Force yes on prompts (if supported)")
	cmd.Flags().AddFlagSet(&flags)
}
