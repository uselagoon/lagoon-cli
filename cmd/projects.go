package cmd

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

type ProjectByName struct {
	ProjectByName Project `json:"projectByName"`
}
type WhatIsThere struct {
	AllProjects []Project `json:"allProjects"`
}
type Environments struct {
	Name            string `json:"name"`
	EnvironmentType string `json:"environmentType"`
	DeployType      string `json:"deployType"`
	Route           string `json:"route"`
}
type Project struct {
	ID                           int            `json:"id"`
	GitURL                       string         `json:"gitUrl"`
	Subfolder                    string         `json:"subfolder"`
	Name                         string         `json:"name"`
	Branches                     string         `json:"branches"`
	Pullrequests                 string         `json:"pullrequests"`
	ProductionEnvironment        string         `json:"productionEnvironment"`
	Environments                 []Environments `json:"environments"`
	AutoIdle                     int            `json:"autoIdle"`
	DevelopmentEnvironmentsLimit int            `json:"developmentEnvironmentsLimit"`
}

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Show your projects, or details about a project",
	Run: func(cmd *cobra.Command, args []string) {
		apiToken := viper.GetString("lagoon_token")
		if apiToken == "" {
			fmt.Println("Need to run `lagoon login` first")
			os.Exit(1)
		}

		if len(args) == 0 {
			listProjects()
		} else {
			projectName := args[0]
			getProject(projectName)
		}
	},
}

func init() {
	rootCmd.AddCommand(projectsCmd)
}

func getProject(name string) {
	client := graphql.NewClient(viper.GetString("lagoon_graphql"))
	req := graphql.NewRequest(fmt.Sprintf(`query {
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
}`, name))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", viper.GetString("lagoon_token")))
	var responseData ProjectByName
	ctx := context.Background()
	if err := client.Run(ctx, req, &responseData); err != nil {
		panic(err)
	}
	project := responseData.ProjectByName
	var currentDevEnvironments int = 0
	for _, environment := range project.Environments {
		if environment.EnvironmentType == "development" {
			currentDevEnvironments++
		}
	}

	fmt.Println(name)
	fmt.Println()
	fmt.Println(fmt.Sprintf("Git URL: %s", project.GitURL))
	fmt.Println(fmt.Sprintf("Branches Pattern: %s", project.Branches))
	fmt.Println(fmt.Sprintf("Pull Requests: %s", project.Pullrequests))
	fmt.Println(fmt.Sprintf("Production Environment: %s", project.ProductionEnvironment))
	fmt.Println(fmt.Sprintf("Development Environments: %d / %d", currentDevEnvironments, project.DevelopmentEnvironmentsLimit))
	fmt.Println()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(true)
	table.SetHeader([]string{"Name", "Deploy Type", "Environment Type", "Route"})
	for _, environment := range project.Environments {
		table.Append([]string{
			environment.Name,
			environment.DeployType,
			environment.EnvironmentType,
			environment.Route,
		})
	}
	table.Render()
}

func listProjects() {
	client := graphql.NewClient(viper.GetString("lagoon_graphql"))
	req := graphql.NewRequest(`
query whatIsThere {
  allProjects {
    gitUrl
    name
  }
}
`)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", viper.GetString("lagoon_token")))
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
}
