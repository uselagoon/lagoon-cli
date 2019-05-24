package cmd

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"os"
	"time"

	"github.com/mglaman/lagoon/graphql"

	"code.cloudfoundry.org/bytefmt"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

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

		year, month, _ := time.Now().Date()
		var responseData EnvironmentByOpenshiftProjectName
		err := graphql.GraphQLRequest(fmt.Sprintf(`query {
	environmentByOpenshiftProjectName(openshiftProjectName: "%[1]s") {
		deployType,
		environmentType
		hitsMonth(month: "%[2]d-%[3]d"){
			total
		}
		storageMonth(month: "%[2]d-%[3]d") {
			bytesUsed
		}
	}
}
`, environmentInfoName, year, month), &responseData)
		if err != nil {
			panic(err)
		}
		p := message.NewPrinter(language.English)
		fmt.Println()
		fmt.Println(p.Sprintf("%s: %s", aurora.Yellow("Mode"), responseData.Environment.EnvironmentType))
		fmt.Println(p.Sprintf("%s: %s", aurora.Yellow("Type"), responseData.Environment.DeployType))
		fmt.Println()
		_, _ = p.Println("Usage this month")
		_, _ = p.Println("----------------")
		fmt.Println(p.Sprintf("%s: %d", aurora.Yellow("Hits"), responseData.Environment.HitsMonth.Total))
		fmt.Println(p.Sprintf("%s: %s", aurora.Yellow("Storage"), bytefmt.ByteSize(uint64(responseData.Environment.StorageMonth.BytesUsed))))
	},
}

func init() {
	environmentCmd.AddCommand(environmentInfoCmd)
}
