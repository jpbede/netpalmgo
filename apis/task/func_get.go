package task

import (
	"context"
	"github.com/jpbede/netpalmgo/models"
)

// WithTaskResponse gets the current infos for a given TaskResponse
func (c *client) WithTaskResponse(ctx context.Context, response models.TaskResponse) (*models.Response, error) {
	var resp *models.Response
	err := c.transport.Get(ctx, "/task/"+response.TaskID, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
