package getconfig

import (
	"context"
	"github.com/jpbede/netpalmgo/models"
)

// Get runs a command with a given library on a given device (args)
func (c *client) WithCommand(ctx context.Context, command string, library models.Library, args models.ConnectionArgs) (*models.Response, error) {
	getconfigReq := models.GetConfigRequest{
		Library:        library,
		ConnectionArgs: args,
		Command:        []string{command},
		QueueStrategy:  models.QueueStrategyFIFO,
	}
	return c.WithRequest(ctx, getconfigReq)
}
