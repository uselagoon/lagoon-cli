package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/mglaman/lagoon/graphql"

	"github.com/logrusorgru/aurora"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var projectInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Details about a project",
	Run: func(cmd *cobra.Command, args []string) {
		if cmdProject.Name == "" {
			if len(args) == 0 {
				// @todo list current projects and allow choosing?
				fmt.Println("You must provide a project name.")
				os.Exit(1)
			}
			if len(args) > 1 {
				fmt.Println("Too many arguments.")
				os.Exit(1)
			}
			cmdProject.Name = args[0]
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
}`, cmdProject.Name), &responseData)

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

		fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Project"), cmdProject.Name))
		if cmdProject.Environment != "" {
			if cmdProject.Environment == strings.TrimSpace(project.ProductionEnvironment) {
				fmt.Println(fmt.Sprintf("%s: %s (production)", aurora.Yellow("Environment"), cmdProject.Environment))
			} else {
				fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Environment"), cmdProject.Environment))
			}
		}
		fmt.Println()
		fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Git"), project.GitURL))
		fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Branches"), project.Branches))
		fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Pull Requests"), project.Pullrequests))
		fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Production Environment"), project.ProductionEnvironment))
		fmt.Println(fmt.Sprintf("%s: %d / %d", aurora.Yellow("Development Environments"), currentDevEnvironments, project.DevelopmentEnvironmentsLimit))
		fmt.Println()
		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(false)
		table.SetAutoFormatHeaders(false)
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		table.SetHeader([]string{"Name", "Deploy Type", "Environment", "Route", "SSH"})
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
