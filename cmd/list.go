package cmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/pkg/output"
	"github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	"github.com/uselagoon/machinery/api/schema"
)

// ListFlags .
type ListFlags struct {
	Project     string `json:"project,omitempty"`
	Environment string `json:"environment,omitempty"`
	Reveal      bool   `json:"reveal,omitempty"`
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
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)

		projects, err := lagoon.ListAllProjects(context.TODO(), lc)
		if err != nil {
			return err
		}

		data := []output.Data{}
		for _, project := range *projects {
			var devEnvironments = 0
			productionRoute := ""
			for _, environment := range project.Environments {
				if environment.EnvironmentType == "development" {
					devEnvironments++
				}
				if environment.EnvironmentType == "production" {
					productionRoute = environment.Route
				}
			}

			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%d", project.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", project.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", project.GitURL)),
				returnNonEmptyString(fmt.Sprintf("%v", project.ProductionEnvironment)),
				returnNonEmptyString(fmt.Sprintf("%v", productionRoute)),
				returnNonEmptyString(fmt.Sprintf("%v/%v", devEnvironments, project.DevelopmentEnvironmentsLimit)),
			})
		}
		if len(data) == 0 {
			outputOptions.Error = "No access to any projects in Lagoon\n"
		}
		dataMain := output.Table{
			Header: []string{"ID", "ProjectName", "GitUrl", "ProductionEnvironment", "ProductionRoute", "DevEnvironments"},
			Data:   data,
		}
		r := output.RenderOutput(dataMain, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		return nil
	},
}

