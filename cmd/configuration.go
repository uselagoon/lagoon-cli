package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/uselagoon/lagoon-cli/internal/lagoon"
	"github.com/uselagoon/lagoon-cli/pkg/output"
	config "github.com/uselagoon/machinery/utils/config"
	"github.com/uselagoon/machinery/utils/discovery"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v3"
)

// configv2 related globals
var lConfig *config.Config
var lContext *config.Context
var lUser *config.User
var configFeaturePrefix = "cli"
var cliFeatures = []string{
	"environment-from-directory",  // feature that enables the sourcing of project/environment from within a git repository
	"disable-update-check",        // feature that disables the automatic update check
	"ssh-token",                   // feature that forces the cli to use ssh token base authentication instead of keycloak based
	"no-strict-host-key-checking", // feature that enables or disables strict host key checking, default is enabled
}

var configCmd = &cobra.Command{
	Use:     "configuration",
	Aliases: []string{"config", "conf", "c"},
	Short:   "Manage or view the contexts and users for interacting with Lagoon",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}

var configListUsersCmd = &cobra.Command{
	Use:     "list-users",
	Aliases: []string{"lu"},
	Short:   "View all configured Lagoon context users",
	RunE: func(cmd *cobra.Command, args []string) error {
		data := []output.Data{}
		for _, user := range lConfig.Users {
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%v", user.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", user.UserConfig.SSHKey)),
			})
		}
		r := output.RenderOutput(output.Table{
			Header: []string{
				"Name",
				"SSH-Key",
			},
			Data: data,
		}, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		return nil
	},
}

var configListContextsCmd = &cobra.Command{
	Use:     "list-contexts",
	Aliases: []string{"lc"},
	Short:   "View all configured Lagoon contexts",
	RunE: func(cmd *cobra.Command, args []string) error {
		data := []output.Data{}
		featurePrefix := fmt.Sprintf("%s-", configFeaturePrefix)
		for _, con := range lConfig.Contexts {
			defa := false
			if con.Name == lConfig.DefaultContext {
				defa = true
			}
			selected := ""
			if con.Name == cmdLagoon {
				selected = "(selected)"
			}
			contextFeatures := map[string]bool{}
			for f, b := range con.ContextConfig.Features {
				// only add cli prefixed features in the cli
				if strings.Contains(f, featurePrefix) {
					// could include if it is context or global sourced here
					contextFeatures[strings.TrimPrefix(f, featurePrefix)] = b
				}
			}
			for f, b := range lConfig.Features {
				// only add cli prefixed features in the cli
				if strings.Contains(f, featurePrefix) {
					if _, ok := con.ContextConfig.Features[f]; !ok {
						// could include if it is context or global sourced here
						contextFeatures[strings.TrimPrefix(f, featurePrefix)] = b
					}
				}
			}
			features := []string{}
			// transform the features into a string slice for printing
			for f, b := range contextFeatures {
				features = append(features, fmt.Sprintf("%s=%t", f, b))
			}
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%v%s", con.Name, selected)),
				returnNonEmptyString(fmt.Sprintf("%v", defa)),
				returnNonEmptyString(fmt.Sprintf("%v", con.User)),
				returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.APIHostname)),
				returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.TokenHost)),
				returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.TokenPort)),
				returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.AuthenticationEndpoint)),
				returnNonEmptyString(fmt.Sprintf("%v", strings.Join(features, ","))),
			})
		}
		r := output.RenderOutput(output.Table{
			Header: []string{
				"Name",
				"Default",
				"User",
				"API-Hostname",
				"Token-Hostname",
				"Token-Port",
				"Authentication-Hostname",
				"Features",
			},
			Data: data,
		}, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		return nil
	},
}

