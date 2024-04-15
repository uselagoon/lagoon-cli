package cmd

import (
	"context"
	"fmt"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	"strconv"

	"github.com/spf13/cobra"
	l "github.com/uselagoon/machinery/api/lagoon"
	ls "github.com/uselagoon/machinery/api/schema"
)

var deployCmd = &cobra.Command{
	Use:     "deploy",
	Aliases: []string{"d", "dep"},
	Short:   "Actions for deploying or promoting branches or environments in lagoon",
}

var deployBranchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Deploy a new branch",
	Long: `Deploy a new branch
This branch may or may not already exist in lagoon, if it already exists you may want to
use 'lagoon deploy latest' instead`,
	Aliases: []string{"b"},
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		branch, err := cmd.Flags().GetString("branch")
		if err != nil {
			return err
		}
		branchRef, err := cmd.Flags().GetString("branchRef")
		if err != nil {
			return err
		}
		returnData, err := cmd.Flags().GetBool("returnData")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Project name", cmdProjectName, "Branch name", branch); err != nil {
			return err
		}

		buildVarStrings, err := cmd.Flags().GetStringArray("buildvar")
		if err != nil {
			return err
		}
		buildVarMap, err := buildVarsToMap(buildVarStrings)
		if err != nil {
			return err
		}

		if yesNo(fmt.Sprintf("You are attempting to deploy branch '%s' for project '%s', are you sure?", branch, cmdProjectName)) {
			current := lagoonCLIConfig.Current
			token := lagoonCLIConfig.Lagoons[current].Token
			lc := lclient.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIVersion,
				lagoonCLIConfig.Lagoons[current].Version,
				&token,
				debug)
			depBranch := &ls.DeployEnvironmentBranchInput{
				Branch:         branch,
				Project:        cmdProjectName,
				ReturnData:     returnData,
				BuildVariables: buildVarMap,
			}
			if branchRef != "" {
				depBranch.BranchRef = branchRef
			}
			result, err := l.DeployBranch(context.TODO(), depBranch, lc)
			if err != nil {
				return err
			}
			fmt.Println(result.DeployEnvironmentBranch)
		}
		return nil
	},
}

var deployPromoteCmd = &cobra.Command{
	Use:     "promote",
	Aliases: []string{"p"},
	Short:   "Promote an environment",
	Long:    "Promote one environment to another",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		sourceEnvironment, err := cmd.Flags().GetString("source")
		if err != nil {
			return err
		}
		destinationEnvironment, err := cmd.Flags().GetString("destination")
		if err != nil {
			return err
		}
		returnData, err := cmd.Flags().GetBool("returnData")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Project name", cmdProjectName, "Source environment", sourceEnvironment, "Destination environment", destinationEnvironment); err != nil {
			return err
		}

		buildVarStrings, err := cmd.Flags().GetStringArray("buildvar")
		if err != nil {
			return err
		}
		buildVarMap, err := buildVarsToMap(buildVarStrings)
		if err != nil {
			return err
		}

		if yesNo(fmt.Sprintf("You are attempting to promote environment '%s' to '%s' for project '%s', are you sure?", sourceEnvironment, destinationEnvironment, cmdProjectName)) {
			current := lagoonCLIConfig.Current
			token := lagoonCLIConfig.Lagoons[current].Token
			lc := lclient.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIVersion,
				lagoonCLIConfig.Lagoons[current].Version,
				&token,
				debug)
			result, err := l.DeployPromote(context.TODO(), &ls.DeployEnvironmentPromoteInput{
				SourceEnvironment:      sourceEnvironment,
				DestinationEnvironment: destinationEnvironment,
				Project:                cmdProjectName,
				BuildVariables:         buildVarMap,
				ReturnData:             returnData,
			}, lc)
			if err != nil {
				return err
			}
			fmt.Println(result.DeployEnvironmentPromote)
		}
		return nil
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
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		returnData, err := cmd.Flags().GetBool("returnData")
		if err != nil {
			return err
		}
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}

		buildVarStrings, err := cmd.Flags().GetStringArray("buildvar")
		if err := requiredInputCheck("Project name", cmdProjectName, "Environment name", cmdProjectEnvironment); err != nil {
			return err
		}
		buildVarMap, err := buildVarsToMap(buildVarStrings)
		if err != nil {
			return err
		}

		if yesNo(fmt.Sprintf("You are attempting to deploy the latest environment '%s' for project '%s', are you sure?", cmdProjectEnvironment, cmdProjectName)) {
			current := lagoonCLIConfig.Current
			token := lagoonCLIConfig.Lagoons[current].Token
			lc := lclient.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIVersion,
				lagoonCLIConfig.Lagoons[current].Version,
				&token,
				debug)
			result, err := l.DeployLatest(context.TODO(), &ls.DeployEnvironmentLatestInput{
				Environment: ls.EnvironmentInput{
					Name: cmdProjectEnvironment,
					Project: ls.ProjectInput{
						Name: cmdProjectName,
					},
				},
				ReturnData:     returnData,
				BuildVariables: buildVarMap,
			}, lc)
			if err != nil {
				return err
			}
			fmt.Println(result.DeployEnvironmentLatest)
		}
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
		return validateTokenE(lagoonCLIConfig.Current)
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
		if err := requiredInputCheck("Project name", cmdProjectName, "Pullrequest title", prTitle, "Pullrequest number", strconv.Itoa(int(prNumber)), "baseBranchName", baseBranchName, "baseBranchRef", baseBranchRef, "headBranchName", headBranchName, "headBranchRef", headBranchRef); err != nil {
			return err
		}
		buildVarStrings, err := cmd.Flags().GetStringArray("buildvar")
		if err != nil {
			return err
		}
		buildVarMap, err := buildVarsToMap(buildVarStrings)
		if err != nil {
			return err
		}

		returnData, err := cmd.Flags().GetBool("returnData")
		if err != nil {
			return err
		}
		if yesNo(fmt.Sprintf("You are attempting to deploy pull request '%v' for project '%s', are you sure?", prNumber, cmdProjectName)) {
			current := lagoonCLIConfig.Current
			token := lagoonCLIConfig.Lagoons[current].Token
			lc := lclient.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIVersion,
				lagoonCLIConfig.Lagoons[current].Version,
				&token,
				debug)

			result, err := l.DeployPullRequest(context.TODO(), &ls.DeployEnvironmentPullrequestInput{
				Project: ls.ProjectInput{
					Name: cmdProjectName,
				},
				Title:          prTitle,
				Number:         prNumber,
				BaseBranchName: baseBranchName,
				BaseBranchRef:  baseBranchRef,
				HeadBranchName: headBranchName,
				HeadBranchRef:  headBranchRef,
				ReturnData:     returnData,
				BuildVariables: buildVarMap,
			}, lc)
			if err != nil {
				return err
			}
			fmt.Println(result.DeployEnvironmentPullrequest)
		}
		return nil
	},
}

