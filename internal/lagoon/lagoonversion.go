// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"
	"fmt"

	"github.com/amazeeio/lagoon-cli/internal/schema"
)

// APIVersion interface contains methods for getting info on the current version of lagoon.
type APIVersion interface {
	LagoonAPIVersion(ctx context.Context, apiVersion *schema.LagoonVersion) error
	LagoonSchema(ctx context.Context, lagoonSchema *schema.LagoonSchema) error
}

// GetLagoonAPIVersion gets info on the current API version of lagoon, supported in lagoon v1.4.1+
func GetLagoonAPIVersion(ctx context.Context, l APIVersion) (*schema.LagoonVersion, error) {
	apiVersion := schema.LagoonVersion{}
	return &apiVersion, l.LagoonAPIVersion(ctx, &apiVersion)
}

// GetLagoonSchema gets the current schema from lagoon, useful as a backup in determining what versions of lagoon things can support
func GetLagoonSchema(ctx context.Context,
	l APIVersion) (*schema.LagoonSchema, error) {

	lagoonSchema := &schema.LagoonSchema{}

	err := l.LagoonSchema(ctx, lagoonSchema)
	if err != nil {
		return lagoonSchema, fmt.Errorf("couldn't perform request: %w", err)
	}
	return lagoonSchema, nil
}
