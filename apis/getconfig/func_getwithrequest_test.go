package getconfig

import (
	"github.com/jpbede/netpalmgo/models"
	"os"
	"testing"
)

func TestClient_GetWithRequest(t *testing.T) {
	cl := GetClient()
	req := models.GetConfigRequest{
		Library: models.LibraryNetmiko,
		ConnectionArgs: models.ConnectionArgs{
			DeviceType: "vyos",
			Host:       os.Getenv("DEVICE_HOST"),
			Username:   os.Getenv("DEVICE_USER"),
			Password:   os.Getenv("DEVICE_PASSWORD"),
		},
		Command:       "show int",
		QueueStrategy: models.QueueStrategyFIFO,
	}

	resp, err := cl.GetWithRequest(req)
	if err != nil {
		t.Errorf("Got error while running GetWithRequest(): %s", err.Error())
	}

	if resp != nil && resp.Status != models.StatusSuccess {
		t.Error("Task wasn't successful")
	}
}
