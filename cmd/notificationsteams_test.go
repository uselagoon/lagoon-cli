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
//func TestMicrosoftTeamsNotificationCommands(t *testing.T) {
//	tests := []struct {
//		name      string
//		cmdArgs   []string
//		setupCmd  func(*cobra.Command, pflag.FlagSet)
//		expectOut []string
//		expectErr bool
//	}{
//		{
//			name:    "Add MicrosoftTeams Notification",
//			cmdArgs: []string{"add", "notification", "microsoftteams", "--name=microsoftteams-notification", "--webhook=test-webhook", "--force"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(addCmd)
//				addCmd.AddCommand(addNotificationMicrosoftTeamsCmd)
//			},
//			expectOut: []string{"microsoftteams-notification", "test-webhook"},
//			expectErr: false,
//		},
//		{
//			name:    "Add MicrosoftTeams Notification to Project",
//			cmdArgs: []string{"add", "notification", "project-microsoftteams", "--name=microsoftteams-notification", "--project=lagoon-demo", "--force"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(addCmd)
//				addCmd.AddCommand(addProjectNotificationMicrosoftTeamsCmd)
//			},
//			expectOut: []string{"success"},
//			expectErr: false,
//		},
//		{
//			name:    "List Project MicrosoftTeams Notifications",
//			cmdArgs: []string{"list", "notification", "project-microsoftteams", "--project=lagoon-demo"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(listCmd)
//				listCmd.AddCommand(listProjectMicrosoftTeamsCmd)
//			},
//			expectOut: []string{"microsoftteams-notification", "test-webhook"},
//			expectErr: false,
//		},
//		{
//			name:    "List all MicrosoftTeams Notifications",
//			cmdArgs: []string{"list", "notification", "microsoftteams"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(listCmd)
//				listCmd.AddCommand(listAllMicrosoftTeamsCmd)
//			},
//			expectOut: []string{"lagoon-demo", "microsoftteams-notification", "test-webhook"},
//			expectErr: false,
//		},
//		// Unable test newname as incorrect data is returned via the API (fixed in PR#3706)
//		{
//			name:    "Update a MicrosoftTeams Notification",
//			cmdArgs: []string{"update", "notification", "microsoftteams", "--name=microsoftteams-notification", "--webhook=new-webhook-test", "--force"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(updateCmd)
//				updateCmd.AddCommand(updateMicrosoftTeamsNotificationCmd)
//			},
//			expectOut: []string{"microsoftteams-notification", "new-webhook-test"},
//			expectErr: false,
//		},
//		{
//			name:    "Delete a MicrosoftTeams Notification from a Project",
//			cmdArgs: []string{"delete", "notification", "project-microsoftteams", "--name=microsoftteams-notification", "--project=lagoon-demo", "--force"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(deleteCmd)
//				deleteCmd.AddCommand(deleteProjectMicrosoftTeamsNotificationCmd)
//			},
//			expectOut: []string{"success"},
//			expectErr: false,
//		},
//		{
//			name:    "Delete a MicrosoftTeams Notification",
//			cmdArgs: []string{"delete", "notification", "microsoftteams", "--name=microsoftteams-notification", "--force"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(deleteCmd)
//				deleteCmd.AddCommand(deleteMicrosoftTeamsNotificationCmd)
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
