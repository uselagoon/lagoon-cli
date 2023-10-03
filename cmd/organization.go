package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/pkg/output"
	l "github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	s "github.com/uselagoon/machinery/api/schema"
	"os"
)

// TODO
var addOrganizationCmd = &cobra.Command{
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
		organizationFriendlyName, err := cmd.Flags().GetString("friendlyName")
		if err != nil {
			return err
		}
		organizationDescription, err := cmd.Flags().GetString("description")
		if err != nil {
			return err
		}
		organizationQuotaProject, err := cmd.Flags().GetUint("quotaProject")
		if err != nil {
			return err
		}
		organizationQuotaGroup, err := cmd.Flags().GetUint("quotaGroup")
		if err != nil {
			return err
		}
		organizationQuotaNotification, err := cmd.Flags().GetUint("quotaNotification")
		if err != nil {
			return err
		}
		organizationQuotaEnvironment, err := cmd.Flags().GetUint("quotaEnvironment")
		if err != nil {
			return err
		}
		organizationQuotaRoute, err := cmd.Flags().GetUint("quotaRoute")
		if err != nil {
			return err
		}

		if organizationName == "" {
			fmt.Println("Missing arguments: Organization name is not defined")
			cmd.Help()
			os.Exit(1)
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

var deleteOrganizationCmd = &cobra.Command{
	Use:     "organization",
	Aliases: []string{"o"},
	Short:   "Delete an organization",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		organizationId, err := cmd.Flags().GetUint("id")
		if err != nil {
			return err
		}
		if organizationId == 0 {
			fmt.Println("Missing arguments: Organization ID is not defined")
			cmd.Help()
			os.Exit(1)
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)

		organization, err := l.GetOrganizationByID(context.TODO(), organizationId, lc)
		handleError(err)
		if yesNo(fmt.Sprintf("You are attempting to delete organization '%s', are you sure?", organization.Name)) {
			_, err := l.DeleteOrganization(context.TODO(), organizationId, lc)
			handleError(err)
			resultData := output.Result{
				Result: organization.Name,
			}
			output.RenderResult(resultData, outputOptions)
		}
		return nil
	},
}

// TODO - update once the API is updated
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
		organizationId, err := cmd.Flags().GetUint("id")
		if err != nil {
			return err
		}
		organizationFriendlyName, err := cmd.Flags().GetString("friendlyName")
		if err != nil {
			return err
		}
		organizationDescription, err := cmd.Flags().GetString("description")
		if err != nil {
			return err
		}
		organizationQuotaProject, err := cmd.Flags().GetUint("quotaProject")
		if err != nil {
			return err
		}
		organizationQuotaGroup, err := cmd.Flags().GetUint("quotaGroup")
		if err != nil {
			return err
		}
		organizationQuotaNotification, err := cmd.Flags().GetUint("quotaNotification")
		if err != nil {
			return err
		}
		organizationQuotaEnvironment, err := cmd.Flags().GetUint("quotaEnvironment")
		if err != nil {
			return err
		}
		organizationQuotaRoute, err := cmd.Flags().GetUint("quotaRoute")
		if err != nil {
			return err
		}

		if organizationId == 0 {
			fmt.Println("Missing arguments: Organization ID is not defined")
			cmd.Help()
			os.Exit(1)
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)

		organizationInput := s.UpdateOrganizationPatchInput{
			Description:       nullStrCheck(organizationDescription),
			FriendlyName:      nullStrCheck(organizationFriendlyName),
			QuotaProject:      nullUintCheck(organizationQuotaProject),
			QuotaGroup:        nullUintCheck(organizationQuotaGroup),
			QuotaNotification: nullUintCheck(organizationQuotaNotification),
			QuotaEnvironment:  nullUintCheck(organizationQuotaEnvironment),
			QuotaRoute:        nullUintCheck(organizationQuotaRoute),
		}
		result, err := l.UpdateOrganization(context.TODO(), organizationId, organizationInput, lc)
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
	addOrganizationCmd.Flags().String("name", "", "Name of the organization")
	addOrganizationCmd.Flags().String("friendlyName", "", "Friendly name of the organization")
	addOrganizationCmd.Flags().String("description", "", "Description of the organization")
	addOrganizationCmd.Flags().Uint("quotaProject", 0, "Project quota for the organization")
	addOrganizationCmd.Flags().Uint("quotaGroup", 0, "Group quota for the organization")
	addOrganizationCmd.Flags().Uint("quotaNotification", 0, "Notification quota for the organization")
	addOrganizationCmd.Flags().Uint("quotaEnvironment", 0, "Environment quota for the organization")
	addOrganizationCmd.Flags().Uint("quotaRoute", 0, "Route quota for the organization")

	updateOrganizationCmd.Flags().Uint("id", 0, "ID of the organization to update")
	updateOrganizationCmd.Flags().String("friendlyName", "", "Friendly name of the organization")
	updateOrganizationCmd.Flags().String("description", "", "Description of the organization")
	updateOrganizationCmd.Flags().Uint("quotaProject", 0, "Project quota for the organization")
	updateOrganizationCmd.Flags().Uint("quotaGroup", 0, "Group quota for the organization")
	updateOrganizationCmd.Flags().Uint("quotaNotification", 0, "Notification quota for the organization")
	updateOrganizationCmd.Flags().Uint("quotaEnvironment", 0, "Environment quota for the organization")
	updateOrganizationCmd.Flags().Uint("quotaRoute", 0, "Route quota for the organization")

	deleteOrganizationCmd.Flags().Uint("id", 0, "ID of the organization")
}
