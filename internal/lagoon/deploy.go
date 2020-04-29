// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"

	"github.com/amazeeio/lagoon-cli/internal/schema"
)

// Deploy interface contains methods for deploying branches and environments in lagoon
type Deploy interface {
	DeployEnvironmentLatest(ctx context.Context, deploy *schema.DeployEnvironmentLatestInput, result *schema.DeployEnvironmentLatest) error
	DeployEnvironmentPullrequest(ctx context.Context, deploy *schema.DeployEnvironmentPullrequestInput, result *schema.DeployEnvironmentPullrequest) error
}

// DeployLatest deploys the latest environment
func DeployLatest(ctx context.Context, deploy *schema.DeployEnvironmentLatestInput, m Deploy) (*schema.DeployEnvironmentLatest, error) {
	result := schema.DeployEnvironmentLatest{}
	return &result, m.DeployEnvironmentLatest(ctx, deploy, &result)
}

// DeployPullRequest deploys a pull request
func DeployPullRequest(ctx context.Context, deploy *schema.DeployEnvironmentPullrequestInput, m Deploy) (*schema.DeployEnvironmentPullrequest, error) {
	result := schema.DeployEnvironmentPullrequest{}
	return &result, m.DeployEnvironmentPullrequest(ctx, deploy, &result)
}
