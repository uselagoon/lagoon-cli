package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/spf13/cobra"
)

func TestGetCommands(t *testing.T) {
	tests := []struct {
		name      string
		cmdArgs   []string
		setupCmd  func(*cobra.Command, pflag.FlagSet)
		expectOut []string
		expectErr bool
	}{
		{
			name:    "Get Organization",
			cmdArgs: []string{"get", "organization", "--organization-name=lagoon-demo-organization", "--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(getCmd)
				getCmd.AddCommand(getOrganizationCmd)
				AddGenericFlags(getOrganizationCmd)
			},
			expectOut: []string{"lagoon-demo-organization", "An organization for testing"},
			expectErr: false,
		},
		{
			name:    "Get Project",
			cmdArgs: []string{"get", "project", "--project=lagoon-demo", "--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(getCmd)
				getCmd.AddCommand(getProjectCmd)
				AddGenericFlags(getProjectCmd)
			},
			expectOut: []string{"lagoon-demo", "ssh://git@example.com/lagoon-demo.git"},
			expectErr: false,
		},
		{
			name:    "Get Environment",
			cmdArgs: []string{"get", "environment", "--project=lagoon-demo", "--environment=staging", "--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(getCmd)
				getCmd.AddCommand(getEnvironmentCmd)
				AddGenericFlags(getEnvironmentCmd)
			},
			expectOut: []string{"staging", "development", "branch"},
			expectErr: false,
		},
		// TODO: set static value for project-key in seed
		{
			name:    "Get Project Key",
			cmdArgs: []string{"get", "project-key", "--project=lagoon-demo", "--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(getCmd)
				getCmd.AddCommand(getProjectKeyCmd)
				AddGenericFlags(getProjectKeyCmd)
			},
			expectOut: []string{"ssh-ed25519"},
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
