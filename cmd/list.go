package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/uselagoon/lagoon-cli/pkg/output"
	l "github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	ls "github.com/uselagoon/machinery/api/schema"
	"strconv"
)

// ListFlags .
type ListFlags struct {
	Project     string `json:"project,omitempty"`
	Environment string `json:"environment,omitempty"`
	Reveal      bool   `json:"reveal,omitempty"`
}

func parseListFlags(flags pflag.FlagSet) ListFlags {
	configMap := make(map[string]interface{})
	flags.VisitAll(func(f *pflag.Flag) {
		if flags.Changed(f.Name) {
			configMap[f.Name] = f.Value
		}
	})
	jsonStr, _ := json.Marshal(configMap)
	parsedFlags := ListFlags{}
	json.Unmarshal(jsonStr, &parsedFlags)
	return parsedFlags
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects, environments, deployments, variables or notifications",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateToken(lagoonCLIConfig.Current) // get a new token if the current one is invalid
	},
}

var listProjectsCmd = &cobra.Command{
	Use:     "projects",
	Aliases: []string{"p"},
	Short:   "List all projects you have access to (alias: p)",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
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

		projects, err := l.ListAllProjects(context.TODO(), lc)
		handleError(err)

		data := []output.Data{}
		for _, project := range *projects {
			var devEnvironments = 0
			for _, environment := range project.Environments {
				if environment.EnvironmentType == "development" {
					devEnvironments++
				}
			}

			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%d", project.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", project.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", project.GitURL)),
				returnNonEmptyString(fmt.Sprintf("%v", project.ProductionEnvironment)),
				returnNonEmptyString(fmt.Sprintf("%v/%v", devEnvironments, project.DevelopmentEnvironmentsLimit)),
			})
		}
		if len(data) == 0 {
			outputOptions.Error = "No access to any projects in Lagoon\n"
		}
		dataMain := output.Table{
			Header: []string{"ID", "ProjectName", "GitUrl", "ProductionEnvironment", "DevEnvironments"},
			Data:   data,
		}

		output.RenderOutput(dataMain, outputOptions)
		return nil
	},
}

var listDeployTargetsCmd = &cobra.Command{
	Use:     "deploytargets",
	Aliases: []string{"deploytarget", "dt"},
	Short:   "List all DeployTargets in Lagoon",
	Long:    "List all DeployTargets (kubernetes or openshift) in lagoon, this requires admin level permissions",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
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
		deploytargets, err := l.ListDeployTargets(context.TODO(), lc)
		if err != nil {
			return err
		}
		data := []output.Data{}
		for _, deploytarget := range *deploytargets {
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%v", deploytarget.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", deploytarget.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", deploytarget.ConsoleURL)),
				returnNonEmptyString(fmt.Sprintf("%v", deploytarget.BuildImage)),
				returnNonEmptyString(fmt.Sprintf("%v", deploytarget.Token)),
				returnNonEmptyString(fmt.Sprintf("%v", deploytarget.SSHHost)),
				returnNonEmptyString(fmt.Sprintf("%v", deploytarget.SSHPort)),
				returnNonEmptyString(fmt.Sprintf("%v", deploytarget.CloudRegion)),
				returnNonEmptyString(fmt.Sprintf("%v", deploytarget.CloudProvider)),
				returnNonEmptyString(fmt.Sprintf("%v", deploytarget.FriendlyName)),
				returnNonEmptyString(fmt.Sprintf("%v", deploytarget.RouterPattern)),
				returnNonEmptyString(fmt.Sprintf("%v", deploytarget.Created)),
				returnNonEmptyString(fmt.Sprintf("%v", deploytarget.MonitoringConfig)),
			})
		}
		output.RenderOutput(output.Table{
			Header: []string{
				"ID",
				"Name",
				"ConsoleUrl",
				"BuildImage",
				"Token",
				"SshHost",
				"SshPort",
				"CloudRegion",
				"CloudProvider",
				"FriendlyName",
				"RouterPattern",
				"Created",
				"MonitoringConfig",
			},
			Data: data,
		}, outputOptions)
		return nil
	},
}

