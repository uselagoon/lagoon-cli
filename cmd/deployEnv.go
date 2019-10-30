package cmd

import (
	"fmt"
	"os"

	"encoding/json"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/graphql"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var deployEnvCmd = &cobra.Command{
	Use:   "deploy [project name] [environment name]",
	Short: "Deploy an environment",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			fmt.Println("Not enough arguments. Requires: project name and environment.")
			cmd.Help()
			os.Exit(1)
		}
		projectName := args[0]
		projectEnvironment := args[1]

		lagoonAPI, err := graphql.LagoonAPI()
		if err != nil {
			fmt.Println(err)
			return
		}

		var jsonBytes []byte

		fmt.Println(fmt.Sprintf("Deploying %s-%s", projectName, projectEnvironment))

		if yesNo() {

			CustomRequest := api.CustomRequest{
				Query: `mutation {
					deployEnvironmentBranch(
					  input: {
						project:{name: $project}
						branchName: $branch
					  }
					)
				  }`,
				Variables: map[string]interface{}{
					"project": projectName,
					"branch":  projectEnvironment,
				},
			}
			CustomRequestResult, err := lagoonAPI.Request(CustomRequest)
			if err != nil {
				fmt.Println(err)
			}
			jsonBytes, _ = json.Marshal(CustomRequestResult)
			fmt.Println(string(jsonBytes))

			if string(jsonBytes) == "success" {
				fmt.Println(fmt.Sprintf("Result: %s", aurora.Green(string(jsonBytes))))
			} else {
				fmt.Println(fmt.Sprintf("Result: %s", aurora.Yellow(string(jsonBytes))))
			}
		}

	},
}

// func init() {
// 	projectCmd.AddCommand(deployEnvCmd)
// }
