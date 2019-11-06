package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/lagoon/environments"
	"github.com/amazeeio/lagoon-cli/output"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// EnvironmentFlags .
type EnvironmentFlags struct {
	Project     string `json:"project,omitempty"`
	Environment string `json:"environment,omitempty"`
}

func parseEnvironmentFlags(flags pflag.FlagSet) EnvironmentFlags {
	configMap := make(map[string]interface{})
	flags.VisitAll(func(f *pflag.Flag) {
		if flags.Changed(f.Name) {
			configMap[f.Name] = f.Value
		}
	})
	jsonStr, _ := json.Marshal(configMap)
	parsedFlags := EnvironmentFlags{}
	json.Unmarshal(jsonStr, &parsedFlags)
	return parsedFlags
}

var deleteEnvCmd = &cobra.Command{
	Use:   "environment [project name] [environment name]",
	Short: "Delete an environment",
	Run: func(cmd *cobra.Command, args []string) {
		environmentFlags := parseEnvironmentFlags(*cmd.Flags())
		if environmentFlags.Project == "" || environmentFlags.Environment == "" {
			fmt.Println("Not enough arguments. Requires: project name and environment name")
			cmd.Help()
			os.Exit(1)
		}
		fmt.Println(fmt.Sprintf("Deleting %s-%s", environmentFlags.Project, environmentFlags.Environment))

		if yesNo() {
			projectByName, err := environments.DeleteEnvironment(environmentFlags.Project, environmentFlags.Environment)
			if err != nil {
				output.RenderError(err.Error(), outputOptions)
				os.Exit(1)
			}
			resultData := output.Result{
				Result: string(projectByName),
			}
			output.RenderResult(resultData, outputOptions)
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteEnvCmd)
}
