package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/amazeeio/lagoon-cli/pkg/api"
	"github.com/amazeeio/lagoon-cli/pkg/output"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// EnvironmentVariableFlags .
type EnvironmentVariableFlags struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
	Scope string `json:"scope,omitempty"`
}

func parseEnvVars(flags pflag.FlagSet) api.EnvVariable {
	configMap := make(map[string]interface{})
	flags.VisitAll(func(f *pflag.Flag) {
		if flags.Changed(f.Name) {
			configMap[f.Name] = f.Value
		}
	})
	jsonStr, _ := json.Marshal(configMap)
	parsedFlags := api.EnvVariable{}
	json.Unmarshal(jsonStr, &parsedFlags)
	return parsedFlags
}

var addVariableCmd = &cobra.Command{
	Use:     "variable",
	Aliases: []string{"v"},
	Short:   "Add a variable to an environment or project",
	Run: func(cmd *cobra.Command, args []string) {
		envVarFlags := parseEnvVars(*cmd.Flags())
		if cmdProjectName == "" {
			fmt.Println("Missing arguments: Project name is not defined")
			cmd.Help()
			os.Exit(1)
		}
		if jsonPatch != "" {
			err := json.Unmarshal([]byte(jsonPatch), &envVarFlags)
			handleError(err)
		}
		if envVarFlags.Name == "" || envVarFlags.Value == "" || envVarFlags.Scope == "" {
			output.RenderError("Must define a variable name, value and scope", outputOptions)
			os.Exit(1)
		}
		if strings.EqualFold(string(envVarFlags.Scope), "global") {
			envVarFlags.Scope = api.GlobalVar
		} else if strings.EqualFold(string(envVarFlags.Scope), "build") {
			envVarFlags.Scope = api.BuildVar
		} else if strings.EqualFold(string(envVarFlags.Scope), "runtime") {
			envVarFlags.Scope = api.RuntimeVar
		} else if strings.EqualFold(string(envVarFlags.Scope), "container_registry") {
			envVarFlags.Scope = api.ContainerRegistryVar
		} else if strings.EqualFold(string(envVarFlags.Scope), "internal_container_registry") {
			envVarFlags.Scope = api.InternalContainerRegistryVar
		} else {
			output.RenderError("Unknown scope: "+string(envVarFlags.Scope), outputOptions)
			os.Exit(1)
		}
		var customReqResult []byte
		var err error
		returnResultData := map[string]interface{}{}
		if cmdProjectEnvironment != "" {
			customReqResult, err = eClient.AddEnvironmentVariableToEnvironment(cmdProjectName, cmdProjectEnvironment, envVarFlags)
			returnResultData["Project"] = cmdProjectName
			returnResultData["Environment"] = cmdProjectEnvironment
		} else {
			customReqResult, err = pClient.AddEnvironmentVariableToProject(cmdProjectName, envVarFlags)
			returnResultData["Project"] = cmdProjectName
		}
		handleError(err)
		var updatedVariable api.EnvVariable
		err = json.Unmarshal([]byte(customReqResult), &updatedVariable)
		handleError(err)
		returnResultData["ID"] = strconv.Itoa(updatedVariable.ID)
		resultData := output.Result{
			Result:     "success",
			ResultData: returnResultData,
		}
		output.RenderResult(resultData, outputOptions)
	},
}

// var deleteVariableEnvCmd = &cobra.Command{
var deleteVariableCmd = &cobra.Command{
	Use:     "variable",
	Aliases: []string{"v"},
	Short:   "Delete a variable from an environment",
	Long:    `This allows you to delete an environment variable from a project.`,
	Run: func(cmd *cobra.Command, args []string) {
		envVarFlags := parseEnvVars(*cmd.Flags())
		if cmdProjectName == "" {
			fmt.Println("Missing arguments: Project name is not defined")
			cmd.Help()
			os.Exit(1)
		}
		if jsonPatch != "" {
			err := json.Unmarshal([]byte(jsonPatch), &envVarFlags)
			handleError(err)
		}
		if envVarFlags.Name == "" {
			output.RenderError("Must define a variable name", outputOptions)
			os.Exit(1)
		}
		deleteMsg := fmt.Sprintf("You are attempting to delete variable '%s' from project '%s', are you sure?", envVarFlags.Name, cmdProjectName)
		if cmdProjectEnvironment != "" {
			deleteMsg = fmt.Sprintf("You are attempting to delete variable '%s' from environment '%s' in project '%s', are you sure?", envVarFlags.Name, cmdProjectEnvironment, cmdProjectName)
		}
		if yesNo(deleteMsg) {
			var deleteResult []byte
			var err error
			if cmdProjectEnvironment != "" {
				deleteResult, err = eClient.DeleteEnvironmentVariableFromEnvironment(cmdProjectName, cmdProjectEnvironment, envVarFlags)
			} else {
				deleteResult, err = pClient.DeleteEnvironmentVariableFromProject(cmdProjectName, envVarFlags)
			}
			handleError(err)
			resultData := output.Result{
				Result: string(deleteResult),
			}
			output.RenderResult(resultData, outputOptions)
		}
	},
}

func init() {
	addVariableCmd.Flags().StringVarP(&variableName, "name", "N", "", "Name of the variable to add")
	addVariableCmd.Flags().StringVarP(&variableValue, "value", "V", "", "Value of the variable to add")
	addVariableCmd.Flags().StringVarP(&variableScope, "scope", "S", "", "Scope of the variable[global, build, runtime, container_registry, internal_container_registry]")
	addVariableCmd.Flags().StringVarP(&jsonPatch, "json", "j", "", "JSON string to patch")
	deleteVariableCmd.Flags().StringVarP(&variableName, "name", "N", "", "Name of the variable to delete")
}