var configAddUserCmd = &cobra.Command{
	Use:     "add-user",
	Aliases: []string{"au"},
	Short:   "Add a new Lagoon context user",
	RunE: func(cmd *cobra.Command, args []string) error {
		uName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		uSSHKey, err := cmd.Flags().GetString("ssh-key")
		if err != nil {
			return err
		}
		// create the requested user
		uConfig := config.UserConfig{
			Name:  uName,
			Grant: &oauth2.Token{},
		}
		// visit the flags and check for any defined flags to set
		cmd.Flags().Visit(func(f *pflag.Flag) {
			if f.Name == "ssh-key" {
				uConfig.SSHKey = uSSHKey
			}
		})
		err = lConfig.NewUser(uConfig)
		if err != nil {
			return err
		}
		data := []output.Data{}
		for _, user := range lConfig.Users {
			if user.Name == uName {
				data = append(data, []string{
					returnNonEmptyString(fmt.Sprintf("%v", user.Name)),
					returnNonEmptyString(fmt.Sprintf("%v", user.UserConfig.SSHKey)),
				})
			}
		}
		r := output.RenderOutput(output.Table{
			Header: []string{
				"Name",
				"SSH-Key",
			},
			Data: data,
		}, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		// then save it
		return lConfig.WriteConfig()
	},
}

var configUpdateUserCmd = &cobra.Command{
	Use:     "update-user",
	Aliases: []string{"uu"},
	Short:   "Update a Lagoon context user",
	RunE: func(cmd *cobra.Command, args []string) error {
		uName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		uSSHKey, err := cmd.Flags().GetString("ssh-key")
		if err != nil {
			return err
		}
		// get the requested user
		u, err := lConfig.GetUser(uName)
		if err != nil {
			return err
		}
		// visit the flags and check for any defined flags to set
		cmd.Flags().Visit(func(f *pflag.Flag) {
			if f.Name == "ssh-key" {
				u.UserConfig.SSHKey = uSSHKey
			}
		})
		err = lConfig.UpdateUser(u.UserConfig)
		if err != nil {
			return err
		}
		data := []output.Data{}
		for _, user := range lConfig.Users {
			if user.Name == uName {
				data = append(data, []string{
					returnNonEmptyString(fmt.Sprintf("%v", user.Name)),
					returnNonEmptyString(fmt.Sprintf("%v", user.UserConfig.SSHKey)),
				})
			}
		}
		r := output.RenderOutput(output.Table{
			Header: []string{
				"Name",
				"SSH-Key",
			},
			Data: data,
		}, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		// then save it
		return lConfig.WriteConfig()
	},
}

var configAddContextCmd = &cobra.Command{
	Use:     "add-context",
	Aliases: []string{"ac"},
	Short:   "Add a new Lagoon context",
	RunE: func(cmd *cobra.Command, args []string) error {
		cName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		cUser, err := cmd.Flags().GetString("user")
		if err != nil {
			return err
		}
		cAPIHost, err := cmd.Flags().GetString("api-hostname")
		if err != nil {
			return err
		}
		cTokenHost, err := cmd.Flags().GetString("token-hostname")
		if err != nil {
			return err
		}
		cTokenPort, err := cmd.Flags().GetInt("token-port")
		if err != nil {
			return err
		}
		cAuthHost, err := cmd.Flags().GetString("authentication-hostname")
		if err != nil {
			return err
		}
		cUIHost, err := cmd.Flags().GetString("ui-hostname")
		if err != nil {
			return err
		}
		cWebhookHost, err := cmd.Flags().GetString("webhook-hostname")
		if err != nil {
			return err
		}
		// create the requested context
		cConfig := config.ContextConfig{
			Name: cName,
		}
		// visit the flags and check for any defined flags to set
		cmd.Flags().Visit(func(f *pflag.Flag) {
			if f.Name == "api-hostname" {
				// strip /graphql from the host
				cConfig.APIHostname = strings.TrimSuffix(cAPIHost, "/graphql")
			}
			if f.Name == "token-hostname" {
				cConfig.TokenHost = cTokenHost
			}
			if f.Name == "token-port" {
				cConfig.TokenPort = cTokenPort
			}
			if f.Name == "authentication-hostname" {
				cConfig.AuthenticationEndpoint = cAuthHost
			}
			if f.Name == "ui-hostname" {
				cConfig.UIHostname = cUIHost
			}
			if f.Name == "webhook-hostname" {
				cConfig.WebhookEndpoint = cWebhookHost
			}
		})
		apiDiscoverEndpoint, err := discovery.Discover(cConfig.APIHostname)
		if err != nil {
			// throw error here, even if the endpoint is unreachable
			return err
		}
		if apiDiscoverEndpoint != nil {
			outputOptions.MessagePrefix = fmt.Sprintf("context %s", cName)
			// try and extract information from the endpoint
			if apiDiscoverEndpoint.AuthorizationEndpoint != "" && cConfig.AuthenticationEndpoint == "" {
				cConfig.AuthenticationEndpoint = apiDiscoverEndpoint.AuthorizationEndpoint
				output.RenderInfo(fmt.Sprintf("using authentication endpoint '%s' from api discovery for context\n", cConfig.AuthenticationEndpoint), outputOptions)
			}
			if apiDiscoverEndpoint.SSHTokenExchange.TokenHost != "" && cConfig.TokenHost == "" {
				cConfig.TokenHost = apiDiscoverEndpoint.SSHTokenExchange.TokenHost
				output.RenderInfo(fmt.Sprintf("using token host endpoint '%s' from api discovery for context\n", cConfig.TokenHost), outputOptions)
			}
			if apiDiscoverEndpoint.SSHTokenExchange.TokenPort != 0 && cConfig.TokenPort == 0 {
				cConfig.TokenPort = apiDiscoverEndpoint.SSHTokenExchange.TokenPort
				output.RenderInfo(fmt.Sprintf("using token host port '%d' from api discovery for context\n", cConfig.TokenPort), outputOptions)
			}
			if apiDiscoverEndpoint.UIHostname != "" && cConfig.UIHostname == "" {
				cConfig.UIHostname = apiDiscoverEndpoint.UIHostname
				output.RenderInfo(fmt.Sprintf("using ui host '%s' from api discovery for context\n", cConfig.UIHostname), outputOptions)
			}
			if apiDiscoverEndpoint.WebhookEndpoint != "" && cConfig.WebhookEndpoint == "" {
				cConfig.WebhookEndpoint = apiDiscoverEndpoint.WebhookEndpoint
				output.RenderInfo(fmt.Sprintf("using webhook endpoint '%s' from api discovery for context\n", cConfig.WebhookEndpoint), outputOptions)
			}
		}
		err = lConfig.NewContext(cConfig, cUser)
		if err != nil {
			return err
		}
		data := []output.Data{}
		for _, con := range lConfig.Contexts {
			defa := false
			if con.Name == lConfig.DefaultContext {
				defa = true
			}
			if con.Name == cName {
				data = append(data, []string{
					returnNonEmptyString(fmt.Sprintf("%v", con.Name)),
					returnNonEmptyString(fmt.Sprintf("%v", defa)),
					returnNonEmptyString(fmt.Sprintf("%v", con.User)),
					returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.APIHostname)),
					returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.TokenHost)),
					returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.TokenPort)),
					returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.AuthenticationEndpoint)),
				})
			}
		}
		r := output.RenderOutput(output.Table{
			Header: []string{
				"Name",
				"Default",
				"User",
				"API-Hostname",
				"Token-Hostname",
				"Token-Port",
				"Authentication-Hostname",
			},
			Data: data,
		}, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		// then save it
		return lConfig.WriteConfig()
	},
}

