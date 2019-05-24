package cmd

import (
	"fmt"
	"os"

	"github.com/mglaman/lagoon/app"
	"github.com/mglaman/lagoon/graphql"

	"github.com/spf13/cobra"
)

var environmentCmd = &cobra.Command{
	Use:   "environment",
	Short: "Show a project's environment",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if !graphql.HasValidToken() {
			fmt.Println("Need to run `lagoon login` first")
			os.Exit(1)
		}
		cmdProject, _ = app.GetLocalProject()
	},
}
