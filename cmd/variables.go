package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	"github.com/uselagoon/machinery/api/schema"

	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/pkg/output"
)

var addVariableCmd = &cobra.Command{
	Use:     "variable",
	Aliases: []string{"v"},
	Short:   "Add or update a variable to an environment or project",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: addOrUpdateVariable,
}

var updateVariableCmd = &cobra.Command{
	Use:     "variable",
	Aliases: []string{"v"},
	Short:   "Add or update a variable to an environment or project",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: addOrUpdateVariable,
}

func addOrUpdateVariable(cmd *cobra.Command, args []string) error {
	debug, err := cmd.Flags().GetBool("debug")
	if err != nil {
		return err
	}
	varName, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}
	varValue, err := cmd.Flags().GetString("value")
	if err != nil {
		return err
	}
	varScope, err := cmd.Flags().GetString("scope")
	if err != nil {
		return err
	}
	organizationName, err := cmd.Flags().GetString("organization-name")
	if err != nil {
		return err
	}
	_, err = getEnvVarType(organizationName, cmdProjectName, cmdProjectEnvironment)
	if err != nil {
		return err
	}
	if err := requiredInputCheck("Variable name", varName); err != nil {
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

	in := &schema.EnvVariableByNameInput{
		Organization: organizationName,
		Project:      cmdProjectName,
		Environment:  cmdProjectEnvironment,
		Scope:        schema.EnvVariableScope(strings.ToUpper(varScope)),
		Name:         varName,
		Value:        varValue,
	}
	envvar, err := lagoon.AddOrUpdateEnvVariableByName(context.TODO(), in, lc)
	if err != nil {
		return err
	}

	if envvar.ID != 0 {
		header := []string{
			"ID",
		}
		env := []string{
			returnNonEmptyString(fmt.Sprintf("%v", envvar.ID)),
		}
		if organizationName != "" {
			header = append(header, "Organization")
			env = append(env, returnNonEmptyString(fmt.Sprintf("%v", organizationName)))
		}
		if cmdProjectName != "" {
			header = append(header, "Project")
			env = append(env, returnNonEmptyString(fmt.Sprintf("%v", cmdProjectName)))
		}
		if cmdProjectEnvironment != "" {
			header = append(header, "Environment")
			env = append(env, returnNonEmptyString(fmt.Sprintf("%v", cmdProjectEnvironment)))
		}
		header = append(header, "Scope")
		env = append(env, returnNonEmptyString(fmt.Sprintf("%v", envvar.Scope)))
		header = append(header, "Name")
		env = append(env, returnNonEmptyString(fmt.Sprintf("%v", envvar.Name)))
		header = append(header, "Value")
		env = append(env, fmt.Sprintf("%v", envvar.Value))
		data := []output.Data{env}
		r := output.RenderOutput(output.Table{
			Header: header,
			Data:   data,
		}, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
	} else {
		return handleNilResults("Variable '%s' remained unchanged\n", cmd, varName)
	}
	return nil
}

var deleteVariableCmd = &cobra.Command{
	Use:     "variable",
	Aliases: []string{"v"},
	Short:   "Delete a variable from an environment",
	Long:    `This allows you to delete an environment variable from a project.`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		varName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		organizationName, err := cmd.Flags().GetString("organization-name")
		if err != nil {
			return err
		}
		envVarType, err := getEnvVarType(organizationName, cmdProjectName, cmdProjectEnvironment)
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Variable name", varName); err != nil {
			return err
		}

		var deleteMsg string
		if envVarType == "organization" {
			deleteMsg = fmt.Sprintf("You are attempting to delete variable '%s' from organization '%s', are you sure?", varName, organizationName)
		} else if envVarType == "project" {
			deleteMsg = fmt.Sprintf("You are attempting to delete variable '%s' from project '%s', are you sure?", varName, cmdProjectName)
		} else if envVarType == "environment" {
			deleteMsg = fmt.Sprintf("You are attempting to delete variable '%s' from environment '%s' in project '%s', are you sure?", varName, cmdProjectEnvironment, cmdProjectName)
		}
		if yesNo(deleteMsg) {
			current := lagoonCLIConfig.Current
			token := lagoonCLIConfig.Lagoons[current].Token
			lc := lclient.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIVersion,
				lagoonCLIConfig.Lagoons[current].Version,
				&token,
				debug)
			in := &schema.DeleteEnvVariableByNameInput{
				Organization: organizationName,
				Project:      cmdProjectName,
				Environment:  cmdProjectEnvironment,
				Name:         varName,
			}
			deleteResult, err := lagoon.DeleteEnvVariableByName(context.TODO(), in, lc)
			if err != nil {
				return err
			}
			resultData := output.Result{
				Result: deleteResult.DeleteEnvVar,
			}
			r := output.RenderResult(resultData, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		}
		return nil
	},
}

func init() {
	addVariableCmd.Flags().StringP("name", "N", "", "Name of the variable to add")
	addVariableCmd.Flags().StringP("value", "V", "", "Value of the variable to add")
	addVariableCmd.Flags().StringP("scope", "S", "", "Scope of the variable[global, build, runtime, container_registry, internal_container_registry]")
	addVariableCmd.Flags().StringP("organization-name", "O", "", "Name of the organization to add variable to")
	updateVariableCmd.Flags().StringP("name", "N", "", "Name of the variable to update")
	updateVariableCmd.Flags().StringP("value", "V", "", "Value of the variable to update")
	updateVariableCmd.Flags().StringP("scope", "S", "", "Scope of the variable[global, build, runtime, container_registry, internal_container_registry]")
	updateVariableCmd.Flags().StringP("organization-name", "O", "", "Name of the organization to update variable for")
	deleteVariableCmd.Flags().StringP("name", "N", "", "Name of the variable to delete")
	deleteVariableCmd.Flags().StringP("organization-name", "O", "", "Name of the organization to delete variable from")
}