var listGroupsCmd = &cobra.Command{
	Use:     "groups",
	Aliases: []string{"g"},
	Short:   "List groups you have access to (alias: g)",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
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

		groups, err := l.ListAllGroups(context.TODO(), lc)
		handleError(err)

		data := []output.Data{}
		for _, group := range *groups {
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%v", group.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", group.Name)),
			})
		}
		if len(data) == 0 {
			outputOptions.Error = "This account is not in any groups\n"
		}
		dataMain := output.Table{
			Header: []string{"ID", "Name"},
			Data:   data,
		}
		output.RenderOutput(dataMain, outputOptions)
		return nil
	},
}

var listGroupProjectsCmd = &cobra.Command{
	Use:     "group-projects",
	Aliases: []string{"gp"},
	Short:   "List projects in a group (alias: gp)",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
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
		listAllProjects, err := cmd.Flags().GetBool("all-projects")
		if err != nil {
			return err
		}

		if !listAllProjects {
			if err := requiredInputCheck("Group name", groupName); err != nil {
				return err
			}
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)

		var groupProjects *[]ls.Group

		if listAllProjects {
			groupProjects, err = l.GetGroupProjects(context.TODO(), "", lc)
		} else {
			groupProjects, err = l.GetGroupProjects(context.TODO(), groupName, lc)
		}
		handleError(err)
		var data []output.Data
		idx := 0
		for _, group := range *groupProjects {
			for _, project := range group.Projects {
				data = append(data, []string{
					returnNonEmptyString(fmt.Sprintf("%d", project.ID)),
					returnNonEmptyString(fmt.Sprintf("%s", project.Name)),
				})
				if listAllProjects {
					data[idx] = append(data[idx], returnNonEmptyString(fmt.Sprintf("%s", group.Name)))
				}
				idx++
			}
		}
		if len(data) == 0 {
			if !listAllProjects {
				outputOptions.Error = fmt.Sprintf("There are no projects in group '%s'\n", groupName)
			} else {
				outputOptions.Error = "There are no projects in any groups\n"
			}
		}

		dataMain := output.Table{
			Header: []string{"ID", "ProjectName"},
			Data:   data,
		}
		if listAllProjects {
			dataMain.Header = append(dataMain.Header, "GroupName")
		}
		output.RenderOutput(dataMain, outputOptions)
		return nil
	},
}

var listEnvironmentsCmd = &cobra.Command{
	Use:     "environments",
	Aliases: []string{"e"},
	Short:   "List environments for a project (alias: e)",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Project name", cmdProjectName); err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)
		environments, err := l.GetEnvironmentsByProjectName(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}

		if len(*environments) == 0 {
			outputOptions.Error = fmt.Sprintf("No environments found for project '%s'\n", cmdProjectName)
		}

		data := []output.Data{}
		for _, environment := range *environments {
			var envRoute = "none"
			if environment.Route != "" {
				envRoute = environment.Route
			}
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%d", environment.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", environment.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", environment.DeployType)),
				returnNonEmptyString(fmt.Sprintf("%v", environment.EnvironmentType)),
				returnNonEmptyString(fmt.Sprintf("%v", environment.OpenshiftProjectName)),
				returnNonEmptyString(fmt.Sprintf("%v", envRoute)),
				returnNonEmptyString(fmt.Sprintf("%v", environment.DeployTarget.Name)),
			})
		}
		dataMain := output.Table{
			Header: []string{"ID", "Name", "DeployType", "Environment", "Namespace", "Route", "DeployTarget"},
			Data:   data,
		}
		output.RenderOutput(dataMain, outputOptions)
		return nil
	},
}

