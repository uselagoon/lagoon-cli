package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/pkg/output"
	l "github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	s "github.com/uselagoon/machinery/api/schema"
)

var addOrgCmd = &cobra.Command{
	Use:     "organization",
	Aliases: []string{"o"},
	Short:   "Add a new organization to Lagoon",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
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
			return err
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

var deleteOrgCmd = &cobra.Command{
	Use:     "organization",
	Aliases: []string{"o"},
	Short:   "Delete an organization",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
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
			return err
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)

		organization, err := l.GetOrganizationByName(context.TODO(), organizationName, lc)
		handleError(err)
		if yesNo(fmt.Sprintf("You are attempting to delete organization '%s', are you sure?", organization.Name)) {
			_, err := l.DeleteOrganization(context.TODO(), organization.ID, lc)
			handleError(err)
			resultData := output.Result{
				Result: organization.Name,
			}
			output.RenderResult(resultData, outputOptions)
		}
		return nil
	},
}

var updateOrganizationCmd = &cobra.Command{
	Use:     "organization",
	Aliases: []string{"o"},
	Short:   "Update an organization",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
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
			return err
		}
		if err != nil {
			return err
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

		organization, err := l.GetOrganizationByName(context.TODO(), organizationName, lc)
		organizationInput := s.UpdateOrganizationPatchInput{
			Description:       nullStrCheck(organizationDescription),
			FriendlyName:      nullStrCheck(organizationFriendlyName),
			QuotaProject:      nullIntCheck(organizationQuotaProject),
			QuotaGroup:        nullIntCheck(organizationQuotaGroup),
			QuotaNotification: nullIntCheck(organizationQuotaNotification),
			QuotaEnvironment:  nullIntCheck(organizationQuotaEnvironment),
			QuotaRoute:        nullIntCheck(organizationQuotaRoute),
		}
		result, err := l.UpdateOrganization(context.TODO(), organization.ID, organizationInput, lc)
		handleError(err)

		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"Organization Name": result.Name,
			},
		}
		output.RenderResult(resultData, outputOptions)
		return nil
	},
}

func init() {
	addOrganizationCmd.AddCommand(addOrgCmd)
	addOrganizationCmd.AddCommand(addGroupToOrganizationCmd)
	addOrganizationCmd.AddCommand(addProjectToOrganizationCmd)
	addOrganizationCmd.AddCommand(addDeployTargetToOrganizationCmd)
	addOrganizationCmd.AddCommand(addUserToOrganizationCmd)

	deleteOrganizationCmd.AddCommand(deleteOrgCmd)
	deleteOrganizationCmd.AddCommand(RemoveDeployTargetFromOrganizationCmd)
	deleteOrganizationCmd.AddCommand(RemoveProjectFromOrganizationCmd)
	deleteOrganizationCmd.AddCommand(RemoveUserFromOrganization)

	addOrgCmd.Flags().StringP("name", "O", "", "Name of the organization")
	addOrgCmd.Flags().String("friendly-name", "", "Friendly name of the organization")
	addOrgCmd.Flags().String("description", "", "Description of the organization")
	addOrgCmd.Flags().Int("project-quota", 0, "Project quota for the organization")
	addOrgCmd.Flags().Int("group-quota", 0, "Group quota for the organization")
	addOrgCmd.Flags().Int("notification-quota", 0, "Notification quota for the organization")
	addOrgCmd.Flags().Int("environment-quota", 0, "Environment quota for the organization")
	addOrgCmd.Flags().Int("route-quota", 0, "Route quota for the organization")

	updateOrganizationCmd.Flags().StringP("name", "O", "", "Name of the organization to update")
	updateOrganizationCmd.Flags().String("friendly-name", "", "Friendly name of the organization")
	updateOrganizationCmd.Flags().String("description", "", "Description of the organization")
	updateOrganizationCmd.Flags().Int("project-quota", 0, "Project quota for the organization")
	updateOrganizationCmd.Flags().Int("group-quota", 0, "Group quota for the organization")
	updateOrganizationCmd.Flags().Int("notification-quota", 0, "Notification quota for the organization")
	updateOrganizationCmd.Flags().Int("environment-quota", 0, "Environment quota for the organization")
	updateOrganizationCmd.Flags().Int("route-quota", 0, "Route quota for the organization")

	deleteOrgCmd.Flags().StringP("name", "O", "", "Name of the organization to delete")
}
