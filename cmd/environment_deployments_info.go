package cmd

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/mglaman/lagoon/graphql"
	"github.com/spf13/cobra"
	"os"
)

var deploymentInfoRemoteId = ""
var environmentDeploymentsInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Deployment details",
	Run: func(cmd *cobra.Command, args []string) {
		if cmdProject.Name == "" || cmdProject.Environment == "" {
			if len(args) != 2 {
				fmt.Println("You must provide an environment name and deployment ID.")
				os.Exit(1)
			}
			environmentInfoName = args[0]
			deploymentInfoRemoteId = args[1]

		} else {
			if len(args) != 1 {
				fmt.Println("You must provide a deployment ID")
				os.Exit(1)
			}
			environmentInfoName = fmt.Sprintf("%s-%s", cmdProject.Name, cmdProject.Environment)
			deploymentInfoRemoteId = args[0]
		}
		fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Deployment:"), deploymentInfoRemoteId))
		var responseData DeploymentByRemoteId
		err := graphql.GraphQLRequest(fmt.Sprintf(`query {
  deploymentByRemoteId(id:"%s") {
    id,
    name,
    status,
    buildLog
  }
}
`, deploymentInfoRemoteId), &responseData)
		if err != nil {
			panic(err)
		}
		status := responseData.Deployment.Status
		if status == "failed" {
			status = aurora.Red(status).String()
		}
		if status == "complete" {
			status = aurora.Green(status).String()
		}
		fmt.Println(fmt.Sprintf("Status: %s", status))
		fmt.Println(responseData.Deployment.BuildLog)
	},
}

func init() {
	environmentDeploymentsCmd.AddCommand(environmentDeploymentsInfoCmd)
}
