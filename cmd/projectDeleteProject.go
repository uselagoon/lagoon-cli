package cmd

import (
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/graphql"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var projectDeleteProjectCmd = &cobra.Command{
	Use:   "project [project name]",
	Short: "Delete a project",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Println("Not enough arguments. Requires: project name.")
			cmd.Help()
			os.Exit(1)
		}
		projectName := args[0]

		fmt.Println(fmt.Sprintf("Deleting %s", projectName))

		if yesNo() {
			var responseData DeleteProjectResult
			err := graphql.GraphQLRequest(fmt.Sprintf(`mutation {
    deleteProject(
      input: {
        project:"%s"
      }
    )
  }`, projectName), &responseData)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				if responseData.DeleteProject == "success" {
					fmt.Println(fmt.Sprintf("Result: %s", aurora.Green(responseData.DeleteProject)))
				} else {
					fmt.Println(fmt.Sprintf("Result: %s", aurora.Yellow(responseData.DeleteProject)))
				}
			}
		}

	},
}

func init() {
	projectDeleteCmd.AddCommand(projectDeleteProjectCmd)
}
