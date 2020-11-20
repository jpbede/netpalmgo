package getconfig

import (
	"context"
	"github.com/jpbede/netpalmgo/models"
)

// Client represents the functions implemented by this API
type Client interface {
	WithCommand(ctx context.Context, command string, library models.Library, args models.ConnectionArgs) (*models.Response, error)
	WithRequest(ctx context.Context, request models.GetConfigRequest) (*models.Response, error)
}
