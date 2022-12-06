package lagoon

import (
	"context"

	"github.com/uselagoon/lagoon-cli/internal/schema"
)

type Variables interface {
	AddOrUpdateEnvVariableByName(ctx context.Context, in *schema.EnvVariableByNameInput, envvar *schema.UpdateEnvVarResponse) error
	DeleteEnvVariableByName(ctx context.Context, in *schema.DeleteEnvVariableByNameInput, envvar *schema.DeleteEnvVarResponse) error
	GetEnvVariablesByProjectEnvironmentName(ctx context.Context, in *schema.EnvVariableByProjectEnvironmentNameInput, envvar *[]schema.EnvKeyValue) error
}

func AddOrUpdateEnvVariableByName(ctx context.Context, in *schema.EnvVariableByNameInput, v Variables) (*schema.UpdateEnvVarResponse, error) {
	envvar := schema.UpdateEnvVarResponse{}
	return &envvar, v.AddOrUpdateEnvVariableByName(ctx, in, &envvar)
}

func DeleteEnvVariableByName(ctx context.Context, in *schema.DeleteEnvVariableByNameInput, v Variables) (*schema.DeleteEnvVarResponse, error) {
	envvar := schema.DeleteEnvVarResponse{}
	return &envvar, v.DeleteEnvVariableByName(ctx, in, &envvar)
}

// ListEnvVars gets info of envvars in lagoon.
func GetEnvVariablesByProjectEnvironmentName(ctx context.Context, in *schema.EnvVariableByProjectEnvironmentNameInput, v Variables) (*[]schema.EnvKeyValue, error) {
	envvar := []schema.EnvKeyValue{}
	return &envvar, v.GetEnvVariablesByProjectEnvironmentName(ctx, in, &envvar)
}
