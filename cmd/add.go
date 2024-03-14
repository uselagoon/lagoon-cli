package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/pkg/output"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	s "github.com/uselagoon/machinery/api/schema"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add a project, or add notifications and variables to projects or environments",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateToken(lagoonCLIConfig.Current) // get a new token if the current one is invalid
	},
}

var addNotificationCmd = &cobra.Command{
	Use:     "notification",
	Aliases: []string{"n"},
	Short:   "Add notifications or add notifications to projects",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateToken(lagoonCLIConfig.Current) // get a new token if the current one is invalid
	},
}

var addOrganizationCmd = &cobra.Command{
	Use:     "organization",
	Aliases: []string{"o"},
	Short:   "Add an organization, or add a deploytarget/group/project/user to an organization",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateToken(lagoonCLIConfig.Current) // get a new token if the current one is invalid
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		organizationName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Organization name", organizationName); err != nil {
			return fmt.Errorf("%v | Additional subcommands for deploytarget, group, project & user are available. --help for more information", err)
		}
		organizationFriendlyName, err := cmd.Flags().GetString("friendly-name")
		if err != nil {
			return err
		}
		organizationDescription, err := cmd.Flags().GetString("description")
		if err != nil {
			return err
		}
		organizationQuotaProject, err := cmd.Flags().GetInt("project-quota")
		if err != nil {
			return err
		}
		organizationQuotaGroup, err := cmd.Flags().GetInt("group-quota")
		if err != nil {
			return err
		}
		organizationQuotaNotification, err := cmd.Flags().GetInt("notification-quota")
		if err != nil {
			return err
		}
		organizationQuotaEnvironment, err := cmd.Flags().GetInt("environment-quota")
		if err != nil {
			return err
		}
		organizationQuotaRoute, err := cmd.Flags().GetInt("route-quota")
		if err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)

		organizationInput := s.AddOrganizationInput{
			Name:              organizationName,
			FriendlyName:      organizationFriendlyName,
			Description:       organizationDescription,
			QuotaProject:      organizationQuotaProject,
			QuotaGroup:        organizationQuotaGroup,
			QuotaNotification: organizationQuotaNotification,
			QuotaEnvironment:  organizationQuotaEnvironment,
			QuotaRoute:        organizationQuotaRoute,
		}
		org := s.Organization{}
		err = lc.AddOrganization(context.TODO(), &organizationInput, &org)
		handleError(err)

		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"Organization Name": organizationName,
			},
		}
		output.RenderResult(resultData, outputOptions)
		return nil
	},
}

func init() {
	addCmd.AddCommand(addDeployTargetCmd)
	addCmd.AddCommand(addGroupCmd)
	addCmd.AddCommand(addProjectCmd)
	addCmd.AddCommand(addProjectToGroupCmd)
	addCmd.AddCommand(addNotificationCmd)
	addCmd.AddCommand(addUserCmd)
	addCmd.AddCommand(addOrganizationCmd)
	addCmd.AddCommand(addUserToGroupCmd)
	addCmd.AddCommand(addUserSSHKeyCmd)
	addCmd.AddCommand(addVariableCmd)
	addCmd.AddCommand(addDeployTargetConfigCmd)

	addOrganizationCmd.Flags().StringP("name", "O", "", "Name of the organization")
	addOrganizationCmd.Flags().String("friendly-name", "", "Friendly name of the organization")
	addOrganizationCmd.Flags().String("description", "", "Description of the organization")
	addOrganizationCmd.Flags().Int("project-quota", 0, "Project quota for the organization")
	addOrganizationCmd.Flags().Int("group-quota", 0, "Group quota for the organization")
	addOrganizationCmd.Flags().Int("notification-quota", 0, "Notification quota for the organization")
	addOrganizationCmd.Flags().Int("environment-quota", 0, "Environment quota for the organization")
	addOrganizationCmd.Flags().Int("route-quota", 0, "Route quota for the organization")
}
