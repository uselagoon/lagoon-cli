package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/spf13/cobra"
)

func TestProjectCommands(t *testing.T) {
	tests := []struct {
		name              string
		cmdArgs           []string
		setupCmd          func(*cobra.Command, pflag.FlagSet)
		expectOut         []string
		expectErr         bool
		expectedErrString string
	}{
		{
			name:    "Add Project",
			cmdArgs: []string{"add", "project", "--project=test-project", "--production-environment=main", "--openshift=4", "--git-url=https://github.com/lagoon-examples/drupal10-base", "--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(addCmd)
				addCmd.AddCommand(addProjectCmd)
				AddGenericFlags(addProjectCmd)
			},
			expectOut: []string{"success", "test-project", "https://github.com/lagoon-examples/drupal10-base"},
			expectErr: false,
		},
		{
			name:    "Add Project to an Organization",
			cmdArgs: []string{"add", "project", "--project=test-organization-project", "--organization-name=lagoon-demo-organization", "--production-environment=main", "--openshift=4", "--git-url=https://github.com/lagoon-examples/drupal10-base", "--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(addCmd)
				addCmd.AddCommand(addProjectCmd)
				AddGenericFlags(addProjectCmd)
			},
			expectOut: []string{"success", "test-organization-project", "https://github.com/lagoon-examples/drupal10-base", "lagoon-demo-organization"},
			expectErr: false,
		},
		{
			name:    "Update a Project",
			cmdArgs: []string{"update", "project", "--project=lagoon-demo", "--auto-idle=0", "--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(updateCmd)
				updateCmd.AddCommand(updateProjectCmd)
				AddGenericFlags(updateProjectCmd)
			},
			expectOut: []string{"success", "lagoon-demo"},
			expectErr: false,
		},
		{
			name:    "Delete a Project",
			cmdArgs: []string{"delete", "project", "--project=test-project", "--output-json", "--force"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(deleteCmd)
				deleteCmd.AddCommand(deleteProjectCmd)
				AddGenericFlags(deleteProjectCmd)
			},
			expectOut: []string{"success"},
			expectErr: false,
		},
		{
			name:    "Remove a Project from an Organization",
			cmdArgs: []string{"delete", "organization-project", "--project=test-organization-project", "--organization-name=lagoon-demo-organization", "--output-json", "--force"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(deleteCmd)
				deleteCmd.AddCommand(removeProjectFromOrganizationCmd)
				AddGenericFlags(removeProjectFromOrganizationCmd)
			},
			expectOut: []string{"success", "test-organization-project", "lagoon-demo-organization"},
			expectErr: false,
		},
		// TODO: Add tests for metadata commands
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
