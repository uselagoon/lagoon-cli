package cmd

import (
	"context"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
)

var projectListCmd = &cobra.Command{
	Use:   "list",
	Short: "Show your projects",
	Run: func(cmd *cobra.Command, args []string) {
		client := GraphQLClient()
		req := GraphQLRequest(`
query whatIsThere {
  allProjects {
    gitUrl
    name
  }
}
`)
		//var responseData map[string]interface{}
		var responseData WhatIsThere
		ctx := context.Background()
		if err := client.Run(ctx, req, &responseData); err != nil {
			panic(err)
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(true)
		table.SetHeader([]string{"Name", "Git URL"})
		for _, project := range responseData.AllProjects {
			table.Append([]string{
				project.Name,
				project.GitURL,
			})
		}
		table.Render()
		fmt.Println()
		fmt.Println("To view a project's details, run `lagoon project info {name}`.")
	},
}

func init() {
	projectCmd.AddCommand(projectListCmd)
}
