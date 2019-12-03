package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/lagoon/users"
	"github.com/amazeeio/lagoon-cli/output"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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
	Short:   "Add group to lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		groupFlags := parseGroup(*cmd.Flags())
		if groupFlags.Name == "" {
			fmt.Println("Not enough arguments. Requires: group name")
			cmd.Help()
			os.Exit(1)
		}
		var customReqResult []byte
		var err error
		customReqResult, err = users.AddGroup(groupFlags)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		returnResultData := map[string]interface{}{}
		err = json.Unmarshal([]byte(customReqResult), &returnResultData)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
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
	Short:   "Add user to a group in lagoon",
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
				Email: userEmail,
			},
			Group: api.Group{
				Name: groupName,
			},
			Role: roleType,
		}
		if userGroupRole.User.Email == "" || userGroupRole.Group.Name == "" || userGroupRole.Role == "" {
			output.RenderError("Must define an email address", outputOptions)
			cmd.Help()
			os.Exit(1)
		}
		var customReqResult []byte
		var err error
		customReqResult, err = users.AddUserToGroup(userGroupRole)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		returnResultData := map[string]interface{}{}
		err = json.Unmarshal([]byte(customReqResult), &returnResultData)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
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
	Short:   "Add project to a group in lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		projectGroup := api.ProjectGroups{
			Project: api.Project{
				Name: cmdProjectName,
			},
			Groups: []api.Group{
				api.Group{
					Name: groupName,
				},
			},
		}
		if projectGroup.Project.Name == "" || len(projectGroup.Groups) == 0 {
			output.RenderError("Must define a project name and group", outputOptions)
			cmd.Help()
			os.Exit(1)
		}
		var customReqResult []byte
		var err error
		customReqResult, err = users.AddProjectToGroup(projectGroup)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		returnResultData := map[string]interface{}{}
		err = json.Unmarshal([]byte(customReqResult), &returnResultData)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		resultData := output.Result{
			Result:     "success",
			ResultData: returnResultData,
		}
		output.RenderResult(resultData, outputOptions)
	},
}

var delUserFromGroupCmd = &cobra.Command{
	Use:     "user-group",
	Aliases: []string{"ug"},
	Short:   "Delete user from a group in lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		userGroupRole := api.UserGroup{
			User: api.User{
				Email: userEmail,
			},
			Group: api.Group{
				Name: groupName,
			},
		}
		if userGroupRole.User.Email == "" || userGroupRole.Group.Name == "" {
			output.RenderError("Must define an email address and group name", outputOptions)
			cmd.Help()
			os.Exit(1)
		}
		var customReqResult []byte
		var err error
		customReqResult, err = users.RemoveUserFromGroup(userGroupRole)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		returnResultData := map[string]interface{}{}
		err = json.Unmarshal([]byte(customReqResult), &returnResultData)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		resultData := output.Result{
			Result:     "success",
			ResultData: returnResultData,
		}
		output.RenderResult(resultData, outputOptions)
	},
}

var delProjectFromGroupCmd = &cobra.Command{
	Use:     "project-group",
	Aliases: []string{"pg"},
	Short:   "Delete project from a group in lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		projectGroup := api.ProjectGroups{
			Project: api.Project{
				Name: cmdProjectName,
			},
			Groups: []api.Group{
				api.Group{
					Name: groupName,
				},
			},
		}
		if projectGroup.Project.Name == "" || len(projectGroup.Groups) == 0 {
			output.RenderError("Must define a project name and group", outputOptions)
			cmd.Help()
			os.Exit(1)
		}
		var customReqResult []byte
		var err error
		customReqResult, err = users.RemoveGroupsFromProject(projectGroup)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		returnResultData := map[string]interface{}{}
		err = json.Unmarshal([]byte(customReqResult), &returnResultData)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		resultData := output.Result{
			Result:     "success",
			ResultData: returnResultData,
		}
		output.RenderResult(resultData, outputOptions)
	},
}
var delGroupCmd = &cobra.Command{
	Use:     "group",
	Aliases: []string{"g"},
	Short:   "Delete group from lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		groupFlags := parseGroup(*cmd.Flags())
		if groupFlags.Name == "" {
			fmt.Println("Not enough arguments. Requires: group name")
			cmd.Help()
			os.Exit(1)
		}
		var customReqResult []byte
		var err error
		customReqResult, err = users.DeleteGroup(groupFlags)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		resultData := output.Result{
			Result: string(customReqResult),
		}
		output.RenderResult(resultData, outputOptions)
	},
}

func init() {
	addGroupCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group")
	addUserToGroupCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group")
	addUserToGroupCmd.Flags().StringVarP(&groupRole, "role", "R", "", "Role in the group [owner, maintainer, developer, reporter, guest]")
	addUserToGroupCmd.Flags().StringVarP(&userEmail, "email", "E", "", "Email address of the user")
	addProjectToGroupCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group")
	delUserFromGroupCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group")
	delUserFromGroupCmd.Flags().StringVarP(&userEmail, "email", "E", "", "Email address of the user")
	delProjectFromGroupCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group")
	delGroupCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group")
}
