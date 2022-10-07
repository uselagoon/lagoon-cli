package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/internal/lagoon"
	"github.com/uselagoon/lagoon-cli/internal/lagoon/client"
	"github.com/uselagoon/lagoon-cli/internal/schema"
	"github.com/uselagoon/lagoon-cli/pkg/output"
)

var addDeployTargetCmd = &cobra.Command{
	Use:     "deploytarget",
	Aliases: []string{"dt"},
	Short:   "Add a deploytarget to lagoon",
	Long:    "Add a deploytarget (kubernetes or openshift) to lagoon, this requires admin level permissions",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		token, err := cmd.Flags().GetString("token")
		if err != nil {
			return err
		}
		consoleURL, err := cmd.Flags().GetString("console-url")
		if err != nil {
			return err
		}
		routerPattern, err := cmd.Flags().GetString("router-pattern")
		if err != nil {
			return err
		}
		sshHost, err := cmd.Flags().GetString("ssh-host")
		if err != nil {
			return err
		}
		sshPort, err := cmd.Flags().GetString("ssh-port")
		if err != nil {
			return err
		}
		friendlyName, err := cmd.Flags().GetString("friendly-name")
		if err != nil {
			return err
		}
		cloudProvider, err := cmd.Flags().GetString("cloud-provider")
		if err != nil {
			return err
		}
		cloudRegion, err := cmd.Flags().GetString("cloud-region")
		if err != nil {
			return err
		}

		if name == "" {
			return fmt.Errorf("Missing arguments: name is not defined")
		}
		if token == "" {
			return fmt.Errorf("Missing arguments: token is not defined")
		}
		if consoleURL == "" {
			return fmt.Errorf("Missing arguments: console-url is not defined")
		}

		addDeployTarget := &schema.AddDeployTargetInput{
			Name:          name,
			FriendlyName:  friendlyName,
			CloudProvider: cloudProvider,
			CloudRegion:   cloudRegion,
			Token:         token,
			RouterPattern: routerPattern,
			ConsoleURL:    consoleURL,
			SSHHost:       sshHost,
			SSHPort:       sshPort,
		}
		id, err := cmd.Flags().GetUint("id")
		if err != nil {
			return err
		}
		if id != 0 {
			addDeployTarget.ID = id
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

		if yesNo(fmt.Sprintf("You are attempting to add '%s' deploytarget, are you sure?", addDeployTarget.Name)) {
			addDeployTargetResponse, err := lagoon.AddDeployTarget(context.TODO(), addDeployTarget, lc)
			if err != nil {
				handleError(err)
			}

			data := []output.Data{}
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%v", addDeployTargetResponse.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", addDeployTargetResponse.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", addDeployTargetResponse.ConsoleURL)),
				returnNonEmptyString(fmt.Sprintf("%v", addDeployTargetResponse.Token)),
				returnNonEmptyString(fmt.Sprintf("%v", addDeployTargetResponse.SSHHost)),
				returnNonEmptyString(fmt.Sprintf("%v", addDeployTargetResponse.SSHPort)),
				returnNonEmptyString(fmt.Sprintf("%v", addDeployTargetResponse.CloudRegion)),
				returnNonEmptyString(fmt.Sprintf("%v", addDeployTargetResponse.CloudProvider)),
				returnNonEmptyString(fmt.Sprintf("%v", addDeployTargetResponse.FriendlyName)),
				returnNonEmptyString(fmt.Sprintf("%v", addDeployTargetResponse.RouterPattern)),
				returnNonEmptyString(fmt.Sprintf("%v", addDeployTargetResponse.MonitoringConfig)),
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
		return nil
	},
}

var updateDeployTargetCmd = &cobra.Command{
	Use:     "deploytarget",
	Aliases: []string{"dt"},
	Short:   "Update a deploytarget in lagoon",
	Long:    "Update a deploytarget (kubernetes or openshift) in lagoon, this requires admin level permissions",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetUint("id")
		if err != nil {
			return err
		}
		token, err := cmd.Flags().GetString("token")
		if err != nil {
			return err
		}
		consoleURL, err := cmd.Flags().GetString("console-url")
		if err != nil {
			return err
		}
		routerPattern, err := cmd.Flags().GetString("router-pattern")
		if err != nil {
			return err
		}
		sshHost, err := cmd.Flags().GetString("ssh-host")
		if err != nil {
			return err
		}
		sshPort, err := cmd.Flags().GetString("ssh-port")
		if err != nil {
			return err
		}
		friendlyName, err := cmd.Flags().GetString("friendly-name")
		if err != nil {
			return err
		}
		cloudProvider, err := cmd.Flags().GetString("cloud-provider")
		if err != nil {
			return err
		}
		cloudRegion, err := cmd.Flags().GetString("cloud-region")
		if err != nil {
			return err
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
			debug,
		)
		updateDeployTarget := &schema.UpdateDeployTargetInput{
			AddDeployTargetInput: schema.AddDeployTargetInput{
				ID:            id,
				Token:         token,
				FriendlyName:  friendlyName,
				CloudProvider: cloudProvider,
				CloudRegion:   cloudRegion,
				RouterPattern: routerPattern,
				ConsoleURL:    consoleURL,
				SSHHost:       sshHost,
				SSHPort:       sshPort,
			},
		}
		if yesNo(fmt.Sprintf("You are attempting to update '%d' deploytarget, are you sure?", updateDeployTarget.ID)) {
			updateDeployTargetResponse, err := lagoon.UpdateDeployTarget(context.TODO(), updateDeployTarget, lc)
			if err != nil {
				handleError(err)
			}

			data := []output.Data{}
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%v", updateDeployTargetResponse.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", updateDeployTargetResponse.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", updateDeployTargetResponse.ConsoleURL)),
				returnNonEmptyString(fmt.Sprintf("%v", updateDeployTargetResponse.Token)),
				returnNonEmptyString(fmt.Sprintf("%v", updateDeployTargetResponse.SSHHost)),
				returnNonEmptyString(fmt.Sprintf("%v", updateDeployTargetResponse.SSHPort)),
				returnNonEmptyString(fmt.Sprintf("%v", updateDeployTargetResponse.CloudRegion)),
				returnNonEmptyString(fmt.Sprintf("%v", updateDeployTargetResponse.CloudProvider)),
				returnNonEmptyString(fmt.Sprintf("%v", updateDeployTargetResponse.FriendlyName)),
				returnNonEmptyString(fmt.Sprintf("%v", updateDeployTargetResponse.RouterPattern)),
				returnNonEmptyString(fmt.Sprintf("%v", updateDeployTargetResponse.MonitoringConfig)),
			})
			output.RenderOutput(output.Table{
				Header: []string{
					"ID",
					"Name",
					"ConsoleUrl",
					"Token",
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
		return nil
	},
}

var deleteDeployTargetCmd = &cobra.Command{
	Use:     "deploytarget",
	Aliases: []string{"dt"},
	Short:   "Delete a deploytarget from lagoon",
	Long:    "Delete a deploytarget (kubernetes or openshift) from lagoon, this requires admin level permissions",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
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
			debug,
		)

		deleteDeployTarget := &schema.DeleteDeployTargetInput{
			Name: name,
		}
		if yesNo(fmt.Sprintf("You are attempting to delete deploytarget '%s', are you sure?", deleteDeployTarget.Name)) {
			deleteDeployTargetResponse, err := lagoon.DeleteDeployTarget(context.TODO(), deleteDeployTarget, lc)
			if err != nil {
				handleError(err)
			}

			handleError(err)
			resultData := output.Result{
				Result: deleteDeployTargetResponse.DeleteDeployTarget,
			}
			output.RenderResult(resultData, outputOptions)
		}
		return nil
	},
}

