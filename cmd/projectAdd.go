package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/amazeeio/lagoon-cli/graphql"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var projectAddCmd = &cobra.Command{
	Use:   "add [project name]",
	Short: "Show your projects",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Println("Not enough arguments. Requires: name")
			cmd.Help()
			os.Exit(1)
		}
		projectName := args[0]

		openshift, _ := strconv.Atoi(projectOpenshift)
		if projectGitURL == "" && projectProductionEnv == "" && projectOpenshift == "" {
			projectGitURL = Prompt("Git URL")
			projectProductionEnv = Prompt("Production Environment")
			projectOpenshift = Prompt("Openshift")
		}
		if projectGitURL != "" && projectProductionEnv != "" && projectOpenshift != "" {
			var responseData AddProject
			err := graphql.GraphQLRequest(fmt.Sprintf(`
	mutation {
	  addProject(
	    input: {
	      name: "%s"
	      openshift: %d
	      gitUrl: "%s"
	      productionEnvironment: "%s"
	    }
	  ) {
			id
	    name
	    gitUrl
	  }
	}
	`, projectName, openshift, projectGitURL, projectProductionEnv), &responseData)

			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(fmt.Sprintf("Result: %s\n", aurora.Green("success")))
				fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Project"), responseData.AddProject.Name))
				fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Git"), responseData.AddProject.GitURL))
			}
		} else {
			fmt.Println(fmt.Sprintf("\nMust have giturl, production environment, and openshift"))
		}
	},
}

func init() {
	projectCmd.AddCommand(projectAddCmd)
	projectAddCmd.Flags().StringVarP(&projectGitURL, "giturl", "g", "", "Git URL")
	projectAddCmd.Flags().StringVarP(&projectProductionEnv, "production", "P", "", "Production Environment")
	projectAddCmd.Flags().StringVarP(&projectOpenshift, "openshift", "o", "", "Openshift")
}
