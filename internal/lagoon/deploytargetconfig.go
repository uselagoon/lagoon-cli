// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"

	"github.com/uselagoon/lagoon-cli/internal/schema"
)

// DeployTargetConfigs interface contains methods for getting info on deploytarget configs.
type DeployTargetConfigs interface {
	DeployTargetConfigsByProjectID(ctx context.Context, project int, deployTargets *[]schema.DeployTargetConfig) error
	UpdateDeployTargetConfiguration(ctx context.Context, in *schema.UpdateDeployTargetConfigInput, deployTargets *schema.DeployTargetConfig) error
	DeleteDeployTargetConfiguration(ctx context.Context, id int, project int, deployTargets *schema.DeleteDeployTargetConfig) error
}

// GetDeployTargetConfigs gets deploytarget configs for a specific project.
func GetDeployTargetConfigs(ctx context.Context, project int, dtc DeployTargetConfigs) (*[]schema.DeployTargetConfig, error) {
	deployTargets := []schema.DeployTargetConfig{}
	return &deployTargets, dtc.DeployTargetConfigsByProjectID(ctx, project, &deployTargets)
}

// UpdateDeployTargetConfiguration adds a deploytarget config to a specific project.
func UpdateDeployTargetConfiguration(ctx context.Context, in *schema.UpdateDeployTargetConfigInput, dtc DeployTargetConfigs) (*schema.DeployTargetConfig, error) {
	deployTarget := schema.DeployTargetConfig{}
	return &deployTarget, dtc.UpdateDeployTargetConfiguration(ctx, in, &deployTarget)
}

// DeleteDeployTargetConfiguration deletes a deploytarget config from a specific project.
func DeleteDeployTargetConfiguration(ctx context.Context, id int, project int, dtc DeployTargetConfigs) (*schema.DeleteDeployTargetConfig, error) {
	deployTarget := schema.DeleteDeployTargetConfig{}
	return &deployTarget, dtc.DeleteDeployTargetConfiguration(ctx, id, project, &deployTarget)
}
