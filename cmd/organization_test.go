package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/spf13/cobra"
)

func TestOrganizationCommands(t *testing.T) {
	tests := []struct {
		name      string
		cmdArgs   []string
		setupCmd  func(*cobra.Command, pflag.FlagSet)
		expectOut []string
		expectErr bool
	}{
		{
			name: "Add Organization",
			cmdArgs: []string{"add", "organization",
				"--organization-name=test-organization",
				"--friendly-name=Test Organization",
				"--description=A test organization",
				"--project-quota=10",
				"--environment-quota=20",
				"--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(addCmd)
				addCmd.AddCommand(addOrganizationCmd)
				AddGenericFlags(addOrganizationCmd)
			},
			expectOut: []string{"success", "test-organization"},
			expectErr: false,
		},
		{
			name: "Delete Organization",
			cmdArgs: []string{"delete", "organization",
				"--organization-name=test-organization",
				"--force",
				"--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(deleteCmd)
				deleteCmd.AddCommand(deleteOrganizationCmd)
				AddGenericFlags(deleteOrganizationCmd)
			},
			expectOut: []string{"result", "test-organization"},
			expectErr: false,
		},
		// TODO: Remaining Organization commands
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
