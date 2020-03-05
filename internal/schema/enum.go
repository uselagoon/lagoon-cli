package schema

// ProjectAvailability determines the number of pods used to run a project.
type ProjectAvailability string

// High tells Lagoon to load balance across multiple pods.
// Standard tells Lagoon to use a single pod for the site.
const (
	High     ProjectAvailability = "HIGH"
	Standard ProjectAvailability = "STANDARD"
)

// Currency for billing purposes.
type Currency string

// These are the Currency units supported by Lagoon.
const (
	AUD Currency = "AUD"
	EUR Currency = "EUR"
	GBP Currency = "GBP"
	USD Currency = "USD"
	CHF Currency = "CHF"
	ZAR Currency = "ZAR"
)
