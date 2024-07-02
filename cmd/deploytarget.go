package cmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/pkg/output"
	"github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	"github.com/uselagoon/machinery/api/schema"
)

var addDeployTargetCmd = &cobra.Command{
	Use:     "deploytarget",
	Aliases: []string{"dt"},
	Short:   "Add a DeployTarget to lagoon",
	Long:    "Add a Deploytarget(Kubernetes) to lagoon, this requires admin level permissions",
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
		buildImage, err := cmd.Flags().GetString("build-image")
		if err != nil {
			return err
		}

		if err := requiredInputCheck("Name", name, "Token", token, "Console-url", consoleURL); err != nil {
			return err
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
			BuildImage:    buildImage,
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
			return err
		}
		current := lagoonCLIConfig.Current
		lagoonToken := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			lagoonCLIConfig.Lagoons[current].Version,
			&lagoonToken,
			debug)

		if yesNo(fmt.Sprintf("You are attempting to add '%s' DeployTarget, are you sure?", addDeployTarget.Name)) {
			addDeployTargetResponse, err := lagoon.AddDeployTarget(context.TODO(), addDeployTarget, lc)
			if err != nil {
				return err
			}

			data := []output.Data{}
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%v", addDeployTargetResponse.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", addDeployTargetResponse.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", addDeployTargetResponse.ConsoleURL)),
				returnNonEmptyString(fmt.Sprintf("%v", addDeployTargetResponse.BuildImage)),
				returnNonEmptyString(fmt.Sprintf("%v", addDeployTargetResponse.Token)),
				returnNonEmptyString(fmt.Sprintf("%v", addDeployTargetResponse.SSHHost)),
				returnNonEmptyString(fmt.Sprintf("%v", addDeployTargetResponse.SSHPort)),
				returnNonEmptyString(fmt.Sprintf("%v", addDeployTargetResponse.CloudRegion)),
				returnNonEmptyString(fmt.Sprintf("%v", addDeployTargetResponse.CloudProvider)),
				returnNonEmptyString(fmt.Sprintf("%v", addDeployTargetResponse.FriendlyName)),
				returnNonEmptyString(fmt.Sprintf("%v", addDeployTargetResponse.RouterPattern)),
				returnNonEmptyString(fmt.Sprintf("%v", addDeployTargetResponse.Created)),
				returnNonEmptyString(fmt.Sprintf("%v", addDeployTargetResponse.MonitoringConfig)),
			})
			r := output.RenderOutput(output.Table{
				Header: []string{
					"ID",
					"Name",
					"ConsoleUrl",
					"BuildImage",
					"Token",
					"SshHost",
					"SshPort",
					"CloudRegion",
					"CloudProvider",
					"FriendlyName",
					"RouterPattern",
					"Created",
					"MonitoringConfig",
				},
				Data: data,
			}, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		}
		return nil
	},
}

