package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/spf13/cobra"
)

func TestEnvironmentCommands(t *testing.T) {
	tests := []struct {
		name              string
		cmdArgs           []string
		setupCmd          func(*cobra.Command, pflag.FlagSet)
		expectOut         []string
		expectErr         bool
		expectedErrString string
	}{
		{
			name:    "List Backups",
			cmdArgs: []string{"list", "backups", "--project=lagoon-demo", "--environment=main", "--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(listCmd)
				listCmd.AddCommand(listBackupsCmd)
				AddGenericFlags(listBackupsCmd)
			},
			expectOut: []string{"e2e1d31b4a7dfc1687f469b6673f6bf2c0aabee0cc6b3f1bdbde710a9bc6280f", "files", "e2e1d31b4a7dfc1687f469b6673f6bf2c0aabee0cc6b3f1bdbde710a9bc6280d", "mariadb"},
			expectErr: false,
		},
		{
			name:    "Get Backup - Error: no download file found",
			cmdArgs: []string{"get", "backup", "--project=lagoon-demo", "--environment=main", "--backup-id=e260f07c374e4a3319c5d46e688dab6f1eb23c3e61c166a37609d55762d2ee0b", "--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(getCmd)
				getCmd.AddCommand(getBackupCmd)
				AddGenericFlags(getBackupCmd)
			},
			expectOut:         []string{""},
			expectErr:         true,
			expectedErrString: "no download file found, status of backups restoration is failed",
		},
		{
			name:    "Get Backup - Error: backup has not been restored",
			cmdArgs: []string{"get", "backup", "--project=lagoon-demo", "--environment=main", "--backup-id=bf072a09e17726da54adc79936ec8745521993599d41211dfc9466dfd5bc32a5", "--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(getCmd)
				getCmd.AddCommand(getBackupCmd)
				AddGenericFlags(getBackupCmd)
			},
			expectOut:         []string{""},
			expectErr:         true,
			expectedErrString: "backup has not been restored",
		},
		// TODO: Seed backup data to test success path (getBackup)
		{
			name:    "Update Environment",
			cmdArgs: []string{"update", "environment", "--project=lagoon-demo", "--environment=pr-175", "--auto-idle=0", "--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(updateCmd)
				updateCmd.AddCommand(updateEnvironmentCmd)
				AddGenericFlags(updateEnvironmentCmd)
			},
			expectOut: []string{"success", "pr-175"},
			expectErr: false,
		},
		{
			name:    "Delete Environment",
			cmdArgs: []string{"delete", "environment", "--project=lagoon-demo", "--environment=pr-175", "--output-json", "--force"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(deleteCmd)
				deleteCmd.AddCommand(deleteEnvCmd)
				AddGenericFlags(deleteEnvCmd)
			},
			expectOut: []string{"success"},
			expectErr: false,
		},
	}
	SetUpRootCmdFlags()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{Use: "root"}
			cmd.SetArgs(tt.cmdArgs)
			flags := pflag.FlagSet{}
			tt.setupCmd(cmd, flags)

			var out bytes.Buffer
			cmd.SetOut(&out)
			cmd.SetErr(&out)

			err := cmd.Execute()
			if err != nil && tt.expectErr {
				assert.NotEmpty(t, err)
				fmt.Println("err:", err)
				return
			} else if err != nil {
				t.Fatalf("Error executing command: %v", err)
			}

			for _, eo := range tt.expectOut {
				assert.Contains(t, out.String(), eo)
			}

		})
	}
}
