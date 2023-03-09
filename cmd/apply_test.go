package cmd

import (
	"github.com/uselagoon/lagoon-cli/internal/lagoon/client"
	"github.com/uselagoon/lagoon-cli/internal/schema"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestApplyAdvancedTaskDefinitions(t *testing.T) {
	forceAction = true
	type testInput struct {
		in       []AdvancedTasksFileInput
		response string
	}
	var testCases = map[string]struct {
		input       *testInput
		name        string
		description string
		wantErr     bool
	}{
		"valid-tasks": {
			input: &testInput{
				in: []AdvancedTasksFileInput{
					{
						ID:            1,
						Name:          "Custom task",
						Description:   "A custom advanced task definition",
						Type:          "COMMAND",
						Command:       "$BIN_PATH/lagoon-cli",
						Service:       "cli",
						GroupName:     "project-high-cotton",
						Project:       "high-cotton",
						Environment:   "Master",
						EnvironmentID: 3,
						Permission:    "DEVELOPER",
						Arguments: []schema.AdvancedTaskDefinitionArgument{
							{
								Name: "scan",
								Type: "STRING",
							},
						},
					},
				},
				response: "testdata/addAdvancedTaskResponse.json",
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			testServer := httptest.NewServer(http.HandlerFunc(
				func(w http.ResponseWriter, _ *http.Request) {
					file, err := os.Open(tt.input.response)
					if err != nil {
						t.Fatalf("couldn't open file: %v", err)
					}
					_, err = io.Copy(w, file)
					if err != nil {
						t.Fatalf("couldn't write file contents to HTTP: %v", err)
					}
				}))
			defer testServer.Close()
			lc := client.New(testServer.URL, "", "", "", false)

			if err := ApplyAdvancedTaskDefinitions(lc, tt.input.in); (err != nil) != tt.wantErr {
				t.Errorf("ApplyAdvancedTaskDefinitions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReadConfigFile(t *testing.T) {
	type testInput struct {
		in string
	}
	var testCases = map[string]struct {
		input       *testInput
		name        string
		description string
		want        *FileConfigRoot
		wantErr     bool
	}{
		"valid-tasks-yaml": {
			input: &testInput{
				in: "testdata/tasks.yml",
			},
			wantErr: false,
			want: &FileConfigRoot{
				Tasks: []AdvancedTasksFileInput{
					{
						Name:        "Custom Advanced Task",
						Description: "A custom advanced task defined inside a yml file",
						Project:     "high-cotton",
						Environment: "Master",
						Type:        "COMMAND",
						Command:     "$BIN_PATH/lagoon-cli",
						Service:     "cli",
						Permission:  "DEVELOPER",
					}, {
						Name:        "ClamAV scan",
						Description: "An anti-virus scan ran from an image",
						Project:     "high-cotton",
						Environment: "Master",
						Type:        "IMAGE",
						Image:       "uselagoon/clamav:latest",
						Service:     "cli",
						Permission:  "MAINTAINER",
						Arguments: []schema.AdvancedTaskDefinitionArgument{
							{
								Name: "scan",
								Type: "STRING",
							},
						},
					},
				},
			},
		},
		"invalid-tasks-yaml": {
			input: &testInput{
				in: "testdata/tasks.yml.invalid",
			},
			wantErr: true,
			want:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadConfigFile(tt.input.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reading config file error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reading config file does not match: got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateTasks(t *testing.T) {
	type testInput struct {
		in string
	}
	var testCases = map[string]struct {
		input       *testInput
		name        string
		description string
		want        string
		wantErr     bool
	}{
		"valid-tasks-yaml": {
			input: &testInput{
				in: "testdata/tasks.yml",
			},
			description: "Checks to see if valid config does not return an error",
			wantErr:     false,
			want:        "",
		},
		"missing-tasks-config-yaml": {
			input: &testInput{
				in: "testdata/tasks.yml.missing",
			},
			description: "Checks to see if invalid config throws an exception",
			wantErr:     true,
			want:        "validation checks failed",
		},
	}

	for _, tt := range testCases {
		config, err := ReadConfigFile(tt.input.in)
		if err != nil {
			t.Errorf("Reading config file error = %v", err)
			return
		}

		t.Run(tt.name, func(t *testing.T) {
			_, err = PreprocessAdvancedTaskDefinitionsInputValidation(config.Tasks)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validating config file error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && !reflect.DeepEqual(err.Error(), tt.want) {
				t.Errorf("Validating config file checks failed: got = %v, want %v", err, tt.want)
			}
		})
	}
}
