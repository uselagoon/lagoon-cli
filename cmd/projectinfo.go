package cmd

import (
	"fmt"
	"github.com/mglaman/lagoon/graphql"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var projectInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Details about a project",
	Run: func(cmd *cobra.Command, args []string) {
		if cmdProjectName == "" {
			if len(args) == 0 {
				// @todo list current projects and allow choosing?
				fmt.Println("You must provide a project name.")
				os.Exit(1)
			}
			if len(args) > 1 {
				fmt.Println("Too many arguments.")
				os.Exit(1)
			}
			cmdProjectName = args[0]
		}
		var responseData ProjectByName
		err := graphql.GraphQLRequest(fmt.Sprintf(`query {
  projectByName(name: "%s") {
    id,
    name,
    gitUrl,
    subfolder,
    branches,
    pullrequests,
    productionEnvironment,
    environments {
      name
      environmentType
      deployType
      route
    }
    autoIdle,
    storageCalc,
    developmentEnvironmentsLimit,
  }
}`, cmdProjectName), &responseData)

		if err != nil {
			panic(err)
		}
		project := responseData.ProjectByName
		var currentDevEnvironments = 0
		for _, environment := range project.Environments {
			if environment.EnvironmentType == "development" {
				currentDevEnvironments++
			}
		}

		fmt.Println(project.Name)
		fmt.Println()
		fmt.Println(fmt.Sprintf("Git URL: %s", project.GitURL))
		fmt.Println(fmt.Sprintf("Branches Pattern: %s", project.Branches))
		fmt.Println(fmt.Sprintf("Pull Requests: %s", project.Pullrequests))
		fmt.Println(fmt.Sprintf("Production Environment: %s", project.ProductionEnvironment))
		fmt.Println(fmt.Sprintf("Development Environments: %d / %d", currentDevEnvironments, project.DevelopmentEnvironmentsLimit))
		fmt.Println()
		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(true)
		table.SetHeader([]string{"Name", "Deploy Type", "Environment Type", "Route", "SSH"})
		for _, environment := range project.Environments {
			table.Append([]string{
				environment.Name,
				environment.DeployType,
				environment.EnvironmentType,
				environment.Route,
				fmt.Sprintf("ssh -p %s -t %s-%s@%s", viper.GetString("lagoon_port"), project.Name, environment.Name, viper.GetString("lagoon_hostname")),
			})
		}
		table.Render()
	},
}

func init() {
	projectCmd.AddCommand(projectInfoCmd)
}
