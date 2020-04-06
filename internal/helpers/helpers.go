package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

/*
	Place to put any helper functions for the cli
*/

// FileExists check if a file exists
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// GetLagoonConfigFile get the configpath, name and extension from the config init so we can pass it back to be used by the rest of the cli
func GetLagoonConfigFile(configPath *string, configName *string, configExtension *string, createConfig bool, cmd *cobra.Command) error {
	// check if we have an envvar or flag to define our confg file
	var configFilePath string
	configFilePath, err := cmd.Flags().GetString("config-file")
	if err != nil {
		return fmt.Errorf("Error reading flag `config-file`: %v", err)
	}
	if configFilePath == "" {
		if lagoonConfigEnvar, ok := os.LookupEnv("LAGOONCONFIG"); ok {
			configFilePath = lagoonConfigEnvar
		}
	}
	if configFilePath != "" {
		if FileExists(configFilePath) || createConfig {
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

// GetLagoonContext get the lagoon cluster to use
func GetLagoonContext(lagoon *string, cmd *cobra.Command) error {
	// check if we have an envvar or flag to define our lagoon context
	var lagoonContext string
	lagoonContext, err := cmd.Flags().GetString("lagoon")
	if err != nil {
		return fmt.Errorf("Error reading flag `lagoon`: %v", err)
	}
	if lagoonContext == "" {
		if lagoonContextEnvar, ok := os.LookupEnv("LAGOONCONTEXT"); ok {
			lagoonContext = lagoonContextEnvar
		}
	}
	if lagoonContext != "" {
		*lagoon = lagoonContext
	} else {
		if viper.GetString("default") == "" {
			*lagoon = "amazeeio"
		} else {
			*lagoon = viper.GetString("default")
		}
	}
	return nil
}

// StripNewLines will strip new lines from strings helper
func StripNewLines(stripString string) string {
	return strings.TrimSuffix(stripString, "\n")
}
