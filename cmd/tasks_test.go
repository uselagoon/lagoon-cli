package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/spf13/cobra"
)

func TestTaskCommands(t *testing.T) {
	tests := []struct {
		name              string
		cmdArgs           []string
		setupCmd          func(*cobra.Command, pflag.FlagSet)
		expectOut         []string
		expectErr         bool
		expectedErrString string
	}{
		// TODO: Seed task data & include cli service for envs
		{
			name:    "Get Task",
			cmdArgs: []string{"get", "task-by-id", "--id=1", "--output-json"},
			setupCmd: func(cmd *cobra.Command, flags pflag.FlagSet) {
				cmd.AddCommand(getCmd)
				getCmd.AddCommand(getTaskByID)
				AddGenericFlags(getTaskByID)
			},
			expectOut: []string{""},
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