var updateDeployTargetCmd = &cobra.Command{
	Use:     "deploytarget",
	Aliases: []string{"dt"},
	Short:   "Update a DeployTarget in lagoon",
	Long:    "Update a Deploytarget(Kubernetes) in lagoon, this requires admin level permissions",
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
		// since there needs to be a way to unset the build image (using `null`)
		// use this helper function to get the `null` representation
		// the buildimage field in the schema is *null.String so that it is omit if it is empty
		// but if it is set to "" to clear the value, will pass the json `null` representation
		// or if set to a string, will pass this into the payload
		buildImage, err := flagStringNullValueOrNil(cmd.Flags(), "build-image")
		if err != nil {
			return err
		}

		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}

		if err := requiredInputCheck("Deploytarget ID", strconv.Itoa(int(id))); err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		lagoonToken := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			lagoonCLIConfig.Lagoons[current].Version,
			&lagoonToken,
			debug)

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
			BuildImage: buildImage,
		}
		if yesNo(fmt.Sprintf("You are attempting to update '%d' DeployTarget, are you sure?", updateDeployTarget.ID)) {
			updateDeployTargetResponse, err := lagoon.UpdateDeployTarget(context.TODO(), updateDeployTarget, lc)
			if err != nil {
				return err
			}

			data := []output.Data{}
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%v", updateDeployTargetResponse.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", updateDeployTargetResponse.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", updateDeployTargetResponse.ConsoleURL)),
				returnNonEmptyString(fmt.Sprintf("%v", updateDeployTargetResponse.BuildImage)),
				returnNonEmptyString(fmt.Sprintf("%v", updateDeployTargetResponse.Token)),
				returnNonEmptyString(fmt.Sprintf("%v", updateDeployTargetResponse.SSHHost)),
				returnNonEmptyString(fmt.Sprintf("%v", updateDeployTargetResponse.SSHPort)),
				returnNonEmptyString(fmt.Sprintf("%v", updateDeployTargetResponse.CloudRegion)),
				returnNonEmptyString(fmt.Sprintf("%v", updateDeployTargetResponse.CloudProvider)),
				returnNonEmptyString(fmt.Sprintf("%v", updateDeployTargetResponse.FriendlyName)),
				returnNonEmptyString(fmt.Sprintf("%v", updateDeployTargetResponse.RouterPattern)),
				returnNonEmptyString(fmt.Sprintf("%v", updateDeployTargetResponse.Created)),
				returnNonEmptyString(fmt.Sprintf("%v", updateDeployTargetResponse.MonitoringConfig)),
			})
			r := output.RenderOutput(output.Table{
				Header: []string{
					"ID",
					"Name",
					"ConsoleUrl",
					"BuildImage",
					"Token",
					"SshHost",
					"SshPort",
					"CloudRegion",
					"CloudProvider",
					"FriendlyName",
					"RouterPattern",
					"Created",
					"MonitoringConfig",
				},
				Data: data,
			}, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		}
		return nil
	},
}

var deleteDeployTargetCmd = &cobra.Command{
	Use:     "deploytarget",
	Aliases: []string{"dt"},
	Short:   "Delete a DeployTarget from lagoon",
	Long:    "Delete a Deploytarget(Kubernetes) from lagoon, this requires admin level permissions",
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
			return err
		}

		if err := requiredInputCheck("Deploytarget name", name); err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)

		deleteDeployTarget := &schema.DeleteDeployTargetInput{
			Name: name,
		}
		if yesNo(fmt.Sprintf("You are attempting to delete DeployTarget '%s', are you sure?", deleteDeployTarget.Name)) {
			deleteDeployTargetResponse, err := lagoon.DeleteDeployTarget(context.TODO(), deleteDeployTarget, lc)
			if err != nil {
				return err
			}

			resultData := output.Result{
				Result: deleteDeployTargetResponse.DeleteDeployTarget,
			}
			r := output.RenderResult(resultData, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		}
		return nil
	},
}

var addDeployTargetToOrganizationCmd = &cobra.Command{
	Use:     "organization-deploytarget",
	Aliases: []string{"org-dt"},
	Short:   "Add a deploy target to an Organization",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		organizationName, err := cmd.Flags().GetString("organization-name")
		if err != nil {
			return err
		}
		deploytarget, err := cmd.Flags().GetUint("deploytarget")
		if err != nil {
			return err
		}

		if err := requiredInputCheck("Organization name", organizationName, "Deploy Target", strconv.Itoa(int(deploytarget))); err != nil {
			return err
		}
		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)

		organization, err := lagoon.GetOrganizationByName(context.TODO(), organizationName, lc)
		if err != nil {
			return err
		}
		if organization.Name == "" {
			return fmt.Errorf("error querying organization by name")
		}

		deployTargetInput := schema.AddDeployTargetToOrganizationInput{
			DeployTarget: deploytarget,
			Organization: organization.ID,
		}

		deployTargetResponse, err := lagoon.AddDeployTargetToOrganization(context.TODO(), &deployTargetInput, lc)
		if err != nil {
			return err
		}

		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"Deploy Target":     deploytarget,
				"Organization Name": deployTargetResponse.Name,
			},
		}
		r := output.RenderResult(resultData, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		return nil
	},
}

