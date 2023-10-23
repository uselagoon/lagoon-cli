package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	l "github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	s "github.com/uselagoon/machinery/api/schema"
	"os"
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
	Run: func(cmd *cobra.Command, args []string) {
		groupFlags := parseGroup(*cmd.Flags())
		if groupFlags.Name == "" {
			fmt.Println("Missing arguments: Group name is not defined")
			cmd.Help()
			os.Exit(1)
		}
		var customReqResult []byte
		var err error
		customReqResult, err = uClient.AddGroup(groupFlags)
		handleError(err)
		returnResultData := map[string]interface{}{}
		err = json.Unmarshal([]byte(customReqResult), &returnResultData)
		handleError(err)
		resultData := output.Result{
			Result:     "success",
			ResultData: returnResultData,
		}
		output.RenderResult(resultData, outputOptions)
	},
}

var addUserToGroupCmd = &cobra.Command{
	Use:     "user-group",
	Aliases: []string{"ug"},
	Short:   "Add a user to a group in lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		var roleType api.GroupRole
		roleType = api.GuestRole
		if strings.EqualFold(string(groupRole), "guest") {
			roleType = api.GuestRole
		} else if strings.EqualFold(string(groupRole), "reporter") {
			roleType = api.ReporterRole
		} else if strings.EqualFold(string(groupRole), "developer") {
			roleType = api.DeveloperRole
		} else if strings.EqualFold(string(groupRole), "maintainer") {
			roleType = api.MaintainerRole
		} else if strings.EqualFold(string(groupRole), "owner") {
			roleType = api.OwnerRole
		}
		userGroupRole := api.UserGroupRole{
			User: api.User{
				Email: strings.ToLower(userEmail),
			},
			Group: api.Group{
				Name: groupName,
			},
			Role: roleType,
		}
		if userGroupRole.User.Email == "" || userGroupRole.Group.Name == "" || userGroupRole.Role == "" {
			output.RenderError("Missing arguments: Email address, group name, or role is not defined", outputOptions)
			cmd.Help()
			os.Exit(1)
		}
		var customReqResult []byte
		var err error
		customReqResult, err = uClient.AddUserToGroup(userGroupRole)
		handleError(err)
		returnResultData := map[string]interface{}{}
		err = json.Unmarshal([]byte(customReqResult), &returnResultData)
		handleError(err)
		resultData := output.Result{
			Result:     "success",
			ResultData: returnResultData,
		}
		output.RenderResult(resultData, outputOptions)
	},
}

var addProjectToGroupCmd = &cobra.Command{
	Use:     "project-group",
	Aliases: []string{"pg"},
	Short:   "Add a project to a group in lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		projectGroup := api.ProjectGroups{
			Project: api.Project{
				Name: cmdProjectName,
			},
			Groups: []api.Group{
				{
					Name: groupName,
				},
			},
		}
		if projectGroup.Project.Name == "" || len(projectGroup.Groups) == 0 {
			output.RenderError("Missing arguments: Project name or group name is not defined", outputOptions)
			cmd.Help()
			os.Exit(1)
		}
		var customReqResult []byte
		var err error
		customReqResult, err = uClient.AddProjectToGroup(projectGroup)
		handleError(err)
		returnResultData := map[string]interface{}{}
		err = json.Unmarshal([]byte(customReqResult), &returnResultData)
		handleError(err)
		resultData := output.Result{
			Result:     "success",
			ResultData: returnResultData,
		}
		output.RenderResult(resultData, outputOptions)
	},
}

