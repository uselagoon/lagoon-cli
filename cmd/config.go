package cmd

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/amazeeio/lagoon-cli/output"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure Lagoon CLI",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}

var configDefaultCmd = &cobra.Command{
	Use:   "default [lagoon name]",
	Short: "Set the default Lagoon to use",
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
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}

		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"default-lagoon": lagoonName,
			},
		}
		output.RenderResult(resultData, outputOptions)
		//fmt.Println(fmt.Sprintf("Updating default lagoon to %s", lagoonName))
	},
}

var configLagoonsCmd = &cobra.Command{
	Use:   "list",
	Short: "View all configured Lagoon instances",
	Run: func(cmd *cobra.Command, args []string) {
		lagoons := viper.Get("lagoons")
		lagoonsMap := reflect.ValueOf(lagoons).MapKeys()
		if !outputOptions.CSV && !outputOptions.JSON {
			fmt.Println("You have the following Lagoon instances configured:")
			for _, lagoon := range lagoonsMap {
				fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Name"), lagoon))
				fmt.Println(fmt.Sprintf(" - %s: %s", aurora.Yellow("Hostname"), viper.GetString("lagoons."+lagoon.String()+".hostname")))
				fmt.Println(fmt.Sprintf(" - %s: %s", aurora.Yellow("GraphQL"), viper.GetString("lagoons."+lagoon.String()+".graphql")))
				fmt.Println(fmt.Sprintf(" - %s: %d", aurora.Yellow("Port"), viper.GetInt("lagoons."+lagoon.String()+".port")))
			}
			fmt.Println("\nYour default Lagoon is:")
			fmt.Println(fmt.Sprintf("%s: %s\n", aurora.Yellow("Name"), viper.Get("default")))
			fmt.Println("Your current lagoon is:")
			fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Name"), viper.Get("current")))
		} else {
			var lagoonsData []map[string]interface{}
			for _, lagoon := range lagoonsMap {
				lagoonMapData := map[string]interface{}{
					"name":     fmt.Sprintf("%s", lagoon),
					"hostname": viper.GetString("lagoons." + lagoon.String() + ".hostname"),
					"graphql":  viper.GetString("lagoons." + lagoon.String() + ".graphql"),
					"port":     viper.GetString("lagoons." + lagoon.String() + ".port"),
				}
				lagoonsData = append(lagoonsData, lagoonMapData)
			}
			returnedData := map[string]interface{}{
				"lagoons":        lagoonsData,
				"default-lagoon": viper.Get("default"),
				"current-lagoon": viper.Get("current"),
			}
			output.RenderJSON(returnedData, outputOptions)
		}
	},
}
var configAddCmd = &cobra.Command{
	Use:   "add [lagoon name]",
	Short: "Add information about an additional Lagoon instance to use",
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
				output.RenderError(err.Error(), outputOptions)
				os.Exit(1)
			}
			resultData := output.Result{
				Result: "success",
				ResultData: map[string]interface{}{
					"lagoon":   lagoonName,
					"hostname": lagoonHostname,
					"graphql":  lagoonGraphQL,
					"port":     lagoonPort,
				},
			}
			output.RenderResult(resultData, outputOptions)
		} else {
			output.RenderError("Must have Hostname, Port, and GraphQL endpoint", outputOptions)
		}
	},
}

var configDeleteCmd = &cobra.Command{
	Use:   "delete [lagoon name]",
	Short: "Delete a Lagoon instance configuration",
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
				output.RenderError(err.Error(), outputOptions)
				os.Exit(1)
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
