package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/internal/lagoon"
	"github.com/uselagoon/lagoon-cli/internal/lagoon/client"
	"github.com/uselagoon/lagoon-cli/pkg/output"
)

// @TODO re-enable this at some point if more environment based commands are made availab;e
// EnvironmentFlags .
// type EnvironmentFlags struct {
// 	Name string `json:"name,omitempty"`
// }

// func parseEnvironmentFlags(flags pflag.FlagSet) EnvironmentFlags {
// 	configMap := make(map[string]interface{})
// 	flags.VisitAll(func(f *pflag.Flag) {
// 		if flags.Changed(f.Name) {
// 			configMap[f.Name] = f.Value
// 		}
// 	})
// 	jsonStr, _ := json.Marshal(configMap)
// 	parsedFlags := EnvironmentFlags{}
// 	json.Unmarshal(jsonStr, &parsedFlags)
// 	return parsedFlags
// }

var deleteEnvCmd = &cobra.Command{
	Use:     "environment",
	Aliases: []string{"e"},
	Short:   "Delete an environment",
	Run: func(cmd *cobra.Command, args []string) {
		// environmentFlags := parseEnvironmentFlags(*cmd.Flags()) //@TODO re-enable this at some point if more environment based commands are made availab;e
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

var addRestoreCmd = &cobra.Command{
	Use:     "restore",
	Aliases: []string{"r"},
	Hidden:  false,
	Short:   "Restore a backup",
	Long: `Restore a backup
Given a backup-id, you can initiate a restore for it.
You can check the status of the backup using the list backups or get backup command.`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
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
		if backupID == "" {
			return fmt.Errorf("Missing arguments: backup-id is not defined")
		}
		if yesNo(fmt.Sprintf("You are attempting to trigger a restore for backup ID '%s', are you sure?", backupID)) {
			current := lagoonCLIConfig.Current
			lc := client.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIConfig.Lagoons[current].Token,
				lagoonCLIConfig.Lagoons[current].Version,
				lagoonCLIVersion,
				debug)
			result, err := lagoon.AddBackupRestore(context.TODO(), backupID, lc)
			if err != nil {
				return err
			}
			fmt.Println(result)
		}
		return nil
	},
}

func init() {
	addRestoreCmd.Flags().StringP("backup-id", "B", "", "The backup ID you want to restore")
	getCmd.AddCommand(getBackupCmd)
	getBackupCmd.Flags().StringP("backup-id", "B", "", "The backup ID you want to restore")
}
