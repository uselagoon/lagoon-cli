package schema

type Openshift struct {
	AddOpenshiftInput
	ID int `json:"id,omitempty"`
}

// AddOpenshiftInput is based on the input to addOpenshift.
type AddOpenshiftInput struct {
	Name             string `json:"name"`
	ConsoleUrl       string `json:"consoleUrl"`
	Token            string `json:"token"`
	RouterPattern    string `json:"routerPattern"`
	SshHost          string `json:"sshHost"`
	SshPort          string `json:"sshPort"`
	Created          string `json:"created"`
	MonitoringConfig string `json:"monitoringConfig"`
	FriendlyName     string `json:"friendlyName"`
	CloudProvider    string `json:"cloudProvider"`
	CloudRegion      string `json:"cloudRegion"`
}

type AddOpenshiftResponse struct {
	Openshift
}

type DeleteOpenshiftInput struct {
	Name string `json:"name"`
}

type DeleteOpenshiftResponse struct {
	DeleteOpenshift string `json:"deleteOpenshift"`
}
