// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"

	"github.com/amazeeio/lagoon-cli/internal/schema"
)

// Me interface contains methods for getting info on the current user of lagoon.
type Me interface {
	Me(ctx context.Context, user *schema.User) error
	CanUserSSH(ctx context.Context, openshiftProjectName string, environment *schema.Environment) error
}

// GetMeInfo gets info on the current user of lagoon.
func GetMeInfo(ctx context.Context, m Me) (*schema.User, error) {
	user := schema.User{}
	return &user, m.Me(ctx, &user)
}

// CanUserSSHToEnvironment returns the environment if the user can ssh to it
func CanUserSSHToEnvironment(ctx context.Context, openshiftProjectName string, m Me) (*schema.Environment, error) {
	environment := schema.Environment{}
	return &environment, m.CanUserSSH(ctx, openshiftProjectName, &environment)
}
