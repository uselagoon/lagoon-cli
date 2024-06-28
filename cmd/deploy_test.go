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
//func TestDeployCommands(t *testing.T) {
//	tests := []struct {
//		name      string
//		cmdArgs   []string
//		setupCmd  func(*cobra.Command, pflag.FlagSet)
//		expectOut []string
//		expectErr bool
//	}{
//		{
//			name:    "Deploy Branch",
//			cmdArgs: []string{"deploy", "branch", "--project=lagoon-demo", "--branch=dev", "--force"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(deployCmd)
//				deployCmd.AddCommand(deployBranchCmd)
//			},
//			expectOut: []string{"success"},
//			expectErr: false,
//		},
//		{
//			name:    "Get Deployment",
//			cmdArgs: []string{"get", "deployment", "--project=lagoon-demo", "--environment=main", "--name=lagoon-build-def456"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(getCmd)
//				getCmd.AddCommand(getDeploymentByNameCmd)
//			},
//			expectOut: []string{"lagoon-build-def456", "85e36e3c-a755-11ed-abf6-df28d8a74109"},
//			expectErr: false,
//		},
//		{
//			name:    "Deploy Promote",
//			cmdArgs: []string{"deploy", "promote", "--project=lagoon-demo-org", "--source=pr-15", "--destination=staging", "--force"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(deployCmd)
//				deployCmd.AddCommand(deployPromoteCmd)
//			},
//			expectOut: []string{"success"},
//			expectErr: false,
//		},
//		{
//			name:    "Deploy Latest",
//			cmdArgs: []string{"deploy", "latest", "--project=lagoon-demo", "--environment=dev", "--force"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(deployCmd)
//				deployCmd.AddCommand(deployLatestCmd)
//			},
//			expectOut: []string{"lagoon-build-"},
//			expectErr: false,
//		},
//		{
//			name:    "Deploy Pullrequest",
//			cmdArgs: []string{"deploy", "pullrequest", "--project=lagoon-demo-org", "--title=pr-15", "--number=15", "--baseBranchName=pr-15", "--baseBranchRef=branchRef", "--headBranchName=branchName", "--headBranchRef=headBranchRef", "--force"},
//			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
//				cmd.AddCommand(deployCmd)
//				deployCmd.AddCommand(deployPullrequestCmd)
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
