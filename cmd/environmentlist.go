package cmd

import (
	"fmt"
	"os"

	"github.com/mglaman/lagoon/graphql"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var environmentListCmd = &cobra.Command{
	Use:   "list",
	Short: "List environments for a project",
	Run: func(cmd *cobra.Command, args []string) {
		if cmdProject.Name == "" {
			if len(args) == 0 {
				// @todo list current projects and allow choosing?
				fmt.Println("You must provide a project name.")
				os.Exit(1)
			}
			cmdProject.Name = args[0]
		}
		var responseData ProjectByName
		err := graphql.GraphQLRequest(fmt.Sprintf(`query {
  projectByName(name: "%s") {
    environments {
		name
		environmentType
		deployType
		route
		openshiftProjectName
		envVariables {
		  id
		  name
		  value
		}
		routes
		monitoringUrls
		deployments {
		  name
		  status
		  created
		  started
		  completed
		  remoteId
		}
		services {
		  name
		  id
		}
	  }
  }
}`, cmdProject.Name), &responseData)
		if err != nil {
			panic(err)
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(true)
		table.SetHeader([]string{"Name", "Environment", "Type", "URL"})
		for _, environment := range responseData.ProjectByName.Environments {
			table.Append([]string{
				environment.OpenshiftProjectName,
				environment.EnvironmentType,
				environment.DeployType,
				environment.Route,
			})
		}
		table.Render()
		fmt.Println()
		fmt.Println("To view an environment's details, run `lagoon environment info {name}`.")
	},
}

func init() {
	environmentCmd.AddCommand(environmentListCmd)
}
