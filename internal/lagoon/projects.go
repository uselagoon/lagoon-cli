// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"

	"github.com/amazeeio/lagoon-cli/internal/schema"
)

// Projects interface contains methods for getting info on projects.
type Projects interface {
	ProjectsByMetadata(ctx context.Context, key string, value string, user *[]schema.ProjectMetadata) error
}

// GetProjectsByMetadata gets info of projects in lagoon that have matching metadata.
func GetProjectsByMetadata(ctx context.Context, key string, value string, p Projects) (*[]schema.ProjectMetadata, error) {
	project := []schema.ProjectMetadata{}
	return &project, p.ProjectsByMetadata(ctx, key, value, &project)
}
