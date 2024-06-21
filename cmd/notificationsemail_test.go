package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/spf13/cobra"
)

func TestEmailNotificationCommands(t *testing.T) {
	tests := []struct {
		name      string
		cmdArgs   []string
		setupCmd  func(*cobra.Command, pflag.FlagSet)
		expectOut []string
		expectErr bool
	}{
		{
			name:    "Add Email Notification",
			cmdArgs: []string{"add", "notification", "email", "--name=email-notification", "--email=test@test.com", "--output-json", "--force"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(addCmd)
				addCmd.AddCommand(addNotificationEmailCmd)
				AddGenericFlags(addNotificationEmailCmd)
			},
			expectOut: []string{"email-notification", "test@test.com"},
			expectErr: false,
		},
		{
			name:    "Add Email Notification to Project",
			cmdArgs: []string{"add", "notification", "project-email", "--name=email-notification", "--project=lagoon-demo", "--output-json", "--force"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(addCmd)
				addCmd.AddCommand(addProjectNotificationEmailCmd)
				AddGenericFlags(addProjectNotificationEmailCmd)
			},
			expectOut: []string{"success"},
			expectErr: false,
		},
		{
			name:    "List Project Email Notifications",
			cmdArgs: []string{"list", "notification", "project-email", "--project=lagoon-demo", "--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(listCmd)
				listCmd.AddCommand(listProjectEmailsCmd)
				AddGenericFlags(listProjectEmailsCmd)
			},
			expectOut: []string{"email-notification", "test@test.com"},
			expectErr: false,
		},
		{
			name:    "List all Email Notifications",
			cmdArgs: []string{"list", "notification", "email", "--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(listCmd)
				listCmd.AddCommand(listAllEmailsCmd)
				AddGenericFlags(listAllEmailsCmd)
			},
			expectOut: []string{"lagoon-demo", "email-notification", "test@test.com"},
			expectErr: false,
		},
		// Unable test newname as incorrect data is returned via the API (fixed in PR#3706)
		{
			name:    "Update an Email Notification",
			cmdArgs: []string{"update", "notification", "email", "--name=email-notification", "--email=newemail@test.com", "--output-json", "--force"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(updateCmd)
				updateCmd.AddCommand(updateEmailNotificationCmd)
				AddGenericFlags(updateEmailNotificationCmd)
			},
			expectOut: []string{"email-notification", "newemail@test.com"},
			expectErr: false,
		},
		{
			name:    "Delete an Email Notification from a Project",
			cmdArgs: []string{"delete", "notification", "project-email", "--name=email-notification", "--project=lagoon-demo", "--output-json", "--force"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(deleteCmd)
				deleteCmd.AddCommand(deleteProjectEmailNotificationCmd)
				AddGenericFlags(deleteProjectEmailNotificationCmd)
			},
			expectOut: []string{"success"},
			expectErr: false,
		},
		{
			name:    "Delete an Email Notification",
			cmdArgs: []string{"delete", "notification", "email", "--name=email-notification", "--output-json", "--force"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(deleteCmd)
				deleteCmd.AddCommand(deleteEmailNotificationCmd)
				AddGenericFlags(deleteEmailNotificationCmd)
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
