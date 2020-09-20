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
	FactsforEnvironment(ctx context.Context, projectId uint, environmentName string, facts *[]schema.Fact) error
	AddFact(ctx context.Context, environmentId int, name string, value string, fact *schema.Fact) error
}

// GetMinimalProjectByNameForFacts gets info of projects in lagoon that have matching metadata.
func GetProjectByNameForFacts(ctx context.Context, projectName string, f Facts) (*schema.Project, error) {
	project := schema.Project{}
	return &project, f.ProjectByName(ctx, projectName, &project)
}

func GetEnvironmentFacts(ctx context.Context, projectId uint, environmentName string, f Facts) (*[]schema.Fact, error) {
	facts := []schema.Fact{}
	ret := f.FactsforEnvironment(ctx, projectId, environmentName, &facts)
	return &facts, ret
}

func AddFact(ctx context.Context, environmentId int, name string, value string, f Facts) (*schema.Fact, error) {
	fact := schema.Fact{}
	err := f.AddFact(ctx, environmentId, name, value, &fact)
	return &fact, err
}