var configUpdateContextCmd = &cobra.Command{
	Use:     "update-context",
	Aliases: []string{"uc"},
	Short:   "Update a Lagoon context",
	RunE: func(cmd *cobra.Command, args []string) error {
		cName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		cUser, err := cmd.Flags().GetString("user")
		if err != nil {
			return err
		}
		cAPIHost, err := cmd.Flags().GetString("api-hostname")
		if err != nil {
			return err
		}
		cTokenHost, err := cmd.Flags().GetString("token-hostname")
		if err != nil {
			return err
		}
		cTokenPort, err := cmd.Flags().GetInt("token-port")
		if err != nil {
			return err
		}
		cAuthHost, err := cmd.Flags().GetString("authentication-hostname")
		if err != nil {
			return err
		}
		cUIHost, err := cmd.Flags().GetString("ui-hostname")
		if err != nil {
			return err
		}
		cWebhookHost, err := cmd.Flags().GetString("webhook-hostname")
		if err != nil {
			return err
		}
		// get the requested context
		cConfig, err := lConfig.GetContext(cName)
		if err != nil {
			return err
		}
		// visit the flags and check for any defined flags to set
		cmd.Flags().Visit(func(f *pflag.Flag) {
			// these will override the context ones if they are defined, otherwise the existing
			// fields will remain untouched
			if f.Name == "api-hostname" {
				cConfig.ContextConfig.APIHostname = cAPIHost
			}
			if f.Name == "token-hostname" {
				cConfig.ContextConfig.TokenHost = cTokenHost
			}
			if f.Name == "token-port" {
				cConfig.ContextConfig.TokenPort = cTokenPort
			}
			if f.Name == "authentication-hostname" {
				cConfig.ContextConfig.AuthenticationEndpoint = cAuthHost
			}
			if f.Name == "ui-hostname" {
				cConfig.ContextConfig.UIHostname = cUIHost
			}
			if f.Name == "webhook-hostname" {
				cConfig.ContextConfig.WebhookEndpoint = cWebhookHost
			}
			if f.Name == "user" {
				cConfig.User = cUser
			}
		})
		// update the context within the configuration
		err = lConfig.UpdateContext(cConfig.ContextConfig, cConfig.User)
		if err != nil {
			return err
		}
		data := []output.Data{}
		for _, con := range lConfig.Contexts {
			defa := false
			if con.Name == lConfig.DefaultContext {
				defa = true
			}
			if con.Name == cName {
				data = append(data, []string{
					returnNonEmptyString(fmt.Sprintf("%v", con.Name)),
					returnNonEmptyString(fmt.Sprintf("%v", defa)),
					returnNonEmptyString(fmt.Sprintf("%v", con.User)),
					returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.APIHostname)),
					returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.TokenHost)),
					returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.TokenPort)),
					returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.AuthenticationEndpoint)),
				})
			}
		}
		r := output.RenderOutput(output.Table{
			Header: []string{
				"Name",
				"Default",
				"User",
				"API-Hostname",
				"Token-Hostname",
				"Token-Port",
				"Authentication-Hostname",
			},
			Data: data,
		}, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		// then save it
		return lConfig.WriteConfig()
	},
}

