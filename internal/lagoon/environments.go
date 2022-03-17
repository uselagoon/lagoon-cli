// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"

	"github.com/uselagoon/lagoon-cli/internal/schema"
)

// Environments interface contains methods for getting info on environments.
type Environments interface {
	BackupsForEnvironmentByName(context.Context, string, uint, *schema.Environment) error
	AddRestore(context.Context, string, *schema.Restore) error
}

// GetBackupsForEnvironmentByName gets backup info in lagoon for specific environment.
func GetBackupsForEnvironmentByName(ctx context.Context, name string, project uint, e Environments) (*schema.Environment, error) {
	environment := schema.Environment{}
	return &environment, e.BackupsForEnvironmentByName(ctx, name, project, &environment)
}

// AddBackupRestore adds a backup restore based on backup ID.
func AddBackupRestore(ctx context.Context, backupID string, e Environments) (*schema.Restore, error) {
	restore := schema.Restore{}
	return &restore, e.AddRestore(ctx, backupID, &restore)
}
