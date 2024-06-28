package cmd

//import (
//	"bytes"
//	"fmt"
//	"github.com/spf13/pflag"
//	"github.com/stretchr/testify/assert"
//	"testing"
//
//	"github.com/spf13/cobra"
//)
//
//func TestUserCommands(t *testing.T) {
//	tests := []struct {
//		name              string
//		cmdArgs           []string
//		setupCmd          func(*cobra.Command, pflag.FlagSet)
//		expectOut         []string
//		expectErr         bool
//		expectedErrString string
//	}{
//		{
//			name:    "Add User",
//			cmdArgs: []string{"add", "user", "--email=user@test.com"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(addCmd)
//				addCmd.AddCommand(addUserCmd)
//			},
//			expectOut: []string{"success"},
//			expectErr: false,
//		},
//		{
//			name:    "Add User SSH Key",
//			cmdArgs: []string{"add", "user-sshkey", "--email=user@test.com", "--keyvalue=ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAINA0ITV2gbDc6noYeWaqfxTYpaEKq7HzU3+F71XGhSL/"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(addCmd)
//				addCmd.AddCommand(addUserSSHKeyCmd)
//			},
//			expectOut: []string{"success"},
//			expectErr: false,
//		},
//		{
//			name:    "Delete User SSH Key",
//			cmdArgs: []string{"delete", "user-sshkey", "--id=2", "--force"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(deleteCmd)
//				deleteCmd.AddCommand(deleteSSHKeyCmd)
//			},
//			expectOut: []string{"success"},
//			expectErr: false,
//		},
//		{
//			name:    "Update User",
//			cmdArgs: []string{"update", "user", "--current-email=user@test.com", "--first-name=test", "--last-name=user"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(updateCmd)
//				updateCmd.AddCommand(updateUserCmd)
//			},
//			expectOut: []string{"success"},
//			expectErr: false,
//		},
//		{
//			name:    "Delete User",
//			cmdArgs: []string{"delete", "user", "--email=user@test.com", "--force"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(deleteCmd)
//				deleteCmd.AddCommand(deleteUserCmd)
//			},
//			expectOut: []string{"success"},
//			expectErr: false,
//		},
//		{
//			name:    "Get User SSK Keys",
//			cmdArgs: []string{"get", "user-sshkeys", "--email=default-user@lagoon-demo"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(getCmd)
//				getCmd.AddCommand(getUserKeysCmd)
//			},
//			expectOut: []string{"default-user@lagoon-demo", "ssh-ed25519"},
//			expectErr: false,
//		},
//		{
//			name:    "Get All User SSK Keys",
//			cmdArgs: []string{"get", "all-user-sshkeys"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(getCmd)
//				getCmd.AddCommand(getAllUserKeysCmd)
//			},
//			expectOut: []string{"default-user@lagoon-demo", "ci-customer-user-rsa@example.com", "ci-customer-user-ecdsa@example.com"},
//			expectErr: false,
//		},
//		{
//			name:    "Get All User SSK Keys in group",
//			cmdArgs: []string{"get", "all-user-sshkeys", "--name=ci-group"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(getCmd)
//				getCmd.AddCommand(getAllUserKeysCmd)
//			},
//			expectOut: []string{"ci-customer-user-rsa@example.com", "ci-customer-user-ecdsa@example.com"},
//			expectErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			cmd := rootCmd
//			tt.cmdArgs = append(tt.cmdArgs, "--output-json", "--config-file=../temp_config.yaml")
//			cmd.SetArgs(tt.cmdArgs)
//			flags := pflag.FlagSet{}
//			tt.setupCmd(cmd, flags)
//
//			var out bytes.Buffer
//			cmd.SetOut(&out)
//			cmd.SetErr(&out)
//
//			cmd.SetArgs(tt.cmdArgs)
//
//			err := cmd.Execute()
//			if err != nil && tt.expectErr {
//				assert.Contains(t, err.Error(), tt.expectedErrString)
//				fmt.Println("err:", err)
//				return
//			} else if err != nil {
//				t.Fatalf("Error executing command: %v", err)
//			}
//
//			for _, eo := range tt.expectOut {
//				assert.Contains(t, out.String(), eo)
//			}
//		})
//	}
//}
