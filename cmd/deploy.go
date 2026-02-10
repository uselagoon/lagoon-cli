package cmd

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/logrusorgru/aurora"
	"github.com/uselagoon/lagoon-cli/pkg/output"

	lclient "github.com/uselagoon/machinery/api/lagoon/client"

	"github.com/spf13/cobra"
	lagoonssh "github.com/uselagoon/lagoon-cli/pkg/lagoon/ssh"
	"github.com/uselagoon/machinery/api/lagoon"
	"github.com/uselagoon/machinery/api/schema"
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
		branchRef, err := cmd.Flags().GetString("branch-ref")
		if err != nil {
			return err
		}
		returnData, err := cmd.Flags().GetBool("returndata")
		if err != nil {
			return err
		}
		follow, err := cmd.Flags().GetBool("follow")
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
			depBranch := &schema.DeployEnvironmentBranchInput{
				Branch:         branch,
				Project:        cmdProjectName,
				ReturnData:     returnData,
				BuildVariables: buildVarMap,
			}
			if branchRef != "" {
				depBranch.BranchRef = branchRef
			}
			result, err := lagoon.DeployBranch(context.TODO(), depBranch, lc)
			if err != nil {
				return err
			}
			resultData := output.Result{Result: result.DeployEnvironmentBranch}
			r := output.RenderResult(resultData, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)

			if follow {
				return followDeployLogs(cmd, cmdProjectName, branch, resultData.Result, debug)
			}
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
		returnData, err := cmd.Flags().GetBool("returndata")
		if err != nil {
			return err
		}
		follow, err := cmd.Flags().GetBool("follow")
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
			result, err := lagoon.DeployPromote(context.TODO(), &schema.DeployEnvironmentPromoteInput{
				SourceEnvironment:      sourceEnvironment,
				DestinationEnvironment: destinationEnvironment,
				Project:                cmdProjectName,
				BuildVariables:         buildVarMap,
				ReturnData:             returnData,
			}, lc)
			if err != nil {
				return err
			}
			resultData := output.Result{Result: result.DeployEnvironmentPromote}
			r := output.RenderResult(resultData, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)

			if follow {
				return followDeployLogs(cmd, cmdProjectName, destinationEnvironment, resultData.Result, debug)
			}
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

		returnData, err := cmd.Flags().GetBool("returndata")
		if err != nil {
			return err
		}
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		follow, err := cmd.Flags().GetBool("follow")
		if err != nil {
			return err
		}

		buildVarStrings, err := cmd.Flags().GetStringArray("buildvar")
		if err != nil {
			return err
		}
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
			result, err := lagoon.DeployLatest(context.TODO(), &schema.DeployEnvironmentLatestInput{
				Environment: schema.EnvironmentInput{
					Name: cmdProjectEnvironment,
					Project: schema.ProjectInput{
						Name: cmdProjectName,
					},
				},
				ReturnData:     returnData,
				BuildVariables: buildVarMap,
			}, lc)
			if err != nil {
				return err
			}
			resultData := output.Result{Result: result.DeployEnvironmentLatest}
			r := output.RenderResult(resultData, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)

			if follow {
				return followDeployLogs(cmd, cmdProjectName, cmdProjectEnvironment, resultData.Result, debug)
			}
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
		baseBranchName, err := cmd.Flags().GetString("base-branch-name")
		if err != nil {
			return err
		}
		baseBranchRef, err := cmd.Flags().GetString("base-branch-ref")
		if err != nil {
			return err
		}
		headBranchName, err := cmd.Flags().GetString("head-branch-name")
		if err != nil {
			return err
		}
		headBranchRef, err := cmd.Flags().GetString("head-branch-ref")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Project name", cmdProjectName, "Pullrequest title", prTitle, "Pullrequest number", strconv.Itoa(int(prNumber)), "Base branch name", baseBranchName, "Base branch ref", baseBranchRef, "Head branch name", headBranchName, "Head branch ref", headBranchRef); err != nil {
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

		returnData, err := cmd.Flags().GetBool("returndata")
		if err != nil {
			return err
		}
		follow, err := cmd.Flags().GetBool("follow")
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
				ReturnData:     returnData,
				BuildVariables: buildVarMap,
			}, lc)
			if err != nil {
				return err
			}
			resultData := output.Result{Result: result.DeployEnvironmentPullrequest}
			r := output.RenderResult(resultData, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)

			if follow {
				return followDeployLogs(cmd, cmdProjectName, fmt.Sprintf("pr-%d", prNumber), resultData.Result, debug)
			}
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
	deployLatestCmd.Flags().Bool("returndata", false, returnDataUsageText)
	deployLatestCmd.Flags().Bool("follow", false, "Follow the deploy logs")
	deployLatestCmd.Flags().StringArray("buildvar", []string{}, "Add one or more build variables to deployment (--buildvar KEY1=VALUE1 [--buildvar KEY2=VALUE2])")

	deployBranchCmd.Flags().StringP("branch", "b", "", "Branch name to deploy")
	deployBranchCmd.Flags().StringP("branch-ref", "r", "", "Branch ref to deploy")
	deployBranchCmd.Flags().Bool("returndata", false, returnDataUsageText)
	deployBranchCmd.Flags().Bool("follow", false, "Follow the deploy logs")
	deployBranchCmd.Flags().StringArray("buildvar", []string{}, "Add one or more build variables to deployment (--buildvar KEY1=VALUE1 [--buildvar KEY2=VALUE2])")

	deployPromoteCmd.Flags().StringP("destination", "d", "", "Destination environment name to create")
	deployPromoteCmd.Flags().StringP("source", "s", "", "Source environment name to use as the base to deploy from")
	deployPromoteCmd.Flags().Bool("returndata", false, returnDataUsageText)
	deployPromoteCmd.Flags().Bool("follow", false, "Follow the deploy logs")
	deployPromoteCmd.Flags().StringArray("buildvar", []string{}, "Add one or more build variables to deployment (--buildvar KEY1=VALUE1 [--buildvar KEY2=VALUE2])")

	deployPullrequestCmd.Flags().StringP("title", "t", "", "Pullrequest title")
	deployPullrequestCmd.Flags().UintP("number", "n", 0, "Pullrequest number")
	deployPullrequestCmd.Flags().StringP("base-branch-name", "N", "", "Pullrequest base branch name")
	deployPullrequestCmd.Flags().StringP("base-branch-ref", "R", "", "Pullrequest base branch reference hash")
	deployPullrequestCmd.Flags().StringP("head-branch-name", "H", "", "Pullrequest head branch name")
	deployPullrequestCmd.Flags().StringP("head-branch-ref", "M", "", "Pullrequest head branch reference hash")
	deployPullrequestCmd.Flags().Bool("returndata", false, returnDataUsageText)
	deployPullrequestCmd.Flags().Bool("follow", false, "Follow the deploy logs")
	deployPullrequestCmd.Flags().StringArray("buildvar", []string{}, "Add one or more build variables to deployment (--buildvar KEY1=VALUE1 [--buildvar KEY2=VALUE2])")
}

