package cmd

import (
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/graphql"
	// "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var addVariableProjectCmd = &cobra.Command{
	Use:   "project [project name]",
	Short: "Add variable to a project",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Println("Not enough arguments. Requires: project name.")
			cmd.Help()
			os.Exit(1)
		}
		projectName := args[0]

		fmt.Println(fmt.Sprintf("Deleting %s-%s", projectName))

		if yesNo() {
			var responseData interface{}
			err := graphql.GraphQLRequest(fmt.Sprintf(`mutation addEnvironmentVariableToEnvironment{
  addEnvVariable(input:{type:ENVIRONMENT, typeId:10026, scope:GLOBAL, name:"VARIABLE_NAME", value:"variablevalue"}) {
    id
  }
}`, projectName), &responseData)
			//var err error
			err = nil
			fmt.Println(fmt.Sprintf("%s", projectName))
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
	addVariableCmd.AddCommand(addVariableProjectCmd)
}
