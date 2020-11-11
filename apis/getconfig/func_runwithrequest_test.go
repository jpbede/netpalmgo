package getconfig

import (
	"github.com/jpbede/netpalmgo/models"
	"os"
	"testing"
)

func TestClient_RunWithRequest(t *testing.T) {
	cl := GetClient()
	req := models.GetConfigRequest{
		Library: models.LibraryNapalm,
		ConnectionArgs: models.ConnectionArgs{
			DeviceType: "vyos",
			Host:       os.Getenv("DEVICE_HOST"),
			Username:   os.Getenv("DEVICE_USER"),
			Password:   os.Getenv("DEVICE_PASSWORD"),
		},
		Command:       "show int",
		QueueStrategy: models.QueueStrategyFIFO,
	}

	resp, err := cl.RunWithRequest(req)
	if err != nil {
		t.Errorf("Got error while running Run(): %s", err.Error())
	}

	if resp != nil && resp.Status != models.StatusSuccess {
		t.Error("Task wasn't successful")
	}
}
