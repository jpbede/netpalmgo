package setconfig

import "github.com/jpbede/netpalmgo/models"

type Client interface {
	Set(command string, library models.Library, args models.ConnectionArgs)

	SetWithRequest(request models.GetConfigRequest)
}
