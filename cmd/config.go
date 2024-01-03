package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/uselagoon/lagoon-cli/internal/lagoon"
	"github.com/uselagoon/lagoon-cli/internal/lagoon/client"
	"github.com/uselagoon/lagoon-cli/pkg/output"
	"gopkg.in/yaml.v3"
)

// LagoonConfigFlags .
type LagoonConfigFlags struct {
	Lagoon    string `json:"lagoon,omitempty"`
	Hostname  string `json:"hostname,omitempty"`
	Port      string `json:"port,omitempty"`
	GraphQL   string `json:"graphql,omitempty"`
	Token     string `json:"token,omitempty"`
	UI        string `json:"ui,omitempty"`
	Kibana    string `json:"kibana,omitempty"`
	SSHKey    string `json:"sshkey,omitempty"`
	SSHPortal bool   `json:"sshportal,omitempty"`
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
		lagoonCLIConfig.Default = strings.TrimSpace(string(lagoonConfig.Lagoon))
		contextExists := false
		for l := range lagoonCLIConfig.Lagoons {
			if l == lagoonCLIConfig.Current {
				contextExists = true
			}
		}
		if !contextExists {
			fmt.Println(fmt.Printf("Chosen context '%s' doesn't exist in config file", lagoonCLIConfig.Current))
			os.Exit(1)
		}
		err := writeLagoonConfig(&lagoonCLIConfig, filepath.Join(configFilePath, configName+configExtension))
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
		for l, lc := range lagoonCLIConfig.Lagoons {
			var isDefault, isCurrent string
			if l == lagoonCLIConfig.Default {
				isDefault = "(default)"
			}
			if l == lagoonCLIConfig.Current {
				isCurrent = "(current)"
			}
			mapData := []string{
				returnNonEmptyString(fmt.Sprintf("%s%s%s", l, isDefault, isCurrent)),
				returnNonEmptyString(lc.Version),
				returnNonEmptyString(lc.GraphQL),
				returnNonEmptyString(lc.HostName),
				returnNonEmptyString(lc.Port),
			}
			if fullConfigList {
				mapData = append(mapData, returnNonEmptyString(lc.UI))
				mapData = append(mapData, returnNonEmptyString(lc.Kibana))
			}
			mapData = append(mapData, returnNonEmptyString(lc.SSHKey))
			data = append(data, mapData)
		}
		sort.Slice(data, func(i, j int) bool {
			return data[i][0] < data[j][0]
		})
		tableHeader := []string{
			"Name",
			"Version",
			"GraphQL",
			"SSH-Hostname",
			"SSH-Port",
		}
		if fullConfigList {
			tableHeader = append(tableHeader, "UI-URL")
			tableHeader = append(tableHeader, "Kibana-URL")
		}
		tableHeader = append(tableHeader, "SSH-Key")
		output.RenderOutput(output.Table{
			Header: tableHeader,
			Data:   data,
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
			lc := lagoonCLIConfig.Lagoons[lagoonConfig.Lagoon]
			lc.HostName = lagoonConfig.Hostname
			lc.Port = lagoonConfig.Port
			lc.GraphQL = lagoonConfig.GraphQL
			if lagoonConfig.UI != "" {
				lc.UI = lagoonConfig.UI
			}
			if lagoonConfig.Kibana != "" {
				lc.Kibana = lagoonConfig.Kibana
			}
			if lagoonConfig.Token != "" {
				lc.Token = lagoonConfig.Token
			}
			if lagoonConfig.SSHKey != "" {
				lc.SSHKey = lagoonConfig.SSHKey
			}
			lagoonCLIConfig.Lagoons[lagoonConfig.Lagoon] = lc
			if err := writeLagoonConfig(&lagoonCLIConfig, filepath.Join(configFilePath, configName+configExtension)); err != nil {
				return fmt.Errorf("couldn't write config: %v", err)
			}
			output.RenderOutput(output.Table{
				Header: []string{
					"Name",
					"GraphQL",
					"SSH-Hostname",
					"SSH-Port",
					"UI-URL",
					"Kibana-URL",
					"SSH-Key",
				},
				Data: []output.Data{
					[]string{
						lagoonConfig.Lagoon,
						lagoonConfig.GraphQL,
						lagoonConfig.Hostname,
						lagoonConfig.Port,
						lagoonConfig.UI,
						lagoonConfig.Kibana,
						lagoonConfig.SSHKey,
					},
				},
			}, outputOptions)
		} else {
			return fmt.Errorf("Must have Hostname, Port, and GraphQL endpoint")
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
			err := removeConfig(lagoonConfig.Lagoon)
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
			lagoonCLIConfig.UpdateCheckDisable = true
		case "false":
			lagoonCLIConfig.UpdateCheckDisable = false
		}
		switch environmentFromDirectory {
		case "true":
			lagoonCLIConfig.EnvironmentFromDirectory = true
		case "false":
			lagoonCLIConfig.EnvironmentFromDirectory = false
		}
		if err := writeLagoonConfig(&lagoonCLIConfig, filepath.Join(configFilePath, configName+configExtension)); err != nil {
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
		fmt.Println(lagoonCLIConfig.Current)
	},
}

var configLagoonVersionCmd = &cobra.Command{
	Use:     "lagoon-version",
	Aliases: []string{"l"},
	Short:   "Checks the current Lagoon for its version and sets it in the config file",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		current := lagoonCLIConfig.Current
		lc := client.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIConfig.Lagoons[current].Token,
			lagoonCLIConfig.Lagoons[current].Version,
			lagoonCLIVersion,
			debug)
		lagoonVersion, err := lagoon.GetLagoonAPIVersion(context.TODO(), lc)
		if err != nil {
			return err
		}
		lagu := lagoonCLIConfig.Lagoons[current]
		lagu.Version = lagoonVersion.LagoonVersion
		lagoonCLIConfig.Lagoons[current] = lagu
		if err = writeLagoonConfig(&lagoonCLIConfig, filepath.Join(configFilePath, configName+configExtension)); err != nil {
			return fmt.Errorf("couldn't write config: %v", err)
		}
		fmt.Println(lagoonVersion.LagoonVersion)
		return nil
	},
}