var configGetConfigPathCmd = &cobra.Command{
	Use:     "config-path",
	Aliases: []string{"cp"},
	Short:   "Get the path of where the config file lives",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := config.Config{}
		path, err := c.GetConfigLocation()
		if err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "%s", path)
		return nil
	},
}

var configSetDefaultCmd = &cobra.Command{
	Use:     "default-context",
	Aliases: []string{"dc"},
	Short:   "Change which context is the default",
	RunE: func(cmd *cobra.Command, args []string) error {
		cName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		if err = lConfig.SetDefaultContext(cName); err != nil {
			return err
		}
		data := []output.Data{}
		for _, con := range lConfig.Contexts {
			defa := false
			if con.Name == lConfig.DefaultContext {
				defa = true
			}
			selected := ""
			if con.Name == cmdLagoon {
				selected = " (selected)"
			}
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%v%s", con.Name, selected)),
				returnNonEmptyString(fmt.Sprintf("%v", defa)),
				returnNonEmptyString(fmt.Sprintf("%v", con.User)),
				returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.APIHostname)),
				returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.TokenHost)),
				returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.TokenPort)),
				returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.AuthenticationEndpoint)),
			})
		}
		r := output.RenderOutput(output.Table{
			Header: []string{
				"Name",
				"Default",
				"User",
				"API-Hostname",
				"Token-Hostname",
				"Token-Port",
				"Authentication-Hostname",
			},
			Data: data,
		}, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		return lConfig.WriteConfig()
	},
}

