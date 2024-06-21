package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/spf13/cobra"
)

func TestGroupCommands(t *testing.T) {
	tests := []struct {
		name              string
		cmdArgs           []string
		setupCmd          func(*cobra.Command, pflag.FlagSet)
		expectOut         []string
		expectErr         bool
		expectedErrString string
	}{
		{
			name:    "Add Group",
			cmdArgs: []string{"add", "group", "--name=test-group", "--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(addCmd)
				addCmd.AddCommand(addGroupCmd)
				AddGenericFlags(addGroupCmd)
			},
			expectOut: []string{"success", "test-group"},
			expectErr: false,
		},
		{
			name:    "Delete Group",
			cmdArgs: []string{"delete", "group", "--name=test-group", "--output-json", "--force"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(deleteCmd)
				deleteCmd.AddCommand(deleteGroupCmd)
				AddGenericFlags(deleteGroupCmd)
			},
			expectOut: []string{"success"},
			expectErr: false,
		},
		{
			name:    "Add Group to Organization",
			cmdArgs: []string{"add", "group", "--name=test-organization-group", "--organization-name=lagoon-demo-organization", "--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(addCmd)
				addCmd.AddCommand(addGroupCmd)
				AddGenericFlags(addGroupCmd)
			},
			expectOut: []string{"success", "test-organization-group", "lagoon-demo-organization"},
			expectErr: false,
		},
		{
			name:    "Add User to Group",
			cmdArgs: []string{"add", "user-group", "--name=lagoon-demo-group", "--email=ci-customer-user-ecdsa@example.com", "--role=guest", "--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(addCmd)
				addCmd.AddCommand(addUserToGroupCmd)
				AddGenericFlags(addUserToGroupCmd)
			},
			expectOut: []string{"success"},
			expectErr: false,
		},
		{
			name:    "Delete User from Group",
			cmdArgs: []string{"delete", "user-group", "--name=lagoon-demo-group", "--email=ci-customer-user-ecdsa@example.com", "--output-json", "--force"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(deleteCmd)
				deleteCmd.AddCommand(deleteUserFromGroupCmd)
				AddGenericFlags(deleteUserFromGroupCmd)
			},
			expectOut: []string{"success"},
			expectErr: false,
		},
		{
			name:    "Add Project to Group",
			cmdArgs: []string{"add", "project-group", "--name=ci-group", "--project=lagoon-demo", "--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(addCmd)
				addCmd.AddCommand(addProjectToGroupCmd)
				AddGenericFlags(addProjectToGroupCmd)
			},
			expectOut: []string{"success"},
			expectErr: false,
		},
		{
			name:    "Delete Project from Group",
			cmdArgs: []string{"delete", "project-group", "--name=ci-group", "--project=lagoon-demo", "--output-json", "--force"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(deleteCmd)
				deleteCmd.AddCommand(deleteProjectFromGroupCmd)
				AddGenericFlags(deleteProjectFromGroupCmd)
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
