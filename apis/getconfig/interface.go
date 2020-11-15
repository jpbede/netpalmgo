package getconfig

import "github.com/jpbede/netpalmgo/models"

// Client represents the functions implemented by this API
type Client interface {
	Get(command string, library models.Library, args models.ConnectionArgs) (*models.Response, error)

	GetWithRequest(request models.GetConfigRequest) (*models.Response, error)
}
