// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"

	"github.com/uselagoon/lagoon-cli/internal/schema"
)

// Projects interface contains methods for getting info on projects.
type Projects interface {
	MinimalProjectByName(ctx context.Context, name string, project *schema.Project) error
	ProjectByNameMetadata(ctx context.Context, name string, project *schema.ProjectMetadata) error
	SSHEndpointsByProject(ctx context.Context, name string, project *schema.Project) error
}

// GetMinimalProjectByName gets info of projects in lagoon that have matching metadata.
func GetMinimalProjectByName(ctx context.Context, name string, p Projects) (*schema.Project, error) {
	project := schema.Project{}
	return &project, p.MinimalProjectByName(ctx, name, &project)
}

// GetProjectMetadata gets the metadata key:values for a lagoon project.
func GetProjectMetadata(ctx context.Context, name string, p Projects) (*schema.ProjectMetadata, error) {
	project := schema.ProjectMetadata{}
	return &project, p.ProjectByNameMetadata(ctx, name, &project)
}

// GetSSHEndpointsByProject gets info of projects in lagoon that have matching metadata.
func GetSSHEndpointsByProject(ctx context.Context, name string, p Projects) (*schema.Project, error) {
	project := schema.Project{}
	return &project, p.SSHEndpointsByProject(ctx, name, &project)
}
