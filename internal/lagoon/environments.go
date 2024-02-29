// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"

	"github.com/uselagoon/lagoon-cli/internal/schema"
)

// Environments interface contains methods for getting info on environments.
type Environments interface {
	AddRestore(context.Context, string, *schema.Restore) error
}

// AddBackupRestore adds a backup restore based on backup ID.
func AddBackupRestore(ctx context.Context, backupID string, e Environments) (*schema.Restore, error) {
	restore := schema.Restore{}
	return &restore, e.AddRestore(ctx, backupID, &restore)
}
