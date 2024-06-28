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
//func TestWebhookNotificationCommands(t *testing.T) {
//	tests := []struct {
//		name      string
//		cmdArgs   []string
//		setupCmd  func(*cobra.Command, pflag.FlagSet)
//		expectOut []string
//		expectErr bool
//	}{
//		{
//			name:    "Add Webhook Notification",
//			cmdArgs: []string{"add", "notification", "webhook", "--name=webhook-notification", "--webhook=test-webhook", "--force"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(addCmd)
//				addCmd.AddCommand(addNotificationWebhookCmd)
//			},
//			expectOut: []string{"webhook-notification", "test-webhook"},
//			expectErr: false,
//		},
//		{
//			name:    "Add Webhook Notification to Project",
//			cmdArgs: []string{"add", "notification", "project-webhook", "--name=webhook-notification", "--project=lagoon-demo", "--force"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(addCmd)
//				addCmd.AddCommand(addProjectNotificationWebhookCmd)
//			},
//			expectOut: []string{"success"},
//			expectErr: false,
//		},
//		{
//			name:    "List Project Webhook Notifications",
//			cmdArgs: []string{"list", "notification", "project-webhook", "--project=lagoon-demo"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(listCmd)
//				listCmd.AddCommand(listProjectWebhooksCmd)
//			},
//			expectOut: []string{"webhook-notification", "test-webhook"},
//			expectErr: false,
//		},
//		{
//			name:    "List all Webhook Notifications",
//			cmdArgs: []string{"list", "notification", "webhook"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(listCmd)
//				listCmd.AddCommand(listAllWebhooksCmd)
//			},
//			expectOut: []string{"lagoon-demo", "webhook-notification", "test-webhook"},
//			expectErr: false,
//		},
//		// Unable test newname as incorrect data is returned via the API (fixed in PR#3706)
//		{
//			name:    "Update a Webhook Notification",
//			cmdArgs: []string{"update", "notification", "webhook", "--name=webhook-notification", "--webhook=new-webhook-test", "--force"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(updateCmd)
//				updateCmd.AddCommand(updateWebhookNotificationCmd)
//			},
//			expectOut: []string{"webhook-notification", "new-webhook-test"},
//			expectErr: false,
//		},
//		{
//			name:    "Delete a Webhook Notification from a Project",
//			cmdArgs: []string{"delete", "notification", "project-webhook", "--name=webhook-notification", "--project=lagoon-demo", "--force"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(deleteCmd)
//				deleteCmd.AddCommand(deleteProjectWebhookNotificationCmd)
//			},
//			expectOut: []string{"success"},
//			expectErr: false,
//		},
//		{
//			name:    "Delete a Webhook Notification",
//			cmdArgs: []string{"delete", "notification", "webhook", "--name=webhook-notification", "--force"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(deleteCmd)
//				deleteCmd.AddCommand(deleteWebhookNotificationCmd)
//			},
//			expectOut: []string{"success"},
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
//			err := cmd.Execute()
//			if err != nil && tt.expectErr {
//				assert.NotEmpty(t, err)
//				fmt.Println("err:", err)
//				return
//			} else if err != nil {
//				t.Fatalf("Error executing command: %v", err)
//			}
//
//			for _, eo := range tt.expectOut {
//				assert.Contains(t, out.String(), eo)
//			}
//
//		})
//	}
//}
