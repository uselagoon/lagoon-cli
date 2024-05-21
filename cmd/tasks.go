package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/pkg/output"

	l "github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	ls "github.com/uselagoon/machinery/api/schema"
)

var getTaskByID = &cobra.Command{
	Use:     "task-by-id",
	Short:   "Get information about a task by its ID",
	Long:    `Get information about a task by its ID`,
	Aliases: []string{"t", "tbi"},
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}

		taskID, err := cmd.Flags().GetInt("id")
		if err != nil {
			return err
		}
		showLogs, err := cmd.Flags().GetBool("logs")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("ID", strconv.Itoa(taskID)); err != nil {
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
		result, err := l.TaskByID(context.TODO(), taskID, lc)
		if err != nil {
			return err
		}
		dataMain := output.Table{
			Header: []string{
				"ID",
				"Name",
				"Status",
				"Created",
				"Started",
				"Completed",
			},
			Data: []output.Data{
				{
					fmt.Sprintf("%d", result.ID),
					returnNonEmptyString(result.Name),
					returnNonEmptyString(result.Status),
					returnNonEmptyString(result.Created),
					returnNonEmptyString(result.Started),
					returnNonEmptyString(result.Completed),
				},
			},
		}
		if showLogs {
			dataMain.Header = append(dataMain.Header, "Logs")
			dataMain.Data[0] = append(dataMain.Data[0], returnNonEmptyString(result.Logs))
		}
		output.RenderOutput(dataMain, outputOptions)
		return nil
	},
}

var runActiveStandbySwitch = &cobra.Command{
	Use:   "activestandby",
	Short: "Run the active/standby switch for a project",
	Long: `Run the active/standby switch for a project
You should only run this once and then check the status of the task that gets created.
If the task fails or fails to update, contact your Lagoon administrator for assistance.`,
	Aliases: []string{"as", "standby"},
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
		if yesNo(fmt.Sprintf("You are attempting to run the active/standby switch for project '%s', are you sure?", cmdProjectName)) {
			current := lagoonCLIConfig.Current
			token := lagoonCLIConfig.Lagoons[current].Token
			lc := lclient.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIVersion,
				lagoonCLIConfig.Lagoons[current].Version,
				&token,
				debug)
			result, err := l.ActiveStandbySwitch(context.TODO(), cmdProjectName, lc)
			if err != nil {
				return err
			}
			fmt.Printf("Created a new task with ID %d \nYou can use the following command to query the task status: \nlagoon -l %s get task-by-id --id %d --logs \n", result.ID, current, result.ID)
		}
		return nil
	},
}

var runDrushArchiveDump = &cobra.Command{
	Use:     "drush-archivedump",
	Aliases: []string{"dard"},
	Short:   "Run a drush archive dump on an environment",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requiredInputCheck("Project name", cmdProjectName, "Environment name", cmdProjectEnvironment); err != nil {
			return err
		}
		taskResult, err := eClient.RunDrushArchiveDump(cmdProjectName, cmdProjectEnvironment)
		if err != nil {
			return err
		}
		var resultMap map[string]interface{}
		err = json.Unmarshal([]byte(taskResult), &resultMap)
		if err != nil {
			return err
		}
		resultData := output.Result{
			Result:     "success",
			ResultData: resultMap,
		}
		output.RenderResult(resultData, outputOptions)
		return nil
	},
}

var runDrushSQLDump = &cobra.Command{
	Use:     "drush-sqldump",
	Aliases: []string{"dsqld"},
	Short:   "Run a drush sql dump on an environment",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requiredInputCheck("Project name", cmdProjectName, "Environment name", cmdProjectEnvironment); err != nil {
			return err
		}
		taskResult, err := eClient.RunDrushSQLDump(cmdProjectName, cmdProjectEnvironment)
		if err != nil {
			return err
		}
		var resultMap map[string]interface{}
		err = json.Unmarshal([]byte(taskResult), &resultMap)
		if err != nil {
			return err
		}
		resultData := output.Result{
			Result:     "success",
			ResultData: resultMap,
		}
		output.RenderResult(resultData, outputOptions)
		return nil
	},
}

var runDrushCacheClear = &cobra.Command{
	Use:     "drush-cacheclear",
	Aliases: []string{"dcc"},
	Short:   "Run a drush cache clear on an environment",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := requiredInputCheck("Project name", cmdProjectName, "Environment name", cmdProjectEnvironment); err != nil {
			return err
		}
		taskResult, err := eClient.RunDrushCacheClear(cmdProjectName, cmdProjectEnvironment)
		if err != nil {
			return err
		}
		var resultMap map[string]interface{}
		err = json.Unmarshal([]byte(taskResult), &resultMap)
		if err != nil {
			return err
		}
		resultData := output.Result{
			Result:     "success",
			ResultData: resultMap,
		}
		output.RenderResult(resultData, outputOptions)
		return nil
	},
}

