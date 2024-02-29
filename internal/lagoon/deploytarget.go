package lagoon

import (
	"context"

	"github.com/uselagoon/lagoon-cli/internal/schema"
)

type DeployTargets interface {
	AddDeployTarget(ctx context.Context, in *schema.AddDeployTargetInput, out *schema.AddDeployTargetResponse) error
	UpdateDeployTarget(ctx context.Context, in *schema.UpdateDeployTargetInput, out *schema.UpdateDeployTargetResponse) error
	DeleteDeployTarget(ctx context.Context, in *schema.DeleteDeployTargetInput, out *schema.DeleteDeployTargetResponse) error
}

func AddDeployTarget(ctx context.Context, in *schema.AddDeployTargetInput, out DeployTargets) (*schema.AddDeployTargetResponse, error) {
	response := schema.AddDeployTargetResponse{}
	return &response, out.AddDeployTarget(ctx, in, &response)
}

func UpdateDeployTarget(ctx context.Context, in *schema.UpdateDeployTargetInput, out DeployTargets) (*schema.UpdateDeployTargetResponse, error) {
	response := schema.UpdateDeployTargetResponse{}
	return &response, out.UpdateDeployTarget(ctx, in, &response)
}

func DeleteDeployTarget(ctx context.Context, in *schema.DeleteDeployTargetInput, out DeployTargets) (*schema.DeleteDeployTargetResponse, error) {
	response := schema.DeleteDeployTargetResponse{}
	return &response, out.DeleteDeployTarget(ctx, in, &response)
}
