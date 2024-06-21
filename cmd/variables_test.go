package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/spf13/cobra"
)

func TestVariableCommands(t *testing.T) {
	tests := []struct {
		name              string
		cmdArgs           []string
		setupCmd          func(*cobra.Command, pflag.FlagSet)
		expectOut         []string
		expectErr         bool
		expectedErrString string
	}{
		{
			name:    "Add Variable to project",
			cmdArgs: []string{"add", "variable", "--project=lagoon-demo", "--name=testProjectVariable", "--value=testProjectValue", "--scope=build", "--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(addCmd)
				addCmd.AddCommand(addVariableCmd)
				AddGenericFlags(addVariableCmd)
			},
			expectOut: []string{"testProjectVariable", "testProjectValue", "build"},
			expectErr: false,
		},
		{
			name:    "Add Variable to environment",
			cmdArgs: []string{"add", "variable", "--project=lagoon-demo", "--environment=dev", "--name=testEnvironmentVariable", "--value=testEnvironmentValue", "--scope=runtime", "--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(addCmd)
				addCmd.AddCommand(addVariableCmd)
				AddGenericFlags(addVariableCmd)
			},
			expectOut: []string{"testEnvironmentVariable", "testEnvironmentValue", "runtime"},
			expectErr: false,
		},
		{
			name:    "Delete Variable from project",
			cmdArgs: []string{"delete", "variable", "--project=lagoon-demo", "--name=testProjectVariable", "--output-json", "--force"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(deleteCmd)
				deleteCmd.AddCommand(deleteVariableCmd)
				AddGenericFlags(deleteVariableCmd)
			},
			expectOut: []string{"success"},
			expectErr: false,
		},
		{
			name:    "Delete Variable from environment",
			cmdArgs: []string{"delete", "variable", "--project=lagoon-demo", "--environment=dev", "--name=testEnvironmentVariable", "--output-json", "--force"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(deleteCmd)
				deleteCmd.AddCommand(deleteVariableCmd)
				AddGenericFlags(deleteVariableCmd)
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
