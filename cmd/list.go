package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/uselagoon/lagoon-cli/internal/lagoon"
	"github.com/uselagoon/lagoon-cli/internal/lagoon/client"
	"github.com/uselagoon/lagoon-cli/internal/schema"
	"github.com/uselagoon/lagoon-cli/pkg/api"
	"github.com/uselagoon/lagoon-cli/pkg/output"
	l "github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
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
	Run: func(cmd *cobra.Command, args []string) {
		returnedJSON, err := pClient.ListAllProjects()
		handleError(err)
		var dataMain output.Table
		err = json.Unmarshal([]byte(returnedJSON), &dataMain)
		handleError(err)
		if len(dataMain.Data) == 0 {
			output.RenderInfo("No access to any projects in Lagoon", outputOptions)
			os.Exit(0)
		}
		output.RenderOutput(dataMain, outputOptions)

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
		lc := client.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIConfig.Lagoons[current].Token,
			lagoonCLIConfig.Lagoons[current].Version,
			lagoonCLIVersion,
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
	Run: func(cmd *cobra.Command, args []string) {
		returnedJSON, err := uClient.ListGroups("")
		handleError(err)
		var dataMain output.Table
		err = json.Unmarshal([]byte(returnedJSON), &dataMain)
		handleError(err)
		if len(dataMain.Data) == 0 {
			output.RenderInfo("This account is not in any groups", outputOptions)
			os.Exit(0)
		}
		output.RenderOutput(dataMain, outputOptions)

	},
}

var listGroupProjectsCmd = &cobra.Command{
	Use:     "group-projects",
	Aliases: []string{"gp"},
	Short:   "List projects in a group (alias: gp)",
	Run: func(cmd *cobra.Command, args []string) {
		if !listAllProjects {
			if groupName == "" {
				fmt.Println("Missing arguments: Group name is not defined")
				cmd.Help()
				os.Exit(1)
			}
		}
		var returnedJSON []byte
		var err error
		if listAllProjects {
			returnedJSON, err = uClient.ListGroupProjects("", listAllProjects)
		} else {
			returnedJSON, err = uClient.ListGroupProjects(groupName, listAllProjects)
		}
		handleError(err)
		var dataMain output.Table
		err = json.Unmarshal([]byte(returnedJSON), &dataMain)
		handleError(err)
		if len(dataMain.Data) == 0 {
			if !listAllProjects {
				output.RenderInfo(fmt.Sprintf("There are no projects in group '%s'", groupName), outputOptions)
			} else {
				output.RenderInfo("There are no projects in any groups", outputOptions)
			}
			os.Exit(0)
		}
		output.RenderOutput(dataMain, outputOptions)

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
		if cmdProjectName == "" {
			fmt.Println("Missing arguments: Project name is not defined")
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
		environments, err := l.GetEnvironmentsByProjectName(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
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
		if cmdProjectName == "" {
			fmt.Println("Missing arguments: Project name is not defined")
			cmd.Help()
			os.Exit(1)
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
		lc := client.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIConfig.Lagoons[current].Token,
			lagoonCLIConfig.Lagoons[current].Version,
			lagoonCLIVersion,
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
	Run: func(cmd *cobra.Command, args []string) {
		if cmdProjectName == "" || cmdProjectEnvironment == "" {
			fmt.Println("Missing arguments: Project name or environment name is not defined")
			cmd.Help()
			os.Exit(1)
		}
		returnedJSON, err := eClient.GetEnvironmentDeployments(cmdProjectName, cmdProjectEnvironment)
		handleError(err)

		var dataMain output.Table
		err = json.Unmarshal([]byte(returnedJSON), &dataMain)
		handleError(err)
		if len(dataMain.Data) == 0 {
			output.RenderInfo(fmt.Sprintf("There are no deployments for environment '%s' in project '%s'", cmdProjectEnvironment, cmdProjectName), outputOptions)
			os.Exit(0)
		}
		output.RenderOutput(dataMain, outputOptions)
	},
}

var listTasksCmd = &cobra.Command{
	Use:     "tasks",
	Aliases: []string{"t"},
	Short:   "List tasks for an environment (alias: t)",
	Run: func(cmd *cobra.Command, args []string) {
		if cmdProjectName == "" || cmdProjectEnvironment == "" {
			fmt.Println("Missing arguments: Project name or environment name is not defined")
			cmd.Help()
			os.Exit(1)
		}
		returnedJSON, err := eClient.GetEnvironmentTasks(cmdProjectName, cmdProjectEnvironment)
		handleError(err)

		var dataMain output.Table
		err = json.Unmarshal([]byte(returnedJSON), &dataMain)
		handleError(err)
		if len(dataMain.Data) == 0 {
			output.RenderInfo(fmt.Sprintf("There are no tasks for environment '%s' in project '%s'", cmdProjectEnvironment, cmdProjectName), outputOptions)
			os.Exit(0)
		}
		output.RenderOutput(dataMain, outputOptions)
	},
}

var listUsersCmd = &cobra.Command{
	//@TODO: once individual user interaction comes in, this will need to be adjusted
	Use:     "users",
	Aliases: []string{"u"},
	Short:   "List all users in groups (alias: u)",
	Long:    `List all users in groups in lagoon, this only shows users that are in groups.`,
	Run: func(cmd *cobra.Command, args []string) {
		returnedJSON, err := uClient.ListUsers(groupName)
		handleError(err)

		var dataMain output.Table
		err = json.Unmarshal([]byte(returnedJSON), &dataMain)
		handleError(err)
		if len(dataMain.Data) == 0 {
			output.RenderInfo("There are no users in any groups", outputOptions)
			os.Exit(0)
		}
		output.RenderOutput(dataMain, outputOptions)

	},
}

var listInvokableTasks = &cobra.Command{
	Use:     "invokable-tasks",
	Aliases: []string{"dcc"},
	Short:   "Print a list of invokable tasks",
	Long:    "Print a list of invokable user defined tasks registered against an environment",
	Run: func(cmd *cobra.Command, args []string) {
		if cmdProjectName == "" || cmdProjectEnvironment == "" {
			fmt.Println("Missing arguments: Project name or environment name are not defined")
			cmd.Help()
			os.Exit(1)
		}
		taskResult, err := eClient.ListInvokableAdvancedTaskDefinitions(cmdProjectName, cmdProjectEnvironment)
		handleError(err)

		var taskList []api.AdvancedTask
		err = json.Unmarshal([]byte(taskResult), &taskList)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var taskListData []output.Data
		for _, task := range taskList {
			taskListData = append(taskListData, []string{task.Name, task.Description})
		}

		var dataMain output.Table
		dataMain.Header = []string{"Task Name", "Description"}

		dataMain.Data = taskListData

		if len(dataMain.Data) == 0 {
			output.RenderInfo("There are no user defined tasks for this environment", outputOptions)
			os.Exit(0)
		}
		output.RenderOutput(dataMain, outputOptions)
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
		if cmdProjectName == "" {
			fmt.Println("Missing arguments: Project is not defined")
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
		projectGroups, err := l.GetProjectGroups(context.TODO(), cmdProjectName, lc)
		handleError(err)

		if len(projectGroups.Groups) == 0 {
			output.RenderInfo(fmt.Sprintf("There are no projects in group '%s'", groupName), outputOptions)
			os.Exit(0)
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
	listCmd.Flags().BoolVarP(&listAllProjects, "all-projects", "", false, "All projects (if supported)")
	listUsersCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group to list users in (if not specified, will default to all groups)")
	listGroupProjectsCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group to list projects in")
	listVariablesCmd.Flags().BoolP("reveal", "", false, "Reveal the variable values")
}
