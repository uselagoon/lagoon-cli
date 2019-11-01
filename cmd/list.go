package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/graphql"
	"github.com/amazeeio/lagoon-cli/lagoon/projects"
	"github.com/amazeeio/lagoon-cli/output"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list projects, environment, etc..",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// get a new token if the current one is invalid
		valid := graphql.VerifyTokenExpiry()
		if valid == false {
			loginErr := loginToken()
			if loginErr != nil {
				fmt.Println("Unable to refresh token, you may need to run `lagoon login` first")
				os.Exit(1)
			}
		}
	},
}

var listProjectCmd = &cobra.Command{
	Use:   "projects",
	Short: "Show your projects",
	Run: func(cmd *cobra.Command, args []string) {

		returnedJSON, err := projects.ListAllProjects()
		if err != nil {
			fmt.Println(err)
			return
		}

		var dataMain output.Table
		err = json.Unmarshal([]byte(returnedJSON), &dataMain)
		if err != nil {
			fmt.Println(err)
			return
		}
		output.RenderTable(dataMain)

	},
}

var listVariablesCmd = &cobra.Command{
	Use:   "variables",
	Short: "Show your variables for a project",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Println("Not enough arguments. Requires: project name.")
			cmd.Help()
			os.Exit(1)
		}
		projectName := args[0]

		returnedJSON, err := projects.ListEnvironmentVariables(projectName, revealValue)
		if err != nil {
			fmt.Println(err)
			return
		}

		var dataMain output.Table
		err = json.Unmarshal([]byte(returnedJSON), &dataMain)
		if err != nil {
			fmt.Println(err)
			return
		}
		output.RenderTable(dataMain)
	},
}

func init() {
	listCmd.AddCommand(listProjectCmd)
	listCmd.AddCommand(listVariablesCmd)
	listVariablesCmd.Flags().BoolVarP(&revealValue, "reveal", "r", false, "Reveal the variable values")
}
