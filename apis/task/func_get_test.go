package task

import (
	"github.com/go-resty/resty/v2"
	"github.com/jpbede/netpalmgo/apis/getconfig"
	"github.com/jpbede/netpalmgo/models"
	"os"
	"testing"
	"time"
)

func TestClient_GetWithTaskResponse(t *testing.T) {
	r := resty.New().SetHostURL(os.Getenv("NETPALM_APIURL")).SetHeader("x-api-key", os.Getenv("NETPALM_APIKEY"))
	taskClient := &client{
		transport: r,
	}
	getConfigClient := getconfig.New(r)

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
	req.UseTextFSM()

	getconfigResp, err := getConfigClient.GetWithRequest(req)
	if err != nil {
		t.Errorf("Got error while creating task: %s", err.Error())
	}

	time.Sleep(10 * time.Second) // give netpalm some time to process

	_, err = taskClient.GetWithTaskResponse(getconfigResp.Data)
	if err != nil {
		t.Errorf("Got error while running Get(): %s", err.Error())
	}
}
