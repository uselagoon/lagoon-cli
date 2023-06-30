package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	l "github.com/uselagoon/machinery/api/lagoon"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/uselagoon/lagoon-cli/internal/lagoon"
	"github.com/uselagoon/lagoon-cli/internal/lagoon/client"
	"github.com/uselagoon/lagoon-cli/pkg/output"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	s "github.com/uselagoon/machinery/api/schema"
)

// @TODO re-enable this at some point if more environment based commands are made available

var deployType string
var environmentType string
var deployBaseRef string
var deployHeadRef string
var deployTitle string
var openshiftProjectName string
var route string
var routes string
var autoIdle int
var openshift int
var created string

//type EnvironmentFlags struct {
//	Name string `json:"name,omitempty"`
//}

func parseEnvironmentFlags(flags pflag.FlagSet) s.UpdateEnvironmentPatchInput {
	configMap := make(map[string]interface{})
	flags.VisitAll(func(f *pflag.Flag) {
		if flags.Changed(f.Name) {
			configMap[f.Name] = f.Value
		}
	})
	jsonStr, _ := json.Marshal(configMap)
	parsedFlags := s.UpdateEnvironmentPatchInput{}
	json.Unmarshal(jsonStr, &parsedFlags)
	return parsedFlags
}

var deleteEnvCmd = &cobra.Command{
	Use:     "environment",
	Aliases: []string{"e"},
	Short:   "Delete an environment",
	Run: func(cmd *cobra.Command, args []string) {
		// environmentFlags := parseEnvironmentFlags(*cmd.Flags()) //@TODO re-enable this at some point if more environment based commands are made available
		if cmdProjectName == "" || cmdProjectEnvironment == "" {
			fmt.Println("Missing arguments: Project name or environment name is not defined")
			cmd.Help()
			os.Exit(1)
		}
		if yesNo(fmt.Sprintf("You are attempting to delete environment '%s' from project '%s', are you sure?", cmdProjectEnvironment, cmdProjectName)) {
			projectByName, err := eClient.DeleteEnvironment(cmdProjectName, cmdProjectEnvironment)
			handleError(err)
			resultData := output.Result{
				Result: string(projectByName),
			}
			output.RenderResult(resultData, outputOptions)
		}
	},
}

var updateEnvironmentCmd = &cobra.Command{
	Use:     "environment",
	Aliases: []string{"e"},
	Short:   "Update an environment",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}

		environmentFlags := parseEnvironmentFlags(*cmd.Flags())
		if cmdProjectName == "" || cmdProjectEnvironment == "" {
			fmt.Println("Missing arguments: Project name or environment name is not defined")
			cmd.Help()
			os.Exit(1)
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)
		project, err := l.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
		environment, err := l.GetEnvironmentByName(context.TODO(), cmdProjectEnvironment, project.ID, lc)
		handleError(err)

		result, err := l.UpdateEnvironment(context.TODO(), environment.ID, environmentFlags, lc)
		handleError(err)
		fmt.Println("res", result)
		//var updatedEnvironment s.Environment
		//err = json.Unmarshal([]byte(result), &updatedEnvironment)
		//handleError(err)
		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"Environment Name": result.Name,
			},
		}
		output.RenderResult(resultData, outputOptions)
		return nil
	},
}

var listBackupsCmd = &cobra.Command{
	Use:     "backups",
	Aliases: []string{"b"},
	Short:   "List an environments backups",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		if cmdProjectEnvironment == "" || cmdProjectName == "" {
			return fmt.Errorf("Missing arguments: Project name or environment name is not defined")
		}
		current := lagoonCLIConfig.Current
		lc := client.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIConfig.Lagoons[current].Token,
			lagoonCLIConfig.Lagoons[current].Version,
			lagoonCLIVersion,
			debug)
		project, err := lagoon.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		backupsResult, err := lagoon.GetBackupsForEnvironmentByName(context.TODO(), cmdProjectEnvironment, project.ID, lc)
		if err != nil {
			return err
		}
		data := []output.Data{}
		for _, backup := range backupsResult.Backups {
			alreadyRestored := "false"
			switch backup.Restore.Status {
			case "pending":
			case "failed":
			case "successful":
				alreadyRestored = "true"
			}
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%v", backup.BackupID)),
				returnNonEmptyString(fmt.Sprintf("%v", backup.Source)),
				returnNonEmptyString(fmt.Sprintf("%v", backup.Created)),
				alreadyRestored,
				returnNonEmptyString(fmt.Sprintf("%v", backup.Restore.Status)),
			})
		}
		output.RenderOutput(output.Table{
			Header: []string{
				"BackupID",
				"Source",
				"Created",
				"Restored",
				"RestoreStatus",
			},
			Data: data,
		}, outputOptions)
		return nil
	},
}

var getBackupCmd = &cobra.Command{
	Use:     "backup",
	Aliases: []string{"b"},
	Short:   "Get a backup download link",
	Long: `Get a backup download link
This returns a direct URL to the backup, this is a signed download link with a limited time to initiate the download (usually 5 minutes).`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		backupID, err := cmd.Flags().GetString("backup-id")
		if err != nil {
			return err
		}
		if cmdProjectEnvironment == "" || cmdProjectName == "" {
			return fmt.Errorf("Missing arguments: Project name or environment name is not defined")
		}
		current := lagoonCLIConfig.Current
		lc := client.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIConfig.Lagoons[current].Token,
			lagoonCLIConfig.Lagoons[current].Version,
			lagoonCLIVersion,
			debug)
		project, err := lagoon.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		backupsResult, err := lagoon.GetBackupsForEnvironmentByName(context.TODO(), cmdProjectEnvironment, project.ID, lc)
		if err != nil {
			return err
		}
		status := ""
		for _, backup := range backupsResult.Backups {
			if backup.BackupID == backupID {
				if backup.Restore.RestoreLocation != "" {
					fmt.Println(backup.Restore.RestoreLocation)
					return nil
				}
				status = backup.Restore.Status
			}
		}
		if status != "" {
			return fmt.Errorf("no download file found, status of backups restoration is %s", status)
		}
		return fmt.Errorf("backup has not been restored")
	},
}

//var deployType s.DeployType
//var environmentType s.EnvType

func init() {
	getCmd.AddCommand(getBackupCmd)
	getBackupCmd.Flags().StringP("backup-id", "B", "", "The backup ID you want to restore")
	updateEnvironmentCmd.Flags().StringVarP(&deployBaseRef, "deployBaseRef", "", "", "TODO")
	updateEnvironmentCmd.Flags().StringVarP(&deployHeadRef, "deployHeadRef", "", "", "TODO")
	updateEnvironmentCmd.Flags().StringVarP(&deployTitle, "deployTitle", "", "", "TODO")
	updateEnvironmentCmd.Flags().StringVarP(&openshiftProjectName, "openshiftProjectName", "", "", "TODO")
	updateEnvironmentCmd.Flags().StringVarP(&route, "route", "", "", "TODO")
	updateEnvironmentCmd.Flags().StringVarP(&routes, "routes", "", "", "TODO")
	updateEnvironmentCmd.Flags().IntVarP(&autoIdle, "autoIdle", "a", 0, "TODO")
	updateEnvironmentCmd.Flags().IntVarP(&openshift, "openshift", "", 0, "TODO")
	updateEnvironmentCmd.Flags().StringVarP(&created, "created", "", "", "TODO")
	updateEnvironmentCmd.Flags().StringVarP(&environmentType, "environmentType", "t", "", "TODO")
	updateEnvironmentCmd.Flags().StringVarP(&deployType, "deployType", "d", "", "TODO")
}
