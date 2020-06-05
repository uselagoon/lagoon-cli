package schema

// Openshift is based on the Lagoon API type.
type Openshift struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	ConsoleURL    string `json:"consoleUrl"`
	RouterPattern string `json:"routerPattern"`
	ProjectUser   string `json:"projectUser"`
	SSHHost       string `json:"sshHost"`
	SSHPort       string `json:"sshPort"`
	Created       string `json:"created"`
	Token         string `json:"token"`
}
