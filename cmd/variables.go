package cmd

import (
	"context"
	"fmt"
	l "github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	ls "github.com/uselagoon/machinery/api/schema"
	"strings"

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
	RunE: func(cmd *cobra.Command, args []string) error {
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
		if err := requiredInputCheck("Project name", cmdProjectName, "Variable name", varName); err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)

		in := &ls.EnvVariableByNameInput{
			Project:     cmdProjectName,
			Environment: cmdProjectEnvironment,
			Scope:       ls.EnvVariableScope(strings.ToUpper(varScope)),
			Name:        varName,
			Value:       varValue,
		}
		envvar, err := l.AddOrUpdateEnvVariableByName(context.TODO(), in, lc)
		if err != nil {
			return err
		}

		if envvar.ID != 0 {
			data := []output.Data{}
			env := []string{
				returnNonEmptyString(fmt.Sprintf("%v", envvar.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", cmdProjectName)),
			}
			if cmdProjectEnvironment != "" {
				env = append(env, returnNonEmptyString(fmt.Sprintf("%v", cmdProjectEnvironment)))
			}
			env = append(env, returnNonEmptyString(fmt.Sprintf("%v", envvar.Scope)))
			env = append(env, returnNonEmptyString(fmt.Sprintf("%v", envvar.Name)))
			env = append(env, returnNonEmptyString(fmt.Sprintf("%v", envvar.Value)))
			data = append(data, env)
			header := []string{
				"ID",
				"Project",
			}
			if cmdProjectEnvironment != "" {
				header = append(header, "Environment")
			}
			header = append(header, "Scope")
			header = append(header, "Name")
			header = append(header, "Value")
			output.RenderOutput(output.Table{
				Header: header,
				Data:   data,
			}, outputOptions)
		} else {
			output.RenderInfo(fmt.Sprintf("variable %s remained unchanged", varName), outputOptions)
		}
		return nil
	},
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
		if err := requiredInputCheck("Project name", cmdProjectName, "Variable name", varName); err != nil {
			return err
		}

		deleteMsg := fmt.Sprintf("You are attempting to delete variable '%s' from project '%s', are you sure?", varName, cmdProjectName)
		if cmdProjectEnvironment != "" {
			deleteMsg = fmt.Sprintf("You are attempting to delete variable '%s' from environment '%s' in project '%s', are you sure?", varName, cmdProjectEnvironment, cmdProjectName)
		}
		if yesNo(deleteMsg) {
			current := lagoonCLIConfig.Current
			token := lagoonCLIConfig.Lagoons[current].Token
			lc := lclient.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIVersion,
				&token,
				debug)
			in := &ls.DeleteEnvVariableByNameInput{
				Project:     cmdProjectName,
				Environment: cmdProjectEnvironment,
				Name:        varName,
			}
			deleteResult, err := l.DeleteEnvVariableByName(context.TODO(), in, lc)
			if err != nil {
				return err
			}
			resultData := output.Result{
				Result: deleteResult.DeleteEnvVar,
			}
			output.RenderResult(resultData, outputOptions)
		}
		return nil
	},
}

var updateVariableCmd = addVariableCmd

func init() {
	updateCmd.AddCommand(updateVariableCmd)
	addVariableCmd.Flags().StringP("name", "N", "", "Name of the variable to add")
	addVariableCmd.Flags().StringP("value", "V", "", "Value of the variable to add")
	addVariableCmd.Flags().StringP("scope", "S", "", "Scope of the variable[global, build, runtime, container_registry, internal_container_registry]")
	deleteVariableCmd.Flags().StringP("name", "N", "", "Name of the variable to delete")
}
