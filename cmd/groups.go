package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	l "github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	ls "github.com/uselagoon/machinery/api/schema"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/uselagoon/lagoon-cli/pkg/api"
	"github.com/uselagoon/lagoon-cli/pkg/output"
)

func parseGroup(flags pflag.FlagSet) api.Group {
	configMap := make(map[string]interface{})
	flags.VisitAll(func(f *pflag.Flag) {
		if flags.Changed(f.Name) {
			configMap[f.Name] = f.Value
		}
	})
	jsonStr, _ := json.Marshal(configMap)
	parsedFlags := api.Group{}
	json.Unmarshal(jsonStr, &parsedFlags)
	return parsedFlags
}

var addGroupCmd = &cobra.Command{
	Use:     "group",
	Aliases: []string{"g"},
	Short:   "Add a group to lagoon",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		groupName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Group name", groupName); err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)

		grp := &ls.AddGroupInput{Name: groupName}
		group, err := l.AddGroup(context.TODO(), grp, lc)
		handleError(err)

		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"Group Name": group.Name,
			},
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
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
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

		var roleType ls.GroupRole
		roleType = ls.GuestRole
		switch strings.ToLower(groupRole) {
		case "guest":
			roleType = ls.GuestRole
		case "reporter":
			roleType = ls.ReporterRole
		case "developer":
			roleType = ls.DeveloperRole
		case "maintainer":
			roleType = ls.MaintainerRole
		case "owner":
			roleType = ls.OwnerRole
		}

		if err := requiredInputCheck("Group name", groupName, "Email address", userEmail, "Role", string(roleType)); err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)

		userGroupRole := &ls.UserGroupRoleInput{
			UserEmail: userEmail,
			GroupName: groupName,
			GroupRole: roleType,
		}
		_, err = l.AddUserToGroup(context.TODO(), userGroupRole, lc)
		handleError(err)

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
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		groupName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Group name", groupName, "Project name", cmdProjectName); err != nil {
			return err
		}

		projectGroup := &ls.ProjectGroupsInput{
			Project: ls.ProjectInput{
				Name: cmdProjectName,
			},
			Groups: []ls.GroupInput{
				{
					Name: groupName,
				},
			},
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)

		project, err := l.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
		if len(project.Name) == 0 {
			outputOptions.Error = fmt.Sprintf("Project '%s' not found", cmdProjectName)
			output.RenderError(outputOptions.Error, outputOptions)
			return nil
		}
		_, err = l.AddProjectToGroup(context.TODO(), projectGroup, lc)
		handleError(err)

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
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
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

		user := &ls.UserGroupInput{
			UserEmail: userEmail,
			GroupName: groupName,
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)

		if yesNo(fmt.Sprintf("You are attempting to delete user '%s' from group '%s', are you sure?", userEmail, groupName)) {
			result, err := l.RemoveUserFromGroup(context.TODO(), user, lc)
			handleError(err)

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
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		groupName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Group name", groupName, "Project name", cmdProjectName); err != nil {
			return err
		}

		projectGroup := &ls.ProjectGroupsInput{
			Project: ls.ProjectInput{
				Name: cmdProjectName,
			},
			Groups: []ls.GroupInput{
				{
					Name: groupName,
				},
			},
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)

		project, err := l.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
		if len(project.Name) == 0 {
			outputOptions.Error = fmt.Sprintf("Project '%s' not found", cmdProjectName)
			output.RenderError(outputOptions.Error, outputOptions)
			return nil
		}

		if yesNo(fmt.Sprintf("You are attempting to delete project '%s' from group '%s', are you sure?", projectGroup.Project.Name, projectGroup.Groups[0].Name)) {
			_, err = l.RemoveGroupsFromProject(context.TODO(), projectGroup, lc)
			handleError(err)

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
		groupName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Group name", groupName); err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)

		if yesNo(fmt.Sprintf("You are attempting to delete group '%s', are you sure?", groupName)) {
			_, err := l.DeleteGroup(context.TODO(), groupName, lc)
			handleError(err)
			resultData := output.Result{
				Result: "success",
			}
			output.RenderResult(resultData, outputOptions)
		}
		return nil
	},
}

var addGroupToOrganizationCmd = &cobra.Command{
	Use:     "group",
	Aliases: []string{"g"},
	Short:   "Add a group to an Organization",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		handleError(err)
		orgOwner, err := cmd.Flags().GetBool("org-owner")
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
		groupName, err := cmd.Flags().GetString("group")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Group name", groupName); err != nil {
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

		groupInput := ls.AddGroupToOrganizationInput{
			Name:         groupName,
			Organization: organization.ID,
			AddOrgOwner:  orgOwner,
		}
		group := ls.OrgGroup{}
		err = lc.AddGroupToOrganization(context.TODO(), &groupInput, &group)
		handleError(err)

		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"Group Name":        group.Name,
				"Organization Name": organizationName,
			},
		}
		output.RenderResult(resultData, outputOptions)
		return nil
	},
}

func init() {
	addGroupCmd.Flags().StringP("name", "N", "", "Name of the group")
	addUserToGroupCmd.Flags().StringP("name", "N", "", "Name of the group")
	addUserToGroupCmd.Flags().StringP("role", "R", "", "Role in the group [owner, maintainer, developer, reporter, guest]")
	addUserToGroupCmd.Flags().StringP("email", "E", "", "Email address of the user")
	addProjectToGroupCmd.Flags().StringP("name", "N", "", "Name of the group")
	deleteUserFromGroupCmd.Flags().StringP("name", "N", "", "Name of the group")
	deleteUserFromGroupCmd.Flags().StringP("email", "E", "", "Email address of the user")
	deleteProjectFromGroupCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group")
	deleteGroupCmd.Flags().StringP("name", "N", "", "Name of the group")
	addGroupToOrganizationCmd.Flags().StringP("name", "O", "", "Name of the organization")
	addGroupToOrganizationCmd.Flags().StringP("group", "G", "", "Name of the group")
	addGroupToOrganizationCmd.Flags().Bool("org-owner", false, "Flag to add the user to the group as an owner")
}
