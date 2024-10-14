package cmd

import (
	"context"
	"fmt"
	"github.com/uselagoon/lagoon-cli/pkg/output"
	"strings"

	"github.com/spf13/cobra"
	"github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
)

var retrieveCmd = &cobra.Command{
	Use:     "retrieve",
	Aliases: []string{"re", "ret"},
	Short:   "Trigger a retrieval operation on backups",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateToken(lContext.Name) // get a new token if the current one is invalid
	},
}

var retrieveBackupCmd = &cobra.Command{
	Use:     "backup",
	Aliases: []string{"b"},
	Hidden:  false,
	Short:   "Retrieve a backup",
	Long: `Retrieve a backup
Given a backup-id, you can initiate a retrieval for it.
You can check the status of the backup using the list backups or get backup command.`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lContext.Name)
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
		if err := requiredInputCheck("Backup ID", backupID); err != nil {
			return err
		}
		if yesNo(fmt.Sprintf("You are attempting to trigger a retrieval for backup ID '%s', are you sure?", backupID)) {
			utoken := lUser.UserConfig.Grant.AccessToken
			lc := lclient.New(
				fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
				lagoonCLIVersion,
				lContext.ContextConfig.Version,
				&utoken,
				debug)
			result, err := lagoon.AddBackupRestore(context.TODO(), backupID, lc)
			if err != nil {
				if strings.Contains(err.Error(), "Duplicate entry") {
					// this error reports a lot about the sql backup, need to fix that in Lagoon upstream
					return fmt.Errorf("retrieval for %s has already been created", backupID)
				}
				return err
			}
			resultData := output.Result{Result: fmt.Sprintf("successfully created restore with ID: %d", result.ID)}
			r := output.RenderResult(resultData, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		}
		return nil
	},
}

func init() {
	retrieveCmd.AddCommand(retrieveBackupCmd)
	retrieveBackupCmd.Flags().StringP("backup-id", "B", "", "The backup ID you want to commence a retrieval for")
}
