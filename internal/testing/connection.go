package testing

import "github.com/jpbede/netpalmgo/models"

func GetConnectionArgs() models.ConnectionArgs {
	return models.ConnectionArgs{
		DeviceType: "vyos",
		Host:       "10.10.10.10",
		Username:   "demo",
		Password:   "demo",
	}
}
