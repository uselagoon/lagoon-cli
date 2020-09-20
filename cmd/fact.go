package cmd

import (
	"context"
	// "encoding/json"
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/internal/lagoon"
	"github.com/amazeeio/lagoon-cli/internal/lagoon/client"
	// "github.com/amazeeio/lagoon-cli/pkg/output"
	"github.com/spf13/cobra"
	// "github.com/spf13/pflag"
	"github.com/spf13/viper"	
)

var factCmd = &cobra.Command{
	Use:   "fact",
	Short: "Add and update facts",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateToken(viper.GetString("current")) // get a new token if the current one is invalid
	},
}

var addFactCommand = &cobra.Command{
	Use: "add",
	// Aliases: []String{"f"},
	Short: "Add a fact",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if cmdProjectName == "" || cmdProjectEnvironment == "" {
			fmt.Println("Missing arguments: Project name or environment name is not defined")
			cmd.Help()
			os.Exit(1)
		}

		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}

		current := viper.GetString("current")
		lc := client.New(
			viper.GetString("lagoons."+current+".graphql"),
			viper.GetString("lagoons."+current+".token"),
			viper.GetString("lagoons."+current+".version"),
			lagoonCLIVersion,
			debug)

		projectDetails, err := lagoon.GetProjectByNameForFacts(
			context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		if projectDetails.ID != 1 {
			fmt.Println(projectDetails.ID)
		}

		retval, errorval := lagoon.AddFact(context.TODO(), 5, "testname", "testval", lc)
		cmd.Println(retval)
		cmd.Println(errorval)
		return nil
	},
}


func init() {
	factCmd.AddCommand(addFactCommand)
}