package cmd

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/output"
	"github.com/spf13/cobra"
)

var runDrushArchiveDump = &cobra.Command{
	Use:     "drush-archivedump",
	Aliases: []string{"dard"},
	Short:   "Run a drush archive dump on an environment",
	Run: func(cmd *cobra.Command, args []string) {
		if cmdProjectName == "" || cmdProjectEnvironment == "" {
			fmt.Println("Not enough arguments. Requires: project name and environment name")
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
			fmt.Println("Not enough arguments. Requires: project name and environment name")
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
			fmt.Println("Not enough arguments. Requires: project name and environment name")
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

var runCustomTask = &cobra.Command{
	Use:     "custom",
	Aliases: []string{"c"},
	Short:   "Run a custom command on an environment",
	Long: `Run a custom command on an environment
The following are supported methods to use
Direct:
  lagoon run custom -p example -e master -N "My Task" -S cli -c "ps -ef"

STDIN:
  cat /path/to/my-script.sh | lagoon run custom -p example -e master -N "My Task" -S cli

Path:
  lagoon run custom -p example -e master -N "My Task" -S cli -s /path/to/my-script.sh
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
			fmt.Println("Not enough arguments. Requires: project name, environment name, and command")
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
	taskService     string
	taskCommand     string
	taskCommandFile string
)

func init() {
	runCustomTask.Flags().StringVarP(&taskName, "name", "N", "Custom Task", "Name of the task that will show in the UI (default: Custom Task)")
	runCustomTask.Flags().StringVarP(&taskService, "service", "S", "cli", "Name of the service (cli, nginx, other) that should run the task (default: cli)")
	runCustomTask.Flags().StringVarP(&taskCommand, "command", "c", "", "The command to run in the task")
	runCustomTask.Flags().StringVarP(&taskCommandFile, "script", "s", "", "Path to bash script to run (will use this before command(-c) if both are defined)")
}
