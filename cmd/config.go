package cmd

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configDefaultCmd = &cobra.Command{
	Use:   "default [lagoon name]",
	Short: "Set the default lagoon to use",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Not enough arguments. Requires: lagoon name")
			cmd.Help()
			os.Exit(1)
		}
		lagoonName := args[0]
		viper.Set("default", strings.TrimSpace(string(lagoonName)))
		err := viper.WriteConfig()
		if err != nil {
			panic(err)
		}

		fmt.Println(fmt.Sprintf("Updating default lagoon to %s", lagoonName))
	},
}

var configLagoonsCmd = &cobra.Command{
	Use:   "lagoons",
	Short: "View lagoons",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You have the following lagoons configured:")
		lagoons := viper.Get("lagoons")
		lagoonsMap := reflect.ValueOf(lagoons).MapKeys()
		for _, lagoon := range lagoonsMap {
			fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Name"), lagoon))
			fmt.Println(fmt.Sprintf(" - %s: %s", aurora.Yellow("Hostname"), viper.GetString("lagoons."+lagoon.String()+".hostname")))
			fmt.Println(fmt.Sprintf(" - %s: %s", aurora.Yellow("GraphQL"), viper.GetString("lagoons."+lagoon.String()+".graphql")))
			fmt.Println(fmt.Sprintf(" - %s: %d", aurora.Yellow("Port"), viper.GetInt("lagoons."+lagoon.String()+".port")))
		}
		fmt.Println("\nYour default lagoon is:")
		fmt.Println(fmt.Sprintf("%s: %s\n", aurora.Yellow("Name"), viper.Get("default")))
		fmt.Println("Your current lagoon is:")
		fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Name"), viper.Get("current")))
	},
}
var configAddCmd = &cobra.Command{
	Use:   "add [lagoon name]",
	Short: "Add a lagoon configuration to use",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Not enough arguments. Requires: lagoon name")
			cmd.Help()
			os.Exit(1)
		}
		lagoonName := args[0]

		if lagoonHostname == "" && lagoonPort == "" && lagoonGraphQL == "" {
			lagoonHostname = Prompt(fmt.Sprintf("Lagoon Hostname (%s)", viper.GetString("lagoons."+lagoonName+".hostname")))
			lagoonPort = Prompt(fmt.Sprintf("Lagoon Port (%d)", viper.GetInt("lagoons."+lagoonName+".port")))
			lagoonGraphQL = Prompt(fmt.Sprintf("Lagoon GraphQL endpoint (%s)", viper.GetString("lagoons."+lagoonName+".graphql")))
		}
		if lagoonHostname != "" && lagoonPort != "" && lagoonGraphQL != "" {
			viper.Set("lagoons."+lagoonName+".hostname", lagoonHostname)
			viper.Set("lagoons."+lagoonName+".port", lagoonPort)
			viper.Set("lagoons."+lagoonName+".graphql", lagoonGraphQL)
			if lagoonToken != "" {
				viper.Set("lagoons."+lagoonName+".token", lagoonToken)
			}
			err := viper.WriteConfig()
			if err != nil {
				panic(err)
			}
			fmt.Println(fmt.Sprintf("\nAdded a new lagoon named: %s", lagoonName))
		} else {
			fmt.Println(fmt.Sprintf("\nMust have Hostname, Port, and GraphQL endpoint"))
		}
	},
}

var configDeleteCmd = &cobra.Command{
	Use:   "delete [lagoon name]",
	Short: "Delete a lagoon configuration",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Not enough arguments. Requires: lagoon name")
			cmd.Help()
			os.Exit(1)
		}
		lagoonName := args[0]
		fmt.Println(fmt.Sprintf("Deleting config for lagoon: %s", lagoonName))
		if yesNo() {
			err := unset(lagoonName)
			if err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	configCmd.AddCommand(configDefaultCmd)
	configCmd.AddCommand(configLagoonsCmd)
	configCmd.AddCommand(configAddCmd)
	configCmd.AddCommand(configDeleteCmd)
	configAddCmd.Flags().StringVarP(&lagoonHostname, "hostname", "H", "", "Lagoon SSH hostname")
	configAddCmd.Flags().StringVarP(&lagoonPort, "port", "P", "", "Lagoon SSH port")
	configAddCmd.Flags().StringVarP(&lagoonGraphQL, "graphql", "g", "", "Lagoon GraphQL endpoint")
	configAddCmd.Flags().StringVarP(&lagoonToken, "token", "t", "", "Lagoon GraphQL token")
}
