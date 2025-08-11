package config

import "strings"

// Config is used for the lagoon configuration.
type Config struct {
	Current                  string             `json:"current"`
	Flags                    []string           `json:"flags,omitempty"`
	Default                  string             `json:"default"`
	Lagoons                  map[string]Context `json:"lagoons"`
	UpdateCheckDisable       bool               `json:"updatecheckdisable,omitempty"`
	EnvironmentFromDirectory bool               `json:"environmentfromdirectory,omitempty"`
	StrictHostKeyChecking    string             `json:"stricthostkeychecking,omitempty"`
}

// Context is used for each lagoon context in the config file.
type Context struct {
	GraphQL             string   `json:"graphql"`
	HostName            string   `json:"hostname"`
	UI                  string   `json:"ui,omitempty"`
	Kibana              string   `json:"kibana,omitempty"`
	Port                string   `json:"port"`
	Token               string   `json:"token,omitempty"`
	Version             string   `json:"version,omitempty"`
	SSHKey              string   `json:"sshkey,omitempty"`
	SSHPortal           bool     `json:"sshPortal,omitempty"`
	PublicKeyIdentities []string `json:"publickeyidentities,omitempty"`
}

// IsFlagSet checks if a flag is set in the config.
func (c *Config) IsFlagSet(flag string) bool {
	flagLower := strings.ToLower(flag)
	for _, f := range c.Flags {
		if strings.ToLower(f) == flagLower {
			return true
		}
	}
	return false
}
