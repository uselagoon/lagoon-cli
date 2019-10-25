package cmd

import (
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/graphql"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var projectListCmd = &cobra.Command{
	Use:   "list",
	Short: "Show your projects",
	Run: func(cmd *cobra.Command, args []string) {

		var responseData WhatIsThere
		err := graphql.GraphQLRequest(`
query whatIsThere {
	allProjects {
		id
		gitUrl
		name,
		developmentEnvironmentsLimit,
		environments {
		  environmentType,
		  route
		}
	  }
}
`, &responseData)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetAutoWrapText(true)
			table.SetHeader([]string{"ID", "Project Name", "Git URL", "Dev Environments"})
			for _, project := range responseData.AllProjects {
				table.Append([]string{
					fmt.Sprintf("%d", project.ID),
					project.Name,
					project.GitURL,
					fmt.Sprintf("%d/%d", len(project.Environments), project.DevelopmentEnvironmentsLimit),
				})
			}
			table.Render()
			fmt.Println()
			fmt.Println("To view a project's details, run `lagoon project info {name}`.")
		}
	},
}

func init() {
	projectCmd.AddCommand(projectListCmd)

}
