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
)

// GetFlags .
type GetFlags struct {
	Project     string `json:"project,omitempty"`
	Environment string `json:"environment,omitempty"`
	RemoteID    string `json:"remoteid,omitempty"`
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

		project, err := lagoon.GetProjectByName(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}

		if project.Name == "" {
			outputOptions.Error = fmt.Sprintf("No details for project '%s'\n", cmdProjectName)
			output.RenderOutput(output.Table{Data: []output.Data{[]string{}}}, outputOptions)
			return nil
		}

		DevEnvironments := 0
		productionRoute := "none"
		deploymentsDisabled, err := strconv.ParseBool(strconv.Itoa(int(project.DeploymentsDisabled)))
		if err != nil {
			return err
		}
		autoIdle, err := strconv.ParseBool(strconv.Itoa(int(project.AutoIdle)))
		if err != nil {
			return err
		}
		factsUI, err := strconv.ParseBool(strconv.Itoa(int(project.FactsUI)))
		if err != nil {
			return err
		}
		problemsUI, err := strconv.ParseBool(strconv.Itoa(int(project.ProblemsUI)))
		if err != nil {
			return err
		}
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
		if err := requiredInputCheck("Project name", cmdProjectName, "Environment name", cmdProjectEnvironment, "Build name", buildName); err != nil {
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
		deployment, err := lagoon.GetDeploymentByName(context.TODO(), cmdProjectName, cmdProjectEnvironment, buildName, showLogs, lc)
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

		data := []output.Data{}
		data = append(data, []string{
			returnNonEmptyString(fmt.Sprintf("%d", environment.ID)),
			returnNonEmptyString(fmt.Sprintf("%v", environment.Name)),
			returnNonEmptyString(fmt.Sprintf("%v", environment.EnvironmentType)),
			returnNonEmptyString(fmt.Sprintf("%v", environment.DeployType)),
			returnNonEmptyString(fmt.Sprintf("%v", environment.Created)),
			returnNonEmptyString(fmt.Sprintf("%v", environment.OpenshiftProjectName)),
			returnNonEmptyString(fmt.Sprintf("%v", environment.Route)),
			returnNonEmptyString(fmt.Sprintf("%v", environment.Routes)),
			returnNonEmptyString(fmt.Sprintf("%d", environment.AutoIdle)),
			returnNonEmptyString(fmt.Sprintf("%v", environment.DeployTitle)),
			returnNonEmptyString(fmt.Sprintf("%v", environment.DeployBaseRef)),
			returnNonEmptyString(fmt.Sprintf("%v", environment.DeployHeadRef)),
		})
		dataMain := output.Table{
			Header: []string{"ID", "EnvironmentName", "EnvironmentType", "DeployType", "Created", "Namespace", "Route", "Routes", "AutoIdle", "DeployTitle", "DeployBaseRef", "DeployHeadRef"},
			Data:   data,
		}
		output.RenderOutput(dataMain, outputOptions)
		return nil
	},
}

var getProjectKeyCmd = &cobra.Command{
	Use:     "project-key",
	Aliases: []string{"pk"},
	Short:   "Get a projects public key",
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

		projectKey, err := lagoon.GetProjectKeyByName(context.TODO(), cmdProjectName, revealValue, lc)
		if err != nil {
			return err
		}
		projectKeys := []string{projectKey.PublicKey}
		if projectKey.PrivateKey != "" {
			projectKeys = append(projectKeys, strings.TrimSuffix(projectKey.PrivateKey, "\n"))
			outputOptions.MultiLine = true
		}

		var data []output.Data
		data = append(data, projectKeys)

		dataMain := output.Table{
			Header: []string{"PublicKey"},
			Data:   data,
		}

		if len(dataMain.Data) == 0 {
			outputOptions.Error = fmt.Sprintf("No project-key for project '%s'", cmdProjectName)
			output.RenderOutput(output.Table{Data: []output.Data{[]string{}}}, outputOptions)
			return nil
		}

		if projectKey.PrivateKey != "" {
			dataMain.Header = append(dataMain.Header, "PrivateKey")
		}
		output.RenderOutput(dataMain, outputOptions)
		return nil
	},
}

var getToken = &cobra.Command{
	Use:     "token",
	Aliases: []string{"tk"},
	Short:   "Generates a Lagoon auth token (for use in, for example, graphQL queries)",
	RunE: func(cmd *cobra.Command, args []string) error {
		token, err := retrieveTokenViaSsh()
		if err != nil {
			return err
		}
		fmt.Println(token)
		return nil
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
			Header: []string{"ID", "Name", "Description", "Project Quota", "Group Quota", "Notification Quota"},
			Data:   data,
		}

		output.RenderOutput(dataMain, outputOptions)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getAllUserKeysCmd)
	getCmd.AddCommand(getEnvironmentCmd)
	getCmd.AddCommand(getOrganizationCmd)
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
	getOrganizationCmd.Flags().StringP("organization-name", "O", "", "Name of the organization")
}