var listVariablesCmd = &cobra.Command{
	Use:     "variables",
	Aliases: []string{"v"},
	Short:   "List variables for a project or environment (alias: v)",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requiredInputCheck("Project name", cmdProjectName); err != nil {
			return err
		}
		reveal, err := cmd.Flags().GetBool("reveal")
		if err != nil {
			return err
		}
		debug, err := cmd.Flags().GetBool("debug")
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
		in := &ls.EnvVariableByProjectEnvironmentNameInput{
			Project:     cmdProjectName,
			Environment: cmdProjectEnvironment,
		}
		envvars, err := l.GetEnvVariablesByProjectEnvironmentName(context.TODO(), in, lc)
		if err != nil {
			return err
		}
		data := []output.Data{}
		for _, envvar := range *envvars {
			env := []string{
				returnNonEmptyString(fmt.Sprintf("%v", envvar.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", cmdProjectName)),
			}
			if cmdProjectEnvironment != "" {
				env = append(env, returnNonEmptyString(fmt.Sprintf("%v", cmdProjectEnvironment)))
			}
			env = append(env, returnNonEmptyString(fmt.Sprintf("%v", envvar.Scope)))
			env = append(env, returnNonEmptyString(fmt.Sprintf("%v", envvar.Name)))
			if reveal {
				env = append(env, fmt.Sprintf("%v", envvar.Value))
			}
			data = append(data, env)
		}
		header := []string{
			"ID",
			"Project",
		}
		if cmdProjectEnvironment != "" {
			header = append(header, "Environment")
		}
		header = append(header, "Scope")
		header = append(header, "Name")
		if reveal {
			header = append(header, "Value")
		}
		if len(data) == 0 {
			if cmdProjectEnvironment != "" {
				outputOptions.Error = fmt.Sprintf("There are no variables for environment '%s' in project '%s'", cmdProjectEnvironment, cmdProjectName)
			} else {
				outputOptions.Error = fmt.Sprintf("There are no variables for project '%s'\n", cmdProjectName)
			}
		}
		output.RenderOutput(output.Table{
			Header: header,
			Data:   data,
		}, outputOptions)
		return nil
	},
}

var listDeploymentsCmd = &cobra.Command{
	Use:     "deployments",
	Aliases: []string{"d"},
	Short:   "List deployments for an environment (alias: d)",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Project name", cmdProjectName, "Environment name", cmdProjectEnvironment); err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)

		project, err := l.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
		handleError(err)

		deployments, err := l.GetDeploymentsByEnvironment(context.TODO(), project.ID, cmdProjectEnvironment, lc)
		handleError(err)

		data := []output.Data{}
		for _, deployment := range deployments.Deployments {
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%d", deployment.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", deployment.RemoteID)),
				returnNonEmptyString(fmt.Sprintf("%v", deployment.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", deployment.Status)),
				returnNonEmptyString(fmt.Sprintf("%v", deployment.Created)),
				returnNonEmptyString(fmt.Sprintf("%v", deployment.Started)),
				returnNonEmptyString(fmt.Sprintf("%v", deployment.Completed)),
			})
		}

		if len(data) == 0 {
			outputOptions.Error = fmt.Sprintf("There are no deployments for environment '%s' in project '%s'\n", cmdProjectEnvironment, cmdProjectName)
		}
		dataMain := output.Table{
			Header: []string{"ID", "RemoteID", "Name", "Status", "Created", "Started", "Completed"},
			Data:   data,
		}
		output.RenderOutput(dataMain, outputOptions)
		return nil
	},
}

