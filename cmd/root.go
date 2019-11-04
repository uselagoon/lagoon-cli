package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/amazeeio/lagoon-cli/app"
	"github.com/amazeeio/lagoon-cli/graphql"
	"github.com/amazeeio/lagoon-cli/output"
	"github.com/manifoldco/promptui"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmdProject app.LagoonProject
var cmdLagoon = ""
var forceAction bool
var cmdSSHKey = ""
var inputScanner = bufio.NewScanner(os.Stdin)
var rootCmd = &cobra.Command{
	Use:   "lagoon",
	Short: "Command line integration for Lagoon",
	Long:  `Lagoon CLI. Manage your Lagoon hosted projects.`,
}

// version/build information
var (
	version string
	build   string
)

// Execute the root command.
func Execute() {
	viper.AutomaticEnv()
	if err := rootCmd.Execute(); err != nil {
		output.RenderError(err.Error(), outputOptions)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.EnableCommandSorting = false

	rootCmd.PersistentFlags().StringVarP(&cmdLagoon, "lagoon", "l", "", "The lagoon instance to interact with")
	rootCmd.PersistentFlags().BoolVarP(&forceAction, "force", "", false, "force")
	rootCmd.PersistentFlags().StringVarP(&cmdSSHKey, "ssh-key", "i", "", "Specify a specific SSH key to use")

	rootCmd.PersistentFlags().BoolVarP(&listAllProjects, "all-projects", "", false, "all projects (if supported)")
	rootCmd.PersistentFlags().BoolVarP(&outputOptions.Header, "no-header", "", false, "no header on table (if supported)")
	rootCmd.PersistentFlags().BoolVarP(&outputOptions.CSV, "output-csv", "", false, "output as csv")
	rootCmd.PersistentFlags().BoolVarP(&outputOptions.JSON, "output-json", "", false, "output as json")
	rootCmd.PersistentFlags().BoolVarP(&outputOptions.Pretty, "pretty", "", false, "make json pretty")

	rootCmd.SetUsageTemplate(`Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}

{{- if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
{{- $current_command:= . }}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}

{{if .HasAvailableLocalFlags -}}
Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(webCmd)
	// rootCmd.AddCommand(projectCmd)
	rootCmd.AddCommand(deployEnvCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(versionCmd)
	// rootCmd.AddCommand(sshEnvCmd) //@TODO
}

// version/build information command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:", version)
		fmt.Println("Build:", build)
		os.Exit(0)
	},
}

func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		output.RenderError(err.Error(), outputOptions)
		os.Exit(1)
	}

	var configName = ".lagoon"

	// Search config in home directory with name ".lagoon" (without extension).
	// @todo see if we can grok the proper info from the cwd .lagoon.yml
	viper.AddConfigPath(home)
	viper.SetConfigName(configName)
	viper.SetDefault("lagoons.amazeeio.hostname", "ssh.lagoon.amazeeio.cloud")
	viper.SetDefault("lagoons.amazeeio.port", 32222)
	viper.SetDefault("lagoons.amazeeio.token", "")
	viper.SetDefault("lagoons.amazeeio.graphql", "https://api.lagoon.amazeeio.cloud/graphql")
	viper.SetDefault("lagoons.amazeeio.ui", "https://ui-lagoon-master.ch.amazee.io")
	viper.SetDefault("lagoons.amazeeio.kibana", "https://logs-db-ui-lagoon-master.ch.amazee.io/")
	viper.SetDefault("default", "amazeeio")
	err = viper.ReadInConfig()
	if err != nil {
		err = viper.WriteConfigAs(filepath.Join(home, configName+".yml"))
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
	}
	if cmdLagoon == "" {
		if viper.GetString("default") == "" {
			cmdLagoon = "amazeeio"
		} else {
			cmdLagoon = viper.GetString("default")
		}
	}
	viper.Set("current", strings.TrimSpace(string(cmdLagoon)))
	err = viper.WriteConfig()
	if err != nil {
		output.RenderError(err.Error(), outputOptions)
		os.Exit(1)
	}

	validateToken(viper.GetString("current")) // get a new token if the current one is invalid
	// if the directory or repository you're in has a valid .lagoon.yml and docker-compose.yml with x-lagoon-project in it
	// we can use that inplaces where projects already exist so you don't have to type it out
	cmdProject, _ = app.GetLocalProject()

	// if !outputOptions.CSV && !outputOptions.JSON {
	// 	fmt.Println("Using Lagoon:", cmdLagoon)
	// }
}

func yesNo() bool {
	if forceAction != true {
		prompt := promptui.Select{
			Label: "Select[Yes/No]",
			Items: []string{"No", "Yes"},
		}
		_, result, err := prompt.Run()
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		return result == "Yes"
	}
	return true
}

func selectList(listItems []string) string {
	if forceAction != true {
		prompt := promptui.Select{
			Label: "Select item",
			Items: listItems,
		}
		_, result, err := prompt.Run()
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		return result
	}
	return ""
}

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

func unset(key string) error {
	delete(viper.Get("lagoons").(map[string]interface{}), key)
	err := viper.WriteConfig()
	if err != nil {
		output.RenderError(err.Error(), outputOptions)
		os.Exit(1)
	}
	return nil
}

// FormatType .
type FormatType string

// . .
const (
	JSON   FormatType = "JSON"
	YAML   FormatType = "YAML"
	STDOUT FormatType = "STDOUT"
)

func validateToken(lagoon string) {
	valid := graphql.VerifyTokenExpiry(lagoon)
	if valid == false {
		loginErr := loginToken()
		if loginErr != nil {
			fmt.Println("Unable to refresh token, you may need to run `lagoon login` first")
			os.Exit(1)
		}
	}
}
