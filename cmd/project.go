package cmd

import (
	"fmt"
	"os"

	"github.com/mglaman/lagoon/app"
	"github.com/mglaman/lagoon/graphql"

	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Show your projects, or details about a project",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if !graphql.HasValidToken() {
			fmt.Println("Need to run `lagoon login` first")
			os.Exit(1)
		}
		cmdProject, err := app.GetLocalProject()
		if err == nil {
			cmdProjectName = cmdProject.Name
		}
	},
}