var listTasksCmd = &cobra.Command{
	Use:     "tasks",
	Aliases: []string{"t"},
	Short:   "List tasks for an environment (alias: t)",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Project name", cmdProjectName, "Environment name", cmdProjectEnvironment); err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)

		project, err := l.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
		handleError(err)

		tasks, err := l.GetTasksByEnvironment(context.TODO(), project.ID, cmdProjectEnvironment, lc)
		handleError(err)

		data := []output.Data{}
		for _, task := range tasks.Tasks {
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%d", task.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", task.RemoteID)),
				returnNonEmptyString(fmt.Sprintf("%v", task.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", task.Status)),
				returnNonEmptyString(fmt.Sprintf("%v", task.Created)),
				returnNonEmptyString(fmt.Sprintf("%v", task.Started)),
				returnNonEmptyString(fmt.Sprintf("%v", task.Completed)),
				returnNonEmptyString(fmt.Sprintf("%v", task.Service)),
			})
		}

		if len(data) == 0 {
			outputOptions.Error = fmt.Sprintf("There are no tasks for environment '%s' in project '%s'\n", cmdProjectEnvironment, cmdProjectName)
		}
		dataMain := output.Table{
			Header: []string{"ID", "RemoteID", "Name", "Status", "Created", "Started", "Completed", "Service"},
			Data:   data,
		}
		output.RenderOutput(dataMain, outputOptions)
		return nil
	},
}

var listUsersCmd = &cobra.Command{
	Use:     "group-users",
	Aliases: []string{"gu"},
	Short:   "List all users in groups",
	Long: `List all users in groups in lagoon, this only shows users that are in groups.
If no group name is provided, all groups are queried.
Without a group name, this query may time out in large Lagoon installs.`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
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
		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)
		data := []output.Data{}
		if groupName != "" {
			// if a groupName is provided, use the groupbyname resolver
			groupMembers, err := l.ListGroupMembers(context.TODO(), groupName, lc)
			if err != nil {
				return err
			}
			for _, member := range groupMembers.Members {
				data = append(data, []string{
					returnNonEmptyString(fmt.Sprintf("%s", groupMembers.ID)),
					returnNonEmptyString(fmt.Sprintf("%s", groupMembers.Name)),
					returnNonEmptyString(fmt.Sprintf("%s", member.User.Email)),
					returnNonEmptyString(fmt.Sprintf("%s", member.Role)),
				})
			}
		} else {
			// otherwise allgroups query
			groupMembers, err := l.ListAllGroupMembers(context.TODO(), groupName, lc)
			if err != nil {
				return err
			}
			for _, group := range *groupMembers {
				for _, member := range group.Members {
					data = append(data, []string{
						returnNonEmptyString(fmt.Sprintf("%s", group.ID)),
						returnNonEmptyString(fmt.Sprintf("%s", group.Name)),
						returnNonEmptyString(fmt.Sprintf("%s", member.User.Email)),
						returnNonEmptyString(fmt.Sprintf("%s", member.Role)),
					})
				}
			}
		}
		dataMain := output.Table{
			Header: []string{"ID", "GroupName", "Email", "Role"},
			Data:   data,
		}
		output.RenderOutput(dataMain, outputOptions)
		return nil
	},
}

var listAllUsersCmd = &cobra.Command{
	Use:     "all-users",
	Aliases: []string{"au"},
	Short:   "List all users",
	Long: `List all users.
This query can take a long time to run if there are a lot of users.`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		emailAddress, err := cmd.Flags().GetString("email-address")
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
		allUsers, err := l.AllUsers(context.TODO(), ls.AllUsersFilter{
			Email: emailAddress,
		}, lc)
		if err != nil {
			return err
		}
		data := []output.Data{}
		for _, user := range *allUsers {
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%s", user.ID)),
				returnNonEmptyString(fmt.Sprintf("%s", user.Email)),
				returnNonEmptyString(fmt.Sprintf("%s", user.FirstName)),
				returnNonEmptyString(fmt.Sprintf("%s", user.LastName)),
				returnNonEmptyString(fmt.Sprintf("%s", user.Comment)),
			})
		}
		dataMain := output.Table{
			Header: []string{"ID", "Email", "FirstName", "LastName", "Comment"},
			Data:   data,
		}
		output.RenderOutput(dataMain, outputOptions)
		return nil
	},
}

