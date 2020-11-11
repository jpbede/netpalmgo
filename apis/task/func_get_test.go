package task

import (
	"github.com/go-resty/resty/v2"
	"github.com/jpbede/netpalmgo/apis/getconfig"
	"github.com/jpbede/netpalmgo/models"
	"os"
	"testing"
)

func TestClient_GetWithTaskResponse(t *testing.T) {
	r := resty.New().SetHostURL(os.Getenv("NETPALM_APIURL")).SetHeader("x-api-key", os.Getenv("NETPALM_APIKEY"))
	taskClient := &client{
		transport: r,
	}
	getConfigClient := getconfig.New(r)

	args := models.ConnectionArgs{
		DeviceType: "vyos",
		Host:       os.Getenv("DEVICE_HOST"),
		Username:   os.Getenv("DEVICE_USER"),
		Password:   os.Getenv("DEVICE_PASSWORD"),
	}
	getconfigResp, err := getConfigClient.Get("show int", models.LibraryNapalm, args)
	if err != nil {
		t.Errorf("Got error while creating task: %s", err.Error())
	}

	resp, err := taskClient.GetWithTaskResponse(getconfigResp.Data)
	if err != nil {
		t.Errorf("Got error while running Get(): %s", err.Error())
	}

	if resp != nil && resp.Status != models.StatusSuccess {
		t.Error("Task wasn't successful")
	}
}
