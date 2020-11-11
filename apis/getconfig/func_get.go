package getconfig

import (
	"github.com/jpbede/netpalmgo/models"
)

func (c *client) Get(command string, library models.Library, args models.ConnectionArgs) (*models.Response, error) {
	getconfigReq := models.GetConfigRequest{
		Library:        library,
		ConnectionArgs: args,
		Command:        command,
		QueueStrategy:  models.QueueStrategyFIFO,
	}
	return c.GetWithRequest(getconfigReq)
}
