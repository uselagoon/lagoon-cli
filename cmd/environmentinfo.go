package cmd

import (
	"fmt"
	"os"

	"github.com/mglaman/lagoon/graphql"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var environmentInfoName = ""
var environmentInfoCmd = &cobra.Command{
	Use:   "info",
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
		fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Environment"), environmentInfoName))

		var responseData EnvironmentByOpenshiftProjectName
		err := graphql.GraphQLRequest(fmt.Sprintf(`query {
	environmentByOpenshiftProjectName(openshiftProjectName: "%s") {
		deployType,
		environmentType
		hitsMonth(month: 0){
			total
		}
		storageMonth(month: "2019-05") {
			bytesUsed
		}
	}
}
`, environmentInfoName), &responseData)
		if err != nil {
			panic(err)
		}
		fmt.Println()
		fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Mode"), responseData.Environment.EnvironmentType))
		fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Type"), responseData.Environment.DeployType))
		fmt.Println()
		fmt.Println(fmt.Sprintf("%s: %d", aurora.Yellow("Hits this month"), responseData.Environment.HitsMonth.Total))
		fmt.Println(fmt.Sprintf("%s: %d", aurora.Yellow("Storage this month"), responseData.Environment.StorageMonth.BytesUsed))
	},
}

func init() {
	environmentCmd.AddCommand(environmentInfoCmd)
}
