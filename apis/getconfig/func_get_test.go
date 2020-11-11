package getconfig

import (
	"github.com/go-resty/resty/v2"
	"github.com/jpbede/netpalmgo/models"
	"os"
	"testing"
)

var nplClient *client

func GetClient() *client {
	if nplClient != nil {
		return nplClient
	}

	nplClient = &client{
		transport: resty.New().SetHostURL(os.Getenv("NETPALM_APIURL")).SetHeader("x-api-key", os.Getenv("NETPALM_APIKEY")),
	}

	return nplClient
}

func TestClient_Get(t *testing.T) {
	cl := GetClient()
	args := models.ConnectionArgs{
		DeviceType: "vyos",
		Host:       os.Getenv("DEVICE_HOST"),
		Username:   os.Getenv("DEVICE_USER"),
		Password:   os.Getenv("DEVICE_PASSWORD"),
	}

	resp, err := cl.Get("show int", models.LibraryNetmiko, args)
	if err != nil {
		t.Errorf("Got error while running Get(): %s", err.Error())
	}

	if resp != nil && resp.Status != models.StatusSuccess {
		t.Error("Task wasn't successful")
	}
}
