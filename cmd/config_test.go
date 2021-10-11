package cmd

import (
	"reflect"
	"testing"

	"github.com/uselagoon/lagoon-cli/internal/lagoon"
)

func TestConfigRead(t *testing.T) {
	var testCases = map[string]struct {
		input           string
		description     string
		readfailallowed bool
		expect          lagoon.Config
	}{
		"valid-yaml": {
			input:           "testdata/lagoon.yml",
			description:     "This test checks that a valid and complete configuration is parsed",
			readfailallowed: false,
			expect: lagoon.Config{
				Current: "amazeeio",
				Default: "amazeeio",
				Lagoons: map[string]lagoon.Context{
					"amazeeio": {
						GraphQL:  "https://api.lagoon.amazeeio.cloud/graphql",
						HostName: "ssh.lagoon.amazeeio.cloud",
						Kibana:   "https://logs-db-ui-lagoon-master.ch.amazee.io/",
						UI:       "https://ui-lagoon-master.ch.amazee.io",
						Port:     "32222",
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
			expect:          lagoon.Config{},
		},
		"missing-yaml": {
			input:           "testdata/lagoon.yml.missing",
			description:     "This test checks if a context is missing the required data (graphql, hostname, port)",
			readfailallowed: true,
			expect: lagoon.Config{
				Current: "amazeeio",
				Default: "amazeeio",
				Lagoons: map[string]lagoon.Context{
					"amazeeio": {
						Kibana: "https://logs-db-ui-lagoon-master.ch.amazee.io/",
						UI:     "https://ui-lagoon-master.ch.amazee.io",
					},
				},
				UpdateCheckDisable:       false,
				EnvironmentFromDirectory: false,
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(tt *testing.T) {
			lc := lagoon.Config{}
			if err := readLagoonConfig(&lc, tc.input); err != nil {
				if tc.readfailallowed == false {
					tt.Fatal(err)
				}
			}
			if !reflect.DeepEqual(lc, tc.expect) {
				tt.Fatalf("Read config does not match expected config")
			}
		})
	}
}
