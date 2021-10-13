package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

var runFactsGatherer = &cobra.Command{
	Use:     "factsgatherer",
	Aliases: []string{"fg"},
	DisableSuggestions: true,
	SuggestionsMinimumDistance: 1,
	Short:   "Run the facts gatherer on an environment",
	Long: `Run the facts gatherer command on an environment
The following are supported methods to use
Direct:
  lagoon run factsgatherer -p example -e main -S cli
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if cmdProjectName == "" || cmdProjectEnvironment == "" {
			fmt.Println("Missing arguments: Project name and/or environment name are not defined")
			cmd.Help()
			os.Exit(1)
		}

		var factsGathererCmd = fmt.Sprintf(`
#!/bin/sh

# Install lagoon-facts binary from Github
set -e

PROJECT=%s
ENVIRONMENT=%s

OS_TYPE=linux
BIN_PATH=/tmp
LATEST_PATH=https://api.github.com/repos/uselagoon/lagoon-facts-app/releases/latest
#BIN_PATH=/usr/bin/

echo "Project: " "$PROJECT"
echo "Environment: " "$ENVIRONMENT"
echo "Binary path: " "$BIN_PATH"

echo "Downloading lagoon-facts from: $LATEST_PATH"

# Get the URL of the latest release.
DOWNLOAD_PATH=$(curl -L -s $LATEST_PATH | grep "browser_download_url" | cut -d '"' -f 4)

# Download and extract the release to the build dir.
wget -O $BIN_PATH/lagoon-facts "$DOWNLOAD_PATH" && chmod +x $BIN_PATH/lagoon-facts

# Run
if $BIN_PATH/lagoon-facts gather -p $PROJECT -e $ENVIRONMENT -v; then
    echo "Pushing facts succeeded"
else
    echo "Pushing facts failed"
fi

# Clean up
#rm -f $BIN_PATH/lagoon-facts
`, cmdProjectName, cmdProjectEnvironment)

		task := api.Task{
			Name:    "Lagoon Facts Gatherer",
			Command: factsGathererCmd,
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

var invokeDefinedTask = &cobra.Command{
	Use:     "invoke",
	Aliases: []string{"i"},
	Short:   "",
	Long: `Invoke a task registered against an environment
The following are supported methods to use
Direct:
 lagoon run invoke -p example -e main -N "advanced task name"
`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmdProjectName == "" || cmdProjectEnvironment == "" || invokedTaskName == "" {
			fmt.Println("Missing arguments: Project name, environment name, or task command are not defined")
			cmd.Help()
			os.Exit(1)
		}

		taskResult, err := eClient.InvokeAdvancedTaskDefinition(cmdProjectName, cmdProjectEnvironment, invokedTaskName)
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
	taskName        string
	invokedTaskName string
	taskService     string
	taskCommand     string
	taskCommandFile string
)

func init() {
	invokeDefinedTask.Flags().StringVarP(&invokedTaskName, "name", "N", "", "Name of the task that will be invoked")
	runCustomTask.Flags().StringVarP(&taskName, "name", "N", "Custom Task", "Name of the task that will show in the UI (default: Custom Task)")
	runCustomTask.Flags().StringVarP(&taskService, "service", "S", "cli", "Name of the service (cli, nginx, other) that should run the task (default: cli)")
	runCustomTask.Flags().StringVarP(&taskCommand, "command", "c", "", "The command to run in the task")
	runCustomTask.Flags().StringVarP(&taskCommandFile, "script", "s", "", "Path to bash script to run (will use this before command(-c) if both are defined)")
}
