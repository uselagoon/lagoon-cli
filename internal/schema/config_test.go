package schema_test

import (
	"bytes"
	"flag"
	"os"
	"testing"

	"github.com/uselagoon/lagoon-cli/internal/schema"
)

var update = flag.Bool("update", false, "update .golden files")

func TestProjectsToConfig(t *testing.T) {
	var testCases = map[string]struct {
		input  string
		expect string
	}{
		"singleProject": {
			input:  "testdata/singleProject.json",
			expect: "testdata/singleProject.golden.yaml",
		},
		"rocketChat": {
			input:  "testdata/rocketChat.json",
			expect: "testdata/rocketChat.golden.yaml",
		},
		"ciBranchPicky": {
			input:  "testdata/ciBranchPicky.json",
			expect: "testdata/ciBranchPicky.golden.yaml",
		},
		"noNewNotifications": {
			input:  "testdata/noNewNotifications.json",
			expect: "testdata/noNewNotifications.golden.yaml",
		},
		"newNotifications": {
			input:  "testdata/newNotifications.json",
			expect: "testdata/newNotifications.golden.yaml",
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(tt *testing.T) {
			// marshal testcase
			testJSON, err := os.ReadFile(tc.input)
			if err != nil {
				tt.Fatalf("couldn't read file: %v", err)
			}
			data := schema.ProjectByNameResponse{}
			if err = schema.UnmarshalProjectByNameResponse(testJSON, &data); err != nil {
				tt.Fatalf("couldn't unmarshal project config: %v", err)
			}
			result, err := schema.ProjectsToConfig(
				[]schema.Project{*data.Response.Project},
				map[string]bool{"project-private-keys": true})
			if err != nil {
				tt.Fatalf("couldn't translate ProjectConfigs: %v", err)
			}

			if *update {
				tt.Logf("update golden file: %s", tc.expect)
				if err = os.WriteFile(tc.expect, result, 0644); err != nil {
					tt.Fatalf("failed to update golden file: %v", err)
				}
			}

			expected, err := os.ReadFile(tc.expect)
			if err != nil {
				tt.Fatalf("failed reading golden file: %v", err)
			}
			if !bytes.Equal(result, expected) {
				tt.Logf("result:\n%s\nexpected:\n%s", result, expected)
				tt.Errorf("result does not match expected")
			}
		})
	}
}
