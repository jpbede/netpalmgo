package task

import (
	"encoding/json"
	"github.com/jpbede/netpalmgo/models"
	"github.com/jpbede/netpalmgo/util"
)

func (c *client) GetWithTaskResponse(response models.TaskResponse) (*models.Response, error) {
	restyResp, err := c.transport.R().Get("/task/" + response.TaskID)
	if err != nil {
		return nil, err
	}

	if err := util.CheckForHTTPError(restyResp); err != nil {
		return nil, err
	}

	var resp *models.Response
	if err := json.Unmarshal(restyResp.Body(), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