var listDeployTargetsCmd = &cobra.Command{
	Use:     "deploytargets",
	Aliases: []string{"deploytarget", "dt"},
	Short:   "List all DeployTargets in Lagoon",
	Long:    "List all Deploytargets(Kubernetes) in lagoon, this requires admin level permissions",
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
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)
		deploytargets, err := lagoon.ListDeployTargets(context.TODO(), lc)
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
		outputOptions.MultiLine = true
		r := output.RenderOutput(output.Table{
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
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)

		groups, err := lagoon.ListAllGroups(context.TODO(), lc)
		if err != nil {
			return err
		}

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
		r := output.RenderOutput(dataMain, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)

		var groupProjects *[]schema.Group

		if listAllProjects {
			groupProjects, err = lagoon.GetGroupProjects(context.TODO(), "", lc)
			if err != nil {
				return err
			}
		} else {
			groupProjects, err = lagoon.GetGroupProjects(context.TODO(), groupName, lc)
			if err != nil {
				return err
			}
		}
		var data []output.Data
		idx := 0
		for _, group := range *groupProjects {
			for _, project := range group.Projects {
				data = append(data, []string{
					returnNonEmptyString(fmt.Sprintf("%d", project.ID)),
					returnNonEmptyString(project.Name),
				})
				if listAllProjects {
					data[idx] = append(data[idx], returnNonEmptyString(group.Name))
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
			r := output.RenderOutput(output.Table{Data: []output.Data{[]string{}}}, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
			return nil
		}

		dataMain := output.Table{
			Header: []string{"ID", "ProjectName"},
			Data:   data,
		}
		if listAllProjects {
			dataMain.Header = append(dataMain.Header, "GroupName")
		}
		r := output.RenderOutput(dataMain, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)
		environments, err := lagoon.GetEnvironmentsByProjectName(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}

		if len(*environments) == 0 {
			outputOptions.Error = fmt.Sprintf("No environments found for project '%s'\n", cmdProjectName)
			r := output.RenderOutput(output.Table{Data: []output.Data{[]string{}}}, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
			return nil
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
			Header: []string{"ID", "Name", "DeployType", "EnvironmentType", "Namespace", "Route", "DeployTarget"},
			Data:   data,
		}
		r := output.RenderOutput(dataMain, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)
		in := &schema.EnvVariableByProjectEnvironmentNameInput{
			Project:     cmdProjectName,
			Environment: cmdProjectEnvironment,
		}
		envvars, err := lagoon.GetEnvVariablesByProjectEnvironmentName(context.TODO(), in, lc)
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
				outputOptions.MultiLine = true
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
				outputOptions.Error = fmt.Sprintf("There are no variables for environment '%s' in project '%s'\n", cmdProjectEnvironment, cmdProjectName)
			} else {
				outputOptions.Error = fmt.Sprintf("There are no variables for project '%s'\n", cmdProjectName)
			}
			r := output.RenderOutput(output.Table{Data: []output.Data{[]string{}}}, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
			return nil
		}
		r := output.RenderOutput(output.Table{
			Header: header,
			Data:   data,
		}, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)

		project, err := lagoon.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}

		deployments, err := lagoon.GetDeploymentsByEnvironment(context.TODO(), project.ID, cmdProjectEnvironment, lc)
		if err != nil {
			return err
		}

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
			r := output.RenderOutput(output.Table{Data: []output.Data{[]string{}}}, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
			return nil
		}
		dataMain := output.Table{
			Header: []string{"ID", "RemoteID", "Name", "Status", "Created", "Started", "Completed"},
			Data:   data,
		}
		r := output.RenderOutput(dataMain, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)

		project, err := lagoon.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}

		tasks, err := lagoon.GetTasksByEnvironment(context.TODO(), project.ID, cmdProjectEnvironment, lc)
		if err != nil {
			return err
		}

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
			r := output.RenderOutput(output.Table{Data: []output.Data{[]string{}}}, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
			return nil
		}
		dataMain := output.Table{
			Header: []string{"ID", "RemoteID", "Name", "Status", "Created", "Started", "Completed", "Service"},
			Data:   data,
		}
		r := output.RenderOutput(dataMain, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		return nil
	},
}

var listUsersCmd = &cobra.Command{
	Use:     "group-users",
	Aliases: []string{"gu"},
	Short:   "List all users in groups",
	Long: `List all users in groups in lagoon, this only shows users that are in groups.
If no group name is provided, all groups are queried.
Without a group name, this query may time out in large Lagoon instalschema.`,
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
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)
		data := []output.Data{}
		if groupName != "" {
			// if a groupName is provided, use the groupbyname resolver
			groupMembers, err := lagoon.ListGroupMembers(context.TODO(), groupName, lc)
			if err != nil {
				return err
			}
			for _, member := range groupMembers.Members {
				data = append(data, []string{
					returnNonEmptyString(groupMembers.ID.String()),
					returnNonEmptyString(groupMembers.Name),
					returnNonEmptyString(member.User.Email),
					returnNonEmptyString(string(member.Role)),
				})
			}
		} else {
			// otherwise allgroups query
			groupMembers, err := lagoon.ListAllGroupMembers(context.TODO(), groupName, lc)
			if err != nil {
				return err
			}
			for _, group := range *groupMembers {
				for _, member := range group.Members {
					data = append(data, []string{
						returnNonEmptyString(group.ID.String()),
						returnNonEmptyString(group.Name),
						returnNonEmptyString(member.User.Email),
						returnNonEmptyString(string(member.Role)),
					})
				}
			}
		}
		dataMain := output.Table{
			Header: []string{"ID", "GroupName", "Email", "Role"},
			Data:   data,
		}
		r := output.RenderOutput(dataMain, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
		emailAddress, err := cmd.Flags().GetString("email")
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
		allUsers, err := lagoon.AllUsers(context.TODO(), schema.AllUsersFilter{
			Email: emailAddress,
		}, lc)
		if err != nil {
			return err
		}
		data := []output.Data{}
		for _, user := range *allUsers {
			data = append(data, []string{
				returnNonEmptyString(user.ID.String()),
				returnNonEmptyString(user.Email),
				returnNonEmptyString(user.FirstName),
				returnNonEmptyString(user.LastName),
				returnNonEmptyString(user.Comment),
			})
		}
		dataMain := output.Table{
			Header: []string{"ID", "Email", "FirstName", "LastName", "Comment"},
			Data:   data,
		}
		r := output.RenderOutput(dataMain, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
		emailAddress, err := cmd.Flags().GetString("email")
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
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)
		allUsers, err := lagoon.GetUserByEmail(context.TODO(), emailAddress, lc)
		if err != nil {
			return err
		}
		data := []output.Data{}
		for _, grouprole := range allUsers.GroupRoles {
			data = append(data, []string{
				returnNonEmptyString(allUsers.ID.String()),
				returnNonEmptyString(allUsers.Email),
				returnNonEmptyString(grouprole.Name),
				returnNonEmptyString(grouprole.Role),
			})
		}
		dataMain := output.Table{
			Header: []string{"ID", "Email", "GroupName", "GroupRole"},
			Data:   data,
		}
		r := output.RenderOutput(dataMain, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)

		project, err := lagoon.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		tasks, err := lagoon.GetInvokableAdvancedTaskDefinitionsByEnvironment(context.TODO(), project.ID, cmdProjectEnvironment, lc)
		if err != nil {
			return err
		}

		data := []output.Data{}
		for _, task := range tasks.AdvancedTasks {
			data = append(data, []string{
				returnNonEmptyString(task.Name),
				returnNonEmptyString(task.Description),
			})
		}

		if len(data) == 0 {
			outputOptions.Error = "There are no user defined tasks for this environment\n"
			r := output.RenderOutput(output.Table{Data: []output.Data{[]string{}}}, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
			return nil
		}
		dataMain := output.Table{
			Header: []string{"Task Name", "Description"},
			Data:   data,
		}
		r := output.RenderOutput(dataMain, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)
		projectGroups, err := lagoon.GetProjectGroups(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}

		if len(projectGroups.Groups) == 0 {
			outputOptions.Error = fmt.Sprintf("There are no groups for project '%s'\n", cmdProjectName)
			r := output.RenderOutput(output.Table{Data: []output.Data{[]string{}}}, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
			return nil
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
		r := output.RenderOutput(dataMain, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		return nil
	},
}

var listOrganizationProjectsCmd = &cobra.Command{
	Use:     "organization-projects",
	Aliases: []string{"org-p"},
	Short:   "List projects in an organization",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		organizationName, err := cmd.Flags().GetString("organization-name")
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
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)

		organization, err := lagoon.GetOrganizationByName(context.TODO(), organizationName, lc)
		if err != nil {
			return err
		}
		if organization.Name == "" {
			return fmt.Errorf("error querying organization by name")
		}
		orgProjects, err := lagoon.ListProjectsByOrganizationID(context.TODO(), organization.ID, lc)
		if err != nil {
			return err
		}

		if len(*orgProjects) == 0 {
			outputOptions.Error = fmt.Sprintf("No associated projects found for organization '%s'\n", organizationName)
			r := output.RenderOutput(output.Table{Data: []output.Data{[]string{}}}, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
			return nil
		}

		data := []output.Data{}
		for _, project := range *orgProjects {
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%d", project.ID)),
				returnNonEmptyString(project.Name),
				returnNonEmptyString(fmt.Sprintf("%d", project.GroupCount)),
			})
		}
		dataMain := output.Table{
			Header: []string{"ID", "Name", "Group Count"},
			Data:   data,
		}
		r := output.RenderOutput(dataMain, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		return nil
	},
}

var listOrganizationGroupsCmd = &cobra.Command{
	Use:     "organization-groups",
	Aliases: []string{"org-g"},
	Short:   "List groups in an organization",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		organizationName, err := cmd.Flags().GetString("organization-name")
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
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)

		organization, err := lagoon.GetOrganizationByName(context.TODO(), organizationName, lc)
		if err != nil {
			return err
		}
		if organization.Name == "" {
			return fmt.Errorf("error querying organization by name")
		}
		orgGroups, err := lagoon.ListGroupsByOrganizationID(context.TODO(), organization.ID, lc)
		if err != nil {
			return err
		}
		if len(*orgGroups) == 0 {
			outputOptions.Error = fmt.Sprintf("No associated groups found for organization '%s'\n", organizationName)
			r := output.RenderOutput(output.Table{Data: []output.Data{[]string{}}}, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
			return nil
		}

		data := []output.Data{}
		for _, group := range *orgGroups {
			data = append(data, []string{
				returnNonEmptyString(group.ID.String()),
				returnNonEmptyString(group.Name),
				returnNonEmptyString(group.Type),
				returnNonEmptyString(fmt.Sprintf("%d", group.MemberCount)),
			})
		}
		dataMain := output.Table{
			Header: []string{"ID", "Name", "Type", "Member Count"},
			Data:   data,
		}
		r := output.RenderOutput(dataMain, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		return nil
	},
}

var listOrganizationDeployTargetsCmd = &cobra.Command{
	Use:     "organization-deploytargets",
	Aliases: []string{"org-dt"},
	Short:   "List deploy targets in an organization",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		organizationName, err := cmd.Flags().GetString("organization-name")
		if err != nil {
			return err
		}
		organizationID, err := cmd.Flags().GetUint("id")
		if err != nil {
			return err
		}
		if organizationName == "" && organizationID == 0 {
			return fmt.Errorf("missing arguments: Organization name or ID is not defined")
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)
		deployTargets, err := lagoon.ListDeployTargetsByOrganizationNameOrID(context.TODO(), nullStrCheck(organizationName), nullUintCheck(organizationID), lc)
		if err != nil {
			return err
		}
		if len(*deployTargets) == 0 {
			outputOptions.Error = fmt.Sprintf("No associated deploy targets found for organization '%s'\n", organizationName)
			r := output.RenderOutput(output.Table{Data: []output.Data{[]string{}}}, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
			return nil
		}

		data := []output.Data{}
		for _, dt := range *deployTargets {
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%d", dt.ID)),
				returnNonEmptyString(dt.Name),
				returnNonEmptyString(dt.RouterPattern),
				returnNonEmptyString(dt.CloudRegion),
				returnNonEmptyString(dt.CloudProvider),
				returnNonEmptyString(dt.SSHHost),
				returnNonEmptyString(dt.SSHPort),
			})
		}
		dataMain := output.Table{
			Header: []string{"ID", "Name", "Router Pattern", "Cloud Region", "Cloud Provider", "SSH Host", "SSH Port"},
			Data:   data,
		}
		r := output.RenderOutput(dataMain, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		return nil
	},
}

var ListOrganizationUsersCmd = &cobra.Command{
	Use:     "organization-users",
	Aliases: []string{"org-u"},
	Short:   "List users in an organization",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		organizationName, err := cmd.Flags().GetString("organization-name")
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
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)
		users, err := lagoon.UsersByOrganizationName(context.TODO(), organizationName, lc)
		if err != nil {
			return err
		}

		data := []output.Data{}
		for _, user := range *users {
			data = append(data, []string{
				returnNonEmptyString(user.ID.String()),
				returnNonEmptyString(user.Email),
				returnNonEmptyString(user.FirstName),
				returnNonEmptyString(user.LastName),
				returnNonEmptyString(user.Comment),
			})
		}
		dataMain := output.Table{
			Header: []string{"ID", "Email", "First Name", "LastName", "Comment"},
			Data:   data,
		}
		r := output.RenderOutput(dataMain, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		return nil
	},
}