var deleteUserFromGroupCmd = &cobra.Command{
	Use:     "user-group",
	Aliases: []string{"ug"},
	Short:   "Delete a user from a group in lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		userGroupRole := api.UserGroup{
			User: api.User{
				Email: strings.ToLower(userEmail),
			},
			Group: api.Group{
				Name: groupName,
			},
		}
		if userGroupRole.User.Email == "" || userGroupRole.Group.Name == "" {
			output.RenderError("Missing arguments: Email address or group name is not defined", outputOptions)
			cmd.Help()
			os.Exit(1)
		}
		var customReqResult []byte
		var err error
		if yesNo(fmt.Sprintf("You are attempting to delete user '%s' from group '%s', are you sure?", userGroupRole.User.Email, userGroupRole.Group.Name)) {
			customReqResult, err = uClient.RemoveUserFromGroup(userGroupRole)
			handleError(err)
			returnResultData := map[string]interface{}{}
			err = json.Unmarshal([]byte(customReqResult), &returnResultData)
			handleError(err)
			resultData := output.Result{
				Result:     "success",
				ResultData: returnResultData,
			}
			output.RenderResult(resultData, outputOptions)
		}
	},
}

var deleteProjectFromGroupCmd = &cobra.Command{
	Use:     "project-group",
	Aliases: []string{"pg"},
	Short:   "Delete a project from a group in lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		projectGroup := api.ProjectGroups{
			Project: api.Project{
				Name: cmdProjectName,
			},
			Groups: []api.Group{
				{
					Name: groupName,
				},
			},
		}
		if projectGroup.Project.Name == "" || len(projectGroup.Groups) == 0 {
			output.RenderError("Missing arguments: Project name or group name is not defined", outputOptions)
			cmd.Help()
			os.Exit(1)
		}
		var customReqResult []byte
		var err error
		if yesNo(fmt.Sprintf("You are attempting to delete project '%s' from group '%s', are you sure?", projectGroup.Project.Name, projectGroup.Groups[0].Name)) {
			customReqResult, err = uClient.RemoveGroupsFromProject(projectGroup)
			handleError(err)
			returnResultData := map[string]interface{}{}
			err = json.Unmarshal([]byte(customReqResult), &returnResultData)
			handleError(err)
			resultData := output.Result{
				Result:     "success",
				ResultData: returnResultData,
			}
			output.RenderResult(resultData, outputOptions)
		}
	},
}
var deleteGroupCmd = &cobra.Command{
	Use:     "group",
	Aliases: []string{"g"},
	Short:   "Delete a group from lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		groupFlags := parseGroup(*cmd.Flags())
		if groupFlags.Name == "" {
			fmt.Println("Missing arguments: Group name is not defined")
			cmd.Help()
			os.Exit(1)
		}
		var customReqResult []byte
		var err error
		if yesNo(fmt.Sprintf("You are attempting to delete group '%s', are you sure?", groupFlags.Name)) {
			customReqResult, err = uClient.DeleteGroup(groupFlags)
			handleError(err)
			resultData := output.Result{
				Result: string(customReqResult),
			}
			output.RenderResult(resultData, outputOptions)
		}
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

		groupInput := s.AddGroupToOrganizationInput{
			Name:         groupName,
			Organization: organization.ID,
			AddOrgOwner:  orgOwner,
		}
		group := s.OrgGroup{}
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
	addGroupCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group")
	addUserToGroupCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group")
	addUserToGroupCmd.Flags().StringVarP(&groupRole, "role", "R", "", "Role in the group [owner, maintainer, developer, reporter, guest]")
	addUserToGroupCmd.Flags().StringVarP(&userEmail, "email", "E", "", "Email address of the user")
	addProjectToGroupCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group")
	deleteUserFromGroupCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group")
	deleteUserFromGroupCmd.Flags().StringVarP(&userEmail, "email", "E", "", "Email address of the user")
	deleteProjectFromGroupCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group")
	deleteGroupCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group")
	addGroupToOrganizationCmd.Flags().StringP("name", "O", "", "Name of the organization")
	addGroupToOrganizationCmd.Flags().StringP("group", "G", "", "Name of the group")
	addGroupToOrganizationCmd.Flags().Bool("org-owner", false, "Flag to add the user to the group as an owner")
}
