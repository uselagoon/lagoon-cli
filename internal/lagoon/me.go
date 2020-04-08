// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/amazeeio/lagoon-cli/internal/schema"
)

// Me interface contains methods for getting info on the current user of lagoon.
type Me interface {
	Me(ctx context.Context, user *schema.User) error
}

// GetMeInfo gets info on the current user of lagoon.
func GetMeInfo(ctx context.Context,
	m Me) ([]byte, error) {

	user := schema.User{}

	err := m.Me(ctx, &user)
	if err != nil {
		return nil, fmt.Errorf("couldn't perform request: %w", err)
	}
	userBytes, _ := json.Marshal(user)

	return userBytes, nil
}
