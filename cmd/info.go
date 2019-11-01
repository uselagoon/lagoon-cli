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

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info on projects, environments, etc..",
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

var infoProjectCmd = &cobra.Command{
	Use:   "project [project]",
	Short: "Details about a project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Not enough arguments. Requires: project name")
			cmd.Help()
			os.Exit(1)
		}
		projectName := args[0]

		returnedJSON, err := projects.ListEnvironmentsForProject(projectName)
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

var infoRocketChatsCmd = &cobra.Command{
	Use:   "rocketchat [project]",
	Short: "Rocketchat details about a project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Not enough arguments. Requires: project name")
			cmd.Help()
			os.Exit(1)
		}
		projectName := args[0]

		returnedJSON, err := projects.ListAllProjectRocketChats(projectName)
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

var infoSlackCmd = &cobra.Command{
	Use:   "slack [project]",
	Short: "Slack details about a project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Not enough arguments. Requires: project name")
			cmd.Help()
			os.Exit(1)
		}
		projectName := args[0]

		returnedJSON, err := projects.ListAllProjectSlacks(projectName)
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
	infoCmd.AddCommand(infoProjectCmd)
	infoCmd.AddCommand(infoRocketChatsCmd)
	infoCmd.AddCommand(infoSlackCmd)
}
