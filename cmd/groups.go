package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	l "github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	s "github.com/uselagoon/machinery/api/schema"

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
		handleError(err)
		groupName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Group name", groupName); err != nil {
			return err
		}
		orgOwner, err := cmd.Flags().GetBool("org-owner")
		if err != nil {
			return err
		}
		organizationName, err := cmd.Flags().GetString("organization")
		if err != nil {
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

		if organizationName != "" {
			organization, err := l.GetOrganizationByName(context.TODO(), organizationName, lc)
			handleError(err)
			groupInput := s.AddGroupToOrganizationInput{
				Name:         groupName,
				Organization: organization.ID,
				AddOrgOwner:  orgOwner,
			}
			_, err = l.AddGroupToOrganization(context.TODO(), &groupInput, lc)
			handleError(err)
		} else {
			groupInput := s.AddGroupInput{
				Name: groupName,
			}
			_, err = l.AddGroup(context.TODO(), &groupInput, lc)
			handleError(err)
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
	RunE: func(cmd *cobra.Command, args []string) error {
		gRole, err := cmd.Flags().GetString("role")
		if err != nil {
			return err
		}
		cmd.Flags().Visit(
			func(f *pflag.Flag) {
				if f.Name == "role" {
					gRole = strings.ToUpper(f.Value.String())
				}
			},
		)
		if gRole == "" {
			// if no role flag is provided, fallback to guest (previous behavior, could be removed though)
			gRole = "GUEST"
		}
		userGroupRole := api.UserGroupRole{
			User: api.User{
				Email: strings.ToLower(userEmail),
			},
			Group: api.Group{
				Name: groupName,
			},
			Role: api.GroupRole(gRole),
		}
		if userGroupRole.User.Email == "" || userGroupRole.Group.Name == "" || userGroupRole.Role == "" {
			return fmt.Errorf("missing arguments: Email address, group name, or role is not defined")
		}
		var customReqResult []byte
		customReqResult, err = uClient.AddUserToGroup(userGroupRole)
		if err != nil {
			return err
		}
		returnResultData := map[string]interface{}{}
		if err = json.Unmarshal([]byte(customReqResult), &returnResultData); err != nil {
			return err
		}
		resultData := output.Result{
			Result:     "success",
			ResultData: returnResultData,
		}
		output.RenderResult(resultData, outputOptions)
		return nil
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

func init() {
	addGroupCmd.Flags().StringP("name", "N", "", "Name of the group")
	addGroupCmd.Flags().StringP("organization", "O", "", "Name of the organization")
	addGroupCmd.Flags().Bool("org-owner", false, "Flag to add the user to the group as an owner")
	addUserToGroupCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group")
	addUserToGroupCmd.Flags().StringP("role", "R", "", "Role in the group [owner, maintainer, developer, reporter, guest]")
	addUserToGroupCmd.Flags().StringVarP(&userEmail, "email", "E", "", "Email address of the user")
	addProjectToGroupCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group")
	deleteUserFromGroupCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group")
	deleteUserFromGroupCmd.Flags().StringVarP(&userEmail, "email", "E", "", "Email address of the user")
	deleteProjectFromGroupCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group")
	deleteGroupCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group")

}
