package lagoon

/*

     _                               _           _
    | |                             | |         | |
  __| | ___ _ __  _ __ ___  ___ __ _| |_ ___  __| |
 / _` |/ _ \ '_ \| '__/ _ \/ __/ _` | __/ _ \/ _` |
| (_| |  __/ |_) | | |  __/ (_| (_| | ||  __/ (_| |
 \__,_|\___| .__/|_|  \___|\___\__,_|\__\___|\__,_|
           | |
           |_|

this configuration has being replaced with github.com/uselagoon/machinery/utils/config

*/

// Config is used for the lagoon configuration.
type Config struct {
	Current                  string             `json:"current"`
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