func followDeployLogs(
	cmd *cobra.Command,
	projectName,
	environmentName,
	buildName string,
	debug bool,
) error {
	safeEnvName := makeSafe(shortenEnvironment(projectName, environmentName))
	sshHost, sshPort, username, _, err := getSSHHostPort(safeEnvName, debug)
	if err != nil {
		return fmt.Errorf("couldn't get SSH endpoint: %v", err)
	}
	ignoreHostKey, acceptNewHostKey :=
		lagoonssh.CheckStrictHostKey(strictHostKeyCheck)
	sshConfig, closeSSHAgent, err := getSSHClientConfig(
		username,
		fmt.Sprintf("%s:%s", sshHost, sshPort),
		ignoreHostKey,
		acceptNewHostKey)
	if err != nil {
		return fmt.Errorf("couldn't get SSH client config: %v", err)
	}
	defer func() {
		err = closeSSHAgent()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error closing ssh agent:%v\n", err)
		}
	}()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// start background ticker to close session when deploy completes
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	go func() {
		defer cancel()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				validateToken(lagoonCLIConfig.Current)
				current := lagoonCLIConfig.Current
				token := lagoonCLIConfig.Lagoons[current].Token
				lc := lclient.New(
					lagoonCLIConfig.Lagoons[current].GraphQL,
					lagoonCLIVersion,
					lagoonCLIConfig.Lagoons[current].Version,
					&token,
					debug)
				// ignore errors here since we can't really do anything about them
				deployment, _ := lagoon.GetDeploymentByName(
					ctx, cmdProjectName, cmdProjectEnvironment, buildName, false, lc)
				if deployment.Completed != "" && deployment.Status != "running" {
					var status string
					switch deployment.Status {
					case "complete":
						status = "complete ✅"
					case "failed":
						status = "failed ❌"
					case "cancelled":
						status = "cancelled 🛑"
					default:
						status = deployment.Status
					}
					fmt.Fprintf(
						cmd.OutOrStdout(),
						"Deployment %s finished with status: %s\n",
						aurora.Yellow(buildName),
						status)
					return
				}
			}
		}
	}()
	fmt.Fprintf(cmd.OutOrStdout(), "Streaming deploy logs...\n")
	return lagoonssh.LogStream(ctx, sshConfig, sshHost, sshPort, []string{
		"lagoonSystem=build",
		"logs=tailLines=32,follow",
	})
}