var listUsersGroupsCmd = &cobra.Command{
	Use:     "user-groups",
	Aliases: []string{"ug"},
	Short:   "List a single users groups and roles",
	Long:    `List a single users groups and roles`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		emailAddress, err := cmd.Flags().GetString("email-address")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Email Address", emailAddress); err != nil {
			return err
		}
		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)
		allUsers, err := l.GetUserByEmail(context.TODO(), emailAddress, lc)
		if err != nil {
			return err
		}
		data := []output.Data{}
		for _, grouprole := range allUsers.GroupRoles {
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%s", allUsers.ID)),
				returnNonEmptyString(fmt.Sprintf("%s", allUsers.Email)),
				returnNonEmptyString(fmt.Sprintf("%s", grouprole.Name)),
				returnNonEmptyString(fmt.Sprintf("%s", grouprole.Role)),
			})
		}
		dataMain := output.Table{
			Header: []string{"ID", "Email", "GroupName", "GroupRole"},
			Data:   data,
		}
		output.RenderOutput(dataMain, outputOptions)
		return nil
	},
}

var listInvokableTasks = &cobra.Command{
	Use:     "invokable-tasks",
	Aliases: []string{"dcc"},
	Short:   "Print a list of invokable tasks",
	Long:    "Print a list of invokable user defined tasks registered against an environment",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Project name", cmdProjectName, "Environment name", cmdProjectEnvironment); err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)

		project, err := l.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
		handleError(err)
		tasks, err := l.GetInvokableAdvancedTaskDefinitionsByEnvironment(context.TODO(), project.ID, cmdProjectEnvironment, lc)
		handleError(err)

		data := []output.Data{}
		for _, task := range tasks.AdvancedTasks {
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%s", task.Name)),
				returnNonEmptyString(fmt.Sprintf("%s", task.Description)),
			})
		}

		if len(data) == 0 {
			outputOptions.Error = "There are no user defined tasks for this environment\n"
		}
		dataMain := output.Table{
			Header: []string{"Task Name", "Description"},
			Data:   data,
		}
		output.RenderOutput(dataMain, outputOptions)
		return nil
	},
}

var listNotificationCmd = &cobra.Command{
	Use:     "notification",
	Aliases: []string{"n"},
	Short:   "List all notifications or notifications on projects",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateToken(lagoonCLIConfig.Current) // get a new token if the current one is invalid
	},
}

var listProjectGroupsCmd = &cobra.Command{
	Use:     "project-groups",
	Aliases: []string{"pg"},
	Short:   "List groups in a project (alias: pg)",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Project name", cmdProjectName); err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)
		projectGroups, err := l.GetProjectGroups(context.TODO(), cmdProjectName, lc)
		handleError(err)

		if len(projectGroups.Groups) == 0 {
			outputOptions.Error = fmt.Sprintf("There are no groups for project '%s'\n", cmdProjectName)
		}

		data := []output.Data{}
		for _, group := range projectGroups.Groups {
			var organization = "null"
			if group.Organization != 0 {
				organization = strconv.Itoa(group.Organization)
			}
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%v", group.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", group.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", organization)),
			})
		}
		dataMain := output.Table{
			Header: []string{"Group ID", "Group Name", "Organization"},
			Data:   data,
		}
		output.RenderOutput(dataMain, outputOptions)
		return nil
	},
}

var listOrganizationCmd = &cobra.Command{
	Use:     "organization",
	Aliases: []string{"o"},
	Short:   "List all organizations projects, groups, deploy targets or users",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateToken(lagoonCLIConfig.Current)
	},
}

var listOrganizationProjectsCmd = &cobra.Command{
	Use:     "projects",
	Aliases: []string{"p"},
	Short:   "List projects in an organization",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
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

		org, err := l.GetOrganizationByName(context.TODO(), organizationName, lc)
		orgProjects, err := l.ListProjectsByOrganizationID(context.TODO(), org.ID, lc)
		handleError(err)

		if len(*orgProjects) == 0 {
			outputOptions.Error = fmt.Sprintf("No associated projects found for organization '%s'\n", organizationName)
		}

		data := []output.Data{}
		for _, project := range *orgProjects {
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%d", project.ID)),
				returnNonEmptyString(fmt.Sprintf("%s", project.Name)),
				returnNonEmptyString(fmt.Sprintf("%d", project.GroupCount)),
			})
		}
		dataMain := output.Table{
			Header: []string{"ID", "Name", "Group Count"},
			Data:   data,
		}
		output.RenderOutput(dataMain, outputOptions)
		return nil
	},
}

