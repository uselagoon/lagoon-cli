package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/pkg/output"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	"gopkg.in/yaml.v3"
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

var emptyCmd = cobra.Command{
	Use:     "none",
	Aliases: []string{""},
	Short:   "none",
	Hidden:  true,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var rawCmd = &cobra.Command{
	Use:     "raw",
	Aliases: []string{"r"},
	Short:   "Run a custom query or mutation",
	Long: `Run a custom query or mutation.
The output of this command will be the JSON response from the API`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		raw, err := cmd.Flags().GetString("raw")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Raw query or mutation", raw); err != nil {
			return err
		}
		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)
		if err != nil {
			return err
		}
		rawResp, err := lc.ProcessRaw(context.TODO(), raw, nil)
		if err != nil {
			return err
		}
		r, err := json.Marshal(rawResp)
		if err != nil {
			return err
		}
		fmt.Println(string(r))
		return nil
	},
}

var customCmd = &cobra.Command{
	Use:     "custom",
	Aliases: []string{"cus", "cust"},
	Short:   "Run a custom command",
	Long: `Run a custom command.
This command alone does nothing, but you can create custom commands and put them into the custom commands directory,
these commands will then be available to use.
The directory for custom commands is ${HOME}/.lagoon-cli/commands.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// just return the help menu for this command as if it is just a normal parent with children commands
		cmd.Help()
		return nil
	},
}

func ReadCustomCommands() ([]*cobra.Command, error) {
	userPath, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("couldn't get $HOME: %v", err)
	}
	customCommandsFilePath := fmt.Sprintf("%s/%s", userPath, commandsFilePath)
	if _, err := os.Stat(customCommandsFilePath); os.IsNotExist(err) {
		err := os.MkdirAll(customCommandsFilePath, 0700)
		if err != nil {
			return nil, fmt.Errorf("couldn't create command directory %s: %v", customCommandsFilePath, err)
		}
	}
	files, err := os.ReadDir(customCommandsFilePath)
	if err != nil {
		return nil, fmt.Errorf("couldn't open command directory %s: %v", customCommandsFilePath, err)
	}
	var cmds []*cobra.Command
	if len(files) != 0 {
		for _, file := range files {
			if !file.IsDir() {
				data, err := os.ReadFile(customCommandsFilePath + "/" + file.Name())
				if err != nil {
					return nil, err
				}
				raw := CustomCommand{}
				err = yaml.Unmarshal(data, &raw)
				if err != nil {
					return nil, fmt.Errorf("unable to unmarshal custom command '%s', yaml is likely invalid: %v", file.Name(), err)
				}
				cCmd := cobra.Command{
					Use:     raw.Name,
					Aliases: []string{""},
					Short:   raw.Description,
					PreRunE: func(_ *cobra.Command, _ []string) error {
						return validateTokenE(lagoonCLIConfig.Current)
					},
					RunE: func(cmd *cobra.Command, args []string) error {
						debug, err := cmd.Flags().GetBool("debug")
						if err != nil {
							return err
						}

						variables := make(map[string]interface{})
						var value interface{}
						// handling reading the custom flags
						for _, flag := range raw.Flags {
							switch flag.Type {
							case "Int":
								value, err = cmd.Flags().GetInt(flag.Name)
								if err != nil {
									return err
								}
								if flag.Required {
									if err := requiredInputCheck(flag.Name, fmt.Sprintf("%d", value.(int))); err != nil {
										return err
									}
								}
							case "String":
								value, err = cmd.Flags().GetString(flag.Name)
								if err != nil {
									return err
								}
								if flag.Required {
									if err := requiredInputCheck(flag.Name, value.(string)); err != nil {
										return err
									}
								}
							case "Boolean":
								value, err = cmd.Flags().GetBool(flag.Name)
								if err != nil {
									return err
								}
							}
							variables[flag.Variable] = value
						}

						current := lagoonCLIConfig.Current
						token := lagoonCLIConfig.Lagoons[current].Token
						lc := lclient.New(
							lagoonCLIConfig.Lagoons[current].GraphQL,
							lagoonCLIVersion,
							&token,
							debug)
						if err != nil {
							return err
						}

						rawResp, err := lc.ProcessRaw(context.TODO(), raw.Query, variables)
						if err != nil {
							return err
						}
						r, err := json.Marshal(rawResp)
						if err != nil {
							return err
						}
						fmt.Println(string(r))
						return nil
					},
				}
				// add custom flags to the command
				for _, flag := range raw.Flags {
					switch flag.Type {
					case "Int":
						if flag.Default != nil {
							cCmd.Flags().Int(flag.Name, (*flag.Default).(int), flag.Description)
						} else {
							cCmd.Flags().Int(flag.Name, 0, flag.Description)
						}
					case "String":
						if flag.Default != nil {
							cCmd.Flags().String(flag.Name, (*flag.Default).(string), flag.Description)
						} else {
							cCmd.Flags().String(flag.Name, "", flag.Description)
						}
					case "Boolean":
						if flag.Default != nil {
							cCmd.Flags().Bool(flag.Name, (*flag.Default).(bool), flag.Description)
						} else {
							cCmd.Flags().Bool(flag.Name, false, flag.Description)
						}
					}
				}
				cmds = append(cmds, &cCmd)
			}
		}
	} else {
		cmds = append(cmds,
			// create a hidden command that does nothing so help and docs can be generated for the custom command
			&emptyCmd)
	}
	return cmds, nil
}

func init() {
	if _, ok := os.LookupEnv("LAGOON_GEN_DOCS"); ok {
		// this is an override for when the docs are generated
		// so that it doesn't include any custom commands
		customCmd.AddCommand(&emptyCmd)
	} else {
		// read any custom commands
		cmds, err := ReadCustomCommands()
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		for _, c := range cmds {
			customCmd.AddCommand(c)
		}
	}
	rawCmd.Flags().String("raw", "", "The raw query or mutation to run")
}
