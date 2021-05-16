package setconfig

import (
	"context"
	"go.bnck.me/netpalm/internal/transport"
	"go.bnck.me/netpalm/models"
)

func (c *client) Set(ctx context.Context, dryRun bool, request models.SetConfigRequest) (*models.Response, error) {
	var resp *models.Response

	// path for operation
	path := "/setconfig"
	if dryRun {
		path += "/dry-run"
	}

	err := c.transport.Post(ctx, path, &resp, transport.WithJSONRequestBody(request))
	if err != nil {
		return nil, err
	}
	return resp, nil
}
