package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	l "github.com/uselagoon/machinery/api/lagoon"
	s "github.com/uselagoon/machinery/api/schema"
	"os"

	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/internal/lagoon"
	"github.com/uselagoon/lagoon-cli/internal/lagoon/client"
	"github.com/uselagoon/lagoon-cli/pkg/output"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
)

// @TODO re-enable this at some point if more environment based commands are made available

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
		//name, err := cmd.Flags().GetString("name")
		//if err != nil {
		//	return err
		//}
		deployBaseRef, err := cmd.Flags().GetString("deploy-base-ref")
		if err != nil {
			return err
		}
		deployHeadRef, err := cmd.Flags().GetString("deploy-head-ref")
		if err != nil {
			return err
		}
		namespace, err := cmd.Flags().GetString("namespace")
		if err != nil {
			return err
		}
		route, err := cmd.Flags().GetString("route")
		if err != nil {
			return err
		}
		routes, err := cmd.Flags().GetString("routes")
		if err != nil {
			return err
		}
		//environmentType, err := cmd.Flags().GetString("environment-type")
		//if err != nil {
		//	return err
		//}
		//deployT, err := cmd.Flags().GetString("deploy-type")
		//if err != nil {
		//	return err
		//}
		autoIdle, err := cmd.Flags().GetUint("auto-idle")
		if err != nil {
			return err
		}
		openShift, err := cmd.Flags().GetUint("deploy-target")
		if err != nil {
			return err
		}
		deployTitle, err := cmd.Flags().GetString("deploy-title")
		if err != nil {
			return err
		}

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

		//envType := s.EnvType(environmentType)
		//deployType := s.DeployType(deployT)
		environmentFlags := s.UpdateEnvironmentPatchInput{
			//Name:                 &name,
			ProjectID:            &project.ID,
			DeployBaseRef:        &deployBaseRef,
			DeployHeadRef:        &deployHeadRef,
			OpenshiftProjectName: &namespace,
			Route:                &route,
			Routes:               &routes,
			//EnvironmentType:      &envType,
			//DeployType:           &deployType,
			AutoIdle:    &autoIdle,
			DeployTitle: &deployTitle,
			Openshift:   &openShift,
		}

		jsonStr, _ := json.Marshal(environmentFlags)
		var parsedFlags map[string]string
		json.Unmarshal(jsonStr, &parsedFlags)
		fmt.Println(parsedFlags)

		result, err := l.UpdateEnvironment(context.TODO(), environment.ID, environmentFlags, lc)
		handleError(err)
		fmt.Println(result)

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
	updateEnvironmentCmd.Flags().String("name", "", "TODO")
	updateEnvironmentCmd.Flags().String("deploy-base-ref", "", "TODO")
	updateEnvironmentCmd.Flags().String("deploy-head-ref", "", "TODO")
	updateEnvironmentCmd.Flags().String("deploy-title", "", "TODO")
	updateEnvironmentCmd.Flags().String("namespace", "", "TODO")
	updateEnvironmentCmd.Flags().String("route", "", "TODO")
	updateEnvironmentCmd.Flags().String("routes", "", "TODO")
	updateEnvironmentCmd.Flags().UintP("auto-idle", "a", 0, "TODO")
	updateEnvironmentCmd.Flags().UintP("deploy-target", "d", 0, "TODO")
	updateEnvironmentCmd.Flags().String("created", "", "TODO")
	updateEnvironmentCmd.Flags().String("environment-type", "", "TODO")
	updateEnvironmentCmd.Flags().String("deploy-type", "", "TODO")
	//updateEnvironmentCmd.Flags().StringVarP(&environmentType, "environmentType", "t", "", "TODO")
	//updateEnvironmentCmd.Flags().Var(&dt, "deployType", "TODO")
	//updateEnvironmentCmd.Flags().StringSlice(environmentType, []string{}, "TODO")
	//updateEnvironmentCmd.Flags().StringSlice(deployType, []string{}, "")
}
