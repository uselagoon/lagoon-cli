// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"

	"github.com/uselagoon/lagoon-cli/internal/schema"
)

// Me interface contains methods for getting info on the current user of lagoon.
type Me interface {
	Me(ctx context.Context, user *schema.User) error
}

// GetMeInfo gets info on the current user of lagoon.
func GetMeInfo(ctx context.Context, m Me) (*schema.User, error) {
	user := schema.User{}
	return &user, m.Me(ctx, &user)
}
