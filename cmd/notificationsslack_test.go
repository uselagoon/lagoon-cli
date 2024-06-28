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
//func TestSlackNotificationCommands(t *testing.T) {
//	tests := []struct {
//		name      string
//		cmdArgs   []string
//		setupCmd  func(*cobra.Command, pflag.FlagSet)
//		expectOut []string
//		expectErr bool
//	}{
//		{
//			name:    "Add Slack Notification",
//			cmdArgs: []string{"add", "notification", "slack", "--name=slack-notification", "--channel=test-channel", "--webhook=test-webhook", "--force"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(addCmd)
//				addCmd.AddCommand(addNotificationSlackCmd)
//			},
//			expectOut: []string{"slack-notification", "test-channel", "test-webhook"},
//			expectErr: false,
//		},
//		{
//			name:    "Add Slack Notification to Project",
//			cmdArgs: []string{"add", "notification", "project-slack", "--name=slack-notification", "--project=lagoon-demo", "--force"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(addCmd)
//				addCmd.AddCommand(addProjectNotificationSlackCmd)
//			},
//			expectOut: []string{"success"},
//			expectErr: false,
//		},
//		{
//			name:    "List Project Slack Notifications",
//			cmdArgs: []string{"list", "notification", "project-slack", "--project=lagoon-demo"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(listCmd)
//				listCmd.AddCommand(listProjectSlacksCmd)
//			},
//			expectOut: []string{"slack-notification", "test-channel", "test-webhook"},
//			expectErr: false,
//		},
//		{
//			name:    "List all Slack Notifications",
//			cmdArgs: []string{"list", "notification", "slack"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(listCmd)
//				listCmd.AddCommand(listAllSlacksCmd)
//			},
//			expectOut: []string{"lagoon-demo", "slack-notification", "test-channel", "test-webhook"},
//			expectErr: false,
//		},
//		// Unable test newname as incorrect data is returned via the API (fixed in PR#3706)
//		{
//			name:    "Update a Slack Notification",
//			cmdArgs: []string{"update", "notification", "slack", "--name=slack-notification", "--webhook=new-webhook-test", "--force"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(updateCmd)
//				updateCmd.AddCommand(updateSlackNotificationCmd)
//			},
//			expectOut: []string{"slack-notification", "new-webhook-test", "test-channel"},
//			expectErr: false,
//		},
//		{
//			name:    "Delete a Slack Notification from a Project",
//			cmdArgs: []string{"delete", "notification", "project-slack", "--name=slack-notification", "--project=lagoon-demo", "--force"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(deleteCmd)
//				deleteCmd.AddCommand(deleteProjectSlackNotificationCmd)
//			},
//			expectOut: []string{"success"},
//			expectErr: false,
//		},
//		{
//			name:    "Delete a Slack Notification",
//			cmdArgs: []string{"delete", "notification", "slack", "--name=slack-notification", "--force"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(deleteCmd)
//				deleteCmd.AddCommand(deleteSlackNotificationCmd)
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
