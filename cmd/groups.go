package cmd

import (
	"context"
	"fmt"
	"strings"

	"slices"

	lclient "github.com/uselagoon/machinery/api/lagoon/client"

	"github.com/uselagoon/machinery/api/lagoon"
	"github.com/uselagoon/machinery/api/schema"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/uselagoon/lagoon-cli/pkg/output"
)

var addGroupCmd = &cobra.Command{
	Use:     "group",
	Aliases: []string{"g"},
	Short:   "Add a group to Lagoon, or add a group to an organization",
	Long:    "To add a group to an organization, you'll need to include the `organization` flag and provide the name of the organization. You need to be an owner of this organization to do this.\nIf you're the organization owner and want to grant yourself ownership to this group to be able to deploy projects that may be added to it, specify the `owner` flag, otherwise you will still be able to add and remove users without being an owner",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lContext.Name)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		groupName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Group name", groupName); err != nil {
			return err
		}
		orgOwner, err := cmd.Flags().GetBool("owner")
		if err != nil {
			return err
		}
		organizationName, err := cmd.Flags().GetString("organization-name")
		if err != nil {
			return err
		}

		utoken := lUser.UserConfig.Grant.AccessToken
		lc := lclient.New(
			fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
			lagoonCLIVersion,
			lContext.ContextConfig.Version,
			&utoken,
			debug)

		if organizationName != "" {
			organization, err := lagoon.GetOrganizationByName(context.TODO(), organizationName, lc)
			if err != nil {
				return err
			}
			if organization.Name == "" {
				return fmt.Errorf("error querying organization by name")
			}
			groupInput := schema.AddGroupToOrganizationInput{
				Name:         groupName,
				Organization: organization.ID,
				AddOrgOwner:  orgOwner,
			}
			_, err = lagoon.AddGroupToOrganization(context.TODO(), &groupInput, lc)
			if err != nil {
				return err
			}
		} else {
			groupInput := schema.AddGroupInput{
				Name: groupName,
			}
			_, err = lagoon.AddGroup(context.TODO(), &groupInput, lc)
			if err != nil {
				return err
			}
		}

		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"Group Name": groupName,
			},
		}
		if organizationName != "" {
			resultData.ResultData["Organization"] = organizationName
		}
		output.RenderResult(resultData, outputOptions)
		return nil
	},
}

var addUserToGroupCmd = &cobra.Command{
	Use:     "user-group",
	Aliases: []string{"ug"},
	Short:   "Add a user to a group in lagoon",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lContext.Name)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		groupName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		groupRole, err := cmd.Flags().GetString("role")
		if err != nil {
			return err
		}
		userEmail, err := cmd.Flags().GetString("email")
		if err != nil {
			return err
		}

		cmd.Flags().Visit(
			func(f *pflag.Flag) {
				if f.Name == "role" {
					groupRole = strings.ToUpper(f.Value.String())
				}
			},
		)

		if groupRole == "" {
			// if no role flag is provided, fallback to guest (previous behavior, could be removed though)
			groupRole = "GUEST"
		}

		if groupRole != "" && !slices.Contains(groupRoles, strings.ToLower(groupRole)) {
			return fmt.Errorf("role '%s' is not valid - valid roles include \"guest\", \"reporter\", \"developer\", \"maintainer\", \"owner\"", groupRole)
		}

		if err := requiredInputCheck("Group name", groupName, "Email address", userEmail); err != nil {
			return err
		}

		utoken := lUser.UserConfig.Grant.AccessToken
		lc := lclient.New(
			fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
			lagoonCLIVersion,
			lContext.ContextConfig.Version,
			&utoken,
			debug)

		userGroupRole := &schema.UserGroupRoleInput{
			UserEmail: userEmail,
			GroupName: groupName,
			GroupRole: schema.GroupRole(groupRole),
		}
		_, err = lagoon.AddUserToGroup(context.TODO(), userGroupRole, lc)
		if err != nil {
			return err
		}

		resultData := output.Result{
			Result: "success",
		}
		output.RenderResult(resultData, outputOptions)
		return nil
	},
}

var addProjectToGroupCmd = &cobra.Command{
	Use:     "project-group",
	Aliases: []string{"pg"},
	Short:   "Add a project to a group in lagoon",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lContext.Name)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		groupName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Group name", groupName, "Project name", cmdProjectName); err != nil {
			return err
		}

		projectGroup := &schema.ProjectGroupsInput{
			Project: schema.ProjectInput{
				Name: cmdProjectName,
			},
			Groups: []schema.GroupInput{
				{
					Name: groupName,
				},
			},
		}

		utoken := lUser.UserConfig.Grant.AccessToken
		lc := lclient.New(
			fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
			lagoonCLIVersion,
			lContext.ContextConfig.Version,
			&utoken,
			debug)

		project, err := lagoon.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		if len(project.Name) == 0 {
			outputOptions.Error = fmt.Sprintf("Project '%s' not found", cmdProjectName)
			output.RenderError(outputOptions.Error, outputOptions)
			return nil
		}
		_, err = lagoon.AddProjectToGroup(context.TODO(), projectGroup, lc)
		if err != nil {
			return err
		}

		resultData := output.Result{
			Result: "success",
		}
		output.RenderResult(resultData, outputOptions)
		return nil
	},
}

