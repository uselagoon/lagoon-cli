package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/spf13/cobra"
)

func TestDeployTargetConfigCommands(t *testing.T) {
	tests := []struct {
		name      string
		cmdArgs   []string
		setupCmd  func(*cobra.Command, pflag.FlagSet)
		expectOut []string
		expectErr bool
	}{
		{
			name:    "Add Deploytarget-config",
			cmdArgs: []string{"add", "deploytarget-config", "--project=lagoon-demo", "--deploytarget=2001", "--pullrequests=true", "--branches=false", "--output-json", "--force"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(addCmd)
				addCmd.AddCommand(addDeployTargetConfigCmd)
				AddGenericFlags(addDeployTargetConfigCmd)
			},
			expectOut: []string{"true", "false", "ci-local-control-k8s"},
			expectErr: false,
		},
		// TODO: Seed deploytarget-config data
		{
			name:    "List Deploytarget-configs",
			cmdArgs: []string{"list", "deploytarget-configs", "--project=lagoon-demo", "--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(listCmd)
				listCmd.AddCommand(listDeployTargetConfigsCmd)
				AddGenericFlags(listDeployTargetConfigsCmd)
			},
			expectOut: []string{"true", "false", "ci-local-control-k8s"},
			expectErr: false,
		},
		{
			name:    "Update Deploytarget-config",
			cmdArgs: []string{"update", "deploytarget-config", "--id=1", "--weight=2", "--output-json", "--force"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(updateCmd)
				updateCmd.AddCommand(updateDeployTargetConfigCmd)
				AddGenericFlags(updateDeployTargetConfigCmd)
			},
			expectOut: []string{"2", "ci-local-control-k8s"},
			expectErr: false,
		},
		{
			name:    "Delete Deploytarget-config",
			cmdArgs: []string{"delete", "deploytarget-config", "--project=lagoon-demo", "--id=1", "--output-json", "--force"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(deleteCmd)
				deleteCmd.AddCommand(deleteDeployTargetConfigCmd)
				AddGenericFlags(deleteDeployTargetConfigCmd)
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
