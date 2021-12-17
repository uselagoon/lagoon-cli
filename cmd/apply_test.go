package cmd

import (
	client "github.com/uselagoon/lagoon-cli/internal/lagoon/client"
	"github.com/uselagoon/lagoon-cli/internal/schema"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestApplyAdvancedTaskDefinitions(t *testing.T) {
	forceAction = true
	type testInput struct {
		in       []AdvancedTasksFileInput
		response string
	}
	var testCases = map[string]struct {
		input           *testInput
		name            string
		description     string
		readfailallowed bool
		wantErr         bool
	}{
		"valid-tasks-yaml": {
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

//func TestApplyWorkflows(t *testing.T) {
//	type args struct {
//		lc        *client.Client
//		workflows []WorkflowsFileInput
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := ApplyWorkflows(tt.args.lc, tt.args.workflows); (err != nil) != tt.wantErr {
//				t.Errorf("ApplyWorkflows() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

//func TestReadConfigFile(t *testing.T) {
//	type args struct {
//		fileName string
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    *FileConfigRoot
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := ReadConfigFile(tt.args.fileName)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("ReadConfigFile() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("ReadConfigFile() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_parseFile(t *testing.T) {
//	type args struct {
//		file string
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    *FileConfigRoot
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := parseFile(tt.args.file)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("parseFile() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("parseFile() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
