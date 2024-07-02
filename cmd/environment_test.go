package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/spf13/cobra"
)

func TestAPIEnvironmentCommands(t *testing.T) {
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
			cmdArgs: []string{"list", "backups", "--project=lagoon-demo", "--environment=main"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(listCmd)
				listCmd.AddCommand(listBackupsCmd)
			},
			expectOut: []string{"e2e1d31b4a7dfc1687f469b6673f6bf2c0aabee0cc6b3f1bdbde710a9bc6280f", "files", "e2e1d31b4a7dfc1687f469b6673f6bf2c0aabee0cc6b3f1bdbde710a9bc6280d", "mariadb"},
			expectErr: false,
		},
		{
			name:    "Update Environment",
			cmdArgs: []string{"update", "environment", "--project=lagoon-demo", "--environment=pr-175", "--auto-idle=0"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(updateCmd)
				updateCmd.AddCommand(updateEnvironmentCmd)
			},
			expectOut: []string{"success", "pr-175"},
			expectErr: false,
		},
		{
			name:    "Delete Environment",
			cmdArgs: []string{"delete", "environment", "--project=lagoon-demo", "--environment=pr-175", "--force"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(deleteCmd)
				deleteCmd.AddCommand(deleteEnvCmd)
			},
			expectOut: []string{"success"},
			expectErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := rootCmd
			tt.cmdArgs = append(tt.cmdArgs, "--output-json", "--config-file=../temp_config.yaml")
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
