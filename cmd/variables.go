package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/lagoon/environments"
	"github.com/amazeeio/lagoon-cli/lagoon/projects"
	"github.com/amazeeio/lagoon-cli/output"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// EnvironmentVariableFlags .
type EnvironmentVariableFlags struct {
	VarName     string `json:"varname,omitempty"`
	VarValue    string `json:"varvalue,omitempty"`
	VarScope    string `json:"varscope,omitempty"`
	Project     string `json:"project,omitempty"`
	Environment string `json:"environment,omitempty"`
}

func parseEnvVars(flags pflag.FlagSet) EnvironmentVariableFlags {
	configMap := make(map[string]interface{})
	flags.VisitAll(func(f *pflag.Flag) {
		if flags.Changed(f.Name) {
			configMap[f.Name] = f.Value
		}
	})
	jsonStr, _ := json.Marshal(configMap)
	parsedFlags := EnvironmentVariableFlags{}
	json.Unmarshal(jsonStr, &parsedFlags)
	return parsedFlags
}

var addVariableCmd = &cobra.Command{
	Use:   "variable",
	Short: "Add variables on environments or projects",
	Run: func(cmd *cobra.Command, args []string) {
		envVarFlags := parseEnvVars(*cmd.Flags())
		if envVarFlags.Project == "" {
			fmt.Println("Not enough arguments. Requires: project name")
			cmd.Help()
			os.Exit(1)
		}

		// setup the envvar
		var envVar api.EnvVariable
		// check if we have a jsonpatch or not
		if jsonPatch != "" {
			// unmarshall the json patch into a tempvar
			var tempEnvVar api.EnvVariable
			err := json.Unmarshal([]byte(jsonPatch), &tempEnvVar)
			if err != nil {
				output.RenderError(err.Error(), outputOptions)
				os.Exit(1)
			}
			if tempEnvVar.Name == "" || tempEnvVar.Value == "" || string(tempEnvVar.Scope) == "" {
				output.RenderError("Must define a variable name, value and scope", outputOptions)
				os.Exit(1)
			}
			envVar.Name = tempEnvVar.Name
			envVar.Value = tempEnvVar.Value
			variableScope = string(tempEnvVar.Scope)
		} else {
			if envVarFlags.VarName == "" || envVarFlags.VarValue == "" || envVarFlags.VarScope == "" {
				output.RenderError("Must define a variable name, value and scope", outputOptions)
				os.Exit(1)
			}
			envVar.Name = envVarFlags.VarName
			envVar.Value = envVarFlags.VarValue
		}
		if strings.EqualFold(string(envVarFlags.VarScope), "global") {
			envVar.Scope = api.GlobalVar
		} else if strings.EqualFold(string(envVarFlags.VarScope), "build") {
			envVar.Scope = api.BuildVar
		} else if strings.EqualFold(string(envVarFlags.VarScope), "runtime") {
			envVar.Scope = api.RuntimeVar
		} else {
			output.RenderError("Unknown scope: "+envVarFlags.VarScope, outputOptions)
			os.Exit(1)
		}
		var customReqResult []byte
		var err error
		returnResultData := map[string]interface{}{}
		if envVarFlags.Environment != "" {
			customReqResult, err = environments.AddEnvironmentVariableToEnvironment(envVarFlags.Project, envVarFlags.Environment, envVar)
			returnResultData["Project"] = envVarFlags.Project
			returnResultData["Environment"] = envVarFlags.Environment
		} else {
			customReqResult, err = projects.AddEnvironmentVariableToProject(envVarFlags.Project, envVar)
			returnResultData["Project"] = envVarFlags.Project
		}
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		var updatedVariable api.EnvVariable
		err = json.Unmarshal([]byte(customReqResult), &updatedVariable)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
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
	Use:   "variable",
	Short: "Delete a variable from an environment",
	Long:  `This allows you to delete an environment variable from a project.`,
	Run: func(cmd *cobra.Command, args []string) {
		envVarFlags := parseEnvVars(*cmd.Flags())
		if envVarFlags.Project == "" {
			fmt.Println("Not enough arguments. Requires: project name")
			cmd.Help()
			os.Exit(1)
		}

		// setup the envvar
		var envVar api.EnvVariable
		// check if we have a jsonpatch or not
		if jsonPatch != "" {
			// unmarshall the json patch into a tempvar
			var tempEnvVar api.EnvVariable
			err := json.Unmarshal([]byte(jsonPatch), &tempEnvVar)
			if err != nil {
				output.RenderError(err.Error(), outputOptions)
				os.Exit(1)
			}
			if tempEnvVar.Name == "" {
				output.RenderError("Must define a variable name", outputOptions)
				os.Exit(1)
			}
			envVar.Name = tempEnvVar.Name
		} else {
			if envVarFlags.VarName == "" {
				output.RenderError("Must define a variable name", outputOptions)
				os.Exit(1)
			}
			envVar.Name = envVarFlags.VarName
		}

		if yesNo() {
			var deleteResult []byte
			var err error
			if envVarFlags.Environment != "" {
				deleteResult, err = environments.DeleteEnvironmentVariableFromEnvironment(envVarFlags.Project, envVarFlags.Environment, envVar)
			} else {
				deleteResult, err = projects.DeleteEnvironmentVariableFromProject(envVarFlags.Project, envVar)
			}
			if err != nil {
				output.RenderError(err.Error(), outputOptions)
				os.Exit(1)
			}
			resultData := output.Result{
				Result: string(deleteResult),
			}
			output.RenderResult(resultData, outputOptions)
		}
	},
}

func init() {
	addVariableCmd.Flags().StringVarP(&variableName, "varname", "N", "", "Name of the variable to add")
	addVariableCmd.Flags().StringVarP(&variableValue, "varvalue", "V", "", "Value of the variable to add")
	addVariableCmd.Flags().StringVarP(&variableScope, "varscope", "S", "", "Scope of the variable[global, build, runtime]")
	addVariableCmd.Flags().StringVarP(&jsonPatch, "json", "j", "", "JSON string to patch")
	deleteVariableCmd.Flags().StringVarP(&variableName, "varname", "N", "", "Name of the variable to add")
}
