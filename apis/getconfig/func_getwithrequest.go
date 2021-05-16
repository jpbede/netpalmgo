package getconfig

import (
	"context"
	"go.bnck.me/netpalm/internal/transport"
	"go.bnck.me/netpalm/models"
)

// GetWithRequest run a already created request
func (c *client) WithRequest(ctx context.Context, request models.GetConfigRequest) (*models.Response, error) {
	var resp *models.Response
	err := c.transport.Post(ctx, "/getconfig", &resp, transport.WithJSONRequestBody(request))
	if err != nil {
		return nil, err
	}
	return resp, nil
}
