package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/amazeeio/lagoon-cli/output"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// LagoonConfigFlags .
type LagoonConfigFlags struct {
	Lagoon   string `json:"lagoon,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	Port     string `json:"port,omitempty"`
	GraphQL  string `json:"graphql,omitempty"`
	Token    string `json:"token,omitempty"`
}

func parseLagoonConfig(flags pflag.FlagSet) LagoonConfigFlags {
	configMap := make(map[string]interface{})
	flags.VisitAll(func(f *pflag.Flag) {
		if flags.Changed(f.Name) {
			configMap[f.Name] = f.Value
		}
	})
	jsonStr, _ := json.Marshal(configMap)
	parsedFlags := LagoonConfigFlags{}
	json.Unmarshal(jsonStr, &parsedFlags)
	return parsedFlags
}

var configCmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{"c"},
	Short:   "Configure Lagoon CLI",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}

var configDefaultCmd = &cobra.Command{
	Use:   "default [-l lagoonname]",
	Short: "Set the default Lagoon to use",
	Run: func(cmd *cobra.Command, args []string) {
		lagoonConfig := parseLagoonConfig(*cmd.Flags())
		if lagoonConfig.Lagoon == "" {
			fmt.Println("Not enough arguments")
			cmd.Help()
			os.Exit(1)
		}
		viper.Set("default", strings.TrimSpace(string(lagoonConfig.Lagoon)))
		err := viper.WriteConfig()
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}

		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"default-lagoon": lagoonConfig.Lagoon,
			},
		}
		output.RenderResult(resultData, outputOptions)
		//fmt.Println(fmt.Sprintf("Updating default lagoon to %s", lagoonName))
	},
}

var configLagoonsCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "View all configured Lagoon instances",
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
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add information about an additional Lagoon instance to use",
	Run: func(cmd *cobra.Command, args []string) {
		lagoonConfig := parseLagoonConfig(*cmd.Flags())
		if lagoonConfig.Lagoon == "" {
			fmt.Println("Not enough arguments. Requires: lagoon name")
			cmd.Help()
			os.Exit(1)
		}

		if lagoonConfig.Hostname != "" && lagoonConfig.Port != "" && lagoonConfig.GraphQL != "" {
			viper.Set("lagoons."+lagoonConfig.Lagoon+".hostname", lagoonConfig.Hostname)
			viper.Set("lagoons."+lagoonConfig.Lagoon+".port", lagoonConfig.Port)
			viper.Set("lagoons."+lagoonConfig.Lagoon+".graphql", lagoonConfig.GraphQL)
			if lagoonConfig.Token != "" {
				viper.Set("lagoons."+lagoonConfig.Lagoon+".token", lagoonConfig.Token)
			}
			err := viper.WriteConfig()
			if err != nil {
				output.RenderError(err.Error(), outputOptions)
				os.Exit(1)
			}
			resultData := output.Result{
				Result: "success",
				ResultData: map[string]interface{}{
					"lagoon":   lagoonConfig.Lagoon,
					"hostname": lagoonConfig.Hostname,
					"graphql":  lagoonConfig.GraphQL,
					"port":     lagoonConfig.Port,
				},
			}
			output.RenderResult(resultData, outputOptions)
		} else {
			output.RenderError("Must have Hostname, Port, and GraphQL endpoint", outputOptions)
		}
	},
}

var configDeleteCmd = &cobra.Command{
	Use:     "delete [-l lagoonname]",
	Aliases: []string{"d"},
	Short:   "Delete a Lagoon instance configuration",
	Run: func(cmd *cobra.Command, args []string) {
		lagoonConfig := parseLagoonConfig(*cmd.Flags())

		if lagoonConfig.Lagoon == "" {
			fmt.Println("Not enough arguments. Requires: lagoon name")
			cmd.Help()
			os.Exit(1)
		}
		fmt.Println(fmt.Sprintf("Deleting config for lagoon: %s", lagoonConfig.Lagoon))
		if yesNo() {
			err := unset(lagoonConfig.Lagoon)
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
