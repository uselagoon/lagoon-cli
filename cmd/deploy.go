package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/internal/lagoon"
	"github.com/amazeeio/lagoon-cli/internal/lagoon/client"
	"github.com/amazeeio/lagoon-cli/internal/schema"
	"github.com/amazeeio/lagoon-cli/pkg/output"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// DeployFlags .
type DeployFlags struct {
	Branch      string `json:"branch,omitempty"`
	Source      string `json:"source,omitempty"`
	Destination string `json:"destination,omitempty"`
}

func parseDeployFlags(flags pflag.FlagSet) DeployFlags {
	configMap := make(map[string]interface{})
	flags.VisitAll(func(f *pflag.Flag) {
		if flags.Changed(f.Name) {
			configMap[f.Name] = f.Value
		}
	})
	jsonStr, _ := json.Marshal(configMap)
	parsedFlags := DeployFlags{}
	json.Unmarshal(jsonStr, &parsedFlags)
	return parsedFlags
}

var deployCmd = &cobra.Command{
	Use:     "deploy",
	Aliases: []string{"d"},
	Short:   "Actions for deploying or promoting branches or environments in lagoon",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateToken(viper.GetString("current")) // get a new token if the current one is invalid
	},
}

var deployBranchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Deploy a new branch",
	Long: `Deploy a new branch
This branch may or may not already exist in lagoon, if it already exists you may want to
use 'lagoon deploy latest' instead`,
	Aliases: []string{"b"},
	Run: func(cmd *cobra.Command, args []string) {
		validateToken(viper.GetString("current")) // get a new token if the current one is invalid
		deployBranch := parseDeployFlags(*cmd.Flags())
		if cmdProjectName == "" || deployBranch.Branch == "" {
			fmt.Println("Missing arguments: Project name or branch name is not defined")
			cmd.Help()
			os.Exit(1)
		}
		if yesNo(fmt.Sprintf("You are attempting to deploy branch '%s' for project '%s', are you sure?", deployBranch.Branch, cmdProjectName)) {
			deployResult, err := eClient.DeployEnvironmentBranch(cmdProjectName, deployBranch.Branch)
			handleError(err)
			resultData := output.Result{
				Result: string(deployResult),
			}
			output.RenderResult(resultData, outputOptions)
		}

	},
}

var deployPromoteCmd = &cobra.Command{
	Use:     "promote",
	Aliases: []string{"p"},
	Short:   "Promote an environment",
	Long:    "Promote one environment to another",
	Run: func(cmd *cobra.Command, args []string) {
		validateToken(viper.GetString("current")) // get a new token if the current one is invalid
		promoteEnv := parseDeployFlags(*cmd.Flags())
		if cmdProjectName == "" || promoteEnv.Source == "" || promoteEnv.Destination == "" {
			fmt.Println("Missing arguments: Project name, source environment name, or destination environment name is not defined")
			cmd.Help()
			os.Exit(1)
		}
		if yesNo(fmt.Sprintf("You are attempting to promote environment '%s' to '%s' for project '%s', are you sure?", promoteEnv.Source, promoteEnv.Destination, cmdProjectName)) {
			deployResult, err := eClient.PromoteEnvironment(cmdProjectName, promoteEnv.Source, promoteEnv.Destination)
			handleError(err)
			resultData := output.Result{
				Result: string(deployResult),
			}
			output.RenderResult(resultData, outputOptions)
		}

	},
}

var deployLatestCmd = &cobra.Command{
	Use:     "latest",
	Aliases: []string{"l"},
	Hidden:  false,
	Short:   "Deploy latest environment",
	Long: `Deploy latest environment
This environment should already exist in lagoon. It is analogous with the 'Deploy' button in the Lagoon UI`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(viper.GetString("current"))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		if cmdProjectName == "" || cmdProjectEnvironment == "" {
			return fmt.Errorf("Missing arguments: Project name or environment name is not defined")
		}
		current := viper.GetString("current")
		lc := client.New(
			viper.GetString("lagoons."+current+".graphql"),
			viper.GetString("lagoons."+current+".token"),
			viper.GetString("lagoons."+current+".version"),
			lagoonCLIVersion,
			debug)
		result, err := lagoon.DeployLatest(context.TODO(), &schema.DeployEnvironmentLatestInput{
			Environment: schema.EnvironmentInput{
				Name: cmdProjectEnvironment,
				Project: schema.ProjectInput{
					Name: cmdProjectName,
				},
			},
		}, lc)
		if err != nil {
			return err
		}
		fmt.Println(result.DeployEnvironmentLatest)
		return nil
	},
}

