package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListCommands(t *testing.T) {
	tests := []struct {
		name      string
		cmdArgs   []string
		setupCmd  func(*cobra.Command, pflag.FlagSet)
		expectOut []string
		expectErr bool
	}{
		{
			name:    "List Projects",
			cmdArgs: []string{"list", "projects"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(listCmd)
				listCmd.AddCommand(listProjectsCmd)
			},
			expectOut: []string{"lagoon-demo", "lagoon-demo-org"},
			expectErr: false,
		},
		//{
		//	name:    "List Deploy Targets",
		//	cmdArgs: []string{"list", "deploytargets"},
		//	setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
		//		cmd.AddCommand(listCmd)
		//		listCmd.AddCommand(listDeployTargetsCmd)
		//	},
		//	expectOut: []string{"ui-kubernetes", "ci-local-control-k8s"},
		//	expectErr: false,
		//},
		{
			name:    "List Groups",
			cmdArgs: []string{"list", "groups"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(listCmd)
				listCmd.AddCommand(listGroupsCmd)
			},
			expectOut: []string{"ci-group", "lagoon-demo-group", "lagoon-demo-organization-group", "project-lagoon-demo"},
			expectErr: false,
		},
		{
			name:    "List Environments",
			cmdArgs: []string{"list", "environments", "--project=lagoon-demo"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(listCmd)
				listCmd.AddCommand(listEnvironmentsCmd)
			},
			expectOut: []string{"main", "staging", "dev"},
			expectErr: false,
		},
		//{
		//	name:    "List Deployments",
		//	cmdArgs: []string{"list", "deployments", "--project=lagoon-demo", "--environment=main"},
		//	setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
		//		cmd.AddCommand(listCmd)
		//		listCmd.AddCommand(listDeploymentsCmd)
		//	},
		//	expectOut: []string{"lagoon-build-7g8h9i", "lagoon-build-def456", "lagoon-build-123abc"},
		//	expectErr: false,
		//},
		//{
		//	name:    "List Tasks",
		//	cmdArgs: []string{"list", "tasks", "--project=lagoon-demo", "--environment=main"},
		//	setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
		//		cmd.AddCommand(listCmd)
		//		listCmd.AddCommand(listTasksCmd)
		//	},
		//	expectOut: []string{"5b21aff1-689c-41b7-80d7-6de9f5bff1f3", "Developer task"},
		//	expectErr: false,
		//},
		{
			name:    "List Users - all groups",
			cmdArgs: []string{"list", "group-users"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(listCmd)
				listCmd.AddCommand(listUsersCmd)
			},
			expectOut: []string{"ci-customer-user-rsa@example.com", "default-user@lagoon-demo"},
			expectErr: false,
		},
		{
			name:    "List Users",
			cmdArgs: []string{"list", "group-users", "--name=lagoon-demo-group"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(listCmd)
				listCmd.AddCommand(listUsersCmd)
			},
			expectOut: []string{"lagoon-demo-group", "guest@example.com"},
			expectErr: false,
		},
		{
			name:    "List all users",
			cmdArgs: []string{"list", "all-users"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(listCmd)
				listCmd.AddCommand(listAllUsersCmd)
			},
			expectOut: []string{"default-user@lagoon-demo", "ci-customer-user-rsa@example.com", "developer@example.com"},
			expectErr: false,
		},
		//{
		//	name:    "List group-projects",
		//	cmdArgs: []string{"list", "group-projects", "--name=lagoon-demo-group"},
		//	setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
		//		cmd.AddCommand(listCmd)
		//		listCmd.AddCommand(listGroupProjectsCmd)
		//	},
		//	expectOut: []string{"18", "lagoon-demo"},
		//	expectErr: false,
		//},
		//// TODO: Seed variable data
		//{
		//	name:    "List Environment Variables",
		//	cmdArgs: []string{"list", "variables", "--project=lagoon-demo", "--environment=main"},
		//	setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
		//		cmd.AddCommand(listCmd)
		//		listCmd.AddCommand(listVariablesCmd)
		//	},
		//	expectOut: []string{""},
		//	expectErr: false,
		//},
		//// TODO: Seed variable data
		//{
		//	name:    "List Project Variables",
		//	cmdArgs: []string{"list", "variables", "--project=lagoon-demo"},
		//	setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
		//		cmd.AddCommand(listCmd)
		//		listCmd.AddCommand(listVariablesCmd)
		//	},
		//	expectOut: []string{""},
		//	expectErr: false,
		//},
		//{
		//	name:    "List user-groups",
		//	cmdArgs: []string{"list", "user-groups", "--email-address=default-user@lagoon-demo"},
		//	setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
		//		cmd.AddCommand(listCmd)
		//		listCmd.AddCommand(listUsersGroupsCmd)
		//	},
		//	expectOut: []string{"project-lagoon-demo", "MAINTAINER"},
		//	expectErr: false,
		//},
		//{
		//	name:    "List invokable-tasks",
		//	cmdArgs: []string{"list", "invokable-tasks", "--project=lagoon-demo", "--environment=main"},
		//	setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
		//		cmd.AddCommand(listCmd)
		//		listCmd.AddCommand(listInvokableTasks)
		//	},
		//	expectOut: []string{"Maintainer task", "Developer task"},
		//	expectErr: false,
		//},
		//{
		//	name:    "List project-groups",
		//	cmdArgs: []string{"list", "project-groups", "--project=lagoon-demo"},
		//	setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
		//		cmd.AddCommand(listCmd)
		//		listCmd.AddCommand(listProjectGroupsCmd)
		//	},
		//	expectOut: []string{"lagoon-demo-group", "lagoon-group", "project-lagoon-demo"},
		//	expectErr: false,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := rootCmd
			tt.cmdArgs = append(tt.cmdArgs, "--output-json", "--config-file=../temp_config.yaml")
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
