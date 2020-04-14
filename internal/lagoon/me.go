// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"
	"fmt"

	"github.com/amazeeio/lagoon-cli/internal/schema"
)

// Me interface contains methods for getting info on the current user of lagoon.
type Me interface {
	Me(ctx context.Context, user *schema.User) error
}

// GetMeInfo gets info on the current user of lagoon.
func GetMeInfo(ctx context.Context,
	m Me) (*schema.User, error) {

	user := &schema.User{}

	err := m.Me(ctx, user)
	if err != nil {
		return user, fmt.Errorf("couldn't perform request: %w", err)
	}

	return user, nil
}
