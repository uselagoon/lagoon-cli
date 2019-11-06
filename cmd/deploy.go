package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/lagoon/environments"
	"github.com/amazeeio/lagoon-cli/output"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// DeployFlags .
type DeployFlags struct {
	Branch string `json:"branch,omitempty"`
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

var deployEnvCmd = &cobra.Command{
	Use:   "deploy [project name] [branch name]",
	Short: "Deploy a latest branch",
	Long:  "Deploy a latest branch",
	Run: func(cmd *cobra.Command, args []string) {
		validateToken(viper.GetString("current")) // get a new token if the current one is invalid
		deployBranch := parseDeployFlags(*cmd.Flags())
		if cmdProjectName == "" || deployBranch.Branch == "" {
			fmt.Println("Not enough arguments. Requires: lagoon name and branch name")
			cmd.Help()
			os.Exit(1)
		}

		if !outputOptions.JSON {
			fmt.Println(fmt.Sprintf("Deploying %s %s", cmdProjectName, deployBranch.Branch))
		}

		if yesNo() {
			deployResult, err := environments.DeployEnvironmentBranch(cmdProjectName, deployBranch.Branch)
			if err != nil {
				output.RenderError(err.Error(), outputOptions)
				os.Exit(1)
			}
			resultData := output.Result{
				Result: string(deployResult),
			}
			output.RenderResult(resultData, outputOptions)
		}

	},
}

/* @TODO
Need to be able to support more than just deploying the latest branch, like deploying pull requests?
*/
