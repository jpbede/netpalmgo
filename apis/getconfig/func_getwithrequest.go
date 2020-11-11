package getconfig

import (
	"encoding/json"
	"github.com/jpbede/netpalmgo/models"
	"github.com/jpbede/netpalmgo/util"
)

func (c *client) GetWithRequest(request models.GetConfigRequest) (*models.Response, error) {
	restyResp, err := c.transport.R().SetBody(request).Post("/getconfig")
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