var listOrganizationGroupsCmd = &cobra.Command{
	Use:     "groups",
	Aliases: []string{"g"},
	Short:   "List groups in an organization",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
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

		org, err := l.GetOrganizationByName(context.TODO(), organizationName, lc)
		orgGroups, err := l.ListGroupsByOrganizationID(context.TODO(), org.ID, lc)
		handleError(err)

		if len(*orgGroups) == 0 {
			outputOptions.Error = fmt.Sprintf("No associated groups found for organization '%s'\n", organizationName)
		}

		data := []output.Data{}
		for _, group := range *orgGroups {
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%s", group.ID.String())),
				returnNonEmptyString(fmt.Sprintf("%s", group.Name)),
				returnNonEmptyString(fmt.Sprintf("%s", group.Type)),
				returnNonEmptyString(fmt.Sprintf("%d", group.MemberCount)),
			})
		}
		dataMain := output.Table{
			Header: []string{"ID", "Name", "Type", "Member Count"},
			Data:   data,
		}
		output.RenderOutput(dataMain, outputOptions)
		return nil
	},
}

var listOrganizationDeployTargetsCmd = &cobra.Command{
	Use:     "deploytargets",
	Aliases: []string{"d"},
	Short:   "List deploy targets in an organization",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
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
		organizationID, err := cmd.Flags().GetUint("id")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Organization name", organizationName, "Organization ID", strconv.Itoa(int(organizationID))); err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)
		deployTargets, err := l.ListDeployTargetsByOrganizationNameOrID(context.TODO(), nullStrCheck(organizationName), nullUintCheck(organizationID), lc)
		handleError(err)

		if len(*deployTargets) == 0 {
			outputOptions.Error = fmt.Sprintf("No associated deploy targets found for organization '%s'\n", organizationName)
		}

		data := []output.Data{}
		for _, dt := range *deployTargets {
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%d", dt.ID)),
				returnNonEmptyString(fmt.Sprintf("%s", dt.Name)),
				returnNonEmptyString(fmt.Sprintf("%s", dt.RouterPattern)),
				returnNonEmptyString(fmt.Sprintf("%s", dt.CloudRegion)),
				returnNonEmptyString(fmt.Sprintf("%s", dt.CloudProvider)),
				returnNonEmptyString(fmt.Sprintf("%s", dt.SSHHost)),
				returnNonEmptyString(fmt.Sprintf("%s", dt.SSHPort)),
			})
		}
		dataMain := output.Table{
			Header: []string{"ID", "Name", "Router Pattern", "Cloud Region", "Cloud Provider", "SSH Host", "SSH Port"},
			Data:   data,
		}
		output.RenderOutput(dataMain, outputOptions)
		return nil
	},
}

var ListOrganizationUsersCmd = &cobra.Command{
	Use:     "users",
	Aliases: []string{"u"},
	Short:   "List users in an organization",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
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
		organization, err := l.GetOrganizationByName(context.Background(), organizationName, lc)
		handleError(err)
		users, err := l.UsersByOrganization(context.TODO(), organization.ID, lc)
		handleError(err)

		data := []output.Data{}
		for _, user := range *users {
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%s", user.ID)),
				returnNonEmptyString(fmt.Sprintf("%s", user.Email)),
				returnNonEmptyString(fmt.Sprintf("%s", user.FirstName)),
				returnNonEmptyString(fmt.Sprintf("%s", user.LastName)),
				returnNonEmptyString(fmt.Sprintf("%s", user.Comment)),
				returnNonEmptyString(fmt.Sprintf("%v", user.Owner)),
			})
		}
		dataMain := output.Table{
			Header: []string{"ID", "Email", "First Name", "LastName", "Comment", "Owner"},
			Data:   data,
		}
		output.RenderOutput(dataMain, outputOptions)
		return nil
	},
}