var invokeDefinedTask = &cobra.Command{
	Use:     "invoke",
	Aliases: []string{"i"},
	Short:   "Invoke a task registered against an environment",
	Long: `Invoke a task registered against an environment
The following are supported methods to use
Direct:
 lagoon run invoke -p example -e main -N "advanced task name"
`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		invokedTaskName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Project name", cmdProjectName, "Environment name", cmdProjectEnvironment, "Task command", invokedTaskName); err != nil {
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

		project, err := l.GetProjectByName(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		environment, err := l.GetAdvancedTasksByEnvironment(context.TODO(), project.ID, cmdProjectEnvironment, lc)
		if err != nil {
			return err
		}

		var taskId uint
		for _, task := range environment.AdvancedTasks {
			if invokedTaskName == task.Name {
				taskId = uint(task.ID)
			}
		}

		taskResult, err := l.InvokeAdvancedTaskDefinition(context.TODO(), environment.ID, taskId, lc)
		if err != nil {
			return err
		}

		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"id":     taskResult.ID,
				"name":   taskResult.Name,
				"status": taskResult.Status,
			},
		}
		output.RenderResult(resultData, outputOptions)
		return nil
	},
}

var runCustomTask = &cobra.Command{
	Use:     "custom",
	Aliases: []string{"c"},
	Short:   "Run a custom command on an environment",
	Long: `Run a custom command on an environment
The following are supported methods to use
Direct:
  lagoon run custom -p example -e main -N "My Task" -S cli -c "ps -ef"

STDIN:
  cat /path/to/my-script.sh | lagoon run custom -p example -e main -N "My Task" -S cli

Path:
  lagoon run custom -p example -e main -N "My Task" -S cli -s /path/to/my-script.sh
`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		taskName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		taskService, err := cmd.Flags().GetString("service")
		if err != nil {
			return err
		}
		taskCommand, err := cmd.Flags().GetString("command")
		if err != nil {
			return err
		}
		taskCommandFile, err := cmd.Flags().GetString("script")
		if err != nil {
			return err
		}
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			// check if we are getting data froms stdin
			scanner := bufio.NewScanner(os.Stdin)
			taskCommand = ""
			for scanner.Scan() {
				taskCommand = taskCommand + scanner.Text() + "\n"
			}
			if err := scanner.Err(); err != nil {
				handleError(errors.New("reading standard input:" + err.Error()))
			}
		} else {
			// otherwise we can read from a file
			if taskCommandFile != "" {
				taskCommandBytes, err := os.ReadFile(taskCommandFile) // just pass the file name
				if err != nil {
					return err
				}
				taskCommand = string(taskCommandBytes)
			}
		}

		if err := requiredInputCheck("Project name", cmdProjectName, "Environment name", cmdProjectEnvironment, "Task command", taskCommand, "Task name", taskName, "Task service", taskService); err != nil {
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

		task := ls.Task{
			Name:    taskName,
			Command: taskCommand,
			Service: taskService,
		}
		fmt.Println(task.Name)
		project, err := l.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
		environment, err := l.GetEnvironmentByName(context.TODO(), cmdProjectEnvironment, project.ID, lc)
		taskResult, err := l.AddTask(context.TODO(), environment.ID, task, lc)
		if err != nil {
			return err
		}
		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"id": taskResult.ID,
			},
		}
		output.RenderResult(resultData, outputOptions)
		return nil
	},
}

var uploadFilesToTask = &cobra.Command{
	Use:     "task-files",
	Short:   "Upload files to a task by its ID",
	Long:    `Upload files to a task by its ID`,
	Aliases: []string{"tf"},
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}

		taskID, err := cmd.Flags().GetInt("id")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("ID", strconv.Itoa(taskID)); err != nil {
			return err
		}
		files, err := cmd.Flags().GetStringSlice("file")
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
		result, err := l.UploadFilesForTask(context.TODO(), taskID, files, lc)
		if err != nil {
			return err
		}
		taskFiles := []string{}
		for _, f := range result.Files {
			taskFiles = append(taskFiles, f.Filename)
		}
		dataMain := output.Table{
			Header: []string{
				"ID",
				"Name",
				"Files",
			},
			Data: []output.Data{
				{
					fmt.Sprintf("%d", result.ID),
					returnNonEmptyString(result.Name),
					returnNonEmptyString(strings.Join(taskFiles, ",")),
				},
			},
		}
		output.RenderOutput(dataMain, outputOptions)
		return nil
	},
}

func init() {
	uploadFilesToTask.Flags().IntP("id", "I", 0, "ID of the task")
	uploadFilesToTask.Flags().StringSliceP("file", "F", []string{}, "File to upload (add multiple flags to upload multiple files)")
	invokeDefinedTask.Flags().StringP("name", "N", "", "Name of the task that will be invoked")
	runCustomTask.Flags().StringP("name", "N", "Custom Task", "Name of the task that will show in the UI (default: Custom Task)")
	runCustomTask.Flags().StringP("service", "S", "cli", "Name of the service (cli, nginx, other) that should run the task (default: cli)")
	runCustomTask.Flags().StringP("command", "c", "", "The command to run in the task")
	runCustomTask.Flags().StringP("script", "s", "", "Path to bash script to run (will use this before command(-c) if both are defined)")
}
