package schema

// ProjectAvailability determines the number of pods used to run a project.
type ProjectAvailability string

// High tells Lagoon to load balance across multiple pods.
// Standard tells Lagoon to use a single pod for the site.
const (
	High     ProjectAvailability = "HIGH"
	Standard ProjectAvailability = "STANDARD"
)