func init() {
	deployCmd.AddCommand(deployBranchCmd)
	deployCmd.AddCommand(deployPromoteCmd)
	deployCmd.AddCommand(deployLatestCmd)
	deployCmd.AddCommand(deployPullrequestCmd)

	const returnDataUsageText = "Returns the build name instead of success text"
	deployLatestCmd.Flags().Bool("returnData", false, returnDataUsageText)
	deployLatestCmd.Flags().StringArray("buildvar", []string{}, "Add one or more build variables to deployment (--buildvar KEY1=VALUE1 [--buildvar KEY2=VALUE2])")

	deployBranchCmd.Flags().StringP("branch", "b", "", "Branch name to deploy")
	deployBranchCmd.Flags().StringP("branchRef", "r", "", "Branch ref to deploy")
	deployBranchCmd.Flags().Bool("returnData", false, returnDataUsageText)
	deployBranchCmd.Flags().StringArray("buildvar", []string{}, "Add one or more build variables to deployment (--buildvar KEY1=VALUE1 [--buildvar KEY2=VALUE2])")

	deployPromoteCmd.Flags().StringP("destination", "d", "", "Destination environment name to create")
	deployPromoteCmd.Flags().StringP("source", "s", "", "Source environment name to use as the base to deploy from")
	deployPromoteCmd.Flags().Bool("returnData", false, returnDataUsageText)
	deployPromoteCmd.Flags().StringArray("buildvar", []string{}, "Add one or more build variables to deployment (--buildvar KEY1=VALUE1 [--buildvar KEY2=VALUE2])")

	deployPullrequestCmd.Flags().StringP("title", "t", "", "Pullrequest title")
	deployPullrequestCmd.Flags().UintP("number", "n", 0, "Pullrequest number")
	deployPullrequestCmd.Flags().StringP("baseBranchName", "N", "", "Pullrequest base branch name")
	deployPullrequestCmd.Flags().StringP("baseBranchRef", "R", "", "Pullrequest base branch reference hash")
	deployPullrequestCmd.Flags().StringP("headBranchName", "H", "", "Pullrequest head branch name")
	deployPullrequestCmd.Flags().StringP("headBranchRef", "M", "", "Pullrequest head branch reference hash")
	deployPullrequestCmd.Flags().Bool("returnData", false, returnDataUsageText)
	deployPullrequestCmd.Flags().StringArray("buildvar", []string{}, "Add one or more build variables to deployment (--buildvar KEY1=VALUE1 [--buildvar KEY2=VALUE2])")
}
