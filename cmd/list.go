package cmd

import (
	"context"
	"fmt"
	"strconv"
	"strings"

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
		wide, err := cmd.Flags().GetBool("wide")
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
			projData := []string{
				returnNonEmptyString(fmt.Sprintf("%d", project.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", project.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", project.GitURL)),
				returnNonEmptyString(fmt.Sprintf("%v", project.ProductionEnvironment)),
				returnNonEmptyString(fmt.Sprintf("%v", productionRoute)),
			}
			if project.DevelopmentEnvironmentsLimit != nil {
				projData = append(projData, returnNonEmptyString(fmt.Sprintf("%v/%v", devEnvironments, *project.DevelopmentEnvironmentsLimit)))
			} else {
				projData = append(projData, returnNonEmptyString(fmt.Sprintf("%v/%v", devEnvironments, 0)))
			}

			if wide {
				autoIdle, err := strconv.ParseBool(strconv.Itoa(int(*project.AutoIdle)))
				if err != nil {
					return err
				}
				factsUI, err := strconv.ParseBool(strconv.Itoa(int(*project.FactsUI)))
				if err != nil {
					return err
				}
				problemsUI, err := strconv.ParseBool(strconv.Itoa(int(*project.ProblemsUI)))
				if err != nil {
					return err
				}
				deploymentsDisabled, err := strconv.ParseBool(strconv.Itoa(int(*project.DeploymentsDisabled)))
				if err != nil {
					return err
				}

				projData = append(projData, returnNonEmptyString(fmt.Sprintf("%v", autoIdle)))
				projData = append(projData, returnNonEmptyString(fmt.Sprintf("%v", project.Branches)))
				projData = append(projData, returnNonEmptyString(fmt.Sprintf("%v", project.PullRequests)))
				projData = append(projData, returnNonEmptyString(fmt.Sprintf("%v", project.RouterPattern)))
				projData = append(projData, returnNonEmptyString(fmt.Sprintf("%v", factsUI)))
				projData = append(projData, returnNonEmptyString(fmt.Sprintf("%v", problemsUI)))
				projData = append(projData, returnNonEmptyString(fmt.Sprintf("%v", deploymentsDisabled)))
			}
			data = append(data, projData)
		}
		if len(data) == 0 {
			return handleNilResults("No access to any projects in Lagoon\n", cmd)
		}
		projHeader := []string{"ID", "ProjectName", "GitUrl", "ProductionEnvironment", "ProductionRoute", "DevEnvironments"}
		if wide {
			projHeader = append(projHeader, "AutoIdle")
			projHeader = append(projHeader, "Branches")
			projHeader = append(projHeader, "PullRequests")
			projHeader = append(projHeader, "RouterPattern")
			projHeader = append(projHeader, "FactsUI")
			projHeader = append(projHeader, "ProblemsUI")
			projHeader = append(projHeader, "DeploymentsDisabled")
		}
		dataMain := output.Table{
			Header: projHeader,
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
		wide, err := cmd.Flags().GetBool("wide")
		if err != nil {
			return err
		}
		showToken, err := cmd.Flags().GetBool("show-token")
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
			depTarget := []string{
				returnNonEmptyString(fmt.Sprintf("%v", deploytarget.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", deploytarget.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", deploytarget.RouterPattern)),
				returnNonEmptyString(fmt.Sprintf("%v", deploytarget.SSHHost)),
				returnNonEmptyString(fmt.Sprintf("%v", deploytarget.SSHPort)),
				returnNonEmptyString(fmt.Sprintf("%v", deploytarget.BuildImage)),
				returnNonEmptyString(fmt.Sprintf("%v", deploytarget.ConsoleURL)),
			}
			if wide {
				depTarget = append(depTarget, returnNonEmptyString(fmt.Sprintf("%v", deploytarget.CloudRegion)))
				depTarget = append(depTarget, returnNonEmptyString(fmt.Sprintf("%v", deploytarget.CloudProvider)))
				depTarget = append(depTarget, returnNonEmptyString(fmt.Sprintf("%v", deploytarget.FriendlyName)))
				depTarget = append(depTarget, returnNonEmptyString(fmt.Sprintf("%v", deploytarget.MonitoringConfig)))
				depTarget = append(depTarget, returnNonEmptyString(fmt.Sprintf("%v", deploytarget.Created)))
			}
			if showToken {
				depTarget = append(depTarget, returnNonEmptyString(fmt.Sprintf("%v", deploytarget.Token)))
			}
			data = append(data, depTarget)
		}
		outputOptions.MultiLine = true
		header := []string{
			"ID",
			"Name",
			"RouterPattern",
			"SshHost",
			"SshPort",
			"BuildImage",
			"ConsoleUrl",
		}
		if wide {
			header = append(header, "CloudRegion")
			header = append(header, "CloudProvider")
			header = append(header, "FriendlyName")
			header = append(header, "MonitoringConfig")
			header = append(header, "Created")
		}
		if showToken {
			header = append(header, "Token")
		}
		r := output.RenderOutput(output.Table{
			Header: header,
			Data:   data,
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
			return handleNilResults("This account is not in any groups\n", cmd)
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
				return handleNilResults("There are no projects in group '%s'\n", cmd, groupName)
			} else {
				return handleNilResults("There are no projects in any groups\n", cmd)
			}
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
		wide, err := cmd.Flags().GetBool("wide")
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
			return handleNilResults("No environments found for project '%s'\n", cmd, cmdProjectName)
		}

		data := []output.Data{}
		for _, environment := range *environments {
			var envRoute = "none"
			if environment.Route != "" {
				envRoute = environment.Route
			}
			envData := []string{
				returnNonEmptyString(fmt.Sprintf("%d", environment.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", environment.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", environment.DeployType)),
				returnNonEmptyString(fmt.Sprintf("%v", environment.EnvironmentType)),
				returnNonEmptyString(fmt.Sprintf("%v", environment.OpenshiftProjectName)),
				returnNonEmptyString(fmt.Sprintf("%v", envRoute)),
				returnNonEmptyString(fmt.Sprintf("%v", environment.DeployTarget.Name)),
			}

			if wide {
				autoIdle, err := strconv.ParseBool(strconv.Itoa(int(*environment.AutoIdle)))
				if err != nil {
					return err
				}

				envData = append(envData, returnNonEmptyString(fmt.Sprintf("%v", environment.Created)))
				envData = append(envData, returnNonEmptyString(fmt.Sprintf("%v", autoIdle)))
				envData = append(envData, returnNonEmptyString(fmt.Sprintf("%v", environment.DeployTitle)))
				envData = append(envData, returnNonEmptyString(fmt.Sprintf("%v", environment.DeployBaseRef)))
				envData = append(envData, returnNonEmptyString(fmt.Sprintf("%v", environment.DeployHeadRef)))
				envData = append(envData, returnNonEmptyString(fmt.Sprintf("%v", environment.Routes)))
			}
			data = append(data, envData)
		}
		environmentHeaders := []string{"ID", "Name", "DeployType", "EnvironmentType", "Namespace", "Route", "DeployTarget"}
		if wide {
			environmentHeaders = append(environmentHeaders, "Created")
			environmentHeaders = append(environmentHeaders, "AutoIdle")
			environmentHeaders = append(environmentHeaders, "DeployTitle")
			environmentHeaders = append(environmentHeaders, "DeployBaseRef")
			environmentHeaders = append(environmentHeaders, "DeployHeadRef")
			environmentHeaders = append(environmentHeaders, "Routes")
		}
		dataMain := output.Table{
			Header: environmentHeaders,
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
		envvars, err := lagoon.GetEnvVariablesByProjectEnvironmentName(context.TODO(), in, reveal, lc)
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
				return handleNilResults("There are no variables for environment '%s' in project '%s'\n", cmd, cmdProjectEnvironment, cmdProjectName)
			} else {
				return handleNilResults("There are no variables for project '%s'\n", cmd, cmdProjectName)
			}
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
		wide, err := cmd.Flags().GetBool("wide")
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

		deployments, err := lagoon.GetDeploymentsByEnvironmentAndProjectName(context.TODO(), cmdProjectName, cmdProjectEnvironment, lc)
		if err != nil {
			return fmt.Errorf("%v: check if the project or environment exists", err.Error())
		}

		data := []output.Data{}
		for _, deployment := range deployments.Deployments {
			dep := []string{
				returnNonEmptyString(fmt.Sprintf("%d", deployment.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", deployment.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", deployment.Status)),
				returnNonEmptyString(fmt.Sprintf("%v", deployment.BuildStep)),
				returnNonEmptyString(fmt.Sprintf("%v", deployment.Started)),
				returnNonEmptyString(fmt.Sprintf("%v", deployment.Completed)),
			}
			if wide {
				dep = append(dep, returnNonEmptyString(fmt.Sprintf("%v", deployment.Created)))
				dep = append(dep, returnNonEmptyString(fmt.Sprintf("%v", deployment.RemoteID)))
			}
			data = append(data, dep)
		}

		if len(data) == 0 {
			return handleNilResults("There are no deployments for environment '%s' in project '%s'\n", cmd, cmdProjectEnvironment, cmdProjectName)
		}
		header := []string{
			"ID",
			"Name",
			"Status",
			"BuildStep",
			"Started",
			"Completed",
		}
		if wide {
			header = append(header, "Created")
			header = append(header, "RemoteID")
		}
		dataMain := output.Table{
			Header: header,
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

		tasks, err := lagoon.GetTasksByEnvironmentAndProjectName(context.TODO(), cmdProjectName, cmdProjectEnvironment, lc)
		if err != nil {
			return fmt.Errorf("%v: check if the project or environment exists", err.Error())
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
			return handleNilResults("There are no tasks for environment '%s' in project '%s'\n", cmd, cmdProjectEnvironment, cmdProjectName)
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

		tasks, err := lagoon.GetInvokableAdvancedTaskDefinitionsByEnvironmentAndProjectName(context.TODO(), cmdProjectName, cmdProjectEnvironment, lc)
		if err != nil {
			return fmt.Errorf("%v: check if the project or environment exists", err.Error())
		}

		data := []output.Data{}
		for _, task := range tasks.AdvancedTasks {
			data = append(data, []string{
				returnNonEmptyString(task.Name),
				returnNonEmptyString(task.Description),
			})
		}

		if len(data) == 0 {
			return handleNilResults("There are no user defined tasks for environment %s\n", cmd, cmdProjectEnvironment)
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
			return handleNilResults("There are no groups for project '%s'\n", cmd, cmdProjectName)
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
			return handleNilResults("No associated projects found for organization '%s'\n", cmd, organizationName)
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
			return handleNilResults("No associated groups found for organization '%s'\n", cmd, organizationName)
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
			return handleNilResults("No associated deploy targets found for organization '%s'\n", cmd, organizationName)
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

var listOrganizationUsersCmd = &cobra.Command{
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

var listOrganizationAdminsCmd = &cobra.Command{
	Use:     "organization-admininstrators",
	Aliases: []string{"organization-admins", "org-admins", "org-a"},
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
		if len(*users) == 0 {
			return handleNilResults("No associated users found for organization '%s'\n", cmd, organizationName)
		}
		data := []output.Data{}
		for _, user := range *users {
			role := "viewer"
			if user.Owner {
				role = "owner"
			}
			if user.Admin {
				role = "admin"
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
				returnNonEmptyString(quotaCheck(organization.QuotaProject)),
				returnNonEmptyString(quotaCheck(organization.QuotaGroup)),
				returnNonEmptyString(quotaCheck(organization.QuotaNotification)),
				returnNonEmptyString(quotaCheck(organization.QuotaEnvironment)),
				returnNonEmptyString(quotaCheck(organization.QuotaRoute)),
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

var listOrganizationVariablesCmd = &cobra.Command{
	Use:     "organization-variables",
	Aliases: []string{"org-v"},
	Short:   "List variables for an organization (alias: org-v)",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		organizationName, err := cmd.Flags().GetString("organization-name")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Organization name", organizationName); err != nil {
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
		envvars, err := lagoon.GetEnvVariablesByOrganizationName(context.TODO(), organizationName, lc)
		if err != nil {
			return err
		}
		data := []output.Data{}
		for _, envvar := range *envvars {
			env := []string{
				returnNonEmptyString(fmt.Sprintf("%v", envvar.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", organizationName)),
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
			"Organization",
		}
		header = append(header, "Scope")
		header = append(header, "Name")
		if reveal {
			header = append(header, "Value")
		}
		if len(data) == 0 {
			return handleNilResults("There are no variables for organization '%s'\n", cmd, organizationName)
		}
		r := output.RenderOutput(output.Table{
			Header: header,
			Data:   data,
		}, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		return nil
	},
}

var listEnvironmentServicesCmd = &cobra.Command{
	Use:     "environment-services",
	Aliases: []string{"es"},
	Short:   "Get information about an environments services",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
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

		project, err := lagoon.GetProjectByName(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		environment, err := lagoon.GetEnvironmentByName(context.TODO(), cmdProjectEnvironment, project.ID, lc)
		if err != nil {
			return err
		}

		if project.Name == "" || environment.Name == "" {
			if project.Name == "" {
				return handleNilResults("Project '%s' not found\n", cmd, cmdProjectName)
			} else {
				return handleNilResults("Environment '%s' not found in project '%s'\n", cmd, cmdProjectEnvironment, cmdProjectName)
			}
		}

		data := []output.Data{}
		envHeader := []string{"EnvironmentID", "EnvironmentName", "ServiceID", "ServiceName", "ServiceType", "Containers", "Updated", "Created"}
		for _, es := range environment.Services {
			containers := []string{}
			for _, c := range es.Containers {
				containers = append(containers, c.Name)
			}
			envData := []string{
				returnNonEmptyString(fmt.Sprintf("%d", environment.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", environment.Name)),
				returnNonEmptyString(fmt.Sprintf("%d", es.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", es.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", es.Type)),
				returnNonEmptyString(fmt.Sprintf("%v", strings.Join(containers, ","))),
				returnNonEmptyString(fmt.Sprintf("%v", es.Updated)),
				returnNonEmptyString(fmt.Sprintf("%v", es.Created)),
			}
			data = append(data, envData)
		}
		dataMain := output.Table{
			Header: envHeader,
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
	listDeploymentsCmd.Flags().Bool("wide", false, "Display additional information about deployments")
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
	listCmd.AddCommand(listOrganizationUsersCmd)
	listCmd.AddCommand(listOrganizationAdminsCmd)
	listCmd.AddCommand(listOrganizationGroupsCmd)
	listCmd.AddCommand(listOrganizationDeployTargetsCmd)
	listCmd.AddCommand(listOrganizationsCmd)
	listCmd.AddCommand(listOrganizationVariablesCmd)
	listCmd.AddCommand(listEnvironmentServicesCmd)
	listAllUsersCmd.Flags().StringP("email", "E", "", "The email address of a user")
	listUsersGroupsCmd.Flags().StringP("email", "E", "", "The email address of a user")
	listCmd.Flags().BoolVarP(&listAllProjects, "all-projects", "", false, "All projects (if supported)")
	listUsersCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group to list users in")
	listGroupProjectsCmd.Flags().StringP("name", "N", "", "Name of the group to list projects in")
	listGroupProjectsCmd.Flags().BoolP("all-projects", "", false, "All projects")
	listVariablesCmd.Flags().BoolP("reveal", "", false, "Reveal the variable values")
	listOrganizationProjectsCmd.Flags().StringP("organization-name", "O", "", "Name of the organization to list associated projects for")
	listOrganizationUsersCmd.Flags().StringP("organization-name", "O", "", "Name of the organization to list associated users for")
	listOrganizationAdminsCmd.Flags().StringP("organization-name", "O", "", "Name of the organization to list associated users for")
	listOrganizationGroupsCmd.Flags().StringP("organization-name", "O", "", "Name of the organization to list associated groups for")
	listOrganizationDeployTargetsCmd.Flags().StringP("organization-name", "O", "", "Name of the organization to list associated deploy targets for")
	listOrganizationDeployTargetsCmd.Flags().Uint("id", 0, "ID of the organization to list associated deploy targets for")
	listOrganizationVariablesCmd.Flags().StringP("organization-name", "O", "", "Name of the organization to list associated variables for")
	listOrganizationVariablesCmd.Flags().BoolP("reveal", "", false, "Reveal the variable values")
	listDeployTargetsCmd.Flags().Bool("wide", false, "Display additional information about deploytargets")
	listDeployTargetsCmd.Flags().Bool("show-token", false, "Display the token for deploytargets")
	listProjectsCmd.Flags().Bool("wide", false, "Display additional information about projects")
	listEnvironmentsCmd.Flags().Bool("wide", false, "Display additional information about environments")
}
