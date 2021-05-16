package testing

import "go.bnck.me/netpalm/models"

// GetConnectionArgs returns a connection for teesting
func GetConnectionArgs() models.ConnectionArgs {
	return models.ConnectionArgs{
		DeviceType: "vyos",
		Host:       "10.10.10.10",
		Username:   "demo",
		Password:   "demo",
	}
}
