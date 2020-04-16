package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"

	"github.com/amazeeio/lagoon-cli/internal/helpers"
	"github.com/amazeeio/lagoon-cli/internal/lagoon/client"
	"github.com/amazeeio/lagoon-cli/pkg/output"
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
	UI       string `json:"ui,omitempty"`
	Kibana   string `json:"kibana,omitempty"`
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
	Use:     "default",
	Aliases: []string{"d"},
	Short:   "Set the default Lagoon to use",
	Run: func(cmd *cobra.Command, args []string) {
		lagoonConfig := parseLagoonConfig(*cmd.Flags())
		if lagoonConfig.Lagoon == "" {
			fmt.Println("Not enough arguments")
			cmd.Help()
			os.Exit(1)
		}
		viper.Set("default", strings.TrimSpace(string(lagoonConfig.Lagoon)))
		err := viper.WriteConfigAs(filepath.Join(configFilePath, configName+configExtension))
		handleError(err)

		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"default-lagoon": lagoonConfig.Lagoon,
			},
		}
		output.RenderResult(resultData, outputOptions)
	},
}

var configLagoonsCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "View all configured Lagoon instances",
	RunE: func(cmd *cobra.Command, args []string) error {
		var data []output.Data
		for _, lagoon := range reflect.ValueOf(viper.Get("lagoons")).MapKeys() {
			var isDefault, isCurrent string
			if lagoon.String() == viper.Get("default") {
				isDefault = "(default)"
			}
			if lagoon.String() == viper.Get("current") {
				isCurrent = "(current)"
			}
			mapData := []string{
				helpers.ReturnNonEmptyString(fmt.Sprintf("%s%s%s", lagoon, isDefault, isCurrent)),
				helpers.ReturnNonEmptyString(viper.GetString("lagoons." + lagoon.String() + ".version")),
				helpers.ReturnNonEmptyString(viper.GetString("lagoons." + lagoon.String() + ".graphql")),
				helpers.ReturnNonEmptyString(viper.GetString("lagoons." + lagoon.String() + ".hostname")),
				helpers.ReturnNonEmptyString(viper.GetString("lagoons." + lagoon.String() + ".port")),
				helpers.ReturnNonEmptyString(viper.GetString("lagoons." + lagoon.String() + ".ui")),
				helpers.ReturnNonEmptyString(viper.GetString("lagoons." + lagoon.String() + ".Kibana")),
			}
			data = append(data, mapData)
		}
		sort.Slice(data, func(i, j int) bool {
			return data[i][0] < data[j][0]
		})
		output.RenderOutput(output.Table{
			Header: []string{
				"Name",
				"Version",
				"GraphQL",
				"SSH-Hostname",
				"SSH-Port",
				"UI-URL",
				"Kibana-URL",
			},
			Data: data,
		}, outputOptions)
		return nil
	},
}

var configAddCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add information about an additional Lagoon instance to use",
	RunE: func(cmd *cobra.Command, args []string) error {
		lagoonConfig := parseLagoonConfig(*cmd.Flags())
		if lagoonConfig.Lagoon == "" {
			return fmt.Errorf("Missing arguments: Lagoon name is not defined")
		}

		if lagoonConfig.Hostname != "" && lagoonConfig.Port != "" && lagoonConfig.GraphQL != "" {
			viper.Set("lagoons."+lagoonConfig.Lagoon+".hostname", lagoonConfig.Hostname)
			viper.Set("lagoons."+lagoonConfig.Lagoon+".port", lagoonConfig.Port)
			viper.Set("lagoons."+lagoonConfig.Lagoon+".graphql", lagoonConfig.GraphQL)
			if lagoonConfig.UI != "" {
				viper.Set("lagoons."+lagoonConfig.Lagoon+".ui", lagoonConfig.UI)
			}
			if lagoonConfig.Kibana != "" {
				viper.Set("lagoons."+lagoonConfig.Lagoon+".kibana", lagoonConfig.Kibana)
			}
			if lagoonConfig.Token != "" {
				viper.Set("lagoons."+lagoonConfig.Lagoon+".token", lagoonConfig.Token)
			}
			err := viper.WriteConfigAs(filepath.Join(configFilePath, configName+configExtension))
			if err != nil {
				return err
			}
			output.RenderOutput(output.Table{
				Header: []string{
					"Name",
					"GraphQL",
					"SSH-Hostname",
					"SSH-Port",
					"UI-URL",
					"Kibana-URL",
				},
				Data: []output.Data{
					[]string{

						lagoonConfig.Lagoon,
						lagoonConfig.GraphQL,
						lagoonConfig.Hostname,
						lagoonConfig.Port,
						lagoonConfig.UI,
						lagoonConfig.Kibana,
					},
				},
			}, outputOptions)
		} else {
			output.RenderError("Must have Hostname, Port, and GraphQL endpoint", outputOptions)
		}
		return nil
	},
}

var configDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	Short:   "Delete a Lagoon instance configuration",
	Run: func(cmd *cobra.Command, args []string) {
		lagoonConfig := parseLagoonConfig(*cmd.Flags())

		if lagoonConfig.Lagoon == "" {
			fmt.Println("Missing arguments: Lagoon name is not defined")
			cmd.Help()
			os.Exit(1)
		}
		if yesNo(fmt.Sprintf("You are attempting to delete config for lagoon '%s', are you sure?", lagoonConfig.Lagoon)) {
			err := unset(lagoonConfig.Lagoon)
			if err != nil {
				output.RenderError(err.Error(), outputOptions)
				os.Exit(1)
			}
		}
	},
}

var configFeatureSwitch = &cobra.Command{
	Use:     "feature",
	Aliases: []string{"f"},
	Short:   "Enable or disable CLI features",
	Run: func(cmd *cobra.Command, args []string) {
		switch updateCheck {
		case "true":
			viper.Set("updateCheckDisable", true)
		case "false":
			viper.Set("updateCheckDisable", false)
		}
		switch projectDirectoryCheck {
		case "true":
			viper.Set("projectDirectoryCheckDisable", true)
		case "false":
			viper.Set("projectDirectoryCheckDisable", false)
		}
		err := viper.WriteConfigAs(filepath.Join(configFilePath, configName+configExtension))
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
	},
}

var configGetCurrent = &cobra.Command{
	Use:     "current",
	Aliases: []string{"cur"},
	Short:   "Display the current Lagoon that commands would be executed against",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.GetString("current"))
	},
}

var configLagoonVersionCmd = &cobra.Command{
	Use:     "lagoon-version",
	Aliases: []string{"l"},
	Hidden:  false,
	Short:   "Checks the current Lagoon for its version and sets it in the config file",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
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
		lagoonVersion, err := helpers.GetLagoonAPIVersion(lc)
		if err != nil {
			return err
		}
		viper.Set("lagoons."+cmdLagoon+".version", lagoonVersion)
		if err = viper.WriteConfig(); err != nil {
			return fmt.Errorf("couldn't write config: %v", err)
		}
		fmt.Println(lagoonVersion)
		return nil
	},
}

var updateCheck string
var projectDirectoryCheck string

func init() {
	configCmd.AddCommand(configAddCmd)
	configCmd.AddCommand(configGetCurrent)
	configCmd.AddCommand(configDefaultCmd)
	configCmd.AddCommand(configDeleteCmd)
	configCmd.AddCommand(configFeatureSwitch)
	configCmd.AddCommand(configLagoonsCmd)
	configCmd.AddCommand(configLagoonVersionCmd)
	configAddCmd.Flags().StringVarP(&lagoonHostname, "hostname", "H", "", "Lagoon SSH hostname")
	configAddCmd.Flags().StringVarP(&lagoonPort, "port", "P", "", "Lagoon SSH port")
	configAddCmd.Flags().StringVarP(&lagoonGraphQL, "graphql", "g", "", "Lagoon GraphQL endpoint")
	configAddCmd.Flags().StringVarP(&lagoonToken, "token", "t", "", "Lagoon GraphQL token")
	configAddCmd.Flags().StringVarP(&lagoonUI, "ui", "u", "", "Lagoon UI location (https://ui-lagoon-master.ch.amazee.io)")
	configAddCmd.PersistentFlags().BoolVarP(&createConfig, "create-config", "", false, "Create the config file if it is non existent (to be used with --config-file)")
	configAddCmd.Flags().StringVarP(&lagoonKibana, "kibana", "k", "", "Lagoon Kibana URL (https://logs-db-ui-lagoon-master.ch.amazee.io)")
	configFeatureSwitch.Flags().StringVarP(&updateCheck, "disable-update-check", "", "", "Enable or disable checking of updates (true/false)")
	configFeatureSwitch.Flags().StringVarP(&projectDirectoryCheck, "disable-project-directory-check", "", "", "Enable or disable checking of local directory for Lagoon project (true/false)")
}
