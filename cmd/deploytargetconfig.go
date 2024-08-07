package cmd

import (
	"context"
	"fmt"

	"strconv"

	"github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	"github.com/uselagoon/machinery/api/schema"

	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/pkg/output"
)

var addDeployTargetConfigCmd = &cobra.Command{
	Use:     "deploytarget-config",
	Aliases: []string{"dtc"},
	Hidden:  false,
	Short:   "Add deploytarget config to a project",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		branches, err := cmd.Flags().GetString("branches")
		if err != nil {
			return err
		}
		pullrequests, err := cmd.Flags().GetString("pullrequests")
		if err != nil {
			return err
		}
		weight, err := cmd.Flags().GetUint("weight")
		if err != nil {
			return err
		}
		deploytarget, err := cmd.Flags().GetUint("deploytarget")
		if err != nil {
			return err
		}

		if err := requiredInputCheck("Project name", cmdProjectName, "Deploytarget", strconv.Itoa(int(deploytarget)), "Pullrequests", pullrequests, "Branches", branches); err != nil {
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

		project, err := lagoon.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		addDeployTargetConfig := &schema.AddDeployTargetConfigInput{
			Project: project.ID,
			Weight:  weight,
		}
		if branches != "" {
			addDeployTargetConfig.Branches = branches
		}
		if pullrequests != "" {
			addDeployTargetConfig.Pullrequests = pullrequests
		}
		if deploytarget != 0 {
			addDeployTargetConfig.DeployTarget = deploytarget
		}
		if yesNo(fmt.Sprintf("You are attempting to add a deploytarget configuration to project '%s', are you sure?", cmdProjectName)) {
			deployTargetConfig, err := lagoon.AddDeployTargetConfiguration(context.TODO(), addDeployTargetConfig, lc)
			if err != nil {
				return err
			}
			data := []output.Data{}
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%v", deployTargetConfig.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", deployTargetConfig.Weight)),
				returnNonEmptyString(fmt.Sprintf("%v", deployTargetConfig.Branches)),
				returnNonEmptyString(fmt.Sprintf("%v", deployTargetConfig.Pullrequests)),
				returnNonEmptyString(fmt.Sprintf("%v", deployTargetConfig.DeployTarget.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", deployTargetConfig.DeployTarget.FriendlyName)),
				returnNonEmptyString(fmt.Sprintf("%v", deployTargetConfig.DeployTarget.CloudProvider)),
				returnNonEmptyString(fmt.Sprintf("%v", deployTargetConfig.DeployTarget.CloudRegion)),
			})
			r := output.RenderOutput(output.Table{
				Header: []string{
					"ID",
					"Weight",
					"Branches",
					"Pullrequests",
					"Name",
					"FriendlyName",
					"CloudProvider",
					"CloudRegion",
				},
				Data: data,
			}, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		}
		return nil
	},
}

var updateDeployTargetConfigCmd = &cobra.Command{
	Use:     "deploytarget-config",
	Aliases: []string{"dtc"},
	Hidden:  false,
	Short:   "Update a deploytarget config",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		id, err := cmd.Flags().GetUint("id")
		if err != nil {
			return err
		}
		branches, err := cmd.Flags().GetString("branches")
		if err != nil {
			return err
		}
		pullrequests, err := cmd.Flags().GetString("pullrequests")
		if err != nil {
			return err
		}
		weight, err := cmd.Flags().GetUint("weight")
		if err != nil {
			return err
		}
		deploytarget, err := cmd.Flags().GetUint("deploytarget")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Deploytarget config id", strconv.Itoa(int(id))); err != nil {
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

		updateDeployTargetConfig := &schema.UpdateDeployTargetConfigInput{
			ID:     id,
			Weight: weight,
		}
		if branches != "" {
			updateDeployTargetConfig.Branches = branches
		}
		if pullrequests != "" {
			updateDeployTargetConfig.Pullrequests = pullrequests
		}
		if deploytarget != 0 {
			updateDeployTargetConfig.DeployTarget = deploytarget
		}

		if yesNo(fmt.Sprintf("You are attempting to update a deploytarget configuration with id '%d', are you sure?", id)) {
			deployTargetConfig, err := lagoon.UpdateDeployTargetConfiguration(context.TODO(), updateDeployTargetConfig, lc)
			if err != nil {
				return err
			}
			data := []output.Data{}
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%v", deployTargetConfig.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", deployTargetConfig.Weight)),
				returnNonEmptyString(fmt.Sprintf("%v", deployTargetConfig.Branches)),
				returnNonEmptyString(fmt.Sprintf("%v", deployTargetConfig.Pullrequests)),
				returnNonEmptyString(fmt.Sprintf("%v", deployTargetConfig.DeployTarget.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", deployTargetConfig.DeployTarget.FriendlyName)),
				returnNonEmptyString(fmt.Sprintf("%v", deployTargetConfig.DeployTarget.CloudProvider)),
				returnNonEmptyString(fmt.Sprintf("%v", deployTargetConfig.DeployTarget.CloudRegion)),
			})
			r := output.RenderOutput(output.Table{
				Header: []string{
					"ID",
					"Weight",
					"Branches",
					"Pullrequests",
					"Name",
					"FriendlyName",
					"CloudProvider",
					"CloudRegion",
				},
				Data: data,
			}, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		}
		return nil
	},
}

