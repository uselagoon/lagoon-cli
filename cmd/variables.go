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
)

var addVariableEnvCmd = &cobra.Command{
	Use:   "environment [project name] [environment name]",
	Short: "Add variable to environment",
	Run: func(cmd *cobra.Command, args []string) {
		var projectName string
		var environmentName string
		if len(args) < 2 {
			if cmdProject.Name != "" && len(args) == 1 {
				projectName = cmdProject.Name
				environmentName = args[0]
			} else {
				fmt.Println("Not enough arguments. Requires: project name and environment name")
				cmd.Help()
				os.Exit(1)
			}
		} else {
			projectName = args[0]
			environmentName = args[1]
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
			if variableName == "" || variableValue == "" || variableScope == "" {
				output.RenderError("Must define a variable name, value and scope", outputOptions)
				os.Exit(1)
			}
			envVar.Name = variableName
			envVar.Value = variableValue
		}
		if strings.EqualFold(string(variableScope), "global") {
			envVar.Scope = api.GlobalVar
		} else if strings.EqualFold(string(variableScope), "build") {
			envVar.Scope = api.BuildVar
		} else if strings.EqualFold(string(variableScope), "runtime") {
			envVar.Scope = api.RuntimeVar
		} else {
			output.RenderError("Unknown scope: "+variableScope, outputOptions)
			os.Exit(1)
		}

		customReqResult, err := environments.AddEnvironmentVariableToEnvironment(projectName, environmentName, envVar)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		var updatedProject api.EnvVariable
		err = json.Unmarshal([]byte(customReqResult), &updatedProject)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"ID": strconv.Itoa(updatedProject.ID),
			},
		}
		output.RenderResult(resultData, outputOptions)
	},
}

var deleteVariableEnvCmd = &cobra.Command{
	Use:   "environment [project name] [environment name]",
	Short: "Delete a variable from an environment",
	Long:  `This allows you to delete an environment variable from a project.`,
	Run: func(cmd *cobra.Command, args []string) {
		var projectName string
		var environmentName string
		if len(args) < 2 {
			if cmdProject.Name != "" && len(args) == 1 {
				projectName = cmdProject.Name
				environmentName = args[0]
			} else {
				fmt.Println("Not enough arguments. Requires: project name and environment name")
				cmd.Help()
				os.Exit(1)
			}
		} else {
			projectName = args[0]
			environmentName = args[1]
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
			if variableName == "" {
				output.RenderError("Must define a variable name", outputOptions)
				os.Exit(1)
			}
			envVar.Name = variableName
		}

		if yesNo() {
			deleteResult, err := environments.DeleteEnvironmentVariableFromEnvironment(projectName, environmentName, envVar)
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
var addVariableProjectCmd = &cobra.Command{
	Use:   "project [project name]",
	Short: "Add variable to a project",
	Long: `This allows you to add an environment variable to a project.

This can be done via flags.
    $ lagoon add variable project my-project --varname VARNAME --varvalue varvalue --scope build
    $ lagoon add variable project my-project -N VARNAME -V varvalue -S build

Or via JSON
    $ lagoon add variable project my-project --json '{"name":"VARNAME", "value":"varvalue", "scope":"build"}'
    $ lagoon add variable project my-project -j '{"name":"VARNAME", "value":"varvalue", "scope":"build"}'
`,
	Run: func(cmd *cobra.Command, args []string) {
		var projectName string
		if len(args) < 1 {
			if cmdProject.Name != "" {
				projectName = cmdProject.Name
			} else {
				fmt.Println("Not enough arguments. Requires: project name")
				cmd.Help()
				os.Exit(1)
			}
		} else {
			projectName = args[0]
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
			if variableName == "" || variableValue == "" || variableScope == "" {
				output.RenderError("Must define a variable name, value and scope", outputOptions)
				os.Exit(1)
			}
			envVar.Name = variableName
			envVar.Value = variableValue
		}
		if strings.EqualFold(string(variableScope), "global") {
			envVar.Scope = api.GlobalVar
		} else if strings.EqualFold(string(variableScope), "build") {
			envVar.Scope = api.BuildVar
		} else if strings.EqualFold(string(variableScope), "runtime") {
			envVar.Scope = api.RuntimeVar
		} else {
			fmt.Println("Unknown scope:", variableScope)
			os.Exit(1)
		}
		customReqResult, err := projects.AddEnvironmentVariableToProject(projectName, envVar)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		var updatedProject api.EnvVariable
		err = json.Unmarshal([]byte(customReqResult), &updatedProject)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"ID": strconv.Itoa(updatedProject.ID),
			},
		}
		output.RenderResult(resultData, outputOptions)
	},
}

var deleteVariableProjectCmd = &cobra.Command{
	Use:   "project [project name]",
	Short: "Delete a variable from a project",
	Long: `This allows you to delete an environment variable from a project.
`,
	Run: func(cmd *cobra.Command, args []string) {
		var projectName string
		if len(args) < 1 {
			if cmdProject.Name != "" {
				projectName = cmdProject.Name
			} else {
				fmt.Println("Not enough arguments. Requires: project name")
				cmd.Help()
				os.Exit(1)
			}
		} else {
			projectName = args[0]
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
			if variableName == "" {
				output.RenderError("Must define a variable name", outputOptions)
				os.Exit(1)
			}
			envVar.Name = variableName
		}

		if yesNo() {
			deleteResult, err := projects.DeleteEnvironmentVariableFromProject(projectName, envVar)
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
	addVariableCmd.AddCommand(addVariableProjectCmd)
	addVariableProjectCmd.Flags().StringVarP(&variableName, "varname", "N", "", "Name of the variable to add")
	addVariableProjectCmd.Flags().StringVarP(&variableValue, "varvalue", "V", "", "Value of the variable to add")
	addVariableProjectCmd.Flags().StringVarP(&variableScope, "varscope", "S", "", "Scope of the variable[global, build, runtime]")
	addVariableProjectCmd.Flags().StringVarP(&jsonPatch, "json", "j", "", "JSON string to patch")
	deleteVariableCmd.AddCommand(deleteVariableProjectCmd)
	deleteVariableProjectCmd.Flags().StringVarP(&variableName, "varname", "N", "", "Name of the variable to add")

	addVariableCmd.AddCommand(addVariableEnvCmd)
	addVariableEnvCmd.Flags().StringVarP(&variableName, "varname", "N", "", "Name of the variable to add")
	addVariableEnvCmd.Flags().StringVarP(&variableValue, "varvalue", "V", "", "Value of the variable to add")
	addVariableEnvCmd.Flags().StringVarP(&variableScope, "varscope", "S", "", "Scope of the variable[global, build, runtime]")
	addVariableEnvCmd.Flags().StringVarP(&jsonPatch, "json", "j", "", "JSON string to patch")
	deleteVariableCmd.AddCommand(deleteVariableEnvCmd)
	deleteVariableEnvCmd.Flags().StringVarP(&variableName, "varname", "N", "", "Name of the variable to add")
}
