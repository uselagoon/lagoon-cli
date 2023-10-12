package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/uselagoon/lagoon-cli/pkg/output"
	l "github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
)

// GetFlags .
type GetFlags struct {
	Project     string `json:"project,omitempty"`
	Environment string `json:"environment,omitempty"`
	RemoteID    string `json:"remoteid,omitempty"`
}

func parseGetFlags(flags pflag.FlagSet) GetFlags {
	configMap := make(map[string]interface{})
	flags.VisitAll(func(f *pflag.Flag) {
		if flags.Changed(f.Name) {
			configMap[f.Name] = f.Value
		}
	})
	jsonStr, _ := json.Marshal(configMap)
	parsedFlags := GetFlags{}
	json.Unmarshal(jsonStr, &parsedFlags)
	return parsedFlags
}

var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "Get info on a resource",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateToken(lagoonCLIConfig.Current) // get a new token if the current one is invalid
	},
}

var getProjectCmd = &cobra.Command{
	Use:     "project",
	Aliases: []string{"p"},
	Short:   "Get details about a project",
	Run: func(cmd *cobra.Command, args []string) {
		getProjectFlags := parseGetFlags(*cmd.Flags())
		if getProjectFlags.Project == "" {
			fmt.Println("Missing arguments: Project name is not defined")
			cmd.Help()
			os.Exit(1)
		}
		returnedJSON, err := pClient.GetProjectInfo(getProjectFlags.Project)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		var dataMain output.Table
		err = json.Unmarshal([]byte(returnedJSON), &dataMain)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		if len(dataMain.Data) == 0 {
			output.RenderInfo(fmt.Sprintf("No details for project '%s'", getProjectFlags.Project), outputOptions)
			os.Exit(0)
		}
		output.RenderOutput(dataMain, outputOptions)

	},
}

var getDeploymentByNameCmd = &cobra.Command{
	Use:     "deployment",
	Aliases: []string{"d"},
	Short:   "Get a deployment by name",
	Long: `Get a deployment by name
This returns information about a deployment, the logs of this build can also be retrieved`,
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}

		buildName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		showLogs, err := cmd.Flags().GetBool("logs")
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
		deployment, err := l.GetDeploymentByName(context.TODO(), cmdProjectName, cmdProjectEnvironment, buildName, showLogs, lc)
		if err != nil {
			return err
		}
		if showLogs {
			dataMain := output.Table{
				Header: []string{
					"Logs",
				},
				Data: []output.Data{
					{
						returnNonEmptyString(deployment.BuildLog),
					},
				},
			}
			output.RenderOutput(dataMain, outputOptions)
			return nil
		}
		dataMain := output.Table{
			Header: []string{
				"ID",
				"RemoteID",
				"Name",
				"Status",
				"Created",
				"Started",
				"Completed",
			},
			Data: []output.Data{
				{
					returnNonEmptyString(fmt.Sprintf("%v", deployment.ID)),
					returnNonEmptyString(fmt.Sprintf("%v", deployment.RemoteID)),
					returnNonEmptyString(fmt.Sprintf("%v", deployment.Name)),
					returnNonEmptyString(fmt.Sprintf("%v", deployment.Status)),
					returnNonEmptyString(fmt.Sprintf("%v", deployment.Created)),
					returnNonEmptyString(fmt.Sprintf("%v", deployment.Started)),
					returnNonEmptyString(fmt.Sprintf("%v", deployment.Completed)),
				},
			},
		}
		output.RenderOutput(dataMain, outputOptions)
		return nil
	},
}
var getEnvironmentCmd = &cobra.Command{
	Use:     "environment",
	Aliases: []string{"e"},
	Short:   "Get details about an environment",
	Run: func(cmd *cobra.Command, args []string) {
		if cmdProjectName == "" || cmdProjectEnvironment == "" {
			fmt.Println("Missing arguments: Project name or environment name is not defined")
			cmd.Help()
			os.Exit(1)
		}
		returnedJSON, err := eClient.GetEnvironmentInfo(cmdProjectName, cmdProjectEnvironment)
		handleError(err)
		var dataMain output.Table
		err = json.Unmarshal([]byte(returnedJSON), &dataMain)
		handleError(err)
		if len(dataMain.Data) == 0 {
			output.RenderInfo(fmt.Sprintf("No environment '%s' for project '%s'", cmdProjectEnvironment, cmdProjectName), outputOptions)
			os.Exit(0)
		}
		output.RenderOutput(dataMain, outputOptions)

	},
}

var getProjectKeyCmd = &cobra.Command{
	Use:     "project-key",
	Aliases: []string{"pk"},
	Short:   "Get a projects public key",
	Run: func(cmd *cobra.Command, args []string) {
		getProjectFlags := parseGetFlags(*cmd.Flags())
		if getProjectFlags.Project == "" {
			fmt.Println("Missing arguments: Project name is not defined")
			cmd.Help()
			os.Exit(1)
		}
		returnedJSON, err := pClient.GetProjectKey(getProjectFlags.Project, revealValue)
		handleError(err)
		var dataMain output.Table
		err = json.Unmarshal([]byte(returnedJSON), &dataMain)
		handleError(err)
		if len(dataMain.Data) == 0 {
			output.RenderInfo(fmt.Sprintf("No project-key for project '%s'", getProjectFlags.Project), outputOptions)
			os.Exit(0)
		}
		output.RenderOutput(dataMain, outputOptions)

	},
}

var getToken = &cobra.Command{
	Use:     "token",
	Aliases: []string{"tk"},
	Short:   "Generates a Lagoon auth token (for use in, for example, graphQL queries)",
	Run: func(cmd *cobra.Command, args []string) {
		token, err := retrieveTokenViaSsh()
		handleError(err)
		fmt.Println(token)
	},
}

func init() {
	getCmd.AddCommand(getAllUserKeysCmd)
	getCmd.AddCommand(getEnvironmentCmd)
	getCmd.AddCommand(getProjectCmd)
	getCmd.AddCommand(getProjectKeyCmd)
	getCmd.AddCommand(getUserKeysCmd)
	getCmd.AddCommand(getTaskByID)
	getCmd.AddCommand(getToken)
	getCmd.AddCommand(getDeploymentByNameCmd)
	getTaskByID.Flags().IntP("id", "I", 0, "ID of the task")
	getTaskByID.Flags().BoolP("logs", "L", false, "Show the task logs if available")
	getProjectKeyCmd.Flags().BoolVarP(&revealValue, "reveal", "", false, "Reveal the variable values")
	getDeploymentByNameCmd.Flags().StringP("name", "N", "", "The name of the deployment (eg, lagoon-build-abcdef)")
	getDeploymentByNameCmd.Flags().BoolP("logs", "L", false, "Show the build logs if available")
}
