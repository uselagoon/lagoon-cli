package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/uselagoon/lagoon-cli/internal/lagoon"
	"github.com/uselagoon/lagoon-cli/internal/lagoon/client"
	"github.com/uselagoon/lagoon-cli/internal/schema"
	"github.com/uselagoon/lagoon-cli/pkg/output"
	"os"
)

func parseOSFlags[T interface{}](flags pflag.FlagSet, t T) T {
	configMap := make(map[string]interface{})
	flags.VisitAll(func(f *pflag.Flag) {
		if flags.Changed(f.Name) {
			configMap[f.Name] = f.Value
		}
	})
	jsonStr, _ := json.Marshal(configMap)
	parsedFlags := t
	json.Unmarshal(jsonStr, &parsedFlags)
	return parsedFlags
}

var addOSCmd = &cobra.Command{
	Use:     "os",
	Aliases: []string{"os"},
	Short:   "Add an openshift to lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		parsedFlags := parseOSFlags[schema.AddOpenshiftInput](*cmd.Flags(), schema.AddOpenshiftInput{})

		if parsedFlags.Name == "" {
			fmt.Println("Missing arguments: Openshift name is not defined")
			cmd.Help()
			os.Exit(1)
		}

		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			handleError(err)
		}
		current := lagoonCLIConfig.Current
		lc := client.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIConfig.Lagoons[current].Token,
			lagoonCLIConfig.Lagoons[current].Version,
			lagoonCLIVersion,
			debug)

		if yesNo(fmt.Sprintf("You are attempting to add '%s' openshift, are you sure?", parsedFlags.Name)) {
			addOpenshiftResponse, err := lagoon.AddOpenshift(context.TODO(), &parsedFlags, lc)
			if err != nil {
				handleError(err)
			}

			data := []output.Data{}
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%v", addOpenshiftResponse.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", addOpenshiftResponse.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", addOpenshiftResponse.ConsoleUrl)),
				returnNonEmptyString(fmt.Sprintf("%v", addOpenshiftResponse.Token)),
				returnNonEmptyString(fmt.Sprintf("%v", addOpenshiftResponse.SshHost)),
				returnNonEmptyString(fmt.Sprintf("%v", addOpenshiftResponse.SshPort)),
				returnNonEmptyString(fmt.Sprintf("%v", addOpenshiftResponse.CloudRegion)),
				returnNonEmptyString(fmt.Sprintf("%v", addOpenshiftResponse.CloudProvider)),
				returnNonEmptyString(fmt.Sprintf("%v", addOpenshiftResponse.FriendlyName)),
				returnNonEmptyString(fmt.Sprintf("%v", addOpenshiftResponse.RouterPattern)),
				returnNonEmptyString(fmt.Sprintf("%v", addOpenshiftResponse.MonitoringConfig)),
			})
			output.RenderOutput(output.Table{
				Header: []string{
					"ID",
					"Name",
					"ConsoleUrl",
					"Token",
					"ConsoleUrl",
					"SshHost",
					"SshPort",
					"CloudRegion",
					"CloudProvider",
					"FriendlyName",
					"RouterPattern",
					"MonitoringConfig",
				},
				Data: data,
			}, outputOptions)
		}
	},
}

var deleteOpenshiftCmd = &cobra.Command{
	Use:     "os",
	Aliases: []string{"g"},
	Short:   "Delete an openshift from lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		parsedFlags := parseOSFlags[schema.DeleteOpenshiftInput](*cmd.Flags(), schema.DeleteOpenshiftInput{})
		if parsedFlags.Name == "" {
			fmt.Println("Missing arguments: Openshift name is not defined")
			cmd.Help()
			os.Exit(1)
		}

		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			handleError(err)
		}
		current := lagoonCLIConfig.Current
		lc := client.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIConfig.Lagoons[current].Token,
			lagoonCLIConfig.Lagoons[current].Version,
			lagoonCLIVersion,
			debug)

		if yesNo(fmt.Sprintf("You are attempting to delete openshift '%s', are you sure?", parsedFlags.Name)) {
			deleteOpenshiftResponse, err := lagoon.DeleteOpenshift(context.TODO(), &parsedFlags, lc)
			if err != nil {
				handleError(err)
			}

			handleError(err)
			resultData := output.Result{
				Result: deleteOpenshiftResponse.DeleteOpenshift,
			}
			output.RenderResult(resultData, outputOptions)
		}
	},
}

func init() {
	addOSCmd.Flags().StringVarP(&osName, "name", "N", "", "Name of openshift")
	addOSCmd.Flags().StringVarP(&osConsoleUrl, "console-url", "U", "", "Console URL")
	addOSCmd.Flags().StringVarP(&osToken, "token", "T", "", "Openshift token")
	deleteOpenshiftCmd.Flags().StringVarP(&osName, "name", "N", "", "Name of openshift")
}
