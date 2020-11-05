// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"

	"github.com/amazeeio/lagoon-cli/internal/schema"
)

// Tasks interface contains methods for running tasks in projects and environments in lagoon.
type Tasks interface {
	RunActiveStandbySwitch(ctx context.Context, project string, result *schema.Task) error
	GetTaskByID(ctx context.Context, id int, result *schema.Task) error
}

// ActiveStandbySwitch runs the activestandby switch.
func ActiveStandbySwitch(ctx context.Context, project string, t Tasks) (*schema.Task, error) {
	result := schema.Task{}
	return &result, t.RunActiveStandbySwitch(ctx, project, &result)
}

func TaskByID(ctx context.Context, id int, t Tasks) (*schema.Task, error) {
	result := schema.Task{}
	return &result, t.GetTaskByID(ctx, id, &result)
}
