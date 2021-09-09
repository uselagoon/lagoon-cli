package schema

// Task is based on the Lagoon API type.
type Task struct {
	ID          uint        `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Status      string      `json:"status,omitempty"`
	Created     string      `json:"created,omitempty"`
	Started     string      `json:"started,omitempty"`
	Completed   string      `json:"completed,omitempty"`
	Service     string      `json:"service,omitempty"`
	Command     string      `json:"command,omitempty"`
	RemoteID    string      `json:"remoteId,omitempty"`
	Logs        string      `json:"logs,omitempty"`
	Environment Environment `json:"environment,omitempty"`
}
