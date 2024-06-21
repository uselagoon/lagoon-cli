package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/spf13/cobra"
)

func TestDeployTargetCommands(t *testing.T) {
	tests := []struct {
		name      string
		cmdArgs   []string
		setupCmd  func(*cobra.Command, pflag.FlagSet)
		expectOut []string
		expectErr bool
	}{
		{
			name:    "Add Deploytarget",
			cmdArgs: []string{"add", "deploytarget", "--name=test-deploytarget", "--console-url=https://localhost:3000/", "--token=abcd1234", "--output-json", "--force"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(addCmd)
				addCmd.AddCommand(addDeployTargetCmd)
				AddGenericFlags(addDeployTargetCmd)
			},
			expectOut: []string{"test-deploytarget", "https://localhost:3000/", "abcd1234"},
			expectErr: false,
		},
		{
			name:    "Update Deploytarget",
			cmdArgs: []string{"update", "deploytarget", "--id=4", "--friendly-name=ui-kubernetes-deploytarget", "--output-json", "--force"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(updateCmd)
				updateCmd.AddCommand(updateDeployTargetCmd)
				AddGenericFlags(updateDeployTargetCmd)
			},
			expectOut: []string{"ui-kubernetes", "ui-kubernetes-deploytarget"},
			expectErr: false,
		},
		{
			name:    "Delete Deploytarget",
			cmdArgs: []string{"delete", "deploytarget", "--name=test-deploytarget", "--output-json", "--force"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(deleteCmd)
				deleteCmd.AddCommand(deleteDeployTargetCmd)
				AddGenericFlags(deleteDeployTargetCmd)
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
