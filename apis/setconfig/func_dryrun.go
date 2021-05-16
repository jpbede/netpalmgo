package setconfig

import (
	"context"
	"go.bnck.me/netpalm/models"
)

func (c *client) DryRun(ctx context.Context, request models.SetConfigRequest) (*models.Response, error) {
	return c.Set(ctx, true, request)
}
