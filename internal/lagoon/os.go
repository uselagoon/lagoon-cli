package lagoon

import (
	"context"
	"github.com/uselagoon/lagoon-cli/internal/schema"
)

type Openshifts interface {
	AddOpenshift(ctx context.Context, in *schema.AddOpenshiftInput, out *schema.AddOpenshiftResponse) error
	DeleteOpenshift(ctx context.Context, in *schema.DeleteOpenshiftInput, out *schema.DeleteOpenshiftResponse) error
}

func AddOpenshift(ctx context.Context, in *schema.AddOpenshiftInput, os Openshifts) (*schema.AddOpenshiftResponse, error) {
	response := schema.AddOpenshiftResponse{}
	return &response, os.AddOpenshift(ctx, in, &response)
}

func DeleteOpenshift(ctx context.Context, in *schema.DeleteOpenshiftInput, os Openshifts) (*schema.DeleteOpenshiftResponse, error) {
	response := schema.DeleteOpenshiftResponse{}

	return &response, os.DeleteOpenshift(ctx, in, &response)
}
