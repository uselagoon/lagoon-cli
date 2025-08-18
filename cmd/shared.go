package cmd

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/uselagoon/lagoon-cli/pkg/output"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	"github.com/uselagoon/machinery/api/schema"
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

func nullBoolToUint(b bool) *uint {
	t := uint(0)
	if b {
		t = uint(1)
	}
	return &t
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

// Outputs the message in a way that can be captured by testing
func handleNilResults(message string, cmd *cobra.Command, fields ...interface{}) error {
	outputOptions.Error = fmt.Sprintf(message, fields...)
	r := output.RenderOutput(output.Table{Data: []output.Data{[]string{}}}, outputOptions)
	fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
	return nil
}

func quotaCheck(quota int) string {
	quotaRoute := strconv.Itoa(quota)
	if quota < 0 {
		quotaRoute = "∞"
	}
	return quotaRoute
}

func environmentCheck(project string, environment string, debug bool) error {

	current := lagoonCLIConfig.Current
	token := lagoonCLIConfig.Lagoons[current].Token
	lc := lclient.New(
		lagoonCLIConfig.Lagoons[current].GraphQL,
		lagoonCLIVersion,
		lagoonCLIConfig.Lagoons[current].Version,
		&token,
		debug)
	
	var environments []schema.Environment	
	err := lc.EnvironmentsByProjectName(context.TODO(), project, &environments)
	if err != nil {
		return err
	}

	names := make([]string, len(environments))
	for i, e := range environments {
		names[i] = e.AddEnvironmentInput.Name
		if environment == names[i] {
			return nil
		}
	}
	
	return fmt.Errorf("invalid environment for project %s: %s. Valid options are %s.\n", project, environment, strings.Join(names, ", "))
}
