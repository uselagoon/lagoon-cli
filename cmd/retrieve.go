package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	l "github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
)

var retrieveCmd = &cobra.Command{
	Use:     "retrieve",
	Aliases: []string{"re", "ret"},
	Short:   "Trigger a retrieval operation on backups",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateToken(lagoonCLIConfig.Current) // get a new token if the current one is invalid
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
		if err := requiredInputCheck("Backup ID", backupID); err != nil {
			return err
		}
		if yesNo(fmt.Sprintf("You are attempting to trigger a retrieval for backup ID '%s', are you sure?", backupID)) {
			current := lagoonCLIConfig.Current
			token := lagoonCLIConfig.Lagoons[current].Token
			lc := lclient.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIVersion,
				&token,
				debug)
			result, err := l.AddBackupRestore(context.TODO(), backupID, lc)
			if err != nil {
				if strings.Contains(err.Error(), "Duplicate entry") {
					// this error reports a lot about the sql backup, need to fix that in Lagoon upstream
					return fmt.Errorf("retrieval for %s has already been created", backupID)
				}
				return err
			}
			fmt.Println("successfully created restore with ID:", result.ID)
		}
		return nil
	},
}

func init() {
	retrieveCmd.AddCommand(retrieveBackupCmd)
	retrieveBackupCmd.Flags().StringP("backup-id", "B", "", "The backup ID you want to commence a retrieval for")
}
