package cmd

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var environmentInfoName = ""
var environmentInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Environment information",
	Run: func(cmd *cobra.Command, args []string) {
		if cmdProject.Name == "" || cmdProject.Environment == "" {
			if len(args) == 0 {
				fmt.Println("You must provide an environment name.")
				os.Exit(1)
			}
			environmentInfoName = args[0]
		} else {
			environmentInfoName = fmt.Sprintf("%s-%s", cmdProject.Name, cmdProject.Environment)
		}
		fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Environment"), environmentInfoName))
	},
}

func init() {
	environmentCmd.AddCommand(environmentInfoCmd)
}