var deleteDeployTargetConfigCmd = &cobra.Command{
	Use:     "deploytarget-config",
	Aliases: []string{"dtc"},
	Hidden:  false,
	Short:   "Delete a deploytarget config",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		id, err := cmd.Flags().GetUint("id")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Deploytarget config id", strconv.Itoa(int(id)), "Project name", cmdProjectName); err != nil {
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

		project, err := lagoon.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		if project.Name == "" {
			outputOptions.Error = fmt.Sprintf("No details for project '%s'", cmdProjectName)
			output.RenderError(outputOptions.Error, outputOptions)
			return nil
		}

		if yesNo(fmt.Sprintf("You are attempting to delete deploytarget configuration with id '%d' from project '%s', are you sure?", id, cmdProjectName)) {
			result, err := lagoon.DeleteDeployTargetConfiguration(context.TODO(), int(id), int(project.ID), lc)
			if err != nil {
				return err
			}
			resultData := output.Result{
				Result: result.DeleteDeployTargetConfig,
			}
			r := output.RenderResult(resultData, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		}
		return nil
	},
}

var listDeployTargetConfigsCmd = &cobra.Command{
	Use:     "deploytarget-configs",
	Aliases: []string{"dtc"},
	Hidden:  false,
	Short:   "List deploytarget configs for a project",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}

		if err := requiredInputCheck("Project name", cmdProjectName); err != nil {
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

		project, err := lagoon.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		if project.Name == "" {
			outputOptions.Error = fmt.Sprintf("No details for project '%s'", cmdProjectName)
			output.RenderError(outputOptions.Error, outputOptions)
			return nil
		}
		deployTargetConfigs, err := lagoon.GetDeployTargetConfigs(context.TODO(), int(project.ID), lc)
		if err != nil {
			return err
		}
		data := []output.Data{}
		for _, deployTargetConfig := range *deployTargetConfigs {
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%v", deployTargetConfig.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", deployTargetConfig.Weight)),
				returnNonEmptyString(fmt.Sprintf("%v", deployTargetConfig.Branches)),
				returnNonEmptyString(fmt.Sprintf("%v", deployTargetConfig.Pullrequests)),
				returnNonEmptyString(fmt.Sprintf("%v", deployTargetConfig.DeployTarget.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", deployTargetConfig.DeployTarget.FriendlyName)),
				returnNonEmptyString(fmt.Sprintf("%v", deployTargetConfig.DeployTarget.CloudProvider)),
				returnNonEmptyString(fmt.Sprintf("%v", deployTargetConfig.DeployTarget.CloudRegion)),
			})
		}
		r := output.RenderOutput(output.Table{
			Header: []string{
				"ID",
				"Weight",
				"Branches",
				"Pullrequests",
				"Name",
				"FriendlyName",
				"CloudProvider",
				"CloudRegion",
			},
			Data: data,
		}, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		return nil
	},
}

func init() {
	addDeployTargetConfigCmd.Flags().StringP("branches", "b", "", "Branches regex")
	addDeployTargetConfigCmd.Flags().StringP("pullrequests", "P", "", "Pullrequests title regex")
	addDeployTargetConfigCmd.Flags().UintP("weight", "w", 1, "Deploytarget config weighting (default:1)")
	addDeployTargetConfigCmd.Flags().UintP("deploytarget", "d", 0, "Deploytarget id")

	updateDeployTargetConfigCmd.Flags().StringP("branches", "b", "", "Branches regex")
	updateDeployTargetConfigCmd.Flags().StringP("pullrequests", "P", "", "Pullrequests title regex")
	updateDeployTargetConfigCmd.Flags().UintP("weight", "w", 1, "Deploytarget config weighting (default:1)")
	updateDeployTargetConfigCmd.Flags().UintP("deploytarget", "d", 0, "Deploytarget id")
	updateDeployTargetConfigCmd.Flags().UintP("id", "I", 0, "Deploytarget config id")

	deleteDeployTargetConfigCmd.Flags().UintP("id", "I", 0, "Deploytarget config id")
}
