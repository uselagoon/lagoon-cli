package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/uselagoon/lagoon-cli/pkg/api"
	"github.com/uselagoon/lagoon-cli/pkg/output"
	"os"
	"strconv"

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
			return nil
		}
		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)

		project, err := l.GetProjectByName(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}

		if project == nil {
			output.RenderInfo(fmt.Sprintf("No details for project '%s'", cmdProjectName), outputOptions)
			return nil
		}

		DevEnvironments := 0
		productionRoute := "none"
		deploymentsDisabled, err := strconv.ParseBool(strconv.Itoa(int(project.DeploymentsDisabled)))
		handleError(err)
		autoIdle, err := strconv.ParseBool(strconv.Itoa(int(project.AutoIdle)))
		handleError(err)
		factsUI, err := strconv.ParseBool(strconv.Itoa(int(project.FactsUI)))
		handleError(err)
		problemsUI, err := strconv.ParseBool(strconv.Itoa(int(project.ProblemsUI)))
		handleError(err)
		for _, environment := range project.Environments {
			if environment.EnvironmentType == "development" {
				DevEnvironments++
			}
			if environment.EnvironmentType == "production" {
				productionRoute = environment.Route
			}
		}

		data := []output.Data{}
		data = append(data, []string{
			returnNonEmptyString(fmt.Sprintf("%d", project.ID)),
			returnNonEmptyString(fmt.Sprintf("%v", project.Name)),
			returnNonEmptyString(fmt.Sprintf("%v", project.GitURL)),
			returnNonEmptyString(fmt.Sprintf("%v", project.Branches)),
			returnNonEmptyString(fmt.Sprintf("%v", project.PullRequests)),
			returnNonEmptyString(fmt.Sprintf("%v", productionRoute)),
			returnNonEmptyString(fmt.Sprintf("%v/%v", DevEnvironments, project.DevelopmentEnvironmentsLimit)),
			returnNonEmptyString(fmt.Sprintf("%v", project.DevelopmentEnvironmentsLimit)),
			returnNonEmptyString(fmt.Sprintf("%v", project.ProductionEnvironment)),
			returnNonEmptyString(fmt.Sprintf("%v", project.RouterPattern)),
			returnNonEmptyString(fmt.Sprintf("%v", autoIdle)),
			returnNonEmptyString(fmt.Sprintf("%v", factsUI)),
			returnNonEmptyString(fmt.Sprintf("%v", problemsUI)),
			returnNonEmptyString(fmt.Sprintf("%v", deploymentsDisabled)),
		})
		dataMain := output.Table{
			Header: []string{"ID", "ProjectName", "GitURL", "Branches", "PullRequests", "ProductionRoute", "DevEnvironments", "DevEnvLimit", "ProductionEnv", "RouterPattern", "AutoIdle", "FactsUI", "ProblemsUI", "DeploymentsDisabled"},
			Data:   data,
		}
		output.RenderOutput(dataMain, outputOptions)
		return nil
	},
}

var getDeploymentCmd = &cobra.Command{
	Use:     "deployment",
	Aliases: []string{"d"},
	Short:   "Get a build log by remote id",
	Run: func(cmd *cobra.Command, args []string) {
		getProjectFlags := parseGetFlags(*cmd.Flags())
		if getProjectFlags.RemoteID == "" {
			fmt.Println("Missing arguments: Remote ID is not defined")
			cmd.Help()
			os.Exit(1)
		}
		returnedJSON, err := eClient.GetDeploymentLog(getProjectFlags.RemoteID)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		if string(returnedJSON) == "null" {
			output.RenderInfo(fmt.Sprintf("No deployment for remoteId '%s'", getProjectFlags.RemoteID), outputOptions)
			os.Exit(0)
		}
		var deployment api.Deployment
		err = json.Unmarshal([]byte(returnedJSON), &deployment)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		if deployment.BuildLog != "" {
			fmt.Println(deployment.BuildLog)
		} else {
			fmt.Println("Log data is not available")
		}

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
	getCmd.AddCommand(getDeploymentCmd)
	getCmd.AddCommand(getEnvironmentCmd)
	getCmd.AddCommand(getProjectCmd)
	getCmd.AddCommand(getProjectKeyCmd)
	getCmd.AddCommand(getUserKeysCmd)
	getCmd.AddCommand(getTaskByID)
	getCmd.AddCommand(getToken)
	getTaskByID.Flags().IntP("id", "I", 0, "ID of the task")
	getTaskByID.Flags().BoolP("logs", "L", false, "Show the task logs if available")
	getProjectKeyCmd.Flags().BoolVarP(&revealValue, "reveal", "", false, "Reveal the variable values")
	getDeploymentCmd.Flags().StringVarP(&remoteID, "remoteid", "R", "", "The remote ID of the deployment")
}
