package getconfig

import (
	"context"
	"github.com/jpbede/netpalmgo/internal/transport"
	"github.com/jpbede/netpalmgo/models"
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
