package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/uselagoon/machinery/api/schema"

	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/pkg/output"
	"github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
)

var deleteEnvCmd = &cobra.Command{
	Use:     "environment",
	Aliases: []string{"e"},
	Short:   "Delete an environment",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		// environmentFlags := parseEnvironmentFlags(*cmd.Flags()) //@TODO re-enable this at some point if more environment based commands are made available
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
		if yesNo(fmt.Sprintf("You are attempting to delete environment '%s' from project '%s', are you sure?", cmdProjectEnvironment, cmdProjectName)) {
			environment, err := lagoon.DeleteEnvironment(context.TODO(), cmdProjectEnvironment, cmdProjectName, true, lc)
			if err != nil {
				return err
			}
			resultData := output.Result{
				Result: environment.DeleteEnvironment,
			}
			r := output.RenderResult(resultData, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		}
		return nil
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
		environmentType, err := cmd.Flags().GetString("environment-type")
		if err != nil {
			return err
		}
		deployT, err := cmd.Flags().GetString("deploy-type")
		if err != nil {
			return err
		}
		deploytarget, err := cmd.Flags().GetUint("deploytarget")
		if err != nil {
			return err
		}
		deployTitle, err := cmd.Flags().GetString("deploy-title")
		if err != nil {
			return err
		}
		autoIdle, err := cmd.Flags().GetBool("auto-idle")
		if err != nil {
			return err
		}
		autoIdleProvided := cmd.Flags().Lookup("auto-idle").Changed

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

		environment, err := lagoon.GetEnvironmentByNameAndProjectName(context.TODO(), cmdProjectEnvironment, cmdProjectName, lc)
		if err != nil {
			return fmt.Errorf("%v: check if the project or environment exists", err.Error())
		}

		environmentFlags := schema.UpdateEnvironmentPatchInput{
			DeployBaseRef:        nullStrCheck(deployBaseRef),
			DeployHeadRef:        nullStrCheck(deployHeadRef),
			OpenshiftProjectName: nullStrCheck(namespace),
			Route:                nullStrCheck(route),
			Routes:               nullStrCheck(routes),
			DeployTitle:          nullStrCheck(deployTitle),
			Openshift:            nullUintCheck(deploytarget),
		}
		if autoIdleProvided {
			environmentFlags.AutoIdle = nullBoolToUint(autoIdle)
		}
		if environmentType != "" {
			envType := schema.EnvType(strings.ToUpper(environmentType))
			if validationErr := schema.ValidateType(envType); validationErr != nil {
				handleError(validationErr)
			}
			environmentFlags.EnvironmentType = &envType
		}
		if deployT != "" {
			deployType := schema.DeployType(strings.ToUpper(deployT))
			if validationErr := schema.ValidateType(deployType); validationErr != nil {
				handleError(validationErr)
			}
			environmentFlags.DeployType = &deployType
		}

		result, err := lagoon.UpdateEnvironment(context.TODO(), environment.ID, environmentFlags, lc)
		if err != nil {
			return err
		}

		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"Environment Name": result.Name,
			},
		}
		r := output.RenderResult(resultData, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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

		backupsResult, err := lagoon.GetBackupsForEnvironmentByNameAndProjectName(context.TODO(), cmdProjectEnvironment, cmdProjectName, lc)
		if err != nil {
			return fmt.Errorf("%v: check if the project or environment exists", err.Error())
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
		r := output.RenderOutput(output.Table{
			Header: []string{
				"BackupID",
				"Source",
				"Created",
				"Restored",
				"RestoreStatus",
			},
			Data: data,
		}, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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

		backupsResult, err := lagoon.GetBackupsForEnvironmentByNameAndProjectName(context.TODO(), cmdProjectEnvironment, cmdProjectName, lc)
		if err != nil {
			return fmt.Errorf("%v: check if the project or environment exists", err.Error())
		}
		status := ""
		for _, backup := range backupsResult.Backups {
			if backup.BackupID == backupID {
				if backup.Restore.RestoreLocation != "" {
					resultData := output.Result{Result: backup.Restore.RestoreLocation}
					r := output.RenderResult(resultData, outputOptions)
					fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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

func init() {
	getCmd.AddCommand(getBackupCmd)
	getBackupCmd.Flags().StringP("backup-id", "B", "", "The backup ID you want to restore")
	updateEnvironmentCmd.Flags().String("deploy-base-ref", "", "Updates the deploy base ref for the selected environment")
	updateEnvironmentCmd.Flags().String("deploy-head-ref", "", "Updates the deploy head ref for the selected environment")
	updateEnvironmentCmd.Flags().String("deploy-title", "", "Updates the deploy title for the selected environment")
	updateEnvironmentCmd.Flags().String("namespace", "", "Update the namespace for the selected environment")
	updateEnvironmentCmd.Flags().String("route", "", "Update the route for the selected environment")
	updateEnvironmentCmd.Flags().String("routes", "", "Update the routes for the selected environment")
	updateEnvironmentCmd.Flags().BoolP("auto-idle", "a", false, "Auto idle setting of the environment. Set to enable, --auto-idle=false to disable")
	updateEnvironmentCmd.Flags().UintP("deploytarget", "d", 0, "Reference to Deploytarget(Kubernetes) this Environment should be deployed to")
	updateEnvironmentCmd.Flags().String("environment-type", "", "Update the environment type - production | development")
	updateEnvironmentCmd.Flags().String("deploy-type", "", "Update the deploy type - branch | pullrequest | promote")
}
