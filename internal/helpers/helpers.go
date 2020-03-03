package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
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

// GetLagoonConfigFile .
func GetLagoonConfigFile(userPath *string, configName *string, configExtension *string, cmd *cobra.Command) error {
	// check if we have an envvar to define our confg file
	var configFilePath string
	createConfig, err := cmd.Flags().GetBool("create-config")
	if err != nil {
		return err
	}
	if lagoonConfigEnvar, ok := os.LookupEnv("LAGOONCONFIG"); !ok {
		configFilePath = lagoonConfigEnvar
	}
	// otherwise check the flag
	configFilePath, err = cmd.Flags().GetString("config-file")
	if err != nil {
		return err
	}
	if configFilePath != "" {
		if FileExists(configFilePath) || createConfig {
			*userPath = filepath.Dir(configFilePath)
			*configExtension = filepath.Ext(configFilePath)
			*configName = strings.TrimSuffix(filepath.Base(configFilePath), *configExtension)
			return nil
		}
		return fmt.Errorf("%s/%s File doesn't exist", *userPath, configFilePath)

	}
	// no config file found
	return nil
}
