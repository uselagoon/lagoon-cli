package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/lagoon/environments"
	"github.com/amazeeio/lagoon-cli/lagoon/projects"
	"github.com/amazeeio/lagoon-cli/output"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
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
		validateToken(viper.GetString("current")) // get a new token if the current one is invalid
	},
}

var listProjectsCmd = &cobra.Command{
	Use:     "projects",
	Aliases: []string{"p"},
	Short:   "Show your projects (alias: p)",
	Run: func(cmd *cobra.Command, args []string) {
		returnedJSON, err := projects.ListAllProjects()
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
			output.RenderError("no data returned", outputOptions)
			os.Exit(1)
		}
		output.RenderOutput(dataMain, outputOptions)

	},
}

var listProjectCmd = &cobra.Command{
	Use:     "project-environments",
	Aliases: []string{"pe"},
	Short:   "List environments for a project (alias: pe)",
	Run: func(cmd *cobra.Command, args []string) {
		getProjectFlags := parseListFlags(*cmd.Flags())
		if getProjectFlags.Project == "" {
			fmt.Println("Not enough arguments. Requires: project name")
			cmd.Help()
			os.Exit(1)
		}

		returnedJSON, err := projects.ListEnvironmentsForProject(getProjectFlags.Project)
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
			output.RenderError("no data returned", outputOptions)
			os.Exit(1)
		}
		output.RenderOutput(dataMain, outputOptions)

	},
}

var listVariablesCmd = &cobra.Command{
	Use:     "variables",
	Aliases: []string{"v"},
	Short:   "Show your variables for a project or environment (alias: v)",
	Run: func(cmd *cobra.Command, args []string) {
		getListFlags := parseListFlags(*cmd.Flags())
		if getListFlags.Project == "" {
			fmt.Println("Not enough arguments. Requires: project name")
			cmd.Help()
			os.Exit(1)
		}
		var returnedJSON []byte
		var err error
		if getListFlags.Environment != "" {
			returnedJSON, err = environments.ListEnvironmentVariables(getListFlags.Project, getListFlags.Environment, getListFlags.Reveal)
		} else {
			returnedJSON, err = projects.ListProjectVariables(getListFlags.Project, getListFlags.Reveal)
		}
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
			output.RenderError("no data returned", outputOptions)
			os.Exit(1)
		}
		output.RenderOutput(dataMain, outputOptions)
	},
}

var listDeploymentsCmd = &cobra.Command{
	Use:     "deployments",
	Aliases: []string{"d"},
	Short:   "Show your deployments for an environment (alias: d)",
	Run: func(cmd *cobra.Command, args []string) {
		getListFlags := parseListFlags(*cmd.Flags())
		if getListFlags.Project == "" || getListFlags.Environment == "" {
			fmt.Println("Not enough arguments. Requires: project name and environment name")
			cmd.Help()
			os.Exit(1)
		}

		returnedJSON, err := environments.GetEnvironmentDeployments(getListFlags.Project, getListFlags.Environment)
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
			output.RenderError("no data returned", outputOptions)
			os.Exit(1)
		}
		output.RenderOutput(dataMain, outputOptions)
	},
}

func init() {
	listCmd.AddCommand(listProjectCmd)
	listCmd.AddCommand(listProjectsCmd)
	listCmd.AddCommand(listVariablesCmd)
	listCmd.AddCommand(listDeploymentsCmd)
	listCmd.AddCommand(listRocketChatsCmd)
	listCmd.AddCommand(listSlackCmd)
	listVariablesCmd.Flags().BoolVarP(&revealValue, "reveal", "", false, "Reveal the variable values")
}
