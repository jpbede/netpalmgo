package setconfig

import (
	"context"
	"github.com/jpbede/netpalmgo/models"
)

func (c *client) DryRun(ctx context.Context, request models.SetConfigRequest) (*models.Response, error) {
	return c.Set(ctx, true, request)
}
