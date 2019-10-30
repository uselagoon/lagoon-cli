package cmd

import (
	"fmt"
	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/graphql"
	"os"
	"strconv"
	"strings"

	"encoding/json"
	"github.com/logrusorgru/aurora"

	"github.com/spf13/cobra"
)

var addVariableEnvCmd = &cobra.Command{
	Use:   "environment [project name] [environment name]",
	Short: "Add variable to environment",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			fmt.Println("Not enough arguments. Requires: project name, environment name.")
			cmd.Help()
			os.Exit(1)
		}
		projectName := args[0]
		environmentName := args[1]

		// setup the envvar
		var envVar api.EnvVariable
		var jsonBytes []byte

		// check if we have a jsonpatch or not
		if jsonPatch != "" {
			// unmarshall the json patch into a tempvar
			var tempEnvVar api.EnvVariable
			err := json.Unmarshal([]byte(jsonPatch), &tempEnvVar)
			if err != nil {
				fmt.Println("{\"error\":\"" + err.Error() + "\"}")
				return
			}
			if tempEnvVar.Name == "" || tempEnvVar.Value == "" || string(tempEnvVar.Scope) == "" {
				fmt.Println("{\"error\":\"Must define a variable name, value and scope\"}")
				return
			}
			envVar.Name = tempEnvVar.Name
			envVar.Value = tempEnvVar.Value
			variableScope = string(tempEnvVar.Scope)
		} else {
			if variableName == "" || variableValue == "" || variableScope == "" {
				fmt.Println("{\"error\":\"Must define a variable name, value and scope\"}")
				return
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
			fmt.Println("{\"error\":\"Unknown scope: " + variableScope + "\"}")
			return
		}

		// set up a lagoonapi client
		lagoonAPI, err := graphql.LagoonAPI()
		if err != nil {
			fmt.Println("{\"error\":\"" + err.Error() + "\"}")
			return
		}

		// get project info from lagoon, we need the project ID for later
		project := api.Project{
			Name: projectName,
		}
		projectByName, err := lagoonAPI.GetProjectByName(project, graphql.ProjectByNameFragment)
		if err != nil {
			fmt.Println("{\"error\":\"" + err.Error() + "\"}")
			return
		}
		jsonBytes, _ = json.Marshal(projectByName)
		reMappedResult := projectByName.(map[string]interface{})
		var projectInfo api.Project
		jsonBytes, _ = json.Marshal(reMappedResult["project"])
		err = json.Unmarshal([]byte(jsonBytes), &projectInfo)
		if err != nil {
			fmt.Println("{\"error\":\"" + err.Error() + "\"}")
			return
		}

		// get the environment info from lagoon, we need the environment ID for later
		// we consume the project ID here
		environment := api.EnvironmentByName{
			Name:    environmentName,
			Project: projectInfo.ID,
		}
		environmentByName, err := lagoonAPI.GetEnvironmentByName(environment)
		if err != nil {
			fmt.Println("{\"error\":\"" + err.Error() + "\"}")
			return
		}
		jsonBytes, _ = json.Marshal(environmentByName)
		reMappedResult = environmentByName.(map[string]interface{})
		var environmentInfo api.Environment
		jsonBytes, _ = json.Marshal(reMappedResult["environmentByName"])
		err = json.Unmarshal([]byte(jsonBytes), &environmentInfo)
		if err != nil {
			fmt.Println("{\"error\":\"" + err.Error() + "\"}")
			return
		}

		// run the query to add the environment variable to lagoon
		customReq := api.CustomRequest{
			Query: `mutation addEnvironmentVariableToProject ($type: EnvVariableType!, $typeId: Int!, $scope: EnvVariableScope!, $name: String!, $value: String!) {
				addEnvVariable(input:{type: $type, typeId: $typeId, scope: $scope, name: $name, value: $value}) {
					id
				}
			}`,
			Variables: map[string]interface{}{
				"type":   api.EnvironmentVar,
				"typeId": environmentInfo.ID,
				"scope":  envVar.Scope,
				"name":   envVar.Name,
				"value":  envVar.Value,
			},
		}
		customReqResult, err := lagoonAPI.Request(customReq)
		if err != nil {
			fmt.Println("{\"error\":\"" + err.Error() + "\"}")
			return
		}
		jsonBytes, _ = json.Marshal(customReqResult)

		// remap the returned data for processing
		reMappedResult = customReqResult.(map[string]interface{})
		var updatedProject api.EnvVariable
		jsonBytes, _ = json.Marshal(reMappedResult["addEnvVariable"])
		err = json.Unmarshal([]byte(jsonBytes), &updatedProject)
		if err != nil {
			fmt.Println("{\"error\":\"" + err.Error() + "\"}")
			return
		}

		fmt.Println(fmt.Sprintf("Result: %s", aurora.Green("success")))
		fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("ID"), strconv.Itoa(updatedProject.ID)))
	},
}

func init() {
	addVariableCmd.AddCommand(addVariableEnvCmd)
	addVariableEnvCmd.Flags().StringVarP(&variableName, "varname", "N", "", "Name of the variable to add")
	addVariableEnvCmd.Flags().StringVarP(&variableValue, "varvalue", "V", "", "Value of the variable to add")
	addVariableEnvCmd.Flags().StringVarP(&variableScope, "varscope", "S", "", "Scope of the variable[global, build, runtime]")
	addVariableEnvCmd.Flags().StringVarP(&jsonPatch, "json", "j", "", "JSON string to patch")
}
