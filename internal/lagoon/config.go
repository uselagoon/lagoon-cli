package lagoon

import (
	"golang.org/x/oauth2"
)

// Config is used for the lagoon configuration.
type Config struct {
	Current                  string             `json:"current"`
	Default                  string             `json:"default"`
	Lagoons                  map[string]Context `json:"lagoons"`
	UpdateCheckDisable       bool               `json:"updatecheckdisable,omitempty"`
	EnvironmentFromDirectory bool               `json:"environmentfromdirectory,omitempty"`
}

// Context is used for each lagoon context in the config file.
type Context struct {
	GraphQL     string        `json:"graphql"`
	HostName    string        `json:"hostname"`
	UI          string        `json:"ui,omitempty"`
	Kibana      string        `json:"kibana,omitempty"`
	Port        string        `json:"port"`
	Token       string        `json:"token,omitempty"`
	Grant       *oauth2.Token `json:"grant,omitempty"`
	Version     string        `json:"version,omitempty"`
	SSHKey      string        `json:"sshkey,omitempty"`
	SSHToken    bool          `json:"sshToken,omitempty"`
	KeycloakURL string        `json:"keycloakUrl,omitempty"`
	KeycloakIDP string        `json:"keycloakidp,omitempty"`
}
