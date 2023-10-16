package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/uselagoon/lagoon-cli/pkg/api"
	"github.com/uselagoon/lagoon-cli/pkg/output"
	l "github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	"os"
	"strconv"
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

var getOrganizationCmd = &cobra.Command{
	Use:     "organization",
	Aliases: []string{"o"},
	Short:   "Get details about an organization",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		organizationName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		if organizationName == "" {
			return fmt.Errorf("missing arguments: Organization is not defined")
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)
		organization, err := l.GetOrganizationByName(context.TODO(), organizationName, lc)
		handleError(err)

		if organization.Name == "" {
			output.RenderInfo(fmt.Sprintf("No organization found for '%s'", organizationName), outputOptions)
			os.Exit(0)
		}

		data := []output.Data{}
		data = append(data, []string{
			strconv.Itoa(int(organization.ID)),
			organization.Name,
			organization.Description,
			strconv.Itoa(int(organization.QuotaProject)),
			strconv.Itoa(int(organization.QuotaGroup)),
			strconv.Itoa(int(organization.QuotaNotification)),
		})

		dataMain := output.Table{
			Header: []string{"ID", "Name", "Description", "quotaProject", "quotaGroup", "quotaNotification"},
			Data:   data,
		}

		output.RenderOutput(dataMain, outputOptions)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getAllUserKeysCmd)
	getCmd.AddCommand(getDeploymentCmd)
	getCmd.AddCommand(getEnvironmentCmd)
	getCmd.AddCommand(getOrganizationCmd)
	getCmd.AddCommand(getProjectCmd)
	getCmd.AddCommand(getProjectKeyCmd)
	getCmd.AddCommand(getUserKeysCmd)
	getCmd.AddCommand(getTaskByID)
	getCmd.AddCommand(getToken)
	getTaskByID.Flags().IntP("id", "I", 0, "ID of the task")
	getTaskByID.Flags().BoolP("logs", "L", false, "Show the task logs if available")
	getProjectKeyCmd.Flags().BoolVarP(&revealValue, "reveal", "", false, "Reveal the variable values")
	getDeploymentCmd.Flags().StringVarP(&remoteID, "remoteid", "R", "", "The remote ID of the deployment")
	getOrganizationCmd.Flags().StringP("organization", "O", "", "Name of the organization")
}
