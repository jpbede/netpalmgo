package setconfig

import (
	"context"
	"go.bnck.me/netpalm/models"
)

// Client represents the functions implemented by this API
type Client interface {
	Set(ctx context.Context, dryRun bool, request models.SetConfigRequest) (*models.Response, error)
	DryRun(ctx context.Context, request models.SetConfigRequest) (*models.Response, error)
}
