package cmd

import (
	"errors"
	"fmt"
	"github.com/mglaman/lagoon/graphql"
	"os"

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
		gitUrl
		name,
		customer {
		  id,
		  name
		}
		environments {
		  environmentType,
		  route
		}
	  }
}
`, &responseData)
		if err != nil {
			panic(err)
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(true)
		table.SetHeader([]string{"Name", "Customer", "Git URL", "URL"})
		for _, project := range responseData.AllProjects {
			productionEnvironment, err := getProductionEnvironment(project.Environments)
			if err != nil {
				panic(err)
			}
			table.Append([]string{
				project.Name,
				project.Customer.Name,
				project.GitURL,
				productionEnvironment.Route,
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

func getProductionEnvironment(environments []Environments) (*Environments, error) {
	for _, environment := range environments {
		if environment.EnvironmentType == "production" {
			return &environment, nil
		}
	}
	return nil, errors.New("unable to determine production environment")
}