func init() {
	addDeployTargetCmd.Flags().UintP("id", "", 0, "Deploytarget id")
	addDeployTargetCmd.Flags().StringP("name", "", "", "Name of deploytarget")
	addDeployTargetCmd.Flags().StringP("console-url", "", "", "Console URL")
	addDeployTargetCmd.Flags().StringP("token", "", "", "deploytarget token")
	addDeployTargetCmd.Flags().StringP("router-pattern", "", "", "deploytarget router-pattern")
	addDeployTargetCmd.Flags().StringP("friendly-name", "", "", "deploytarget friendly name")
	addDeployTargetCmd.Flags().StringP("cloud-provider", "", "", "deploytarget cloud provider")
	addDeployTargetCmd.Flags().StringP("cloud-region", "", "", "deploytarget cloud region")
	addDeployTargetCmd.Flags().StringP("ssh-host", "", "", "deploytarget ssh host")
	addDeployTargetCmd.Flags().StringP("ssh-port", "", "", "deploytarget ssh port")

	deleteDeployTargetCmd.Flags().UintP("id", "", 0, "Deploytarget id")
	deleteDeployTargetCmd.Flags().StringP("name", "", "", "Name of deploytarget")

	updateDeployTargetCmd.Flags().UintP("id", "", 0, "Deploytarget id")
	updateDeployTargetCmd.Flags().StringP("console-url", "", "", "Console URL")
	updateDeployTargetCmd.Flags().StringP("token", "", "", "deploytarget token")
	updateDeployTargetCmd.Flags().StringP("router-pattern", "", "", "deploytarget router-pattern")
	updateDeployTargetCmd.Flags().StringP("friendly-name", "", "", "deploytarget friendly name")
	updateDeployTargetCmd.Flags().StringP("cloud-provider", "", "", "deploytarget cloud provider")
	updateDeployTargetCmd.Flags().StringP("cloud-region", "", "", "deploytarget cloud region")
	updateDeployTargetCmd.Flags().StringP("ssh-host", "", "", "deploytarget ssh host")
	updateDeployTargetCmd.Flags().StringP("ssh-port", "", "", "deploytarget ssh port")
}