var deployPullrequestCmd = &cobra.Command{
	Use:     "pullrequest",
	Aliases: []string{"r"},
	Hidden:  false,
	Short:   "Deploy a pullrequest",
	Long: `Deploy a pullrequest
This pullrequest may not already exist as an environment in lagoon.`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(viper.GetString("current"))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		prTitle, err := cmd.Flags().GetString("title")
		if err != nil {
			return err
		}
		prNumber, err := cmd.Flags().GetUint("number")
		if err != nil {
			return err
		}
		baseBranchName, err := cmd.Flags().GetString("baseBranchName")
		if err != nil {
			return err
		}
		baseBranchRef, err := cmd.Flags().GetString("baseBranchRef")
		if err != nil {
			return err
		}
		headBranchName, err := cmd.Flags().GetString("headBranchName")
		if err != nil {
			return err
		}
		headBranchRef, err := cmd.Flags().GetString("headBranchRef")
		if err != nil {
			return err
		}
		if cmdProjectName == "" || prTitle == "" || prNumber == 0 || baseBranchName == "" ||
			baseBranchRef == "" || headBranchName == "" || headBranchRef == "" {
			return fmt.Errorf("Missing arguments: Project name, title, number, baseBranchName, baseBranchRef, headBranchName, or headBranchRef is not defined")
		}
		current := viper.GetString("current")
		lc := client.New(
			viper.GetString("lagoons."+current+".graphql"),
			viper.GetString("lagoons."+current+".token"),
			viper.GetString("lagoons."+current+".version"),
			lagoonCLIVersion,
			debug)

		result, err := lagoon.DeployPullRequest(context.TODO(), &schema.DeployEnvironmentPullrequestInput{
			Project: schema.ProjectInput{
				Name: cmdProjectName,
			},
			Title:          prTitle,
			Number:         prNumber,
			BaseBranchName: baseBranchName,
			BaseBranchRef:  baseBranchRef,
			HeadBranchName: headBranchName,
			HeadBranchRef:  headBranchRef,
		}, lc)
		if err != nil {
			return err
		}
		fmt.Println(result.DeployEnvironmentPullrequest)
		return nil
	},
}

var (
	promoteSourceEnv string
	promoteDestEnv   string
)

func init() {
	deployCmd.AddCommand(deployBranchCmd)
	deployCmd.AddCommand(deployPromoteCmd)
	deployCmd.AddCommand(deployLatestCmd)
	deployCmd.AddCommand(deployPullrequestCmd)
	deployBranchCmd.Flags().StringVarP(&deployBranchName, "branch", "b", "", "branch name")
	deployPromoteCmd.Flags().StringVarP(&promoteDestEnv, "destination", "d", "", "destination environment name")
	deployPromoteCmd.Flags().StringVarP(&promoteSourceEnv, "source", "s", "", "source environment name")

	deployPullrequestCmd.Flags().StringP("title", "t", "", "Pullrequest title")
	deployPullrequestCmd.Flags().UintP("number", "n", 0, "Pullrequest number")
	deployPullrequestCmd.Flags().StringP("baseBranchName", "N", "", "Pullrequest base branch name")
	deployPullrequestCmd.Flags().StringP("baseBranchRef", "R", "", "Pullrequest base branch reference hash")
	deployPullrequestCmd.Flags().StringP("headBranchName", "H", "", "Pullrequest head branch name")
	deployPullrequestCmd.Flags().StringP("headBranchRef", "M", "", "Pullrequest head branch reference hash")
}
