package custom

import (
	"fmt"
	"os"
	"strings"

	"github.com/adrg/xdg"
	"gopkg.in/yaml.v3"
)

const (
	// this is the directory within the XDG path that command files will be stored
	commandDir  = "lagoon-commands"
	commandFile = ".commands"
)

// CustomCommand is the custom command data structure, this is what can be used to define custom commands
type CustomCommand struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Query       string `yaml:"query"`
	Flags       []struct {
		Name        string       `yaml:"name"`
		Description string       `yaml:"description"`
		Variable    string       `yaml:"variable"`
		Type        string       `yaml:"type"`
		Required    bool         `yaml:"required"`
		Default     *interface{} `yaml:"default,omitempty"`
	} `yaml:"flags"`
}

type Commands struct {
	Commands []CustomCommand `yaml:"commands"`
}

// LoadCommands will load custom commands if they exist
func LoadCommands(create bool) (*Commands, error) {
	c := &Commands{}
	if !create {
		_, err := xdg.SearchDataFile(fmt.Sprintf("%s/%s", commandDir, commandFile))
		if err != nil {
			// if the command directory can't be created, the read will fail
			return c, err
		}
	}
	commandFilePath, err := GetCommandsLocation()
	if err != nil {
		return c, err
	}
	files, err := os.ReadDir(commandFilePath)
	if err != nil {
		return c, err
	}
	for _, f := range files {
		cc, err := readCommandfile(fmt.Sprintf("%s/%s", commandFilePath, f.Name()))
		if err != nil {
			return c, err
		}
		c.Commands = append(c.Commands, *cc)
	}
	return c, nil
}

// GetCommandsLocation will return the commands file locations using XDG
func GetCommandsLocation() (string, error) {
	commandFilePath, err := xdg.DataFile(fmt.Sprintf("%s/%s", commandDir, commandFile))
	if err != nil {
		return "", err
	}
	return strings.Replace(commandFilePath, fmt.Sprintf("/%s", commandFile), "", -1), nil
}

// helper for reading the commands file from a defined path
func readCommandfile(file string) (*CustomCommand, error) {
	rawYAML, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("couldn't read %v: %v", file, err)
	}
	cc := &CustomCommand{}
	err = yaml.Unmarshal(rawYAML, cc)
	if err != nil {
		return nil, err
	}
	return cc, nil
}
