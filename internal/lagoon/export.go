// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"
	"fmt"

	"github.com/amazeeio/lagoon-cli/internal/schema"
)

// Exporter interface contains methods for exporting data from Lagoon.
type Exporter interface {
	ProjectByName(ctx context.Context, name string, project *schema.Project) error
}

// ExportProject exports the given project by name.
func ExportProject(ctx context.Context,
	e Exporter, name string, exclude map[string]bool) ([]byte, error) {

	project := schema.Project{}

	err := e.ProjectByName(ctx, name, &project)
	if err != nil {
		return nil, fmt.Errorf("couldn't perform request: %w", err)
	}

	return schema.ProjectsToConfig([]schema.Project{project}, exclude)
}
