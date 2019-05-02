package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure Lagoon CLI",
	Run: func(cmd *cobra.Command, args []string) {
		lagoonHostname := Prompt(fmt.Sprintf("Lagoon Hostname (%s)", viper.GetString("lagoon_hostname")))
		lagoonPort := Prompt(fmt.Sprintf("Lagoon Port (%s)", viper.GetString("lagoon_port")))
		lagoonGraphQL := Prompt(fmt.Sprintf("Lagoon GraphQL endpoint (%s)", viper.GetString("lagoon_graphql")))

		viper.Set("lagoon_hostname", lagoonHostname)
		viper.Set("lagoon_port", lagoonPort)
		viper.Set("lagoon_graphql", lagoonGraphQL)

		fmt.Println("Lagoon CLI is now configured, run `lagoon login` to generate your JWT access token.")
	},
}

var inputScanner = bufio.NewScanner(os.Stdin)

// GetInput reads input from an input buffer and returns the result as a string.
func GetInput() string {
	inputScanner.Scan()
	return strings.TrimSpace(inputScanner.Text())
}

// Prompt gets input with a prompt and returns the input
func Prompt(prompt string) string {
	fullPrompt := fmt.Sprintf("%s", prompt)
	fmt.Print(fullPrompt + ": ")
	return GetInput()
}