var configConvertLegacyConfig = &cobra.Command{
	Use:     "convert-config",
	Aliases: []string{"convert"},
	Short:   "Convert legacy .lagoon.yml config to the new configuration format",
	Long: `Convert legacy .lagoon.yml config to the new configuration format.
This will prompt you to provide any required information if it is missing from your legacy configuration.
Running this command initially will run in dry-run mode, if you're happy with the result you can run it again
with the --write-config flag to save the new configuration.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		writeConfig, err := cmd.Flags().GetBool("write-config")
		if err != nil {
			return err
		}
		return convertConfig(writeConfig)
	},
}

var configSetFeatureStatus = &cobra.Command{
	Use:     "feature",
	Aliases: []string{"feat", "f"},
	Short:   "Enable or disable a feature for all contexts or a specific context",
	RunE: func(cmd *cobra.Command, args []string) error {
		cContext, err := cmd.Flags().GetString("context")
		if err != nil {
			return err
		}
		cFeature, err := cmd.Flags().GetString("feature")
		if err != nil {
			return err
		}
		cFeatureState, err := cmd.Flags().GetBool("state")
		if err != nil {
			return err
		}
		contextFeature := false
		cmd.Flags().Visit(func(f *pflag.Flag) {
			if f.Name == "context" {
				contextFeature = true
			}
		})
		if cFeature == "" {
			return fmt.Errorf("feature name must be provided")
		}
		if !slices.Contains(cliFeatures, cFeature) {
			return fmt.Errorf("feature %s is not one of the supported features: %s", cFeature, cliFeatures)
		}
		if cContext != "" && contextFeature {
			err := lConfig.SetContextFeature(cContext, configFeaturePrefix, cFeature, cFeatureState)
			if err != nil {
				return fmt.Errorf("unable to set context feature: %v", err)
			}
			return lConfig.WriteConfig()
		} else if cContext == "" && contextFeature {
			return fmt.Errorf("context name must be provided")
		}
		err = lConfig.SetGlobalFeature(configFeaturePrefix, cFeature, cFeatureState)
		if err != nil {
			return fmt.Errorf("unable to set global feature: %v", err)
		}
		return lConfig.WriteConfig()
	},
}

func init() {
	configCmd.AddCommand(configListContextsCmd)
	configCmd.AddCommand(configListUsersCmd)
	configCmd.AddCommand(configGetConfigPathCmd)

	configCmd.AddCommand(configConvertLegacyConfig)
	configConvertLegacyConfig.Flags().Bool("write-config", false, "Whether the config should be written to the config file or not")

	configCmd.AddCommand(configSetDefaultCmd)
	configSetDefaultCmd.Flags().String("name", "", "The name of the context to be default")

	configCmd.AddCommand(configAddContextCmd)
	configAddContextCmd.Flags().String("name", "", "The name to reference this context as")
	configAddContextCmd.Flags().String("api-hostname", "", "Lagoon API hostname (eg: https://api.lagoon.sh)")
	configAddContextCmd.Flags().String("token-hostname", "", "Lagoon Token endpoint hostname (eg: token.lagoon.sh)")
	configAddContextCmd.Flags().Int("token-port", 0, "Lagoon Token endpoint port (eg: 22)")
	configAddContextCmd.Flags().String("authentication-hostname", "", "Lagoon authentication hostname (eg: https://keycloak.lagoon.sh)")
	configAddContextCmd.Flags().String("user", "", "The user to associate to this context")
	configAddContextCmd.Flags().String("ui-hostname", "", "Lagoon UI hostname (eg: https://ui.lagoon.sh)")
	configAddContextCmd.Flags().String("webhook-hostname", "", "Lagoon webhook hostname (eg: https://webhook.lagoon.sh)")

	configCmd.AddCommand(configAddUserCmd)
	configAddUserCmd.Flags().String("name", "", "The name to reference this user as")
	configAddUserCmd.Flags().String("ssh-key", "", "The full path to this users ssh-key")

	configCmd.AddCommand(configUpdateUserCmd)
	configUpdateUserCmd.Flags().String("name", "", "The name to reference this user as")
	configUpdateUserCmd.Flags().String("ssh-key", "", "The full path to this users ssh-key")

	configCmd.AddCommand(configUpdateContextCmd)
	configUpdateContextCmd.Flags().String("name", "", "The name to reference this context as")
	configUpdateContextCmd.Flags().String("api-hostname", "", "Lagoon API hostname (eg: https://api.lagoon.sh)")
	configUpdateContextCmd.Flags().String("token-hostname", "", "Lagoon Token endpoint hostname (eg: token.lagoon.sh)")
	configUpdateContextCmd.Flags().Int("token-port", 0, "Lagoon Token endpoint port (eg: 22)")
	configUpdateContextCmd.Flags().String("authentication-hostname", "", "Lagoon authentication hostname (eg: https://keycloak.lagoon.sh)")
	configUpdateContextCmd.Flags().String("user", "", "The user to associate to this context")
	configUpdateContextCmd.Flags().String("ui-hostname", "", "Lagoon UI hostname (eg: https://ui.lagoon.sh)")
	configUpdateContextCmd.Flags().String("webhook-hostname", "", "Lagoon webhook hostname (eg: https://webhook.lagoon.sh)")

	configCmd.AddCommand(configSetFeatureStatus)
	configSetFeatureStatus.Flags().Bool("state", false, "The state of the feature (--state=true or --state=false)")
	configSetFeatureStatus.Flags().String("context", "", "If provided the feature will be enabled for this context, otherwise globally")
	configSetFeatureStatus.Flags().String("feature", "", fmt.Sprintf("The name of the feature to enable or disable [%s]", strings.Join(cliFeatures, ",")))
}

func getLegacyConfigFile(configPath *string, configName *string, configExtension *string, cmd *cobra.Command) error {
	// check if we have an envvar or flag to define our confg file
	var configFilePath string
	configFilePath, err := cmd.Flags().GetString("config-file")
	if err != nil {
		return fmt.Errorf("error reading flag `config-file`: %v", err)
	}
	if configFilePath == "" {
		if lagoonConfigEnvar, ok := os.LookupEnv("LAGOONCONFIG"); ok {
			configFilePath = lagoonConfigEnvar
		}
		// prefer LAGOON_CONFIG_FILE
		if lagoonConfigEnvar, ok := os.LookupEnv("LAGOON_CONFIG_FILE"); ok {
			configFilePath = lagoonConfigEnvar
		}
	}
	if configFilePath != "" {
		if fileExists(configFilePath) || createConfig {
			*configPath = filepath.Dir(configFilePath)
			*configExtension = filepath.Ext(configFilePath)
			*configName = strings.TrimSuffix(filepath.Base(configFilePath), *configExtension)
			return nil
		}
		return fmt.Errorf("%s/%s File doesn't exist", *configPath, configFilePath)
	}
	// no config file found
	return nil
}

func readLegacyConfig() ([]byte, error) {
	// check for the legacy config file
	userPath, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	configFilePath = userPath
	err = getLegacyConfigFile(&configFilePath, &configName, &configExtension, rootCmd)
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(filepath.Join(configFilePath, configName+configExtension))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func convertConfig(writeConfig bool) error {
	data, err := readLegacyConfig()
	if err != nil {
		return err
	}
	lc := lagoon.Config{}
	err = yaml.Unmarshal(data, &lc)
	if err != nil {
		return fmt.Errorf("unable to unmarshal config, yaml is likely invalid: %v", err)
	}
	// convert the legacy config into the new user/context config file
	cc := config.Config{}
	for n, l := range lc.Lagoons {
		apiHostname := strings.TrimSuffix(l.GraphQL, "/graphql")
		// try and extract the auth endpoint from the discovery endpoint of the api
		var authEndpoint string
		apiDiscoverEndpoint, err := discovery.Discover(apiHostname)
		if err != nil {
			// if the endpoint isn't reachable during a conversion either because its old or incorrect
			// ignore the error and prompt the user to enter missing values manually
			if !strings.Contains(err.Error(), "no route") && !strings.Contains(err.Error(), "connection refused") {
				return err
			}
		}
		tokenHost := l.HostName
		tokenPort := l.Port
		uiHostname := l.UI
		var webhookEndpoint string
		iTokenPort, _ := strconv.Atoi(tokenPort)
		if apiDiscoverEndpoint != nil {
			// if the discovery endpoint returns data, check for the auth endpoint
			outputOptions.MessagePrefix = fmt.Sprintf("context %s", n)
			if apiDiscoverEndpoint.AuthorizationEndpoint != "" {
				authEndpoint = apiDiscoverEndpoint.AuthorizationEndpoint
				output.RenderInfo(fmt.Sprintf("using authentication endpoint '%s' from api discovery for context\n", authEndpoint), outputOptions)
			}
			if apiDiscoverEndpoint.SSHTokenExchange.TokenHost != "" {
				tokenHost = apiDiscoverEndpoint.SSHTokenExchange.TokenHost
				output.RenderInfo(fmt.Sprintf("using token host endpoint '%s' from api discovery for context\n", tokenHost), outputOptions)
			}
			if apiDiscoverEndpoint.SSHTokenExchange.TokenPort != 0 {
				iTokenPort = apiDiscoverEndpoint.SSHTokenExchange.TokenPort
				output.RenderInfo(fmt.Sprintf("using token host port '%d' from api discovery for context\n", iTokenPort), outputOptions)
			}
			if apiDiscoverEndpoint.UIHostname != "" {
				uiHostname = apiDiscoverEndpoint.UIHostname
				output.RenderInfo(fmt.Sprintf("using ui host '%s' from api discovery for context\n", uiHostname), outputOptions)
			}
			if apiDiscoverEndpoint.WebhookEndpoint != "" {
				webhookEndpoint = apiDiscoverEndpoint.WebhookEndpoint
				output.RenderInfo(fmt.Sprintf("using webhook endpoint '%s' from api discovery for context\n", webhookEndpoint), outputOptions)
			}
		}
		if authEndpoint == "" {
			prompt := promptui.Prompt{
				Label:     fmt.Sprintf("Enter authentication endpoint for context %s / %s", n, apiHostname),
				Default:   "https://keycloak.example.com/auth",
				AllowEdit: true,
			}
			authEndpoint, err = prompt.Run()
			if err != nil {
				return err
			}
		}
		uConfig := config.UserConfig{
			Name: n,
			Grant: &oauth2.Token{
				AccessToken: l.Token,
			},
			SSHKey: l.SSHKey,
		}
		if l.PublicKeyIdentities != nil {
			uConfig.PublicKeyIdentities = l.PublicKeyIdentities
		}
		cConfig := config.ContextConfig{
			Name:                   n,
			APIHostname:            apiHostname,
			TokenHost:              tokenHost,
			TokenPort:              iTokenPort,
			AuthenticationEndpoint: authEndpoint,
			Version:                l.Version,
			UIHostname:             uiHostname,
			WebhookEndpoint:        webhookEndpoint,
		}
		err = cc.NewUser(uConfig)
		if err != nil {
			return err
		}
		err = cc.NewContext(cConfig, uConfig.Name)
		if err != nil {
			return err
		}
	}
	// convert old global feature flags if defined to replacement feature flags
	if lc.EnvironmentFromDirectory {
		err := cc.SetGlobalFeature("cli", "environment-from-directory", true)
		if err != nil {
			return fmt.Errorf("unable to set global feature: %v", err)
		}
	}
	if lc.UpdateCheckDisable {
		err := cc.SetGlobalFeature("cli", "disable-update-check", true)
		if err != nil {
			return fmt.Errorf("unable to set global feature: %v", err)
		}
	}
	if lc.StrictHostKeyChecking == "no" {
		err := cc.SetGlobalFeature("cli", "no-strict-host-key-checking", true)
		if err != nil {
			return fmt.Errorf("unable to set global feature: %v", err)
		}
	}
	err = cc.SetDefaultContext(lc.Default)
	if err != nil {
		return fmt.Errorf("unable to set default context: %v", err)
	}
	if writeConfig {
		if err := cc.WriteConfig(); err != nil {
			return err
		}
	} else {
		cb, _ := yaml.Marshal(cc)
		fmt.Println(string(cb))
		fmt.Println("configuration file not written, to save converted config run this again with the flag --write-config")
	}
	return nil
}

// helper function for creating an initial configuration with prompts if no legacy or new config is detected
func createInitialConfig() error {
	lConfig, _ = config.LoadConfig(true)
	uPrompt := promptui.Prompt{
		Label:     "Enter a user name",
		Default:   "user",
		AllowEdit: true,
	}
	userName, err := uPrompt.Run()
	if err != nil {
		return fmt.Errorf("unable to create configuration: %s", err.Error())
	}
	err = lConfig.NewUser(config.UserConfig{
		Name: userName,
	})
	if err != nil {
		return fmt.Errorf("unable to create new user: %s", err.Error())
	}
	cPrompt := promptui.Prompt{
		Label:     "Enter a context name",
		Default:   "lagoon",
		AllowEdit: true,
	}
	cName, err := cPrompt.Run()
	if err != nil {
		return fmt.Errorf("unable to create configuration: %s", err.Error())
	}
	prompt := promptui.Prompt{
		Label:     "Enter API hostname, omit the /graphql path",
		Default:   "https://api.example.com",
		AllowEdit: true,
	}
	apiHostname, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("unable to create configuration: %s", err.Error())
	}
	var authEndpoint, tokenHost, uiHostname, webhookEndpoint string
	var iTokenPort int
	apiHostname = strings.TrimSuffix(apiHostname, "/graphql") // remove the suffix just incase
	apiDiscoverEndpoint, err := discovery.Discover(apiHostname)
	if err != nil {
		// throw error here, even if the endpoint is unreachable
		return err
	}
	if apiDiscoverEndpoint != nil {
		// try and extract information from the endpoint
		outputOptions.MessagePrefix = fmt.Sprintf("context %s", cName)
		if apiDiscoverEndpoint.AuthorizationEndpoint != "" {
			authEndpoint = apiDiscoverEndpoint.AuthorizationEndpoint
			output.RenderInfo(fmt.Sprintf("using authentication endpoint '%s' from api discovery for context\n", authEndpoint), outputOptions)
		}
		if apiDiscoverEndpoint.SSHTokenExchange.TokenHost != "" {
			tokenHost = apiDiscoverEndpoint.SSHTokenExchange.TokenHost
			output.RenderInfo(fmt.Sprintf("using token host endpoint '%s' from api discovery for context\n", tokenHost), outputOptions)
		}
		if apiDiscoverEndpoint.SSHTokenExchange.TokenPort != 0 {
			iTokenPort = apiDiscoverEndpoint.SSHTokenExchange.TokenPort
			output.RenderInfo(fmt.Sprintf("using token host port '%d' from api discovery for context\n", iTokenPort), outputOptions)
		}
		if apiDiscoverEndpoint.UIHostname != "" {
			uiHostname = apiDiscoverEndpoint.UIHostname
			output.RenderInfo(fmt.Sprintf("using ui host '%s' from api discovery for context\n", uiHostname), outputOptions)
		}
		if apiDiscoverEndpoint.WebhookEndpoint != "" {
			webhookEndpoint = apiDiscoverEndpoint.WebhookEndpoint
			output.RenderInfo(fmt.Sprintf("using webhook endpoint '%s' from api discovery for context\n", webhookEndpoint), outputOptions)
		}
	}
	// if the endpoint has no values, prompt for them
	if authEndpoint == "" {
		prompt2 := promptui.Prompt{
			Label:     fmt.Sprintf("Enter authentication endpoint for context %s / %s", cName, apiHostname),
			Default:   "https://keycloak.example.com/auth",
			AllowEdit: true,
		}
		authEndpoint, err = prompt2.Run()
		if err != nil {
			return err
		}
	}
	if tokenHost == "" {
		prompt3 := promptui.Prompt{
			Label:     fmt.Sprintf("Enter SSH token endpoint for context %s / %s", cName, apiHostname),
			Default:   "token.example.com",
			AllowEdit: true,
		}
		tokenHost, err = prompt3.Run()
		if err != nil {
			return fmt.Errorf("unable to create configuration: %s", err.Error())
		}
	}
	if iTokenPort == 0 {
		prompt4 := promptui.Prompt{
			Label:     fmt.Sprintf("Enter SSH token endpoint port for context %s / %s", cName, apiHostname),
			Default:   "22",
			AllowEdit: true,
		}
		tokenPort, err := prompt4.Run()
		if err != nil {
			return fmt.Errorf("unable to create configuration: %s", err.Error())
		}
		iTokenPort, err = strconv.Atoi(tokenPort)
		if err != nil {
			return fmt.Errorf("provided port is not a number: %s", err.Error())
		}
	}
	err = lConfig.NewContext(config.ContextConfig{
		Name:                   cName,
		APIHostname:            apiHostname,
		AuthenticationEndpoint: authEndpoint,
		TokenHost:              tokenHost,
		TokenPort:              iTokenPort,
		UIHostname:             uiHostname,
		WebhookEndpoint:        webhookEndpoint,
	}, userName)
	if err != nil {
		return fmt.Errorf("unable to create new context: %v", err)
	}
	err = lConfig.SetDefaultContext(cName)
	if err != nil {
		return fmt.Errorf("unable to set default context: %v", err)
	}
	err = lConfig.WriteConfig()
	if err != nil {
		return fmt.Errorf("unable to create configuration: %v", err)
	}
	return nil
}
