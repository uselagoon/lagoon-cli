// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"

	"github.com/amazeeio/lagoon-cli/internal/schema"
)
// Deploy interface contains methods for deploying branches and environments in lagoon.
type Facts interface {
	ProjectByName(ctx context.Context, name string, project *schema.Project) error
}

// GetMinimalProjectByNameForFacts gets info of projects in lagoon that have matching metadata.
func GetProjectByNameForFacts(ctx context.Context, projectName string, f Facts) (*schema.Project, error) {
	project := schema.Project{}
	return &project, f.ProjectByName(ctx, projectName, &project)
}