package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/internal/lagoon"
	"github.com/uselagoon/lagoon-cli/internal/lagoon/client"
	"github.com/uselagoon/lagoon-cli/pkg/api"
	"github.com/uselagoon/lagoon-cli/pkg/output"
	"io/ioutil"
	"os"
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
		if taskID == 0 {
			return fmt.Errorf("Missing arguments: ID is not defined")
		}
		current := lagoonCLIConfig.Current
		lc := client.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIConfig.Lagoons[current].Token,
			lagoonCLIConfig.Lagoons[current].Version,
			lagoonCLIVersion,
			debug)
		result, err := lagoon.TaskByID(context.TODO(), taskID, lc)
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
		if cmdProjectName == "" {
			return fmt.Errorf("Missing arguments: Project name is not defined")
		}
		if yesNo(fmt.Sprintf("You are attempting to run the active/standby switch for project '%s', are you sure?", cmdProjectName)) {
			current := lagoonCLIConfig.Current
			lc := client.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIConfig.Lagoons[current].Token,
				lagoonCLIConfig.Lagoons[current].Version,
				lagoonCLIVersion,
				debug)
			result, err := lagoon.ActiveStandbySwitch(context.TODO(), cmdProjectName, lc)
			if err != nil {
				return err
			}
			fmt.Println(fmt.Sprintf(`Created a new task with ID %d
You can use the following command to query the task status:
lagoon -l %s get task-by-id --id %d --logs`, result.ID, current, result.ID))
		}
		return nil
	},
}

var runDrushArchiveDump = &cobra.Command{
	Use:     "drush-archivedump",
	Aliases: []string{"dard"},
	Short:   "Run a drush archive dump on an environment",
	Run: func(cmd *cobra.Command, args []string) {
		if cmdProjectName == "" || cmdProjectEnvironment == "" {
			fmt.Println("Missing arguments: Project name or environment name are not defined")
			cmd.Help()
			os.Exit(1)
		}
		taskResult, err := eClient.RunDrushArchiveDump(cmdProjectName, cmdProjectEnvironment)
		handleError(err)
		var resultMap map[string]interface{}
		err = json.Unmarshal([]byte(taskResult), &resultMap)
		handleError(err)
		resultData := output.Result{
			Result:     "success",
			ResultData: resultMap,
		}
		output.RenderResult(resultData, outputOptions)
	},
}

var runDrushSQLDump = &cobra.Command{
	Use:     "drush-sqldump",
	Aliases: []string{"dsqld"},
	Short:   "Run a drush sql dump on an environment",
	Run: func(cmd *cobra.Command, args []string) {
		if cmdProjectName == "" || cmdProjectEnvironment == "" {
			fmt.Println("Missing arguments: Project name or environment name are not defined")
			cmd.Help()
			os.Exit(1)
		}
		taskResult, err := eClient.RunDrushSQLDump(cmdProjectName, cmdProjectEnvironment)
		handleError(err)
		var resultMap map[string]interface{}
		err = json.Unmarshal([]byte(taskResult), &resultMap)
		handleError(err)
		resultData := output.Result{
			Result:     "success",
			ResultData: resultMap,
		}
		output.RenderResult(resultData, outputOptions)
	},
}

var runDrushCacheClear = &cobra.Command{
	Use:     "drush-cacheclear",
	Aliases: []string{"dcc"},
	Short:   "Run a drush cache clear on an environment",
	Run: func(cmd *cobra.Command, args []string) {
		if cmdProjectName == "" || cmdProjectEnvironment == "" {
			fmt.Println("Missing arguments: Project name or environment name are not defined")
			cmd.Help()
			os.Exit(1)
		}
		taskResult, err := eClient.RunDrushCacheClear(cmdProjectName, cmdProjectEnvironment)
		handleError(err)
		var resultMap map[string]interface{}
		err = json.Unmarshal([]byte(taskResult), &resultMap)
		handleError(err)
		resultData := output.Result{
			Result:     "success",
			ResultData: resultMap,
		}
		output.RenderResult(resultData, outputOptions)
	},
}

var runDefinedTask = &cobra.Command{
	Use:     "task",
	Aliases: []string{"i"},
	Short:   "Run a custom task registered against an environment",
	Long: `Run a custom task registered against an environment
The following are supported methods to use
Direct:
 lagoon run run -p example -e main -N "advanced task name" [--argument=NAME=VALUE|..]
`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmdProjectName == "" || cmdProjectEnvironment == "" || invokedTaskName == "" {
			fmt.Println("Missing arguments: Project name, environment name, or task command are not defined")
			cmd.Help()
			os.Exit(1)
		}

		taskArguments, err := splitInvokeTaskArguments(invokedTaskArguments)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		taskResult, err := eClient.InvokeAdvancedTaskDefinition(cmdProjectName, cmdProjectEnvironment, invokedTaskName, taskArguments)
		handleError(err)
		var resultMap map[string]interface{}
		err = json.Unmarshal([]byte(taskResult), &resultMap)
		handleError(err)
		resultData := output.Result{
			Result:     "success",
			ResultData: resultMap,
		}
		output.RenderResult(resultData, outputOptions)
	},
}