var ListOrganizationAdminsCmd = &cobra.Command{
	Use:     "organization-admins",
	Aliases: []string{"org-a"},
	Short:   "List admins in an organization",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		organizationName, err := cmd.Flags().GetString("organization-name")
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
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)
		users, err := lagoon.ListOrganizationAdminsByName(context.TODO(), organizationName, lc)
		if err != nil {
			return err
		}

		data := []output.Data{}
		for _, user := range *users {
			role := "viewer"
			if user.Owner {
				role = "owner"
			}
			data = append(data, []string{
				returnNonEmptyString(user.ID.String()),
				returnNonEmptyString(user.Email),
				returnNonEmptyString(user.FirstName),
				returnNonEmptyString(user.LastName),
				returnNonEmptyString(role),
			})
		}
		dataMain := output.Table{
			Header: []string{"ID", "Email", "First Name", "LastName", "OrganizationRole"},
			Data:   data,
		}
		r := output.RenderOutput(dataMain, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)

		organizations, err := lagoon.AllOrganizations(context.TODO(), lc)
		if err != nil {
			return err
		}

		data := []output.Data{}
		for _, organization := range *organizations {
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%d", organization.ID)),
				returnNonEmptyString(organization.Name),
				returnNonEmptyString(organization.Description),
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
		r := output.RenderOutput(dataMain, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
	listCmd.AddCommand(listOrganizationProjectsCmd)
	listCmd.AddCommand(ListOrganizationUsersCmd)
	listCmd.AddCommand(ListOrganizationAdminsCmd)
	listCmd.AddCommand(listOrganizationGroupsCmd)
	listCmd.AddCommand(listOrganizationDeployTargetsCmd)
	listCmd.AddCommand(listOrganizationsCmd)
	listAllUsersCmd.Flags().StringP("email", "E", "", "The email address of a user")
	listUsersGroupsCmd.Flags().StringP("email", "E", "", "The email address of a user")
	listCmd.Flags().BoolVarP(&listAllProjects, "all-projects", "", false, "All projects (if supported)")
	listUsersCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group to list users in")
	listGroupProjectsCmd.Flags().StringP("name", "N", "", "Name of the group to list projects in")
	listGroupProjectsCmd.Flags().BoolP("all-projects", "", false, "All projects")
	listVariablesCmd.Flags().BoolP("reveal", "", false, "Reveal the variable values")
	listOrganizationProjectsCmd.Flags().StringP("organization-name", "O", "", "Name of the organization to list associated projects for")
	ListOrganizationUsersCmd.Flags().StringP("organization-name", "O", "", "Name of the organization to list associated users for")
	ListOrganizationAdminsCmd.Flags().StringP("organization-name", "O", "", "Name of the organization to list associated users for")
	listOrganizationGroupsCmd.Flags().StringP("organization-name", "O", "", "Name of the organization to list associated groups for")
	listOrganizationDeployTargetsCmd.Flags().StringP("organization-name", "O", "", "Name of the organization to list associated deploy targets for")
	listOrganizationDeployTargetsCmd.Flags().Uint("id", 0, "ID of the organization to list associated deploy targets for")
}
