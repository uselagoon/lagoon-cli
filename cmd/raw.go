package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/internal/custom"
	"github.com/uselagoon/lagoon-cli/pkg/output"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
)

var emptyCmd = cobra.Command{
	Use:     "none",
	Aliases: []string{""},
	Short:   "none",
	Hidden:  true,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lContext.Name)
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
		return validateTokenE(lContext.Name)
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
		utoken := lUser.UserConfig.Grant.AccessToken
		lc := lclient.New(
			fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
			lagoonCLIVersion,
			lContext.ContextConfig.Version,
			&utoken,
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
The directory for custom commands uses ${XDG_DATA_HOME}/lagoon-commands.
If XDG_DATA_HOME is not defined, a directory will be created with the defaults, this command will output the location at the end.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		commandsDir, err := custom.GetCommandsLocation()
		if err != nil {
			return err
		}
		// just return the help menu for this command as if it is just a normal parent with children commands
		cmd.Help()
		fmt.Println("Save your command YAML files into the following directory")
		fmt.Println(commandsDir)
		return nil
	},
}

func ConvertToCobra(raw custom.CustomCommand) *cobra.Command {
	cCmd := &cobra.Command{
		Use:   raw.Name,
		Short: raw.Description,
		PreRunE: func(_ *cobra.Command, _ []string) error {
			return validateTokenE(lContext.Name)
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

			utoken := lUser.UserConfig.Grant.AccessToken
			lc := lclient.New(
				fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
				lagoonCLIVersion,
				lContext.ContextConfig.Version,
				&utoken,
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
	return cCmd
}

func init() {
	if _, ok := os.LookupEnv("LAGOON_GEN_DOCS"); ok {
		// this is an override for when the docs are generated
		// so that it doesn't include any custom commands
		customCmd.AddCommand(&emptyCmd)
	} else {
		// read any custom commands
		cmds2, err := custom.LoadCommands(true)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		for _, cm := range cmds2.Commands {
			customCmd.AddCommand(ConvertToCobra(cm))
		}
	}
	rawCmd.Flags().String("raw", "", "The raw query or mutation to run")
}
