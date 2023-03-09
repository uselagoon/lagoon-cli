// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"

	"github.com/uselagoon/lagoon-cli/internal/schema"
)

// Tasks interface contains methods for running tasks in projects and environments in lagoon.
type Tasks interface {
	RunActiveStandbySwitch(ctx context.Context, project string, result *schema.Task) error
	GetTaskByID(ctx context.Context, id int, result *schema.Task) error
	AddAdvancedTaskDefinition(ctx context.Context, input *schema.AdvancedTaskDefinitionInput, taskDefinition *schema.AdvancedTaskDefinitionResponse) error
	GetTaskDefinitionByID(ctx context.Context, id int, taskDefinition *schema.AdvancedTaskDefinitionResponse) error
	GetAdvancedTasksByEnvironment(ctx context.Context, environment int, tasks *[]schema.AdvancedTaskDefinitionResponse) error
	UpdateAdvancedTaskDefinition(ctx context.Context, id int, patch *schema.AdvancedTaskDefinitionInput, taskDefinition *schema.AdvancedTaskDefinitionResponse) error
}

// ActiveStandbySwitch runs the activestandby switch.
func ActiveStandbySwitch(ctx context.Context, project string, t Tasks) (*schema.Task, error) {
	result := schema.Task{}
	return &result, t.RunActiveStandbySwitch(ctx, project, &result)
}

// TaskByID returns a task by the associated id
func TaskByID(ctx context.Context, id int, t Tasks) (*schema.Task, error) {
	result := schema.Task{}
	return &result, t.GetTaskByID(ctx, id, &result)
}

// AddAdvancedTaskDefinition Add Advanced Task Definition from yaml file.
func AddAdvancedTaskDefinition(ctx context.Context, input *schema.AdvancedTaskDefinitionInput, t Tasks) (*schema.AdvancedTaskDefinitionResponse, error) {
	taskDefinition := schema.AdvancedTaskDefinitionResponse{}
	return &taskDefinition, t.AddAdvancedTaskDefinition(ctx, input, &taskDefinition)
}

// GetTaskDefinitionByID returns an advanced task definition by the associated id
func GetTaskDefinitionByID(ctx context.Context, id int, t Tasks) (*schema.AdvancedTaskDefinitionResponse, error) {
	taskDefinition := schema.AdvancedTaskDefinitionResponse{}
	return &taskDefinition, t.GetTaskDefinitionByID(ctx, id, &taskDefinition)
}

// GetAdvancedTasksByEnvironment returns an advanced task definition by the associated id
func GetAdvancedTasksByEnvironment(ctx context.Context, environment int, t Tasks) (*[]schema.AdvancedTaskDefinitionResponse, error) {
	var tasks []schema.AdvancedTaskDefinitionResponse
	return &tasks, t.GetAdvancedTasksByEnvironment(ctx, environment, &tasks)
}

// UpdateAdvancedTaskDefinition updates an advanced task definition by ID.
func UpdateAdvancedTaskDefinition(ctx context.Context, id int, patch *schema.AdvancedTaskDefinitionInput, t Tasks) (*schema.AdvancedTaskDefinitionResponse, error) {
	taskDefinition := schema.AdvancedTaskDefinitionResponse{}
	return &taskDefinition, t.UpdateAdvancedTaskDefinition(ctx, id, patch, &taskDefinition)
}
