package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/spf13/cobra"
)

func TestUserCommands(t *testing.T) {
	tests := []struct {
		name              string
		cmdArgs           []string
		setupCmd          func(*cobra.Command, pflag.FlagSet)
		expectOut         []string
		expectErr         bool
		expectedErrString string
	}{
		{
			name:    "Add User",
			cmdArgs: []string{"add", "user", "--email=user@test.com", "--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(addCmd)
				addCmd.AddCommand(addUserCmd)
				AddGenericFlags(addUserCmd)
			},
			expectOut: []string{"success"},
			expectErr: false,
		},
		{
			name:    "Add User SSH Key",
			cmdArgs: []string{"add", "user-sshkey", "--email=user@test.com", "--keyvalue=ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAINA0ITV2gbDc6noYeWaqfxTYpaEKq7HzU3+F71XGhSL/", "--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(addCmd)
				addCmd.AddCommand(addUserSSHKeyCmd)
				AddGenericFlags(addUserSSHKeyCmd)
			},
			expectOut: []string{"success"},
			expectErr: false,
		},
		{
			name:    "Delete User SSH Key",
			cmdArgs: []string{"delete", "user-sshkey", "--id=2", "--output-json", "--force"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(deleteCmd)
				deleteCmd.AddCommand(deleteSSHKeyCmd)
				AddGenericFlags(deleteSSHKeyCmd)
			},
			expectOut: []string{"success"},
			expectErr: false,
		},
		{
			name:    "Update User",
			cmdArgs: []string{"update", "user", "--current-email=user@test.com", "--first-name=test", "--last-name=user", "--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(updateCmd)
				updateCmd.AddCommand(updateUserCmd)
				AddGenericFlags(updateUserCmd)
			},
			expectOut: []string{"success"},
			expectErr: false,
		},
		{
			name:    "Delete User",
			cmdArgs: []string{"delete", "user", "--email=user@test.com", "--output-json", "--force"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(deleteCmd)
				deleteCmd.AddCommand(deleteUserCmd)
				AddGenericFlags(deleteUserCmd)
			},
			expectOut: []string{"success"},
			expectErr: false,
		},
		{
			name:    "Get User SSK Keys",
			cmdArgs: []string{"get", "user-sshkeys", "--email=default-user@lagoon-demo", "--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(getCmd)
				getCmd.AddCommand(getUserKeysCmd)
				AddGenericFlags(getUserKeysCmd)
			},
			expectOut: []string{"default-user@lagoon-demo", "ssh-ed25519"},
			expectErr: false,
		},
		{
			name:    "Get All User SSK Keys",
			cmdArgs: []string{"get", "all-user-sshkeys", "--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(getCmd)
				getCmd.AddCommand(getAllUserKeysCmd)
				AddGenericFlags(getAllUserKeysCmd)
			},
			expectOut: []string{"default-user@lagoon-demo", "ci-customer-user-rsa@example.com", "ci-customer-user-ecdsa@example.com"},
			expectErr: false,
		},
		{
			name:    "Get All User SSK Keys in group",
			cmdArgs: []string{"get", "all-user-sshkeys", "--name=ci-group", "--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(getCmd)
				getCmd.AddCommand(getAllUserKeysCmd)
				AddGenericFlags(getAllUserKeysCmd)
			},
			expectOut: []string{"ci-customer-user-rsa@example.com", "ci-customer-user-ecdsa@example.com"},
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

			cmd.SetArgs(tt.cmdArgs)

			err := cmd.Execute()
			if err != nil && tt.expectErr {
				assert.Contains(t, err.Error(), tt.expectedErrString)
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