var deleteUserFromGroupCmd = &cobra.Command{
	Use:     "user-group",
	Aliases: []string{"ug"},
	Short:   "Delete a user from a group in lagoon",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lContext.Name)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		groupName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		userEmail, err := cmd.Flags().GetString("email")
		if err != nil {
			return err
		}

		if err := requiredInputCheck("Group name", groupName, "Email address", userEmail); err != nil {
			return err
		}

		user := &schema.UserGroupInput{
			UserEmail: userEmail,
			GroupName: groupName,
		}

		utoken := lUser.UserConfig.Grant.AccessToken
		lc := lclient.New(
			fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
			lagoonCLIVersion,
			lContext.ContextConfig.Version,
			&utoken,
			debug)

		if yesNo(fmt.Sprintf("You are attempting to delete user '%s' from group '%s', are you sure?", userEmail, groupName)) {
			result, err := lagoon.RemoveUserFromGroup(context.TODO(), user, lc)
			if err != nil {
				return err
			}

			resultData := output.Result{
				Result: "success",
				ResultData: map[string]interface{}{
					"id": result.ID,
				},
			}
			output.RenderResult(resultData, outputOptions)
		}
		return nil
	},
}

var deleteProjectFromGroupCmd = &cobra.Command{
	Use:     "project-group",
	Aliases: []string{"pg"},
	Short:   "Delete a project from a group in lagoon",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lContext.Name)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		groupName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Group name", groupName, "Project name", cmdProjectName); err != nil {
			return err
		}

		projectGroup := &schema.ProjectGroupsInput{
			Project: schema.ProjectInput{
				Name: cmdProjectName,
			},
			Groups: []schema.GroupInput{
				{
					Name: groupName,
				},
			},
		}

		utoken := lUser.UserConfig.Grant.AccessToken
		lc := lclient.New(
			fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
			lagoonCLIVersion,
			lContext.ContextConfig.Version,
			&utoken,
			debug)

		project, err := lagoon.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		if len(project.Name) == 0 {
			outputOptions.Error = fmt.Sprintf("Project '%s' not found", cmdProjectName)
			output.RenderError(outputOptions.Error, outputOptions)
			return nil
		}

		if yesNo(fmt.Sprintf("You are attempting to delete project '%s' from group '%s', are you sure?", projectGroup.Project.Name, projectGroup.Groups[0].Name)) {
			_, err = lagoon.RemoveGroupsFromProject(context.TODO(), projectGroup, lc)
			if err != nil {
				return err
			}

			resultData := output.Result{
				Result: "success",
			}
			output.RenderResult(resultData, outputOptions)
		}
		return nil
	},
}
var deleteGroupCmd = &cobra.Command{
	Use:     "group",
	Aliases: []string{"g"},
	Short:   "Delete a group from lagoon",
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		groupName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Group name", groupName); err != nil {
			return err
		}

		utoken := lUser.UserConfig.Grant.AccessToken
		lc := lclient.New(
			fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
			lagoonCLIVersion,
			lContext.ContextConfig.Version,
			&utoken,
			debug)

		if yesNo(fmt.Sprintf("You are attempting to delete group '%s', are you sure?", groupName)) {
			_, err := lagoon.DeleteGroup(context.TODO(), groupName, lc)
			if err != nil {
				return err
			}
			resultData := output.Result{
				Result: "success",
			}
			output.RenderResult(resultData, outputOptions)
		}
		return nil
	},
}

func init() {
	addGroupCmd.Flags().StringP("name", "N", "", "Name of the group")
	addGroupCmd.Flags().StringP("organization-name", "O", "", "Name of the organization")
	addGroupCmd.Flags().Bool("owner", false, "Organization owner only: Flag to grant yourself ownership of this group")
	addUserToGroupCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group")
	addUserToGroupCmd.Flags().StringP("role", "R", "", "Role in the group [owner, maintainer, developer, reporter, guest]")
	addUserToGroupCmd.Flags().StringP("email", "E", "", "Email address of the user")
	addProjectToGroupCmd.Flags().StringP("name", "N", "", "Name of the group")
	deleteUserFromGroupCmd.Flags().StringP("name", "N", "", "Name of the group")
	deleteUserFromGroupCmd.Flags().StringP("email", "E", "", "Email address of the user")
	deleteProjectFromGroupCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group")
	deleteGroupCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group")

}