var updateCheck string
var environmentFromDirectory string
var fullConfigList bool

func init() {
	configCmd.AddCommand(configAddCmd)
	configCmd.AddCommand(configGetCurrent)
	configCmd.AddCommand(configDefaultCmd)
	configCmd.AddCommand(configDeleteCmd)
	configCmd.AddCommand(configFeatureSwitch)
	configCmd.AddCommand(configLagoonsCmd)
	configCmd.AddCommand(configLagoonVersionCmd)
	configAddCmd.Flags().StringVarP(&lagoonHostname, "hostname", "H", "",
		"Lagoon SSH hostname")
	configAddCmd.Flags().StringVarP(&lagoonPort, "port", "P", "",
		"Lagoon SSH port")
	configAddCmd.Flags().StringVarP(&lagoonGraphQL, "graphql", "g", "",
		"Lagoon GraphQL endpoint")
	configAddCmd.Flags().StringVarP(&lagoonToken, "token", "t", "",
		"Lagoon GraphQL token")
	configAddCmd.Flags().StringVarP(&lagoonUI, "ui", "u", "",
		"Lagoon UI location (https://dashboard.amazeeio.cloud)")
	configAddCmd.PersistentFlags().BoolVarP(&createConfig, "create-config", "", false,
		"Create the config file if it is non existent (to be used with --config-file)")
	configAddCmd.Flags().StringVarP(&lagoonKibana, "kibana", "k", "",
		"Lagoon Kibana URL (https://logs.amazeeio.cloud)")
	configAddCmd.Flags().StringVarP(&lagoonSSHKey, "ssh-key", "", "",
		"SSH Key to use for this cluster for generating tokens")
	configLagoonsCmd.Flags().BoolVarP(&fullConfigList, "show-full", "", false,
		"Show full config output when listing Lagoon configurations")
	configFeatureSwitch.Flags().StringVarP(&updateCheck, "disable-update-check", "", "",
		"Enable or disable checking of updates (true/false)")
	configFeatureSwitch.Flags().StringVarP(&environmentFromDirectory, "enable-local-dir-check", "", "",
		"Enable or disable checking of local directory for Lagoon project (true/false)")
}

// readLagoonConfig reads the lagoon config from specified file.
func readLagoonConfig(lc *lagoon.Config, file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		// if there is no file found in the specified location, prompt the user to create it with the default
		// configuration to point to the amazeeio lagoon instance
		if yesNo(fmt.Sprintf("Config file '%s' does not exist, do you want to create it with defaults?", file)) {
			l := lagoon.Context{
				GraphQL:  "https://api.lagoon.amazeeio.cloud/graphql",
				HostName: "ssh.lagoon.amazeeio.cloud",
				Token:    "",
				Port:     "32222",
				UI:       "https://dashboard.amazeeio.cloud",
				Kibana:   "https://logs.amazeeio.cloud/",
			}
			lc.Lagoons = map[string]lagoon.Context{}
			lc.Lagoons["amazeeio"] = l
			lc.Default = "amazeeio"
			return writeLagoonConfig(lc, file)
		}
		return err
	}
	err = yaml.Unmarshal(data, &lc)
	if err != nil {
		return fmt.Errorf("unable to unmarshal config, yaml is likely invalid: %v", err)
	}
	for ln, l := range lc.Lagoons {
		if l.GraphQL == "" || l.HostName == "" || l.Port == "" {
			return fmt.Errorf("configured lagoon %s is missing required configuration for graphql, hostname, or port", ln)
		}
	}
	return nil

}

func analyze(file string) error {
	handle, err := os.Open(file)

	if err != nil {
		return err
	}
	defer handle.Close()
	return doSomething(handle)
}

func doSomething(handle io.Reader) error {
	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		// Do something with line
		d := scanner.Text()
		fmt.Println(d)
	}
	return nil
}

// functions to handle read/write of configuration file

// writeLagoonConfig writes the lagoon config to specified file.
func writeLagoonConfig(lc *lagoon.Config, file string) error {
	d, err := yaml.Marshal(&lc)
	if err != nil {
		return fmt.Errorf("unable to marshal config into valid yaml: %v", err)
	}
	err = os.WriteFile(file, d, 0777)
	if err != nil {
		return err
	}
	return nil
}

func setConfigDefaultVersion(lc *lagoon.Config, lagoon string, version string) error {
	if lc.Lagoons[lagoon].Version == "" {
		l := lc.Lagoons[lagoon]
		l.Version = version
		lc.Lagoons[lagoon] = l
		if err := writeLagoonConfig(&lagoonCLIConfig, filepath.Join(configFilePath, configName+configExtension)); err != nil {
			return fmt.Errorf("couldn't write config: %v", err)
		}
	}
	return nil
}

func removeConfig(key string) error {
	delete(lagoonCLIConfig.Lagoons, key)
	if err := writeLagoonConfig(&lagoonCLIConfig, filepath.Join(configFilePath, configName+configExtension)); err != nil {
		output.RenderError(err.Error(), outputOptions)
		os.Exit(1)
	}
	return nil
}
