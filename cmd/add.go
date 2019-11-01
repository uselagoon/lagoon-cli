package cmd

import (
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/graphql"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a project, environment, or variable",
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

var addVariableCmd = &cobra.Command{
	Use:   "variable",
	Short: "Add variables on environments or projects",
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

func init() {
	addCmd.AddCommand(addVariableCmd)
}
