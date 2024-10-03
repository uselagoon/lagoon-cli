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
		wide, err := cmd.Flags().GetBool("wide")
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
			return handleNilResults("No details for project '%s'\n", cmd, cmdProjectName)
		}

		devEnvironments := 0
		productionRoute := ""
		deploymentsDisabled, err := strconv.ParseBool(strconv.Itoa(int(*project.DeploymentsDisabled)))
		if err != nil {
			return err
		}
		autoIdle, err := strconv.ParseBool(strconv.Itoa(int(*project.AutoIdle)))
		if err != nil {
			return err
		}
		factsUI, err := strconv.ParseBool(strconv.Itoa(int(*project.FactsUI)))
		if err != nil {
			return err
		}
		problemsUI, err := strconv.ParseBool(strconv.Itoa(int(*project.ProblemsUI)))
		if err != nil {
			return err
		}
		for _, environment := range project.Environments {
			if environment.EnvironmentType == "development" {
				devEnvironments++
			}
			if environment.EnvironmentType == "production" {
				productionRoute = environment.Route
			}
		}

		projData := []string{
			returnNonEmptyString(fmt.Sprintf("%d", project.ID)),
			returnNonEmptyString(fmt.Sprintf("%v", project.Name)),
			returnNonEmptyString(fmt.Sprintf("%v", project.GitURL)),
			returnNonEmptyString(fmt.Sprintf("%v", project.ProductionEnvironment)),
			returnNonEmptyString(fmt.Sprintf("%v", productionRoute)),
		}
		if project.DevelopmentEnvironmentsLimit != nil {
			projData = append(projData, returnNonEmptyString(fmt.Sprintf("%v/%v", devEnvironments, *project.DevelopmentEnvironmentsLimit)))
		} else {
			projData = append(projData, returnNonEmptyString(fmt.Sprintf("%v/%v", devEnvironments, 0)))
		}
		projHeader := []string{"ID", "ProjectName", "GitUrl", "ProductionEnvironment", "ProductionRoute", "DevEnvironments"}
		if wide {
			projHeader = append(projHeader, "AutoIdle")
			projData = append(projData, returnNonEmptyString(fmt.Sprintf("%v", autoIdle)))
			projHeader = append(projHeader, "Branches")
			projData = append(projData, returnNonEmptyString(fmt.Sprintf("%v", project.Branches)))
			projHeader = append(projHeader, "PullRequests")
			projData = append(projData, returnNonEmptyString(fmt.Sprintf("%v", project.PullRequests)))
			projHeader = append(projHeader, "RouterPattern")
			projData = append(projData, returnNonEmptyString(fmt.Sprintf("%v", project.RouterPattern)))
			projHeader = append(projHeader, "FactsUI")
			projData = append(projData, returnNonEmptyString(fmt.Sprintf("%v", factsUI)))
			projHeader = append(projHeader, "ProblemsUI")
			projData = append(projData, returnNonEmptyString(fmt.Sprintf("%v", problemsUI)))
			projHeader = append(projHeader, "DeploymentsDisabled")
			projData = append(projData, returnNonEmptyString(fmt.Sprintf("%v", deploymentsDisabled)))
		}
		data := []output.Data{}
		data = append(data, projData)
		dataMain := output.Table{
			Header: projHeader,
			Data:   data,
		}
		r := output.RenderOutput(dataMain, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
			r := output.RenderOutput(dataMain, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
		r := output.RenderOutput(dataMain, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
		wide, err := cmd.Flags().GetBool("wide")
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

		if project.Name == "" || environment.Name == "" {
			if project.Name == "" {
				return handleNilResults("Project '%s' not found\n", cmd, cmdProjectName)
			} else {
				return handleNilResults("Environment '%s' not found in project '%s'\n", cmd, cmdProjectEnvironment, cmdProjectName)
			}
		}

		autoIdle, err := strconv.ParseBool(strconv.Itoa(int(*environment.AutoIdle)))
		if err != nil {
			return err
		}

		data := []output.Data{}
		var envRoute = "none"
		if environment.Route != "" {
			envRoute = environment.Route
		}
		envData := []string{
			returnNonEmptyString(fmt.Sprintf("%d", environment.ID)),
			returnNonEmptyString(fmt.Sprintf("%v", environment.Name)),
			returnNonEmptyString(fmt.Sprintf("%v", environment.DeployType)),
			returnNonEmptyString(fmt.Sprintf("%v", environment.EnvironmentType)),
			returnNonEmptyString(fmt.Sprintf("%v", environment.OpenshiftProjectName)),
			returnNonEmptyString(fmt.Sprintf("%v", envRoute)),
			returnNonEmptyString(fmt.Sprintf("%v", environment.DeployTarget.Name)),
		}
		envHeader := []string{"ID", "Name", "DeployType", "EnvironmentType", "Namespace", "Route", "DeployTarget"}
		// if wide, add additional fields to the result
		if wide {
			envHeader = append(envHeader, "Created")
			envData = append(envData, returnNonEmptyString(fmt.Sprintf("%v", environment.Created)))
			envHeader = append(envHeader, "AutoIdle")
			envData = append(envData, returnNonEmptyString(fmt.Sprintf("%v", autoIdle)))
			envHeader = append(envHeader, "DeployTitle")
			envData = append(envData, returnNonEmptyString(fmt.Sprintf("%v", environment.DeployTitle)))
			envHeader = append(envHeader, "DeployBaseRef")
			envData = append(envData, returnNonEmptyString(fmt.Sprintf("%v", environment.DeployBaseRef)))
			envHeader = append(envHeader, "DeployHeadRef")
			envData = append(envData, returnNonEmptyString(fmt.Sprintf("%v", environment.DeployHeadRef)))
			envHeader = append(envHeader, "Routes")
			envData = append(envData, returnNonEmptyString(fmt.Sprintf("%v", environment.Routes)))
		}
		data = append(data, envData)
		dataMain := output.Table{
			Header: envHeader,
			Data:   data,
		}
		r := output.RenderOutput(dataMain, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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

		project, err := lagoon.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		if project.Name == "" {
			return handleNilResults("No project found for '%s'\n", cmd, cmdProjectName)
		}

		projectKey, err := lagoon.GetProjectKeyByName(context.TODO(), cmdProjectName, revealValue, lc)
		if err != nil {
			return err
		}
		if projectKey.PublicKey == "" {
			return handleNilResults("No project-key for project '%s'\n", cmd, cmdProjectName)
		}
		if revealValue && projectKey.PrivateKey == "" {
			return handleNilResults("No private-key for project '%s'\n", cmd, cmdProjectName)
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

		if projectKey.PrivateKey != "" {
			dataMain.Header = append(dataMain.Header, "PrivateKey")
		}
		r := output.RenderOutput(dataMain, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
			strconv.Itoa(organization.QuotaProject),
			strconv.Itoa(organization.QuotaGroup),
			strconv.Itoa(organization.QuotaNotification),
		})

		dataMain := output.Table{
			Header: []string{"ID", "Name", "Description", "Project Quota", "Group Quota", "Notification Quota"},
			Data:   data,
		}

		r := output.RenderOutput(dataMain, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
	getEnvironmentCmd.Flags().Bool("wide", false, "Display additional information about the environment")
	getProjectCmd.Flags().Bool("wide", false, "Display additional information about the project")
}
