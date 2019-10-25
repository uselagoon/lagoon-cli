package cmd

import (
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/graphql"
	// "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var projectVariableEnvCmd = &cobra.Command{
	Use:   "environment [add|delete] [project name] [environment name]",
	Short: "Add or delete variable on environment",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 3 {
			fmt.Println("Not enough arguments. Requires: command, project name and environment.")
			cmd.Help()
			os.Exit(1)
		}
		cmdValue := args[0]
		if cmdValue != "delete" && cmdValue != "add" {
			fmt.Println("Command must be add or delete")
			cmd.Help()
			os.Exit(1)
		}
		projectName := args[1]
		projectEnvironment := args[2]

		fmt.Println(fmt.Sprintf("Deleting %s-%s", projectName, projectEnvironment))

		if yesNo() {
			var responseData interface{}
			err := graphql.GraphQLRequest(fmt.Sprintf(`mutation addEnvironmentVariableToEnvironment{
  addEnvVariable(input:{type:ENVIRONMENT, typeId:10026, scope:GLOBAL, name:"VARIABLE_NAME", value:"variablevalue"}) {
    id
  }
}`, projectName, projectEnvironment), &responseData)
			//var err error
			err = nil
			fmt.Println(fmt.Sprintf("%s %s-%s", cmdValue, projectName, projectEnvironment))
			if err != nil {
				fmt.Println(err.Error())
			} else {
				// if responseData.DeleteEnvironment == "success" {
				// 	fmt.Println(fmt.Sprintf("Result: %s", aurora.Green(responseData.DeleteEnvironment)))
				// } else {
				// 	fmt.Println(fmt.Sprintf("Result: %s", aurora.Yellow(responseData.DeleteEnvironment)))
				// }
			}
		}

	},
}

func init() {
	projectVariableCmd.AddCommand(projectVariableEnvCmd)
}
