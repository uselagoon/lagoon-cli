// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"

	"github.com/amazeeio/lagoon-cli/internal/schema"
)

// Openshifts interface contains methods for interacting with openshifts in lagoon.
type Openshifts interface {
	AllOpenshifts(ctx context.Context, result *[]schema.Openshift) error
}

// GetAllOpenshifts deploys the latest environment.
func GetAllOpenshifts(ctx context.Context, m Openshifts) (*[]schema.Openshift, error) {
	result := []schema.Openshift{}
	return &result, m.AllOpenshifts(ctx, &result)
}
