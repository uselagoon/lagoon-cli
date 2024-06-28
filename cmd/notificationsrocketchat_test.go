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
//func TestRocketChatNotificationCommands(t *testing.T) {
//	tests := []struct {
//		name      string
//		cmdArgs   []string
//		setupCmd  func(*cobra.Command, pflag.FlagSet)
//		expectOut []string
//		expectErr bool
//	}{
//		{
//			name:    "Add RocketChat Notification",
//			cmdArgs: []string{"add", "notification", "rocketchat", "--name=rocketchat-notification", "--channel=test-channel", "--webhook=test-webhook", "--force"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(addCmd)
//				addCmd.AddCommand(addNotificationRocketchatCmd)
//			},
//			expectOut: []string{"rocketchat-notification", "test-channel", "test-webhook"},
//			expectErr: false,
//		},
//		{
//			name:    "Add RocketChat Notification to Project",
//			cmdArgs: []string{"add", "notification", "project-rocketchat", "--name=rocketchat-notification", "--project=lagoon-demo", "--force"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(addCmd)
//				addCmd.AddCommand(addProjectNotificationRocketChatCmd)
//			},
//			expectOut: []string{"success"},
//			expectErr: false,
//		},
//		{
//			name:    "List Project RocketChat Notifications",
//			cmdArgs: []string{"list", "notification", "project-rocketchat", "--project=lagoon-demo"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(listCmd)
//				listCmd.AddCommand(listProjectRocketChatsCmd)
//			},
//			expectOut: []string{"rocketchat-notification", "test-channel", "test-webhook"},
//			expectErr: false,
//		},
//		{
//			name:    "List all RocketChat Notifications",
//			cmdArgs: []string{"list", "notification", "rocketchat"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(listCmd)
//				listCmd.AddCommand(listAllRocketChatsCmd)
//			},
//			expectOut: []string{"lagoon-demo", "rocketchat-notification", "test-channel", "test-webhook"},
//			expectErr: false,
//		},
//		// Unable test newname as incorrect data is returned via the API (fixed in PR#3706)
//		{
//			name:    "Update a RocketChat Notification",
//			cmdArgs: []string{"update", "notification", "rocketchat", "--name=rocketchat-notification", "--webhook=new-webhook-test", "--force"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(updateCmd)
//				updateCmd.AddCommand(updateRocketChatNotificationCmd)
//			},
//			expectOut: []string{"rocketchat-notification", "new-webhook-test", "test-channel"},
//			expectErr: false,
//		},
//		{
//			name:    "Delete a RocketChat Notification from a Project",
//			cmdArgs: []string{"delete", "notification", "project-rocketchat", "--name=rocketchat-notification", "--project=lagoon-demo", "--force"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(deleteCmd)
//				deleteCmd.AddCommand(deleteProjectRocketChatNotificationCmd)
//			},
//			expectOut: []string{"success"},
//			expectErr: false,
//		},
//		{
//			name:    "Delete a RocketChat Notification",
//			cmdArgs: []string{"delete", "notification", "rocketchat", "--name=rocketchat-notification", "--force"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(deleteCmd)
//				deleteCmd.AddCommand(deleteRocketChatNotificationCmd)
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
