package client_test

import (
	"bytes"
	"context"
	"flag"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/uselagoon/lagoon-cli/internal/lagoon/client"
	"github.com/uselagoon/lagoon-cli/internal/schema"
)

var update = flag.Bool("update", false, "update .golden files")

func TestAddProjectRequest(t *testing.T) {
	var testCases = map[string]struct {
		input  *schema.AddProjectInput
		expect string
	}{
		"AddProjectInput vars can be sent undefined": {
			input:  &schema.AddProjectInput{},
			expect: "testdata/addProjectRequest0.golden.graphql",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(tt *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(
				func(_ http.ResponseWriter, r *http.Request) {
					requestBody, err := io.ReadAll(r.Body)
					if err != nil {
						tt.Fatalf("couldn't read request body: %v", err)
					}

					if *update {
						tt.Logf("update golden file: %s", tc.expect)
						if err = os.WriteFile(tc.expect, requestBody, 0644); err != nil {
							tt.Fatalf("failed to update golden file: %v", err)
						}
					}

					expected, err := os.ReadFile(tc.expect)
					if err != nil {
						tt.Fatalf("couldn't read file: %v", err)
					}
					if !bytes.Equal(requestBody, expected) {
						tt.Logf("result:\n%s\nexpected:\n%s", requestBody, expected)
						tt.Fatalf("result does not match expected")
					}
				}))
			defer ts.Close()
			c := client.New(ts.URL, "", "", "", false)
			// ignore response error - we're testing the request
			_ = c.AddProject(context.Background(), &schema.AddProjectInput{
				Name:                  "foo",
				Openshift:             999,
				GitURL:                "git@example.com/bar/foo.git",
				ProductionEnvironment: "production",
			}, nil)
		})
	}
}

func TestAddOrUpdateEnvironmentResponse(t *testing.T) {
	type testInput struct {
		in       *schema.AddEnvironmentInput
		response string
	}
	var testCases = map[string]struct {
		input  *testInput
		expect *schema.Environment
	}{
		"simple update": {
			input: &testInput{
				in: &schema.AddEnvironmentInput{
					DeployBaseRef:        "develop",
					DeployType:           "BRANCH",
					EnvironmentType:      "DEVELOPMENT",
					Name:                 "develop",
					OpenshiftProjectName: "governors-develop",
					ProjectID:            24,
				},
				response: "testdata/addOrUpdateEnvironmentResponse0.json",
			},
			expect: &schema.Environment{
				AddEnvironmentInput: schema.AddEnvironmentInput{
					ID:   14,
					Name: "develop",
				},
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(tt *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(
				func(w http.ResponseWriter, _ *http.Request) {
					file, err := os.Open(tc.input.response)
					if err != nil {
						tt.Fatalf("couldn't open file: %v", err)
					}
					_, err = io.Copy(w, file)
					if err != nil {
						tt.Fatalf("couldn't write file contents to HTTP: %v", err)
					}
				}))
			defer ts.Close()
			c := client.New(ts.URL, "", "", "", false)

			out := schema.Environment{}
			err := c.AddOrUpdateEnvironment(context.Background(), tc.input.in, &out)
			if err != nil {
				tt.Fatal(err)
			}
			if !reflect.DeepEqual(&out, tc.expect) {
				tt.Fatalf("expected:\n%v\ngot:\n%v", tc.expect, out)
			}
		})
	}
}

func TestAddUser(t *testing.T) {
	type testInput struct {
		in       *schema.AddUserInput
		response string
	}
	userID := uuid.MustParse("b6d33fb6-6b7c-4144-bf90-f7ac6ec47f2e")
	var testCases = map[string]struct {
		input  *testInput
		expect *schema.User
	}{
		"add new user": {
			input: &testInput{
				in: &schema.AddUserInput{
					Email:     "art@vandelayindustries.example.com",
					FirstName: "Art",
					LastName:  "Vandelay",
					Comment:   "Import/Export",
				},
				response: "testdata/addUserResponse0.json",
			},
			expect: &schema.User{
				ID: &userID,
				AddUserInput: schema.AddUserInput{
					Email: "art@vandelayindustries.example.com",
				},
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(tt *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(
				func(w http.ResponseWriter, _ *http.Request) {
					file, err := os.Open(tc.input.response)
					if err != nil {
						tt.Fatalf("couldn't open file: %v", err)
					}
					_, err = io.Copy(w, file)
					if err != nil {
						tt.Fatalf("couldn't write file contents to HTTP: %v", err)
					}
				}))
			defer ts.Close()
			c := client.New(ts.URL, "", "", "", false)

			out := schema.User{}
			err := c.AddUser(context.Background(), tc.input.in, &out)
			if err != nil {
				tt.Fatal(err)
			}
			if !reflect.DeepEqual(&out, tc.expect) {
				tt.Fatalf("expected:\n%v\ngot:\n%v", tc.expect, out)
			}
		})
	}
}