var invokeInteractiveTask = &cobra.Command{
	Use:     "interactive",
	Aliases: []string{"i"},
	Short:   "Interactively run a custom task against an environment",
	Long: `Interactively run a custom task against an environment
Provides prompts for arguments
example:
 lagoon run invoke interactive -p example -e main
`,
	Run: func(cmd *cobra.Command, args []string) {
		debug, err := cmd.Flags().GetBool("debug")
		if cmdProjectName == "" || cmdProjectEnvironment == "" {
			fmt.Println("Missing arguments: Project name or environment name are not defined")
			cmd.Help()
			os.Exit(1)
		}

		//TODO: get project id for subsequent queries
		current := lagoonCLIConfig.Current
		lc := client.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIConfig.Lagoons[current].Token,
			lagoonCLIConfig.Lagoons[current].Version,
			lagoonCLIVersion,
			debug)
		project, err := lagoon.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//TODO: get list of tasks and their arguments

		environment, err := lagoon.TasksForEnvironment(context.TODO(), project.ID, cmdProjectEnvironment, lc)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if len(environment.AdvancedTasks) == 0 {
			fmt.Println("No custom tasks registered against environment - exiting")
			return
		}

		prompt := promptui.Select{
			Label: "Select Task",
			Items: environment.AdvancedTasks,
			Templates: &promptui.SelectTemplates{
				Active:   fmt.Sprintf("%s {{ .Name | underline }} -- {{ .Description }}", promptui.IconSelect),
				Inactive: "{{ .Name }} -- {{ .Description }}",
				Selected: fmt.Sprintf("%s {{ .Name | green }}", promptui.IconGood),
			},
		}

		taskIndex, _, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		//TODO: fill out any arguments
		task := environment.AdvancedTasks[taskIndex]
		taskArguments := map[string]string{}
		taskArgumentString := ""
		for _, v := range task.AdvancedTaskDefinitionArguments {
			if len(v.Range) != 0 { //we have a selection
				prompt = promptui.Select{
					Label: v.DisplayName,
					Items: v.Range,
				}
				_, argumentValue, err := prompt.Run()
				if err != nil {
					fmt.Printf("Prompt failed %v\n", err)
					return
				}
				taskArguments[v.Name] = argumentValue
				taskArgumentString = taskArgumentString + fmt.Sprintf("%v(%v) : %v", v.DisplayName, v.Name, argumentValue)
			} else { // standard prompt
				prompt := promptui.Prompt{
					Label: fmt.Sprintf("%v", v.DisplayName),
				}
				argumentValue, err := prompt.Run()

				if err != nil {
					fmt.Printf("Prompt failed %v\n", err)
					return
				}
				taskArguments[v.Name] = argumentValue
				taskArgumentString = taskArgumentString + fmt.Sprintf(", %v(%v) : %v", v.DisplayName, v.Name, argumentValue)
			}
		}

		if !yesNo(fmt.Sprintf("Run command `%v` with arguments %v", task.Name, taskArgumentString)) {
			fmt.Println("Exiting")
			return
		}

		taskResult, err := eClient.InvokeAdvancedTaskDefinition(cmdProjectName, cmdProjectEnvironment, task.Name, taskArguments)
		handleError(err)
		var resultMap map[string]interface{}
		err = json.Unmarshal([]byte(taskResult), &resultMap)
		handleError(err)
		resultData := output.Result{
			Result:     "success",
			ResultData: resultMap,
		}
		output.RenderResult(resultData, outputOptions)
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
	Run: func(cmd *cobra.Command, args []string) {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			// check if we are getting data froms stdin
			scanner := bufio.NewScanner(os.Stdin)
			taskCommand = ""
			for scanner.Scan() {
				taskCommand = taskCommand + scanner.Text() + "\n"
			}
			if err := scanner.Err(); err != nil {
				// fmt.Fprintln(os.Stderr, "reading standard input:", err)
				handleError(errors.New("reading standard input:" + err.Error()))
			}
		} else {
			// otherwise we can read from a file
			if taskCommandFile != "" {
				taskCommandBytes, err := ioutil.ReadFile(taskCommandFile) // just pass the file name
				handleError(err)
				taskCommand = string(taskCommandBytes)
			}
		}

		if cmdProjectName == "" || cmdProjectEnvironment == "" || taskCommand == "" {
			fmt.Println("Missing arguments: Project name, environment name, or task command are not defined")
			cmd.Help()
			os.Exit(1)
		}
		task := api.Task{
			Name:    taskName,
			Command: taskCommand,
			Service: taskService,
		}
		taskResult, err := eClient.RunCustomTask(cmdProjectName, cmdProjectEnvironment, task)
		handleError(err)
		var resultMap map[string]interface{}
		err = json.Unmarshal([]byte(taskResult), &resultMap)
		handleError(err)
		resultData := output.Result{
			Result:     "success",
			ResultData: resultMap,
		}
		output.RenderResult(resultData, outputOptions)
	},
}

var (
	taskName             string
	invokedTaskName      string
	invokedTaskArguments []string
	taskService          string
	taskCommand          string
	taskCommandFile      string
)

func init() {
	//register sub tasks
	runDefinedTask.AddCommand(invokeInteractiveTask)
	runDefinedTask.Flags().StringVarP(&invokedTaskName, "name", "N", "", "Name of the task that will be run")
	runDefinedTask.Flags().StringSliceVar(&invokedTaskArguments, "argument", []string{}, "Arguments to be passed to custom task, of the form NAME=VALUE")
	runCustomTask.Flags().StringVarP(&taskName, "name", "N", "Custom Task", "Name of the task that will show in the UI (default: Custom Task)")
	runCustomTask.Flags().StringVarP(&taskService, "service", "S", "cli", "Name of the service (cli, nginx, other) that should run the task (default: cli)")
	runCustomTask.Flags().StringVarP(&taskCommand, "command", "c", "", "The command to run in the task")
	runCustomTask.Flags().StringVarP(&taskCommandFile, "script", "s", "", "Path to bash script to run (will use this before command(-c) if both are defined)")
}
