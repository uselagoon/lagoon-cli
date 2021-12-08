package lagoon

import (
	"context"

	"github.com/uselagoon/lagoon-cli/internal/schema"
)

// Workflow interface contains methods for running workflows in lagoon.
type Workflow interface {
	AddAdvancedTaskDefinition(ctx context.Context, input *schema.AdvancedTaskDefinitionInput, taskDefinition *schema.AdvancedTaskDefinition) error
	GetTaskDefinitionByID(ctx context.Context, id int, taskDefinition *schema.AdvancedTaskDefinition) error
	GetAdvancedTasksByEnvironment(ctx context.Context, environment int, tasks *[]schema.AdvancedTaskDefinition) error
	UpdateAdvancedTaskDefinition(ctx context.Context, id int, patch *schema.AdvancedTaskDefinitionInput, taskDefinition *schema.AdvancedTaskDefinition) error
}

// AddAdvancedTaskDefinition Add Advanced Task Definition from yaml file.
func AddAdvancedTaskDefinition(ctx context.Context, input *schema.AdvancedTaskDefinitionInput, w Workflow) (*schema.AdvancedTaskDefinition, error) {
	taskDefinition := schema.AdvancedTaskDefinition{}
	return &taskDefinition, w.AddAdvancedTaskDefinition(ctx, input, &taskDefinition)
}

// GetTaskDefinitionByID returns an advanced task definition by the associated id
func GetTaskDefinitionByID(ctx context.Context, id int, w Workflow) (*schema.AdvancedTaskDefinition, error) {
	taskDefinition := schema.AdvancedTaskDefinition{}
	return &taskDefinition, w.GetTaskDefinitionByID(ctx, id, &taskDefinition)
}

// GetAdvancedTasksByEnvironment returns an advanced task definition by the associated id
func GetAdvancedTasksByEnvironment(ctx context.Context, environment int, w Workflow) (*[]schema.AdvancedTaskDefinition, error) {
	var tasks []schema.AdvancedTaskDefinition
	return &tasks, w.GetAdvancedTasksByEnvironment(ctx, environment, &tasks)
}

// UpdateAdvancedTaskDefinition updates an advanced task definition by ID.
func UpdateAdvancedTaskDefinition(ctx context.Context, id int, patch *schema.AdvancedTaskDefinitionInput, w Workflow) (*schema.AdvancedTaskDefinition, error) {
	taskDefinition := schema.AdvancedTaskDefinition{}
	return &taskDefinition, w.UpdateAdvancedTaskDefinition(ctx, id, patch, &taskDefinition)
}
