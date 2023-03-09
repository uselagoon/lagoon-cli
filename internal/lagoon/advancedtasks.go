package lagoon

import (
	"context"
	"github.com/uselagoon/lagoon-cli/internal/schema"
)

type AdvancedTask interface {
	AddAdvancedTaskDefinition(ctx context.Context, input *schema.AdvancedTaskDefinitionInput, taskDefinition *schema.AdvancedTaskDefinitionResponse) error
	GetTaskDefinitionByID(ctx context.Context, id int, taskDefinition *schema.AdvancedTaskDefinitionResponse) error
	GetAdvancedTasksByEnvironment(ctx context.Context, environment int, tasks *[]schema.AdvancedTaskDefinitionResponse) error
	UpdateAdvancedTaskDefinition(ctx context.Context, id int, patch *schema.AdvancedTaskDefinitionInput, taskDefinition *schema.AdvancedTaskDefinitionResponse) error
}

// AddAdvancedTaskDefinition Add Advanced Task Definition from yaml file.
func AddAdvancedTaskDefinition(ctx context.Context, input *schema.AdvancedTaskDefinitionInput, a AdvancedTask) (*schema.AdvancedTaskDefinitionResponse, error) {
	taskDefinition := schema.AdvancedTaskDefinitionResponse{}
	return &taskDefinition, a.AddAdvancedTaskDefinition(ctx, input, &taskDefinition)
}

// GetTaskDefinitionByID returns an advanced task definition by the associated id
func GetTaskDefinitionByID(ctx context.Context, id int, a AdvancedTask) (*schema.AdvancedTaskDefinitionResponse, error) {
	taskDefinition := schema.AdvancedTaskDefinitionResponse{}
	return &taskDefinition, a.GetTaskDefinitionByID(ctx, id, &taskDefinition)
}

// GetAdvancedTasksByEnvironment returns an advanced task definition by the associated id
func GetAdvancedTasksByEnvironment(ctx context.Context, environment int, a AdvancedTask) (*[]schema.AdvancedTaskDefinitionResponse, error) {
	var tasks []schema.AdvancedTaskDefinitionResponse
	return &tasks, a.GetAdvancedTasksByEnvironment(ctx, environment, &tasks)
}

// UpdateAdvancedTaskDefinition updates an advanced task definition by ID.
func UpdateAdvancedTaskDefinition(ctx context.Context, id int, patch *schema.AdvancedTaskDefinitionInput, a AdvancedTask) (*schema.AdvancedTaskDefinitionResponse, error) {
	taskDefinition := schema.AdvancedTaskDefinitionResponse{}
	return &taskDefinition, a.UpdateAdvancedTaskDefinition(ctx, id, patch, &taskDefinition)
}
