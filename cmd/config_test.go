package cmd

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/uselagoon/lagoon-cli/internal/config"
)

func TestConfigRead(t *testing.T) {
	var testCases = map[string]struct {
		input           string
		description     string
		readfailallowed bool
		expect          config.Config
	}{
		"valid-yaml": {
			input:           "testdata/lagoon.yml",
			description:     "This test checks that a valid and complete configuration is parsed",
			readfailallowed: false,
			expect: config.Config{
				Current: "amazeeio",
				Default: "amazeeio",
				Lagoons: map[string]config.Context{
					"amazeeio": {
						GraphQL:  "https://api.amazeeio.cloud/graphql",
						HostName: "token.amazeeio.cloud",
						Kibana:   "https://logs.amazeeio.cloud/",
						UI:       "https://dashboard.amazeeio.cloud",
						Port:     "22",
					},
				},
				UpdateCheckDisable:       false,
				EnvironmentFromDirectory: false,
			},
		},
		"invalid-yaml": {
			input:           "testdata/lagoon.yml.invalid",
			description:     "This test checks to see if an invalid yaml config is not parsed",
			readfailallowed: true,
			expect:          config.Config{},
		},
		"missing-yaml": {
			input:           "testdata/lagoon.yml.missing",
			description:     "This test checks if a context is missing the required data (graphql, hostname, port)",
			readfailallowed: true,
			expect: config.Config{
				Current: "amazeeio",
				Default: "amazeeio",
				Lagoons: map[string]config.Context{
					"amazeeio": {
						Kibana: "https://logs.amazeeio.cloud/",
						UI:     "https://dashboard.amazeeio.cloud",
					},
				},
				UpdateCheckDisable:       false,
				EnvironmentFromDirectory: false,
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(tt *testing.T) {
			lc := config.Config{}
			if err := readLagoonConfig(&lc, tc.input); err != nil {
				if tc.readfailallowed == false {
					tt.Fatal(err)
				}
			}
			fmt.Println(lc, tc.expect)
			if !reflect.DeepEqual(lc, tc.expect) {
				tt.Fatalf("Read config does not match expected config")
			}
		})
	}
}
