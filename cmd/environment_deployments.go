package cmd

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"time"

	"github.com/mglaman/lagoon/graphql"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var environmentDeploymentsCmd = &cobra.Command{
	Use:   "deployments",
	Short: "Environment information",
	Run: func(cmd *cobra.Command, args []string) {
		if cmdProject.Name == "" || cmdProject.Environment == "" {
			if len(args) == 0 {
				fmt.Println("You must provide an environment name.")
				os.Exit(1)
			}
			environmentInfoName = args[0]
		} else {
			environmentInfoName = fmt.Sprintf("%s-%s", cmdProject.Name, cmdProject.Environment)
		}
		fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Deployments:"), environmentInfoName))

		year, month, _ := time.Now().Date()
		var responseData EnvironmentByOpenshiftProjectName
		err := graphql.GraphQLRequest(fmt.Sprintf(`query {
	environmentByOpenshiftProjectName(openshiftProjectName: "%[1]s") {
		deployments {
		  id,
		  name,
		  status,
		  created,
		  started,
		  completed,
		  remoteId
		}
	}
}
`, environmentInfoName, year, month), &responseData)
		if err != nil {
			panic(err)
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(true)
		table.SetHeader([]string{"ID", "Status", "Started", "Completed"})
		for _, deployment := range responseData.Environment.Deployments {
			status := deployment.Status
			if status == "failed" {
				status = aurora.Red(status).String()
			}
			if status == "complete" {
				status = aurora.Green(status).String()
			}
			table.Append([]string{
				deployment.RemoteID,
				status,
				deployment.Started,
				deployment.Completed,
			})
		}
		table.Render()
		fmt.Println()
		fmt.Println("To view an environment's details, run `lagoon environment deployments info {id}`.")
	},
}

func init() {
	environmentCmd.AddCommand(environmentDeploymentsCmd)
}
