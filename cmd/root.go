package cmd

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/amazeeio/lagoon-cli/internal/helpers"
	"github.com/amazeeio/lagoon-cli/pkg/app"
	"github.com/amazeeio/lagoon-cli/pkg/graphql"
	"github.com/amazeeio/lagoon-cli/pkg/lagoon/environments"
	"github.com/amazeeio/lagoon-cli/pkg/lagoon/projects"
	"github.com/amazeeio/lagoon-cli/pkg/lagoon/users"
	"github.com/amazeeio/lagoon-cli/pkg/output"
	"github.com/amazeeio/lagoon-cli/pkg/updatecheck"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/viper"
)

var cmdProject app.LagoonProject
var cmdLagoon = ""
var forceAction bool
var cmdSSHKey = ""
var inputScanner = bufio.NewScanner(os.Stdin)
var versionFlag bool
var docsFlag bool
var updateInterval = time.Hour * 24 * 7 // One week interval between updates
var configName = ".lagoon"
var configExtension = ".yml"
var createConfig bool
var userPath string
var configFilePath string
var updateDocURL = "https://amazeeio.github.io/lagoon-cli"

var skipUpdateCheck bool

// version/build information
var (
	version string
	build   string
)

var rootCmd = &cobra.Command{
	Use:               "lagoon",
	Short:             "Command line integration for Lagoon",
	Long:              `Lagoon CLI. Manage your Lagoon hosted projects.`,
	DisableAutoGenTag: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if viper.GetBool("updateCheckDisable") == true {
			skipUpdateCheck = true
		}
		if skipUpdateCheck == false {
			// Using code from https://github.com/drud/ddev/
			updateFile := filepath.Join(userPath, ".lagoon.update")
			// Do periodic detection of whether an update is available for lagoon-cli users.
			timeToCheckForUpdates, err := updatecheck.IsUpdateNeeded(updateFile, updateInterval)
			if err != nil {
				output.RenderInfo(fmt.Sprintf("Could not perform update check %v", err), outputOptions)
			}
			if timeToCheckForUpdates && isInternetActive() {
				// Recreate the updatefile with current time so we won't do this again soon.
				err = updatecheck.ResetUpdateTime(updateFile)
				if err != nil {
					output.RenderInfo(fmt.Sprintf("Failed to update updatecheck file %s", updateFile), outputOptions)
				}
				updateNeeded, updateURL, err := updatecheck.AvailableUpdates("amazeeio", "lagoon-cli", version)
				if err != nil {
					output.RenderInfo("Could not check for updates. This is most often caused by a networking issue.", outputOptions)
					output.RenderError(err.Error(), outputOptions)
					return
				}
				if updateNeeded {
					output.RenderInfo(fmt.Sprintf("A new update is available! please visit %s to download the update.\nFor upgrade help see %s\n\nIf installed using brew, upgrade using `brew upgrade lagoon`\n", updateURL, updateDocURL), outputOptions)
				}
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if docsFlag {
			err := doc.GenMarkdownTree(cmd, "docs/commands")
			if err != nil {
				output.RenderError(err.Error(), outputOptions)
				os.Exit(1)
			}
		}
		if versionFlag {
			displayVersionInfo()
		}
		cmd.Help()
		os.Exit(1)
	},
}

// Execute the root command.
func Execute() {
	viper.AutomaticEnv()
	if err := rootCmd.Execute(); err != nil {
		output.RenderError(err.Error(), outputOptions)
		os.Exit(1)
	}
}

//IsInternetActive() checks to see if we have a viable
// internet connection. It just tries a quick DNS query.
// This requires that the named record be query-able.
func isInternetActive() bool {
	_, err := net.LookupHost("amazee.io")
	return err == nil
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cmdProjectName, "project", "p", "", "Specify a project to use")
	rootCmd.PersistentFlags().StringVarP(&cmdProjectEnvironment, "environment", "e", "", "Specify an environment to use")

	rootCmd.PersistentFlags().StringVarP(&cmdLagoon, "lagoon", "l", "", "The Lagoon instance to interact with")
	rootCmd.PersistentFlags().BoolVarP(&forceAction, "force", "", false, "Force yes on prompts (if supported)")
	rootCmd.PersistentFlags().StringVarP(&cmdSSHKey, "ssh-key", "i", "", "Specify path to a specific SSH key to use for lagoon authentication")

	// rootCmd.PersistentFlags().BoolVarP(&listAllProjects, "all-projects", "", false, "All projects (if supported)")
	rootCmd.PersistentFlags().BoolVarP(&outputOptions.Header, "no-header", "", false, "No header on table (if supported)")
	rootCmd.PersistentFlags().BoolVarP(&outputOptions.CSV, "output-csv", "", false, "Output as CSV (if supported)")
	rootCmd.PersistentFlags().BoolVarP(&outputOptions.JSON, "output-json", "", false, "Output as JSON (if supported)")
	rootCmd.PersistentFlags().BoolVarP(&outputOptions.Pretty, "pretty", "", false, "Make JSON pretty (if supported)")
	rootCmd.PersistentFlags().BoolVarP(&debugEnable, "debug", "", false, "Enable debugging output (if supported)")
	rootCmd.PersistentFlags().BoolVarP(&skipUpdateCheck, "skip-update-check", "", false, "Skip checking for updates")

	// get config-file from flag
	rootCmd.PersistentFlags().StringP("config-file", "", "", "Path to the config file to use (must be *.yml or *.yaml)")

	rootCmd.Flags().BoolVarP(&versionFlag, "version", "", false, "Version information")
	rootCmd.Flags().BoolVarP(&docsFlag, "docs", "", false, "Generate docs")

	rootCmd.Flags().MarkHidden("docs")

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
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(deployCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(kibanaCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(sshEnvCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(webCmd)
	rootCmd.AddCommand(importCmd)
	rootCmd.AddCommand(exportCmd)
	rootCmd.AddCommand(whoamiCmd)
}

// version/build information command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version information",
	Run: func(cmd *cobra.Command, args []string) {
		displayVersionInfo()
	},
}

func displayVersionInfo() {
	fmt.Println("Version:", version)
	fmt.Println("Build:", build)
	os.Exit(0)
}

func initConfig() {
	var err error
	// Find home directory.
	userPath, err = os.UserHomeDir()
	if err != nil {
		output.RenderError(fmt.Errorf("couldn't get $HOME: %v", err).Error(), outputOptions)
		os.Exit(1)
	}
	configFilePath = userPath

	// check if we are being given a path to a different config file
	err = helpers.GetLagoonConfigFile(&configFilePath, &configName, &configExtension, createConfig, rootCmd)
	if err != nil {
		output.RenderError(err.Error(), outputOptions)
		os.Exit(1)
	}

	// Search config in userPath directory with default name ".lagoon" (without extension).
	// @todo see if we can grok the proper info from the cwd .lagoon.yml
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configFilePath)
	viper.SetConfigName(configName)
	err = viper.ReadInConfig()
	if err != nil {
		// if we can't read the file cause it doesn't exist, then we should set the default configuration options and try create it
		viper.SetDefault("lagoons.amazeeio.hostname", "ssh.lagoon.amazeeio.cloud")
		viper.SetDefault("lagoons.amazeeio.port", 32222)
		viper.SetDefault("lagoons.amazeeio.token", "")
		viper.SetDefault("lagoons.amazeeio.graphql", "https://api.lagoon.amazeeio.cloud/graphql")
		viper.SetDefault("lagoons.amazeeio.ui", "https://ui-lagoon-master.ch.amazee.io")
		viper.SetDefault("lagoons.amazeeio.kibana", "https://logs-db-ui-lagoon-master.ch.amazee.io/")
		viper.SetDefault("default", "amazeeio")
		err = viper.WriteConfigAs(filepath.Join(configFilePath, configName+configExtension))
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
	}
	// get the lagoon context to use
	err = helpers.GetLagoonContext(&cmdLagoon, rootCmd)
	if err != nil {
		output.RenderError(err.Error(), outputOptions)
		os.Exit(1)
	}
	viper.Set("current", strings.TrimSpace(string(cmdLagoon))) // set the current lagoon to whatever we defined from config or as override in a flag

	err = viper.WriteConfig()
	if err != nil {
		output.RenderError(err.Error(), outputOptions)
		os.Exit(1)
	}

	// if the directory or repository you're in has a valid .lagoon.yml and docker-compose.yml with x-lagoon-project in it
	// we can use that inplaces where projects already exist so you don't have to type it out
	// and environments too
	if viper.GetBool("projectDirectoryCheckDisable") == false {
		cmdProject, _ = app.GetLocalProject()
	}
	if cmdProject.Name != "" && cmdProjectName == "" {
		cmdProjectName = cmdProject.Name
	}
	if cmdProject.Environment != "" && cmdProjectEnvironment == "" {
		cmdProjectEnvironment = cmdProject.Environment
	}

	// if !outputOptions.CSV && !outputOptions.JSON {
	// 	fmt.Println("Using Lagoon:", cmdLagoon)
	// }

}

func yesNo(message string) bool {
	if forceAction != true {
		prompt := promptui.Select{
			Label: message + "; Select[Yes/No]",
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

// global the clients
var eClient environments.Client
var uClient users.Client
var pClient projects.Client

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
			fmt.Println("Unable to refresh token, you may need to run `lagoon login` first, error was", loginErr.Error())
			os.Exit(1)
		}
	}
	// set up the clients
	var err error
	eClient, err = environments.New(debugEnable)
	if err != nil {
		output.RenderError(err.Error(), outputOptions)
		os.Exit(1)
	}
	uClient, err = users.New(debugEnable)
	if err != nil {
		output.RenderError(err.Error(), outputOptions)
		os.Exit(1)
	}
	pClient, err = projects.New(debugEnable)
	if err != nil {
		output.RenderError(err.Error(), outputOptions)
		os.Exit(1)
	}
	outputOptions.Debug = debugEnable
}

// validateTokenE does the same thing as validateToken, it just returns an
// error instead of exiting on error.
func validateTokenE(lagoon string) error {
	var err error
	if graphql.VerifyTokenExpiry(lagoon) {
		return nil // nothing to do
	}
	if err = loginToken(); err != nil {
		return fmt.Errorf("Couldn't refresh token, try `lagoon login`: %w", err)
	}
	// set up the clients
	eClient, err = environments.New(debugEnable)
	if err != nil {
		output.RenderError(err.Error(), outputOptions)
		return err
	}
	uClient, err = users.New(debugEnable)
	if err != nil {
		output.RenderError(err.Error(), outputOptions)
		return err
	}
	pClient, err = projects.New(debugEnable)
	if err != nil {
		output.RenderError(err.Error(), outputOptions)
		return err
	}
	outputOptions.Debug = debugEnable
	return nil
}
