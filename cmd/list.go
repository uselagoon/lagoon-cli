package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/uselagoon/lagoon-cli/internal/lagoon"
	"github.com/uselagoon/lagoon-cli/internal/lagoon/client"
	"github.com/uselagoon/lagoon-cli/pkg/api"
	"github.com/uselagoon/lagoon-cli/pkg/output"
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
	Short: "List projects, deployments, variables or notifications",
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
	Aliases: []string{"dt"},
	Short:   "List all deploytargets you have access to",
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
				returnNonEmptyString(fmt.Sprintf("%v", deploytarget.FriendlyName)),
				returnNonEmptyString(fmt.Sprintf("%v", deploytarget.CloudProvider)),
				returnNonEmptyString(fmt.Sprintf("%v", deploytarget.CloudRegion)),
			})
		}
		output.RenderOutput(output.Table{
			Header: []string{
				"ID",
				"Name",
				"FriendlyName",
				"CloudProvider",
				"CloudRegion",
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

var listProjectCmd = &cobra.Command{
	Use:     "environments",
	Aliases: []string{"e"},
	Short:   "List environments for a project (alias: e)",
	Run: func(cmd *cobra.Command, args []string) {
		if cmdProjectName == "" {
			fmt.Println("Missing arguments: Project name is not defined")
			cmd.Help()
			os.Exit(1)
		}
		returnedJSON, err := pClient.ListEnvironmentsForProject(cmdProjectName)
		handleError(err)
		var dataMain output.Table
		err = json.Unmarshal([]byte(returnedJSON), &dataMain)
		handleError(err)
		if len(dataMain.Data) == 0 {
			output.RenderInfo(fmt.Sprintf("There are no environments for project '%s'", cmdProjectName), outputOptions)
			os.Exit(0)
		}
		output.RenderOutput(dataMain, outputOptions)

	},
}

var listVariablesCmd = &cobra.Command{
	Use:     "variables",
	Aliases: []string{"v"},
	Short:   "List variables for a project or environment (alias: v)",
	Run: func(cmd *cobra.Command, args []string) {
		getListFlags := parseListFlags(*cmd.Flags())
		if cmdProjectName == "" {
			fmt.Println("Missing arguments: Project name is not defined")
			cmd.Help()
			os.Exit(1)
		}
		var returnedJSON []byte
		var err error
		if cmdProjectEnvironment != "" {
			returnedJSON, err = eClient.ListEnvironmentVariables(cmdProjectName, cmdProjectEnvironment, getListFlags.Reveal)
		} else {
			returnedJSON, err = pClient.ListProjectVariables(cmdProjectName, getListFlags.Reveal)
		}
		handleError(err)
		var dataMain output.Table
		err = json.Unmarshal([]byte(returnedJSON), &dataMain)
		handleError(err)
		if len(dataMain.Data) == 0 {
			if cmdProjectEnvironment != "" {
				output.RenderInfo(fmt.Sprintf("There are no variables for environment '%s' in project '%s'", cmdProjectEnvironment, cmdProjectName), outputOptions)
			} else {
				output.RenderInfo(fmt.Sprintf("There are no variables for project '%s'", cmdProjectName), outputOptions)
			}
			os.Exit(0)
		}
		output.RenderOutput(dataMain, outputOptions)
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

func init() {
	listCmd.AddCommand(listDeployTargetsCmd)
	listCmd.AddCommand(listDeploymentsCmd)
	listCmd.AddCommand(listGroupsCmd)
	listCmd.AddCommand(listGroupProjectsCmd)
	listCmd.AddCommand(listProjectCmd)
	listCmd.AddCommand(listProjectsCmd)
	listCmd.AddCommand(listRocketChatsCmd)
	listCmd.AddCommand(listSlackCmd)
	listCmd.AddCommand(listTasksCmd)
	listCmd.AddCommand(listUsersCmd)
	listCmd.AddCommand(listVariablesCmd)
	listCmd.AddCommand(listInvokableTasks)
	listCmd.AddCommand(listBackupsCmd)
	listCmd.AddCommand(listDeployTargetConfigsCmd)
	listCmd.Flags().BoolVarP(&listAllProjects, "all-projects", "", false, "All projects (if supported)")
	listUsersCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group to list users in (if not specified, will default to all groups)")
	listGroupProjectsCmd.Flags().StringVarP(&groupName, "name", "N", "", "Name of the group to list users in (if not specified, will default to all groups)")
	listVariablesCmd.Flags().BoolVarP(&revealValue, "reveal", "", false, "Reveal the variable values")
}
