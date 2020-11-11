package getconfig

import (
	"encoding/json"
	"github.com/jpbede/netpalmgo/models"
)

func (c *client) RunWithRequest(request models.GetConfigRequest) (models.Response, error) {
	req := c.transport.R()

	restyResp, err := req.SetBody(request).Post("/getconfig")
	if err != nil {
		return models.Response{}, err
	}

	var resp models.Response
	err = json.Unmarshal(restyResp.Body(), &resp)
	if err != nil {
		return models.Response{}, err
	}
	return resp, nil
}
