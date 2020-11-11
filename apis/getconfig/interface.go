package getconfig

import "github.com/jpbede/netpalmgo/models"

type Client interface {
	Run(command string, library models.Library, args models.ConnectionArgs) (models.Response, error)

	RunWithRequest(request models.GetConfigRequest) (models.Response, error)
}
