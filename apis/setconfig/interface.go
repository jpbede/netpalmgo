package setconfig

import "github.com/jpbede/netpalmgo/models"

// Client represents the functions implemented by this API
type Client interface {
	Set(command string, library models.Library, args models.ConnectionArgs)

	SetWithRequest(request models.GetConfigRequest)
}