var removeDeployTargetFromOrganizationCmd = &cobra.Command{
	Use:     "organization-deploytarget",
	Aliases: []string{"org-dt"},
	Short:   "Remove a deploy target from an Organization",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}

		organizationName, err := cmd.Flags().GetString("organization-name")
		if err != nil {
			return err
		}
		deploytarget, err := cmd.Flags().GetUint("deploytarget")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Organization name", organizationName, "Deploy Target", strconv.Itoa(int(deploytarget))); err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)

		organization, err := lagoon.GetOrganizationByName(context.TODO(), organizationName, lc)
		if err != nil {
			return err
		}
		if organization.Name == "" {
			return fmt.Errorf("error querying organization by name")
		}

		deployTargetInput := schema.RemoveDeployTargetFromOrganizationInput{
			DeployTarget: deploytarget,
			Organization: organization.ID,
		}

		if yesNo(fmt.Sprintf("You are attempting to remove deploy target '%d' from organization '%s', are you sure?", deploytarget, organization.Name)) {
			_, err := lagoon.RemoveDeployTargetFromOrganization(context.TODO(), &deployTargetInput, lc)
			if err != nil {
				return err
			}
			resultData := output.Result{
				Result: "success",
				ResultData: map[string]interface{}{
					"Deploy Target":     deploytarget,
					"Organization Name": organizationName,
				},
			}
			r := output.RenderResult(resultData, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		}
		return nil
	},
}

func init() {
	addDeployTargetCmd.Flags().UintP("id", "", 0, "ID of the DeployTarget")
	addDeployTargetCmd.Flags().StringP("name", "", "", "Name of DeployTarget")
	addDeployTargetCmd.Flags().StringP("console-url", "", "", "DeployTarget console URL")
	addDeployTargetCmd.Flags().StringP("token", "", "", "DeployTarget token")
	addDeployTargetCmd.Flags().StringP("router-pattern", "", "", "DeployTarget router-pattern")
	addDeployTargetCmd.Flags().StringP("friendly-name", "", "", "DeployTarget friendly name")
	addDeployTargetCmd.Flags().StringP("cloud-provider", "", "", "DeployTarget cloud provider")
	addDeployTargetCmd.Flags().StringP("cloud-region", "", "", "DeployTarget cloud region")
	addDeployTargetCmd.Flags().StringP("ssh-host", "", "", "DeployTarget ssh host")
	addDeployTargetCmd.Flags().StringP("ssh-port", "", "", "DeployTarget ssh port")
	addDeployTargetCmd.Flags().StringP("build-image", "", "", "DeployTarget build image to use (if different to the default)")

	addDeployTargetToOrganizationCmd.Flags().StringP("organization-name", "O", "", "Name of Organization")
	addDeployTargetToOrganizationCmd.Flags().UintP("deploytarget", "D", 0, "ID of DeployTarget")

	deleteDeployTargetCmd.Flags().UintP("id", "", 0, "ID of the DeployTarget")
	deleteDeployTargetCmd.Flags().StringP("name", "", "", "Name of DeployTarget")

	removeDeployTargetFromOrganizationCmd.Flags().StringP("organization-name", "O", "", "Name of Organization")
	removeDeployTargetFromOrganizationCmd.Flags().UintP("deploytarget", "D", 0, "ID of DeployTarget")

	updateDeployTargetCmd.Flags().UintP("id", "", 0, "ID of the DeployTarget")
	updateDeployTargetCmd.Flags().StringP("console-url", "", "", "DeployTarget console URL")
	updateDeployTargetCmd.Flags().StringP("token", "", "", "DeployTarget token")
	updateDeployTargetCmd.Flags().StringP("router-pattern", "", "", "DeployTarget router-pattern")
	updateDeployTargetCmd.Flags().StringP("friendly-name", "", "", "DeployTarget friendly name")
	updateDeployTargetCmd.Flags().StringP("cloud-provider", "", "", "DeployTarget cloud provider")
	updateDeployTargetCmd.Flags().StringP("cloud-region", "", "", "DeployTarget cloud region")
	updateDeployTargetCmd.Flags().StringP("ssh-host", "", "", "DeployTarget ssh host")
	updateDeployTargetCmd.Flags().StringP("ssh-port", "", "", "DeployTarget ssh port")
	updateDeployTargetCmd.Flags().StringP("build-image", "", "", "DeployTarget build image to use (if different to the default, use \"\" to clear)")
}
