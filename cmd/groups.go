package cmd

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/lagoon/users"
	"github.com/amazeeio/lagoon-cli/output"
	"github.com/spf13/cobra"
)

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
		if userGroupRole.User.Email == "" && userGroupRole.Group.Name == "" && userGroupRole.Role == "" {
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
		if projectGroup.Project.Name == "" && len(projectGroup.Groups) == 0 {
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

func init() {
	addUserToGroupCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group")
	addUserToGroupCmd.Flags().StringVarP(&groupRole, "role", "R", "", "Role in the group [owner, maintainer, developer, reporter, guest]")
	addUserToGroupCmd.Flags().StringVarP(&userEmail, "email", "E", "", "Email address of the user")
	addProjectToGroupCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group")
}
