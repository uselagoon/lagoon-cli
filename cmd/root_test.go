package cmd

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/internal/config"
)

func TestGetLagoonContext(t *testing.T) {
	var testCases = map[string]struct {
		flag   string
		env    map[string]string
		def    string
		expect string
	}{
		"default-from-config": {
			def:    "first",
			expect: "first",
		},
		"no-default": {
			expect: "amazeeio",
		},
		"flag": {
			flag:   "second",
			def:    "first",
			expect: "second",
		},
		"LAGOONCONTEXT": {
			env:    map[string]string{"LAGOONCONTEXT": "second"},
			def:    "first",
			expect: "second",
		},
		"LAGOON_CONTEXT": {
			env:    map[string]string{"LAGOON_CONTEXT": "second"},
			def:    "first",
			expect: "second",
		},
		"LAGOON_CONTEXT-preferred-over-LAGOONCONTEXT": {
			env:    map[string]string{"LAGOONCONTEXT": "third", "LAGOON_CONTEXT": "second"},
			def:    "first",
			expect: "second",
		},
		"flag-preferred-over-envvars": {
			flag:   "second",
			env:    map[string]string{"LAGOONCONTEXT": "third", "LAGOON_CONTEXT": "third"},
			def:    "first",
			expect: "second",
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(tt *testing.T) {
			for _, k := range []string{"LAGOONCONTEXT", "LAGOON_CONTEXT"} {
				if v, ok := tc.env[k]; ok {
					tt.Setenv(k, v)
				} else {
					// t.Setenv restores the original value after the test
					tt.Setenv(k, "")
					os.Unsetenv(k)
				}
			}
			cmd := &cobra.Command{}
			cmd.Flags().String("lagoon", "", "")
			if tc.flag != "" {
				if err := cmd.Flags().Set("lagoon", tc.flag); err != nil {
					tt.Fatal(err)
				}
			}
			originalConfigFilePath := configFilePath
			lc := config.Config{Default: tc.def}
			var lagoon string
			if err := getLagoonContext(&lc, &lagoon, cmd); err != nil {
				tt.Fatal(err)
			}
			if lagoon != tc.expect || lc.Current != tc.expect {
				tt.Fatalf("expected context %q, got %q (current %q)", tc.expect, lagoon, lc.Current)
			}
			if configFilePath != originalConfigFilePath {
				tt.Fatalf("configFilePath changed from %q to %q", originalConfigFilePath, configFilePath)
			}
		})
	}
}
