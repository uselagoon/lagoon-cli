package cmd

import (
	"fmt"
	"reflect"

	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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

func init() {
	configCmd.AddCommand(configLagoonsCmd)
}
