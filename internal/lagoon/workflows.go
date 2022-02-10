package lagoon

import (
	"context"
	"github.com/uselagoon/lagoon-cli/internal/schema"
)

// Workflow interface contains methods for running workflows in lagoon.
type Workflow interface {
	GetWorkflowsByEnvironment(ctx context.Context, environment int, workflows *[]schema.WorkflowResponse) error
	AddWorkflow(ctx context.Context, input *schema.WorkflowInput, taskDefinition *schema.WorkflowResponse) error
	UpdateWorkflow(ctx context.Context, id int, patch *schema.WorkflowInput, taskDefinition *schema.WorkflowResponse) error
	AddAdvancedTaskDefinition(ctx context.Context, input *schema.AdvancedTaskDefinitionInput, taskDefinition *schema.AdvancedTaskDefinitionResponse) error
	GetTaskDefinitionByID(ctx context.Context, id int, taskDefinition *schema.AdvancedTaskDefinitionResponse) error
	GetAdvancedTasksByEnvironment(ctx context.Context, environment int, tasks *[]schema.AdvancedTaskDefinitionResponse) error
	UpdateAdvancedTaskDefinition(ctx context.Context, id int, patch *schema.AdvancedTaskDefinitionInput, taskDefinition *schema.AdvancedTaskDefinitionResponse) error
}

// GetWorkflowsByEnvironment returns workflows by the associated environment id
func GetWorkflowsByEnvironment(ctx context.Context, environment int, w Workflow) (*[]schema.WorkflowResponse, error) {
	var workflows []schema.WorkflowResponse
	return &workflows, w.GetWorkflowsByEnvironment(ctx, environment, &workflows)
}

// AddWorkflow Add Workflow from yaml file.
func AddWorkflow(ctx context.Context, input *schema.WorkflowInput, w Workflow) (*schema.WorkflowResponse, error) {
	workflow := schema.WorkflowResponse{}
	return &workflow, w.AddWorkflow(ctx, input, &workflow)
}

// UpdateWorkflow updates a workflow.
func UpdateWorkflow(ctx context.Context, id int, patch *schema.WorkflowInput, w Workflow) (*schema.WorkflowResponse, error) {
	workflow := schema.WorkflowResponse{}
	return &workflow, w.UpdateWorkflow(ctx, id, patch, &workflow)
}

// AddAdvancedTaskDefinition Add Advanced Task Definition from yaml file.
func AddAdvancedTaskDefinition(ctx context.Context, input *schema.AdvancedTaskDefinitionInput, w Workflow) (*schema.AdvancedTaskDefinitionResponse, error) {
	taskDefinition := schema.AdvancedTaskDefinitionResponse{}
	return &taskDefinition, w.AddAdvancedTaskDefinition(ctx, input, &taskDefinition)
}

// GetTaskDefinitionByID returns an advanced task definition by the associated id
func GetTaskDefinitionByID(ctx context.Context, id int, w Workflow) (*schema.AdvancedTaskDefinitionResponse, error) {
	taskDefinition := schema.AdvancedTaskDefinitionResponse{}
	return &taskDefinition, w.GetTaskDefinitionByID(ctx, id, &taskDefinition)
}

// GetAdvancedTasksByEnvironment returns an advanced task definition by the associated id
func GetAdvancedTasksByEnvironment(ctx context.Context, environment int, w Workflow) (*[]schema.AdvancedTaskDefinitionResponse, error) {
	var tasks []schema.AdvancedTaskDefinitionResponse
	return &tasks, w.GetAdvancedTasksByEnvironment(ctx, environment, &tasks)
}

// UpdateAdvancedTaskDefinition updates an advanced task definition by ID.
func UpdateAdvancedTaskDefinition(ctx context.Context, id int, patch *schema.AdvancedTaskDefinitionInput, w Workflow) (*schema.AdvancedTaskDefinitionResponse, error) {
	taskDefinition := schema.AdvancedTaskDefinitionResponse{}
	return &taskDefinition, w.UpdateAdvancedTaskDefinition(ctx, id, patch, &taskDefinition)
}