var listOrganizationsCmd = &cobra.Command{
	Use:     "organizations",
	Aliases: []string{"o"},
	Short:   "List all organizations",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
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

		organizations, err := l.AllOrganizations(context.TODO(), lc)

		data := []output.Data{}
		for _, organization := range *organizations {
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%d", organization.ID)),
				returnNonEmptyString(fmt.Sprintf("%s", organization.Name)),
				returnNonEmptyString(fmt.Sprintf("%s", organization.Description)),
				returnNonEmptyString(fmt.Sprintf("%d", organization.QuotaProject)),
				returnNonEmptyString(fmt.Sprintf("%d", organization.QuotaGroup)),
				returnNonEmptyString(fmt.Sprintf("%d", organization.QuotaNotification)),
				returnNonEmptyString(fmt.Sprintf("%d", organization.QuotaEnvironment)),
				returnNonEmptyString(fmt.Sprintf("%d", organization.QuotaRoute)),
			})
		}
		dataMain := output.Table{
			Header: []string{"ID", "Name", "Description", "Project Quota", "Group Quota", "Notification Quota", "Environment Quota", "Route Quota"},
			Data:   data,
		}
		output.RenderOutput(dataMain, outputOptions)
		return nil
	},
}

func init() {
	listCmd.AddCommand(listDeployTargetsCmd)
	listCmd.AddCommand(listDeploymentsCmd)
	listCmd.AddCommand(listGroupsCmd)
	listCmd.AddCommand(listGroupProjectsCmd)
	listCmd.AddCommand(listProjectGroupsCmd)
	listCmd.AddCommand(listEnvironmentsCmd)
	listCmd.AddCommand(listProjectsCmd)
	listCmd.AddCommand(listNotificationCmd)
	listCmd.AddCommand(listTasksCmd)
	listCmd.AddCommand(listUsersCmd)
	listCmd.AddCommand(listVariablesCmd)
	listCmd.AddCommand(listInvokableTasks)
	listCmd.AddCommand(listBackupsCmd)
	listCmd.AddCommand(listDeployTargetConfigsCmd)
	listCmd.AddCommand(listAllUsersCmd)
	listCmd.AddCommand(listUsersGroupsCmd)
	listAllUsersCmd.Flags().StringP("email-address", "E", "", "The email address of a user")
	listUsersGroupsCmd.Flags().StringP("email-address", "E", "", "The email address of a user")
	listCmd.AddCommand(listOrganizationCmd)
	listOrganizationCmd.AddCommand(listOrganizationProjectsCmd)
	listOrganizationCmd.AddCommand(ListOrganizationUsersCmd)
	listOrganizationCmd.AddCommand(listOrganizationGroupsCmd)
	listOrganizationCmd.AddCommand(listOrganizationDeployTargetsCmd)
	listOrganizationCmd.AddCommand(listOrganizationsCmd)
	listCmd.Flags().BoolVarP(&listAllProjects, "all-projects", "", false, "All projects (if supported)")
	listUsersCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group to list users in")
	listGroupProjectsCmd.Flags().StringP("name", "N", "", "Name of the group to list projects in")
	listGroupProjectsCmd.Flags().BoolP("all-projects", "", false, "All projects")
	listVariablesCmd.Flags().BoolP("reveal", "", false, "Reveal the variable values")
	listOrganizationProjectsCmd.Flags().StringP("name", "O", "", "Name of the organization to list associated projects for")
	ListOrganizationUsersCmd.Flags().StringP("name", "O", "", "Name of the organization to list associated users for")
	listOrganizationGroupsCmd.Flags().StringP("name", "O", "", "Name of the organization to list associated groups for")
	listOrganizationDeployTargetsCmd.Flags().StringP("name", "O", "", "Name of the organization to list associated deploy targets for")
	listOrganizationDeployTargetsCmd.Flags().Uint("id", 0, "ID of the organization to list associated deploy targets for")
}
