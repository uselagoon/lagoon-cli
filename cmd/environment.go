package cmd

import (
	"fmt"
	"os"

	"github.com/uselagoon/lagoon-cli/pkg/output"
	"github.com/spf13/cobra"
)

// @TODO re-enable this at some point if more environment based commands are made availab;e
// EnvironmentFlags .
// type EnvironmentFlags struct {
// 	Name string `json:"name,omitempty"`
// }

// func parseEnvironmentFlags(flags pflag.FlagSet) EnvironmentFlags {
// 	configMap := make(map[string]interface{})
// 	flags.VisitAll(func(f *pflag.Flag) {
// 		if flags.Changed(f.Name) {
// 			configMap[f.Name] = f.Value
// 		}
// 	})
// 	jsonStr, _ := json.Marshal(configMap)
// 	parsedFlags := EnvironmentFlags{}
// 	json.Unmarshal(jsonStr, &parsedFlags)
// 	return parsedFlags
// }

var deleteEnvCmd = &cobra.Command{
	Use:     "environment",
	Aliases: []string{"e"},
	Short:   "Delete an environment",
	Run: func(cmd *cobra.Command, args []string) {
		// environmentFlags := parseEnvironmentFlags(*cmd.Flags()) //@TODO re-enable this at some point if more environment based commands are made availab;e
		if cmdProjectName == "" || cmdProjectEnvironment == "" {
			fmt.Println("Missing arguments: Project name or environment name is not defined")
			cmd.Help()
			os.Exit(1)
		}
		if yesNo(fmt.Sprintf("You are attempting to delete environment '%s' from project '%s', are you sure?", cmdProjectEnvironment, cmdProjectName)) {
			projectByName, err := eClient.DeleteEnvironment(cmdProjectName, cmdProjectEnvironment)
			handleError(err)
			resultData := output.Result{
				Result: string(projectByName),
			}
			output.RenderResult(resultData, outputOptions)
		}
	},
}
