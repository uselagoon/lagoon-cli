package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mglaman/lagoon/app"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmdProject app.LagoonProject
var cmdProjectName = ""

var rootCmd = &cobra.Command{
	Use:   "lagoon",
	Short: "Command line integration for Lagoon",
	Long:  `Lagoon CLI. Manage your Lagoon hosted projects.`,
}

// Execute the root command.
func Execute() {
	viper.AutomaticEnv()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.EnableCommandSorting = false

	rootCmd.PersistentFlags().StringVarP(&cmdProjectName, "project", "p", "", "The project name to interact with")

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
  {{rpad .Name .NamePadding }} {{.Short}}{{if .HasAvailableSubCommands}}{{range .Commands}}
    {{rpad .Name 9 }} {{.Short}}
{{- end}}{{end}}{{end}}{{end}}{{end}}

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
	rootCmd.AddCommand(kibanaCmd)
	rootCmd.AddCommand(projectCmd)

}

func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var configName = ".lagoon"

	// Search config in home directory with name ".lagoon" (without extension).
	// @todo see if we can grok the proper info from the cwd .lagoon.yml
	viper.AddConfigPath(home)
	viper.SetConfigName(configName)
	viper.SetDefault("lagoon_hostname", "ssh.lagoon.amazeeio.cloud")
	viper.SetDefault("lagoon_port", 32222)
	viper.SetDefault("lagoon_token", "")
	viper.SetDefault("lagoon_graphql", "https://api.lagoon.amazeeio.cloud/graphql")
	viper.SetDefault("lagoon_ui", "https://ui-lagoon-master.ch.amazee.io")
	viper.SetDefault("lagoon_kibana", "https://logs-db-ui-lagoon-master.ch.amazee.io/")
	err = viper.ReadInConfig()
	if err != nil {
		err = viper.WriteConfigAs(filepath.Join(home, configName+".yml"))
		if err != nil {
			panic(err)
		}

	}
}
